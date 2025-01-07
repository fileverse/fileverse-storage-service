// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fileverseportalabi

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

// PortalKeyVerifiersKeyVerifier is an auto generated low-level Go binding around an user-defined struct.
type PortalKeyVerifiersKeyVerifier struct {
	PortalEncryptionKeyVerifier [32]byte
	PortalDecryptionKeyVerifier [32]byte
	MemberEncryptionKeyVerifier [32]byte
	MemberDecryptionKeyVerifier [32]byte
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ownerViewDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ownerEditDid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"portalEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"portalDecryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"memberEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"memberDecryptionKeyVerifier\",\"type\":\"bytes32\"}],\"internalType\":\"structPortalKeyVerifiers.KeyVerifier\",\"name\":\"_keyVerifier\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"AddedCollaborator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataIPFSHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentIPFSHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gateIPFSHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"AddedFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataIPFSHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentIPFSHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gateIPFSHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"EditedFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RegisteredCollaboratorKeys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"RemovedCollaborator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemovedCollaboratorKeys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"portalEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"portalDecryptionKeyVerifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"memberEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"memberDecryptionKeyVerifier\",\"type\":\"bytes32\"}],\"name\":\"UpdatedKeyVerifiers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataIPFSHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"UpdatedPortalMetadata\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collaborator\",\"type\":\"address\"}],\"name\":\"addCollaborator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contentIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_gateIPFSHash\",\"type\":\"string\"},{\"internalType\":\"enumFileversePortal.FileType\",\"name\":\"filetype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"addFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collaboratorKeys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"viewDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"editDid\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_metadataIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contentIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_gateIPFSHash\",\"type\":\"string\"},{\"internalType\":\"enumFileversePortal.FileType\",\"name\":\"filetype\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"editFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"metadataIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contentIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gateIPFSHash\",\"type\":\"string\"},{\"internalType\":\"enumFileversePortal.FileType\",\"name\":\"fileType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaboratorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaboratorKeysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaborators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isCollaborator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keyVerifiers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"portalEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"portalDecryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"memberEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"memberDecryptionKeyVerifier\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataIPFSHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"viewDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"editDid\",\"type\":\"string\"}],\"name\":\"registerCollaboratorKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevCollaborator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collaborator\",\"type\":\"address\"}],\"name\":\"removeCollaborator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeCollaboratorKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"portalEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"portalDecryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"memberEncryptionKeyVerifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"memberDecryptionKeyVerifier\",\"type\":\"bytes32\"}],\"name\":\"updateKeyVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataIPFSHash\",\"type\":\"string\"}],\"name\":\"updateMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// CollaboratorKeys is a free data retrieval call binding the contract method 0x876dbec4.
//
// Solidity: function collaboratorKeys(address ) view returns(string viewDid, string editDid)
func (_Bindings *BindingsCaller) CollaboratorKeys(opts *bind.CallOpts, arg0 common.Address) (struct {
	ViewDid string
	EditDid string
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "collaboratorKeys", arg0)

	outstruct := new(struct {
		ViewDid string
		EditDid string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ViewDid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EditDid = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// CollaboratorKeys is a free data retrieval call binding the contract method 0x876dbec4.
//
// Solidity: function collaboratorKeys(address ) view returns(string viewDid, string editDid)
func (_Bindings *BindingsSession) CollaboratorKeys(arg0 common.Address) (struct {
	ViewDid string
	EditDid string
}, error) {
	return _Bindings.Contract.CollaboratorKeys(&_Bindings.CallOpts, arg0)
}

// CollaboratorKeys is a free data retrieval call binding the contract method 0x876dbec4.
//
// Solidity: function collaboratorKeys(address ) view returns(string viewDid, string editDid)
func (_Bindings *BindingsCallerSession) CollaboratorKeys(arg0 common.Address) (struct {
	ViewDid string
	EditDid string
}, error) {
	return _Bindings.Contract.CollaboratorKeys(&_Bindings.CallOpts, arg0)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, uint8 fileType, uint256 version)
func (_Bindings *BindingsCaller) Files(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MetadataIPFSHash string
	ContentIPFSHash  string
	GateIPFSHash     string
	FileType         uint8
	Version          *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "files", arg0)

	outstruct := new(struct {
		MetadataIPFSHash string
		ContentIPFSHash  string
		GateIPFSHash     string
		FileType         uint8
		Version          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetadataIPFSHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ContentIPFSHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.GateIPFSHash = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.FileType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Version = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, uint8 fileType, uint256 version)
func (_Bindings *BindingsSession) Files(arg0 *big.Int) (struct {
	MetadataIPFSHash string
	ContentIPFSHash  string
	GateIPFSHash     string
	FileType         uint8
	Version          *big.Int
}, error) {
	return _Bindings.Contract.Files(&_Bindings.CallOpts, arg0)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, uint8 fileType, uint256 version)
func (_Bindings *BindingsCallerSession) Files(arg0 *big.Int) (struct {
	MetadataIPFSHash string
	ContentIPFSHash  string
	GateIPFSHash     string
	FileType         uint8
	Version          *big.Int
}, error) {
	return _Bindings.Contract.Files(&_Bindings.CallOpts, arg0)
}

// GetCollaboratorCount is a free data retrieval call binding the contract method 0xa2425927.
//
// Solidity: function getCollaboratorCount() view returns(uint256)
func (_Bindings *BindingsCaller) GetCollaboratorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getCollaboratorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollaboratorCount is a free data retrieval call binding the contract method 0xa2425927.
//
// Solidity: function getCollaboratorCount() view returns(uint256)
func (_Bindings *BindingsSession) GetCollaboratorCount() (*big.Int, error) {
	return _Bindings.Contract.GetCollaboratorCount(&_Bindings.CallOpts)
}

// GetCollaboratorCount is a free data retrieval call binding the contract method 0xa2425927.
//
// Solidity: function getCollaboratorCount() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetCollaboratorCount() (*big.Int, error) {
	return _Bindings.Contract.GetCollaboratorCount(&_Bindings.CallOpts)
}

// GetCollaboratorKeysCount is a free data retrieval call binding the contract method 0x66cc227d.
//
// Solidity: function getCollaboratorKeysCount() view returns(uint256)
func (_Bindings *BindingsCaller) GetCollaboratorKeysCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getCollaboratorKeysCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollaboratorKeysCount is a free data retrieval call binding the contract method 0x66cc227d.
//
// Solidity: function getCollaboratorKeysCount() view returns(uint256)
func (_Bindings *BindingsSession) GetCollaboratorKeysCount() (*big.Int, error) {
	return _Bindings.Contract.GetCollaboratorKeysCount(&_Bindings.CallOpts)
}

// GetCollaboratorKeysCount is a free data retrieval call binding the contract method 0x66cc227d.
//
// Solidity: function getCollaboratorKeysCount() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetCollaboratorKeysCount() (*big.Int, error) {
	return _Bindings.Contract.GetCollaboratorKeysCount(&_Bindings.CallOpts)
}

// GetCollaborators is a free data retrieval call binding the contract method 0xffad0fc6.
//
// Solidity: function getCollaborators() view returns(address[])
func (_Bindings *BindingsCaller) GetCollaborators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getCollaborators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollaborators is a free data retrieval call binding the contract method 0xffad0fc6.
//
// Solidity: function getCollaborators() view returns(address[])
func (_Bindings *BindingsSession) GetCollaborators() ([]common.Address, error) {
	return _Bindings.Contract.GetCollaborators(&_Bindings.CallOpts)
}

// GetCollaborators is a free data retrieval call binding the contract method 0xffad0fc6.
//
// Solidity: function getCollaborators() view returns(address[])
func (_Bindings *BindingsCallerSession) GetCollaborators() ([]common.Address, error) {
	return _Bindings.Contract.GetCollaborators(&_Bindings.CallOpts)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Bindings *BindingsCaller) GetFileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getFileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Bindings *BindingsSession) GetFileCount() (*big.Int, error) {
	return _Bindings.Contract.GetFileCount(&_Bindings.CallOpts)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetFileCount() (*big.Int, error) {
	return _Bindings.Contract.GetFileCount(&_Bindings.CallOpts)
}

// IsCollaborator is a free data retrieval call binding the contract method 0x6af8c4e7.
//
// Solidity: function isCollaborator(address account) view returns(bool)
func (_Bindings *BindingsCaller) IsCollaborator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isCollaborator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollaborator is a free data retrieval call binding the contract method 0x6af8c4e7.
//
// Solidity: function isCollaborator(address account) view returns(bool)
func (_Bindings *BindingsSession) IsCollaborator(account common.Address) (bool, error) {
	return _Bindings.Contract.IsCollaborator(&_Bindings.CallOpts, account)
}

// IsCollaborator is a free data retrieval call binding the contract method 0x6af8c4e7.
//
// Solidity: function isCollaborator(address account) view returns(bool)
func (_Bindings *BindingsCallerSession) IsCollaborator(account common.Address) (bool, error) {
	return _Bindings.Contract.IsCollaborator(&_Bindings.CallOpts, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Bindings *BindingsCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Bindings *BindingsSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Bindings.Contract.IsTrustedForwarder(&_Bindings.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Bindings *BindingsCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Bindings.Contract.IsTrustedForwarder(&_Bindings.CallOpts, forwarder)
}

// KeyVerifiers is a free data retrieval call binding the contract method 0x2df1742a.
//
// Solidity: function keyVerifiers(uint256 ) view returns(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier)
func (_Bindings *BindingsCaller) KeyVerifiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PortalEncryptionKeyVerifier [32]byte
	PortalDecryptionKeyVerifier [32]byte
	MemberEncryptionKeyVerifier [32]byte
	MemberDecryptionKeyVerifier [32]byte
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "keyVerifiers", arg0)

	outstruct := new(struct {
		PortalEncryptionKeyVerifier [32]byte
		PortalDecryptionKeyVerifier [32]byte
		MemberEncryptionKeyVerifier [32]byte
		MemberDecryptionKeyVerifier [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PortalEncryptionKeyVerifier = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.PortalDecryptionKeyVerifier = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.MemberEncryptionKeyVerifier = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.MemberDecryptionKeyVerifier = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// KeyVerifiers is a free data retrieval call binding the contract method 0x2df1742a.
//
// Solidity: function keyVerifiers(uint256 ) view returns(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier)
func (_Bindings *BindingsSession) KeyVerifiers(arg0 *big.Int) (struct {
	PortalEncryptionKeyVerifier [32]byte
	PortalDecryptionKeyVerifier [32]byte
	MemberEncryptionKeyVerifier [32]byte
	MemberDecryptionKeyVerifier [32]byte
}, error) {
	return _Bindings.Contract.KeyVerifiers(&_Bindings.CallOpts, arg0)
}

// KeyVerifiers is a free data retrieval call binding the contract method 0x2df1742a.
//
// Solidity: function keyVerifiers(uint256 ) view returns(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier)
func (_Bindings *BindingsCallerSession) KeyVerifiers(arg0 *big.Int) (struct {
	PortalEncryptionKeyVerifier [32]byte
	PortalDecryptionKeyVerifier [32]byte
	MemberEncryptionKeyVerifier [32]byte
	MemberDecryptionKeyVerifier [32]byte
}, error) {
	return _Bindings.Contract.KeyVerifiers(&_Bindings.CallOpts, arg0)
}

// MetadataIPFSHash is a free data retrieval call binding the contract method 0x93ce2036.
//
// Solidity: function metadataIPFSHash() view returns(string)
func (_Bindings *BindingsCaller) MetadataIPFSHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "metadataIPFSHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataIPFSHash is a free data retrieval call binding the contract method 0x93ce2036.
//
// Solidity: function metadataIPFSHash() view returns(string)
func (_Bindings *BindingsSession) MetadataIPFSHash() (string, error) {
	return _Bindings.Contract.MetadataIPFSHash(&_Bindings.CallOpts)
}

// MetadataIPFSHash is a free data retrieval call binding the contract method 0x93ce2036.
//
// Solidity: function metadataIPFSHash() view returns(string)
func (_Bindings *BindingsCallerSession) MetadataIPFSHash() (string, error) {
	return _Bindings.Contract.MetadataIPFSHash(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bindings *BindingsCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bindings *BindingsSession) PendingOwner() (common.Address, error) {
	return _Bindings.Contract.PendingOwner(&_Bindings.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bindings *BindingsCallerSession) PendingOwner() (common.Address, error) {
	return _Bindings.Contract.PendingOwner(&_Bindings.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bindings *BindingsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bindings *BindingsSession) AcceptOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.AcceptOwnership(&_Bindings.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bindings *BindingsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.AcceptOwnership(&_Bindings.TransactOpts)
}

// AddCollaborator is a paid mutator transaction binding the contract method 0xccdbe6a6.
//
// Solidity: function addCollaborator(address collaborator) returns()
func (_Bindings *BindingsTransactor) AddCollaborator(opts *bind.TransactOpts, collaborator common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addCollaborator", collaborator)
}

// AddCollaborator is a paid mutator transaction binding the contract method 0xccdbe6a6.
//
// Solidity: function addCollaborator(address collaborator) returns()
func (_Bindings *BindingsSession) AddCollaborator(collaborator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddCollaborator(&_Bindings.TransactOpts, collaborator)
}

// AddCollaborator is a paid mutator transaction binding the contract method 0xccdbe6a6.
//
// Solidity: function addCollaborator(address collaborator) returns()
func (_Bindings *BindingsTransactorSession) AddCollaborator(collaborator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddCollaborator(&_Bindings.TransactOpts, collaborator)
}

// AddFile is a paid mutator transaction binding the contract method 0x3bf1bdad.
//
// Solidity: function addFile(string _metadataIPFSHash, string _contentIPFSHash, string _gateIPFSHash, uint8 filetype, uint256 version) returns()
func (_Bindings *BindingsTransactor) AddFile(opts *bind.TransactOpts, _metadataIPFSHash string, _contentIPFSHash string, _gateIPFSHash string, filetype uint8, version *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addFile", _metadataIPFSHash, _contentIPFSHash, _gateIPFSHash, filetype, version)
}

// AddFile is a paid mutator transaction binding the contract method 0x3bf1bdad.
//
// Solidity: function addFile(string _metadataIPFSHash, string _contentIPFSHash, string _gateIPFSHash, uint8 filetype, uint256 version) returns()
func (_Bindings *BindingsSession) AddFile(_metadataIPFSHash string, _contentIPFSHash string, _gateIPFSHash string, filetype uint8, version *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.AddFile(&_Bindings.TransactOpts, _metadataIPFSHash, _contentIPFSHash, _gateIPFSHash, filetype, version)
}

// AddFile is a paid mutator transaction binding the contract method 0x3bf1bdad.
//
// Solidity: function addFile(string _metadataIPFSHash, string _contentIPFSHash, string _gateIPFSHash, uint8 filetype, uint256 version) returns()
func (_Bindings *BindingsTransactorSession) AddFile(_metadataIPFSHash string, _contentIPFSHash string, _gateIPFSHash string, filetype uint8, version *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.AddFile(&_Bindings.TransactOpts, _metadataIPFSHash, _contentIPFSHash, _gateIPFSHash, filetype, version)
}

// EditFile is a paid mutator transaction binding the contract method 0x5eb1cf8f.
//
// Solidity: function editFile(uint256 fileId, string _metadataIPFSHash, string _contentIPFSHash, string _gateIPFSHash, uint8 filetype, uint256 version) returns()
func (_Bindings *BindingsTransactor) EditFile(opts *bind.TransactOpts, fileId *big.Int, _metadataIPFSHash string, _contentIPFSHash string, _gateIPFSHash string, filetype uint8, version *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "editFile", fileId, _metadataIPFSHash, _contentIPFSHash, _gateIPFSHash, filetype, version)
}

// EditFile is a paid mutator transaction binding the contract method 0x5eb1cf8f.
//
// Solidity: function editFile(uint256 fileId, string _metadataIPFSHash, string _contentIPFSHash, string _gateIPFSHash, uint8 filetype, uint256 version) returns()
func (_Bindings *BindingsSession) EditFile(fileId *big.Int, _metadataIPFSHash string, _contentIPFSHash string, _gateIPFSHash string, filetype uint8, version *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.EditFile(&_Bindings.TransactOpts, fileId, _metadataIPFSHash, _contentIPFSHash, _gateIPFSHash, filetype, version)
}

// EditFile is a paid mutator transaction binding the contract method 0x5eb1cf8f.
//
// Solidity: function editFile(uint256 fileId, string _metadataIPFSHash, string _contentIPFSHash, string _gateIPFSHash, uint8 filetype, uint256 version) returns()
func (_Bindings *BindingsTransactorSession) EditFile(fileId *big.Int, _metadataIPFSHash string, _contentIPFSHash string, _gateIPFSHash string, filetype uint8, version *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.EditFile(&_Bindings.TransactOpts, fileId, _metadataIPFSHash, _contentIPFSHash, _gateIPFSHash, filetype, version)
}

// RegisterCollaboratorKeys is a paid mutator transaction binding the contract method 0xcfd9e0c1.
//
// Solidity: function registerCollaboratorKeys(string viewDid, string editDid) returns()
func (_Bindings *BindingsTransactor) RegisterCollaboratorKeys(opts *bind.TransactOpts, viewDid string, editDid string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerCollaboratorKeys", viewDid, editDid)
}

// RegisterCollaboratorKeys is a paid mutator transaction binding the contract method 0xcfd9e0c1.
//
// Solidity: function registerCollaboratorKeys(string viewDid, string editDid) returns()
func (_Bindings *BindingsSession) RegisterCollaboratorKeys(viewDid string, editDid string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterCollaboratorKeys(&_Bindings.TransactOpts, viewDid, editDid)
}

// RegisterCollaboratorKeys is a paid mutator transaction binding the contract method 0xcfd9e0c1.
//
// Solidity: function registerCollaboratorKeys(string viewDid, string editDid) returns()
func (_Bindings *BindingsTransactorSession) RegisterCollaboratorKeys(viewDid string, editDid string) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterCollaboratorKeys(&_Bindings.TransactOpts, viewDid, editDid)
}

// RemoveCollaborator is a paid mutator transaction binding the contract method 0xb4d266bc.
//
// Solidity: function removeCollaborator(address prevCollaborator, address collaborator) returns()
func (_Bindings *BindingsTransactor) RemoveCollaborator(opts *bind.TransactOpts, prevCollaborator common.Address, collaborator common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeCollaborator", prevCollaborator, collaborator)
}

// RemoveCollaborator is a paid mutator transaction binding the contract method 0xb4d266bc.
//
// Solidity: function removeCollaborator(address prevCollaborator, address collaborator) returns()
func (_Bindings *BindingsSession) RemoveCollaborator(prevCollaborator common.Address, collaborator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveCollaborator(&_Bindings.TransactOpts, prevCollaborator, collaborator)
}

// RemoveCollaborator is a paid mutator transaction binding the contract method 0xb4d266bc.
//
// Solidity: function removeCollaborator(address prevCollaborator, address collaborator) returns()
func (_Bindings *BindingsTransactorSession) RemoveCollaborator(prevCollaborator common.Address, collaborator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveCollaborator(&_Bindings.TransactOpts, prevCollaborator, collaborator)
}

// RemoveCollaboratorKeys is a paid mutator transaction binding the contract method 0xde729317.
//
// Solidity: function removeCollaboratorKeys() returns()
func (_Bindings *BindingsTransactor) RemoveCollaboratorKeys(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeCollaboratorKeys")
}

// RemoveCollaboratorKeys is a paid mutator transaction binding the contract method 0xde729317.
//
// Solidity: function removeCollaboratorKeys() returns()
func (_Bindings *BindingsSession) RemoveCollaboratorKeys() (*types.Transaction, error) {
	return _Bindings.Contract.RemoveCollaboratorKeys(&_Bindings.TransactOpts)
}

// RemoveCollaboratorKeys is a paid mutator transaction binding the contract method 0xde729317.
//
// Solidity: function removeCollaboratorKeys() returns()
func (_Bindings *BindingsTransactorSession) RemoveCollaboratorKeys() (*types.Transaction, error) {
	return _Bindings.Contract.RemoveCollaboratorKeys(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// UpdateKeyVerifiers is a paid mutator transaction binding the contract method 0x53689f33.
//
// Solidity: function updateKeyVerifiers(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier) returns()
func (_Bindings *BindingsTransactor) UpdateKeyVerifiers(opts *bind.TransactOpts, portalEncryptionKeyVerifier [32]byte, portalDecryptionKeyVerifier [32]byte, memberEncryptionKeyVerifier [32]byte, memberDecryptionKeyVerifier [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateKeyVerifiers", portalEncryptionKeyVerifier, portalDecryptionKeyVerifier, memberEncryptionKeyVerifier, memberDecryptionKeyVerifier)
}

// UpdateKeyVerifiers is a paid mutator transaction binding the contract method 0x53689f33.
//
// Solidity: function updateKeyVerifiers(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier) returns()
func (_Bindings *BindingsSession) UpdateKeyVerifiers(portalEncryptionKeyVerifier [32]byte, portalDecryptionKeyVerifier [32]byte, memberEncryptionKeyVerifier [32]byte, memberDecryptionKeyVerifier [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateKeyVerifiers(&_Bindings.TransactOpts, portalEncryptionKeyVerifier, portalDecryptionKeyVerifier, memberEncryptionKeyVerifier, memberDecryptionKeyVerifier)
}

// UpdateKeyVerifiers is a paid mutator transaction binding the contract method 0x53689f33.
//
// Solidity: function updateKeyVerifiers(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier) returns()
func (_Bindings *BindingsTransactorSession) UpdateKeyVerifiers(portalEncryptionKeyVerifier [32]byte, portalDecryptionKeyVerifier [32]byte, memberEncryptionKeyVerifier [32]byte, memberDecryptionKeyVerifier [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateKeyVerifiers(&_Bindings.TransactOpts, portalEncryptionKeyVerifier, portalDecryptionKeyVerifier, memberEncryptionKeyVerifier, memberDecryptionKeyVerifier)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x918b5be1.
//
// Solidity: function updateMetadata(string _metadataIPFSHash) returns()
func (_Bindings *BindingsTransactor) UpdateMetadata(opts *bind.TransactOpts, _metadataIPFSHash string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateMetadata", _metadataIPFSHash)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x918b5be1.
//
// Solidity: function updateMetadata(string _metadataIPFSHash) returns()
func (_Bindings *BindingsSession) UpdateMetadata(_metadataIPFSHash string) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateMetadata(&_Bindings.TransactOpts, _metadataIPFSHash)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x918b5be1.
//
// Solidity: function updateMetadata(string _metadataIPFSHash) returns()
func (_Bindings *BindingsTransactorSession) UpdateMetadata(_metadataIPFSHash string) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateMetadata(&_Bindings.TransactOpts, _metadataIPFSHash)
}

// BindingsAddedCollaboratorIterator is returned from FilterAddedCollaborator and is used to iterate over the raw logs and unpacked data for AddedCollaborator events raised by the Bindings contract.
type BindingsAddedCollaboratorIterator struct {
	Event *BindingsAddedCollaborator // Event containing the contract specifics and raw log

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
func (it *BindingsAddedCollaboratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAddedCollaborator)
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
		it.Event = new(BindingsAddedCollaborator)
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
func (it *BindingsAddedCollaboratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAddedCollaboratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAddedCollaborator represents a AddedCollaborator event raised by the Bindings contract.
type BindingsAddedCollaborator struct {
	Account common.Address
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedCollaborator is a free log retrieval operation binding the contract event 0xce61b8b5dd5de3bc40e91f97638290ba020c1f9267272522a018d616f077b166.
//
// Solidity: event AddedCollaborator(address indexed account, address indexed by)
func (_Bindings *BindingsFilterer) FilterAddedCollaborator(opts *bind.FilterOpts, account []common.Address, by []common.Address) (*BindingsAddedCollaboratorIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AddedCollaborator", accountRule, byRule)
	if err != nil {
		return nil, err
	}
	return &BindingsAddedCollaboratorIterator{contract: _Bindings.contract, event: "AddedCollaborator", logs: logs, sub: sub}, nil
}

// WatchAddedCollaborator is a free log subscription operation binding the contract event 0xce61b8b5dd5de3bc40e91f97638290ba020c1f9267272522a018d616f077b166.
//
// Solidity: event AddedCollaborator(address indexed account, address indexed by)
func (_Bindings *BindingsFilterer) WatchAddedCollaborator(opts *bind.WatchOpts, sink chan<- *BindingsAddedCollaborator, account []common.Address, by []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AddedCollaborator", accountRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAddedCollaborator)
				if err := _Bindings.contract.UnpackLog(event, "AddedCollaborator", log); err != nil {
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

// ParseAddedCollaborator is a log parse operation binding the contract event 0xce61b8b5dd5de3bc40e91f97638290ba020c1f9267272522a018d616f077b166.
//
// Solidity: event AddedCollaborator(address indexed account, address indexed by)
func (_Bindings *BindingsFilterer) ParseAddedCollaborator(log types.Log) (*BindingsAddedCollaborator, error) {
	event := new(BindingsAddedCollaborator)
	if err := _Bindings.contract.UnpackLog(event, "AddedCollaborator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsAddedFileIterator is returned from FilterAddedFile and is used to iterate over the raw logs and unpacked data for AddedFile events raised by the Bindings contract.
type BindingsAddedFileIterator struct {
	Event *BindingsAddedFile // Event containing the contract specifics and raw log

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
func (it *BindingsAddedFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAddedFile)
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
		it.Event = new(BindingsAddedFile)
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
func (it *BindingsAddedFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAddedFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAddedFile represents a AddedFile event raised by the Bindings contract.
type BindingsAddedFile struct {
	FileId           *big.Int
	MetadataIPFSHash string
	ContentIPFSHash  string
	GateIPFSHash     string
	By               common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddedFile is a free log retrieval operation binding the contract event 0x7c868b0d1657677f41ebc5d6a76e66b11862720a3800b9b0895b9457c181d8d9.
//
// Solidity: event AddedFile(uint256 indexed fileId, string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) FilterAddedFile(opts *bind.FilterOpts, fileId []*big.Int, by []common.Address) (*BindingsAddedFileIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AddedFile", fileIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &BindingsAddedFileIterator{contract: _Bindings.contract, event: "AddedFile", logs: logs, sub: sub}, nil
}

// WatchAddedFile is a free log subscription operation binding the contract event 0x7c868b0d1657677f41ebc5d6a76e66b11862720a3800b9b0895b9457c181d8d9.
//
// Solidity: event AddedFile(uint256 indexed fileId, string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) WatchAddedFile(opts *bind.WatchOpts, sink chan<- *BindingsAddedFile, fileId []*big.Int, by []common.Address) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AddedFile", fileIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAddedFile)
				if err := _Bindings.contract.UnpackLog(event, "AddedFile", log); err != nil {
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

// ParseAddedFile is a log parse operation binding the contract event 0x7c868b0d1657677f41ebc5d6a76e66b11862720a3800b9b0895b9457c181d8d9.
//
// Solidity: event AddedFile(uint256 indexed fileId, string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) ParseAddedFile(log types.Log) (*BindingsAddedFile, error) {
	event := new(BindingsAddedFile)
	if err := _Bindings.contract.UnpackLog(event, "AddedFile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsEditedFileIterator is returned from FilterEditedFile and is used to iterate over the raw logs and unpacked data for EditedFile events raised by the Bindings contract.
type BindingsEditedFileIterator struct {
	Event *BindingsEditedFile // Event containing the contract specifics and raw log

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
func (it *BindingsEditedFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEditedFile)
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
		it.Event = new(BindingsEditedFile)
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
func (it *BindingsEditedFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEditedFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEditedFile represents a EditedFile event raised by the Bindings contract.
type BindingsEditedFile struct {
	FileId           *big.Int
	MetadataIPFSHash string
	ContentIPFSHash  string
	GateIPFSHash     string
	By               common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEditedFile is a free log retrieval operation binding the contract event 0x038bf72a38023435f249f01545d426ebccb3348739aab6415949426e9c5b1ca0.
//
// Solidity: event EditedFile(uint256 indexed fileId, string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) FilterEditedFile(opts *bind.FilterOpts, fileId []*big.Int, by []common.Address) (*BindingsEditedFileIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EditedFile", fileIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &BindingsEditedFileIterator{contract: _Bindings.contract, event: "EditedFile", logs: logs, sub: sub}, nil
}

// WatchEditedFile is a free log subscription operation binding the contract event 0x038bf72a38023435f249f01545d426ebccb3348739aab6415949426e9c5b1ca0.
//
// Solidity: event EditedFile(uint256 indexed fileId, string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) WatchEditedFile(opts *bind.WatchOpts, sink chan<- *BindingsEditedFile, fileId []*big.Int, by []common.Address) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EditedFile", fileIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEditedFile)
				if err := _Bindings.contract.UnpackLog(event, "EditedFile", log); err != nil {
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

// ParseEditedFile is a log parse operation binding the contract event 0x038bf72a38023435f249f01545d426ebccb3348739aab6415949426e9c5b1ca0.
//
// Solidity: event EditedFile(uint256 indexed fileId, string metadataIPFSHash, string contentIPFSHash, string gateIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) ParseEditedFile(log types.Log) (*BindingsEditedFile, error) {
	event := new(BindingsEditedFile)
	if err := _Bindings.contract.UnpackLog(event, "EditedFile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Bindings contract.
type BindingsOwnershipTransferStartedIterator struct {
	Event *BindingsOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferStarted)
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
		it.Event = new(BindingsOwnershipTransferStarted)
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
func (it *BindingsOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Bindings contract.
type BindingsOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferStartedIterator{contract: _Bindings.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferStarted)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) ParseOwnershipTransferStarted(log types.Log) (*BindingsOwnershipTransferStarted, error) {
	event := new(BindingsOwnershipTransferStarted)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
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
		it.Event = new(BindingsOwnershipTransferred)
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
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRegisteredCollaboratorKeysIterator is returned from FilterRegisteredCollaboratorKeys and is used to iterate over the raw logs and unpacked data for RegisteredCollaboratorKeys events raised by the Bindings contract.
type BindingsRegisteredCollaboratorKeysIterator struct {
	Event *BindingsRegisteredCollaboratorKeys // Event containing the contract specifics and raw log

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
func (it *BindingsRegisteredCollaboratorKeysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRegisteredCollaboratorKeys)
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
		it.Event = new(BindingsRegisteredCollaboratorKeys)
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
func (it *BindingsRegisteredCollaboratorKeysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRegisteredCollaboratorKeysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRegisteredCollaboratorKeys represents a RegisteredCollaboratorKeys event raised by the Bindings contract.
type BindingsRegisteredCollaboratorKeys struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegisteredCollaboratorKeys is a free log retrieval operation binding the contract event 0x402d924b9bd03718220c51647be3ad8f7da00914b8aba403753caa37f4944a02.
//
// Solidity: event RegisteredCollaboratorKeys(address indexed account)
func (_Bindings *BindingsFilterer) FilterRegisteredCollaboratorKeys(opts *bind.FilterOpts, account []common.Address) (*BindingsRegisteredCollaboratorKeysIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RegisteredCollaboratorKeys", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRegisteredCollaboratorKeysIterator{contract: _Bindings.contract, event: "RegisteredCollaboratorKeys", logs: logs, sub: sub}, nil
}

// WatchRegisteredCollaboratorKeys is a free log subscription operation binding the contract event 0x402d924b9bd03718220c51647be3ad8f7da00914b8aba403753caa37f4944a02.
//
// Solidity: event RegisteredCollaboratorKeys(address indexed account)
func (_Bindings *BindingsFilterer) WatchRegisteredCollaboratorKeys(opts *bind.WatchOpts, sink chan<- *BindingsRegisteredCollaboratorKeys, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RegisteredCollaboratorKeys", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRegisteredCollaboratorKeys)
				if err := _Bindings.contract.UnpackLog(event, "RegisteredCollaboratorKeys", log); err != nil {
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

// ParseRegisteredCollaboratorKeys is a log parse operation binding the contract event 0x402d924b9bd03718220c51647be3ad8f7da00914b8aba403753caa37f4944a02.
//
// Solidity: event RegisteredCollaboratorKeys(address indexed account)
func (_Bindings *BindingsFilterer) ParseRegisteredCollaboratorKeys(log types.Log) (*BindingsRegisteredCollaboratorKeys, error) {
	event := new(BindingsRegisteredCollaboratorKeys)
	if err := _Bindings.contract.UnpackLog(event, "RegisteredCollaboratorKeys", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRemovedCollaboratorIterator is returned from FilterRemovedCollaborator and is used to iterate over the raw logs and unpacked data for RemovedCollaborator events raised by the Bindings contract.
type BindingsRemovedCollaboratorIterator struct {
	Event *BindingsRemovedCollaborator // Event containing the contract specifics and raw log

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
func (it *BindingsRemovedCollaboratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRemovedCollaborator)
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
		it.Event = new(BindingsRemovedCollaborator)
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
func (it *BindingsRemovedCollaboratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRemovedCollaboratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRemovedCollaborator represents a RemovedCollaborator event raised by the Bindings contract.
type BindingsRemovedCollaborator struct {
	Account common.Address
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedCollaborator is a free log retrieval operation binding the contract event 0xa3cf1c45ba4d2d0a3d4d7a5fdd2a75b6feec8d0902c676a228fbb2cada34c318.
//
// Solidity: event RemovedCollaborator(address indexed account, address indexed by)
func (_Bindings *BindingsFilterer) FilterRemovedCollaborator(opts *bind.FilterOpts, account []common.Address, by []common.Address) (*BindingsRemovedCollaboratorIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RemovedCollaborator", accountRule, byRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRemovedCollaboratorIterator{contract: _Bindings.contract, event: "RemovedCollaborator", logs: logs, sub: sub}, nil
}

// WatchRemovedCollaborator is a free log subscription operation binding the contract event 0xa3cf1c45ba4d2d0a3d4d7a5fdd2a75b6feec8d0902c676a228fbb2cada34c318.
//
// Solidity: event RemovedCollaborator(address indexed account, address indexed by)
func (_Bindings *BindingsFilterer) WatchRemovedCollaborator(opts *bind.WatchOpts, sink chan<- *BindingsRemovedCollaborator, account []common.Address, by []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RemovedCollaborator", accountRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRemovedCollaborator)
				if err := _Bindings.contract.UnpackLog(event, "RemovedCollaborator", log); err != nil {
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

// ParseRemovedCollaborator is a log parse operation binding the contract event 0xa3cf1c45ba4d2d0a3d4d7a5fdd2a75b6feec8d0902c676a228fbb2cada34c318.
//
// Solidity: event RemovedCollaborator(address indexed account, address indexed by)
func (_Bindings *BindingsFilterer) ParseRemovedCollaborator(log types.Log) (*BindingsRemovedCollaborator, error) {
	event := new(BindingsRemovedCollaborator)
	if err := _Bindings.contract.UnpackLog(event, "RemovedCollaborator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRemovedCollaboratorKeysIterator is returned from FilterRemovedCollaboratorKeys and is used to iterate over the raw logs and unpacked data for RemovedCollaboratorKeys events raised by the Bindings contract.
type BindingsRemovedCollaboratorKeysIterator struct {
	Event *BindingsRemovedCollaboratorKeys // Event containing the contract specifics and raw log

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
func (it *BindingsRemovedCollaboratorKeysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRemovedCollaboratorKeys)
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
		it.Event = new(BindingsRemovedCollaboratorKeys)
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
func (it *BindingsRemovedCollaboratorKeysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRemovedCollaboratorKeysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRemovedCollaboratorKeys represents a RemovedCollaboratorKeys event raised by the Bindings contract.
type BindingsRemovedCollaboratorKeys struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedCollaboratorKeys is a free log retrieval operation binding the contract event 0xe43263e34fd2c5f045a6bb8f67fc65b9bdfd615beb14b6036a76aedf940e1876.
//
// Solidity: event RemovedCollaboratorKeys(address indexed account)
func (_Bindings *BindingsFilterer) FilterRemovedCollaboratorKeys(opts *bind.FilterOpts, account []common.Address) (*BindingsRemovedCollaboratorKeysIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RemovedCollaboratorKeys", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRemovedCollaboratorKeysIterator{contract: _Bindings.contract, event: "RemovedCollaboratorKeys", logs: logs, sub: sub}, nil
}

// WatchRemovedCollaboratorKeys is a free log subscription operation binding the contract event 0xe43263e34fd2c5f045a6bb8f67fc65b9bdfd615beb14b6036a76aedf940e1876.
//
// Solidity: event RemovedCollaboratorKeys(address indexed account)
func (_Bindings *BindingsFilterer) WatchRemovedCollaboratorKeys(opts *bind.WatchOpts, sink chan<- *BindingsRemovedCollaboratorKeys, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RemovedCollaboratorKeys", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRemovedCollaboratorKeys)
				if err := _Bindings.contract.UnpackLog(event, "RemovedCollaboratorKeys", log); err != nil {
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

// ParseRemovedCollaboratorKeys is a log parse operation binding the contract event 0xe43263e34fd2c5f045a6bb8f67fc65b9bdfd615beb14b6036a76aedf940e1876.
//
// Solidity: event RemovedCollaboratorKeys(address indexed account)
func (_Bindings *BindingsFilterer) ParseRemovedCollaboratorKeys(log types.Log) (*BindingsRemovedCollaboratorKeys, error) {
	event := new(BindingsRemovedCollaboratorKeys)
	if err := _Bindings.contract.UnpackLog(event, "RemovedCollaboratorKeys", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdatedKeyVerifiersIterator is returned from FilterUpdatedKeyVerifiers and is used to iterate over the raw logs and unpacked data for UpdatedKeyVerifiers events raised by the Bindings contract.
type BindingsUpdatedKeyVerifiersIterator struct {
	Event *BindingsUpdatedKeyVerifiers // Event containing the contract specifics and raw log

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
func (it *BindingsUpdatedKeyVerifiersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdatedKeyVerifiers)
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
		it.Event = new(BindingsUpdatedKeyVerifiers)
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
func (it *BindingsUpdatedKeyVerifiersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdatedKeyVerifiersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdatedKeyVerifiers represents a UpdatedKeyVerifiers event raised by the Bindings contract.
type BindingsUpdatedKeyVerifiers struct {
	PortalEncryptionKeyVerifier [32]byte
	PortalDecryptionKeyVerifier [32]byte
	MemberEncryptionKeyVerifier [32]byte
	MemberDecryptionKeyVerifier [32]byte
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedKeyVerifiers is a free log retrieval operation binding the contract event 0xaab31fb710d6d5455bc6c4780d47909b9be91efb723613e2f5dd42ae47dd7069.
//
// Solidity: event UpdatedKeyVerifiers(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier)
func (_Bindings *BindingsFilterer) FilterUpdatedKeyVerifiers(opts *bind.FilterOpts) (*BindingsUpdatedKeyVerifiersIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdatedKeyVerifiers")
	if err != nil {
		return nil, err
	}
	return &BindingsUpdatedKeyVerifiersIterator{contract: _Bindings.contract, event: "UpdatedKeyVerifiers", logs: logs, sub: sub}, nil
}

// WatchUpdatedKeyVerifiers is a free log subscription operation binding the contract event 0xaab31fb710d6d5455bc6c4780d47909b9be91efb723613e2f5dd42ae47dd7069.
//
// Solidity: event UpdatedKeyVerifiers(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier)
func (_Bindings *BindingsFilterer) WatchUpdatedKeyVerifiers(opts *bind.WatchOpts, sink chan<- *BindingsUpdatedKeyVerifiers) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdatedKeyVerifiers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdatedKeyVerifiers)
				if err := _Bindings.contract.UnpackLog(event, "UpdatedKeyVerifiers", log); err != nil {
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

// ParseUpdatedKeyVerifiers is a log parse operation binding the contract event 0xaab31fb710d6d5455bc6c4780d47909b9be91efb723613e2f5dd42ae47dd7069.
//
// Solidity: event UpdatedKeyVerifiers(bytes32 portalEncryptionKeyVerifier, bytes32 portalDecryptionKeyVerifier, bytes32 memberEncryptionKeyVerifier, bytes32 memberDecryptionKeyVerifier)
func (_Bindings *BindingsFilterer) ParseUpdatedKeyVerifiers(log types.Log) (*BindingsUpdatedKeyVerifiers, error) {
	event := new(BindingsUpdatedKeyVerifiers)
	if err := _Bindings.contract.UnpackLog(event, "UpdatedKeyVerifiers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdatedPortalMetadataIterator is returned from FilterUpdatedPortalMetadata and is used to iterate over the raw logs and unpacked data for UpdatedPortalMetadata events raised by the Bindings contract.
type BindingsUpdatedPortalMetadataIterator struct {
	Event *BindingsUpdatedPortalMetadata // Event containing the contract specifics and raw log

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
func (it *BindingsUpdatedPortalMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdatedPortalMetadata)
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
		it.Event = new(BindingsUpdatedPortalMetadata)
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
func (it *BindingsUpdatedPortalMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdatedPortalMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdatedPortalMetadata represents a UpdatedPortalMetadata event raised by the Bindings contract.
type BindingsUpdatedPortalMetadata struct {
	MetadataIPFSHash string
	By               common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPortalMetadata is a free log retrieval operation binding the contract event 0xf7e4f835d77656f8760a2c1a3e841fd8a50759e50c615130349625666a46acd5.
//
// Solidity: event UpdatedPortalMetadata(string metadataIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) FilterUpdatedPortalMetadata(opts *bind.FilterOpts, by []common.Address) (*BindingsUpdatedPortalMetadataIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdatedPortalMetadata", byRule)
	if err != nil {
		return nil, err
	}
	return &BindingsUpdatedPortalMetadataIterator{contract: _Bindings.contract, event: "UpdatedPortalMetadata", logs: logs, sub: sub}, nil
}

// WatchUpdatedPortalMetadata is a free log subscription operation binding the contract event 0xf7e4f835d77656f8760a2c1a3e841fd8a50759e50c615130349625666a46acd5.
//
// Solidity: event UpdatedPortalMetadata(string metadataIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) WatchUpdatedPortalMetadata(opts *bind.WatchOpts, sink chan<- *BindingsUpdatedPortalMetadata, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdatedPortalMetadata", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdatedPortalMetadata)
				if err := _Bindings.contract.UnpackLog(event, "UpdatedPortalMetadata", log); err != nil {
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

// ParseUpdatedPortalMetadata is a log parse operation binding the contract event 0xf7e4f835d77656f8760a2c1a3e841fd8a50759e50c615130349625666a46acd5.
//
// Solidity: event UpdatedPortalMetadata(string metadataIPFSHash, address indexed by)
func (_Bindings *BindingsFilterer) ParseUpdatedPortalMetadata(log types.Log) (*BindingsUpdatedPortalMetadata, error) {
	event := new(BindingsUpdatedPortalMetadata)
	if err := _Bindings.contract.UnpackLog(event, "UpdatedPortalMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
