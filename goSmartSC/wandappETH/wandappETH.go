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

// WandappETHABI is the input ABI used to generate the binding from.
const WandappETHABI = "[]=======./proofVerify.sol:proofVerify=======[]=======./wandappBase.sol:DemoERC20=======[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]=======./wandappBase.sol:wandappBase=======[{\"constant\":false,\"inputs\":[],\"name\":\"getChlID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenOrigAddr\",\"type\":\"address\"},{\"name\":\"operAddr\",\"type\":\"address\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"HashEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"deposition\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTransfer\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"chlID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errcode\",\"type\":\"uint8\"}],\"name\":\"BadProofEvent\",\"type\":\"event\"}]=======./wandappETH.sol:wandappETH=======[{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getChlID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"operatorWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenOrigAddr\",\"type\":\"address\"},{\"name\":\"operAddr\",\"type\":\"address\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"chlId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBalance\",\"type\":\"uint256\"}],\"name\":\"DepositeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"chlID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"chlstate\",\"type\":\"uint8\"}],\"name\":\"ExitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"target\",\"type\":\"address\"}],\"name\":\"WrongAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"operNonce\",\"type\":\"uint256\"}],\"name\":\"SettleEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"HashEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"deposition\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTransfer\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawCnt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timeStamp_\",\"type\":\"int64\"},{\"indexed\":false,\"name\":\"chlID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errcode\",\"type\":\"uint8\"}],\"name\":\"BadProofEvent\",\"type\":\"event\"}]"

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
