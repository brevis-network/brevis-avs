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

// BLSApkRegMetaData contains all meta data concerning the BLSApkReg contract.
var BLSApkRegMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// BLSApkRegABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSApkRegMetaData.ABI instead.
var BLSApkRegABI = BLSApkRegMetaData.ABI

// BLSApkReg is an auto generated Go binding around an Ethereum contract.
type BLSApkReg struct {
	BLSApkRegCaller     // Read-only binding to the contract
	BLSApkRegTransactor // Write-only binding to the contract
	BLSApkRegFilterer   // Log filterer for contract events
}

// BLSApkRegCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSApkRegCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSApkRegTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSApkRegFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSApkRegSession struct {
	Contract     *BLSApkReg        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSApkRegCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSApkRegCallerSession struct {
	Contract *BLSApkRegCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BLSApkRegTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSApkRegTransactorSession struct {
	Contract     *BLSApkRegTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BLSApkRegRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSApkRegRaw struct {
	Contract *BLSApkReg // Generic contract binding to access the raw methods on
}

// BLSApkRegCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSApkRegCallerRaw struct {
	Contract *BLSApkRegCaller // Generic read-only contract binding to access the raw methods on
}

// BLSApkRegTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSApkRegTransactorRaw struct {
	Contract *BLSApkRegTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSApkReg creates a new instance of BLSApkReg, bound to a specific deployed contract.
func NewBLSApkReg(address common.Address, backend bind.ContractBackend) (*BLSApkReg, error) {
	contract, err := bindBLSApkReg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSApkReg{BLSApkRegCaller: BLSApkRegCaller{contract: contract}, BLSApkRegTransactor: BLSApkRegTransactor{contract: contract}, BLSApkRegFilterer: BLSApkRegFilterer{contract: contract}}, nil
}

// NewBLSApkRegCaller creates a new read-only instance of BLSApkReg, bound to a specific deployed contract.
func NewBLSApkRegCaller(address common.Address, caller bind.ContractCaller) (*BLSApkRegCaller, error) {
	contract, err := bindBLSApkReg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegCaller{contract: contract}, nil
}

// NewBLSApkRegTransactor creates a new write-only instance of BLSApkReg, bound to a specific deployed contract.
func NewBLSApkRegTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSApkRegTransactor, error) {
	contract, err := bindBLSApkReg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegTransactor{contract: contract}, nil
}

// NewBLSApkRegFilterer creates a new log filterer instance of BLSApkReg, bound to a specific deployed contract.
func NewBLSApkRegFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSApkRegFilterer, error) {
	contract, err := bindBLSApkReg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegFilterer{contract: contract}, nil
}

// bindBLSApkReg binds a generic wrapper to an already deployed contract.
func bindBLSApkReg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BLSApkRegABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSApkReg *BLSApkRegRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSApkReg.Contract.BLSApkRegCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSApkReg *BLSApkRegRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkReg.Contract.BLSApkRegTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSApkReg *BLSApkRegRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSApkReg.Contract.BLSApkRegTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSApkReg *BLSApkRegCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSApkReg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSApkReg *BLSApkRegTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkReg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSApkReg *BLSApkRegTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSApkReg.Contract.contract.Transact(opts, method, params...)
}

// BLSApkRegNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the BLSApkReg contract.
type BLSApkRegNewPubkeyRegistrationIterator struct {
	Event *BLSApkRegNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *BLSApkRegNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegNewPubkeyRegistration)
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
		it.Event = new(BLSApkRegNewPubkeyRegistration)
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
func (it *BLSApkRegNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the BLSApkReg contract.
type BLSApkRegNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkReg *BLSApkRegFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*BLSApkRegNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BLSApkReg.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegNewPubkeyRegistrationIterator{contract: _BLSApkReg.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkReg *BLSApkRegFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *BLSApkRegNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BLSApkReg.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegNewPubkeyRegistration)
				if err := _BLSApkReg.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkReg *BLSApkRegFilterer) ParseNewPubkeyRegistration(log types.Log) (*BLSApkRegNewPubkeyRegistration, error) {
	event := new(BLSApkRegNewPubkeyRegistration)
	if err := _BLSApkReg.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the BLSApkReg contract.
type BLSApkRegOperatorAddedToQuorumsIterator struct {
	Event *BLSApkRegOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *BLSApkRegOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegOperatorAddedToQuorums)
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
		it.Event = new(BLSApkRegOperatorAddedToQuorums)
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
func (it *BLSApkRegOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the BLSApkReg contract.
type BLSApkRegOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkReg *BLSApkRegFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*BLSApkRegOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _BLSApkReg.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegOperatorAddedToQuorumsIterator{contract: _BLSApkReg.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkReg *BLSApkRegFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *BLSApkRegOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _BLSApkReg.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegOperatorAddedToQuorums)
				if err := _BLSApkReg.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkReg *BLSApkRegFilterer) ParseOperatorAddedToQuorums(log types.Log) (*BLSApkRegOperatorAddedToQuorums, error) {
	event := new(BLSApkRegOperatorAddedToQuorums)
	if err := _BLSApkReg.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the BLSApkReg contract.
type BLSApkRegOperatorRemovedFromQuorumsIterator struct {
	Event *BLSApkRegOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *BLSApkRegOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegOperatorRemovedFromQuorums)
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
		it.Event = new(BLSApkRegOperatorRemovedFromQuorums)
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
func (it *BLSApkRegOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the BLSApkReg contract.
type BLSApkRegOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkReg *BLSApkRegFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*BLSApkRegOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _BLSApkReg.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegOperatorRemovedFromQuorumsIterator{contract: _BLSApkReg.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkReg *BLSApkRegFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *BLSApkRegOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _BLSApkReg.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegOperatorRemovedFromQuorums)
				if err := _BLSApkReg.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_BLSApkReg *BLSApkRegFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*BLSApkRegOperatorRemovedFromQuorums, error) {
	event := new(BLSApkRegOperatorRemovedFromQuorums)
	if err := _BLSApkReg.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
