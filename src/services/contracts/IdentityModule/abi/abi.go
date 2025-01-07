// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identitymoduleabi
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

// IIdentityModuleGuardian is an auto generated low-level Go binding around an user-defined struct.
type IIdentityModuleGuardian struct {
	EncryptedShard []byte
	Account        common.Address
	GuardianType   uint8
}

// IIdentityModuleIdentityInput is an auto generated low-level Go binding around an user-defined struct.
type IIdentityModuleIdentityInput struct {
	Salt              *big.Int
	SigningDid        string
	AccountPublicKey  []byte
	AccountPrivateKey []byte
	AgentAddress      common.Address
}

// IIdentityModuleIdentityOutput is an auto generated low-level Go binding around an user-defined struct.
type IIdentityModuleIdentityOutput struct {
	Salt              *big.Int
	SigningDid        string
	AccountPublicKey  []byte
	AccountPrivateKey []byte
	AgentAddress      common.Address
	Guardians         []IIdentityModuleGuardian
}

// IIdentityModuleIdentityPublicDetails is an auto generated low-level Go binding around an user-defined struct.
type IIdentityModuleIdentityPublicDetails struct {
	Salt             *big.Int
	SigningDid       string
	AccountPublicKey []byte
	AgentAddress     common.Address
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"AddIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"AddIdentityGuardians\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"RemoveIdentityGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"RemoveKeyStoreContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"UpdateIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"UpdateIdentityAccountKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpdateIdentityAgent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"UpdateIdentitySaltSigningDid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identityAddress\",\"type\":\"address\"}],\"name\":\"UpdateKeyStoreContract\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"encryptedShard\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumIIdentityModule.GuardianType\",\"name\":\"guardianType\",\"type\":\"uint8\"}],\"internalType\":\"structIIdentityModule.Guardian[]\",\"name\":\"orgGuardians\",\"type\":\"tuple[]\"}],\"name\":\"addIdentityGuardians\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentityModuleDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signingDid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"accountPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountPrivateKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agentAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"encryptedShard\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumIIdentityModule.GuardianType\",\"name\":\"guardianType\",\"type\":\"uint8\"}],\"internalType\":\"structIIdentityModule.Guardian[]\",\"name\":\"guardians\",\"type\":\"tuple[]\"}],\"internalType\":\"structIIdentityModule.IdentityOutput\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getIdentityModuleKeyStoreDetails\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentityModulePublicDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signingDid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"accountPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agentAddress\",\"type\":\"address\"}],\"internalType\":\"structIIdentityModule.IdentityPublicDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"guardianEncryptedShard\",\"type\":\"bytes\"}],\"name\":\"removeIdentityGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contractAddress\",\"type\":\"address[]\"}],\"name\":\"removeKeyStoreContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signingDid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"accountPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountPrivateKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agentAddress\",\"type\":\"address\"}],\"internalType\":\"structIIdentityModule.IdentityInput\",\"name\":\"identity\",\"type\":\"tuple\"}],\"name\":\"updateIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"accountPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountPrivateKey\",\"type\":\"bytes\"}],\"name\":\"updateIdentityAccountKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentAddress\",\"type\":\"address\"}],\"name\":\"updateIdentityAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signingDid\",\"type\":\"string\"}],\"name\":\"updateIdentitySaltSigningDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"ipfsHash\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"contractAddress\",\"type\":\"address[]\"}],\"name\":\"updateKeyStoreContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var IdentitymoduleABI = BindingsMetaData.ABI

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
	Contract     *Bindings      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
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

// GetIdentityModuleDetail is a free data retrieval call binding the contract method 0xdca576ef.
//
// Solidity: function getIdentityModuleDetail() view returns((uint256,string,bytes,bytes,address,(bytes,address,uint8)[]))
func (_Bindings *BindingsCaller) GetIdentityModuleDetail(opts *bind.CallOpts) (IIdentityModuleIdentityOutput, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getIdentityModuleDetail")

	if err != nil {
		return *new(IIdentityModuleIdentityOutput), err
	}

	out0 := *abi.ConvertType(out[0], new(IIdentityModuleIdentityOutput)).(*IIdentityModuleIdentityOutput)

	return out0, err

}

// GetIdentityModuleDetail is a free data retrieval call binding the contract method 0xdca576ef.
//
// Solidity: function getIdentityModuleDetail() view returns((uint256,string,bytes,bytes,address,(bytes,address,uint8)[]))
func (_Bindings *BindingsSession) GetIdentityModuleDetail() (IIdentityModuleIdentityOutput, error) {
	return _Bindings.Contract.GetIdentityModuleDetail(&_Bindings.CallOpts)
}

// GetIdentityModuleDetail is a free data retrieval call binding the contract method 0xdca576ef.
//
// Solidity: function getIdentityModuleDetail() view returns((uint256,string,bytes,bytes,address,(bytes,address,uint8)[]))
func (_Bindings *BindingsCallerSession) GetIdentityModuleDetail() (IIdentityModuleIdentityOutput, error) {
	return _Bindings.Contract.GetIdentityModuleDetail(&_Bindings.CallOpts)
}

// GetIdentityModuleKeyStoreDetails is a free data retrieval call binding the contract method 0xcf2d3488.
//
// Solidity: function getIdentityModuleKeyStoreDetails(address contractAddress) view returns(bytes)
func (_Bindings *BindingsCaller) GetIdentityModuleKeyStoreDetails(opts *bind.CallOpts, contractAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getIdentityModuleKeyStoreDetails", contractAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetIdentityModuleKeyStoreDetails is a free data retrieval call binding the contract method 0xcf2d3488.
//
// Solidity: function getIdentityModuleKeyStoreDetails(address contractAddress) view returns(bytes)
func (_Bindings *BindingsSession) GetIdentityModuleKeyStoreDetails(contractAddress common.Address) ([]byte, error) {
	return _Bindings.Contract.GetIdentityModuleKeyStoreDetails(&_Bindings.CallOpts, contractAddress)
}

// GetIdentityModuleKeyStoreDetails is a free data retrieval call binding the contract method 0xcf2d3488.
//
// Solidity: function getIdentityModuleKeyStoreDetails(address contractAddress) view returns(bytes)
func (_Bindings *BindingsCallerSession) GetIdentityModuleKeyStoreDetails(contractAddress common.Address) ([]byte, error) {
	return _Bindings.Contract.GetIdentityModuleKeyStoreDetails(&_Bindings.CallOpts, contractAddress)
}

// GetIdentityModulePublicDetails is a free data retrieval call binding the contract method 0x897e56fd.
//
// Solidity: function getIdentityModulePublicDetails() view returns((uint256,string,bytes,address))
func (_Bindings *BindingsCaller) GetIdentityModulePublicDetails(opts *bind.CallOpts) (IIdentityModuleIdentityPublicDetails, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getIdentityModulePublicDetails")

	if err != nil {
		return *new(IIdentityModuleIdentityPublicDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IIdentityModuleIdentityPublicDetails)).(*IIdentityModuleIdentityPublicDetails)

	return out0, err

}

// GetIdentityModulePublicDetails is a free data retrieval call binding the contract method 0x897e56fd.
//
// Solidity: function getIdentityModulePublicDetails() view returns((uint256,string,bytes,address))
func (_Bindings *BindingsSession) GetIdentityModulePublicDetails() (IIdentityModuleIdentityPublicDetails, error) {
	return _Bindings.Contract.GetIdentityModulePublicDetails(&_Bindings.CallOpts)
}

// GetIdentityModulePublicDetails is a free data retrieval call binding the contract method 0x897e56fd.
//
// Solidity: function getIdentityModulePublicDetails() view returns((uint256,string,bytes,address))
func (_Bindings *BindingsCallerSession) GetIdentityModulePublicDetails() (IIdentityModuleIdentityPublicDetails, error) {
	return _Bindings.Contract.GetIdentityModulePublicDetails(&_Bindings.CallOpts)
}

// AddIdentityGuardians is a paid mutator transaction binding the contract method 0x4ba4deb1.
//
// Solidity: function addIdentityGuardians((bytes,address,uint8)[] orgGuardians) returns()
func (_Bindings *BindingsTransactor) AddIdentityGuardians(opts *bind.TransactOpts, orgGuardians []IIdentityModuleGuardian) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addIdentityGuardians", orgGuardians)
}

// AddIdentityGuardians is a paid mutator transaction binding the contract method 0x4ba4deb1.
//
// Solidity: function addIdentityGuardians((bytes,address,uint8)[] orgGuardians) returns()
func (_Bindings *BindingsSession) AddIdentityGuardians(orgGuardians []IIdentityModuleGuardian) (*types.Transaction, error) {
	return _Bindings.Contract.AddIdentityGuardians(&_Bindings.TransactOpts, orgGuardians)
}

// AddIdentityGuardians is a paid mutator transaction binding the contract method 0x4ba4deb1.
//
// Solidity: function addIdentityGuardians((bytes,address,uint8)[] orgGuardians) returns()
func (_Bindings *BindingsTransactorSession) AddIdentityGuardians(orgGuardians []IIdentityModuleGuardian) (*types.Transaction, error) {
	return _Bindings.Contract.AddIdentityGuardians(&_Bindings.TransactOpts, orgGuardians)
}

// RemoveIdentityGuardian is a paid mutator transaction binding the contract method 0x109515f1.
//
// Solidity: function removeIdentityGuardian(bytes guardianEncryptedShard) returns()
func (_Bindings *BindingsTransactor) RemoveIdentityGuardian(opts *bind.TransactOpts, guardianEncryptedShard []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeIdentityGuardian", guardianEncryptedShard)
}

// RemoveIdentityGuardian is a paid mutator transaction binding the contract method 0x109515f1.
//
// Solidity: function removeIdentityGuardian(bytes guardianEncryptedShard) returns()
func (_Bindings *BindingsSession) RemoveIdentityGuardian(guardianEncryptedShard []byte) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveIdentityGuardian(&_Bindings.TransactOpts, guardianEncryptedShard)
}

