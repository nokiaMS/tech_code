// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wandappETH

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DemoERC20ABI is the input ABI used to generate the binding from.
const DemoERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DemoERC20Bin is the compiled bytecode used for deploying new contracts.
const DemoERC20Bin = `0x`

// DeployDemoERC20 deploys a new Ethereum contract, binding an instance of DemoERC20 to it.
func DeployDemoERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DemoERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(DemoERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DemoERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DemoERC20{DemoERC20Caller: DemoERC20Caller{contract: contract}, DemoERC20Transactor: DemoERC20Transactor{contract: contract}, DemoERC20Filterer: DemoERC20Filterer{contract: contract}}, nil
}

// DemoERC20 is an auto generated Go binding around an Ethereum contract.
type DemoERC20 struct {
	DemoERC20Caller     // Read-only binding to the contract
	DemoERC20Transactor // Write-only binding to the contract
	DemoERC20Filterer   // Log filterer for contract events
}

// DemoERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type DemoERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DemoERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DemoERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DemoERC20Session struct {
	Contract     *DemoERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DemoERC20CallerSession struct {
	Contract *DemoERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DemoERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DemoERC20TransactorSession struct {
	Contract     *DemoERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DemoERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type DemoERC20Raw struct {
	Contract *DemoERC20 // Generic contract binding to access the raw methods on
}

// DemoERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DemoERC20CallerRaw struct {
	Contract *DemoERC20Caller // Generic read-only contract binding to access the raw methods on
}

// DemoERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DemoERC20TransactorRaw struct {
	Contract *DemoERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDemoERC20 creates a new instance of DemoERC20, bound to a specific deployed contract.
func NewDemoERC20(address common.Address, backend bind.ContractBackend) (*DemoERC20, error) {
	contract, err := bindDemoERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DemoERC20{DemoERC20Caller: DemoERC20Caller{contract: contract}, DemoERC20Transactor: DemoERC20Transactor{contract: contract}, DemoERC20Filterer: DemoERC20Filterer{contract: contract}}, nil
}

// NewDemoERC20Caller creates a new read-only instance of DemoERC20, bound to a specific deployed contract.
func NewDemoERC20Caller(address common.Address, caller bind.ContractCaller) (*DemoERC20Caller, error) {
	contract, err := bindDemoERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DemoERC20Caller{contract: contract}, nil
}

// NewDemoERC20Transactor creates a new write-only instance of DemoERC20, bound to a specific deployed contract.
func NewDemoERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*DemoERC20Transactor, error) {
	contract, err := bindDemoERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DemoERC20Transactor{contract: contract}, nil
}

// NewDemoERC20Filterer creates a new log filterer instance of DemoERC20, bound to a specific deployed contract.
func NewDemoERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*DemoERC20Filterer, error) {
	contract, err := bindDemoERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DemoERC20Filterer{contract: contract}, nil
}

