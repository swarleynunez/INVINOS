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

// EntityInfoMetaData contains all meta data concerning the EntityInfo contract.
var EntityInfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"companies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"created\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"containers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"created\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"createCompany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"createContainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_info\",\"type\":\"string\"}],\"name\":\"createProductType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"existCompany\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"existContainer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"existProductType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"productTypes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"created\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c09380380610c0983398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610b76806100936000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063a14d489d1161005b578063a14d489d1461019f578063a844babb146101b2578063f14513f2146101c557600080fd5b80638da5cb5b146101475780638e01a4b21461018c57600080fd5b80637a3d0277116100a75780637a3d02771461010c5780638534d3011461011f578063857d386b1461013457600080fd5b80634c5231d4146100c3578063763852e9146100eb575b600080fd5b6100d66100d136600461084a565b6101d8565b60405190151581526020015b60405180910390f35b6100fe6100f936600461084a565b610206565b6040516100e29291906108ab565b6100fe61011a36600461084a565b6102b8565b61013261012d366004610904565b6102de565b005b6100d661014236600461084a565b61047c565b6000546101679073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e2565b61013261019a366004610904565b61048e565b6100d66101ad36600461084a565b6105e3565b6100fe6101c036600461084a565b6105f5565b6101326101d3366004610904565b61061b565b60006001826040516101ea9190610968565b9081526040519081900360200190206001015460ff1692915050565b805160208183018101805160038252928201919093012091528054819061022c90610984565b80601f016020809104026020016040519081016040528092919081815260200182805461025890610984565b80156102a55780601f1061027a576101008083540402835291602001916102a5565b820191906000526020600020905b81548152906001019060200180831161028857829003601f168201915b5050506001909301549192505060ff1682565b805160208183018101805160028252928201919093012091528054819061022c90610984565b60005473ffffffffffffffffffffffffffffffffffffffff16331461038a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c60448201527f20746869732066756e6374696f6e00000000000000000000000000000000000060648201526084015b60405180910390fd5b610393826101d8565b156103fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f50726f64756374207479706520616c72656164792063726561746564000000006044820152606401610381565b8060018360405161040b9190610968565b908152604051908190036020019020906104259082610a26565b50600180836040516104379190610968565b90815260405190819003602001902060010180549115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009092169190911790555050565b60006002826040516101ea9190610968565b60005473ffffffffffffffffffffffffffffffffffffffff163314610535576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c60448201527f20746869732066756e6374696f6e0000000000000000000000000000000000006064820152608401610381565b61053e826105e3565b156105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f436f6e7461696e657220616c72656164792063726561746564000000000000006044820152606401610381565b806003836040516105b69190610968565b908152604051908190036020019020906105d09082610a26565b5060016003836040516104379190610968565b60006003826040516101ea9190610968565b805160208183018101805160018252928201919093012091528054819061022c90610984565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c60448201527f20746869732066756e6374696f6e0000000000000000000000000000000000006064820152608401610381565b6106cb8261047c565b15610732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f436f6d70616e7920616c726561647920637265617465640000000000000000006044820152606401610381565b806002836040516107439190610968565b9081526040519081900360200190209061075d9082610a26565b5060016002836040516104379190610968565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126107b057600080fd5b813567ffffffffffffffff808211156107cb576107cb610770565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561081157610811610770565b8160405283815286602085880101111561082a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561085c57600080fd5b813567ffffffffffffffff81111561087357600080fd5b61087f8482850161079f565b949350505050565b60005b838110156108a257818101518382015260200161088a565b50506000910152565b60408152600083518060408401526108ca816060850160208801610887565b921515602083015250601f919091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601606001919050565b6000806040838503121561091757600080fd5b823567ffffffffffffffff8082111561092f57600080fd5b61093b8683870161079f565b9350602085013591508082111561095157600080fd5b5061095e8582860161079f565b9150509250929050565b6000825161097a818460208701610887565b9190910192915050565b600181811c9082168061099857607f821691505b6020821081036109d1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115610a2157600081815260208120601f850160051c810160208610156109fe5750805b601f850160051c820191505b81811015610a1d57828155600101610a0a565b5050505b505050565b815167ffffffffffffffff811115610a4057610a40610770565b610a5481610a4e8454610984565b846109d7565b602080601f831160018114610aa75760008415610a715750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610a1d565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015610af457888601518255948401946001909101908401610ad5565b5085821015610b3057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212200503eab396577f46b1f6cf18e54b3315ae782c88ea84d36809706709060404f664736f6c63430008150033",
}

// EntityInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use EntityInfoMetaData.ABI instead.
var EntityInfoABI = EntityInfoMetaData.ABI

// EntityInfoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EntityInfoMetaData.Bin instead.
var EntityInfoBin = EntityInfoMetaData.Bin

// DeployEntityInfo deploys a new Ethereum contract, binding an instance of EntityInfo to it.
func DeployEntityInfo(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *EntityInfo, error) {
	parsed, err := EntityInfoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EntityInfoBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EntityInfo{EntityInfoCaller: EntityInfoCaller{contract: contract}, EntityInfoTransactor: EntityInfoTransactor{contract: contract}, EntityInfoFilterer: EntityInfoFilterer{contract: contract}}, nil
}

// EntityInfo is an auto generated Go binding around an Ethereum contract.
type EntityInfo struct {
	EntityInfoCaller     // Read-only binding to the contract
	EntityInfoTransactor // Write-only binding to the contract
	EntityInfoFilterer   // Log filterer for contract events
}

// EntityInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntityInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntityInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntityInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntityInfoSession struct {
	Contract     *EntityInfo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntityInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntityInfoCallerSession struct {
	Contract *EntityInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EntityInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntityInfoTransactorSession struct {
	Contract     *EntityInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EntityInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntityInfoRaw struct {
	Contract *EntityInfo // Generic contract binding to access the raw methods on
}

// EntityInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntityInfoCallerRaw struct {
	Contract *EntityInfoCaller // Generic read-only contract binding to access the raw methods on
}

// EntityInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntityInfoTransactorRaw struct {
	Contract *EntityInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntityInfo creates a new instance of EntityInfo, bound to a specific deployed contract.
func NewEntityInfo(address common.Address, backend bind.ContractBackend) (*EntityInfo, error) {
	contract, err := bindEntityInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntityInfo{EntityInfoCaller: EntityInfoCaller{contract: contract}, EntityInfoTransactor: EntityInfoTransactor{contract: contract}, EntityInfoFilterer: EntityInfoFilterer{contract: contract}}, nil
}

// NewEntityInfoCaller creates a new read-only instance of EntityInfo, bound to a specific deployed contract.
func NewEntityInfoCaller(address common.Address, caller bind.ContractCaller) (*EntityInfoCaller, error) {
	contract, err := bindEntityInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntityInfoCaller{contract: contract}, nil
}

// NewEntityInfoTransactor creates a new write-only instance of EntityInfo, bound to a specific deployed contract.
func NewEntityInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*EntityInfoTransactor, error) {
	contract, err := bindEntityInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntityInfoTransactor{contract: contract}, nil
}

// NewEntityInfoFilterer creates a new log filterer instance of EntityInfo, bound to a specific deployed contract.
func NewEntityInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*EntityInfoFilterer, error) {
	contract, err := bindEntityInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntityInfoFilterer{contract: contract}, nil
}

// bindEntityInfo binds a generic wrapper to an already deployed contract.
func bindEntityInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EntityInfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityInfo *EntityInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntityInfo.Contract.EntityInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityInfo *EntityInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityInfo.Contract.EntityInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityInfo *EntityInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityInfo.Contract.EntityInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityInfo *EntityInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntityInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityInfo *EntityInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityInfo *EntityInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityInfo.Contract.contract.Transact(opts, method, params...)
}

// Companies is a free data retrieval call binding the contract method 0x7a3d0277.
//
// Solidity: function companies(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoCaller) Companies(opts *bind.CallOpts, arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "companies", arg0)

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