// RemoveIdentityGuardian is a paid mutator transaction binding the contract method 0x109515f1.
//
// Solidity: function removeIdentityGuardian(bytes guardianEncryptedShard) returns()
func (_Bindings *BindingsTransactorSession) RemoveIdentityGuardian(guardianEncryptedShard []byte) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveIdentityGuardian(&_Bindings.TransactOpts, guardianEncryptedShard)
}

// RemoveKeyStoreContract is a paid mutator transaction binding the contract method 0x0a1660ad.
//
// Solidity: function removeKeyStoreContract(address[] contractAddress) returns()
func (_Bindings *BindingsTransactor) RemoveKeyStoreContract(opts *bind.TransactOpts, contractAddress []common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeKeyStoreContract", contractAddress)
}

// RemoveKeyStoreContract is a paid mutator transaction binding the contract method 0x0a1660ad.
//
// Solidity: function removeKeyStoreContract(address[] contractAddress) returns()
func (_Bindings *BindingsSession) RemoveKeyStoreContract(contractAddress []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveKeyStoreContract(&_Bindings.TransactOpts, contractAddress)
}

// RemoveKeyStoreContract is a paid mutator transaction binding the contract method 0x0a1660ad.
//
// Solidity: function removeKeyStoreContract(address[] contractAddress) returns()
func (_Bindings *BindingsTransactorSession) RemoveKeyStoreContract(contractAddress []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveKeyStoreContract(&_Bindings.TransactOpts, contractAddress)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x232b7fc7.
//
// Solidity: function updateIdentity((uint256,string,bytes,bytes,address) identity) returns()
func (_Bindings *BindingsTransactor) UpdateIdentity(opts *bind.TransactOpts, identity IIdentityModuleIdentityInput) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateIdentity", identity)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x232b7fc7.
//
// Solidity: function updateIdentity((uint256,string,bytes,bytes,address) identity) returns()
func (_Bindings *BindingsSession) UpdateIdentity(identity IIdentityModuleIdentityInput) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentity(&_Bindings.TransactOpts, identity)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x232b7fc7.
//
// Solidity: function updateIdentity((uint256,string,bytes,bytes,address) identity) returns()
func (_Bindings *BindingsTransactorSession) UpdateIdentity(identity IIdentityModuleIdentityInput) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentity(&_Bindings.TransactOpts, identity)
}