// bindDemoERC20 binds a generic wrapper to an already deployed contract.
func bindDemoERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DemoERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DemoERC20 *DemoERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DemoERC20.Contract.DemoERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DemoERC20 *DemoERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoERC20.Contract.DemoERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DemoERC20 *DemoERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DemoERC20.Contract.DemoERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DemoERC20 *DemoERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DemoERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DemoERC20 *DemoERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DemoERC20 *DemoERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DemoERC20.Contract.contract.Transact(opts, method, params...)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Approve(&_DemoERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Approve(&_DemoERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Transfer(&_DemoERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Transfer(&_DemoERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.TransferFrom(&_DemoERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_DemoERC20 *DemoERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.TransferFrom(&_DemoERC20.TransactOpts, _from, _to, _value)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058205009b50ef18b0459a5aef77c9d756a391d442ca4208e0632582b533e3f0a6be60029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ProofVerifyABI is the input ABI used to generate the binding from.
const ProofVerifyABI = "[]"

// ProofVerifyBin is the compiled bytecode used for deploying new contracts.
const ProofVerifyBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206141c7f5279239401e4bfab0218ad737d7d8668560d621b4ba6d664269d9f8d60029`

// DeployProofVerify deploys a new Ethereum contract, binding an instance of ProofVerify to it.
func DeployProofVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProofVerify, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofVerifyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProofVerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProofVerify{ProofVerifyCaller: ProofVerifyCaller{contract: contract}, ProofVerifyTransactor: ProofVerifyTransactor{contract: contract}, ProofVerifyFilterer: ProofVerifyFilterer{contract: contract}}, nil
}

// ProofVerify is an auto generated Go binding around an Ethereum contract.
type ProofVerify struct {
	ProofVerifyCaller     // Read-only binding to the contract
	ProofVerifyTransactor // Write-only binding to the contract
	ProofVerifyFilterer   // Log filterer for contract events
}

// ProofVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofVerifySession struct {
	Contract     *ProofVerify      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofVerifyCallerSession struct {
	Contract *ProofVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ProofVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofVerifyTransactorSession struct {
	Contract     *ProofVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ProofVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofVerifyRaw struct {
	Contract *ProofVerify // Generic contract binding to access the raw methods on
}

// ProofVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofVerifyCallerRaw struct {
	Contract *ProofVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// ProofVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofVerifyTransactorRaw struct {
	Contract *ProofVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofVerify creates a new instance of ProofVerify, bound to a specific deployed contract.
func NewProofVerify(address common.Address, backend bind.ContractBackend) (*ProofVerify, error) {
	contract, err := bindProofVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofVerify{ProofVerifyCaller: ProofVerifyCaller{contract: contract}, ProofVerifyTransactor: ProofVerifyTransactor{contract: contract}, ProofVerifyFilterer: ProofVerifyFilterer{contract: contract}}, nil
}

// NewProofVerifyCaller creates a new read-only instance of ProofVerify, bound to a specific deployed contract.
func NewProofVerifyCaller(address common.Address, caller bind.ContractCaller) (*ProofVerifyCaller, error) {
	contract, err := bindProofVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofVerifyCaller{contract: contract}, nil
}

// NewProofVerifyTransactor creates a new write-only instance of ProofVerify, bound to a specific deployed contract.
func NewProofVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofVerifyTransactor, error) {
	contract, err := bindProofVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofVerifyTransactor{contract: contract}, nil
}

// NewProofVerifyFilterer creates a new log filterer instance of ProofVerify, bound to a specific deployed contract.
func NewProofVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofVerifyFilterer, error) {
	contract, err := bindProofVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofVerifyFilterer{contract: contract}, nil
}

// bindProofVerify binds a generic wrapper to an already deployed contract.
func bindProofVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofVerifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofVerify *ProofVerifyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProofVerify.Contract.ProofVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofVerify *ProofVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofVerify.Contract.ProofVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofVerify *ProofVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofVerify.Contract.ProofVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofVerify *ProofVerifyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProofVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofVerify *ProofVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofVerify *ProofVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofVerify.Contract.contract.Transact(opts, method, params...)
}

// WandappBaseABI is the input ABI used to generate the binding from.
const WandappBaseABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getChlID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenOrigAddr\",\"type\":\"address\"},{\"name\":\"operAddr\",\"type\":\"address\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"HashEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"deposition\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTransfer\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"chlID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errcode\",\"type\":\"uint8\"}],\"name\":\"BadProofEvent\",\"type\":\"event\"}]"

// WandappBaseBin is the compiled bytecode used for deploying new contracts.
const WandappBaseBin = `0x608060405234801561001057600080fd5b5060405160608061013183398101604090815281516020830151919092015160008054600160a060020a03948516600160a060020a0319918216179091556001805494909316931692909217905560025560c28061006f6000396000f300608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631469a07381146043575b600080fd5b348015604e57600080fd5b5060556067565b60408051918252519081900360200190f35b600060746003546001607e565b6003819055905090565b600082820183811015608f57600080fd5b93925050505600a165627a7a723058202bb0c2802275f55b2f8114ff0f48d38446dbdcb49f94640ca667539ef6da67350029`

// DeployWandappBase deploys a new Ethereum contract, binding an instance of WandappBase to it.
func DeployWandappBase(auth *bind.TransactOpts, backend bind.ContractBackend, tokenOrigAddr common.Address, operAddr common.Address, duration *big.Int) (common.Address, *types.Transaction, *WandappBase, error) {
	parsed, err := abi.JSON(strings.NewReader(WandappBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WandappBaseBin), backend, tokenOrigAddr, operAddr, duration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WandappBase{WandappBaseCaller: WandappBaseCaller{contract: contract}, WandappBaseTransactor: WandappBaseTransactor{contract: contract}, WandappBaseFilterer: WandappBaseFilterer{contract: contract}}, nil
}

// WandappBase is an auto generated Go binding around an Ethereum contract.
type WandappBase struct {
	WandappBaseCaller     // Read-only binding to the contract
	WandappBaseTransactor // Write-only binding to the contract
	WandappBaseFilterer   // Log filterer for contract events
}

// WandappBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type WandappBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WandappBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WandappBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WandappBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WandappBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WandappBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WandappBaseSession struct {
	Contract     *WandappBase      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WandappBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WandappBaseCallerSession struct {
	Contract *WandappBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WandappBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WandappBaseTransactorSession struct {
	Contract     *WandappBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WandappBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type WandappBaseRaw struct {
	Contract *WandappBase // Generic contract binding to access the raw methods on
}

// WandappBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WandappBaseCallerRaw struct {
	Contract *WandappBaseCaller // Generic read-only contract binding to access the raw methods on
}

// WandappBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WandappBaseTransactorRaw struct {
	Contract *WandappBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWandappBase creates a new instance of WandappBase, bound to a specific deployed contract.
func NewWandappBase(address common.Address, backend bind.ContractBackend) (*WandappBase, error) {
	contract, err := bindWandappBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WandappBase{WandappBaseCaller: WandappBaseCaller{contract: contract}, WandappBaseTransactor: WandappBaseTransactor{contract: contract}, WandappBaseFilterer: WandappBaseFilterer{contract: contract}}, nil
}

// NewWandappBaseCaller creates a new read-only instance of WandappBase, bound to a specific deployed contract.
func NewWandappBaseCaller(address common.Address, caller bind.ContractCaller) (*WandappBaseCaller, error) {
	contract, err := bindWandappBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WandappBaseCaller{contract: contract}, nil
}

// NewWandappBaseTransactor creates a new write-only instance of WandappBase, bound to a specific deployed contract.
func NewWandappBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*WandappBaseTransactor, error) {
	contract, err := bindWandappBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WandappBaseTransactor{contract: contract}, nil
}

// NewWandappBaseFilterer creates a new log filterer instance of WandappBase, bound to a specific deployed contract.
func NewWandappBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*WandappBaseFilterer, error) {
	contract, err := bindWandappBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WandappBaseFilterer{contract: contract}, nil
}

// bindWandappBase binds a generic wrapper to an already deployed contract.
func bindWandappBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WandappBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WandappBase *WandappBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WandappBase.Contract.WandappBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WandappBase *WandappBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappBase.Contract.WandappBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WandappBase *WandappBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WandappBase.Contract.WandappBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WandappBase *WandappBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WandappBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WandappBase *WandappBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WandappBase *WandappBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WandappBase.Contract.contract.Transact(opts, method, params...)
}

// GetChlID is a paid mutator transaction binding the contract method 0x1469a073.
//
// Solidity: function getChlID() returns(uint256)
func (_WandappBase *WandappBaseTransactor) GetChlID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappBase.contract.Transact(opts, "getChlID")
}

// GetChlID is a paid mutator transaction binding the contract method 0x1469a073.
//
// Solidity: function getChlID() returns(uint256)
func (_WandappBase *WandappBaseSession) GetChlID() (*types.Transaction, error) {
	return _WandappBase.Contract.GetChlID(&_WandappBase.TransactOpts)
}

// GetChlID is a paid mutator transaction binding the contract method 0x1469a073.
//
// Solidity: function getChlID() returns(uint256)
func (_WandappBase *WandappBaseTransactorSession) GetChlID() (*types.Transaction, error) {
	return _WandappBase.Contract.GetChlID(&_WandappBase.TransactOpts)
}

// WandappBaseBadProofEventIterator is returned from FilterBadProofEvent and is used to iterate over the raw logs and unpacked data for BadProofEvent events raised by the WandappBase contract.
type WandappBaseBadProofEventIterator struct {
	Event *WandappBaseBadProofEvent // Event containing the contract specifics and raw log

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
func (it *WandappBaseBadProofEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappBaseBadProofEvent)
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
		it.Event = new(WandappBaseBadProofEvent)
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
func (it *WandappBaseBadProofEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappBaseBadProofEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappBaseBadProofEvent represents a BadProofEvent event raised by the WandappBase contract.
type WandappBaseBadProofEvent struct {
	From        common.Address
	To          common.Address
	Operator    common.Address
	TransferCnt *big.Int
	WithdrawCnt *big.Int
	Nonce       uint64
	TimeStamp   int64
	ChlID       *big.Int
	Errcode     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBadProofEvent is a free log retrieval operation binding the contract event 0x9b797295befa0a1ed24f6cb1cdbb99f4fcba5989b39a7d79108a6282f4039a9b.
//
// Solidity: e BadProofEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, chlID uint256, errcode uint8)
func (_WandappBase *WandappBaseFilterer) FilterBadProofEvent(opts *bind.FilterOpts) (*WandappBaseBadProofEventIterator, error) {

	logs, sub, err := _WandappBase.contract.FilterLogs(opts, "BadProofEvent")
	if err != nil {
		return nil, err
	}
	return &WandappBaseBadProofEventIterator{contract: _WandappBase.contract, event: "BadProofEvent", logs: logs, sub: sub}, nil
}

// WatchBadProofEvent is a free log subscription operation binding the contract event 0x9b797295befa0a1ed24f6cb1cdbb99f4fcba5989b39a7d79108a6282f4039a9b.
//
// Solidity: e BadProofEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, chlID uint256, errcode uint8)
func (_WandappBase *WandappBaseFilterer) WatchBadProofEvent(opts *bind.WatchOpts, sink chan<- *WandappBaseBadProofEvent) (event.Subscription, error) {

	logs, sub, err := _WandappBase.contract.WatchLogs(opts, "BadProofEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappBaseBadProofEvent)
				if err := _WandappBase.contract.UnpackLog(event, "BadProofEvent", log); err != nil {
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

// WandappBaseHashEventIterator is returned from FilterHashEvent and is used to iterate over the raw logs and unpacked data for HashEvent events raised by the WandappBase contract.
type WandappBaseHashEventIterator struct {
	Event *WandappBaseHashEvent // Event containing the contract specifics and raw log

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
func (it *WandappBaseHashEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappBaseHashEvent)
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
		it.Event = new(WandappBaseHashEvent)
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
func (it *WandappBaseHashEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappBaseHashEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappBaseHashEvent represents a HashEvent event raised by the WandappBase contract.
type WandappBaseHashEvent struct {
	Index [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHashEvent is a free log retrieval operation binding the contract event 0x45fffb12f7c572056d713ff90194e4559d57f06a2f6b91e50da7cf0f3b5824bc.
//
// Solidity: e HashEvent(index bytes32)
func (_WandappBase *WandappBaseFilterer) FilterHashEvent(opts *bind.FilterOpts) (*WandappBaseHashEventIterator, error) {

	logs, sub, err := _WandappBase.contract.FilterLogs(opts, "HashEvent")
	if err != nil {
		return nil, err
	}
	return &WandappBaseHashEventIterator{contract: _WandappBase.contract, event: "HashEvent", logs: logs, sub: sub}, nil
}

// WatchHashEvent is a free log subscription operation binding the contract event 0x45fffb12f7c572056d713ff90194e4559d57f06a2f6b91e50da7cf0f3b5824bc.
//
// Solidity: e HashEvent(index bytes32)
func (_WandappBase *WandappBaseFilterer) WatchHashEvent(opts *bind.WatchOpts, sink chan<- *WandappBaseHashEvent) (event.Subscription, error) {

	logs, sub, err := _WandappBase.contract.WatchLogs(opts, "HashEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappBaseHashEvent)
				if err := _WandappBase.contract.UnpackLog(event, "HashEvent", log); err != nil {
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

// WandappBaseWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the WandappBase contract.
type WandappBaseWithdrawEventIterator struct {
	Event *WandappBaseWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *WandappBaseWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappBaseWithdrawEvent)
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
		it.Event = new(WandappBaseWithdrawEvent)
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
func (it *WandappBaseWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappBaseWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappBaseWithdrawEvent represents a WithdrawEvent event raised by the WandappBase contract.
type WandappBaseWithdrawEvent struct {
	From          common.Address
	To            common.Address
	Operator      common.Address
	TransferCnt   *big.Int
	WithdrawCnt   *big.Int
	Nonce         uint64
	TimeStamp     int64
	Deposition    *big.Int
	TotalTransfer *big.Int
	TotalWithdraw *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0xb4aff9b5153482df5e297bc29c40cf9031b4c4750fb492217c2235c9ead50dc3.
//
// Solidity: e WithdrawEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, deposition uint256, totalTransfer uint256, totalWithdraw uint256)
func (_WandappBase *WandappBaseFilterer) FilterWithdrawEvent(opts *bind.FilterOpts) (*WandappBaseWithdrawEventIterator, error) {

	logs, sub, err := _WandappBase.contract.FilterLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return &WandappBaseWithdrawEventIterator{contract: _WandappBase.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0xb4aff9b5153482df5e297bc29c40cf9031b4c4750fb492217c2235c9ead50dc3.
//
// Solidity: e WithdrawEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, deposition uint256, totalTransfer uint256, totalWithdraw uint256)
func (_WandappBase *WandappBaseFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *WandappBaseWithdrawEvent) (event.Subscription, error) {

	logs, sub, err := _WandappBase.contract.WatchLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappBaseWithdrawEvent)
				if err := _WandappBase.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// WandappETHABI is the input ABI used to generate the binding from.
const WandappETHABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getChlID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"operatorWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenOrigAddr\",\"type\":\"address\"},{\"name\":\"operAddr\",\"type\":\"address\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"chlId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"DepositeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"chlID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"chlstate\",\"type\":\"uint8\"}],\"name\":\"ExitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"target\",\"type\":\"address\"}],\"name\":\"WrongAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"operNonce\",\"type\":\"uint256\"}],\"name\":\"SettleEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"HashEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"deposition\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTransfer\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"chlID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errcode\",\"type\":\"uint8\"}],\"name\":\"BadProofEvent\",\"type\":\"event\"}]"

// WandappETHBin is the compiled bytecode used for deploying new contracts.
const WandappETHBin = `0x608060405234801561001057600080fd5b506040516060806125ab83398101604090815281516020830151919092015160008054600160a060020a03948516600160a060020a0319918216179091556001805494909316931692909217905560025561253b806100706000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630968f264811461008757806311da60b4146100e25780631469a073146100f75780633104562b1461011e5780633805550f146101365780639d955d791461018f578063f97a25c0146101f5575b600080fd5b34801561009357600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100e094369492936024939284019190819084018382808284375094975061024e9650505050505050565b005b3480156100ee57600080fd5b506100e061064b565b34801561010357600080fd5b5061010c610878565b60408051918252519081900360200190f35b34801561012a57600080fd5b506100e0600435610891565b34801561014257600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100e0943694929360249392840191908190840183828082843750949750610a709650505050505050565b34801561019b57600080fd5b506101b0600160a060020a0360043516611086565b60408051978852602088019690965286860194909452606086019290925267ffffffffffffffff90811660808601521660a084015260c0830152519081900360e00190f35b34801561020157600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100e09436949293602493928401919081908401838280828437509497506110df9650505050505050565b61025661249b565b60006102618361153d565b915060018251600160a060020a031660009081526014602052604090206006015460ff16600281111561029057fe5b14156102e6576040805160e560020a62461bcd02815260206004820152601a60248201527f57726f6e67206f7074696f6e20696e207468652073746174652e000000000000604482015290519081900360640190fd5b6102ef826119d4565b8151600160a060020a031660009081526014602052604090206004015460a083015167ffffffffffffffff918216911611610362576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206124f0833981519152604482015290519081900360640190fd5b60808201518251600160a060020a031660009081526014602052604090206002015461038e9190611d61565b600080548451604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015260248101869052905194955091169263a9059cbb92604480840193602093929083900390910190829087803b15801561040357600080fd5b505af1158015610417573d6000803e3d6000fd5b505050506040513d602081101561042d57600080fd5b5051151561043a57600080fd5b8160800151601460008460000151600160a060020a0316600160a060020a03168152602001908152602001600020600201819055508160a00151601460008460000151600160a060020a0316600160a060020a0316815260200190815260200160002060040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fb4aff9b5153482df5e297bc29c40cf9031b4c4750fb492217c2235c9ead50dc382600001518360200151600160009054906101000a9004600160a060020a0316856060015186608001518760a001518860c00151601460008b60000151600160a060020a0316600160a060020a0316815260200190815260200160002060010154601460008c60000151600160a060020a0316600160a060020a0316815260200190815260200160002060020154601460008d60000151600160a060020a0316600160a060020a0316815260200190815260200160002060030154604051808b600160a060020a0316600160a060020a031681526020018a600160a060020a0316600160a060020a0316815260200189600160a060020a0316600160a060020a031681526020018881526020018781526020018667ffffffffffffffff1667ffffffffffffffff1681526020018560070b60070b81526020018481526020018381526020018281526020019a505050505050505050505060405180910390a1505050565b336000908152601460205260409020805415156106b2576040805160e560020a62461bcd02815260206004820152601060248201527f756e6b6e6f776564206368616e6e656c00000000000000000000000000000000604482015290519081900360640190fd5b6005810154421015610734576040805160e560020a62461bcd02815260206004820152602b60248201527f7468652074696d65206f6620736574746c696e6e67206368616e6e656c20697360448201527f6e27742020636f6d696e67000000000000000000000000000000000000000000606482015290519081900360840190fd5b6001600682015460ff16600281111561074957fe5b148061077c57506002600682015460ff16600281111561076557fe5b14801561077c5750600154600160a060020a031633145b15156107d2576040805160e560020a62461bcd02815260206004820152601f60248201527f5573657220686173206e6f206175746820696e20746869732073746174652e00604482015290519081900360640190fd5b6040805161014081018252600780840154600160a060020a03908116835260088501548116602084015260098501541692820192909252600a8301546060820152600b8301546080820152600c83015467ffffffffffffffff811660a0830152680100000000000000009004820b820b90910b60c0820152600d82015460e0820152600e820154610100820152600f820154610120820152610875903390611d78565b50565b60006108876003546001612180565b6003819055905090565b8060008110156108eb576040805160e560020a62461bcd02815260206004820152601360248201527f4e756d626572206973206e656761746976652e00000000000000000000000000604482015290519081900360640190fd5b60008054604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018690529051600160a060020a03909216926323b872dd926064808401936020939083900390910190829087803b15801561095f57600080fd5b505af1158015610973573d6000803e3d6000fd5b505050506040513d602081101561098957600080fd5b5051151561099657600080fd5b3360009081526014602052604090205415156109d2576109b4610878565b336000908152601460205260409020908155600601805460ff191690555b336000908152601460205260409020600101546109ef9083612180565b3360008181526014602090815260408083206001808201879055905490549354825191825292810194909452600160a060020a0392831684820152911660608301526080820185905260a082019290925290517f6fd22d40cda41d7dbbf23f88347434a3907b286562d061724be4200d1da8dd089181900360c00190a15050565b610a7861249b565b6000610a838361153d565b9150610a8e826119d4565b508051600160a060020a0316600090815260146020526040812090600682015460ff166002811115610abc57fe5b1415610cfe5760025442016005820155600154600160a060020a0316331415610b73578151600160a060020a031660009081526014602052604090206004015460a083015167ffffffffffffffff68010000000000000000909204821691161015610b5f576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206124f0833981519152604482015290519081900360640190fd5b60068101805460ff19166002179055610bf7565b8151600160a060020a031660009081526014602052604090206004015460a083015167ffffffffffffffff91821691161015610be7576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206124f0833981519152604482015290519081900360640190fd5b60068101805460ff191660011790555b8151600160a060020a0390811660009081526014602090815260409182902085516007808301805492871673ffffffffffffffffffffffffffffffffffffffff199384161790559287015160088301805491871691831691909117905592860151600982018054919095169316929092179092556060840151600a8201556080840151600b82015560a0840151600c8201805460c087015190940b67ffffffffffffffff90811668010000000000000000026fffffffffffffffff0000000000000000199190931667ffffffffffffffff1990951694909417939093161790915560e0830151600d820155610100830151600e820155610120830151600f90910155610eca565b6002600682015460ff166002811115610d1357fe5b1415610d7a578151600160a060020a03163314610d7a576040805160e560020a62461bcd02815260206004820152601b60248201527f57726f6e672063616c6c657220696e20746869732073746174652e0000000000604482015290519081900360640190fd5b6001600682015460ff166002811115610d8f57fe5b1415610df757600154600160a060020a03163314610df7576040805160e560020a62461bcd02815260206004820152601b60248201527f57726f6e672063616c6c657220696e20746869732073746174652e0000000000604482015290519081900360640190fd5b600c81015460a083015167ffffffffffffffff91821691161115610e26578151610e219083611d78565b610eca565b81516040805161014081018252600780850154600160a060020a03908116835260088601548116602084015260098601541692820192909252600a8401546060820152600b8401546080820152600c84015467ffffffffffffffff811660a0830152680100000000000000009004820b820b90910b60c0820152600d83015460e0820152600e830154610100820152600f830154610120820152610eca9190611d78565b7f90b0279af0b82a8471a26044997dc3e95eb00db377daa4806cacba5730f07c1582600001518360200151600160009054906101000a9004600160a060020a0316601460008760000151600160a060020a0316600160a060020a0316815260200190815260200160002060030154601460008860000151600160a060020a0316600160a060020a03168152602001908152602001600020600201548760a001518860c00151601460008b60000151600160a060020a0316600160a060020a0316815260200190815260200160002060000154601460008c60000151600160a060020a0316600160a060020a0316815260200190815260200160002060060160009054906101000a900460ff16604051808a600160a060020a0316600160a060020a0316815260200189600160a060020a0316600160a060020a0316815260200188600160a060020a0316600160a060020a031681526020018781526020018681526020018567ffffffffffffffff1667ffffffffffffffff1681526020018460070b60070b815260200183815260200182600281111561106657fe5b60ff168152602001995050505050505050505060405180910390a1505050565b600160a060020a0316600090815260146020526040902080546001820154600283015460038401546004850154600590950154939592949193909267ffffffffffffffff80841693680100000000000000009004169190565b6110e761249b565b600154600090600160a060020a031633811461114d576040805160e560020a62461bcd02815260206004820152601660248201527f53656e646572206e6f7420617574686f72697a65642e00000000000000000000604482015290519081900360640190fd5b6111568461153d565b925060028351600160a060020a031660009081526014602052604090206006015460ff16600281111561118557fe5b14156111db576040805160e560020a62461bcd02815260206004820152601a60248201527f57726f6e67206f7074696f6e20696e207468652073746174652e000000000000604482015290519081900360640190fd5b6111e4836119d4565b8251600160a060020a031660009081526014602052604090206004015460a084015167ffffffffffffffff680100000000000000009092048216911611611263576040805160e560020a62461bcd02815260206004820152600e60248201526000805160206124f0833981519152604482015290519081900360640190fd5b60608301518351600160a060020a031660009081526014602052604090206003015461128f9190611d61565b60008054600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03928316600482015260248101869052905194965091169263a9059cbb92604480840193602093929083900390910190829087803b15801561130557600080fd5b505af1158015611319573d6000803e3d6000fd5b505050506040513d602081101561132f57600080fd5b5051151561133c57600080fd5b8260600151601460008560000151600160a060020a0316600160a060020a03168152602001908152602001600020600301819055508260a00151601460008560000151600160a060020a0316600160a060020a0316815260200190815260200160002060040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fb4aff9b5153482df5e297bc29c40cf9031b4c4750fb492217c2235c9ead50dc3836000015184602001518560400151866060015187608001518860a001518960c00151601460008c60000151600160a060020a0316600160a060020a0316815260200190815260200160002060010154601460008d60000151600160a060020a0316600160a060020a0316815260200190815260200160002060020154601460008e60000151600160a060020a0316600160a060020a0316815260200190815260200160002060030154604051808b600160a060020a0316600160a060020a031681526020018a600160a060020a0316600160a060020a0316815260200189600160a060020a0316600160a060020a031681526020018881526020018781526020018667ffffffffffffffff1667ffffffffffffffff1681526020018560070b60070b81526020018481526020018381526020018281526020019a505050505050505050505060405180910390a150505050565b61154561249b565b60608061155061249b565b84516101ae1461155f57600080fd5b6040805160ec80825261012082019092529060208201611d80803883390190505092506020850151602084015260408501516040840152606085015160608401526080850151608084015260a085015160a084015260c085015160c084015260e085015160e08401526101008501516101008401526115dd85612199565b600160a060020a0316611654846040518082805190602001908083835b602083106116195780518252601f1990920191602091820191016115fa565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902061164f886121a9565b6121fc565b600160a060020a0316146116b2576040805160e560020a62461bcd02815260206004820152601d60248201527f746865207369676e61757265206f6620757365722069732077726f6e67000000604482015290519081900360640190fd5b6040805161010c80825261014082019092529060208201612180803883390190505091506020850151602083015260408501516040830152606085015160608301526080850151608083015260a085015160a083015260c085015160c083015260e085015160e083015261010085015161010083015261012085015161012083015261173d856123d3565b600160a060020a03166117af836040518082805190602001908083835b602083106117795780518252601f19909201916020918201910161175a565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902061164f886123e3565b600160a060020a031614611833576040805160e560020a62461bcd02815260206004820152602160248201527f746865207369676e61757265206f66206f70657261746f722069732077726f6e60448201527f6700000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61183c85612436565b6040805160208082019390935281518082038401815290820191829052805190928291908401908083835b602083106118865780518252601f199092019160209182019101611867565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019166118c08661243e565b14611915576040805160e560020a62461bcd02815260206004820152601960248201527f6261642073656372657420666f7220746869732070726f6f6600000000000000604482015290519081900360640190fd5b61191e85612199565b600160a060020a0316815261193285612445565b600160a060020a03166020820152611949856123d3565b600160a060020a0316604082015261196085612455565b606082015261196e8561245c565b608082015261197c85612463565b67ffffffffffffffff1660a082015261199485612474565b600790810b900b60c08201526119a985612485565b60e08201526119b78561248c565b6101008201526119c685612493565b610120820152949350505050565b60008160a0015167ffffffffffffffff16111515611a3c576040805160e560020a62461bcd02815260206004820152600d60248201527f496c6c6567616c206e6f6e636500000000000000000000000000000000000000604482015290519081900360640190fd5b61010081015160001115611a9a576040805160e560020a62461bcd02815260206004820152600d60248201527f496c6c6567616c206e6f6e636500000000000000000000000000000000000000604482015290519081900360640190fd5b6001546040820151600160a060020a03908116911614611b04576040805160e560020a62461bcd02815260206004820152601960248201527f756e6b6e6f776564206f70657261746f72206164647265737300000000000000604482015290519081900360640190fd5b8051600160a060020a031660009081526014602052604090205415801590611b48575060e08101518151600160a060020a0316600090815260146020526040902054145b1515611b9e576040805160e560020a62461bcd02815260206004820152601360248201527f6368616e6e656c206f7574206f66206461746500000000000000000000000000604482015290519081900360640190fd5b8051600160a060020a031660009081526014602052604090206001015460808201516060830151611bcf9190612180565b1115611c25576040805160e560020a62461bcd02815260206004820152601a60248201527f68617665206e6f7420656e6f756768206465706f736974696f6e000000000000604482015290519081900360640190fd5b8051600160a060020a031660009081526014602052604090206002015460808201511015611cc3576040805160e560020a62461bcd02815260206004820152602b60248201527f776974686472617720636f756e7420697320736d616c6c6572207468616e207460448201527f6865206c617374206f6e65000000000000000000000000000000000000000000606482015290519081900360840190fd5b8051600160a060020a031660009081526014602052604090206003015460608201511015610875576040805160e560020a62461bcd02815260206004820152602b60248201527f7472616e7366657220636f756e7420697320736d616c6c6572207468616e207460448201527f6865206c617374206f6e65000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008083831115611d7157600080fd5b5050900390565b60008054600160a060020a0384811683526014602052604083206001810154600282015460608701519390941694919391928392611dc09291611dbb9190612180565b611d61565b9150611dd485606001518460030154611d61565b90506000821115611f4a5783600160a060020a031663a9059cbb87846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015611e5b57600080fd5b505af1158015611e6f573d6000803e3d6000fd5b505050506040513d6020811015611e8557600080fd5b50511515611edd576040805160e560020a62461bcd02815260206004820152601f60248201527f7472616e66657220746f6b656e20746f2063726561746f72206661696c656400604482015290519081900360640190fd5b600483015460408051600160a060020a03891681526020810185905267ffffffffffffffff80841682840152680100000000000000009093049092166060830152517ff299e08d054ee3c65880d58af77cc2d5e3eb14272c78cc3edb4aa9ce19c075919181900360800190a15b60008111156120b557600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810184905290519186169163a9059cbb916044808201926020929091908290030181600087803b158015611fc257600080fd5b505af1158015611fd6573d6000803e3d6000fd5b505050506040513d6020811015611fec57600080fd5b50511515612044576040805160e560020a62461bcd02815260206004820181905260248201527f7472616e66657220746f6b656e20746f206f70657261746f72206661696c6564604482015290519081900360640190fd5b600154600484015460408051600160a060020a0390931683526020830184905267ffffffffffffffff80831684830152680100000000000000009092049091166060830152517ff299e08d054ee3c65880d58af77cc2d5e3eb14272c78cc3edb4aa9ce19c075919181900360800190a15b505050600160a060020a0390921660009081526014602052604081208181556001810182905560028101829055600381018290556004810180546fffffffffffffffffffffffffffffffff199081169091556005820183905560068201805460ff1916905560078201805473ffffffffffffffffffffffffffffffffffffffff19908116909155600883018054821690556009830180549091169055600a8201839055600b8201839055600c820180549091169055600d8101829055600e8101829055600f01555050565b60008282018381101561219257600080fd5b9392505050565b60140151600160a060020a031690565b60408051604180825260808201909252606091602082016108208038833901905050905061012c820151602082015261014c820151604082015260ff61014d830151166041820151176041820152919050565b60008060008060608551604114151561221457600080fd5b602086015160408701516041880151919550935060ff169150601b82101561223d57601b820191505b8160ff16601b148061225257508160ff16601c145b151561225d57600080fd5b6040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600181886040516020018083805190602001908083835b602083106122cb5780518252601f1990920191602091820191016122ac565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193945092839250908401908083835b6020831061232b5780518252601f19909201916020918201910161230c565b51815160209384036101000a60001901801990921691161790526040805192909401829003822060008084528383018087529190915260ff8a1683860152606083018c9052608083018b9052935160a08084019750919550601f1981019492819003909101925090865af11580156123a7573d6000803e3d6000fd5b5050604051601f190151955050600160a060020a03851615156123c957600080fd5b5050505092915050565b603c0151600160a060020a031690565b60408051604180825260808201909252606091602082016108208038833901905050905061016d820151602082015261018d820151604082015260ff61018e830151166041820151176041820152919050565b6101ae015190565b60ac015190565b60280151600160a060020a031690565b605c015190565b607c015190565b6084015167ffffffffffffffff1690565b608c015167ffffffffffffffff1690565b60cc015190565b60ec015190565b61010c015190565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810191909152905600496c6c6567616c206e6f6e63652e000000000000000000000000000000000000a165627a7a723058201b434b1d382da11386f4cb78758a07408d6e6cd297b741ad1690b38de17c55e20029`

// DeployWandappETH deploys a new Ethereum contract, binding an instance of WandappETH to it.
func DeployWandappETH(auth *bind.TransactOpts, backend bind.ContractBackend, tokenOrigAddr common.Address, operAddr common.Address, duration *big.Int) (common.Address, *types.Transaction, *WandappETH, error) {
	parsed, err := abi.JSON(strings.NewReader(WandappETHABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WandappETHBin), backend, tokenOrigAddr, operAddr, duration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WandappETH{WandappETHCaller: WandappETHCaller{contract: contract}, WandappETHTransactor: WandappETHTransactor{contract: contract}, WandappETHFilterer: WandappETHFilterer{contract: contract}}, nil
}

// WandappETH is an auto generated Go binding around an Ethereum contract.
type WandappETH struct {
	WandappETHCaller     // Read-only binding to the contract
	WandappETHTransactor // Write-only binding to the contract
	WandappETHFilterer   // Log filterer for contract events
}

// WandappETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type WandappETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WandappETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WandappETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WandappETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WandappETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WandappETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WandappETHSession struct {
	Contract     *WandappETH       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WandappETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WandappETHCallerSession struct {
	Contract *WandappETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WandappETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WandappETHTransactorSession struct {
	Contract     *WandappETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WandappETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type WandappETHRaw struct {
	Contract *WandappETH // Generic contract binding to access the raw methods on
}

// WandappETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WandappETHCallerRaw struct {
	Contract *WandappETHCaller // Generic read-only contract binding to access the raw methods on
}

// WandappETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WandappETHTransactorRaw struct {
	Contract *WandappETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWandappETH creates a new instance of WandappETH, bound to a specific deployed contract.
func NewWandappETH(address common.Address, backend bind.ContractBackend) (*WandappETH, error) {
	contract, err := bindWandappETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WandappETH{WandappETHCaller: WandappETHCaller{contract: contract}, WandappETHTransactor: WandappETHTransactor{contract: contract}, WandappETHFilterer: WandappETHFilterer{contract: contract}}, nil
}

// NewWandappETHCaller creates a new read-only instance of WandappETH, bound to a specific deployed contract.
func NewWandappETHCaller(address common.Address, caller bind.ContractCaller) (*WandappETHCaller, error) {
	contract, err := bindWandappETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WandappETHCaller{contract: contract}, nil
}

// NewWandappETHTransactor creates a new write-only instance of WandappETH, bound to a specific deployed contract.
func NewWandappETHTransactor(address common.Address, transactor bind.ContractTransactor) (*WandappETHTransactor, error) {
	contract, err := bindWandappETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WandappETHTransactor{contract: contract}, nil
}

// NewWandappETHFilterer creates a new log filterer instance of WandappETH, bound to a specific deployed contract.
func NewWandappETHFilterer(address common.Address, filterer bind.ContractFilterer) (*WandappETHFilterer, error) {
	contract, err := bindWandappETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WandappETHFilterer{contract: contract}, nil
}

// bindWandappETH binds a generic wrapper to an already deployed contract.
func bindWandappETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WandappETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WandappETH *WandappETHRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WandappETH.Contract.WandappETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WandappETH *WandappETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappETH.Contract.WandappETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WandappETH *WandappETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WandappETH.Contract.WandappETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WandappETH *WandappETHCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WandappETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WandappETH *WandappETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WandappETH *WandappETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WandappETH.Contract.contract.Transact(opts, method, params...)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0x9d955d79.
//
// Solidity: function getChannelInfo(user address) constant returns(uint256, uint256, uint256, uint256, uint64, uint64, uint256)
func (_WandappETH *WandappETHCaller) GetChannelInfo(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, uint64, uint64, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(uint64)
		ret5 = new(uint64)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _WandappETH.contract.Call(opts, out, "getChannelInfo", user)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetChannelInfo is a free data retrieval call binding the contract method 0x9d955d79.
//
// Solidity: function getChannelInfo(user address) constant returns(uint256, uint256, uint256, uint256, uint64, uint64, uint256)
func (_WandappETH *WandappETHSession) GetChannelInfo(user common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, uint64, uint64, *big.Int, error) {
	return _WandappETH.Contract.GetChannelInfo(&_WandappETH.CallOpts, user)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0x9d955d79.
//
// Solidity: function getChannelInfo(user address) constant returns(uint256, uint256, uint256, uint256, uint64, uint64, uint256)
func (_WandappETH *WandappETHCallerSession) GetChannelInfo(user common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, uint64, uint64, *big.Int, error) {
	return _WandappETH.Contract.GetChannelInfo(&_WandappETH.CallOpts, user)
}

// Deposite is a paid mutator transaction binding the contract method 0x3104562b.
//
// Solidity: function deposite(amount uint256) returns()
func (_WandappETH *WandappETHTransactor) Deposite(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WandappETH.contract.Transact(opts, "deposite", amount)
}

// Deposite is a paid mutator transaction binding the contract method 0x3104562b.
//
// Solidity: function deposite(amount uint256) returns()
func (_WandappETH *WandappETHSession) Deposite(amount *big.Int) (*types.Transaction, error) {
	return _WandappETH.Contract.Deposite(&_WandappETH.TransactOpts, amount)
}

// Deposite is a paid mutator transaction binding the contract method 0x3104562b.
//
// Solidity: function deposite(amount uint256) returns()
func (_WandappETH *WandappETHTransactorSession) Deposite(amount *big.Int) (*types.Transaction, error) {
	return _WandappETH.Contract.Deposite(&_WandappETH.TransactOpts, amount)
}

// Exit is a paid mutator transaction binding the contract method 0x3805550f.
//
// Solidity: function exit(proof bytes) returns()
func (_WandappETH *WandappETHTransactor) Exit(opts *bind.TransactOpts, proof []byte) (*types.Transaction, error) {
	return _WandappETH.contract.Transact(opts, "exit", proof)
}

// Exit is a paid mutator transaction binding the contract method 0x3805550f.
//
// Solidity: function exit(proof bytes) returns()
func (_WandappETH *WandappETHSession) Exit(proof []byte) (*types.Transaction, error) {
	return _WandappETH.Contract.Exit(&_WandappETH.TransactOpts, proof)
}

// Exit is a paid mutator transaction binding the contract method 0x3805550f.
//
// Solidity: function exit(proof bytes) returns()
func (_WandappETH *WandappETHTransactorSession) Exit(proof []byte) (*types.Transaction, error) {
	return _WandappETH.Contract.Exit(&_WandappETH.TransactOpts, proof)
}

// GetChlID is a paid mutator transaction binding the contract method 0x1469a073.
//
// Solidity: function getChlID() returns(uint256)
func (_WandappETH *WandappETHTransactor) GetChlID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappETH.contract.Transact(opts, "getChlID")
}

// GetChlID is a paid mutator transaction binding the contract method 0x1469a073.
//
// Solidity: function getChlID() returns(uint256)
func (_WandappETH *WandappETHSession) GetChlID() (*types.Transaction, error) {
	return _WandappETH.Contract.GetChlID(&_WandappETH.TransactOpts)
}

// GetChlID is a paid mutator transaction binding the contract method 0x1469a073.
//
// Solidity: function getChlID() returns(uint256)
func (_WandappETH *WandappETHTransactorSession) GetChlID() (*types.Transaction, error) {
	return _WandappETH.Contract.GetChlID(&_WandappETH.TransactOpts)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0xf97a25c0.
//
// Solidity: function operatorWithdraw(proof bytes) returns()
func (_WandappETH *WandappETHTransactor) OperatorWithdraw(opts *bind.TransactOpts, proof []byte) (*types.Transaction, error) {
	return _WandappETH.contract.Transact(opts, "operatorWithdraw", proof)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0xf97a25c0.
//
// Solidity: function operatorWithdraw(proof bytes) returns()
func (_WandappETH *WandappETHSession) OperatorWithdraw(proof []byte) (*types.Transaction, error) {
	return _WandappETH.Contract.OperatorWithdraw(&_WandappETH.TransactOpts, proof)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0xf97a25c0.
//
// Solidity: function operatorWithdraw(proof bytes) returns()
func (_WandappETH *WandappETHTransactorSession) OperatorWithdraw(proof []byte) (*types.Transaction, error) {
	return _WandappETH.Contract.OperatorWithdraw(&_WandappETH.TransactOpts, proof)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_WandappETH *WandappETHTransactor) Settle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WandappETH.contract.Transact(opts, "settle")
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_WandappETH *WandappETHSession) Settle() (*types.Transaction, error) {
	return _WandappETH.Contract.Settle(&_WandappETH.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_WandappETH *WandappETHTransactorSession) Settle() (*types.Transaction, error) {
	return _WandappETH.Contract.Settle(&_WandappETH.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(proof bytes) returns()
func (_WandappETH *WandappETHTransactor) Withdraw(opts *bind.TransactOpts, proof []byte) (*types.Transaction, error) {
	return _WandappETH.contract.Transact(opts, "withdraw", proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(proof bytes) returns()
func (_WandappETH *WandappETHSession) Withdraw(proof []byte) (*types.Transaction, error) {
	return _WandappETH.Contract.Withdraw(&_WandappETH.TransactOpts, proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(proof bytes) returns()
func (_WandappETH *WandappETHTransactorSession) Withdraw(proof []byte) (*types.Transaction, error) {
	return _WandappETH.Contract.Withdraw(&_WandappETH.TransactOpts, proof)
}

// WandappETHBadProofEventIterator is returned from FilterBadProofEvent and is used to iterate over the raw logs and unpacked data for BadProofEvent events raised by the WandappETH contract.
type WandappETHBadProofEventIterator struct {
	Event *WandappETHBadProofEvent // Event containing the contract specifics and raw log

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
func (it *WandappETHBadProofEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHBadProofEvent)
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
		it.Event = new(WandappETHBadProofEvent)
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
func (it *WandappETHBadProofEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHBadProofEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHBadProofEvent represents a BadProofEvent event raised by the WandappETH contract.
type WandappETHBadProofEvent struct {
	From        common.Address
	To          common.Address
	Operator    common.Address
	TransferCnt *big.Int
	WithdrawCnt *big.Int
	Nonce       uint64
	TimeStamp   int64
	ChlID       *big.Int
	Errcode     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBadProofEvent is a free log retrieval operation binding the contract event 0x9b797295befa0a1ed24f6cb1cdbb99f4fcba5989b39a7d79108a6282f4039a9b.
//
// Solidity: e BadProofEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, chlID uint256, errcode uint8)
func (_WandappETH *WandappETHFilterer) FilterBadProofEvent(opts *bind.FilterOpts) (*WandappETHBadProofEventIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "BadProofEvent")
	if err != nil {
		return nil, err
	}
	return &WandappETHBadProofEventIterator{contract: _WandappETH.contract, event: "BadProofEvent", logs: logs, sub: sub}, nil
}

// WatchBadProofEvent is a free log subscription operation binding the contract event 0x9b797295befa0a1ed24f6cb1cdbb99f4fcba5989b39a7d79108a6282f4039a9b.
//
// Solidity: e BadProofEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, chlID uint256, errcode uint8)
func (_WandappETH *WandappETHFilterer) WatchBadProofEvent(opts *bind.WatchOpts, sink chan<- *WandappETHBadProofEvent) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "BadProofEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHBadProofEvent)
				if err := _WandappETH.contract.UnpackLog(event, "BadProofEvent", log); err != nil {
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

// WandappETHDepositeEventIterator is returned from FilterDepositeEvent and is used to iterate over the raw logs and unpacked data for DepositeEvent events raised by the WandappETH contract.
type WandappETHDepositeEventIterator struct {
	Event *WandappETHDepositeEvent // Event containing the contract specifics and raw log

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
func (it *WandappETHDepositeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHDepositeEvent)
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
		it.Event = new(WandappETHDepositeEvent)
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
func (it *WandappETHDepositeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHDepositeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHDepositeEvent represents a DepositeEvent event raised by the WandappETH contract.
type WandappETHDepositeEvent struct {
	ChlId        *big.Int
	Creator      common.Address
	Operator     common.Address
	TokenAddr    common.Address
	Amount       *big.Int
	TotalBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositeEvent is a free log retrieval operation binding the contract event 0x6fd22d40cda41d7dbbf23f88347434a3907b286562d061724be4200d1da8dd08.
//
// Solidity: e DepositeEvent(chlId uint256, creator address, operator address, tokenAddr address, amount uint256, totalBalance uint256)
func (_WandappETH *WandappETHFilterer) FilterDepositeEvent(opts *bind.FilterOpts) (*WandappETHDepositeEventIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "DepositeEvent")
	if err != nil {
		return nil, err
	}
	return &WandappETHDepositeEventIterator{contract: _WandappETH.contract, event: "DepositeEvent", logs: logs, sub: sub}, nil
}

// WatchDepositeEvent is a free log subscription operation binding the contract event 0x6fd22d40cda41d7dbbf23f88347434a3907b286562d061724be4200d1da8dd08.
//
// Solidity: e DepositeEvent(chlId uint256, creator address, operator address, tokenAddr address, amount uint256, totalBalance uint256)
func (_WandappETH *WandappETHFilterer) WatchDepositeEvent(opts *bind.WatchOpts, sink chan<- *WandappETHDepositeEvent) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "DepositeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHDepositeEvent)
				if err := _WandappETH.contract.UnpackLog(event, "DepositeEvent", log); err != nil {
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

// WandappETHExitEventIterator is returned from FilterExitEvent and is used to iterate over the raw logs and unpacked data for ExitEvent events raised by the WandappETH contract.
type WandappETHExitEventIterator struct {
	Event *WandappETHExitEvent // Event containing the contract specifics and raw log

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
func (it *WandappETHExitEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHExitEvent)
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
		it.Event = new(WandappETHExitEvent)
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
func (it *WandappETHExitEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHExitEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHExitEvent represents a ExitEvent event raised by the WandappETH contract.
type WandappETHExitEvent struct {
	From        common.Address
	To          common.Address
	Operator    common.Address
	TransferCnt *big.Int
	WithdrawCnt *big.Int
	Nonce       uint64
	TimeStamp   int64
	ChlID       *big.Int
	Chlstate    uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExitEvent is a free log retrieval operation binding the contract event 0x90b0279af0b82a8471a26044997dc3e95eb00db377daa4806cacba5730f07c15.
//
// Solidity: e ExitEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp int64, chlID uint256, chlstate uint8)
func (_WandappETH *WandappETHFilterer) FilterExitEvent(opts *bind.FilterOpts) (*WandappETHExitEventIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "ExitEvent")
	if err != nil {
		return nil, err
	}
	return &WandappETHExitEventIterator{contract: _WandappETH.contract, event: "ExitEvent", logs: logs, sub: sub}, nil
}

// WatchExitEvent is a free log subscription operation binding the contract event 0x90b0279af0b82a8471a26044997dc3e95eb00db377daa4806cacba5730f07c15.
//
// Solidity: e ExitEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp int64, chlID uint256, chlstate uint8)
func (_WandappETH *WandappETHFilterer) WatchExitEvent(opts *bind.WatchOpts, sink chan<- *WandappETHExitEvent) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "ExitEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHExitEvent)
				if err := _WandappETH.contract.UnpackLog(event, "ExitEvent", log); err != nil {
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

// WandappETHHashEventIterator is returned from FilterHashEvent and is used to iterate over the raw logs and unpacked data for HashEvent events raised by the WandappETH contract.
type WandappETHHashEventIterator struct {
	Event *WandappETHHashEvent // Event containing the contract specifics and raw log

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
func (it *WandappETHHashEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHHashEvent)
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
		it.Event = new(WandappETHHashEvent)
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
func (it *WandappETHHashEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHHashEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHHashEvent represents a HashEvent event raised by the WandappETH contract.
type WandappETHHashEvent struct {
	Index [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHashEvent is a free log retrieval operation binding the contract event 0x45fffb12f7c572056d713ff90194e4559d57f06a2f6b91e50da7cf0f3b5824bc.
//
// Solidity: e HashEvent(index bytes32)
func (_WandappETH *WandappETHFilterer) FilterHashEvent(opts *bind.FilterOpts) (*WandappETHHashEventIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "HashEvent")
	if err != nil {
		return nil, err
	}
	return &WandappETHHashEventIterator{contract: _WandappETH.contract, event: "HashEvent", logs: logs, sub: sub}, nil
}

// WatchHashEvent is a free log subscription operation binding the contract event 0x45fffb12f7c572056d713ff90194e4559d57f06a2f6b91e50da7cf0f3b5824bc.
//
// Solidity: e HashEvent(index bytes32)
func (_WandappETH *WandappETHFilterer) WatchHashEvent(opts *bind.WatchOpts, sink chan<- *WandappETHHashEvent) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "HashEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHHashEvent)
				if err := _WandappETH.contract.UnpackLog(event, "HashEvent", log); err != nil {
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

// WandappETHSettleEventIterator is returned from FilterSettleEvent and is used to iterate over the raw logs and unpacked data for SettleEvent events raised by the WandappETH contract.
type WandappETHSettleEventIterator struct {
	Event *WandappETHSettleEvent // Event containing the contract specifics and raw log

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
func (it *WandappETHSettleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHSettleEvent)
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
		it.Event = new(WandappETHSettleEvent)
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
func (it *WandappETHSettleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHSettleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHSettleEvent represents a SettleEvent event raised by the WandappETH contract.
type WandappETHSettleEvent struct {
	To        common.Address
	Value     *big.Int
	UserNonce *big.Int
	OperNonce *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSettleEvent is a free log retrieval operation binding the contract event 0xf299e08d054ee3c65880d58af77cc2d5e3eb14272c78cc3edb4aa9ce19c07591.
//
// Solidity: e SettleEvent(to address, value uint256, userNonce uint256, operNonce uint256)
func (_WandappETH *WandappETHFilterer) FilterSettleEvent(opts *bind.FilterOpts) (*WandappETHSettleEventIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "SettleEvent")
	if err != nil {
		return nil, err
	}
	return &WandappETHSettleEventIterator{contract: _WandappETH.contract, event: "SettleEvent", logs: logs, sub: sub}, nil
}

// WatchSettleEvent is a free log subscription operation binding the contract event 0xf299e08d054ee3c65880d58af77cc2d5e3eb14272c78cc3edb4aa9ce19c07591.
//
// Solidity: e SettleEvent(to address, value uint256, userNonce uint256, operNonce uint256)
func (_WandappETH *WandappETHFilterer) WatchSettleEvent(opts *bind.WatchOpts, sink chan<- *WandappETHSettleEvent) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "SettleEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHSettleEvent)
				if err := _WandappETH.contract.UnpackLog(event, "SettleEvent", log); err != nil {
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

// WandappETHWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the WandappETH contract.
type WandappETHWithdrawEventIterator struct {
	Event *WandappETHWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *WandappETHWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHWithdrawEvent)
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
		it.Event = new(WandappETHWithdrawEvent)
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
func (it *WandappETHWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHWithdrawEvent represents a WithdrawEvent event raised by the WandappETH contract.
type WandappETHWithdrawEvent struct {
	From          common.Address
	To            common.Address
	Operator      common.Address
	TransferCnt   *big.Int
	WithdrawCnt   *big.Int
	Nonce         uint64
	TimeStamp     int64
	Deposition    *big.Int
	TotalTransfer *big.Int
	TotalWithdraw *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0xb4aff9b5153482df5e297bc29c40cf9031b4c4750fb492217c2235c9ead50dc3.
//
// Solidity: e WithdrawEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, deposition uint256, totalTransfer uint256, totalWithdraw uint256)
func (_WandappETH *WandappETHFilterer) FilterWithdrawEvent(opts *bind.FilterOpts) (*WandappETHWithdrawEventIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return &WandappETHWithdrawEventIterator{contract: _WandappETH.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0xb4aff9b5153482df5e297bc29c40cf9031b4c4750fb492217c2235c9ead50dc3.
//
// Solidity: e WithdrawEvent(from address, to address, operator address, transferCnt uint256, withdrawCnt uint256, nonce uint64, timeStamp_ int64, deposition uint256, totalTransfer uint256, totalWithdraw uint256)
func (_WandappETH *WandappETHFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *WandappETHWithdrawEvent) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "WithdrawEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHWithdrawEvent)
				if err := _WandappETH.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// WandappETHWrongAuthIterator is returned from FilterWrongAuth and is used to iterate over the raw logs and unpacked data for WrongAuth events raised by the WandappETH contract.
type WandappETHWrongAuthIterator struct {
	Event *WandappETHWrongAuth // Event containing the contract specifics and raw log

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
func (it *WandappETHWrongAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WandappETHWrongAuth)
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
		it.Event = new(WandappETHWrongAuth)
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
func (it *WandappETHWrongAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WandappETHWrongAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WandappETHWrongAuth represents a WrongAuth event raised by the WandappETH contract.
type WandappETHWrongAuth struct {
	Sender common.Address
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWrongAuth is a free log retrieval operation binding the contract event 0x71cc3f58d82300f5840dd39e0ab6f3ea51b58a0edc6f77e7a0c74a99ddf9a3c6.
//
// Solidity: e WrongAuth(sender address, target address)
func (_WandappETH *WandappETHFilterer) FilterWrongAuth(opts *bind.FilterOpts) (*WandappETHWrongAuthIterator, error) {

	logs, sub, err := _WandappETH.contract.FilterLogs(opts, "WrongAuth")
	if err != nil {
		return nil, err
	}
	return &WandappETHWrongAuthIterator{contract: _WandappETH.contract, event: "WrongAuth", logs: logs, sub: sub}, nil
}

// WatchWrongAuth is a free log subscription operation binding the contract event 0x71cc3f58d82300f5840dd39e0ab6f3ea51b58a0edc6f77e7a0c74a99ddf9a3c6.
//
// Solidity: e WrongAuth(sender address, target address)
func (_WandappETH *WandappETHFilterer) WatchWrongAuth(opts *bind.WatchOpts, sink chan<- *WandappETHWrongAuth) (event.Subscription, error) {

	logs, sub, err := _WandappETH.contract.WatchLogs(opts, "WrongAuth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WandappETHWrongAuth)
				if err := _WandappETH.contract.UnpackLog(event, "WrongAuth", log); err != nil {
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
