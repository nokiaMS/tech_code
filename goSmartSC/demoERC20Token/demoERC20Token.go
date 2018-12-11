// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demoERC20Token

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
const DemoERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// DemoERC20Bin is the compiled bytecode used for deploying new contracts.
const DemoERC20Bin = `0x608060405234801561001057600080fd5b506040805180820190915260028082527f475800000000000000000000000000000000000000000000000000000000000060209092019182526100559160009161013f565b506040805180820190915260038082527f4023400000000000000000000000000000000000000000000000000000000000602090920191825261009a9160019161013f565b506002805460ff19166008179055629896806003556004602052620f42407fa0c2436e9fcad9824cfe275d95bb4732f44c2903b3014cc507d46698c7adf4478190557f879cbbb7e8fc6520e96255287a54b06cf0085bf0f91ead8c6903dd9d5b10a67381905573b4173fddaa6e5a3b496460ba440cff0f984b98b86000527fbd7b651ce5c78cba55bc837347e80409a8936c79debf80740e260931fe7d5d03556101da565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061018057805160ff19168380011785556101ad565b828001600101855582156101ad579182015b828111156101ad578251825591602001919060010190610192565b506101b99291506101bd565b5090565b6101d791905b808211156101b957600081556001016101c3565b90565b6105a6806101e96000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461009d578063095ea7b31461012757806318160ddd1461015f57806323b872dd14610186578063313ce567146101b057806370a08231146101db57806395d89b41146101fc578063a9059cbb14610211578063dd62ed3e14610235575b600080fd5b3480156100a957600080fd5b506100b261025c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100ec5781810151838201526020016100d4565b50505050905090810190601f1680156101195780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013357600080fd5b5061014b600160a060020a03600435166024356102f2565b604080519115158252519081900360200190f35b34801561016b57600080fd5b50610174610359565b60408051918252519081900360200190f35b34801561019257600080fd5b5061014b600160a060020a036004358116906024351660443561035f565b3480156101bc57600080fd5b506101c5610440565b6040805160ff9092168252519081900360200190f35b3480156101e757600080fd5b50610174600160a060020a0360043516610449565b34801561020857600080fd5b506100b2610464565b34801561021d57600080fd5b5061014b600160a060020a03600435166024356104c4565b34801561024157600080fd5b50610174600160a060020a036004358116906024351661054f565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102e85780601f106102bd576101008083540402835291602001916102e8565b820191906000526020600020905b8154815290600101906020018083116102cb57829003601f168201915b5050505050905090565b336000818152600560209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b60035490565b600160a060020a03831660009081526004602052604081205482118015906103aa5750600160a060020a03841660009081526005602090815260408083203384529091529020548211155b1561043557600160a060020a03808416600081815260046020908152604080832080548801905593881680835284832080548890039055600582528483203384528252918490208054879003905583518681529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3506001610439565b5060005b9392505050565b60025460ff1690565b600160a060020a031660009081526004602052604090205490565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156102e85780601f106102bd576101008083540402835291602001916102e8565b3360009081526004602052604081205482116105475733600081815260046020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001610353565b506000610353565b600160a060020a039182166000908152600560209081526040808320939094168252919091522054905600a165627a7a723058207a55b1c865c2dcf6b93639ea7428cbf8eca96f4a9695b020d5b3f39405f91cec0029`

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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_DemoERC20 *DemoERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DemoERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_DemoERC20 *DemoERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DemoERC20.Contract.Allowance(&_DemoERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_DemoERC20 *DemoERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DemoERC20.Contract.Allowance(&_DemoERC20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_DemoERC20 *DemoERC20Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DemoERC20.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_DemoERC20 *DemoERC20Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DemoERC20.Contract.BalanceOf(&_DemoERC20.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_DemoERC20 *DemoERC20CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DemoERC20.Contract.BalanceOf(&_DemoERC20.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DemoERC20 *DemoERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DemoERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DemoERC20 *DemoERC20Session) Decimals() (uint8, error) {
	return _DemoERC20.Contract.Decimals(&_DemoERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DemoERC20 *DemoERC20CallerSession) Decimals() (uint8, error) {
	return _DemoERC20.Contract.Decimals(&_DemoERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DemoERC20 *DemoERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DemoERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DemoERC20 *DemoERC20Session) Name() (string, error) {
	return _DemoERC20.Contract.Name(&_DemoERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DemoERC20 *DemoERC20CallerSession) Name() (string, error) {
	return _DemoERC20.Contract.Name(&_DemoERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DemoERC20 *DemoERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DemoERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DemoERC20 *DemoERC20Session) Symbol() (string, error) {
	return _DemoERC20.Contract.Symbol(&_DemoERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DemoERC20 *DemoERC20CallerSession) Symbol() (string, error) {
	return _DemoERC20.Contract.Symbol(&_DemoERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DemoERC20 *DemoERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DemoERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DemoERC20 *DemoERC20Session) TotalSupply() (*big.Int, error) {
	return _DemoERC20.Contract.TotalSupply(&_DemoERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DemoERC20 *DemoERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _DemoERC20.Contract.TotalSupply(&_DemoERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Approve(&_DemoERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Approve(&_DemoERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Transfer(&_DemoERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.Transfer(&_DemoERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.TransferFrom(&_DemoERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_DemoERC20 *DemoERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DemoERC20.Contract.TransferFrom(&_DemoERC20.TransactOpts, _from, _to, _value)
}

// DemoERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DemoERC20 contract.
type DemoERC20ApprovalIterator struct {
	Event *DemoERC20Approval // Event containing the contract specifics and raw log

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
func (it *DemoERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DemoERC20Approval)
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
		it.Event = new(DemoERC20Approval)
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
func (it *DemoERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DemoERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DemoERC20Approval represents a Approval event raised by the DemoERC20 contract.
type DemoERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_DemoERC20 *DemoERC20Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*DemoERC20ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DemoERC20.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &DemoERC20ApprovalIterator{contract: _DemoERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_DemoERC20 *DemoERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DemoERC20Approval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DemoERC20.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DemoERC20Approval)
				if err := _DemoERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// DemoERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DemoERC20 contract.
type DemoERC20TransferIterator struct {
	Event *DemoERC20Transfer // Event containing the contract specifics and raw log

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
func (it *DemoERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DemoERC20Transfer)
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
		it.Event = new(DemoERC20Transfer)
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
func (it *DemoERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DemoERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DemoERC20Transfer represents a Transfer event raised by the DemoERC20 contract.
type DemoERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_DemoERC20 *DemoERC20Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*DemoERC20TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DemoERC20.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &DemoERC20TransferIterator{contract: _DemoERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_DemoERC20 *DemoERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DemoERC20Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DemoERC20.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DemoERC20Transfer)
				if err := _DemoERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