// Companies is a free data retrieval call binding the contract method 0x7a3d0277.
//
// Solidity: function companies(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoSession) Companies(arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	return _EntityInfo.Contract.Companies(&_EntityInfo.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x7a3d0277.
//
// Solidity: function companies(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoCallerSession) Companies(arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	return _EntityInfo.Contract.Companies(&_EntityInfo.CallOpts, arg0)
}

// Containers is a free data retrieval call binding the contract method 0x763852e9.
//
// Solidity: function containers(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoCaller) Containers(opts *bind.CallOpts, arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "containers", arg0)

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

// Containers is a free data retrieval call binding the contract method 0x763852e9.
//
// Solidity: function containers(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoSession) Containers(arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	return _EntityInfo.Contract.Containers(&_EntityInfo.CallOpts, arg0)
}

// Containers is a free data retrieval call binding the contract method 0x763852e9.
//
// Solidity: function containers(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoCallerSession) Containers(arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	return _EntityInfo.Contract.Containers(&_EntityInfo.CallOpts, arg0)
}

// ExistCompany is a free data retrieval call binding the contract method 0x857d386b.
//
// Solidity: function existCompany(string _id) view returns(bool)
func (_EntityInfo *EntityInfoCaller) ExistCompany(opts *bind.CallOpts, _id string) (bool, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "existCompany", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistCompany is a free data retrieval call binding the contract method 0x857d386b.
//
// Solidity: function existCompany(string _id) view returns(bool)
func (_EntityInfo *EntityInfoSession) ExistCompany(_id string) (bool, error) {
	return _EntityInfo.Contract.ExistCompany(&_EntityInfo.CallOpts, _id)
}

// ExistCompany is a free data retrieval call binding the contract method 0x857d386b.
//
// Solidity: function existCompany(string _id) view returns(bool)
func (_EntityInfo *EntityInfoCallerSession) ExistCompany(_id string) (bool, error) {
	return _EntityInfo.Contract.ExistCompany(&_EntityInfo.CallOpts, _id)
}

// ExistContainer is a free data retrieval call binding the contract method 0xa14d489d.
//
// Solidity: function existContainer(string _id) view returns(bool)
func (_EntityInfo *EntityInfoCaller) ExistContainer(opts *bind.CallOpts, _id string) (bool, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "existContainer", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistContainer is a free data retrieval call binding the contract method 0xa14d489d.
//
// Solidity: function existContainer(string _id) view returns(bool)
func (_EntityInfo *EntityInfoSession) ExistContainer(_id string) (bool, error) {
	return _EntityInfo.Contract.ExistContainer(&_EntityInfo.CallOpts, _id)
}

// ExistContainer is a free data retrieval call binding the contract method 0xa14d489d.
//
// Solidity: function existContainer(string _id) view returns(bool)
func (_EntityInfo *EntityInfoCallerSession) ExistContainer(_id string) (bool, error) {
	return _EntityInfo.Contract.ExistContainer(&_EntityInfo.CallOpts, _id)
}

// ExistProductType is a free data retrieval call binding the contract method 0x4c5231d4.
//
// Solidity: function existProductType(string _id) view returns(bool)
func (_EntityInfo *EntityInfoCaller) ExistProductType(opts *bind.CallOpts, _id string) (bool, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "existProductType", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistProductType is a free data retrieval call binding the contract method 0x4c5231d4.
//
// Solidity: function existProductType(string _id) view returns(bool)
func (_EntityInfo *EntityInfoSession) ExistProductType(_id string) (bool, error) {
	return _EntityInfo.Contract.ExistProductType(&_EntityInfo.CallOpts, _id)
}

// ExistProductType is a free data retrieval call binding the contract method 0x4c5231d4.
//
// Solidity: function existProductType(string _id) view returns(bool)
func (_EntityInfo *EntityInfoCallerSession) ExistProductType(_id string) (bool, error) {
	return _EntityInfo.Contract.ExistProductType(&_EntityInfo.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EntityInfo *EntityInfoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EntityInfo *EntityInfoSession) Owner() (common.Address, error) {
	return _EntityInfo.Contract.Owner(&_EntityInfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EntityInfo *EntityInfoCallerSession) Owner() (common.Address, error) {
	return _EntityInfo.Contract.Owner(&_EntityInfo.CallOpts)
}

// ProductTypes is a free data retrieval call binding the contract method 0xa844babb.
//
// Solidity: function productTypes(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoCaller) ProductTypes(opts *bind.CallOpts, arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	var out []interface{}
	err := _EntityInfo.contract.Call(opts, &out, "productTypes", arg0)

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

// ProductTypes is a free data retrieval call binding the contract method 0xa844babb.
//
// Solidity: function productTypes(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoSession) ProductTypes(arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	return _EntityInfo.Contract.ProductTypes(&_EntityInfo.CallOpts, arg0)
}

// ProductTypes is a free data retrieval call binding the contract method 0xa844babb.
//
// Solidity: function productTypes(string ) view returns(string info, bool created)
func (_EntityInfo *EntityInfoCallerSession) ProductTypes(arg0 string) (struct {
	Info    string
	Created bool
}, error) {
	return _EntityInfo.Contract.ProductTypes(&_EntityInfo.CallOpts, arg0)
}

// CreateCompany is a paid mutator transaction binding the contract method 0xf14513f2.
//
// Solidity: function createCompany(string _id, string _info) returns()
func (_EntityInfo *EntityInfoTransactor) CreateCompany(opts *bind.TransactOpts, _id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.contract.Transact(opts, "createCompany", _id, _info)
}

// CreateCompany is a paid mutator transaction binding the contract method 0xf14513f2.
//
// Solidity: function createCompany(string _id, string _info) returns()
func (_EntityInfo *EntityInfoSession) CreateCompany(_id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.Contract.CreateCompany(&_EntityInfo.TransactOpts, _id, _info)
}

// CreateCompany is a paid mutator transaction binding the contract method 0xf14513f2.
//
// Solidity: function createCompany(string _id, string _info) returns()
func (_EntityInfo *EntityInfoTransactorSession) CreateCompany(_id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.Contract.CreateCompany(&_EntityInfo.TransactOpts, _id, _info)
}

// CreateContainer is a paid mutator transaction binding the contract method 0x8e01a4b2.
//
// Solidity: function createContainer(string _id, string _info) returns()
func (_EntityInfo *EntityInfoTransactor) CreateContainer(opts *bind.TransactOpts, _id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.contract.Transact(opts, "createContainer", _id, _info)
}

// CreateContainer is a paid mutator transaction binding the contract method 0x8e01a4b2.
//
// Solidity: function createContainer(string _id, string _info) returns()
func (_EntityInfo *EntityInfoSession) CreateContainer(_id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.Contract.CreateContainer(&_EntityInfo.TransactOpts, _id, _info)
}

// CreateContainer is a paid mutator transaction binding the contract method 0x8e01a4b2.
//
// Solidity: function createContainer(string _id, string _info) returns()
func (_EntityInfo *EntityInfoTransactorSession) CreateContainer(_id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.Contract.CreateContainer(&_EntityInfo.TransactOpts, _id, _info)
}

// CreateProductType is a paid mutator transaction binding the contract method 0x8534d301.
//
// Solidity: function createProductType(string _id, string _info) returns()
func (_EntityInfo *EntityInfoTransactor) CreateProductType(opts *bind.TransactOpts, _id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.contract.Transact(opts, "createProductType", _id, _info)
}

// CreateProductType is a paid mutator transaction binding the contract method 0x8534d301.
//
// Solidity: function createProductType(string _id, string _info) returns()
func (_EntityInfo *EntityInfoSession) CreateProductType(_id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.Contract.CreateProductType(&_EntityInfo.TransactOpts, _id, _info)
}

// CreateProductType is a paid mutator transaction binding the contract method 0x8534d301.
//
// Solidity: function createProductType(string _id, string _info) returns()
func (_EntityInfo *EntityInfoTransactorSession) CreateProductType(_id string, _info string) (*types.Transaction, error) {
	return _EntityInfo.Contract.CreateProductType(&_EntityInfo.TransactOpts, _id, _info)
}
