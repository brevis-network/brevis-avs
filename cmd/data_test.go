package cmd

import (
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestCalcRequestId(t *testing.T) {
	ec, _ := ethclient.Dial("https://ethereum-holesky-rpc.publicnode.com")

	req := &RequestData{
		ChainId: 17000,
		ReceiptQueryInfos: []*ReceiptQueryInfo{
			{
				TxHash: "0xa84f4b721d7da2b00f78b39e7ef4c992477246c773ccd654fc7b59c9db95585e",
				LogQueryInfos: []*LogQueryInfo{
					{
						LogIndex:         0,
						IsValueFromTopic: false,
						ValueIndex:       2,
					},
				},
			},
		},
		StorageQueryInfos: []*StorageQueryInfo{
			{
				BlkHash: "0xe24e1ab74ada254217daca8dede6147d776cadc911a60ce5e9f0f2ec9e5ca7f2",
				Account: "0x716aa16B01afCCCD4434a8295aA2E845a611cd82",
				Slot:    "0x0000000000000000000000000000000000000000000000000000000000000001",
			},
		},
		TransactionQueryInfos: []*TransactionQueryInfo{
			{
				TxHash: "0xa84f4b721d7da2b00f78b39e7ef4c992477246c773ccd654fc7b59c9db95585e",
			},
		},
	}

	log.Infof("%x", calcRequestId(ec, req))
}
