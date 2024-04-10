/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	avsregclient "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	listenPort int
	quroNums   types.QuorumNums
	quroPerc   types.QuorumThresholdPercentages
)

var GwServer *Server

// gwCmd represents the gw command
var gwCmd = &cobra.Command{
	Use:   "gw",
	Short: "Run as gateway, accept signed data from validators, aggregate and submit onchain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mcc := GetMcc(kMultichain)
		var eigenC *OneChain
		chainCache := make(map[uint64]*OneChain)
		for i, cfg := range mcc {
			onc, err := NewOneChain(cfg)
			if err != nil {
				log.Println("setup err:", err, "config:", *cfg)
				continue
			}
			if i == 0 {
				eigenC = onc
			}
			chainCache[onc.ChainID] = onc
			go onc.MonBrevisReq()
		}

		if eigenC == nil {
			log.Fatalln("eigen chain not setup")
		}

		eigenec := eigenC.ec
		brvEigen, _ := NewBrevEigen(Hex2addr(eigenC.BrvEigen), eigenec)
		quroNumBytes, err := brvEigen.QuorumNumbers(nil)
		for i := 0; i < len(quroNumBytes); i++ {
			quroNums = append(quroNums, types.QuorumNum(quroNumBytes[i]))
			quroPerc = append(quroPerc, 67) // always 2/3
		}
		log.Println("quroNums:", quroNums, "quroPerc:", quroPerc)
		chkErr(err, "QuorumNumbers")

		dal := NewDB()

		logger, _ := logging.NewZapLogger(logging.Development)
		avsreader, _ := avsregclient.BuildAvsRegistryChainReader(Hex2addr(eigenC.BrvRegCo), Hex2addr(eigenC.OpStateRetriever), eigenec, logger)
		oppub, err := NewOpPkMap(eigenec, Hex2addr(eigenC.BlsApkReg), eigenC.BlsApkStartBlk, eigenC.MaxBlkDelta, dal)
		chkErr(err, "newOpPkMap")
		avsSvc := avsregistry.NewAvsRegistryServiceChainCaller(avsreader, oppub, logger)
		blsSvc := blsagg.NewBlsAggregatorService(avsSvc, logger)

		svr := &Server{
			taskIdx:    1,
			idx2req:    make(map[uint32]ReqInfo),
			reqid2idx:  make(map[[32]byte]uint32),
			blsSvc:     blsSvc,
			config:     eigenC.OneChainConfig,
			chainCache: chainCache,
			dal:        dal,
		}
		GwServer = svr

		// submit when blsSvc has enough sig
		go svr.Forever(brvEigen)
		brvReq, _ := NewBrevisRequest(Hex2addr(eigenC.Brevis), eigenec)
		go svr.DoFulfill(brvReq)

		router := httprouter.New()
		router.POST("/brveigen/:reqid/:opid", svr.HandleSig)
		router.POST("/brvgw/prepareQuery", svr.HandlePrepareQuery)
		router.GET("/brvgw/queryInfo/:reqid/:targetchainid", svr.HandleGetQueryInfo)
		// block forever
		err = http.ListenAndServe(fmt.Sprintf(":%d", listenPort), router)
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(gwCmd)
	gwCmd.PersistentFlags().IntVar(&listenPort, "port", 8080, "listen port for rpc")
	// gwCmd.PersistentFlags().Uint64Var(&chid, "chainid", 17000, "chain id to watch for and submit onchain")
}

type ReqInfo struct {
	Id     [32]byte // reqid
	BlkNum uint64   // blk num
}

type Server struct {
	taskIdx uint32 // incre, must begin >= 1 as 0 is default
	// maps due to blsaggsvc limitation, we know reqid and blknum when handle sig req, but blsaggsvc expects taskIdx(uint32 type)
	// and in BlsAggregationServiceResponse, only taskIdx and msg2sign, no req id, but our contract expects reqid
	idx2req map[uint32]ReqInfo
	// so we know if a reqid is already handled
	reqid2idx map[[32]byte]uint32

	// lock for 2 maps and increment taskIdx
	lock sync.RWMutex

	blsSvc blsagg.BlsAggregationService

	config *OneChainConfig

	chainCache map[uint64]*OneChain

	dal *DB
}

