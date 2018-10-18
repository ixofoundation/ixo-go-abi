// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// BasicProjectWalletABI is the input ABI used to generate the binding from.
const BasicProjectWalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_authoriser\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BasicProjectWalletBin is the compiled bytecode used for deploying new contracts.
const BasicProjectWalletBin = `0x608060405234801561001057600080fd5b5060405160608061029983398101604090815281516020830151919092015160008054600160a060020a03948516600160a060020a03199182161790915560018054949093169316929092179055600255610229806100706000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610050578063a9059cbb14610077575b600080fd5b34801561005c57600080fd5b506100656100bc565b60408051918252519081900360200190f35b34801561008357600080fd5b506100a873ffffffffffffffffffffffffffffffffffffffff600435166024356100c2565b604080519115158252519081900360200190f35b60025481565b60015460009073ffffffffffffffffffffffffffffffffffffffff16331461014b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b505050506040513d60208110156101f457600080fd5b509093925050505600a165627a7a72305820a57f44df56894921b295f5260067fe533d856755c79ff5af9d7e46f1324c122b0029`

// DeployBasicProjectWallet deploys a new Ethereum contract, binding an instance of BasicProjectWallet to it.
func DeployBasicProjectWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _authoriser common.Address, _name [32]byte) (common.Address, *types.Transaction, *BasicProjectWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicProjectWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicProjectWalletBin), backend, _token, _authoriser, _name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicProjectWallet{BasicProjectWalletCaller: BasicProjectWalletCaller{contract: contract}, BasicProjectWalletTransactor: BasicProjectWalletTransactor{contract: contract}, BasicProjectWalletFilterer: BasicProjectWalletFilterer{contract: contract}}, nil
}

// BasicProjectWallet is an auto generated Go binding around an Ethereum contract.
type BasicProjectWallet struct {
	BasicProjectWalletCaller     // Read-only binding to the contract
	BasicProjectWalletTransactor // Write-only binding to the contract
	BasicProjectWalletFilterer   // Log filterer for contract events
}

// BasicProjectWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicProjectWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicProjectWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicProjectWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicProjectWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicProjectWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicProjectWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicProjectWalletSession struct {
	Contract     *BasicProjectWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasicProjectWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicProjectWalletCallerSession struct {
	Contract *BasicProjectWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BasicProjectWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicProjectWalletTransactorSession struct {
	Contract     *BasicProjectWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BasicProjectWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicProjectWalletRaw struct {
	Contract *BasicProjectWallet // Generic contract binding to access the raw methods on
}

// BasicProjectWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicProjectWalletCallerRaw struct {
	Contract *BasicProjectWalletCaller // Generic read-only contract binding to access the raw methods on
}

// BasicProjectWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicProjectWalletTransactorRaw struct {
	Contract *BasicProjectWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicProjectWallet creates a new instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWallet(address common.Address, backend bind.ContractBackend) (*BasicProjectWallet, error) {
	contract, err := bindBasicProjectWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWallet{BasicProjectWalletCaller: BasicProjectWalletCaller{contract: contract}, BasicProjectWalletTransactor: BasicProjectWalletTransactor{contract: contract}, BasicProjectWalletFilterer: BasicProjectWalletFilterer{contract: contract}}, nil
}

// NewBasicProjectWalletCaller creates a new read-only instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWalletCaller(address common.Address, caller bind.ContractCaller) (*BasicProjectWalletCaller, error) {
	contract, err := bindBasicProjectWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWalletCaller{contract: contract}, nil
}

// NewBasicProjectWalletTransactor creates a new write-only instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicProjectWalletTransactor, error) {
	contract, err := bindBasicProjectWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWalletTransactor{contract: contract}, nil
}

// NewBasicProjectWalletFilterer creates a new log filterer instance of BasicProjectWallet, bound to a specific deployed contract.
func NewBasicProjectWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicProjectWalletFilterer, error) {
	contract, err := bindBasicProjectWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicProjectWalletFilterer{contract: contract}, nil
}

