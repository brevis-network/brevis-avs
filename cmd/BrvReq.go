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
	_ = abi.ConvertType
)

// BrevisExtractInfos is an auto generated low-level Go binding around an user-defined struct.
type BrevisExtractInfos struct {
	Receipts []BrevisReceiptInfo
	Stores   []BrevisStorageInfo
	Txs      []BrevisTransactionInfo
}

// BrevisLogInfo is an auto generated low-level Go binding around an user-defined struct.
type BrevisLogInfo struct {
	LogIndex        uint64
	Value           [32]byte
	ValueFromTopic  bool
	ValueIndex      uint64
	ContractAddress common.Address
	LogTopic0       [32]byte
}

// BrevisReceiptInfo is an auto generated low-level Go binding around an user-defined struct.
type BrevisReceiptInfo struct {
	BlkNum       uint64
	ReceiptIndex uint64
	Logs         []BrevisLogInfo
}

// BrevisStorageInfo is an auto generated low-level Go binding around an user-defined struct.
type BrevisStorageInfo struct {
	BlockHash   [32]byte
	Account     common.Address
	Slot        [32]byte
	SlotValue   [32]byte
	BlockNumber uint64
}

// BrevisTransactionInfo is an auto generated low-level Go binding around an user-defined struct.
type BrevisTransactionInfo struct {
	TxHash          [32]byte
	HashOfRawTxData [32]byte
	BlockHash       [32]byte
	BlockNumber     uint64
}

// TxTxInfo is an auto generated low-level Go binding around an user-defined struct.
type TxTxInfo struct {
	ChainId   uint64
	Nonce     uint64
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Gas       *big.Int
	To        common.Address
	Value     *big.Int
	Data      []byte
	From      common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea264697066735822122084316ae7f90ab02dedc5c08bd3c2dea313dd1e29a325b4271a7f1a5bcf89953864736f6c63430008140033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BrevisMetaData contains all meta data concerning the Brevis contract.
var BrevisMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea264697066735822122092c5f428ae7a6eefcc4051ab0da62adfb986ac9461eb63dac4f97a3eeb32f1bb64736f6c63430008140033",
}

// BrevisABI is the input ABI used to generate the binding from.
// Deprecated: Use BrevisMetaData.ABI instead.
var BrevisABI = BrevisMetaData.ABI

// BrevisBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrevisMetaData.Bin instead.
var BrevisBin = BrevisMetaData.Bin

// DeployBrevis deploys a new Ethereum contract, binding an instance of Brevis to it.
func DeployBrevis(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Brevis, error) {
	parsed, err := BrevisMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrevisBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Brevis{BrevisCaller: BrevisCaller{contract: contract}, BrevisTransactor: BrevisTransactor{contract: contract}, BrevisFilterer: BrevisFilterer{contract: contract}}, nil
}

// Brevis is an auto generated Go binding around an Ethereum contract.
type Brevis struct {
	BrevisCaller     // Read-only binding to the contract
	BrevisTransactor // Write-only binding to the contract
	BrevisFilterer   // Log filterer for contract events
}

// BrevisCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrevisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrevisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrevisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrevisSession struct {
	Contract     *Brevis           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrevisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrevisCallerSession struct {
	Contract *BrevisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrevisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrevisTransactorSession struct {
	Contract     *BrevisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrevisRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrevisRaw struct {
	Contract *Brevis // Generic contract binding to access the raw methods on
}

// BrevisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrevisCallerRaw struct {
	Contract *BrevisCaller // Generic read-only contract binding to access the raw methods on
}

// BrevisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrevisTransactorRaw struct {
	Contract *BrevisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrevis creates a new instance of Brevis, bound to a specific deployed contract.
func NewBrevis(address common.Address, backend bind.ContractBackend) (*Brevis, error) {
	contract, err := bindBrevis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brevis{BrevisCaller: BrevisCaller{contract: contract}, BrevisTransactor: BrevisTransactor{contract: contract}, BrevisFilterer: BrevisFilterer{contract: contract}}, nil
}

// NewBrevisCaller creates a new read-only instance of Brevis, bound to a specific deployed contract.
func NewBrevisCaller(address common.Address, caller bind.ContractCaller) (*BrevisCaller, error) {
	contract, err := bindBrevis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisCaller{contract: contract}, nil
}

// NewBrevisTransactor creates a new write-only instance of Brevis, bound to a specific deployed contract.
func NewBrevisTransactor(address common.Address, transactor bind.ContractTransactor) (*BrevisTransactor, error) {
	contract, err := bindBrevis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisTransactor{contract: contract}, nil
}

// NewBrevisFilterer creates a new log filterer instance of Brevis, bound to a specific deployed contract.
func NewBrevisFilterer(address common.Address, filterer bind.ContractFilterer) (*BrevisFilterer, error) {
	contract, err := bindBrevis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrevisFilterer{contract: contract}, nil
}

// bindBrevis binds a generic wrapper to an already deployed contract.
func bindBrevis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrevisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brevis *BrevisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brevis.Contract.BrevisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brevis *BrevisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brevis.Contract.BrevisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brevis *BrevisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brevis.Contract.BrevisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brevis *BrevisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brevis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brevis *BrevisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brevis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brevis *BrevisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brevis.Contract.contract.Transact(opts, method, params...)
}

// BrevisRequestMetaData contains all meta data concerning the BrevisRequest contract.
var BrevisRequestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"contractIBrevisProof\",\"name\":\"_brevisProof\",\"type\":\"address\"},{\"internalType\":\"contractIBrevisEigen\",\"name\":\"_brevisEigen\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumBrevisRequest.AskForType\",\"name\":\"askFor\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"AskFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"ChallengeWindowUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"requestIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"queryURLs\",\"type\":\"bytes[]\"}],\"name\":\"OpRequestsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"ProofPost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"QueryDataPost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"RequestRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIBrevisApp\",\"name\":\"callback\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBrevisRequest.Option\",\"name\":\"option\",\"type\":\"uint8\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"RequestTimeoutUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"ResponseTimeoutUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"askForProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"askForQueryData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"brevisEigen\",\"outputs\":[{\"internalType\":\"contractIBrevisEigen\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"brevisProof\",\"outputs\":[{\"internalType\":\"contractIBrevisProof\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"challengeQueryData\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_requestIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_queryURLs\",\"type\":\"bytes[]\"}],\"name\":\"fulfillOpRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_appCircuitOutput\",\"type\":\"bytes\"}],\"name\":\"fulfillRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"keccakToMimc\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"postProof\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"postQueryData\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"queryRequestStatus\",\"outputs\":[{\"internalType\":\"enumBrevisRequest.RequestStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestExts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"canChallengeBefore\",\"type\":\"uint256\"},{\"internalType\":\"enumBrevisRequest.AskForType\",\"name\":\"askFor\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"shouldRespondBefore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundee\",\"type\":\"address\"},{\"internalType\":\"contractIBrevisApp\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"enumBrevisRequest.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumBrevisRequest.Option\",\"name\":\"option\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"responseTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_refundee\",\"type\":\"address\"},{\"internalType\":\"contractIBrevisApp\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"enumBrevisRequest.Option\",\"name\":\"_option\",\"type\":\"uint8\"}],\"name\":\"sendRequest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBrevisEigen\",\"name\":\"_brevisEigen\",\"type\":\"address\"}],\"name\":\"setBrevisEigen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengeWindow\",\"type\":\"uint256\"}],\"name\":\"setChallengeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"setRequestTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_responseTimeout\",\"type\":\"uint256\"}],\"name\":\"setResponseTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080346100d857601f61158c38819003918201601f19168301916001600160401b038311848410176100dc578084926060946040528339810103126100d85780516001600160a01b039190828116908190036100d8576020820151918383168093036100d85760400151918383168093036100d8575f549160018060a01b03199233848216175f55604051953391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3826001541617600155816003541617600355600454161760045561149b90816100f18239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040818152600480361015610020575b505050361561001e575f80fd5b005b5f92833560e01c908162a731ed14610fe05750806301c1aa0d14610f7f5780630de4c3d714610f585780632763ce8a14610f275780632d85fa2014610ee057806332e7ed8114610c125780633f20b4c914610bf357806361b9fd7314610bb8578063622b6af414610b5757806369c228a114610b2057806369dc190314610abf5780636db03b1d146107905780637249fbb6146109a75780637ff7b0d214610917578063861a1412146108f85780638da5cb5b146108d257806395d39656146108ab5780639d8669851461082d578063a17ed96c1461080e578063a42dce8014610795578063a4514ff314610790578063b6979c3e1461075f578063c415b95c14610737578063c7f5aaa01461070f578063da03667a146104a0578063f2a9c666146102305763f2fde38b03610011573461022c57602036600319011261022c576101696110e9565b908354906001600160a01b038083169361018433861461110c565b169384156101c35750506001600160a01b031916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b8280fd5b509190608036600319011261049c5782359061024a6110d3565b90604435906001600160a01b039283831680930361049857606435936003851015610494578587526020916005835283882054610451578116801561040e5761029560025442611157565b9084519160c0830183811067ffffffffffffffff8211176103fb57899389938c96936003938a5282528782013481528983019182528960608401988c8a526080850198818a5260a08601986102e98161102a565b8952815260058b52209251835551600183015583600283019151166001600160a01b031982541617905501935116835492519160088310156103e8579174ff00000000000000000000000000000000000000007f8351e4cc62a96a972da26714bfbb019926334976bb613c36e0dc72b69cf8f5649a989694927fffffffffffffffffffff0000000000000000000000000000000000000000000075ff00000000000000000000000000000000000000000060a09c9a9896516103aa8161102a565b6103b38161102a565b60a81b16941617918a1b161717905581519485523390850152349084015260608301526103df8161102a565b6080820152a180f35b634e487b7160e01b8b5260218c5260248bfd5b60418c634e487b7160e01b5f525260245ffd5b835162461bcd60e51b8152808a01849052601560248201527f726566756e646565206e6f742070726f766964656400000000000000000000006044820152606490fd5b835162461bcd60e51b8152808a01849052601860248201527f7265717565737420616c726561647920696e20717565756500000000000000006044820152606490fd5b8680fd5b8580fd5b5080fd5b5082903461049c57608036600319011261049c5780356104be6110ab565b9067ffffffffffffffff604435818111610498576104df9036908601611048565b9190946064358281116106c3576104f99036908301611048565b9390956001600160a01b03938460035416918b51996371e8f36b60e11b8b5288868c015260209a8b81602481885afa90811561070557926105758f95938f938f98969061054e9186916106d8575b50156111da565b86519889978896879563e044095360e01b8752168c86015260248501526044840191611226565b03925af19081156106ce57889161069d575b50840361065a5750958596838752600586526003828820017401000000000000000000000000000000000000000060ff60a01b198254161790557f85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6868351868152a18387526005865260038288200154169384610602578680f35b86956106336106419288958551958694850198633ceb5b5160e11b8a52602486015260448501526064840191611226565b03601f1981018352826111a0565b51925af15061064e611246565b50808280808080808680f35b875162461bcd60e51b8152908101869052601d60248201527f72657175657374496420616e642070726f6f66206e6f74206d617463680000006044820152606490fd5b90508681813d83116106c7575b6106b481836111a0565b810103126106c3575189610587565b8780fd5b503d6106aa565b89513d8a823e3d90fd5b6106f891508a3d8c116106fe575b6106f081836111a0565b8101906111c2565b5f610547565b503d6106e6565b8e513d8f823e3d90fd5b50503461049c578160031936011261049c576020906001600160a01b03600354169051908152f35b50503461049c578160031936011261049c576020906001600160a01b03600154169051908152f35b50913461078d57602036600319011261078d575061078b610782602093356112d1565b915180926110ff565bf35b80fd5b6110c2565b50503461049c57602036600319011261049c577f5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38906107d26110e9565b6001600160a01b036107e881865416331461110c565b80600154921690816001600160a01b03198416176001558351921682526020820152a180f35b50503461049c578160031936011261049c576020906009549051908152f35b503461022c57602036600319011261022c5760c092829135815260056020522061089b8154926001830154926001600160a01b03906003826002830154169101549260ff8460a81c1695815197885260208801528601528116606085015260ff608085019160a01c166110ff565b6108a48161102a565b60a0820152f35b503461022c57602036600319011261022c5760209282913581526007845220549051908152f35b50503461049c578160031936011261049c576001600160a01b0360209254169051908152f35b50503461049c578160031936011261049c576020906008549051908152f35b503461022c578160031936011261022c576109306110d3565b916001600160a01b036001541633036109645750828080610961948194359061c350f161095b611246565b50611285565b80f35b906020606492519162461bcd60e51b8352820152601160248201527f6e6f742066656520636f6c6c6563746f720000000000000000000000000000006044820152fd5b503461022c5760209081600319360112610abb5783813591828252600584528482205442111561049c576001600160a01b039084846024846003541693895194859384926371e8f36b60e11b84528301525afa908115610ab157928080610a6394610a3d82957ffea410cb461deba9fe807dde02d6641d82e1bf09ecc88ecfa0f2ffadf2a1fdfe9b988491610a9a5750156111da565b878252600589526001878320918383556002830154169101549061c350f161095b611246565b818552600583526003818620017407000000000000000000000000000000000000000060ff60a01b1982541617905551908152a180f35b6106f891508b3d8d116106fe576106f081836111a0565b86513d85823e3d90fd5b8380fd5b50903461022c57602036600319011261022c577f86fe7fc31f35681a1ed77325f0cf24935a5d25b1861e7ce9ceed9cb67f2222709135610b0a6001600160a01b03855416331461110c565b600954908060095582519182526020820152a180f35b50903461022c5736600319011261049c5760243567ffffffffffffffff811161022c57610b509250369101611048565b5050611420565b50903461022c57602036600319011261022c577f87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a639135610ba26001600160a01b03855416331461110c565b600254908060025582519182526020820152a180f35b83823461049c57606036600319011261049c57610bd36110ab565b5060443567ffffffffffffffff811161022c57610b509250369101611048565b50503461049c578160031936011261049c576020906002549051908152f35b503461022c578160031936011261022c5767ffffffffffffffff8135818111610edc57610c42903690840161107a565b909260249384358481116106c357610c5d903690840161107a565b9390958115610e9a57848203610e96576001600160a01b038085541694853b15610e9257895163d1e3471560e01b81528b816020988985830152818381610ca88c8c8c8401916113bf565b03925af18015610e8857610e75575b508a9897989695965b848110610d8b575050505090610cde918751978089528801916113bf565b918583038187015281835280830192818360051b820101958589925b858410610d2b578a7fa1c3ef280bbac31181188d4beffe370ebae03e08d58a3aa4832b107d540c0a838b8b038ca180f35b9091929394959697601f198282030188528835601e1984360301811215610d87578301868101919035858111610d83578036038313610d8357610d7388928392600195611226565b9a01980196959401929190610cfa565b8c80fd5b8b80fd5b60039998999796978c84825416610da384898b6113fc565b35813b1561022c5782918f83928a9151948593849263e26d07d360e01b84528b8401525af18015610e6b57610e57575b5050610de08287896113fc565b358d52600588528b8d20017402000000000000000000000000000000000000000060ff60a01b19825416179055610e1960085442611157565b610e248287896113fc565b358d52600688528b8d20555f198114610e4557600101989798969596610cc0565b634e487b7160e01b8c5260118252838cfd5b610e6090611178565b610d83578c5f610dd3565b8e513d84823e3d90fd5b610e81909b919b611178565b995f610cb7565b8b513d8e823e3d90fd5b8a80fd5b8880fd5b83601260649260208b519362461bcd60e51b85528401528201527f696e76616c6964207265717565737449647300000000000000000000000000006044820152fd5b8480fd5b83823461049c57602036600319011261049c5780356001600160a01b03808216809203610abb57610f1590845416331461110c565b6001600160a01b031982541617905580f35b83823461049c57602036600319011261049c57803567ffffffffffffffff811161022c57610b509250369101611048565b50913461078d578060031936011261078d57506001600160a01b0360209254169051908152f35b50903461022c57602036600319011261022c577fedb9338f4b0faf2b899d2d7f54b90753d2a8ebb34936e381edb91b091c3e45a79135610fca6001600160a01b03855416331461110c565b600854908060085582519182526020820152a180f35b9291905034610abb576020366003190112610abb57606093829135815260066020522090815491600260ff6001830154169101549284526110208161102a565b6020840152820152f35b6003111561103457565b634e487b7160e01b5f52602160045260245ffd5b9181601f840112156110765782359167ffffffffffffffff8311611076576020838186019501011161107657565b5f80fd5b9181601f840112156110765782359167ffffffffffffffff8311611076576020808501948460051b01011161107657565b6024359067ffffffffffffffff8216820361107657565b602036600319011215611420575f80fd5b602435906001600160a01b038216820361107657565b600435906001600160a01b038216820361107657565b9060088210156110345752565b1561111357565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b9190820180921161116457565b634e487b7160e01b5f52601160045260245ffd5b67ffffffffffffffff811161118c57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761118c57604052565b90816020910312611076575180151581036110765790565b156111e157565b60405162461bcd60e51b815260206004820152601760248201527f70726f6f6620616c72656164792067656e6572617465640000000000000000006044820152606490fd5b908060209392818452848401375f828201840152601f01601f1916010190565b3d15611280573d9067ffffffffffffffff821161118c5760405191611275601f8201601f1916602001846111a0565b82523d5f602084013e565b606090565b1561128c57565b60405162461bcd60e51b815260206004820152601260248201527f73656e64206e6174697665206661696c656400000000000000000000000000006044820152606490fd5b5f8181526020906005825260409060ff6003838320015460a01c166008811015611368576002148015611399575b80611385575b61137c578381526005835260ff6003838320015460a01c1660088110156113685760041480611351575b61134857600560039360ff9583525220015460a01c1690565b50505050600590565b50838152600683526002828220015442101561132f565b634e487b7160e01b82526021600452602482fd5b50505050600690565b508381526006835281812054421015611305565b508381526005835260ff6003838320015460a01c166008811015611368576003146112ff565b90918281527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83116110765760209260051b809284830137010190565b919081101561140c5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b60405162461bcd60e51b815260206004820152600f60248201527f6e6f7420696d706c656d656e74656400000000000000000000000000000000006044820152606490fdfea26469706673582212208de76c0ef54023bd4aca98f554f769b4ce2ca47d9a5a2dbfbd1aea3ef253b54f64736f6c63430008140033",
}