// UpdateIdentityAccountKey is a paid mutator transaction binding the contract method 0x2d5efa6e.
//
// Solidity: function updateIdentityAccountKey(bytes accountPublicKey, bytes accountPrivateKey) returns()
func (_Bindings *BindingsTransactor) UpdateIdentityAccountKey(opts *bind.TransactOpts, accountPublicKey []byte, accountPrivateKey []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateIdentityAccountKey", accountPublicKey, accountPrivateKey)
}

// UpdateIdentityAccountKey is a paid mutator transaction binding the contract method 0x2d5efa6e.
//
// Solidity: function updateIdentityAccountKey(bytes accountPublicKey, bytes accountPrivateKey) returns()
func (_Bindings *BindingsSession) UpdateIdentityAccountKey(accountPublicKey []byte, accountPrivateKey []byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentityAccountKey(&_Bindings.TransactOpts, accountPublicKey, accountPrivateKey)
}

// UpdateIdentityAccountKey is a paid mutator transaction binding the contract method 0x2d5efa6e.
//
// Solidity: function updateIdentityAccountKey(bytes accountPublicKey, bytes accountPrivateKey) returns()
func (_Bindings *BindingsTransactorSession) UpdateIdentityAccountKey(accountPublicKey []byte, accountPrivateKey []byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentityAccountKey(&_Bindings.TransactOpts, accountPublicKey, accountPrivateKey)
}

// UpdateIdentityAgent is a paid mutator transaction binding the contract method 0x6d0f32a1.
//
// Solidity: function updateIdentityAgent(address agentAddress) returns()
func (_Bindings *BindingsTransactor) UpdateIdentityAgent(opts *bind.TransactOpts, agentAddress common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateIdentityAgent", agentAddress)
}

// UpdateIdentityAgent is a paid mutator transaction binding the contract method 0x6d0f32a1.
//
// Solidity: function updateIdentityAgent(address agentAddress) returns()
func (_Bindings *BindingsSession) UpdateIdentityAgent(agentAddress common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentityAgent(&_Bindings.TransactOpts, agentAddress)
}

// UpdateIdentityAgent is a paid mutator transaction binding the contract method 0x6d0f32a1.
//
// Solidity: function updateIdentityAgent(address agentAddress) returns()
func (_Bindings *BindingsTransactorSession) UpdateIdentityAgent(agentAddress common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentityAgent(&_Bindings.TransactOpts, agentAddress)
}

// UpdateIdentitySaltSigningDid is a paid mutator transaction binding the contract method 0x776be6e2.
//
// Solidity: function updateIdentitySaltSigningDid(uint256 salt, string signingDid) returns()
func (_Bindings *BindingsTransactor) UpdateIdentitySaltSigningDid(opts *bind.TransactOpts, salt *big.Int, signingDid string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateIdentitySaltSigningDid", salt, signingDid)
}

// UpdateIdentitySaltSigningDid is a paid mutator transaction binding the contract method 0x776be6e2.
//
// Solidity: function updateIdentitySaltSigningDid(uint256 salt, string signingDid) returns()
func (_Bindings *BindingsSession) UpdateIdentitySaltSigningDid(salt *big.Int, signingDid string) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentitySaltSigningDid(&_Bindings.TransactOpts, salt, signingDid)
}

// UpdateIdentitySaltSigningDid is a paid mutator transaction binding the contract method 0x776be6e2.
//
// Solidity: function updateIdentitySaltSigningDid(uint256 salt, string signingDid) returns()
func (_Bindings *BindingsTransactorSession) UpdateIdentitySaltSigningDid(salt *big.Int, signingDid string) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateIdentitySaltSigningDid(&_Bindings.TransactOpts, salt, signingDid)
}

