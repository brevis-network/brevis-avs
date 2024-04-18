package cmd

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/viper"
)

const contractName = "BrevisEigen"

// send sig to brevis gw
type SigReq struct {
	Chainid, Blknum uint64
	// my sig
	Sig bls.Signature
}

// prepare data to be signed by this validator
func msgToSign(chainid uint64, brvEigen string, reqid [32]byte, blkNum uint64) [32]byte {
	packed := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string", "bytes32", "uint64"},
		[]interface{}{new(big.Int).SetUint64(chainid), Hex2addr(brvEigen), contractName, reqid, blkNum},
	)
	return common.BytesToHash(packed)
}

func PostSig(chainid uint64, brvEigen string, reqid [32]byte, blkNum uint64) error {
	msg := msgToSign(chainid, brvEigen, reqid, blkNum)
	// get bls key
	blsKey, err := bls.ReadPrivateKeyFromFile(viper.GetString(kBlsPath), viper.GetString(kBlsPwd))
	if err != nil {
		return fmt.Errorf("read bls key err: %v", err)
	}
	sig := blsKey.SignMessage(msg)
	// post data
	sigreq := SigReq{
		Chainid: chainid,
		Blknum:  blkNum,
		// Reqid:   hex.EncodeToString(reqid[:]),
		// Opid:    viper.GetString(kOpId),
		Sig: *sig,
	}
	raw, _ := json.Marshal(sigreq)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/brveigen/%s/%s", viper.GetString(kBvnGw), hex.EncodeToString(reqid[:]), viper.GetString(kOpId)), bytes.NewBuffer(raw))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("status: %d, msg: %s", resp.StatusCode, string(b))
	}
	return nil
}

// ecdsa keystore json to auth for onchain tx
func Kspath2auth(kspath, pass string, chainid *big.Int) (*bind.TransactOpts, error) {
	ksjson, err := os.ReadFile(kspath)
	if err != nil {
		return nil, err
	}
	return bind.NewTransactorWithChainID(strings.NewReader(string(ksjson)), pass, chainid)
}

func Kspath2priv(kspath, pass string) (*keystore.Key, error) {
	raw, err := os.ReadFile(kspath)
	if err != nil {
		return nil, fmt.Errorf("readfile err: %v", err)
	}
	return keystore.DecryptKey(raw, pass)
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg+":", err)
	}
}