// BrevisRequestABI is the input ABI used to generate the binding from.
// Deprecated: Use BrevisRequestMetaData.ABI instead.
var BrevisRequestABI = BrevisRequestMetaData.ABI

// BrevisRequestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrevisRequestMetaData.Bin instead.
var BrevisRequestBin = BrevisRequestMetaData.Bin

// DeployBrevisRequest deploys a new Ethereum contract, binding an instance of BrevisRequest to it.
func DeployBrevisRequest(auth *bind.TransactOpts, backend bind.ContractBackend, _feeCollector common.Address, _brevisProof common.Address, _brevisEigen common.Address) (common.Address, *types.Transaction, *BrevisRequest, error) {
	parsed, err := BrevisRequestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrevisRequestBin), backend, _feeCollector, _brevisProof, _brevisEigen)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BrevisRequest{BrevisRequestCaller: BrevisRequestCaller{contract: contract}, BrevisRequestTransactor: BrevisRequestTransactor{contract: contract}, BrevisRequestFilterer: BrevisRequestFilterer{contract: contract}}, nil
}

// BrevisRequest is an auto generated Go binding around an Ethereum contract.
type BrevisRequest struct {
	BrevisRequestCaller     // Read-only binding to the contract
	BrevisRequestTransactor // Write-only binding to the contract
	BrevisRequestFilterer   // Log filterer for contract events
}

// BrevisRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrevisRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrevisRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrevisRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrevisRequestSession struct {
	Contract     *BrevisRequest    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrevisRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrevisRequestCallerSession struct {
	Contract *BrevisRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BrevisRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrevisRequestTransactorSession struct {
	Contract     *BrevisRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BrevisRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrevisRequestRaw struct {
	Contract *BrevisRequest // Generic contract binding to access the raw methods on
}

// BrevisRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrevisRequestCallerRaw struct {
	Contract *BrevisRequestCaller // Generic read-only contract binding to access the raw methods on
}

// BrevisRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrevisRequestTransactorRaw struct {
	Contract *BrevisRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrevisRequest creates a new instance of BrevisRequest, bound to a specific deployed contract.
func NewBrevisRequest(address common.Address, backend bind.ContractBackend) (*BrevisRequest, error) {
	contract, err := bindBrevisRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrevisRequest{BrevisRequestCaller: BrevisRequestCaller{contract: contract}, BrevisRequestTransactor: BrevisRequestTransactor{contract: contract}, BrevisRequestFilterer: BrevisRequestFilterer{contract: contract}}, nil
}

// NewBrevisRequestCaller creates a new read-only instance of BrevisRequest, bound to a specific deployed contract.
func NewBrevisRequestCaller(address common.Address, caller bind.ContractCaller) (*BrevisRequestCaller, error) {
	contract, err := bindBrevisRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestCaller{contract: contract}, nil
}

// NewBrevisRequestTransactor creates a new write-only instance of BrevisRequest, bound to a specific deployed contract.
func NewBrevisRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*BrevisRequestTransactor, error) {
	contract, err := bindBrevisRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestTransactor{contract: contract}, nil
}

// NewBrevisRequestFilterer creates a new log filterer instance of BrevisRequest, bound to a specific deployed contract.
func NewBrevisRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*BrevisRequestFilterer, error) {
	contract, err := bindBrevisRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestFilterer{contract: contract}, nil
}

// bindBrevisRequest binds a generic wrapper to an already deployed contract.
func bindBrevisRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrevisRequestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevisRequest *BrevisRequestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevisRequest.Contract.BrevisRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevisRequest *BrevisRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisRequest.Contract.BrevisRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevisRequest *BrevisRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevisRequest.Contract.BrevisRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevisRequest *BrevisRequestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevisRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevisRequest *BrevisRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevisRequest *BrevisRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevisRequest.Contract.contract.Transact(opts, method, params...)
}

// BrevisEigen is a free data retrieval call binding the contract method 0x0de4c3d7.
//
// Solidity: function brevisEigen() view returns(address)
func (_BrevisRequest *BrevisRequestCaller) BrevisEigen(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "brevisEigen")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrevisEigen is a free data retrieval call binding the contract method 0x0de4c3d7.
//
// Solidity: function brevisEigen() view returns(address)
func (_BrevisRequest *BrevisRequestSession) BrevisEigen() (common.Address, error) {
	return _BrevisRequest.Contract.BrevisEigen(&_BrevisRequest.CallOpts)
}

// BrevisEigen is a free data retrieval call binding the contract method 0x0de4c3d7.
//
// Solidity: function brevisEigen() view returns(address)
func (_BrevisRequest *BrevisRequestCallerSession) BrevisEigen() (common.Address, error) {
	return _BrevisRequest.Contract.BrevisEigen(&_BrevisRequest.CallOpts)
}

// BrevisProof is a free data retrieval call binding the contract method 0xc7f5aaa0.
//
// Solidity: function brevisProof() view returns(address)
func (_BrevisRequest *BrevisRequestCaller) BrevisProof(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "brevisProof")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BrevisProof is a free data retrieval call binding the contract method 0xc7f5aaa0.
//
// Solidity: function brevisProof() view returns(address)
func (_BrevisRequest *BrevisRequestSession) BrevisProof() (common.Address, error) {
	return _BrevisRequest.Contract.BrevisProof(&_BrevisRequest.CallOpts)
}

// BrevisProof is a free data retrieval call binding the contract method 0xc7f5aaa0.
//
// Solidity: function brevisProof() view returns(address)
func (_BrevisRequest *BrevisRequestCallerSession) BrevisProof() (common.Address, error) {
	return _BrevisRequest.Contract.BrevisProof(&_BrevisRequest.CallOpts)
}

// ChallengeQueryData is a free data retrieval call binding the contract method 0x2763ce8a.
//
// Solidity: function challengeQueryData(bytes ) pure returns()
func (_BrevisRequest *BrevisRequestCaller) ChallengeQueryData(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "challengeQueryData", arg0)

	if err != nil {
		return err
	}

	return err

}

// ChallengeQueryData is a free data retrieval call binding the contract method 0x2763ce8a.
//
// Solidity: function challengeQueryData(bytes ) pure returns()
func (_BrevisRequest *BrevisRequestSession) ChallengeQueryData(arg0 []byte) error {
	return _BrevisRequest.Contract.ChallengeQueryData(&_BrevisRequest.CallOpts, arg0)
}

// ChallengeQueryData is a free data retrieval call binding the contract method 0x2763ce8a.
//
// Solidity: function challengeQueryData(bytes ) pure returns()
func (_BrevisRequest *BrevisRequestCallerSession) ChallengeQueryData(arg0 []byte) error {
	return _BrevisRequest.Contract.ChallengeQueryData(&_BrevisRequest.CallOpts, arg0)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_BrevisRequest *BrevisRequestCaller) ChallengeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "challengeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_BrevisRequest *BrevisRequestSession) ChallengeWindow() (*big.Int, error) {
	return _BrevisRequest.Contract.ChallengeWindow(&_BrevisRequest.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint256)