// UpdateKeyStoreContract is a paid mutator transaction binding the contract method 0x781a96e8.
//
// Solidity: function updateKeyStoreContract(bytes[] ipfsHash, address[] contractAddress) returns()
func (_Bindings *BindingsTransactor) UpdateKeyStoreContract(opts *bind.TransactOpts, ipfsHash [][]byte, contractAddress []common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateKeyStoreContract", ipfsHash, contractAddress)
}

// UpdateKeyStoreContract is a paid mutator transaction binding the contract method 0x781a96e8.
//
// Solidity: function updateKeyStoreContract(bytes[] ipfsHash, address[] contractAddress) returns()
func (_Bindings *BindingsSession) UpdateKeyStoreContract(ipfsHash [][]byte, contractAddress []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateKeyStoreContract(&_Bindings.TransactOpts, ipfsHash, contractAddress)
}

// UpdateKeyStoreContract is a paid mutator transaction binding the contract method 0x781a96e8.
//
// Solidity: function updateKeyStoreContract(bytes[] ipfsHash, address[] contractAddress) returns()
func (_Bindings *BindingsTransactorSession) UpdateKeyStoreContract(ipfsHash [][]byte, contractAddress []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateKeyStoreContract(&_Bindings.TransactOpts, ipfsHash, contractAddress)
}

