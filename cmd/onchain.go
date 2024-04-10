package cmd

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

const defaultPollSec = 30

// monitor events onchain
type OneChain struct {
	*OneChainConfig
	ec  *ethclient.Client
	mon *mon2.Monitor
	// db  *mon2.DAL // persist mon block, in mem only for now
	brevis *BrevisRequestFilterer // brevis contract binding to decode event, set in MonBrevisReq
}

func NewOneChain(cfg *OneChainConfig) (*OneChain, error) {
	ret := &OneChain{
		OneChainConfig: cfg,
	}
	var err error
	ret.ec, err = ethclient.Dial(cfg.Gateway)
	if err != nil {
		return nil, err
	}
	bgCtx := context.Background()
	chid, _ := ret.ec.ChainID(bgCtx)
	if chid == nil {
		return nil, fmt.Errorf("failed to retrieve rpc chid, cfg: %d", cfg.ChainID)
	}
	if chid.Uint64() != cfg.ChainID {
		return nil, fmt.Errorf("mismatch chainid cfg: %d, rpc: %d", cfg.ChainID, chid.Uint64())
	}
	ret.mon, err = mon2.NewMonitor(ret.ec, make(MemDAL), mon2.PerChainCfg{
		BlkIntv:         time.Duration(cfg.BlkInterval) * time.Second,
		BlkDelay:        cfg.BlkDelay,
		MaxBlkDelta:     cfg.MaxBlkDelta,
		ForwardBlkDelay: cfg.ForwardBlkDelay,
	})
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *OneChain) Close() {
	c.mon.Close()
}

var zeroAddr common.Address

// watch new requests on Brevis contract
func (c *OneChain) MonBrevisReq() {
	brvAddr := Hex2addr(c.Brevis)
	if brvAddr == zeroAddr {
		return
	}
	c.brevis, _ = NewBrevisRequestFilterer(brvAddr, c.ec)
	go c.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    brvAddr,
		ChkIntv: time.Second * defaultPollSec,
		AbiStr:  BrevisRequestABI,
		// FromBlk not needed as it's likely already expired anyway
	}, c.evCallback)
}

func (c *OneChain) evCallback(evname string, elog types.Log) {
	switch evname {
	case "RequestSent":
		evSent, err := c.brevis.ParseRequestSent(elog)
		if err != nil {
			log.Error("parse sent err:", err, elog)
		} else {
			log.Infoln("Found req:", hex.EncodeToString(evSent.RequestId[:]), "from:", evSent.Sender.String())
			if GwServer == nil /*opr mode*/ {
				myCalcResult := selfCalcRequestId(evSent.RequestId, c.ChainID)
				if !bytes.Equal(evSent.RequestId[:], myCalcResult[:]) {
					log.Errorf("invalid request %x", evSent.RequestId)
				} else {
					err = PostSig(c.ChainID, c.BrvEigen, evSent.RequestId, elog.BlockNumber)
					if err != nil {
						log.Errorln("post sig err:", err)
					}
				}
			} else {
				query, _ := GwServer.dal.SelectQuery("0x"+common.Bytes2Hex(evSent.RequestId[:]), c.ChainID)
				if query == nil {
					log.Errorf("failed to find query from DB, reqId %x, %d", evSent.RequestId, c.ChainID)
					return
				}

				requiredFee, _ := new(big.Int).SetString(query.Fee, 0)
				if evSent.Fee.Cmp(requiredFee) >= 0 {
					query.Status = QueryStatus_QS_PAID
					_, err := GwServer.dal.Update(query)
					if err != nil {
						log.Errorf("failed to update query status, err %x", err)
					}
				}
			}
		}
	default:
		// skip unsupported ev
		// log.Error("unsupported evname: ", evname)
		return
	}
}

func selfCalcRequestId(claimedReqId [32]byte, targetChainId uint64) [32]byte {
	req := new(RequestData)
	var myCalcResult [32]byte

	httpReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/brvgw/queryInfo/%s/%d", viper.GetString(kBvnGw), hex.EncodeToString(claimedReqId[:]), targetChainId), nil)
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Errorf("failed to retrieve original query info, err: %s", err)
		return myCalcResult
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(req)
	if err != nil {
		log.Errorf("failed to decode original query info, err: %s", err)
		return myCalcResult
	}

	c, found := ChainCache[req.ChainId]
	if !found {
		log.Errorf("unsupported chain %d", req.ChainId)
		return myCalcResult
	}

	return calcRequestId(c.ec, req)
}

type Addr = common.Address

func Hex2addr(addr string) Addr {
	return common.HexToAddress(addr)
}

type MemDAL map[string]mon2.LogEventID

func (d MemDAL) GetMonitorBlock(key string) (uint64, int64, bool, error) {
	le, ok := d[key]
	return le.BlkNum, le.Index, ok, nil
}

func (d MemDAL) SetMonitorBlock(key string, blockNum uint64, blockIdx int64) error {
	d[key] = mon2.LogEventID{BlkNum: blockNum, Index: blockIdx}
	return nil
}
