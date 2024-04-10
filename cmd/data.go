package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type RequestData struct {
	ChainId               uint64                  `json:"chain_id"`
	ReceiptQueryInfos     []*ReceiptQueryInfo     `json:"receipt_query_infos"`
	StorageQueryInfos     []*StorageQueryInfo     `json:"storage_query_infos"`
	TransactionQueryInfos []*TransactionQueryInfo `json:"transaction_query_infos"`
	TargetChainId         uint64                  `json:"target_chain_id"`
}

type ReceiptQueryInfo struct {
	TxHash        string          `json:"tx_hash"`
	LogQueryInfos []*LogQueryInfo `json:"log_query_infos"`
}

type LogQueryInfo struct {
	LogIndex         uint64 `json:"log_index"`
	IsValueFromTopic bool   `json:"is_value_from_topic"`
	ValueIndex       uint64 `json:"value_index"`
}

type StorageQueryInfo struct {
	BlkHash string `json:"blk_hash"`
	Account string `json:"account"`
	Slot    string `json:"slot"`
}

type TransactionQueryInfo struct {
	TxHash string `json:"tx_hash"`
}

type RequestFilledData struct {
	ChainId          uint64
	ReceiptInfos     []*ReceiptInfo
	StorageInfos     []*StorageInfo
	TransactionInfos []*TransactionInfo
}

type ReceiptInfo struct {
	BlkNum       uint64
	ReceiptIndex uint64
	Logs         []*LogInfo
}

type LogInfo struct {
	LogIndex         uint64
	Value            [32]byte
	IsValueFromTopic bool
	ValueIndex       uint64
	ContractAddr     common.Address
	LogTopic0        common.Hash
}

type StorageInfo struct {
	BlkHash   common.Hash
	Account   common.Address
	Slot      common.Hash
	SlotValue [32]byte
	BlkNum    uint64
}

type TransactionInfo struct {
	TxHash          common.Hash
	HashOfRawTxData common.Hash
	BlkHash         common.Hash
	BlkNum          uint64
}

func (data *RequestFilledData) Hash() []byte {
	var hashes [][]byte
	for _, i := range data.ReceiptInfos {
		hashes = append(hashes, i.Hash())
	}
	for _, i := range data.StorageInfos {
		hashes = append(hashes, i.Hash())
	}
	for _, i := range data.TransactionInfos {
		hashes = append(hashes, i.Hash())
	}
	return solsha3.SoliditySHA3(
		[]string{"uint64", "string[]"},
		[]interface{}{data.ChainId, hashes},
	)
}

func (r *ReceiptInfo) Hash() []byte {
	var logs [][]byte
	for _, l := range r.Logs {
		logs = append(logs, l.Hash())
	}
	return solsha3.SoliditySHA3(
		[]string{"uint64", "uint64", "string[]"},
		[]interface{}{r.BlkNum, r.ReceiptIndex, logs},
	)
}

func (l *LogInfo) Hash() []byte {
	return solsha3.SoliditySHA3(
		[]string{"uint64", "bytes32", "bool", "uint64", "address", "bytes32"},
		[]interface{}{l.LogIndex, l.Value, l.IsValueFromTopic, l.ValueIndex, l.ContractAddr, l.LogTopic0},
	)
}

func (s *StorageInfo) Hash() []byte {
	return solsha3.SoliditySHA3(
		[]string{"bytes32", "address", "bytes32", "bytes32", "uint64"},
		[]interface{}{s.BlkHash, s.Account, s.Slot, s.SlotValue, s.BlkNum},
	)
}

func (t *TransactionInfo) Hash() []byte {
	return solsha3.SoliditySHA3(
		[]string{"bytes32", "bytes32", "bytes32", "uint64"},
		[]interface{}{t.TxHash, t.HashOfRawTxData, t.BlkHash, t.BlkNum},
	)
}

func calcRequestId(c *ethclient.Client, req *RequestData) common.Hash {
	var result common.Hash

	data := &RequestFilledData{
		ChainId: req.ChainId,
	}
	for _, q := range req.ReceiptQueryInfos {
		i, err := getReceiptInfo(c, q)
		if err != nil {
			log.Errorf("getReceiptInfo err: %s", err)
			return result
		}
		data.ReceiptInfos = append(data.ReceiptInfos, i)
	}
	for _, q := range req.StorageQueryInfos {
		i, err := getStorageInfo(c, q)
		if err != nil {
			log.Errorf("getStorageInfo err: %s", err)
			return result
		}
		data.StorageInfos = append(data.StorageInfos, i)
	}
	for _, q := range req.TransactionQueryInfos {
		i, err := getTransactionInfo(c, q)
		if err != nil {
			log.Errorf("getTransactionInfo err: %s", err)
			return result
		}
		data.TransactionInfos = append(data.TransactionInfos, i)
	}

	return common.Hash(data.Hash())
}