// BindingsAddIdentityIterator is returned from FilterAddIdentity and is used to iterate over the raw logs and unpacked data for AddIdentity events raised by the Bindings contract.
type BindingsAddIdentityIterator struct {
	Event *BindingsAddIdentity // Event containing the contract specifics and raw log

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
func (it *BindingsAddIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAddIdentity)
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
		it.Event = new(BindingsAddIdentity)
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
func (it *BindingsAddIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAddIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAddIdentity represents a AddIdentity event raised by the Bindings contract.
type BindingsAddIdentity struct {
	IdentityAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddIdentity is a free log retrieval operation binding the contract event 0x696f554bfa6e8597f6c7e15d57a6d19d4614c602bc21813b1769aebe5d7632e0.
//
// Solidity: event AddIdentity(address identityAddress)
func (_Bindings *BindingsFilterer) FilterAddIdentity(opts *bind.FilterOpts) (*BindingsAddIdentityIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AddIdentity")
	if err != nil {
		return nil, err
	}
	return &BindingsAddIdentityIterator{contract: _Bindings.contract, event: "AddIdentity", logs: logs, sub: sub}, nil
}

// WatchAddIdentity is a free log subscription operation binding the contract event 0x696f554bfa6e8597f6c7e15d57a6d19d4614c602bc21813b1769aebe5d7632e0.
//
// Solidity: event AddIdentity(address identityAddress)
func (_Bindings *BindingsFilterer) WatchAddIdentity(opts *bind.WatchOpts, sink chan<- *BindingsAddIdentity) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AddIdentity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAddIdentity)
				if err := _Bindings.contract.UnpackLog(event, "AddIdentity", log); err != nil {
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

// ParseAddIdentity is a log parse operation binding the contract event 0x696f554bfa6e8597f6c7e15d57a6d19d4614c602bc21813b1769aebe5d7632e0.
//
// Solidity: event AddIdentity(address identityAddress)
func (_Bindings *BindingsFilterer) ParseAddIdentity(log types.Log) (*BindingsAddIdentity, error) {
	event := new(BindingsAddIdentity)
	if err := _Bindings.contract.UnpackLog(event, "AddIdentity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsAddIdentityGuardiansIterator is returned from FilterAddIdentityGuardians and is used to iterate over the raw logs and unpacked data for AddIdentityGuardians events raised by the Bindings contract.
type BindingsAddIdentityGuardiansIterator struct {
	Event *BindingsAddIdentityGuardians // Event containing the contract specifics and raw log

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
func (it *BindingsAddIdentityGuardiansIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAddIdentityGuardians)
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
		it.Event = new(BindingsAddIdentityGuardians)
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
func (it *BindingsAddIdentityGuardiansIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAddIdentityGuardiansIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAddIdentityGuardians represents a AddIdentityGuardians event raised by the Bindings contract.
type BindingsAddIdentityGuardians struct {
	Identity common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddIdentityGuardians is a free log retrieval operation binding the contract event 0x0f8e407b19d13f05c7b7ec64a02fe05416da4f19dc49c787ad7affe8601e8a37.
//
// Solidity: event AddIdentityGuardians(address identity)
func (_Bindings *BindingsFilterer) FilterAddIdentityGuardians(opts *bind.FilterOpts) (*BindingsAddIdentityGuardiansIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AddIdentityGuardians")
	if err != nil {
		return nil, err
	}
	return &BindingsAddIdentityGuardiansIterator{contract: _Bindings.contract, event: "AddIdentityGuardians", logs: logs, sub: sub}, nil
}

// WatchAddIdentityGuardians is a free log subscription operation binding the contract event 0x0f8e407b19d13f05c7b7ec64a02fe05416da4f19dc49c787ad7affe8601e8a37.
//
// Solidity: event AddIdentityGuardians(address identity)
func (_Bindings *BindingsFilterer) WatchAddIdentityGuardians(opts *bind.WatchOpts, sink chan<- *BindingsAddIdentityGuardians) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AddIdentityGuardians")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAddIdentityGuardians)
				if err := _Bindings.contract.UnpackLog(event, "AddIdentityGuardians", log); err != nil {
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

// ParseAddIdentityGuardians is a log parse operation binding the contract event 0x0f8e407b19d13f05c7b7ec64a02fe05416da4f19dc49c787ad7affe8601e8a37.
//
// Solidity: event AddIdentityGuardians(address identity)
func (_Bindings *BindingsFilterer) ParseAddIdentityGuardians(log types.Log) (*BindingsAddIdentityGuardians, error) {
	event := new(BindingsAddIdentityGuardians)
	if err := _Bindings.contract.UnpackLog(event, "AddIdentityGuardians", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the Bindings contract.
type BindingsExecutionFailureIterator struct {
	Event *BindingsExecutionFailure // Event containing the contract specifics and raw log

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
func (it *BindingsExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExecutionFailure)
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
		it.Event = new(BindingsExecutionFailure)
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
func (it *BindingsExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExecutionFailure represents a ExecutionFailure event raised by the Bindings contract.
type BindingsExecutionFailure struct {
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 payment)
func (_Bindings *BindingsFilterer) FilterExecutionFailure(opts *bind.FilterOpts) (*BindingsExecutionFailureIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return &BindingsExecutionFailureIterator{contract: _Bindings.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 payment)
func (_Bindings *BindingsFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *BindingsExecutionFailure) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExecutionFailure)
				if err := _Bindings.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 payment)
func (_Bindings *BindingsFilterer) ParseExecutionFailure(log types.Log) (*BindingsExecutionFailure, error) {
	event := new(BindingsExecutionFailure)
	if err := _Bindings.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the Bindings contract.
type BindingsExecutionSuccessIterator struct {
	Event *BindingsExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *BindingsExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExecutionSuccess)
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
		it.Event = new(BindingsExecutionSuccess)
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
func (it *BindingsExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExecutionSuccess represents a ExecutionSuccess event raised by the Bindings contract.
type BindingsExecutionSuccess struct {
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0xf2788e0d1a8dffbdac057cef8914e5346afd7a09d39d19b60080caf3cc07c4f7.
//
// Solidity: event ExecutionSuccess(uint256 payment)
func (_Bindings *BindingsFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*BindingsExecutionSuccessIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &BindingsExecutionSuccessIterator{contract: _Bindings.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0xf2788e0d1a8dffbdac057cef8914e5346afd7a09d39d19b60080caf3cc07c4f7.
//
// Solidity: event ExecutionSuccess(uint256 payment)
func (_Bindings *BindingsFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *BindingsExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExecutionSuccess)
				if err := _Bindings.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0xf2788e0d1a8dffbdac057cef8914e5346afd7a09d39d19b60080caf3cc07c4f7.
//
// Solidity: event ExecutionSuccess(uint256 payment)
func (_Bindings *BindingsFilterer) ParseExecutionSuccess(log types.Log) (*BindingsExecutionSuccess, error) {
	event := new(BindingsExecutionSuccess)
	if err := _Bindings.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRemoveIdentityGuardianIterator is returned from FilterRemoveIdentityGuardian and is used to iterate over the raw logs and unpacked data for RemoveIdentityGuardian events raised by the Bindings contract.
type BindingsRemoveIdentityGuardianIterator struct {
	Event *BindingsRemoveIdentityGuardian // Event containing the contract specifics and raw log

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
func (it *BindingsRemoveIdentityGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRemoveIdentityGuardian)
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
		it.Event = new(BindingsRemoveIdentityGuardian)
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
func (it *BindingsRemoveIdentityGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRemoveIdentityGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRemoveIdentityGuardian represents a RemoveIdentityGuardian event raised by the Bindings contract.
type BindingsRemoveIdentityGuardian struct {
	Identity common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveIdentityGuardian is a free log retrieval operation binding the contract event 0x7823aea0ddd0f0be22780761d32d5f07778185275bc9dd35a52e44eefb1418a2.
//
// Solidity: event RemoveIdentityGuardian(address identity)
func (_Bindings *BindingsFilterer) FilterRemoveIdentityGuardian(opts *bind.FilterOpts) (*BindingsRemoveIdentityGuardianIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RemoveIdentityGuardian")
	if err != nil {
		return nil, err
	}
	return &BindingsRemoveIdentityGuardianIterator{contract: _Bindings.contract, event: "RemoveIdentityGuardian", logs: logs, sub: sub}, nil
}

// WatchRemoveIdentityGuardian is a free log subscription operation binding the contract event 0x7823aea0ddd0f0be22780761d32d5f07778185275bc9dd35a52e44eefb1418a2.
//
// Solidity: event RemoveIdentityGuardian(address identity)
func (_Bindings *BindingsFilterer) WatchRemoveIdentityGuardian(opts *bind.WatchOpts, sink chan<- *BindingsRemoveIdentityGuardian) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RemoveIdentityGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRemoveIdentityGuardian)
				if err := _Bindings.contract.UnpackLog(event, "RemoveIdentityGuardian", log); err != nil {
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

// ParseRemoveIdentityGuardian is a log parse operation binding the contract event 0x7823aea0ddd0f0be22780761d32d5f07778185275bc9dd35a52e44eefb1418a2.
//
// Solidity: event RemoveIdentityGuardian(address identity)
func (_Bindings *BindingsFilterer) ParseRemoveIdentityGuardian(log types.Log) (*BindingsRemoveIdentityGuardian, error) {
	event := new(BindingsRemoveIdentityGuardian)
	if err := _Bindings.contract.UnpackLog(event, "RemoveIdentityGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRemoveKeyStoreContractIterator is returned from FilterRemoveKeyStoreContract and is used to iterate over the raw logs and unpacked data for RemoveKeyStoreContract events raised by the Bindings contract.
type BindingsRemoveKeyStoreContractIterator struct {
	Event *BindingsRemoveKeyStoreContract // Event containing the contract specifics and raw log

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
func (it *BindingsRemoveKeyStoreContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRemoveKeyStoreContract)
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
		it.Event = new(BindingsRemoveKeyStoreContract)
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
func (it *BindingsRemoveKeyStoreContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRemoveKeyStoreContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRemoveKeyStoreContract represents a RemoveKeyStoreContract event raised by the Bindings contract.
type BindingsRemoveKeyStoreContract struct {
	IdentityAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemoveKeyStoreContract is a free log retrieval operation binding the contract event 0x11710649f7aab8640a7351f1f984b44709b1a727e7ab5ab0de28087e78abfacd.
//
// Solidity: event RemoveKeyStoreContract(address identityAddress)
func (_Bindings *BindingsFilterer) FilterRemoveKeyStoreContract(opts *bind.FilterOpts) (*BindingsRemoveKeyStoreContractIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RemoveKeyStoreContract")
	if err != nil {
		return nil, err
	}
	return &BindingsRemoveKeyStoreContractIterator{contract: _Bindings.contract, event: "RemoveKeyStoreContract", logs: logs, sub: sub}, nil
}

// WatchRemoveKeyStoreContract is a free log subscription operation binding the contract event 0x11710649f7aab8640a7351f1f984b44709b1a727e7ab5ab0de28087e78abfacd.
//
// Solidity: event RemoveKeyStoreContract(address identityAddress)
func (_Bindings *BindingsFilterer) WatchRemoveKeyStoreContract(opts *bind.WatchOpts, sink chan<- *BindingsRemoveKeyStoreContract) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RemoveKeyStoreContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRemoveKeyStoreContract)
				if err := _Bindings.contract.UnpackLog(event, "RemoveKeyStoreContract", log); err != nil {
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

// ParseRemoveKeyStoreContract is a log parse operation binding the contract event 0x11710649f7aab8640a7351f1f984b44709b1a727e7ab5ab0de28087e78abfacd.
//
// Solidity: event RemoveKeyStoreContract(address identityAddress)
func (_Bindings *BindingsFilterer) ParseRemoveKeyStoreContract(log types.Log) (*BindingsRemoveKeyStoreContract, error) {
	event := new(BindingsRemoveKeyStoreContract)
	if err := _Bindings.contract.UnpackLog(event, "RemoveKeyStoreContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdateIdentityIterator is returned from FilterUpdateIdentity and is used to iterate over the raw logs and unpacked data for UpdateIdentity events raised by the Bindings contract.
type BindingsUpdateIdentityIterator struct {
	Event *BindingsUpdateIdentity // Event containing the contract specifics and raw log

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
func (it *BindingsUpdateIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdateIdentity)
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
		it.Event = new(BindingsUpdateIdentity)
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
func (it *BindingsUpdateIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdateIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdateIdentity represents a UpdateIdentity event raised by the Bindings contract.
type BindingsUpdateIdentity struct {
	IdentityAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateIdentity is a free log retrieval operation binding the contract event 0x16d3d34a0f2f74e5f733b88045f080217aaf7c95e8eaf0b8a93e7473e158be26.
//
// Solidity: event UpdateIdentity(address identityAddress)
func (_Bindings *BindingsFilterer) FilterUpdateIdentity(opts *bind.FilterOpts) (*BindingsUpdateIdentityIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdateIdentity")
	if err != nil {
		return nil, err
	}
	return &BindingsUpdateIdentityIterator{contract: _Bindings.contract, event: "UpdateIdentity", logs: logs, sub: sub}, nil
}

// WatchUpdateIdentity is a free log subscription operation binding the contract event 0x16d3d34a0f2f74e5f733b88045f080217aaf7c95e8eaf0b8a93e7473e158be26.
//
// Solidity: event UpdateIdentity(address identityAddress)
func (_Bindings *BindingsFilterer) WatchUpdateIdentity(opts *bind.WatchOpts, sink chan<- *BindingsUpdateIdentity) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdateIdentity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdateIdentity)
				if err := _Bindings.contract.UnpackLog(event, "UpdateIdentity", log); err != nil {
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

// ParseUpdateIdentity is a log parse operation binding the contract event 0x16d3d34a0f2f74e5f733b88045f080217aaf7c95e8eaf0b8a93e7473e158be26.
//
// Solidity: event UpdateIdentity(address identityAddress)
func (_Bindings *BindingsFilterer) ParseUpdateIdentity(log types.Log) (*BindingsUpdateIdentity, error) {
	event := new(BindingsUpdateIdentity)
	if err := _Bindings.contract.UnpackLog(event, "UpdateIdentity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdateIdentityAccountKeyIterator is returned from FilterUpdateIdentityAccountKey and is used to iterate over the raw logs and unpacked data for UpdateIdentityAccountKey events raised by the Bindings contract.
type BindingsUpdateIdentityAccountKeyIterator struct {
	Event *BindingsUpdateIdentityAccountKey // Event containing the contract specifics and raw log

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
func (it *BindingsUpdateIdentityAccountKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdateIdentityAccountKey)
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
		it.Event = new(BindingsUpdateIdentityAccountKey)
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
func (it *BindingsUpdateIdentityAccountKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdateIdentityAccountKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdateIdentityAccountKey represents a UpdateIdentityAccountKey event raised by the Bindings contract.
type BindingsUpdateIdentityAccountKey struct {
	IdentityAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateIdentityAccountKey is a free log retrieval operation binding the contract event 0x787e962c49d675c8e739d0d0cc9f9fe3626edf8a067f53014f9eb602ae702bdb.
//
// Solidity: event UpdateIdentityAccountKey(address identityAddress)
func (_Bindings *BindingsFilterer) FilterUpdateIdentityAccountKey(opts *bind.FilterOpts) (*BindingsUpdateIdentityAccountKeyIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdateIdentityAccountKey")
	if err != nil {
		return nil, err
	}
	return &BindingsUpdateIdentityAccountKeyIterator{contract: _Bindings.contract, event: "UpdateIdentityAccountKey", logs: logs, sub: sub}, nil
}

// WatchUpdateIdentityAccountKey is a free log subscription operation binding the contract event 0x787e962c49d675c8e739d0d0cc9f9fe3626edf8a067f53014f9eb602ae702bdb.
//
// Solidity: event UpdateIdentityAccountKey(address identityAddress)
func (_Bindings *BindingsFilterer) WatchUpdateIdentityAccountKey(opts *bind.WatchOpts, sink chan<- *BindingsUpdateIdentityAccountKey) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdateIdentityAccountKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdateIdentityAccountKey)
				if err := _Bindings.contract.UnpackLog(event, "UpdateIdentityAccountKey", log); err != nil {
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

// ParseUpdateIdentityAccountKey is a log parse operation binding the contract event 0x787e962c49d675c8e739d0d0cc9f9fe3626edf8a067f53014f9eb602ae702bdb.
//
// Solidity: event UpdateIdentityAccountKey(address identityAddress)
func (_Bindings *BindingsFilterer) ParseUpdateIdentityAccountKey(log types.Log) (*BindingsUpdateIdentityAccountKey, error) {
	event := new(BindingsUpdateIdentityAccountKey)
	if err := _Bindings.contract.UnpackLog(event, "UpdateIdentityAccountKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdateIdentityAgentIterator is returned from FilterUpdateIdentityAgent and is used to iterate over the raw logs and unpacked data for UpdateIdentityAgent events raised by the Bindings contract.
type BindingsUpdateIdentityAgentIterator struct {
	Event *BindingsUpdateIdentityAgent // Event containing the contract specifics and raw log

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
func (it *BindingsUpdateIdentityAgentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdateIdentityAgent)
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
		it.Event = new(BindingsUpdateIdentityAgent)
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
func (it *BindingsUpdateIdentityAgentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdateIdentityAgentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdateIdentityAgent represents a UpdateIdentityAgent event raised by the Bindings contract.
type BindingsUpdateIdentityAgent struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateIdentityAgent is a free log retrieval operation binding the contract event 0xb9f280f703f99edd5290fbb726ff5cb0eeb1d7db6f47ab187df55930cdc51e55.
//
// Solidity: event UpdateIdentityAgent(address agent)
func (_Bindings *BindingsFilterer) FilterUpdateIdentityAgent(opts *bind.FilterOpts) (*BindingsUpdateIdentityAgentIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdateIdentityAgent")
	if err != nil {
		return nil, err
	}
	return &BindingsUpdateIdentityAgentIterator{contract: _Bindings.contract, event: "UpdateIdentityAgent", logs: logs, sub: sub}, nil
}

// WatchUpdateIdentityAgent is a free log subscription operation binding the contract event 0xb9f280f703f99edd5290fbb726ff5cb0eeb1d7db6f47ab187df55930cdc51e55.
//
// Solidity: event UpdateIdentityAgent(address agent)
func (_Bindings *BindingsFilterer) WatchUpdateIdentityAgent(opts *bind.WatchOpts, sink chan<- *BindingsUpdateIdentityAgent) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdateIdentityAgent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdateIdentityAgent)
				if err := _Bindings.contract.UnpackLog(event, "UpdateIdentityAgent", log); err != nil {
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

// ParseUpdateIdentityAgent is a log parse operation binding the contract event 0xb9f280f703f99edd5290fbb726ff5cb0eeb1d7db6f47ab187df55930cdc51e55.
//
// Solidity: event UpdateIdentityAgent(address agent)
func (_Bindings *BindingsFilterer) ParseUpdateIdentityAgent(log types.Log) (*BindingsUpdateIdentityAgent, error) {
	event := new(BindingsUpdateIdentityAgent)
	if err := _Bindings.contract.UnpackLog(event, "UpdateIdentityAgent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdateIdentitySaltSigningDidIterator is returned from FilterUpdateIdentitySaltSigningDid and is used to iterate over the raw logs and unpacked data for UpdateIdentitySaltSigningDid events raised by the Bindings contract.
type BindingsUpdateIdentitySaltSigningDidIterator struct {
	Event *BindingsUpdateIdentitySaltSigningDid // Event containing the contract specifics and raw log

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
func (it *BindingsUpdateIdentitySaltSigningDidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdateIdentitySaltSigningDid)
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
		it.Event = new(BindingsUpdateIdentitySaltSigningDid)
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
func (it *BindingsUpdateIdentitySaltSigningDidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdateIdentitySaltSigningDidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdateIdentitySaltSigningDid represents a UpdateIdentitySaltSigningDid event raised by the Bindings contract.
type BindingsUpdateIdentitySaltSigningDid struct {
	IdentityAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateIdentitySaltSigningDid is a free log retrieval operation binding the contract event 0x20f6814e8271ebdb56562b6bcae01eb60332fa7ca8a4cbd0385e77a4a9c11fc0.
//
// Solidity: event UpdateIdentitySaltSigningDid(address identityAddress)
func (_Bindings *BindingsFilterer) FilterUpdateIdentitySaltSigningDid(opts *bind.FilterOpts) (*BindingsUpdateIdentitySaltSigningDidIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdateIdentitySaltSigningDid")
	if err != nil {
		return nil, err
	}
	return &BindingsUpdateIdentitySaltSigningDidIterator{contract: _Bindings.contract, event: "UpdateIdentitySaltSigningDid", logs: logs, sub: sub}, nil
}

// WatchUpdateIdentitySaltSigningDid is a free log subscription operation binding the contract event 0x20f6814e8271ebdb56562b6bcae01eb60332fa7ca8a4cbd0385e77a4a9c11fc0.
//
// Solidity: event UpdateIdentitySaltSigningDid(address identityAddress)
func (_Bindings *BindingsFilterer) WatchUpdateIdentitySaltSigningDid(opts *bind.WatchOpts, sink chan<- *BindingsUpdateIdentitySaltSigningDid) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdateIdentitySaltSigningDid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdateIdentitySaltSigningDid)
				if err := _Bindings.contract.UnpackLog(event, "UpdateIdentitySaltSigningDid", log); err != nil {
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

// ParseUpdateIdentitySaltSigningDid is a log parse operation binding the contract event 0x20f6814e8271ebdb56562b6bcae01eb60332fa7ca8a4cbd0385e77a4a9c11fc0.
//
// Solidity: event UpdateIdentitySaltSigningDid(address identityAddress)
func (_Bindings *BindingsFilterer) ParseUpdateIdentitySaltSigningDid(log types.Log) (*BindingsUpdateIdentitySaltSigningDid, error) {
	event := new(BindingsUpdateIdentitySaltSigningDid)
	if err := _Bindings.contract.UnpackLog(event, "UpdateIdentitySaltSigningDid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpdateKeyStoreContractIterator is returned from FilterUpdateKeyStoreContract and is used to iterate over the raw logs and unpacked data for UpdateKeyStoreContract events raised by the Bindings contract.
type BindingsUpdateKeyStoreContractIterator struct {
	Event *BindingsUpdateKeyStoreContract // Event containing the contract specifics and raw log

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
func (it *BindingsUpdateKeyStoreContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpdateKeyStoreContract)
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
		it.Event = new(BindingsUpdateKeyStoreContract)
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
func (it *BindingsUpdateKeyStoreContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpdateKeyStoreContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpdateKeyStoreContract represents a UpdateKeyStoreContract event raised by the Bindings contract.
type BindingsUpdateKeyStoreContract struct {
	IdentityAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateKeyStoreContract is a free log retrieval operation binding the contract event 0x37fe678ae18145ce41727287e39fd0eaa2a3db98c87e169505fa334b5cc8f8c5.
//
// Solidity: event UpdateKeyStoreContract(address identityAddress)
func (_Bindings *BindingsFilterer) FilterUpdateKeyStoreContract(opts *bind.FilterOpts) (*BindingsUpdateKeyStoreContractIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "UpdateKeyStoreContract")
	if err != nil {
		return nil, err
	}
	return &BindingsUpdateKeyStoreContractIterator{contract: _Bindings.contract, event: "UpdateKeyStoreContract", logs: logs, sub: sub}, nil
}

// WatchUpdateKeyStoreContract is a free log subscription operation binding the contract event 0x37fe678ae18145ce41727287e39fd0eaa2a3db98c87e169505fa334b5cc8f8c5.
//
// Solidity: event UpdateKeyStoreContract(address identityAddress)
func (_Bindings *BindingsFilterer) WatchUpdateKeyStoreContract(opts *bind.WatchOpts, sink chan<- *BindingsUpdateKeyStoreContract) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "UpdateKeyStoreContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpdateKeyStoreContract)
				if err := _Bindings.contract.UnpackLog(event, "UpdateKeyStoreContract", log); err != nil {
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

// ParseUpdateKeyStoreContract is a log parse operation binding the contract event 0x37fe678ae18145ce41727287e39fd0eaa2a3db98c87e169505fa334b5cc8f8c5.
//
// Solidity: event UpdateKeyStoreContract(address identityAddress)
func (_Bindings *BindingsFilterer) ParseUpdateKeyStoreContract(log types.Log) (*BindingsUpdateKeyStoreContract, error) {
	event := new(BindingsUpdateKeyStoreContract)
	if err := _Bindings.contract.UnpackLog(event, "UpdateKeyStoreContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
