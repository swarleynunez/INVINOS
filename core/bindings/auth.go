// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// AuthMetaData contains all meta data concerning the Auth contract.
var AuthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transitionTypes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"created\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161069b38038061069b83398101604081905261002f91610296565b600080546001600160a01b0319166001600160a01b03831617815560408051808201909152600d81526c50524f445543545f454e54525960981b60208083019190915291805260019091527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49906100a69082610365565b507fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4a805460ff1916600190811790915560408051808201909152601281527150524f445543545f50524f43455353494e4760701b6020828101919091526000839052919091527fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f906101389082610365565b507fcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b6887930805460ff19166001908117909155604080518082019091526011815270282927a22aa1aa2fa820a92a24aa24a7a760791b6020828101919091526002600052919091527fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f906101c99082610365565b507fd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec3310805460ff1916600190811790915560408051808201909152600e81526d141493d11550d517d3d55514155560921b6020828101919091526003600052919091527f7dfe757ecd65cbd7922a9c0161e935dd7fdbcc0e999689c7d31633896b1fc60b906102579082610365565b50506003600052600160208190527f7dfe757ecd65cbd7922a9c0161e935dd7fdbcc0e999689c7d31633896b1fc60c805460ff19169091179055610424565b6000602082840312156102a857600080fd5b81516001600160a01b03811681146102bf57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806102f057607f821691505b60208210810361031057634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561036057600081815260208120601f850160051c8101602086101561033d5750805b601f850160051c820191505b8181101561035c57828155600101610349565b5050505b505050565b81516001600160401b0381111561037e5761037e6102c6565b6103928161038c84546102dc565b84610316565b602080601f8311600181146103c757600084156103af5750858301515b600019600386901b1c1916600185901b17855561035c565b600085815260208120601f198616915b828110156103f6578886015182559484019460019091019084016103d7565b50858210156104145787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610268806104336000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637f81cbe91461003b5780638da5cb5b14610065575b600080fd5b61004e610049366004610151565b6100aa565b60405161005c92919061016a565b60405180910390f35b6000546100859073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161005c565b6001602052600090815260409020805481906100c5906101df565b80601f01602080910402602001604051908101604052809291908181526020018280546100f1906101df565b801561013e5780601f106101135761010080835404028352916020019161013e565b820191906000526020600020905b81548152906001019060200180831161012157829003601f168201915b5050506001909301549192505060ff1682565b60006020828403121561016357600080fd5b5035919050565b604081526000835180604084015260005b81811015610198576020818701810151606086840101520161017b565b5060006060828501015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505082151560208301529392505050565b600181811c908216806101f357607f821691505b60208210810361022c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea2646970667358221220561e1b99d61262ea5b8f2aea56157a2fa264bf7e4c06fa2dadfc09de6677ae6d64736f6c63430008150033",
}

// AuthABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthMetaData.ABI instead.
var AuthABI = AuthMetaData.ABI

// AuthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthMetaData.Bin instead.
var AuthBin = AuthMetaData.Bin

// DeployAuth deploys a new Ethereum contract, binding an instance of Auth to it.
func DeployAuth(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Auth, error) {
	parsed, err := AuthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// Auth is an auto generated Go binding around an Ethereum contract.
type Auth struct {
	AuthCaller     // Read-only binding to the contract
	AuthTransactor // Write-only binding to the contract
	AuthFilterer   // Log filterer for contract events
}

// AuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthSession struct {
	Contract     *Auth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthCallerSession struct {
	Contract *AuthCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthTransactorSession struct {
	Contract     *AuthTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthRaw struct {
	Contract *Auth // Generic contract binding to access the raw methods on
}

// AuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthCallerRaw struct {
	Contract *AuthCaller // Generic read-only contract binding to access the raw methods on
}

// AuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthTransactorRaw struct {
	Contract *AuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuth creates a new instance of Auth, bound to a specific deployed contract.
func NewAuth(address common.Address, backend bind.ContractBackend) (*Auth, error) {
	contract, err := bindAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// NewAuthCaller creates a new read-only instance of Auth, bound to a specific deployed contract.
func NewAuthCaller(address common.Address, caller bind.ContractCaller) (*AuthCaller, error) {
	contract, err := bindAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthCaller{contract: contract}, nil
}

// NewAuthTransactor creates a new write-only instance of Auth, bound to a specific deployed contract.
func NewAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthTransactor, error) {
	contract, err := bindAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthTransactor{contract: contract}, nil
}

// NewAuthFilterer creates a new log filterer instance of Auth, bound to a specific deployed contract.
func NewAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthFilterer, error) {
	contract, err := bindAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthFilterer{contract: contract}, nil
}

// bindAuth binds a generic wrapper to an already deployed contract.
func bindAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.AuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auth *AuthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auth *AuthSession) Owner() (common.Address, error) {
	return _Auth.Contract.Owner(&_Auth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auth *AuthCallerSession) Owner() (common.Address, error) {
	return _Auth.Contract.Owner(&_Auth.CallOpts)
}

// TransitionTypes is a free data retrieval call binding the contract method 0x7f81cbe9.
//
// Solidity: function transitionTypes(uint256 ) view returns(string info, bool created)
func (_Auth *AuthCaller) TransitionTypes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Info    string
	Created bool
}, error) {
	var out []interface{}
	err := _Auth.contract.Call(opts, &out, "transitionTypes", arg0)

	outstruct := new(struct {
		Info    string
		Created bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Info = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Created = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TransitionTypes is a free data retrieval call binding the contract method 0x7f81cbe9.
//
// Solidity: function transitionTypes(uint256 ) view returns(string info, bool created)
func (_Auth *AuthSession) TransitionTypes(arg0 *big.Int) (struct {
	Info    string
	Created bool
}, error) {
	return _Auth.Contract.TransitionTypes(&_Auth.CallOpts, arg0)
}

// TransitionTypes is a free data retrieval call binding the contract method 0x7f81cbe9.
//
// Solidity: function transitionTypes(uint256 ) view returns(string info, bool created)
func (_Auth *AuthCallerSession) TransitionTypes(arg0 *big.Int) (struct {
	Info    string
	Created bool
}, error) {
	return _Auth.Contract.TransitionTypes(&_Auth.CallOpts, arg0)
}