func (_BrevisRequest *BrevisRequestCallerSession) ChallengeWindow() (*big.Int, error) {
	return _BrevisRequest.Contract.ChallengeWindow(&_BrevisRequest.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_BrevisRequest *BrevisRequestCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_BrevisRequest *BrevisRequestSession) FeeCollector() (common.Address, error) {
	return _BrevisRequest.Contract.FeeCollector(&_BrevisRequest.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_BrevisRequest *BrevisRequestCallerSession) FeeCollector() (common.Address, error) {
	return _BrevisRequest.Contract.FeeCollector(&_BrevisRequest.CallOpts)
}

// KeccakToMimc is a free data retrieval call binding the contract method 0x95d39656.
//
// Solidity: function keccakToMimc(bytes32 ) view returns(bytes32)
func (_BrevisRequest *BrevisRequestCaller) KeccakToMimc(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "keccakToMimc", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeccakToMimc is a free data retrieval call binding the contract method 0x95d39656.
//
// Solidity: function keccakToMimc(bytes32 ) view returns(bytes32)
func (_BrevisRequest *BrevisRequestSession) KeccakToMimc(arg0 [32]byte) ([32]byte, error) {
	return _BrevisRequest.Contract.KeccakToMimc(&_BrevisRequest.CallOpts, arg0)
}

// KeccakToMimc is a free data retrieval call binding the contract method 0x95d39656.
//
// Solidity: function keccakToMimc(bytes32 ) view returns(bytes32)
func (_BrevisRequest *BrevisRequestCallerSession) KeccakToMimc(arg0 [32]byte) ([32]byte, error) {
	return _BrevisRequest.Contract.KeccakToMimc(&_BrevisRequest.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisRequest *BrevisRequestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisRequest *BrevisRequestSession) Owner() (common.Address, error) {
	return _BrevisRequest.Contract.Owner(&_BrevisRequest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisRequest *BrevisRequestCallerSession) Owner() (common.Address, error) {
	return _BrevisRequest.Contract.Owner(&_BrevisRequest.CallOpts)
}

// PostProof is a free data retrieval call binding the contract method 0x61b9fd73.
//
// Solidity: function postProof(bytes32 , uint64 , bytes ) pure returns()
func (_BrevisRequest *BrevisRequestCaller) PostProof(opts *bind.CallOpts, arg0 [32]byte, arg1 uint64, arg2 []byte) error {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "postProof", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// PostProof is a free data retrieval call binding the contract method 0x61b9fd73.
//
// Solidity: function postProof(bytes32 , uint64 , bytes ) pure returns()
func (_BrevisRequest *BrevisRequestSession) PostProof(arg0 [32]byte, arg1 uint64, arg2 []byte) error {
	return _BrevisRequest.Contract.PostProof(&_BrevisRequest.CallOpts, arg0, arg1, arg2)
}

// PostProof is a free data retrieval call binding the contract method 0x61b9fd73.
//
// Solidity: function postProof(bytes32 , uint64 , bytes ) pure returns()
func (_BrevisRequest *BrevisRequestCallerSession) PostProof(arg0 [32]byte, arg1 uint64, arg2 []byte) error {
	return _BrevisRequest.Contract.PostProof(&_BrevisRequest.CallOpts, arg0, arg1, arg2)
}

// PostQueryData is a free data retrieval call binding the contract method 0x69c228a1.
//
// Solidity: function postQueryData(bytes32 , bytes ) pure returns()
func (_BrevisRequest *BrevisRequestCaller) PostQueryData(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) error {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "postQueryData", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// PostQueryData is a free data retrieval call binding the contract method 0x69c228a1.
//
// Solidity: function postQueryData(bytes32 , bytes ) pure returns()
func (_BrevisRequest *BrevisRequestSession) PostQueryData(arg0 [32]byte, arg1 []byte) error {
	return _BrevisRequest.Contract.PostQueryData(&_BrevisRequest.CallOpts, arg0, arg1)
}

// PostQueryData is a free data retrieval call binding the contract method 0x69c228a1.
//
// Solidity: function postQueryData(bytes32 , bytes ) pure returns()
func (_BrevisRequest *BrevisRequestCallerSession) PostQueryData(arg0 [32]byte, arg1 []byte) error {
	return _BrevisRequest.Contract.PostQueryData(&_BrevisRequest.CallOpts, arg0, arg1)
}

// QueryRequestStatus is a free data retrieval call binding the contract method 0xb6979c3e.
//
// Solidity: function queryRequestStatus(bytes32 _requestId) view returns(uint8)
func (_BrevisRequest *BrevisRequestCaller) QueryRequestStatus(opts *bind.CallOpts, _requestId [32]byte) (uint8, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "queryRequestStatus", _requestId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QueryRequestStatus is a free data retrieval call binding the contract method 0xb6979c3e.
//
// Solidity: function queryRequestStatus(bytes32 _requestId) view returns(uint8)
func (_BrevisRequest *BrevisRequestSession) QueryRequestStatus(_requestId [32]byte) (uint8, error) {
	return _BrevisRequest.Contract.QueryRequestStatus(&_BrevisRequest.CallOpts, _requestId)
}

// QueryRequestStatus is a free data retrieval call binding the contract method 0xb6979c3e.
//
// Solidity: function queryRequestStatus(bytes32 _requestId) view returns(uint8)
func (_BrevisRequest *BrevisRequestCallerSession) QueryRequestStatus(_requestId [32]byte) (uint8, error) {
	return _BrevisRequest.Contract.QueryRequestStatus(&_BrevisRequest.CallOpts, _requestId)
}

// RequestExts is a free data retrieval call binding the contract method 0x00a731ed.
//
// Solidity: function requestExts(bytes32 ) view returns(uint256 canChallengeBefore, uint8 askFor, uint256 shouldRespondBefore)
func (_BrevisRequest *BrevisRequestCaller) RequestExts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CanChallengeBefore  *big.Int
	AskFor              uint8
	ShouldRespondBefore *big.Int
}, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "requestExts", arg0)

	outstruct := new(struct {
		CanChallengeBefore  *big.Int
		AskFor              uint8
		ShouldRespondBefore *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanChallengeBefore = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AskFor = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ShouldRespondBefore = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RequestExts is a free data retrieval call binding the contract method 0x00a731ed.
//
// Solidity: function requestExts(bytes32 ) view returns(uint256 canChallengeBefore, uint8 askFor, uint256 shouldRespondBefore)
func (_BrevisRequest *BrevisRequestSession) RequestExts(arg0 [32]byte) (struct {
	CanChallengeBefore  *big.Int
	AskFor              uint8
	ShouldRespondBefore *big.Int
}, error) {
	return _BrevisRequest.Contract.RequestExts(&_BrevisRequest.CallOpts, arg0)
}

// RequestExts is a free data retrieval call binding the contract method 0x00a731ed.
//
// Solidity: function requestExts(bytes32 ) view returns(uint256 canChallengeBefore, uint8 askFor, uint256 shouldRespondBefore)
func (_BrevisRequest *BrevisRequestCallerSession) RequestExts(arg0 [32]byte) (struct {
	CanChallengeBefore  *big.Int
	AskFor              uint8
	ShouldRespondBefore *big.Int
}, error) {
	return _BrevisRequest.Contract.RequestExts(&_BrevisRequest.CallOpts, arg0)
}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint256)
func (_BrevisRequest *BrevisRequestCaller) RequestTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "requestTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint256)
func (_BrevisRequest *BrevisRequestSession) RequestTimeout() (*big.Int, error) {
	return _BrevisRequest.Contract.RequestTimeout(&_BrevisRequest.CallOpts)
}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint256)
func (_BrevisRequest *BrevisRequestCallerSession) RequestTimeout() (*big.Int, error) {
	return _BrevisRequest.Contract.RequestTimeout(&_BrevisRequest.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 deadline, uint256 fee, address refundee, address callback, uint8 status, uint8 option)
func (_BrevisRequest *BrevisRequestCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Deadline *big.Int
	Fee      *big.Int
	Refundee common.Address
	Callback common.Address
	Status   uint8
	Option   uint8
}, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Deadline *big.Int
		Fee      *big.Int
		Refundee common.Address
		Callback common.Address
		Status   uint8
		Option   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deadline = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Refundee = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Callback = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Option = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 deadline, uint256 fee, address refundee, address callback, uint8 status, uint8 option)
func (_BrevisRequest *BrevisRequestSession) Requests(arg0 [32]byte) (struct {
	Deadline *big.Int
	Fee      *big.Int
	Refundee common.Address
	Callback common.Address
	Status   uint8
	Option   uint8
}, error) {
	return _BrevisRequest.Contract.Requests(&_BrevisRequest.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 deadline, uint256 fee, address refundee, address callback, uint8 status, uint8 option)
func (_BrevisRequest *BrevisRequestCallerSession) Requests(arg0 [32]byte) (struct {
	Deadline *big.Int
	Fee      *big.Int
	Refundee common.Address
	Callback common.Address
	Status   uint8
	Option   uint8
}, error) {
	return _BrevisRequest.Contract.Requests(&_BrevisRequest.CallOpts, arg0)
}

// ResponseTimeout is a free data retrieval call binding the contract method 0xa17ed96c.
//
// Solidity: function responseTimeout() view returns(uint256)
func (_BrevisRequest *BrevisRequestCaller) ResponseTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisRequest.contract.Call(opts, &out, "responseTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResponseTimeout is a free data retrieval call binding the contract method 0xa17ed96c.
//
// Solidity: function responseTimeout() view returns(uint256)
func (_BrevisRequest *BrevisRequestSession) ResponseTimeout() (*big.Int, error) {
	return _BrevisRequest.Contract.ResponseTimeout(&_BrevisRequest.CallOpts)
}

// ResponseTimeout is a free data retrieval call binding the contract method 0xa17ed96c.
//
// Solidity: function responseTimeout() view returns(uint256)
func (_BrevisRequest *BrevisRequestCallerSession) ResponseTimeout() (*big.Int, error) {
	return _BrevisRequest.Contract.ResponseTimeout(&_BrevisRequest.CallOpts)
}

// AskForProof is a paid mutator transaction binding the contract method 0x6db03b1d.
//
// Solidity: function askForProof(bytes32 ) payable returns()
func (_BrevisRequest *BrevisRequestTransactor) AskForProof(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "askForProof", arg0)
}

// AskForProof is a paid mutator transaction binding the contract method 0x6db03b1d.
//
// Solidity: function askForProof(bytes32 ) payable returns()
func (_BrevisRequest *BrevisRequestSession) AskForProof(arg0 [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.AskForProof(&_BrevisRequest.TransactOpts, arg0)
}

// AskForProof is a paid mutator transaction binding the contract method 0x6db03b1d.
//
// Solidity: function askForProof(bytes32 ) payable returns()
func (_BrevisRequest *BrevisRequestTransactorSession) AskForProof(arg0 [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.AskForProof(&_BrevisRequest.TransactOpts, arg0)
}

// AskForQueryData is a paid mutator transaction binding the contract method 0xa4514ff3.
//
// Solidity: function askForQueryData(bytes32 ) payable returns()
func (_BrevisRequest *BrevisRequestTransactor) AskForQueryData(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "askForQueryData", arg0)
}

// AskForQueryData is a paid mutator transaction binding the contract method 0xa4514ff3.
//
// Solidity: function askForQueryData(bytes32 ) payable returns()
func (_BrevisRequest *BrevisRequestSession) AskForQueryData(arg0 [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.AskForQueryData(&_BrevisRequest.TransactOpts, arg0)
}

// AskForQueryData is a paid mutator transaction binding the contract method 0xa4514ff3.
//
// Solidity: function askForQueryData(bytes32 ) payable returns()
func (_BrevisRequest *BrevisRequestTransactorSession) AskForQueryData(arg0 [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.AskForQueryData(&_BrevisRequest.TransactOpts, arg0)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_BrevisRequest *BrevisRequestTransactor) CollectFee(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "collectFee", _amount, _to)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_BrevisRequest *BrevisRequestSession) CollectFee(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.CollectFee(&_BrevisRequest.TransactOpts, _amount, _to)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) CollectFee(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.CollectFee(&_BrevisRequest.TransactOpts, _amount, _to)
}

// FulfillOpRequests is a paid mutator transaction binding the contract method 0x32e7ed81.
//
// Solidity: function fulfillOpRequests(bytes32[] _requestIds, bytes[] _queryURLs) returns()
func (_BrevisRequest *BrevisRequestTransactor) FulfillOpRequests(opts *bind.TransactOpts, _requestIds [][32]byte, _queryURLs [][]byte) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "fulfillOpRequests", _requestIds, _queryURLs)
}

// FulfillOpRequests is a paid mutator transaction binding the contract method 0x32e7ed81.
//
// Solidity: function fulfillOpRequests(bytes32[] _requestIds, bytes[] _queryURLs) returns()
func (_BrevisRequest *BrevisRequestSession) FulfillOpRequests(_requestIds [][32]byte, _queryURLs [][]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.FulfillOpRequests(&_BrevisRequest.TransactOpts, _requestIds, _queryURLs)
}

// FulfillOpRequests is a paid mutator transaction binding the contract method 0x32e7ed81.
//
// Solidity: function fulfillOpRequests(bytes32[] _requestIds, bytes[] _queryURLs) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) FulfillOpRequests(_requestIds [][32]byte, _queryURLs [][]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.FulfillOpRequests(&_BrevisRequest.TransactOpts, _requestIds, _queryURLs)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0xda03667a.
//
// Solidity: function fulfillRequest(bytes32 _requestId, uint64 _chainId, bytes _proof, bytes _appCircuitOutput) returns()
func (_BrevisRequest *BrevisRequestTransactor) FulfillRequest(opts *bind.TransactOpts, _requestId [32]byte, _chainId uint64, _proof []byte, _appCircuitOutput []byte) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "fulfillRequest", _requestId, _chainId, _proof, _appCircuitOutput)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0xda03667a.
//
// Solidity: function fulfillRequest(bytes32 _requestId, uint64 _chainId, bytes _proof, bytes _appCircuitOutput) returns()
func (_BrevisRequest *BrevisRequestSession) FulfillRequest(_requestId [32]byte, _chainId uint64, _proof []byte, _appCircuitOutput []byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.FulfillRequest(&_BrevisRequest.TransactOpts, _requestId, _chainId, _proof, _appCircuitOutput)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0xda03667a.
//
// Solidity: function fulfillRequest(bytes32 _requestId, uint64 _chainId, bytes _proof, bytes _appCircuitOutput) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) FulfillRequest(_requestId [32]byte, _chainId uint64, _proof []byte, _appCircuitOutput []byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.FulfillRequest(&_BrevisRequest.TransactOpts, _requestId, _chainId, _proof, _appCircuitOutput)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _requestId) returns()
func (_BrevisRequest *BrevisRequestTransactor) Refund(opts *bind.TransactOpts, _requestId [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "refund", _requestId)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _requestId) returns()
func (_BrevisRequest *BrevisRequestSession) Refund(_requestId [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.Refund(&_BrevisRequest.TransactOpts, _requestId)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _requestId) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) Refund(_requestId [32]byte) (*types.Transaction, error) {
	return _BrevisRequest.Contract.Refund(&_BrevisRequest.TransactOpts, _requestId)
}

// SendRequest is a paid mutator transaction binding the contract method 0xf2a9c666.
//
// Solidity: function sendRequest(bytes32 _requestId, address _refundee, address _callback, uint8 _option) payable returns()
func (_BrevisRequest *BrevisRequestTransactor) SendRequest(opts *bind.TransactOpts, _requestId [32]byte, _refundee common.Address, _callback common.Address, _option uint8) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "sendRequest", _requestId, _refundee, _callback, _option)
}

// SendRequest is a paid mutator transaction binding the contract method 0xf2a9c666.
//
// Solidity: function sendRequest(bytes32 _requestId, address _refundee, address _callback, uint8 _option) payable returns()
func (_BrevisRequest *BrevisRequestSession) SendRequest(_requestId [32]byte, _refundee common.Address, _callback common.Address, _option uint8) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SendRequest(&_BrevisRequest.TransactOpts, _requestId, _refundee, _callback, _option)
}

// SendRequest is a paid mutator transaction binding the contract method 0xf2a9c666.
//
// Solidity: function sendRequest(bytes32 _requestId, address _refundee, address _callback, uint8 _option) payable returns()
func (_BrevisRequest *BrevisRequestTransactorSession) SendRequest(_requestId [32]byte, _refundee common.Address, _callback common.Address, _option uint8) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SendRequest(&_BrevisRequest.TransactOpts, _requestId, _refundee, _callback, _option)
}

// SetBrevisEigen is a paid mutator transaction binding the contract method 0x2d85fa20.
//
// Solidity: function setBrevisEigen(address _brevisEigen) returns()
func (_BrevisRequest *BrevisRequestTransactor) SetBrevisEigen(opts *bind.TransactOpts, _brevisEigen common.Address) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "setBrevisEigen", _brevisEigen)
}

// SetBrevisEigen is a paid mutator transaction binding the contract method 0x2d85fa20.
//
// Solidity: function setBrevisEigen(address _brevisEigen) returns()
func (_BrevisRequest *BrevisRequestSession) SetBrevisEigen(_brevisEigen common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetBrevisEigen(&_BrevisRequest.TransactOpts, _brevisEigen)
}

// SetBrevisEigen is a paid mutator transaction binding the contract method 0x2d85fa20.
//
// Solidity: function setBrevisEigen(address _brevisEigen) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) SetBrevisEigen(_brevisEigen common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetBrevisEigen(&_BrevisRequest.TransactOpts, _brevisEigen)
}