func getReceiptInfo(c *ethclient.Client, query *ReceiptQueryInfo) (*ReceiptInfo, error) {
	receiptInfo := &ReceiptInfo{}

	receipt, err := c.TransactionReceipt(context.Background(), common.HexToHash(query.TxHash))
	if err != nil {
		return nil, err
	}

	receiptInfo.BlkNum = receipt.BlockNumber.Uint64()
	receiptInfo.ReceiptIndex = uint64(receipt.TransactionIndex)

	var logInfos []*LogInfo
	for _, logQuery := range query.LogQueryInfos {
		for i, log := range receipt.Logs {
			if i == int(logQuery.LogIndex) {
				logInfo := &LogInfo{
					LogIndex:         logQuery.LogIndex,
					IsValueFromTopic: logQuery.IsValueFromTopic,
					ValueIndex:       logQuery.ValueIndex,
				}
				logInfo.ContractAddr = log.Address
				logInfo.LogTopic0 = log.Topics[0]
				if logQuery.IsValueFromTopic {
					logInfo.Value = log.Topics[logQuery.ValueIndex]
				} else {
					copy(logInfo.Value[:], log.Data[logQuery.ValueIndex*32:logQuery.ValueIndex*32+32])
				}

				logInfos = append(logInfos, logInfo)
			}
		}
	}
	receiptInfo.Logs = logInfos

	return receiptInfo, nil
}

func getStorageInfo(c *ethclient.Client, query *StorageQueryInfo) (*StorageInfo, error) {
	header, err := c.HeaderByHash(context.Background(), common.HexToHash(query.BlkHash))
	if err != nil {
		log.Errorf("failed to get block by hash %s: err %s", query.BlkHash, err.Error())
		return nil, err
	}

	value, err := c.StorageAt(context.Background(), common.HexToAddress(query.Account), common.HexToHash(query.Slot), header.Number)
	if err != nil {
		log.Errorf("failed to get proof for account(%s) %s on block %s: %s", query.Account, query.Slot, query.BlkHash, err.Error())
		return nil, err
	}

	storageInfo := &StorageInfo{
		BlkHash: common.HexToHash(query.BlkHash),
		Account: common.HexToAddress(query.Account),
		Slot:    common.HexToHash(query.Slot),
		BlkNum:  header.Number.Uint64(),
	}
	copy(storageInfo.SlotValue[:], value[0:32])

	return storageInfo, nil
}

func getTransactionInfo(c *ethclient.Client, query *TransactionQueryInfo) (*TransactionInfo, error) {
	txInfo := &TransactionInfo{
		TxHash: common.HexToHash(query.TxHash),
	}

	tx, isPending, err := c.TransactionByHash(context.Background(), common.HexToHash(query.TxHash))
	if err != nil {
		log.Errorf("failed to get tx %s, err %s", query.TxHash, err.Error())
		return nil, err
	}
	if isPending {
		log.Errorf("pending transaction not supported: %s", query.TxHash)
		return nil, fmt.Errorf("pending transaction not supported: %s", query.TxHash)
	}

	receipt, err := c.TransactionReceipt(context.Background(), common.HexToHash(query.TxHash))
	if err != nil {
		log.Errorf("failed to get receipt of tx %s, err %s", query.TxHash, err.Error())
		return nil, err
	}

	txInfo.BlkHash = receipt.BlockHash
	txInfo.BlkNum = receipt.BlockNumber.Uint64()
	buf := new(bytes.Buffer)
	tx.EncodeRLP(buf)
	txInfo.HashOfRawTxData = crypto.Keccak256Hash(buf.Bytes())

	return txInfo, nil
}

type PrepareQueryResponse struct {
	QueryHash string `json:"query_hash"`
	Fee       string `json:"fee"`
}

type OpPubKey struct {
	G1X   json.RawMessage `json:"g1_x"`
	G1Y   json.RawMessage `json:"g1_y"`
	G2XA0 json.RawMessage `json:"g2_x_a0"`
	G2XA1 json.RawMessage `json:"g2_x_a1"`
	G2YA0 json.RawMessage `json:"g2_y_a0"`
	G2YA1 json.RawMessage `json:"g2_y_a1"`
}

type OpAddrPubKeyBO struct {
	Addr   common.Address
	PubKey OpPubKey
}
