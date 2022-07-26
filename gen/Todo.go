// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	MsgID   *big.Int
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_MsgID\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MsgID\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"MsgID\",\"type\":\"uint256\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"MsgID\",\"type\":\"uint256\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506109a3806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630f560cd71461005157806336555b851461006f5780638da5cb5b1461008b5780639507d39a146100a9575b600080fd5b6100596100d9565b6040516100669190610603565b60405180910390f35b6100896004803603810190610084919061079a565b6101d4565b005b610093610299565b6040516100a09190610837565b60405180910390f35b6100c360048036038101906100be9190610852565b6102bd565b6040516100d091906108bc565b60405180910390f35b60606001805480602002602001604051908101604052809291908181526020016000905b828210156101cb57838290600052602060002090600202016040518060400160405290816000820180546101309061090d565b80601f016020809104026020016040519081016040528092919081815260200182805461015c9061090d565b80156101a95780601f1061017e576101008083540402835291602001916101a9565b820191906000526020600020905b81548152906001019060200180831161018c57829003601f168201915b50505050508152602001600182015481525050815260200190600101906100fd565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461022c57600080fd5b600160405180604001604052808481526020018381525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000019080519060200190610288929190610395565b506020820151816001015550505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6102c561041b565b600182815481106102d9576102d861093e565b5b90600052602060002090600202016040518060400160405290816000820180546103029061090d565b80601f016020809104026020016040519081016040528092919081815260200182805461032e9061090d565b801561037b5780601f106103505761010080835404028352916020019161037b565b820191906000526020600020905b81548152906001019060200180831161035e57829003601f168201915b505050505081526020016001820154815250509050919050565b8280546103a19061090d565b90600052602060002090601f0160209004810192826103c3576000855561040a565b82601f106103dc57805160ff191683800117855561040a565b8280016001018555821561040a579182015b828111156104095782518255916020019190600101906103ee565b5b5090506104179190610435565b5090565b604051806040016040528060608152602001600081525090565b5b8082111561044e576000816000905550600101610436565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104b857808201518184015260208101905061049d565b838111156104c7576000848401525b50505050565b6000601f19601f8301169050919050565b60006104e98261047e565b6104f38185610489565b935061050381856020860161049a565b61050c816104cd565b840191505092915050565b6000819050919050565b61052a81610517565b82525050565b6000604083016000830151848203600086015261054d82826104de565b91505060208301516105626020860182610521565b508091505092915050565b60006105798383610530565b905092915050565b6000602082019050919050565b600061059982610452565b6105a3818561045d565b9350836020820285016105b58561046e565b8060005b858110156105f157848403895281516105d2858261056d565b94506105dd83610581565b925060208a019950506001810190506105b9565b50829750879550505050505092915050565b6000602082019050818103600083015261061d818461058e565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61067b826104cd565b810181811067ffffffffffffffff8211171561069a57610699610643565b5b80604052505050565b60006106ad610625565b90506106b98282610672565b919050565b600067ffffffffffffffff8211156106d9576106d8610643565b5b6106e2826104cd565b9050602081019050919050565b82818337600083830152505050565b600061071161070c846106be565b6106a3565b90508281526020810184848401111561072d5761072c61063e565b5b6107388482856106ef565b509392505050565b600082601f83011261075557610754610639565b5b81356107658482602086016106fe565b91505092915050565b61077781610517565b811461078257600080fd5b50565b6000813590506107948161076e565b92915050565b600080604083850312156107b1576107b061062f565b5b600083013567ffffffffffffffff8111156107cf576107ce610634565b5b6107db85828601610740565b92505060206107ec85828601610785565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610821826107f6565b9050919050565b61083181610816565b82525050565b600060208201905061084c6000830184610828565b92915050565b6000602082840312156108685761086761062f565b5b600061087684828501610785565b91505092915050565b6000604083016000830151848203600086015261089c82826104de565b91505060208301516108b16020860182610521565b508091505092915050565b600060208201905081810360008301526108d6818461087f565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061092557607f821691505b602082108103610938576109376108de565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220be1bb9391bb48591d04eddc51381437dc51f05834ec5f575eed6c966e0d07d9364736f6c634300080e0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _MsgID) view returns((string,uint256))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _MsgID *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _MsgID)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _MsgID) view returns((string,uint256))
func (_Todo *TodoSession) Get(_MsgID *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _MsgID)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _MsgID) view returns((string,uint256))
func (_Todo *TodoCallerSession) Get(_MsgID *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _MsgID)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,uint256)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,uint256)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,uint256)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x36555b85.
//
// Solidity: function add(string _content, uint256 _MsgID) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string, _MsgID *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content, _MsgID)
}

// Add is a paid mutator transaction binding the contract method 0x36555b85.
//
// Solidity: function add(string _content, uint256 _MsgID) returns()
func (_Todo *TodoSession) Add(_content string, _MsgID *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content, _MsgID)
}

// Add is a paid mutator transaction binding the contract method 0x36555b85.
//
// Solidity: function add(string _content, uint256 _MsgID) returns()
func (_Todo *TodoTransactorSession) Add(_content string, _MsgID *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content, _MsgID)
}