// SetChallengeWindow is a paid mutator transaction binding the contract method 0x01c1aa0d.
//
// Solidity: function setChallengeWindow(uint256 _challengeWindow) returns()
func (_BrevisRequest *BrevisRequestTransactor) SetChallengeWindow(opts *bind.TransactOpts, _challengeWindow *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "setChallengeWindow", _challengeWindow)
}

// SetChallengeWindow is a paid mutator transaction binding the contract method 0x01c1aa0d.
//
// Solidity: function setChallengeWindow(uint256 _challengeWindow) returns()
func (_BrevisRequest *BrevisRequestSession) SetChallengeWindow(_challengeWindow *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetChallengeWindow(&_BrevisRequest.TransactOpts, _challengeWindow)
}

// SetChallengeWindow is a paid mutator transaction binding the contract method 0x01c1aa0d.
//
// Solidity: function setChallengeWindow(uint256 _challengeWindow) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) SetChallengeWindow(_challengeWindow *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetChallengeWindow(&_BrevisRequest.TransactOpts, _challengeWindow)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_BrevisRequest *BrevisRequestTransactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_BrevisRequest *BrevisRequestSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetFeeCollector(&_BrevisRequest.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetFeeCollector(&_BrevisRequest.TransactOpts, _feeCollector)
}

// SetRequestTimeout is a paid mutator transaction binding the contract method 0x622b6af4.
//
// Solidity: function setRequestTimeout(uint256 _timeout) returns()
func (_BrevisRequest *BrevisRequestTransactor) SetRequestTimeout(opts *bind.TransactOpts, _timeout *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "setRequestTimeout", _timeout)
}

// SetRequestTimeout is a paid mutator transaction binding the contract method 0x622b6af4.
//
// Solidity: function setRequestTimeout(uint256 _timeout) returns()
func (_BrevisRequest *BrevisRequestSession) SetRequestTimeout(_timeout *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetRequestTimeout(&_BrevisRequest.TransactOpts, _timeout)
}

// SetRequestTimeout is a paid mutator transaction binding the contract method 0x622b6af4.
//
// Solidity: function setRequestTimeout(uint256 _timeout) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) SetRequestTimeout(_timeout *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetRequestTimeout(&_BrevisRequest.TransactOpts, _timeout)
}

// SetResponseTimeout is a paid mutator transaction binding the contract method 0x69dc1903.
//
// Solidity: function setResponseTimeout(uint256 _responseTimeout) returns()
func (_BrevisRequest *BrevisRequestTransactor) SetResponseTimeout(opts *bind.TransactOpts, _responseTimeout *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "setResponseTimeout", _responseTimeout)
}

// SetResponseTimeout is a paid mutator transaction binding the contract method 0x69dc1903.
//
// Solidity: function setResponseTimeout(uint256 _responseTimeout) returns()
func (_BrevisRequest *BrevisRequestSession) SetResponseTimeout(_responseTimeout *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetResponseTimeout(&_BrevisRequest.TransactOpts, _responseTimeout)
}