// bindBasicProjectWallet binds a generic wrapper to an already deployed contract.
func bindBasicProjectWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicProjectWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicProjectWallet *BasicProjectWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicProjectWallet.Contract.BasicProjectWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicProjectWallet *BasicProjectWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.BasicProjectWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicProjectWallet *BasicProjectWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.BasicProjectWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicProjectWallet *BasicProjectWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicProjectWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicProjectWallet *BasicProjectWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicProjectWallet *BasicProjectWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_BasicProjectWallet *BasicProjectWalletCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BasicProjectWallet.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_BasicProjectWallet *BasicProjectWalletSession) Name() ([32]byte, error) {
	return _BasicProjectWallet.Contract.Name(&_BasicProjectWallet.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_BasicProjectWallet *BasicProjectWalletCallerSession) Name() ([32]byte, error) {
	return _BasicProjectWallet.Contract.Name(&_BasicProjectWallet.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_BasicProjectWallet *BasicProjectWalletTransactor) Transfer(opts *bind.TransactOpts, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _BasicProjectWallet.contract.Transact(opts, "transfer", _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_BasicProjectWallet *BasicProjectWalletSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.Transfer(&_BasicProjectWallet.TransactOpts, _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_BasicProjectWallet *BasicProjectWalletTransactorSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _BasicProjectWallet.Contract.Transfer(&_BasicProjectWallet.TransactOpts, _receiver, _amt)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ProjectWalletABI is the input ABI used to generate the binding from.
const ProjectWalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProjectWalletBin is the compiled bytecode used for deploying new contracts.
const ProjectWalletBin = `0x`

// DeployProjectWallet deploys a new Ethereum contract, binding an instance of ProjectWallet to it.
func DeployProjectWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProjectWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProjectWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProjectWallet{ProjectWalletCaller: ProjectWalletCaller{contract: contract}, ProjectWalletTransactor: ProjectWalletTransactor{contract: contract}, ProjectWalletFilterer: ProjectWalletFilterer{contract: contract}}, nil
}

// ProjectWallet is an auto generated Go binding around an Ethereum contract.
type ProjectWallet struct {
	ProjectWalletCaller     // Read-only binding to the contract
	ProjectWalletTransactor // Write-only binding to the contract
	ProjectWalletFilterer   // Log filterer for contract events
}

// ProjectWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectWalletSession struct {
	Contract     *ProjectWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectWalletCallerSession struct {
	Contract *ProjectWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProjectWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectWalletTransactorSession struct {
	Contract     *ProjectWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProjectWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectWalletRaw struct {
	Contract *ProjectWallet // Generic contract binding to access the raw methods on
}

// ProjectWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectWalletCallerRaw struct {
	Contract *ProjectWalletCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectWalletTransactorRaw struct {
	Contract *ProjectWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectWallet creates a new instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWallet(address common.Address, backend bind.ContractBackend) (*ProjectWallet, error) {
	contract, err := bindProjectWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectWallet{ProjectWalletCaller: ProjectWalletCaller{contract: contract}, ProjectWalletTransactor: ProjectWalletTransactor{contract: contract}, ProjectWalletFilterer: ProjectWalletFilterer{contract: contract}}, nil
}

// NewProjectWalletCaller creates a new read-only instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWalletCaller(address common.Address, caller bind.ContractCaller) (*ProjectWalletCaller, error) {
	contract, err := bindProjectWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletCaller{contract: contract}, nil
}

// NewProjectWalletTransactor creates a new write-only instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectWalletTransactor, error) {
	contract, err := bindProjectWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletTransactor{contract: contract}, nil
}

// NewProjectWalletFilterer creates a new log filterer instance of ProjectWallet, bound to a specific deployed contract.
func NewProjectWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectWalletFilterer, error) {
	contract, err := bindProjectWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFilterer{contract: contract}, nil
}

// bindProjectWallet binds a generic wrapper to an already deployed contract.
func bindProjectWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWallet *ProjectWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWallet.Contract.ProjectWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWallet *ProjectWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWallet.Contract.ProjectWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWallet *ProjectWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWallet.Contract.ProjectWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWallet *ProjectWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWallet *ProjectWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWallet *ProjectWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWallet.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_ProjectWallet *ProjectWalletTransactor) Transfer(opts *bind.TransactOpts, _receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWallet.contract.Transact(opts, "transfer", _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_ProjectWallet *ProjectWalletSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWallet.Contract.Transfer(&_ProjectWallet.TransactOpts, _receiver, _amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_receiver address, _amt uint256) returns(bool)
func (_ProjectWallet *ProjectWalletTransactorSession) Transfer(_receiver common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _ProjectWallet.Contract.Transfer(&_ProjectWallet.TransactOpts, _receiver, _amt)
}

// ProjectWalletFactoryABI is the input ABI used to generate the binding from.
const ProjectWalletFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_authoriser\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProjectWalletFactoryBin is the compiled bytecode used for deploying new contracts.
const ProjectWalletFactoryBin = `0x608060405234801561001057600080fd5b50610491806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663e8eef2708114610045575b600080fd5b34801561005157600080fd5b5061007c73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356100a5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000807fff000000000000000000000000000000000000000000000000000000000000007f010000000000000000000000000000000000000000000000000000000000000084831a0216151561015c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f496e76616c6964206e616d650000000000000000000000000000000000000000604482015290519081900360640190fd5b8484846101676101bc565b73ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152604080820192909252905190819003606001906000f0801580156101b2573d6000803e3d6000fd5b5095945050505050565b604051610299806101cd833901905600608060405234801561001057600080fd5b5060405160608061029983398101604090815281516020830151919092015160008054600160a060020a03948516600160a060020a03199182161790915560018054949093169316929092179055600255610229806100706000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610050578063a9059cbb14610077575b600080fd5b34801561005c57600080fd5b506100656100bc565b60408051918252519081900360200190f35b34801561008357600080fd5b506100a873ffffffffffffffffffffffffffffffffffffffff600435166024356100c2565b604080519115158252519081900360200190f35b60025481565b60015460009073ffffffffffffffffffffffffffffffffffffffff16331461014b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5065726d697373696f6e2064656e696564000000000000000000000000000000604482015290519081900360640190fd5b60008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b505050506040513d60208110156101f457600080fd5b509093925050505600a165627a7a72305820a57f44df56894921b295f5260067fe533d856755c79ff5af9d7e46f1324c122b0029a165627a7a723058206b16a3976414825fc4ed35e8a22423602f47b8e1f8d4090ecbb2782879a6f0b40029`

// DeployProjectWalletFactory deploys a new Ethereum contract, binding an instance of ProjectWalletFactory to it.
func DeployProjectWalletFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProjectWalletFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProjectWalletFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProjectWalletFactory{ProjectWalletFactoryCaller: ProjectWalletFactoryCaller{contract: contract}, ProjectWalletFactoryTransactor: ProjectWalletFactoryTransactor{contract: contract}, ProjectWalletFactoryFilterer: ProjectWalletFactoryFilterer{contract: contract}}, nil
}

// ProjectWalletFactory is an auto generated Go binding around an Ethereum contract.
type ProjectWalletFactory struct {
	ProjectWalletFactoryCaller     // Read-only binding to the contract
	ProjectWalletFactoryTransactor // Write-only binding to the contract
	ProjectWalletFactoryFilterer   // Log filterer for contract events
}

// ProjectWalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectWalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectWalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectWalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectWalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectWalletFactorySession struct {
	Contract     *ProjectWalletFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProjectWalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectWalletFactoryCallerSession struct {
	Contract *ProjectWalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ProjectWalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectWalletFactoryTransactorSession struct {
	Contract     *ProjectWalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ProjectWalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectWalletFactoryRaw struct {
	Contract *ProjectWalletFactory // Generic contract binding to access the raw methods on
}

// ProjectWalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectWalletFactoryCallerRaw struct {
	Contract *ProjectWalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectWalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectWalletFactoryTransactorRaw struct {
	Contract *ProjectWalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectWalletFactory creates a new instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactory(address common.Address, backend bind.ContractBackend) (*ProjectWalletFactory, error) {
	contract, err := bindProjectWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactory{ProjectWalletFactoryCaller: ProjectWalletFactoryCaller{contract: contract}, ProjectWalletFactoryTransactor: ProjectWalletFactoryTransactor{contract: contract}, ProjectWalletFactoryFilterer: ProjectWalletFactoryFilterer{contract: contract}}, nil
}

// NewProjectWalletFactoryCaller creates a new read-only instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*ProjectWalletFactoryCaller, error) {
	contract, err := bindProjectWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactoryCaller{contract: contract}, nil
}

// NewProjectWalletFactoryTransactor creates a new write-only instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectWalletFactoryTransactor, error) {
	contract, err := bindProjectWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactoryTransactor{contract: contract}, nil
}

// NewProjectWalletFactoryFilterer creates a new log filterer instance of ProjectWalletFactory, bound to a specific deployed contract.
func NewProjectWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectWalletFactoryFilterer, error) {
	contract, err := bindProjectWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectWalletFactoryFilterer{contract: contract}, nil
}

// bindProjectWalletFactory binds a generic wrapper to an already deployed contract.
func bindProjectWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectWalletFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletFactory *ProjectWalletFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletFactory.Contract.ProjectWalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletFactory *ProjectWalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.ProjectWalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletFactory *ProjectWalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.ProjectWalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectWalletFactory *ProjectWalletFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProjectWalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectWalletFactory *ProjectWalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectWalletFactory *ProjectWalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(_token address, _authoriser address, _name bytes32) returns(address)
func (_ProjectWalletFactory *ProjectWalletFactoryTransactor) CreateWallet(opts *bind.TransactOpts, _token common.Address, _authoriser common.Address, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletFactory.contract.Transact(opts, "createWallet", _token, _authoriser, _name)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(_token address, _authoriser address, _name bytes32) returns(address)
func (_ProjectWalletFactory *ProjectWalletFactorySession) CreateWallet(_token common.Address, _authoriser common.Address, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.CreateWallet(&_ProjectWalletFactory.TransactOpts, _token, _authoriser, _name)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(_token address, _authoriser address, _name bytes32) returns(address)
func (_ProjectWalletFactory *ProjectWalletFactoryTransactorSession) CreateWallet(_token common.Address, _authoriser common.Address, _name [32]byte) (*types.Transaction, error) {
	return _ProjectWalletFactory.Contract.CreateWallet(&_ProjectWalletFactory.TransactOpts, _token, _authoriser, _name)
}
