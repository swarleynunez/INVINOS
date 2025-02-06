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
	Bin: "0x608060405234801561001057600080fd5b50604051610b36380380610b3683398101604081905261002f91610054565b60008054600160a060020a031916600160a060020a0392909216919091179055610082565b600060208284031215610065578081fd5b8151600160a060020a038116811461007b578182fd5b9392505050565b610aa5806100916000396000f3fe608060405234801561001057600080fd5b50600436106100bb576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b116100835780638da5cb5b146101455780638e01a4b21461015a578063a14d489d1461016d578063a844babb14610180578063f14513f214610193576100bb565b80634c5231d4146100c0578063763852e9146100e95780637a3d02771461010a5780638534d3011461011d578063857d386b14610132575b600080fd5b6100d36100ce3660046106f6565b6101a6565b6040516100e091906107cf565b60405180910390f35b6100fc6100f73660046106f6565b6101d5565b6040516100e09291906107da565b6100fc6101183660046106f6565b610287565b61013061012b366004610731565b6102ad565b005b6100d36101403660046106f6565b6103b8565b61014d6103ca565b6040516100e091906107ae565b610130610168366004610731565b6103e6565b6100d361017b3660046106f6565b6104c2565b6100fc61018e3660046106f6565b6104d4565b6101306101a1366004610731565b6104fa565b60006001826040516101b89190610792565b9081526040519081900360200190206001015460ff169050919050565b80516020818301810180516003825292820191909301209152805481906101fb906109ec565b80601f0160208091040260200160405190810160405280929190818152602001828054610227906109ec565b80156102745780601f1061024957610100808354040283529160200191610274565b820191906000526020600020905b81548152906001019060200180831161025757829003601f168201915b5050506001909301549192505060ff1682565b80516020818301810180516002825292820191909301209152805481906101fb906109ec565b60005473ffffffffffffffffffffffffffffffffffffffff1633146102f05760405160e560020a62461bcd0281526004016102e790610883565b60405180910390fd5b6102f9826101a6565b156103195760405160e560020a62461bcd0281526004016102e790610815565b6000825111801561032b575060008151115b61034a5760405160e560020a62461bcd0281526004016102e79061084c565b8060018360405161035b9190610792565b9081526020016040518091039020600001908051906020019061037f9291906105d6565b50600180836040516103919190610792565b908152604051908190036020019020600101805491151560ff199092169190911790555050565b60006002826040516101b89190610792565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1633146104205760405160e560020a62461bcd0281526004016102e790610883565b610429826104c2565b156104495760405160e560020a62461bcd0281526004016102e790610917565b6000825111801561045b575060008151115b61047a5760405160e560020a62461bcd0281526004016102e7906108e0565b8060038360405161048b9190610792565b908152602001604051809103902060000190805190602001906104af9291906105d6565b5060016003836040516103919190610792565b60006003826040516101b89190610792565b80516020818301810180516001825292820191909301209152805481906101fb906109ec565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105345760405160e560020a62461bcd0281526004016102e790610883565b61053d826103b8565b1561055d5760405160e560020a62461bcd0281526004016102e79061094e565b6000825111801561056f575060008151115b61058e5760405160e560020a62461bcd0281526004016102e790610985565b8060028360405161059f9190610792565b908152602001604051809103902060000190805190602001906105c39291906105d6565b5060016002836040516103919190610792565b8280546105e2906109ec565b90600052602060002090601f016020900481019282610604576000855561064a565b82601f1061061d57805160ff191683800117855561064a565b8280016001018555821561064a579182015b8281111561064a57825182559160200191906001019061062f565b5061065692915061065a565b5090565b5b80821115610656576000815560010161065b565b600082601f83011261067f578081fd5b813567ffffffffffffffff8082111561069a5761069a610a40565b604051601f8301601f19908116603f011681019082821181831017156106c2576106c2610a40565b816040528381528660208588010111156106da578485fd5b8360208701602083013792830160200193909352509392505050565b600060208284031215610707578081fd5b813567ffffffffffffffff81111561071d578182fd5b6107298482850161066f565b949350505050565b60008060408385031215610743578081fd5b823567ffffffffffffffff8082111561075a578283fd5b6107668683870161066f565b9350602085013591508082111561077b578283fd5b506107888582860161066f565b9150509250929050565b600082516107a48184602087016109bc565b9190910192915050565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b901515815260200190565b60006040825283518060408401526107f98160608501602088016109bc565b921515602083015250601f91909101601f191601606001919050565b6020808252601c908201527f50726f64756374207479706520616c7265616479206372656174656400000000604082015260600190565b60208082526017908201527f456d7074792070726f6475637420747970652064617461000000000000000000604082015260600190565b6020808252602e908201527f4f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c60408201527f20746869732066756e6374696f6e000000000000000000000000000000000000606082015260800190565b60208082526014908201527f456d70747920636f6e7461696e65722064617461000000000000000000000000604082015260600190565b60208082526019908201527f436f6e7461696e657220616c7265616479206372656174656400000000000000604082015260600190565b60208082526017908201527f436f6d70616e7920616c72656164792063726561746564000000000000000000604082015260600190565b60208082526012908201527f456d70747920636f6d70616e7920646174610000000000000000000000000000604082015260600190565b60005b838110156109d75781810151838201526020016109bf565b838111156109e6576000848401525b50505050565b600281046001821680610a0057607f821691505b60208210811415610a3a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea264697066735822122034a3894e696e4912dece1404900af30edf6da4099995ec8fe26a3e5ef1826c3f64736f6c63430008010033",
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