// SetResponseTimeout is a paid mutator transaction binding the contract method 0x69dc1903.
//
// Solidity: function setResponseTimeout(uint256 _responseTimeout) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) SetResponseTimeout(_responseTimeout *big.Int) (*types.Transaction, error) {
	return _BrevisRequest.Contract.SetResponseTimeout(&_BrevisRequest.TransactOpts, _responseTimeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisRequest *BrevisRequestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BrevisRequest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisRequest *BrevisRequestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.TransferOwnership(&_BrevisRequest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisRequest *BrevisRequestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevisRequest.Contract.TransferOwnership(&_BrevisRequest.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BrevisRequest *BrevisRequestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisRequest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BrevisRequest *BrevisRequestSession) Receive() (*types.Transaction, error) {
	return _BrevisRequest.Contract.Receive(&_BrevisRequest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BrevisRequest *BrevisRequestTransactorSession) Receive() (*types.Transaction, error) {
	return _BrevisRequest.Contract.Receive(&_BrevisRequest.TransactOpts)
}

// BrevisRequestAskForIterator is returned from FilterAskFor and is used to iterate over the raw logs and unpacked data for AskFor events raised by the BrevisRequest contract.
type BrevisRequestAskForIterator struct {
	Event *BrevisRequestAskFor // Event containing the contract specifics and raw log

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
func (it *BrevisRequestAskForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestAskFor)
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
		it.Event = new(BrevisRequestAskFor)
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
func (it *BrevisRequestAskForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestAskForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestAskFor represents a AskFor event raised by the BrevisRequest contract.
type BrevisRequestAskFor struct {
	RequestId [32]byte
	AskFor    uint8
	From      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAskFor is a free log retrieval operation binding the contract event 0xe77f1579b66737e0839e75b8ee1a189f5b6b5d2e504090f76337ffde5903da3a.
//
// Solidity: event AskFor(bytes32 indexed requestId, uint8 askFor, address from)
func (_BrevisRequest *BrevisRequestFilterer) FilterAskFor(opts *bind.FilterOpts, requestId [][32]byte) (*BrevisRequestAskForIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "AskFor", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestAskForIterator{contract: _BrevisRequest.contract, event: "AskFor", logs: logs, sub: sub}, nil
}

// WatchAskFor is a free log subscription operation binding the contract event 0xe77f1579b66737e0839e75b8ee1a189f5b6b5d2e504090f76337ffde5903da3a.
//
// Solidity: event AskFor(bytes32 indexed requestId, uint8 askFor, address from)
func (_BrevisRequest *BrevisRequestFilterer) WatchAskFor(opts *bind.WatchOpts, sink chan<- *BrevisRequestAskFor, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "AskFor", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestAskFor)
				if err := _BrevisRequest.contract.UnpackLog(event, "AskFor", log); err != nil {
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

// ParseAskFor is a log parse operation binding the contract event 0xe77f1579b66737e0839e75b8ee1a189f5b6b5d2e504090f76337ffde5903da3a.
//
// Solidity: event AskFor(bytes32 indexed requestId, uint8 askFor, address from)
func (_BrevisRequest *BrevisRequestFilterer) ParseAskFor(log types.Log) (*BrevisRequestAskFor, error) {
	event := new(BrevisRequestAskFor)
	if err := _BrevisRequest.contract.UnpackLog(event, "AskFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestChallengeWindowUpdatedIterator is returned from FilterChallengeWindowUpdated and is used to iterate over the raw logs and unpacked data for ChallengeWindowUpdated events raised by the BrevisRequest contract.
type BrevisRequestChallengeWindowUpdatedIterator struct {
	Event *BrevisRequestChallengeWindowUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisRequestChallengeWindowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestChallengeWindowUpdated)
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
		it.Event = new(BrevisRequestChallengeWindowUpdated)
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
func (it *BrevisRequestChallengeWindowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestChallengeWindowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestChallengeWindowUpdated represents a ChallengeWindowUpdated event raised by the BrevisRequest contract.
type BrevisRequestChallengeWindowUpdated struct {
	From *big.Int
	To   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChallengeWindowUpdated is a free log retrieval operation binding the contract event 0xedb9338f4b0faf2b899d2d7f54b90753d2a8ebb34936e381edb91b091c3e45a7.
//
// Solidity: event ChallengeWindowUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) FilterChallengeWindowUpdated(opts *bind.FilterOpts) (*BrevisRequestChallengeWindowUpdatedIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "ChallengeWindowUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestChallengeWindowUpdatedIterator{contract: _BrevisRequest.contract, event: "ChallengeWindowUpdated", logs: logs, sub: sub}, nil
}

// WatchChallengeWindowUpdated is a free log subscription operation binding the contract event 0xedb9338f4b0faf2b899d2d7f54b90753d2a8ebb34936e381edb91b091c3e45a7.
//
// Solidity: event ChallengeWindowUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) WatchChallengeWindowUpdated(opts *bind.WatchOpts, sink chan<- *BrevisRequestChallengeWindowUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "ChallengeWindowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestChallengeWindowUpdated)
				if err := _BrevisRequest.contract.UnpackLog(event, "ChallengeWindowUpdated", log); err != nil {
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

// ParseChallengeWindowUpdated is a log parse operation binding the contract event 0xedb9338f4b0faf2b899d2d7f54b90753d2a8ebb34936e381edb91b091c3e45a7.
//
// Solidity: event ChallengeWindowUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) ParseChallengeWindowUpdated(log types.Log) (*BrevisRequestChallengeWindowUpdated, error) {
	event := new(BrevisRequestChallengeWindowUpdated)
	if err := _BrevisRequest.contract.UnpackLog(event, "ChallengeWindowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestFeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the BrevisRequest contract.
type BrevisRequestFeeCollectorUpdatedIterator struct {
	Event *BrevisRequestFeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisRequestFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestFeeCollectorUpdated)
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
		it.Event = new(BrevisRequestFeeCollectorUpdated)
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
func (it *BrevisRequestFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestFeeCollectorUpdated represents a FeeCollectorUpdated event raised by the BrevisRequest contract.
type BrevisRequestFeeCollectorUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_BrevisRequest *BrevisRequestFilterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts) (*BrevisRequestFeeCollectorUpdatedIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestFeeCollectorUpdatedIterator{contract: _BrevisRequest.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_BrevisRequest *BrevisRequestFilterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *BrevisRequestFeeCollectorUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestFeeCollectorUpdated)
				if err := _BrevisRequest.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
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

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_BrevisRequest *BrevisRequestFilterer) ParseFeeCollectorUpdated(log types.Log) (*BrevisRequestFeeCollectorUpdated, error) {
	event := new(BrevisRequestFeeCollectorUpdated)
	if err := _BrevisRequest.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestOpRequestsFulfilledIterator is returned from FilterOpRequestsFulfilled and is used to iterate over the raw logs and unpacked data for OpRequestsFulfilled events raised by the BrevisRequest contract.
type BrevisRequestOpRequestsFulfilledIterator struct {
	Event *BrevisRequestOpRequestsFulfilled // Event containing the contract specifics and raw log

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
func (it *BrevisRequestOpRequestsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestOpRequestsFulfilled)
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
		it.Event = new(BrevisRequestOpRequestsFulfilled)
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
func (it *BrevisRequestOpRequestsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestOpRequestsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestOpRequestsFulfilled represents a OpRequestsFulfilled event raised by the BrevisRequest contract.
type BrevisRequestOpRequestsFulfilled struct {
	RequestIds [][32]byte
	QueryURLs  [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOpRequestsFulfilled is a free log retrieval operation binding the contract event 0xa1c3ef280bbac31181188d4beffe370ebae03e08d58a3aa4832b107d540c0a83.
//
// Solidity: event OpRequestsFulfilled(bytes32[] requestIds, bytes[] queryURLs)
func (_BrevisRequest *BrevisRequestFilterer) FilterOpRequestsFulfilled(opts *bind.FilterOpts) (*BrevisRequestOpRequestsFulfilledIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "OpRequestsFulfilled")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestOpRequestsFulfilledIterator{contract: _BrevisRequest.contract, event: "OpRequestsFulfilled", logs: logs, sub: sub}, nil
}

// WatchOpRequestsFulfilled is a free log subscription operation binding the contract event 0xa1c3ef280bbac31181188d4beffe370ebae03e08d58a3aa4832b107d540c0a83.
//
// Solidity: event OpRequestsFulfilled(bytes32[] requestIds, bytes[] queryURLs)
func (_BrevisRequest *BrevisRequestFilterer) WatchOpRequestsFulfilled(opts *bind.WatchOpts, sink chan<- *BrevisRequestOpRequestsFulfilled) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "OpRequestsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestOpRequestsFulfilled)
				if err := _BrevisRequest.contract.UnpackLog(event, "OpRequestsFulfilled", log); err != nil {
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

// ParseOpRequestsFulfilled is a log parse operation binding the contract event 0xa1c3ef280bbac31181188d4beffe370ebae03e08d58a3aa4832b107d540c0a83.
//
// Solidity: event OpRequestsFulfilled(bytes32[] requestIds, bytes[] queryURLs)
func (_BrevisRequest *BrevisRequestFilterer) ParseOpRequestsFulfilled(log types.Log) (*BrevisRequestOpRequestsFulfilled, error) {
	event := new(BrevisRequestOpRequestsFulfilled)
	if err := _BrevisRequest.contract.UnpackLog(event, "OpRequestsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BrevisRequest contract.
type BrevisRequestOwnershipTransferredIterator struct {
	Event *BrevisRequestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrevisRequestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestOwnershipTransferred)
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
		it.Event = new(BrevisRequestOwnershipTransferred)
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
func (it *BrevisRequestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestOwnershipTransferred represents a OwnershipTransferred event raised by the BrevisRequest contract.
type BrevisRequestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisRequest *BrevisRequestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrevisRequestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestOwnershipTransferredIterator{contract: _BrevisRequest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisRequest *BrevisRequestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrevisRequestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestOwnershipTransferred)
				if err := _BrevisRequest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BrevisRequest *BrevisRequestFilterer) ParseOwnershipTransferred(log types.Log) (*BrevisRequestOwnershipTransferred, error) {
	event := new(BrevisRequestOwnershipTransferred)
	if err := _BrevisRequest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestProofPostIterator is returned from FilterProofPost and is used to iterate over the raw logs and unpacked data for ProofPost events raised by the BrevisRequest contract.
type BrevisRequestProofPostIterator struct {
	Event *BrevisRequestProofPost // Event containing the contract specifics and raw log

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
func (it *BrevisRequestProofPostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestProofPost)
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
		it.Event = new(BrevisRequestProofPost)
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
func (it *BrevisRequestProofPostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestProofPostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestProofPost represents a ProofPost event raised by the BrevisRequest contract.
type BrevisRequestProofPost struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProofPost is a free log retrieval operation binding the contract event 0x799191eb766b210bea2d9052b39f21140ed9fbf3ff2fe5102fbd6bb5bc134e3e.
//
// Solidity: event ProofPost(bytes32 indexed requestId)
func (_BrevisRequest *BrevisRequestFilterer) FilterProofPost(opts *bind.FilterOpts, requestId [][32]byte) (*BrevisRequestProofPostIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "ProofPost", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestProofPostIterator{contract: _BrevisRequest.contract, event: "ProofPost", logs: logs, sub: sub}, nil
}

// WatchProofPost is a free log subscription operation binding the contract event 0x799191eb766b210bea2d9052b39f21140ed9fbf3ff2fe5102fbd6bb5bc134e3e.
//
// Solidity: event ProofPost(bytes32 indexed requestId)
func (_BrevisRequest *BrevisRequestFilterer) WatchProofPost(opts *bind.WatchOpts, sink chan<- *BrevisRequestProofPost, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "ProofPost", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestProofPost)
				if err := _BrevisRequest.contract.UnpackLog(event, "ProofPost", log); err != nil {
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

// ParseProofPost is a log parse operation binding the contract event 0x799191eb766b210bea2d9052b39f21140ed9fbf3ff2fe5102fbd6bb5bc134e3e.
//
// Solidity: event ProofPost(bytes32 indexed requestId)
func (_BrevisRequest *BrevisRequestFilterer) ParseProofPost(log types.Log) (*BrevisRequestProofPost, error) {
	event := new(BrevisRequestProofPost)
	if err := _BrevisRequest.contract.UnpackLog(event, "ProofPost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestQueryDataPostIterator is returned from FilterQueryDataPost and is used to iterate over the raw logs and unpacked data for QueryDataPost events raised by the BrevisRequest contract.
type BrevisRequestQueryDataPostIterator struct {
	Event *BrevisRequestQueryDataPost // Event containing the contract specifics and raw log

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
func (it *BrevisRequestQueryDataPostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestQueryDataPost)
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
		it.Event = new(BrevisRequestQueryDataPost)
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
func (it *BrevisRequestQueryDataPostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestQueryDataPostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestQueryDataPost represents a QueryDataPost event raised by the BrevisRequest contract.
type BrevisRequestQueryDataPost struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQueryDataPost is a free log retrieval operation binding the contract event 0xe4e1e2afaa20a4e7e37593d3fb41a91d496a30ea2b7ccf67a9aa9450e74af4f0.
//
// Solidity: event QueryDataPost(bytes32 indexed requestId)
func (_BrevisRequest *BrevisRequestFilterer) FilterQueryDataPost(opts *bind.FilterOpts, requestId [][32]byte) (*BrevisRequestQueryDataPostIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "QueryDataPost", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BrevisRequestQueryDataPostIterator{contract: _BrevisRequest.contract, event: "QueryDataPost", logs: logs, sub: sub}, nil
}

// WatchQueryDataPost is a free log subscription operation binding the contract event 0xe4e1e2afaa20a4e7e37593d3fb41a91d496a30ea2b7ccf67a9aa9450e74af4f0.
//
// Solidity: event QueryDataPost(bytes32 indexed requestId)
func (_BrevisRequest *BrevisRequestFilterer) WatchQueryDataPost(opts *bind.WatchOpts, sink chan<- *BrevisRequestQueryDataPost, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "QueryDataPost", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestQueryDataPost)
				if err := _BrevisRequest.contract.UnpackLog(event, "QueryDataPost", log); err != nil {
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

// ParseQueryDataPost is a log parse operation binding the contract event 0xe4e1e2afaa20a4e7e37593d3fb41a91d496a30ea2b7ccf67a9aa9450e74af4f0.
//
// Solidity: event QueryDataPost(bytes32 indexed requestId)
func (_BrevisRequest *BrevisRequestFilterer) ParseQueryDataPost(log types.Log) (*BrevisRequestQueryDataPost, error) {
	event := new(BrevisRequestQueryDataPost)
	if err := _BrevisRequest.contract.UnpackLog(event, "QueryDataPost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the BrevisRequest contract.
type BrevisRequestRequestFulfilledIterator struct {
	Event *BrevisRequestRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *BrevisRequestRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestRequestFulfilled)
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
		it.Event = new(BrevisRequestRequestFulfilled)
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
func (it *BrevisRequestRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestRequestFulfilled represents a RequestFulfilled event raised by the BrevisRequest contract.
type BrevisRequestRequestFulfilled struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 requestId)
func (_BrevisRequest *BrevisRequestFilterer) FilterRequestFulfilled(opts *bind.FilterOpts) (*BrevisRequestRequestFulfilledIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestRequestFulfilledIterator{contract: _BrevisRequest.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 requestId)
func (_BrevisRequest *BrevisRequestFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *BrevisRequestRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestRequestFulfilled)
				if err := _BrevisRequest.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 requestId)
func (_BrevisRequest *BrevisRequestFilterer) ParseRequestFulfilled(log types.Log) (*BrevisRequestRequestFulfilled, error) {
	event := new(BrevisRequestRequestFulfilled)
	if err := _BrevisRequest.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestRequestRefundedIterator is returned from FilterRequestRefunded and is used to iterate over the raw logs and unpacked data for RequestRefunded events raised by the BrevisRequest contract.
type BrevisRequestRequestRefundedIterator struct {
	Event *BrevisRequestRequestRefunded // Event containing the contract specifics and raw log

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
func (it *BrevisRequestRequestRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestRequestRefunded)
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
		it.Event = new(BrevisRequestRequestRefunded)
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
func (it *BrevisRequestRequestRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestRequestRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestRequestRefunded represents a RequestRefunded event raised by the BrevisRequest contract.
type BrevisRequestRequestRefunded struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestRefunded is a free log retrieval operation binding the contract event 0xfea410cb461deba9fe807dde02d6641d82e1bf09ecc88ecfa0f2ffadf2a1fdfe.
//
// Solidity: event RequestRefunded(bytes32 requestId)
func (_BrevisRequest *BrevisRequestFilterer) FilterRequestRefunded(opts *bind.FilterOpts) (*BrevisRequestRequestRefundedIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "RequestRefunded")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestRequestRefundedIterator{contract: _BrevisRequest.contract, event: "RequestRefunded", logs: logs, sub: sub}, nil
}

// WatchRequestRefunded is a free log subscription operation binding the contract event 0xfea410cb461deba9fe807dde02d6641d82e1bf09ecc88ecfa0f2ffadf2a1fdfe.
//
// Solidity: event RequestRefunded(bytes32 requestId)
func (_BrevisRequest *BrevisRequestFilterer) WatchRequestRefunded(opts *bind.WatchOpts, sink chan<- *BrevisRequestRequestRefunded) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "RequestRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestRequestRefunded)
				if err := _BrevisRequest.contract.UnpackLog(event, "RequestRefunded", log); err != nil {
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

// ParseRequestRefunded is a log parse operation binding the contract event 0xfea410cb461deba9fe807dde02d6641d82e1bf09ecc88ecfa0f2ffadf2a1fdfe.
//
// Solidity: event RequestRefunded(bytes32 requestId)
func (_BrevisRequest *BrevisRequestFilterer) ParseRequestRefunded(log types.Log) (*BrevisRequestRequestRefunded, error) {
	event := new(BrevisRequestRequestRefunded)
	if err := _BrevisRequest.contract.UnpackLog(event, "RequestRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the BrevisRequest contract.
type BrevisRequestRequestSentIterator struct {
	Event *BrevisRequestRequestSent // Event containing the contract specifics and raw log

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
func (it *BrevisRequestRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestRequestSent)
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
		it.Event = new(BrevisRequestRequestSent)
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
func (it *BrevisRequestRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestRequestSent represents a RequestSent event raised by the BrevisRequest contract.
type BrevisRequestRequestSent struct {
	RequestId [32]byte
	Sender    common.Address
	Fee       *big.Int
	Callback  common.Address
	Option    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x8351e4cc62a96a972da26714bfbb019926334976bb613c36e0dc72b69cf8f564.
//
// Solidity: event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback, uint8 option)
func (_BrevisRequest *BrevisRequestFilterer) FilterRequestSent(opts *bind.FilterOpts) (*BrevisRequestRequestSentIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestRequestSentIterator{contract: _BrevisRequest.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x8351e4cc62a96a972da26714bfbb019926334976bb613c36e0dc72b69cf8f564.
//
// Solidity: event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback, uint8 option)
func (_BrevisRequest *BrevisRequestFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *BrevisRequestRequestSent) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestRequestSent)
				if err := _BrevisRequest.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x8351e4cc62a96a972da26714bfbb019926334976bb613c36e0dc72b69cf8f564.
//
// Solidity: event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback, uint8 option)
func (_BrevisRequest *BrevisRequestFilterer) ParseRequestSent(log types.Log) (*BrevisRequestRequestSent, error) {
	event := new(BrevisRequestRequestSent)
	if err := _BrevisRequest.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestRequestTimeoutUpdatedIterator is returned from FilterRequestTimeoutUpdated and is used to iterate over the raw logs and unpacked data for RequestTimeoutUpdated events raised by the BrevisRequest contract.
type BrevisRequestRequestTimeoutUpdatedIterator struct {
	Event *BrevisRequestRequestTimeoutUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisRequestRequestTimeoutUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestRequestTimeoutUpdated)
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
		it.Event = new(BrevisRequestRequestTimeoutUpdated)
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
func (it *BrevisRequestRequestTimeoutUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestRequestTimeoutUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestRequestTimeoutUpdated represents a RequestTimeoutUpdated event raised by the BrevisRequest contract.
type BrevisRequestRequestTimeoutUpdated struct {
	From *big.Int
	To   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestTimeoutUpdated is a free log retrieval operation binding the contract event 0x87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a63.
//
// Solidity: event RequestTimeoutUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) FilterRequestTimeoutUpdated(opts *bind.FilterOpts) (*BrevisRequestRequestTimeoutUpdatedIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "RequestTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestRequestTimeoutUpdatedIterator{contract: _BrevisRequest.contract, event: "RequestTimeoutUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestTimeoutUpdated is a free log subscription operation binding the contract event 0x87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a63.
//
// Solidity: event RequestTimeoutUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) WatchRequestTimeoutUpdated(opts *bind.WatchOpts, sink chan<- *BrevisRequestRequestTimeoutUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "RequestTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestRequestTimeoutUpdated)
				if err := _BrevisRequest.contract.UnpackLog(event, "RequestTimeoutUpdated", log); err != nil {
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

// ParseRequestTimeoutUpdated is a log parse operation binding the contract event 0x87a73c061f18ffd513249d1d727921e40e348948b01e2979efb36ef4f5204a63.
//
// Solidity: event RequestTimeoutUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) ParseRequestTimeoutUpdated(log types.Log) (*BrevisRequestRequestTimeoutUpdated, error) {
	event := new(BrevisRequestRequestTimeoutUpdated)
	if err := _BrevisRequest.contract.UnpackLog(event, "RequestTimeoutUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisRequestResponseTimeoutUpdatedIterator is returned from FilterResponseTimeoutUpdated and is used to iterate over the raw logs and unpacked data for ResponseTimeoutUpdated events raised by the BrevisRequest contract.
type BrevisRequestResponseTimeoutUpdatedIterator struct {
	Event *BrevisRequestResponseTimeoutUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisRequestResponseTimeoutUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisRequestResponseTimeoutUpdated)
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
		it.Event = new(BrevisRequestResponseTimeoutUpdated)
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
func (it *BrevisRequestResponseTimeoutUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisRequestResponseTimeoutUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisRequestResponseTimeoutUpdated represents a ResponseTimeoutUpdated event raised by the BrevisRequest contract.
type BrevisRequestResponseTimeoutUpdated struct {
	From *big.Int
	To   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResponseTimeoutUpdated is a free log retrieval operation binding the contract event 0x86fe7fc31f35681a1ed77325f0cf24935a5d25b1861e7ce9ceed9cb67f222270.
//
// Solidity: event ResponseTimeoutUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) FilterResponseTimeoutUpdated(opts *bind.FilterOpts) (*BrevisRequestResponseTimeoutUpdatedIterator, error) {

	logs, sub, err := _BrevisRequest.contract.FilterLogs(opts, "ResponseTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisRequestResponseTimeoutUpdatedIterator{contract: _BrevisRequest.contract, event: "ResponseTimeoutUpdated", logs: logs, sub: sub}, nil
}

// WatchResponseTimeoutUpdated is a free log subscription operation binding the contract event 0x86fe7fc31f35681a1ed77325f0cf24935a5d25b1861e7ce9ceed9cb67f222270.
//
// Solidity: event ResponseTimeoutUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) WatchResponseTimeoutUpdated(opts *bind.WatchOpts, sink chan<- *BrevisRequestResponseTimeoutUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisRequest.contract.WatchLogs(opts, "ResponseTimeoutUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisRequestResponseTimeoutUpdated)
				if err := _BrevisRequest.contract.UnpackLog(event, "ResponseTimeoutUpdated", log); err != nil {
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

// ParseResponseTimeoutUpdated is a log parse operation binding the contract event 0x86fe7fc31f35681a1ed77325f0cf24935a5d25b1861e7ce9ceed9cb67f222270.
//
// Solidity: event ResponseTimeoutUpdated(uint256 from, uint256 to)
func (_BrevisRequest *BrevisRequestFilterer) ParseResponseTimeoutUpdated(log types.Log) (*BrevisRequestResponseTimeoutUpdated, error) {
	event := new(BrevisRequestResponseTimeoutUpdated)
	if err := _BrevisRequest.contract.UnpackLog(event, "ResponseTimeoutUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeVaultMetaData contains all meta data concerning the FeeVault contract.
var FeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080346100a457601f6104bd38819003918201601f19168301916001600160401b038311848410176100a8578084926020946040528339810103126100a457516001600160a01b0390818116908190036100a4575f5460018060a01b03199033828216175f55604051933391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3600154161760015561040090816100bd8239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040818152600480361015610020575b505050361561001e575f80fd5b005b5f92833560e01c9081637ff7b0d214610224575080638da5cb5b146101fe578063a42dce8014610178578063c415b95c1461014c5763f2fde38b03610011573461014857602036600319011261014857610078610365565b908354906001600160a01b038083169361009333861461037f565b169384156100df57505073ffffffffffffffffffffffffffffffffffffffff1916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b8280fd5b5050346101745781600319360112610174576020906001600160a01b03600154169051908152f35b5080fd5b505034610174576020366003190112610174577f5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38906101b5610365565b6001600160a01b036101cb81865416331461037f565b806001549216908173ffffffffffffffffffffffffffffffffffffffff198416176001558351921682526020820152a180f35b5050346101745781600319360112610174576001600160a01b0360209254169051908152f35b8493915034610361578160031936011261036157602435906001600160a01b03808316830361035d5760015416330361031c5750838080809386359061c350f13d156103175767ffffffffffffffff3d81811161030457835191601f8201601f19908116603f01168301908111838210176102f157845281528460203d92013e5b156102ae578280f35b906020606492519162461bcd60e51b8352820152601260248201527f73656e64206e6174697665206661696c656400000000000000000000000000006044820152fd5b634e487b7160e01b875260418652602487fd5b634e487b7160e01b865260418552602486fd5b6102a5565b62461bcd60e51b8152602084820152601160248201527f6e6f742066656520636f6c6c6563746f720000000000000000000000000000006044820152606490fd5b8580fd5b8380fd5b600435906001600160a01b038216820361037b57565b5f80fd5b1561038657565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfea2646970667358221220d3c32a4bd5a9586f200ad879472653aed7e6c9294e86e396564359104800f07364736f6c63430008140033",
}

// FeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeVaultMetaData.ABI instead.
var FeeVaultABI = FeeVaultMetaData.ABI

// FeeVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeVaultMetaData.Bin instead.
var FeeVaultBin = FeeVaultMetaData.Bin

// DeployFeeVault deploys a new Ethereum contract, binding an instance of FeeVault to it.
func DeployFeeVault(auth *bind.TransactOpts, backend bind.ContractBackend, _feeCollector common.Address) (common.Address, *types.Transaction, *FeeVault, error) {
	parsed, err := FeeVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeVaultBin), backend, _feeCollector)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeVault{FeeVaultCaller: FeeVaultCaller{contract: contract}, FeeVaultTransactor: FeeVaultTransactor{contract: contract}, FeeVaultFilterer: FeeVaultFilterer{contract: contract}}, nil
}

// FeeVault is an auto generated Go binding around an Ethereum contract.
type FeeVault struct {
	FeeVaultCaller     // Read-only binding to the contract
	FeeVaultTransactor // Write-only binding to the contract
	FeeVaultFilterer   // Log filterer for contract events
}

// FeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeVaultSession struct {
	Contract     *FeeVault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeVaultCallerSession struct {
	Contract *FeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeVaultTransactorSession struct {
	Contract     *FeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeVaultRaw struct {
	Contract *FeeVault // Generic contract binding to access the raw methods on
}

// FeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeVaultCallerRaw struct {
	Contract *FeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// FeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeVaultTransactorRaw struct {
	Contract *FeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeVault creates a new instance of FeeVault, bound to a specific deployed contract.
func NewFeeVault(address common.Address, backend bind.ContractBackend) (*FeeVault, error) {
	contract, err := bindFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeVault{FeeVaultCaller: FeeVaultCaller{contract: contract}, FeeVaultTransactor: FeeVaultTransactor{contract: contract}, FeeVaultFilterer: FeeVaultFilterer{contract: contract}}, nil
}

// NewFeeVaultCaller creates a new read-only instance of FeeVault, bound to a specific deployed contract.
func NewFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*FeeVaultCaller, error) {
	contract, err := bindFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeVaultCaller{contract: contract}, nil
}

// NewFeeVaultTransactor creates a new write-only instance of FeeVault, bound to a specific deployed contract.
func NewFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeVaultTransactor, error) {
	contract, err := bindFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeVaultTransactor{contract: contract}, nil
}

// NewFeeVaultFilterer creates a new log filterer instance of FeeVault, bound to a specific deployed contract.
func NewFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeVaultFilterer, error) {
	contract, err := bindFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeVaultFilterer{contract: contract}, nil
}

// bindFeeVault binds a generic wrapper to an already deployed contract.
func bindFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeVault *FeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeVault.Contract.FeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeVault *FeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.Contract.FeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeVault *FeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeVault.Contract.FeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeVault *FeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeVault *FeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeVault *FeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeVault.Contract.contract.Transact(opts, method, params...)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_FeeVault *FeeVaultCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeVault.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_FeeVault *FeeVaultSession) FeeCollector() (common.Address, error) {
	return _FeeVault.Contract.FeeCollector(&_FeeVault.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_FeeVault *FeeVaultCallerSession) FeeCollector() (common.Address, error) {
	return _FeeVault.Contract.FeeCollector(&_FeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeVault *FeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeVault *FeeVaultSession) Owner() (common.Address, error) {
	return _FeeVault.Contract.Owner(&_FeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeVault *FeeVaultCallerSession) Owner() (common.Address, error) {
	return _FeeVault.Contract.Owner(&_FeeVault.CallOpts)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_FeeVault *FeeVaultTransactor) CollectFee(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "collectFee", _amount, _to)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_FeeVault *FeeVaultSession) CollectFee(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.CollectFee(&_FeeVault.TransactOpts, _amount, _to)
}

// CollectFee is a paid mutator transaction binding the contract method 0x7ff7b0d2.
//
// Solidity: function collectFee(uint256 _amount, address _to) returns()
func (_FeeVault *FeeVaultTransactorSession) CollectFee(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.CollectFee(&_FeeVault.TransactOpts, _amount, _to)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_FeeVault *FeeVaultTransactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_FeeVault *FeeVaultSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.SetFeeCollector(&_FeeVault.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_FeeVault *FeeVaultTransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.SetFeeCollector(&_FeeVault.TransactOpts, _feeCollector)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeVault *FeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeVault *FeeVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.TransferOwnership(&_FeeVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeVault *FeeVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.TransferOwnership(&_FeeVault.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeVault *FeeVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeVault *FeeVaultSession) Receive() (*types.Transaction, error) {
	return _FeeVault.Contract.Receive(&_FeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeVault *FeeVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _FeeVault.Contract.Receive(&_FeeVault.TransactOpts)
}

// FeeVaultFeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the FeeVault contract.
type FeeVaultFeeCollectorUpdatedIterator struct {
	Event *FeeVaultFeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *FeeVaultFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeVaultFeeCollectorUpdated)
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
		it.Event = new(FeeVaultFeeCollectorUpdated)
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
func (it *FeeVaultFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeVaultFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeVaultFeeCollectorUpdated represents a FeeCollectorUpdated event raised by the FeeVault contract.
type FeeVaultFeeCollectorUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_FeeVault *FeeVaultFilterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts) (*FeeVaultFeeCollectorUpdatedIterator, error) {

	logs, sub, err := _FeeVault.contract.FilterLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeVaultFeeCollectorUpdatedIterator{contract: _FeeVault.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_FeeVault *FeeVaultFilterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *FeeVaultFeeCollectorUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeVault.contract.WatchLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeVaultFeeCollectorUpdated)
				if err := _FeeVault.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
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

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0x5d16ad41baeb009cd23eb8f6c7cde5c2e0cd5acf4a33926ab488875c37c37f38.
//
// Solidity: event FeeCollectorUpdated(address from, address to)
func (_FeeVault *FeeVaultFilterer) ParseFeeCollectorUpdated(log types.Log) (*FeeVaultFeeCollectorUpdated, error) {
	event := new(FeeVaultFeeCollectorUpdated)
	if err := _FeeVault.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeVault contract.
type FeeVaultOwnershipTransferredIterator struct {
	Event *FeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeVaultOwnershipTransferred)
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
		it.Event = new(FeeVaultOwnershipTransferred)
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
func (it *FeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the FeeVault contract.
type FeeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeVaultOwnershipTransferredIterator{contract: _FeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeVaultOwnershipTransferred)
				if err := _FeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeeVault *FeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*FeeVaultOwnershipTransferred, error) {
	event := new(FeeVaultOwnershipTransferred)
	if err := _FeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisAppMetaData contains all meta data concerning the IBrevisApp contract.
var IBrevisAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_appCircuitOutput\",\"type\":\"bytes\"}],\"name\":\"brevisCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBrevisAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IBrevisAppMetaData.ABI instead.
var IBrevisAppABI = IBrevisAppMetaData.ABI

// IBrevisApp is an auto generated Go binding around an Ethereum contract.
type IBrevisApp struct {
	IBrevisAppCaller     // Read-only binding to the contract
	IBrevisAppTransactor // Write-only binding to the contract
	IBrevisAppFilterer   // Log filterer for contract events
}

// IBrevisAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBrevisAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBrevisAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBrevisAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBrevisAppSession struct {
	Contract     *IBrevisApp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBrevisAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBrevisAppCallerSession struct {
	Contract *IBrevisAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IBrevisAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBrevisAppTransactorSession struct {
	Contract     *IBrevisAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBrevisAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBrevisAppRaw struct {
	Contract *IBrevisApp // Generic contract binding to access the raw methods on
}

// IBrevisAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBrevisAppCallerRaw struct {
	Contract *IBrevisAppCaller // Generic read-only contract binding to access the raw methods on
}

// IBrevisAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBrevisAppTransactorRaw struct {
	Contract *IBrevisAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBrevisApp creates a new instance of IBrevisApp, bound to a specific deployed contract.
func NewIBrevisApp(address common.Address, backend bind.ContractBackend) (*IBrevisApp, error) {
	contract, err := bindIBrevisApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBrevisApp{IBrevisAppCaller: IBrevisAppCaller{contract: contract}, IBrevisAppTransactor: IBrevisAppTransactor{contract: contract}, IBrevisAppFilterer: IBrevisAppFilterer{contract: contract}}, nil
}

// NewIBrevisAppCaller creates a new read-only instance of IBrevisApp, bound to a specific deployed contract.
func NewIBrevisAppCaller(address common.Address, caller bind.ContractCaller) (*IBrevisAppCaller, error) {
	contract, err := bindIBrevisApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisAppCaller{contract: contract}, nil
}

// NewIBrevisAppTransactor creates a new write-only instance of IBrevisApp, bound to a specific deployed contract.
func NewIBrevisAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IBrevisAppTransactor, error) {
	contract, err := bindIBrevisApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisAppTransactor{contract: contract}, nil
}

// NewIBrevisAppFilterer creates a new log filterer instance of IBrevisApp, bound to a specific deployed contract.
func NewIBrevisAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IBrevisAppFilterer, error) {
	contract, err := bindIBrevisApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBrevisAppFilterer{contract: contract}, nil
}

// bindIBrevisApp binds a generic wrapper to an already deployed contract.
func bindIBrevisApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBrevisAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisApp *IBrevisAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisApp.Contract.IBrevisAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisApp *IBrevisAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisApp.Contract.IBrevisAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisApp *IBrevisAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisApp.Contract.IBrevisAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisApp *IBrevisAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisApp *IBrevisAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisApp *IBrevisAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisApp.Contract.contract.Transact(opts, method, params...)
}

// BrevisCallback is a paid mutator transaction binding the contract method 0x79d6b6a2.
//
// Solidity: function brevisCallback(bytes32 _requestId, bytes _appCircuitOutput) returns()
func (_IBrevisApp *IBrevisAppTransactor) BrevisCallback(opts *bind.TransactOpts, _requestId [32]byte, _appCircuitOutput []byte) (*types.Transaction, error) {
	return _IBrevisApp.contract.Transact(opts, "brevisCallback", _requestId, _appCircuitOutput)
}

// BrevisCallback is a paid mutator transaction binding the contract method 0x79d6b6a2.
//
// Solidity: function brevisCallback(bytes32 _requestId, bytes _appCircuitOutput) returns()
func (_IBrevisApp *IBrevisAppSession) BrevisCallback(_requestId [32]byte, _appCircuitOutput []byte) (*types.Transaction, error) {
	return _IBrevisApp.Contract.BrevisCallback(&_IBrevisApp.TransactOpts, _requestId, _appCircuitOutput)
}

// BrevisCallback is a paid mutator transaction binding the contract method 0x79d6b6a2.
//
// Solidity: function brevisCallback(bytes32 _requestId, bytes _appCircuitOutput) returns()
func (_IBrevisApp *IBrevisAppTransactorSession) BrevisCallback(_requestId [32]byte, _appCircuitOutput []byte) (*types.Transaction, error) {
	return _IBrevisApp.Contract.BrevisCallback(&_IBrevisApp.TransactOpts, _requestId, _appCircuitOutput)
}

// IBrevisEigenMetaData contains all meta data concerning the IBrevisEigen contract.
var IBrevisEigenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"reqIds\",\"type\":\"bytes32[]\"}],\"name\":\"mustVerified\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBrevisEigenABI is the input ABI used to generate the binding from.
// Deprecated: Use IBrevisEigenMetaData.ABI instead.
var IBrevisEigenABI = IBrevisEigenMetaData.ABI

// IBrevisEigen is an auto generated Go binding around an Ethereum contract.
type IBrevisEigen struct {
	IBrevisEigenCaller     // Read-only binding to the contract
	IBrevisEigenTransactor // Write-only binding to the contract
	IBrevisEigenFilterer   // Log filterer for contract events
}

// IBrevisEigenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBrevisEigenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisEigenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBrevisEigenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisEigenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBrevisEigenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisEigenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBrevisEigenSession struct {
	Contract     *IBrevisEigen     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBrevisEigenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBrevisEigenCallerSession struct {
	Contract *IBrevisEigenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBrevisEigenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBrevisEigenTransactorSession struct {
	Contract     *IBrevisEigenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBrevisEigenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBrevisEigenRaw struct {
	Contract *IBrevisEigen // Generic contract binding to access the raw methods on
}

// IBrevisEigenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBrevisEigenCallerRaw struct {
	Contract *IBrevisEigenCaller // Generic read-only contract binding to access the raw methods on
}

// IBrevisEigenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBrevisEigenTransactorRaw struct {
	Contract *IBrevisEigenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBrevisEigen creates a new instance of IBrevisEigen, bound to a specific deployed contract.
func NewIBrevisEigen(address common.Address, backend bind.ContractBackend) (*IBrevisEigen, error) {
	contract, err := bindIBrevisEigen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBrevisEigen{IBrevisEigenCaller: IBrevisEigenCaller{contract: contract}, IBrevisEigenTransactor: IBrevisEigenTransactor{contract: contract}, IBrevisEigenFilterer: IBrevisEigenFilterer{contract: contract}}, nil
}

// NewIBrevisEigenCaller creates a new read-only instance of IBrevisEigen, bound to a specific deployed contract.
func NewIBrevisEigenCaller(address common.Address, caller bind.ContractCaller) (*IBrevisEigenCaller, error) {
	contract, err := bindIBrevisEigen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisEigenCaller{contract: contract}, nil
}

// NewIBrevisEigenTransactor creates a new write-only instance of IBrevisEigen, bound to a specific deployed contract.
func NewIBrevisEigenTransactor(address common.Address, transactor bind.ContractTransactor) (*IBrevisEigenTransactor, error) {
	contract, err := bindIBrevisEigen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisEigenTransactor{contract: contract}, nil
}

// NewIBrevisEigenFilterer creates a new log filterer instance of IBrevisEigen, bound to a specific deployed contract.
func NewIBrevisEigenFilterer(address common.Address, filterer bind.ContractFilterer) (*IBrevisEigenFilterer, error) {
	contract, err := bindIBrevisEigen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBrevisEigenFilterer{contract: contract}, nil
}

// bindIBrevisEigen binds a generic wrapper to an already deployed contract.
func bindIBrevisEigen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBrevisEigenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisEigen *IBrevisEigenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisEigen.Contract.IBrevisEigenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisEigen *IBrevisEigenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisEigen.Contract.IBrevisEigenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisEigen *IBrevisEigenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisEigen.Contract.IBrevisEigenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisEigen *IBrevisEigenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisEigen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisEigen *IBrevisEigenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisEigen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisEigen *IBrevisEigenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisEigen.Contract.contract.Transact(opts, method, params...)
}

// MustVerified is a paid mutator transaction binding the contract method 0xd1e34715.
//
// Solidity: function mustVerified(bytes32[] reqIds) returns()
func (_IBrevisEigen *IBrevisEigenTransactor) MustVerified(opts *bind.TransactOpts, reqIds [][32]byte) (*types.Transaction, error) {
	return _IBrevisEigen.contract.Transact(opts, "mustVerified", reqIds)
}

// MustVerified is a paid mutator transaction binding the contract method 0xd1e34715.
//
// Solidity: function mustVerified(bytes32[] reqIds) returns()
func (_IBrevisEigen *IBrevisEigenSession) MustVerified(reqIds [][32]byte) (*types.Transaction, error) {
	return _IBrevisEigen.Contract.MustVerified(&_IBrevisEigen.TransactOpts, reqIds)
}

// MustVerified is a paid mutator transaction binding the contract method 0xd1e34715.
//
// Solidity: function mustVerified(bytes32[] reqIds) returns()
func (_IBrevisEigen *IBrevisEigenTransactorSession) MustVerified(reqIds [][32]byte) (*types.Transaction, error) {
	return _IBrevisEigen.Contract.MustVerified(&_IBrevisEigen.TransactOpts, reqIds)
}

// IBrevisProofMetaData contains all meta data concerning the IBrevisProof contract.
var IBrevisProofMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"getProofAppData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"hasProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"submitOpResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_proofWithPubInputs\",\"type\":\"bytes\"}],\"name\":\"submitProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blkNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiptIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"logIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valueFromTopic\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"valueIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"logTopic0\",\"type\":\"bytes32\"}],\"internalType\":\"structBrevis.LogInfo[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"internalType\":\"structBrevis.ReceiptInfo[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slotValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"internalType\":\"structBrevis.StorageInfo[]\",\"name\":\"stores\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfRawTxData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"internalType\":\"structBrevis.TransactionInfo[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"internalType\":\"structBrevis.ExtractInfos\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"validateOpRequest\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBrevisProofABI is the input ABI used to generate the binding from.
// Deprecated: Use IBrevisProofMetaData.ABI instead.
var IBrevisProofABI = IBrevisProofMetaData.ABI

// IBrevisProof is an auto generated Go binding around an Ethereum contract.
type IBrevisProof struct {
	IBrevisProofCaller     // Read-only binding to the contract
	IBrevisProofTransactor // Write-only binding to the contract
	IBrevisProofFilterer   // Log filterer for contract events
}

// IBrevisProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBrevisProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBrevisProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBrevisProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBrevisProofSession struct {
	Contract     *IBrevisProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBrevisProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBrevisProofCallerSession struct {
	Contract *IBrevisProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBrevisProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBrevisProofTransactorSession struct {
	Contract     *IBrevisProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBrevisProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBrevisProofRaw struct {
	Contract *IBrevisProof // Generic contract binding to access the raw methods on
}

// IBrevisProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBrevisProofCallerRaw struct {
	Contract *IBrevisProofCaller // Generic read-only contract binding to access the raw methods on
}

// IBrevisProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBrevisProofTransactorRaw struct {
	Contract *IBrevisProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBrevisProof creates a new instance of IBrevisProof, bound to a specific deployed contract.
func NewIBrevisProof(address common.Address, backend bind.ContractBackend) (*IBrevisProof, error) {
	contract, err := bindIBrevisProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBrevisProof{IBrevisProofCaller: IBrevisProofCaller{contract: contract}, IBrevisProofTransactor: IBrevisProofTransactor{contract: contract}, IBrevisProofFilterer: IBrevisProofFilterer{contract: contract}}, nil
}

// NewIBrevisProofCaller creates a new read-only instance of IBrevisProof, bound to a specific deployed contract.
func NewIBrevisProofCaller(address common.Address, caller bind.ContractCaller) (*IBrevisProofCaller, error) {
	contract, err := bindIBrevisProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisProofCaller{contract: contract}, nil
}

// NewIBrevisProofTransactor creates a new write-only instance of IBrevisProof, bound to a specific deployed contract.
func NewIBrevisProofTransactor(address common.Address, transactor bind.ContractTransactor) (*IBrevisProofTransactor, error) {
	contract, err := bindIBrevisProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisProofTransactor{contract: contract}, nil
}

// NewIBrevisProofFilterer creates a new log filterer instance of IBrevisProof, bound to a specific deployed contract.
func NewIBrevisProofFilterer(address common.Address, filterer bind.ContractFilterer) (*IBrevisProofFilterer, error) {
	contract, err := bindIBrevisProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBrevisProofFilterer{contract: contract}, nil
}

// bindIBrevisProof binds a generic wrapper to an already deployed contract.
func bindIBrevisProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBrevisProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisProof *IBrevisProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisProof.Contract.IBrevisProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisProof *IBrevisProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisProof.Contract.IBrevisProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisProof *IBrevisProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisProof.Contract.IBrevisProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisProof *IBrevisProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisProof *IBrevisProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisProof *IBrevisProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisProof.Contract.contract.Transact(opts, method, params...)
}

// GetProofAppData is a free data retrieval call binding the contract method 0x5984fb92.
//
// Solidity: function getProofAppData(bytes32 _requestId) view returns(bytes32, bytes32)
func (_IBrevisProof *IBrevisProofCaller) GetProofAppData(opts *bind.CallOpts, _requestId [32]byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _IBrevisProof.contract.Call(opts, &out, "getProofAppData", _requestId)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetProofAppData is a free data retrieval call binding the contract method 0x5984fb92.
//
// Solidity: function getProofAppData(bytes32 _requestId) view returns(bytes32, bytes32)
func (_IBrevisProof *IBrevisProofSession) GetProofAppData(_requestId [32]byte) ([32]byte, [32]byte, error) {
	return _IBrevisProof.Contract.GetProofAppData(&_IBrevisProof.CallOpts, _requestId)
}

// GetProofAppData is a free data retrieval call binding the contract method 0x5984fb92.
//
// Solidity: function getProofAppData(bytes32 _requestId) view returns(bytes32, bytes32)
func (_IBrevisProof *IBrevisProofCallerSession) GetProofAppData(_requestId [32]byte) ([32]byte, [32]byte, error) {
	return _IBrevisProof.Contract.GetProofAppData(&_IBrevisProof.CallOpts, _requestId)
}

// HasProof is a free data retrieval call binding the contract method 0xe3d1e6d6.
//
// Solidity: function hasProof(bytes32 _requestId) view returns(bool)
func (_IBrevisProof *IBrevisProofCaller) HasProof(opts *bind.CallOpts, _requestId [32]byte) (bool, error) {
	var out []interface{}
	err := _IBrevisProof.contract.Call(opts, &out, "hasProof", _requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasProof is a free data retrieval call binding the contract method 0xe3d1e6d6.
//
// Solidity: function hasProof(bytes32 _requestId) view returns(bool)
func (_IBrevisProof *IBrevisProofSession) HasProof(_requestId [32]byte) (bool, error) {
	return _IBrevisProof.Contract.HasProof(&_IBrevisProof.CallOpts, _requestId)
}

// HasProof is a free data retrieval call binding the contract method 0xe3d1e6d6.
//
// Solidity: function hasProof(bytes32 _requestId) view returns(bool)
func (_IBrevisProof *IBrevisProofCallerSession) HasProof(_requestId [32]byte) (bool, error) {
	return _IBrevisProof.Contract.HasProof(&_IBrevisProof.CallOpts, _requestId)
}

// ValidateOpRequest is a free data retrieval call binding the contract method 0x2792d466.
//
// Solidity: function validateOpRequest(bytes32 _requestId, uint64 _chainId, ((uint64,uint64,(uint64,bytes32,bool,uint64,address,bytes32)[])[],(bytes32,address,bytes32,bytes32,uint64)[],(bytes32,bytes32,bytes32,uint64)[]) _info) view returns()
func (_IBrevisProof *IBrevisProofCaller) ValidateOpRequest(opts *bind.CallOpts, _requestId [32]byte, _chainId uint64, _info BrevisExtractInfos) error {
	var out []interface{}
	err := _IBrevisProof.contract.Call(opts, &out, "validateOpRequest", _requestId, _chainId, _info)

	if err != nil {
		return err
	}

	return err

}

// ValidateOpRequest is a free data retrieval call binding the contract method 0x2792d466.
//
// Solidity: function validateOpRequest(bytes32 _requestId, uint64 _chainId, ((uint64,uint64,(uint64,bytes32,bool,uint64,address,bytes32)[])[],(bytes32,address,bytes32,bytes32,uint64)[],(bytes32,bytes32,bytes32,uint64)[]) _info) view returns()
func (_IBrevisProof *IBrevisProofSession) ValidateOpRequest(_requestId [32]byte, _chainId uint64, _info BrevisExtractInfos) error {
	return _IBrevisProof.Contract.ValidateOpRequest(&_IBrevisProof.CallOpts, _requestId, _chainId, _info)
}

// ValidateOpRequest is a free data retrieval call binding the contract method 0x2792d466.
//
// Solidity: function validateOpRequest(bytes32 _requestId, uint64 _chainId, ((uint64,uint64,(uint64,bytes32,bool,uint64,address,bytes32)[])[],(bytes32,address,bytes32,bytes32,uint64)[],(bytes32,bytes32,bytes32,uint64)[]) _info) view returns()
func (_IBrevisProof *IBrevisProofCallerSession) ValidateOpRequest(_requestId [32]byte, _chainId uint64, _info BrevisExtractInfos) error {
	return _IBrevisProof.Contract.ValidateOpRequest(&_IBrevisProof.CallOpts, _requestId, _chainId, _info)
}

// SubmitOpResult is a paid mutator transaction binding the contract method 0xe26d07d3.
//
// Solidity: function submitOpResult(bytes32 _requestId) returns()
func (_IBrevisProof *IBrevisProofTransactor) SubmitOpResult(opts *bind.TransactOpts, _requestId [32]byte) (*types.Transaction, error) {
	return _IBrevisProof.contract.Transact(opts, "submitOpResult", _requestId)
}

// SubmitOpResult is a paid mutator transaction binding the contract method 0xe26d07d3.
//
// Solidity: function submitOpResult(bytes32 _requestId) returns()
func (_IBrevisProof *IBrevisProofSession) SubmitOpResult(_requestId [32]byte) (*types.Transaction, error) {
	return _IBrevisProof.Contract.SubmitOpResult(&_IBrevisProof.TransactOpts, _requestId)
}

// SubmitOpResult is a paid mutator transaction binding the contract method 0xe26d07d3.
//
// Solidity: function submitOpResult(bytes32 _requestId) returns()
func (_IBrevisProof *IBrevisProofTransactorSession) SubmitOpResult(_requestId [32]byte) (*types.Transaction, error) {
	return _IBrevisProof.Contract.SubmitOpResult(&_IBrevisProof.TransactOpts, _requestId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xe0440953.
//
// Solidity: function submitProof(uint64 _chainId, bytes _proofWithPubInputs) returns(bytes32 _requestId)
func (_IBrevisProof *IBrevisProofTransactor) SubmitProof(opts *bind.TransactOpts, _chainId uint64, _proofWithPubInputs []byte) (*types.Transaction, error) {
	return _IBrevisProof.contract.Transact(opts, "submitProof", _chainId, _proofWithPubInputs)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xe0440953.
//
// Solidity: function submitProof(uint64 _chainId, bytes _proofWithPubInputs) returns(bytes32 _requestId)
func (_IBrevisProof *IBrevisProofSession) SubmitProof(_chainId uint64, _proofWithPubInputs []byte) (*types.Transaction, error) {
	return _IBrevisProof.Contract.SubmitProof(&_IBrevisProof.TransactOpts, _chainId, _proofWithPubInputs)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xe0440953.
//
// Solidity: function submitProof(uint64 _chainId, bytes _proofWithPubInputs) returns(bytes32 _requestId)
func (_IBrevisProof *IBrevisProofTransactorSession) SubmitProof(_chainId uint64, _proofWithPubInputs []byte) (*types.Transaction, error) {
	return _IBrevisProof.Contract.SubmitProof(&_IBrevisProof.TransactOpts, _chainId, _proofWithPubInputs)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RLPReaderMetaData contains all meta data concerning the RLPReader contract.
var RLPReaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea2646970667358221220f13d5db35952f1f0335d167485b80fdff1af004a8ed3a2a91170c10d1f51493464736f6c63430008140033",
}

// RLPReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPReaderMetaData.ABI instead.
var RLPReaderABI = RLPReaderMetaData.ABI

// RLPReaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RLPReaderMetaData.Bin instead.
var RLPReaderBin = RLPReaderMetaData.Bin

// DeployRLPReader deploys a new Ethereum contract, binding an instance of RLPReader to it.
func DeployRLPReader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPReader, error) {
	parsed, err := RLPReaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RLPReaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// RLPReader is an auto generated Go binding around an Ethereum contract.
type RLPReader struct {
	RLPReaderCaller     // Read-only binding to the contract
	RLPReaderTransactor // Write-only binding to the contract
	RLPReaderFilterer   // Log filterer for contract events
}

// RLPReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPReaderSession struct {
	Contract     *RLPReader        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPReaderCallerSession struct {
	Contract *RLPReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPReaderTransactorSession struct {
	Contract     *RLPReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPReaderRaw struct {
	Contract *RLPReader // Generic contract binding to access the raw methods on
}

// RLPReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPReaderCallerRaw struct {
	Contract *RLPReaderCaller // Generic read-only contract binding to access the raw methods on
}

// RLPReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPReaderTransactorRaw struct {
	Contract *RLPReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPReader creates a new instance of RLPReader, bound to a specific deployed contract.
func NewRLPReader(address common.Address, backend bind.ContractBackend) (*RLPReader, error) {
	contract, err := bindRLPReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPReader{RLPReaderCaller: RLPReaderCaller{contract: contract}, RLPReaderTransactor: RLPReaderTransactor{contract: contract}, RLPReaderFilterer: RLPReaderFilterer{contract: contract}}, nil
}

// NewRLPReaderCaller creates a new read-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderCaller(address common.Address, caller bind.ContractCaller) (*RLPReaderCaller, error) {
	contract, err := bindRLPReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderCaller{contract: contract}, nil
}

// NewRLPReaderTransactor creates a new write-only instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPReaderTransactor, error) {
	contract, err := bindRLPReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPReaderTransactor{contract: contract}, nil
}

// NewRLPReaderFilterer creates a new log filterer instance of RLPReader, bound to a specific deployed contract.
func NewRLPReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPReaderFilterer, error) {
	contract, err := bindRLPReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPReaderFilterer{contract: contract}, nil
}

// bindRLPReader binds a generic wrapper to an already deployed contract.
func bindRLPReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.RLPReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.RLPReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPReader *RLPReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPReader *RLPReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPReader *RLPReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPReader.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea2646970667358221220e4452412b1c842239525529cee107192565885a28e2498d223386853155461ed64736f6c63430008140033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// TxMetaData contains all meta data concerning the Tx contract.
var TxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txRaw\",\"type\":\"bytes\"}],\"name\":\"decodeTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasTipCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFeeCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"internalType\":\"structTx.TxInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001a57610e3a908161001f823930815050f35b5f80fdfe6040608081526004361015610012575f80fd5b5f90813560e01c63dae029d314610027575f80fd5b602090816003193601126103d55767ffffffffffffffff916004358381116103d157366023820112156103d15780600401359384116103d157602481019060248536920101116103d1576102e4946102da926102c992610085610526565b9660ff926100ba6002856100b36100ad61009f8789610591565b356001600160f81b03191690565b60f81c90565b16146105ae565b6100df6100da6100d56100ce8580886105fa565b36916106a7565b6109ff565b610a66565b926101136101056100f86100f2876106dd565b51610b79565b67ffffffffffffffff1690565b67ffffffffffffffff168b52565b6101356101256100f86100f2876106ea565b67ffffffffffffffff168b840152565b6101416100f2856106fa565b898b01526101516100f28561070a565b60608b01526101626100f28561071a565b60808b015261018c61017c6101768661072a565b51610b5b565b6001600160a01b031660a08c0152565b6101986100f28561073a565b60c08b01526101af6101a98561074a565b51610be7565b60e08b01526101c96101c36100f28661075b565b60ff1690565b956101ed6101e26101a96101e76101e26101a98a61076c565b6107a2565b9761077d565b958a6102076102026100ad61009f898861059f565b6107d8565b9582871660010361039b576102336101c361022e6100ad610228858a610637565b90610865565b6107f0565b965b61ffff8816603781116102e85750506102bd939261029d8761028f6102b5958561028961027a6102686102a5998e610648565b9d909361027481610891565b91610656565b929093519c8d968701916108f8565b916108f8565b03601f1981018852876104f5565b859716610814565b60f81b6001600160f81b03191690565b901a91610971565b535b8151910120610981565b6001600160a01b0316610100840152565b5191829182610417565b0390f35b909697945083959391925011155f146103615761035b93828261031d61031761009f61032f9661034d98610591565b98610803565b9261032782610891565b931691610673565b8c5195869491929160f81b6001600160f81b03191690888601610940565b03601f1981018352826104f5565b906102bf565b61035b93828261031d61031761009f61037d9661034d98610591565b8c5195869491929160f01b6001600160f01b03191690888601610905565b6103cb6103c66103c06103ba6103b36101c38c610803565b858a61061a565b90610825565b60f01c90565b610851565b96610235565b8480fd5b8280fd5b91908251928382525f5b848110610403575050825f602080949584010152601f8019910116010190565b6020818301810151848301820152016103e3565b6104bd906020815261043660208201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660408201526040830151606082015260608301516080820152608083015160a082015261048260a084015160c08301906001600160a01b03169052565b60c083015160e082015260e0830151610120906104ac6101009183838601526101408501906103d9565b9401516001600160a01b0316910152565b90565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff8211176104f057604052565b6104c0565b90601f8019910116810190811067ffffffffffffffff8211176104f057604052565b60405190610524826104d4565b565b60405190610120820182811067ffffffffffffffff8211176104f057604052816101005f918281528260208201528260408201528260608201528260808201528260a08201528260c0820152606060e08201520152565b634e487b7160e01b5f52603260045260245ffd5b901561059a5790565b61057d565b906001101561059a5760010190565b156105b557565b60405162461bcd60e51b815260206004820152601660248201527f6e6f7420612044796e616d6963466565547854797065000000000000000000006044820152606490fd5b909291928360011161061657831161061657600101915f190190565b5f80fd5b909291928360021161061657831161061657600201916001190190565b906003116106165760020190600190565b906002116106165790600290565b909291928360031161061657831161061657600301916002190190565b90939293848311610616578411610616578101920390565b67ffffffffffffffff81116104f057601f01601f191660200190565b9291926106b38261068b565b916106c160405193846104f5565b829481845281830111610616578281602093845f960137010152565b80511561059a5760200190565b80516001101561059a5760400190565b80516002101561059a5760600190565b80516003101561059a5760800190565b80516004101561059a5760a00190565b80516005101561059a5760c00190565b80516006101561059a5760e00190565b80516007101561059a576101000190565b80516009101561059a576101400190565b8051600a101561059a576101600190565b8051600b101561059a576101800190565b805182101561059a5760209160051b010190565b6020815191015190602081106107b6575090565b5f199060200360031b1b1690565b634e487b7160e01b5f52601160045260245ffd5b60ff60f6199116019060ff82116107eb57565b6107c4565b60ff6042199116019060ff82116107eb57565b60ff166002019060ff82116107eb57565b60ff1660c0019060ff82116107eb57565b6001600160f01b0319903581811693926002811061084257505050565b60020360031b82901b16169150565b61ffff90811660421901919082116107eb57565b6001600160f81b0319903581811693926001811061088257505050565b60010360031b82901b16169150565b6042198101919082116107eb57565b60bf198101919082116107eb57565b607f198101919082116107eb57565b60200390602082116107eb57565b5f198101919082116107eb57565b60f6198101919082116107eb57565b60b6198101919082116107eb57565b908092918237015f815290565b6001600160f81b0319909116815260f960f81b60018201526001600160f01b031990911660028201526004929182908483013701015f815290565b6001600160f81b03199182168152601f60fb1b6001820152911660028201526003929182908483013701015f815290565b80516001101561059a5760210190565b919260ff8116601b81106109cd575b509160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa156109c2575f5190565b6040513d5f823e3d90fd5b601b9150929192019060ff82116107eb5791906020610990565b604051906109f4826104d4565b5f6020838281520152565b610a076109e7565b50602081519160405192610a1a846104d4565b835201602082015290565b67ffffffffffffffff81116104f05760051b60200190565b90600182018092116107eb57565b919082018092116107eb57565b5f1981146107eb5760010190565b610a6f81610b37565b1561061657610a7d81610c34565b610a8681610a25565b91610a9460405193846104f5565b818352601f19610aa383610a25565b015f5b818110610b20575050610ac7602080920151610ac181610d1a565b90610a4b565b5f905b838210610ad8575050505090565b610b1481610ae8610b1a93610c95565b90610af1610517565b8281528187820152610b03868a61078e565b52610b0e858961078e565b50610a4b565b91610a58565b90610aca565b602090610b2b6109e7565b82828801015201610aa6565b805115610b5657602060c0910151515f1a10610b5257600190565b5f90565b505f90565b601581510361061657610b756001600160a01b0391610b79565b1690565b80518015159081610baf575b501561061657610b9490610bbb565b90519060208110610ba3575090565b6020036101000a900490565b6021915011155f610b85565b906020820191610bcb8351610d1a565b9251908382018092116107eb57519283039283116107eb579190565b80511561061657610bfa6104bd91610bbb565b610c068193929361068b565b92610c1460405194856104f5565b818452601f19610c238361068b565b013660208601378360200190610d87565b805115610b56575f9060208101908151610c4d81610d1a565b81018091116107eb579151905181018091116107eb5791905b828110610c735750905090565b610c7c81610c95565b81018091116107eb57610c8f9091610a58565b90610c66565b80515f1a906080821015610caa575050600190565b60b8821015610cc55750610cc06104bd916108af565b610a3d565b9060c0811015610ce95760b51991600160b783602003016101000a91015104010190565b9060f8821015610d005750610cc06104bd916108a0565b60010151602082900360f7016101000a90040160f5190190565b515f1a6080811015610d2b57505f90565b60b881108015610d62575b15610d415750600190565b60c0811015610d5657610cc06104bd916108e9565b610cc06104bd916108da565b5060c08110158015610d36575060f88110610d36565b601f81116107eb576101000a90565b929091928315610dfe5792915b602093848410610dc957805182528481018091116107eb579381018091116107eb5791601f1981019081116107eb5791610d94565b9193509180610dd757505050565b610deb610de6610df0926108be565b610d78565b6108cc565b905182518216911916179052565b5091505056fea264697066735822122095f0b41ef8e45eae429ac1742a8be864705456cea62fac3fbd17b840b78190e664736f6c63430008140033",
}

// TxABI is the input ABI used to generate the binding from.
// Deprecated: Use TxMetaData.ABI instead.
var TxABI = TxMetaData.ABI

// TxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TxMetaData.Bin instead.
var TxBin = TxMetaData.Bin

// DeployTx deploys a new Ethereum contract, binding an instance of Tx to it.
func DeployTx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tx, error) {
	parsed, err := TxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tx{TxCaller: TxCaller{contract: contract}, TxTransactor: TxTransactor{contract: contract}, TxFilterer: TxFilterer{contract: contract}}, nil
}

// Tx is an auto generated Go binding around an Ethereum contract.
type Tx struct {
	TxCaller     // Read-only binding to the contract
	TxTransactor // Write-only binding to the contract
	TxFilterer   // Log filterer for contract events
}

// TxCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxSession struct {
	Contract     *Tx               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxCallerSession struct {
	Contract *TxCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxTransactorSession struct {
	Contract     *TxTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxRaw struct {
	Contract *Tx // Generic contract binding to access the raw methods on
}

// TxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxCallerRaw struct {
	Contract *TxCaller // Generic read-only contract binding to access the raw methods on
}

// TxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxTransactorRaw struct {
	Contract *TxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTx creates a new instance of Tx, bound to a specific deployed contract.
func NewTx(address common.Address, backend bind.ContractBackend) (*Tx, error) {
	contract, err := bindTx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tx{TxCaller: TxCaller{contract: contract}, TxTransactor: TxTransactor{contract: contract}, TxFilterer: TxFilterer{contract: contract}}, nil
}

// NewTxCaller creates a new read-only instance of Tx, bound to a specific deployed contract.
func NewTxCaller(address common.Address, caller bind.ContractCaller) (*TxCaller, error) {
	contract, err := bindTx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxCaller{contract: contract}, nil
}

// NewTxTransactor creates a new write-only instance of Tx, bound to a specific deployed contract.
func NewTxTransactor(address common.Address, transactor bind.ContractTransactor) (*TxTransactor, error) {
	contract, err := bindTx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxTransactor{contract: contract}, nil
}

// NewTxFilterer creates a new log filterer instance of Tx, bound to a specific deployed contract.
func NewTxFilterer(address common.Address, filterer bind.ContractFilterer) (*TxFilterer, error) {
	contract, err := bindTx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxFilterer{contract: contract}, nil
}

// bindTx binds a generic wrapper to an already deployed contract.
func bindTx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tx *TxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tx.Contract.TxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tx *TxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tx.Contract.TxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tx *TxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tx.Contract.TxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tx *TxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tx *TxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tx *TxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tx.Contract.contract.Transact(opts, method, params...)
}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txRaw) pure returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address) info)
func (_Tx *TxCaller) DecodeTx(opts *bind.CallOpts, txRaw []byte) (TxTxInfo, error) {
	var out []interface{}
	err := _Tx.contract.Call(opts, &out, "decodeTx", txRaw)

	if err != nil {
		return *new(TxTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TxTxInfo)).(*TxTxInfo)

	return out0, err

}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txRaw) pure returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address) info)
func (_Tx *TxSession) DecodeTx(txRaw []byte) (TxTxInfo, error) {
	return _Tx.Contract.DecodeTx(&_Tx.CallOpts, txRaw)
}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txRaw) pure returns((uint64,uint64,uint256,uint256,uint256,address,uint256,bytes,address) info)
func (_Tx *TxCallerSession) DecodeTx(txRaw []byte) (TxTxInfo, error) {
	return _Tx.Contract.DecodeTx(&_Tx.CallOpts, txRaw)
}
