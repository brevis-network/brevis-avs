/*
Copyright © 2024 Brevis Network
*/
package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	avsdir "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	regco "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

var (
	chid   uint64
	quoStr string // flag hex string will parse into []byte
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "join Brevis AVS",
	Long:  `calls Brevis RegistryCoordinator contract`,
	Run: func(cmd *cobra.Command, args []string) {
		mcc := GetMcc(kMultichain)
		for _, onecfg := range mcc {
			if onecfg.ChainID == chid {
				fmt.Println("joining Brevis AVS on chainid:", chid)
				ec, err := ethclient.Dial(onecfg.Gateway)
				chkErr(err, "eth Dial")
				bgCtx := context.Background()
				chidbig, err := ec.ChainID(bgCtx)
				chkErr(err, "ChainID")
				if chidbig.Uint64() != onecfg.ChainID {
					log.Fatalln("chainid mismatch. got:", chidbig, "expect:", onecfg.ChainID)
				}
				// todo: split into funcs
				// prepare for onchain tx
				regContract, _ := regco.NewContractRegistryCoordinator(Hex2addr(onecfg.BrvRegCo), ec)
				// get ecdsa priv key
				ecdsaKey, err := Kspath2priv(viper.GetString(kKsPath), viper.GetString(kKsPwd))
				chkErr(err, "read ecdsa key")
				// get bls key
				blsKey, err := bls.ReadPrivateKeyFromFile(viper.GetString(kBlsPath), viper.GetString(kBlsPwd))
				chkErr(err, "read bls key")
				// query contract to get hashed addr, to be signed by bls
				g1MsgToSign, err := regContract.PubkeyRegistrationMessageHash(nil, ecdsaKey.Address)
				chkErr(err, "PubkeyRegistrationMessageHash")
				signedMsg := utils.ConvertToBN254G1Point(
					blsKey.SignHashedToCurveMessage(utils.ConvertBn254GethToGnark(g1MsgToSign)).G1Point,
				)
				G1pubkeyBN254 := utils.ConvertToBN254G1Point(blsKey.GetPubKeyG1())
				G2pubkeyBN254 := utils.ConvertToBN254G2Point(blsKey.GetPubKeyG2())
				pubkeyRegParams := regco.IBLSApkRegistryPubkeyRegistrationParams{
					PubkeyRegistrationSignature: signedMsg,
					PubkeyG1:                    G1pubkeyBN254,
					PubkeyG2:                    G2pubkeyBN254,
				}
				// now prepare sigsalt, we need to query avsdirectory contract first
				avs, _ := avsdir.NewContractAVSDirectoryCaller(Hex2addr(onecfg.AVSDir), ec)
				var salt [32]byte
				expiry := big.NewInt(time.Now().Unix() + 1e6)
				expiry.FillBytes(salt[:]) // so salt has same bytes

				msg2sign, err := avs.CalculateOperatorAVSRegistrationDigestHash(nil, ecdsaKey.Address, Hex2addr(onecfg.BrvEigen), salt, expiry)
				chkErr(err, "cal digest")
				sig, err := crypto.Sign(msg2sign[:], ecdsaKey.PrivateKey)
				chkErr(err, "ecdsa sign")
				sig[64] += 27 // 0/1 fix

				// ecdsa auth
				auth, _ := bind.NewKeyedTransactorWithChainID(ecdsaKey.PrivateKey, new(big.Int).SetUint64(chid))
				quoBytes := common.FromHex(quoStr)
				log.Println("calling RegisterOperator onchain, quorums:", quoBytes)
				tx, err := regContract.RegisterOperator(auth, quoBytes, "no need", pubkeyRegParams, regco.ISignatureUtilsSignatureWithSaltAndExpiry{
					Signature: sig,
					Salt:      salt,
					Expiry:    expiry,
				})
				chkErr(err, "RegisterOperator")
				log.Println("waiting for tx:", tx.Hash().String())
				waitCtx, cancel := context.WithTimeout(bgCtx, time.Minute*10)
				defer cancel()
				receipt, err := bind.WaitMined(waitCtx, ec, tx)
				chkErr(err, "waitmine:")
				if receipt.Status != types.ReceiptStatusSuccessful {
					log.Fatalln("tx failed")
				}
				// now query operator id (we could also call sdk types/OperatorIdFromKeyPair locally)
				opId, err := regContract.GetOperatorId(nil, ecdsaKey.Address)
				chkErr(err, "GetOperatorId")
				log.Println("Your OperatorID is:", hex.EncodeToString(opId[:]))
				return
			}
		}
		log.Println("chainid", chid, "not found in config")
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
	joinCmd.PersistentFlags().Uint64Var(&chid, "chainid", 17000, "chain id to join")
	joinCmd.PersistentFlags().StringVar(&quoStr, "quorums", "0x00", "quorums as hex string eg. 0x0001 for quorums 0 and 1")
}