func (s *Server) HandleSig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqidStr := ps.ByName("reqid")
	reqId := common.HexToHash(reqidStr)
	var sigReq SigReq
	err := json.NewDecoder(r.Body).Decode(&sigReq)
	if err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	idx, ok := s.TryGetTaskIdx(reqId)
	if !ok {
		// first time see this req, update things
		idx = s.SetNewReq(reqId, sigReq.Blknum)
		log.Println(quroNums, quroPerc)
		err = s.blsSvc.InitializeNewTask(idx, uint32(sigReq.Blknum), quroNums, quroPerc, time.Hour*12)
		if err != nil {
			log.Println("initNewTask err:", err)
			http.Error(w, "initNewTask: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	opidStr := ps.ByName("opid")
	// process new sig
	msg2sign := msgToSign(s.config.ChainID, s.config.BrvEigen, reqId, sigReq.Blknum)
	err = s.blsSvc.ProcessNewSignature(context.Background(), idx, msg2sign, &sigReq.Sig, types.Bytes32(common.HexToHash(opidStr)))
	if err != nil {
		log.Println("ProcessNewSignature err:", err)
		http.Error(w, "ProcessNewSignature: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func (s *Server) HandlePrepareQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req RequestData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	c, found := s.chainCache[req.ChainId]
	if !found {
		http.Error(w, "bad request: not supported chain", http.StatusBadRequest)
		return
	}
	if len(req.ReceiptQueryInfos)+len(req.StorageQueryInfos)+len(req.TransactionQueryInfos) > 100 {
		http.Error(w, "bad request: exceeds max allowed", http.StatusBadRequest)
		return
	}
	reqId := calcRequestId(c.ec, &req)

	if reqId == *new(common.Hash) {
		http.Error(w, "failed to generate reqId", http.StatusInternalServerError)
		return
	}

	fee := "100000000000000" // hardcode fee for now
	queryInDB, _ := s.dal.SelectQuery(reqId.Hex(), req.TargetChainId)
	if queryInDB == nil {
		queryInDB = NewQuery(reqId.Hex(), req.ChainId, req.TargetChainId, &req, fee)
		err := s.dal.Insert(queryInDB)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
	}

	response := &PrepareQueryResponse{
		QueryHash: reqId.Hex(),
		Fee:       queryInDB.Fee,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

func (s *Server) HandleGetQueryInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqidStr := ps.ByName("reqid")
	reqId := common.HexToHash(reqidStr)
	targetChainIdStr := ps.ByName("targetchainid")
	targetChainId, err := strconv.Atoi(targetChainIdStr)
	if err != nil {
		http.Error(w, "bad request: invalid targetchainid", http.StatusBadRequest)
		return
	}

	queryInDB, _ := s.dal.SelectQuery(reqId.Hex(), uint64(targetChainId))
	if queryInDB == nil {
		http.Error(w, "query not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(queryInDB.RequestData))
}

// block on blsSvc resp, and submit onchain
func (s *Server) Forever(brvEigen *BrevEigen) {
	for blsResp := range s.blsSvc.GetResponseChannel() {
		if blsResp.Err != nil {
			log.Println("blsResp err", blsResp.Err)
			continue
		}
		// TODO: not submit onchain if the query status from DB is not PAID

		// log.Printf("%+v", blsResp)
		nonSignerPubkeys := []BN254G1Point{}
		for _, nonSignerPubkey := range blsResp.NonSignersPubkeysG1 {
			nonSignerPubkeys = append(nonSignerPubkeys, ConvertToBN254G1Point(nonSignerPubkey))
		}
		quorumApks := []BN254G1Point{}
		for _, quorumApk := range blsResp.QuorumApksG1 {
			quorumApks = append(quorumApks, ConvertToBN254G1Point(quorumApk))
		}
		nonSignerStakesAndSignature := IBLSSignatureCheckerNonSignerStakesAndSignature{
			NonSignerPubkeys:             nonSignerPubkeys,
			QuorumApks:                   quorumApks,
			ApkG2:                        ConvertToBN254G2Point(blsResp.SignersApkG2),
			Sigma:                        ConvertToBN254G1Point(blsResp.SignersAggSigG1.G1Point),
			NonSignerQuorumBitmapIndices: blsResp.NonSignerQuorumBitmapIndices,
			QuorumApkIndices:             blsResp.QuorumApkIndices,
			TotalStakeIndices:            blsResp.TotalStakeIndices,
			NonSignerStakeIndices:        blsResp.NonSignerStakeIndices,
		}
		// get reqinfo by taskIdx
		s.lock.RLock()
		reqInfo := s.idx2req[blsResp.TaskIndex]
		s.lock.RUnlock()
		log.Println("sending tx for reqid:", hex.EncodeToString(reqInfo.Id[:]))
		// now submit onchain
		auth, _ := Kspath2auth(viper.GetString(kKsPath), viper.GetString(kKsPwd), new(big.Int).SetUint64(s.config.ChainID))
		abi, _ := BrevEigenMetaData.GetAbi()
		calldata, _ := abi.Pack("verifyRequest", reqInfo.Id, reqInfo.BlkNum, nonSignerStakesAndSignature)
		log.Printf("calldata: %x", calldata)
		tx, err := brvEigen.VerifyRequest(auth, reqInfo.Id, reqInfo.BlkNum, nonSignerStakesAndSignature)
		if err != nil {
			log.Println("verify err:", err)
			continue
		}
		log.Println("verify tx:", tx.Hash().String())
	}
}

func (s *Server) DoFulfill(brvReq *BrevisRequest) {
	auth, err := Kspath2auth(viper.GetString(kKsPath), viper.GetString(kKsPwd), new(big.Int).SetUint64(s.config.ChainID))
	chkErr(err, "Kspath2auth")
	for {
		time.Sleep(1 * time.Minute)

		queries, _ := s.dal.GetToBeFulfilledQueries(48 * time.Hour)
		for _, query := range queries {
			tx, err := brvReq.FulfillOpRequests(auth, [][32]byte{common.HexToHash(query.Hash)}, [][]byte{[]byte("")})
			if err != nil {
				log.Printf("fulfill req %s err %s", query.Hash, err)
				continue
			}
			// TODO, ideally should do below in Mined callback
			query.FulfillTx = tx.Hash().Hex()
			query.Status = QueryStatus_QS_COMPLETE
			s.dal.Update(query)
		}
	}
}

func (s *Server) TryGetTaskIdx(reqid [32]byte) (uint32, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	val, ok := s.reqid2idx[reqid]
	return val, ok
}

// lock, incr taskIdx, update maps, return the new taskIdx
func (s *Server) SetNewReq(reqid [32]byte, blk uint64) uint32 {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.taskIdx += 1
	s.reqid2idx[reqid] = s.taskIdx
	s.idx2req[s.taskIdx] = ReqInfo{
		Id:     reqid,
		BlkNum: blk,
	}
	return s.taskIdx
}

// simple map of operator id -> pubkey, impl required interface GetOperatorPubkeys
type OpPubKeyMap struct {
	m    map[Addr]types.OperatorPubkeys
	lock sync.RWMutex
}

func (om *OpPubKeyMap) GetOperatorPubkeys(_ context.Context, operator Addr) (val types.OperatorPubkeys, ok bool) {
	om.lock.RLock()
	defer om.lock.RUnlock()
	val, ok = om.m[operator]
	return
}

func (om *OpPubKeyMap) AddOpPubkey(operator Addr, pk types.OperatorPubkeys) {
	om.lock.Lock()
	defer om.lock.Unlock()
	om.m[operator] = pk
}

// start and fill the map by iter through logs async, so map takes a bit time to fill
func NewOpPkMap(ec *ethclient.Client, blsapkAddr Addr, origBlk, maxBlkDelta uint64, dal *DB) (*OpPubKeyMap, error) {
	ret := new(OpPubKeyMap)
	ret.m = make(map[Addr]types.OperatorPubkeys)
	// go over events and add
	go func() {
		opPubKeys, err := dal.GetAllOpPubKeys()
		chkErr(err, "dal.GetAllOpPubKeys")
		for _, pub := range opPubKeys {
			sdkPubKey := *convertMyToSDK(&pub.PubKey)
			log.Println("add operator:", pub.Addr.Hex(), ", ", common.Hash(types.OperatorIdFromPubkey(sdkPubKey.G1Pubkey)).Hex())
			ret.AddOpPubkey(pub.Addr, sdkPubKey)
		}
		startBlkFromDB, err := dal.GetBlsApkStartBlk()
		chkErr(err, "dal.GetBlsApkStartBlk")

		logfilter, _ := NewBLSApkRegFilterer(blsapkAddr, ec)
		startBlk := origBlk
		if startBlkFromDB != 0 {
			startBlk = startBlkFromDB
		}
		for {
			endBlk, _ := ec.BlockNumber(context.Background())
			if endBlk < startBlk+10 { // don't query too frequent, wait a minute and check again, holesky 12s/block
				time.Sleep(time.Minute * 2)
				continue
			}
			// catch up old events
			if endBlk > startBlk+maxBlkDelta {
				endBlk = startBlk + maxBlkDelta // avoid error like exceed maximum block range: 50000
			}
			iter, err := logfilter.FilterNewPubkeyRegistration(&bind.FilterOpts{
				Start: startBlk,
				End:   &endBlk,
			}, nil)
			chkErr(err, "filter pubkey reg")
			for iter.Next() {
				ev := iter.Event
				pubkey := types.OperatorPubkeys{
					G1Pubkey: bls.NewG1Point(ev.PubkeyG1.X, ev.PubkeyG1.Y),
					G2Pubkey: bls.NewG2Point(ev.PubkeyG2.X, ev.PubkeyG2.Y),
				}
				log.Println("add operator:", ev.Operator.Hex(), ", ", common.Hash(types.OperatorIdFromPubkey(pubkey.G1Pubkey)).Hex())
				ret.AddOpPubkey(ev.Operator, pubkey)
				err := dal.SaveOpPubKey(&OpAddrPubKeyBO{
					Addr:   ev.Operator,
					PubKey: *convertSDKToMy(&pubkey),
				})
				chkErr(err, "dal.SaveOpPubKey")
			}
			startBlk = endBlk // next range of blocks
			dal.UpdateBlsApkStartBlk(endBlk)
		}
	}()
	return ret, nil
}

func convertMyToSDK(pubkey *OpPubKey) *types.OperatorPubkeys {
	var ret types.OperatorPubkeys
	ret.G1Pubkey = &bls.G1Point{&bn254.G1Affine{}}
	ret.G2Pubkey = &bls.G2Point{&bn254.G2Affine{}}
	(&ret.G1Pubkey.X).UnmarshalJSON(pubkey.G1X)
	(&ret.G1Pubkey.Y).UnmarshalJSON(pubkey.G1Y)
	(&ret.G2Pubkey.X.A0).UnmarshalJSON(pubkey.G2XA0)
	(&ret.G2Pubkey.X.A1).UnmarshalJSON(pubkey.G2XA1)
	(&ret.G2Pubkey.Y.A0).UnmarshalJSON(pubkey.G2YA0)
	(&ret.G2Pubkey.Y.A1).UnmarshalJSON(pubkey.G2YA1)
	return &ret
}

func convertSDKToMy(sdkPubKey *types.OperatorPubkeys) *OpPubKey {
	var ret OpPubKey
	ret.G1X, _ = sdkPubKey.G1Pubkey.X.MarshalJSON()
	ret.G1Y, _ = sdkPubKey.G1Pubkey.Y.MarshalJSON()
	ret.G2XA0, _ = sdkPubKey.G2Pubkey.X.A0.MarshalJSON()
	ret.G2XA1, _ = sdkPubKey.G2Pubkey.X.A1.MarshalJSON()
	ret.G2YA0, _ = sdkPubKey.G2Pubkey.Y.A0.MarshalJSON()
	ret.G2YA1, _ = sdkPubKey.G2Pubkey.Y.A1.MarshalJSON()
	return &ret
}

func ConvertToBN254G1Point(input *bls.G1Point) BN254G1Point {
	output := BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) BN254G2Point {
	output := BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
