// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cmd

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// BrevEigenMetaData contains all meta data concerning the BrevEigen contract.
var BrevEigenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"__avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"__registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"__stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mustVerified\",\"inputs\":[{\"name\":\"reqIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumbers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"records\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setQuorums\",\"inputs\":[{\"name\":\"newQ\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyRequest\",\"inputs\":[{\"name\":\"reqid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Verified\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// BrevEigenABI is the input ABI used to generate the binding from.
// Deprecated: Use BrevEigenMetaData.ABI instead.
var BrevEigenABI = BrevEigenMetaData.ABI

// BrevEigen is an auto generated Go binding around an Ethereum contract.
type BrevEigen struct {
	BrevEigenCaller     // Read-only binding to the contract
	BrevEigenTransactor // Write-only binding to the contract
	BrevEigenFilterer   // Log filterer for contract events
}

// BrevEigenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrevEigenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevEigenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrevEigenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevEigenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrevEigenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevEigenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrevEigenSession struct {
	Contract     *BrevEigen        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrevEigenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrevEigenCallerSession struct {
	Contract *BrevEigenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BrevEigenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrevEigenTransactorSession struct {
	Contract     *BrevEigenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BrevEigenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrevEigenRaw struct {
	Contract *BrevEigen // Generic contract binding to access the raw methods on
}

// BrevEigenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrevEigenCallerRaw struct {
	Contract *BrevEigenCaller // Generic read-only contract binding to access the raw methods on
}

// BrevEigenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrevEigenTransactorRaw struct {
	Contract *BrevEigenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrevEigen creates a new instance of BrevEigen, bound to a specific deployed contract.
func NewBrevEigen(address common.Address, backend bind.ContractBackend) (*BrevEigen, error) {
	contract, err := bindBrevEigen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrevEigen{BrevEigenCaller: BrevEigenCaller{contract: contract}, BrevEigenTransactor: BrevEigenTransactor{contract: contract}, BrevEigenFilterer: BrevEigenFilterer{contract: contract}}, nil
}

// NewBrevEigenCaller creates a new read-only instance of BrevEigen, bound to a specific deployed contract.
func NewBrevEigenCaller(address common.Address, caller bind.ContractCaller) (*BrevEigenCaller, error) {
	contract, err := bindBrevEigen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrevEigenCaller{contract: contract}, nil
}

// NewBrevEigenTransactor creates a new write-only instance of BrevEigen, bound to a specific deployed contract.
func NewBrevEigenTransactor(address common.Address, transactor bind.ContractTransactor) (*BrevEigenTransactor, error) {
	contract, err := bindBrevEigen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrevEigenTransactor{contract: contract}, nil
}

// NewBrevEigenFilterer creates a new log filterer instance of BrevEigen, bound to a specific deployed contract.
func NewBrevEigenFilterer(address common.Address, filterer bind.ContractFilterer) (*BrevEigenFilterer, error) {
	contract, err := bindBrevEigen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrevEigenFilterer{contract: contract}, nil
}

// bindBrevEigen binds a generic wrapper to an already deployed contract.
func bindBrevEigen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrevEigenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevEigen *BrevEigenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevEigen.Contract.BrevEigenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevEigen *BrevEigenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevEigen.Contract.BrevEigenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevEigen *BrevEigenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevEigen.Contract.BrevEigenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevEigen *BrevEigenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevEigen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevEigen *BrevEigenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevEigen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevEigen *BrevEigenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevEigen.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BrevEigen *BrevEigenCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BrevEigen *BrevEigenSession) AvsDirectory() (common.Address, error) {
	return _BrevEigen.Contract.AvsDirectory(&_BrevEigen.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BrevEigen *BrevEigenCallerSession) AvsDirectory() (common.Address, error) {
	return _BrevEigen.Contract.AvsDirectory(&_BrevEigen.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BrevEigen *BrevEigenCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BrevEigen *BrevEigenSession) BlsApkRegistry() (common.Address, error) {
	return _BrevEigen.Contract.BlsApkRegistry(&_BrevEigen.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BrevEigen *BrevEigenCallerSession) BlsApkRegistry() (common.Address, error) {
	return _BrevEigen.Contract.BlsApkRegistry(&_BrevEigen.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_BrevEigen *BrevEigenCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_BrevEigen *BrevEigenSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _BrevEigen.Contract.CheckSignatures(&_BrevEigen.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_BrevEigen *BrevEigenCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _BrevEigen.Contract.CheckSignatures(&_BrevEigen.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BrevEigen *BrevEigenCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BrevEigen *BrevEigenSession) Delegation() (common.Address, error) {
	return _BrevEigen.Contract.Delegation(&_BrevEigen.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_BrevEigen *BrevEigenCallerSession) Delegation() (common.Address, error) {
	return _BrevEigen.Contract.Delegation(&_BrevEigen.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BrevEigen *BrevEigenCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BrevEigen *BrevEigenSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _BrevEigen.Contract.GetOperatorRestakedStrategies(&_BrevEigen.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BrevEigen *BrevEigenCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _BrevEigen.Contract.GetOperatorRestakedStrategies(&_BrevEigen.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BrevEigen *BrevEigenCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BrevEigen *BrevEigenSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BrevEigen.Contract.GetRestakeableStrategies(&_BrevEigen.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BrevEigen *BrevEigenCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BrevEigen.Contract.GetRestakeableStrategies(&_BrevEigen.CallOpts)
}

// MustVerified is a free data retrieval call binding the contract method 0xd1e34715.
//
// Solidity: function mustVerified(bytes32[] reqIds) view returns()
func (_BrevEigen *BrevEigenCaller) MustVerified(opts *bind.CallOpts, reqIds [][32]byte) error {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "mustVerified", reqIds)

	if err != nil {
		return err
	}

	return err

}

// MustVerified is a free data retrieval call binding the contract method 0xd1e34715.
//
// Solidity: function mustVerified(bytes32[] reqIds) view returns()
func (_BrevEigen *BrevEigenSession) MustVerified(reqIds [][32]byte) error {
	return _BrevEigen.Contract.MustVerified(&_BrevEigen.CallOpts, reqIds)
}

// MustVerified is a free data retrieval call binding the contract method 0xd1e34715.
//
// Solidity: function mustVerified(bytes32[] reqIds) view returns()
func (_BrevEigen *BrevEigenCallerSession) MustVerified(reqIds [][32]byte) error {
	return _BrevEigen.Contract.MustVerified(&_BrevEigen.CallOpts, reqIds)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevEigen *BrevEigenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevEigen *BrevEigenSession) Owner() (common.Address, error) {
	return _BrevEigen.Contract.Owner(&_BrevEigen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevEigen *BrevEigenCallerSession) Owner() (common.Address, error) {
	return _BrevEigen.Contract.Owner(&_BrevEigen.CallOpts)
}

// QuorumNumbers is a free data retrieval call binding the contract method 0x2a8414fd.
//
// Solidity: function quorumNumbers() view returns(bytes)
func (_BrevEigen *BrevEigenCaller) QuorumNumbers(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "quorumNumbers")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QuorumNumbers is a free data retrieval call binding the contract method 0x2a8414fd.
//
// Solidity: function quorumNumbers() view returns(bytes)
func (_BrevEigen *BrevEigenSession) QuorumNumbers() ([]byte, error) {
	return _BrevEigen.Contract.QuorumNumbers(&_BrevEigen.CallOpts)
}

// QuorumNumbers is a free data retrieval call binding the contract method 0x2a8414fd.
//
// Solidity: function quorumNumbers() view returns(bytes)
func (_BrevEigen *BrevEigenCallerSession) QuorumNumbers() ([]byte, error) {
	return _BrevEigen.Contract.QuorumNumbers(&_BrevEigen.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_BrevEigen *BrevEigenCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "records", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_BrevEigen *BrevEigenSession) Records(arg0 [32]byte) (bool, error) {
	return _BrevEigen.Contract.Records(&_BrevEigen.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(bool)
func (_BrevEigen *BrevEigenCallerSession) Records(arg0 [32]byte) (bool, error) {
	return _BrevEigen.Contract.Records(&_BrevEigen.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BrevEigen *BrevEigenCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BrevEigen *BrevEigenSession) RegistryCoordinator() (common.Address, error) {
	return _BrevEigen.Contract.RegistryCoordinator(&_BrevEigen.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_BrevEigen *BrevEigenCallerSession) RegistryCoordinator() (common.Address, error) {
	return _BrevEigen.Contract.RegistryCoordinator(&_BrevEigen.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BrevEigen *BrevEigenCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BrevEigen *BrevEigenSession) StakeRegistry() (common.Address, error) {
	return _BrevEigen.Contract.StakeRegistry(&_BrevEigen.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BrevEigen *BrevEigenCallerSession) StakeRegistry() (common.Address, error) {
	return _BrevEigen.Contract.StakeRegistry(&_BrevEigen.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_BrevEigen *BrevEigenCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_BrevEigen *BrevEigenSession) StaleStakesForbidden() (bool, error) {
	return _BrevEigen.Contract.StaleStakesForbidden(&_BrevEigen.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_BrevEigen *BrevEigenCallerSession) StaleStakesForbidden() (bool, error) {
	return _BrevEigen.Contract.StaleStakesForbidden(&_BrevEigen.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BrevEigen *BrevEigenCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _BrevEigen.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BrevEigen *BrevEigenSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _BrevEigen.Contract.TrySignatureAndApkVerification(&_BrevEigen.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BrevEigen *BrevEigenCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _BrevEigen.Contract.TrySignatureAndApkVerification(&_BrevEigen.CallOpts, msgHash, apk, apkG2, sigma)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BrevEigen *BrevEigenTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BrevEigen *BrevEigenSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BrevEigen.Contract.DeregisterOperatorFromAVS(&_BrevEigen.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BrevEigen *BrevEigenTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BrevEigen.Contract.DeregisterOperatorFromAVS(&_BrevEigen.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initialOwner) returns()
func (_BrevEigen *BrevEigenTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "initialize", _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initialOwner) returns()
func (_BrevEigen *BrevEigenSession) Initialize(_initialOwner common.Address) (*types.Transaction, error) {
	return _BrevEigen.Contract.Initialize(&_BrevEigen.TransactOpts, _initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initialOwner) returns()
func (_BrevEigen *BrevEigenTransactorSession) Initialize(_initialOwner common.Address) (*types.Transaction, error) {
	return _BrevEigen.Contract.Initialize(&_BrevEigen.TransactOpts, _initialOwner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BrevEigen *BrevEigenTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BrevEigen *BrevEigenSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BrevEigen.Contract.RegisterOperatorToAVS(&_BrevEigen.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BrevEigen *BrevEigenTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BrevEigen.Contract.RegisterOperatorToAVS(&_BrevEigen.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BrevEigen *BrevEigenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BrevEigen *BrevEigenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BrevEigen.Contract.RenounceOwnership(&_BrevEigen.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BrevEigen *BrevEigenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BrevEigen.Contract.RenounceOwnership(&_BrevEigen.TransactOpts)
}

// SetQuorums is a paid mutator transaction binding the contract method 0x266c50a3.
//
// Solidity: function setQuorums(bytes newQ) returns()
func (_BrevEigen *BrevEigenTransactor) SetQuorums(opts *bind.TransactOpts, newQ []byte) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "setQuorums", newQ)
}

// SetQuorums is a paid mutator transaction binding the contract method 0x266c50a3.
//
// Solidity: function setQuorums(bytes newQ) returns()
func (_BrevEigen *BrevEigenSession) SetQuorums(newQ []byte) (*types.Transaction, error) {
	return _BrevEigen.Contract.SetQuorums(&_BrevEigen.TransactOpts, newQ)
}

// SetQuorums is a paid mutator transaction binding the contract method 0x266c50a3.
//
// Solidity: function setQuorums(bytes newQ) returns()
func (_BrevEigen *BrevEigenTransactorSession) SetQuorums(newQ []byte) (*types.Transaction, error) {
	return _BrevEigen.Contract.SetQuorums(&_BrevEigen.TransactOpts, newQ)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_BrevEigen *BrevEigenTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_BrevEigen *BrevEigenSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _BrevEigen.Contract.SetStaleStakesForbidden(&_BrevEigen.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_BrevEigen *BrevEigenTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _BrevEigen.Contract.SetStaleStakesForbidden(&_BrevEigen.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevEigen *BrevEigenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevEigen *BrevEigenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevEigen.Contract.TransferOwnership(&_BrevEigen.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevEigen *BrevEigenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevEigen.Contract.TransferOwnership(&_BrevEigen.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BrevEigen *BrevEigenTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BrevEigen *BrevEigenSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _BrevEigen.Contract.UpdateAVSMetadataURI(&_BrevEigen.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BrevEigen *BrevEigenTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _BrevEigen.Contract.UpdateAVSMetadataURI(&_BrevEigen.TransactOpts, _metadataURI)
}

// VerifyRequest is a paid mutator transaction binding the contract method 0xe88e7868.
//
// Solidity: function verifyRequest(bytes32 reqid, uint64 blockNum, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_BrevEigen *BrevEigenTransactor) VerifyRequest(opts *bind.TransactOpts, reqid [32]byte, blockNum uint64, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _BrevEigen.contract.Transact(opts, "verifyRequest", reqid, blockNum, nonSignerStakesAndSignature)
}

// VerifyRequest is a paid mutator transaction binding the contract method 0xe88e7868.
//
// Solidity: function verifyRequest(bytes32 reqid, uint64 blockNum, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_BrevEigen *BrevEigenSession) VerifyRequest(reqid [32]byte, blockNum uint64, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _BrevEigen.Contract.VerifyRequest(&_BrevEigen.TransactOpts, reqid, blockNum, nonSignerStakesAndSignature)
}

// VerifyRequest is a paid mutator transaction binding the contract method 0xe88e7868.
//
// Solidity: function verifyRequest(bytes32 reqid, uint64 blockNum, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_BrevEigen *BrevEigenTransactorSession) VerifyRequest(reqid [32]byte, blockNum uint64, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _BrevEigen.Contract.VerifyRequest(&_BrevEigen.TransactOpts, reqid, blockNum, nonSignerStakesAndSignature)
}

// BrevEigenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BrevEigen contract.
type BrevEigenInitializedIterator struct {
	Event *BrevEigenInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrevEigenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevEigenInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrevEigenInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrevEigenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevEigenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevEigenInitialized represents a Initialized event raised by the BrevEigen contract.
type BrevEigenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BrevEigen *BrevEigenFilterer) FilterInitialized(opts *bind.FilterOpts) (*BrevEigenInitializedIterator, error) {

	logs, sub, err := _BrevEigen.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BrevEigenInitializedIterator{contract: _BrevEigen.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BrevEigen *BrevEigenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BrevEigenInitialized) (event.Subscription, error) {

	logs, sub, err := _BrevEigen.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevEigenInitialized)
				if err := _BrevEigen.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BrevEigen *BrevEigenFilterer) ParseInitialized(log types.Log) (*BrevEigenInitialized, error) {
	event := new(BrevEigenInitialized)
	if err := _BrevEigen.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevEigenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BrevEigen contract.
type BrevEigenOwnershipTransferredIterator struct {
	Event *BrevEigenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrevEigenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevEigenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrevEigenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrevEigenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevEigenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevEigenOwnershipTransferred represents a OwnershipTransferred event raised by the BrevEigen contract.
type BrevEigenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevEigen *BrevEigenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrevEigenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevEigen.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrevEigenOwnershipTransferredIterator{contract: _BrevEigen.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevEigen *BrevEigenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrevEigenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevEigen.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevEigenOwnershipTransferred)
				if err := _BrevEigen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevEigen *BrevEigenFilterer) ParseOwnershipTransferred(log types.Log) (*BrevEigenOwnershipTransferred, error) {
	event := new(BrevEigenOwnershipTransferred)
	if err := _BrevEigen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevEigenStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the BrevEigen contract.
type BrevEigenStaleStakesForbiddenUpdateIterator struct {
	Event *BrevEigenStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrevEigenStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevEigenStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrevEigenStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrevEigenStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevEigenStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevEigenStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the BrevEigen contract.
type BrevEigenStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_BrevEigen *BrevEigenFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*BrevEigenStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _BrevEigen.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &BrevEigenStaleStakesForbiddenUpdateIterator{contract: _BrevEigen.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_BrevEigen *BrevEigenFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *BrevEigenStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _BrevEigen.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevEigenStaleStakesForbiddenUpdate)
				if err := _BrevEigen.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_BrevEigen *BrevEigenFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*BrevEigenStaleStakesForbiddenUpdate, error) {
	event := new(BrevEigenStaleStakesForbiddenUpdate)
	if err := _BrevEigen.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevEigenVerifiedIterator is returned from FilterVerified and is used to iterate over the raw logs and unpacked data for Verified events raised by the BrevEigen contract.
type BrevEigenVerifiedIterator struct {
	Event *BrevEigenVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrevEigenVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevEigenVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrevEigenVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrevEigenVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevEigenVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevEigenVerified represents a Verified event raised by the BrevEigen contract.
type BrevEigenVerified struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerified is a free log retrieval operation binding the contract event 0xfd80962e9540d35b7e1c5b2289980b16c1b92208bc2a42634e3290ad70fc3b74.
//
// Solidity: event Verified(bytes32 requestId)
func (_BrevEigen *BrevEigenFilterer) FilterVerified(opts *bind.FilterOpts) (*BrevEigenVerifiedIterator, error) {

	logs, sub, err := _BrevEigen.contract.FilterLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return &BrevEigenVerifiedIterator{contract: _BrevEigen.contract, event: "Verified", logs: logs, sub: sub}, nil
}

// WatchVerified is a free log subscription operation binding the contract event 0xfd80962e9540d35b7e1c5b2289980b16c1b92208bc2a42634e3290ad70fc3b74.
//
// Solidity: event Verified(bytes32 requestId)
func (_BrevEigen *BrevEigenFilterer) WatchVerified(opts *bind.WatchOpts, sink chan<- *BrevEigenVerified) (event.Subscription, error) {

	logs, sub, err := _BrevEigen.contract.WatchLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevEigenVerified)
				if err := _BrevEigen.contract.UnpackLog(event, "Verified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerified is a log parse operation binding the contract event 0xfd80962e9540d35b7e1c5b2289980b16c1b92208bc2a42634e3290ad70fc3b74.
//
// Solidity: event Verified(bytes32 requestId)
func (_BrevEigen *BrevEigenFilterer) ParseVerified(log types.Log) (*BrevEigenVerified, error) {
	event := new(BrevEigenVerified)
	if err := _BrevEigen.contract.UnpackLog(event, "Verified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
