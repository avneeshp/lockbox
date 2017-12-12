// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LockboxAppABI is the input ABI used to generate the binding from.
const LockboxAppABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockboxAccounts\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ssn\",\"type\":\"string\"},{\"name\":\"license\",\"type\":\"string\"},{\"name\":\"dob\",\"type\":\"uint256\"},{\"name\":\"initialized\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockboxAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"lname\",\"type\":\"string\"}],\"name\":\"createLockbox\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getLicenseNo\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ssn\",\"type\":\"string\"},{\"name\":\"licenseno\",\"type\":\"string\"},{\"name\":\"dob\",\"type\":\"uint256\"}],\"name\":\"addPersonalInfo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getDob\",\"outputs\":[{\"name\":\"num\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"LockBoxApp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cname\",\"type\":\"string\"},{\"name\":\"cardno\",\"type\":\"uint64\"},{\"name\":\"edate\",\"type\":\"uint256\"},{\"name\":\"cvv\",\"type\":\"uint256\"}],\"name\":\"addCardDetails\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getCardDetails\",\"outputs\":[{\"name\":\"cardno\",\"type\":\"uint64\"},{\"name\":\"edate\",\"type\":\"uint256\"},{\"name\":\"cvv\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getLoginCredentials\",\"outputs\":[{\"name\":\"service\",\"type\":\"string\"},{\"name\":\"uname\",\"type\":\"string\"},{\"name\":\"passwd\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"service\",\"type\":\"string\"},{\"name\":\"uname\",\"type\":\"string\"},{\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"addLoginCredentials\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getLockboxName\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"string\"}],\"name\":\"stringToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSSN\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"initialized\",\"type\":\"bool\"}],\"name\":\"LockboxAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ssn\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"license\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"dob\",\"type\":\"uint256\"}],\"name\":\"PersonalInfoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cname\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"cardno\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"edate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cvv\",\"type\":\"uint256\"}],\"name\":\"CardDetailsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"service\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"uname\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"LoginCredentialsAdded\",\"type\":\"event\"}]"

// LockboxAppBin is the compiled bytecode used for deploying new contracts.
const LockboxAppBin = `0x6060604052341561000f57600080fd5b6116458061001e6000396000f3006060604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663096ce44781146100d4578063106061e51461029257806323bfa8e5146102c457806327ebbc2f1461031e57806336ff4a1b146103b45780635d936ce01461043e5780637bf699dd1461046f57806384ab686d14610484578063870ab12d146104e1578063ab23cceb14610578578063b6cd185614610721578063b87d8ce6146107eb578063cfb519281461080a578063fffcd07c1461085b575b600080fd5b34156100df57600080fd5b6100f3600160a060020a036004351661087a565b604051600160a060020a03871681526080810183905281151560a082015260c060208201818152875460026000196101006001841615020190911604918301829052906040830190606084019060e08501908a9080156101945780601f1061016957610100808354040283529160200191610194565b820191906000526020600020905b81548152906001019060200180831161017757829003601f168201915b50508481038352885460026000196101006001841615020190911604808252602090910190899080156102085780601f106101dd57610100808354040283529160200191610208565b820191906000526020600020905b8154815290600101906020018083116101eb57829003601f168201915b505084810382528754600260001961010060018416150201909116048082526020909101908890801561027c5780601f106102515761010080835404028352916020019161027c565b820191906000526020600020905b81548152906001019060200180831161025f57829003601f168201915b5050995050505050505050505060405180910390f35b341561029d57600080fd5b6102a86004356108b8565b604051600160a060020a03909116815260200160405180910390f35b61030a60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506108e095505050505050565b604051901515815260200160405180910390f35b341561032957600080fd5b61033d600160a060020a0360043516610a9b565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610379578082015183820152602001610361565b50505050905090810190601f1680156103a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61030a60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350610b7492505050565b341561044957600080fd5b61045d600160a060020a0360043516610cfa565b60405190815260200160405180910390f35b341561047a57600080fd5b610482610d2f565b005b61030a60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505067ffffffffffffffff8535169460208101359450604001359250610d59915050565b34156104ec57600080fd5b61054060048035600160a060020a03169060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610ed095505050505050565b604051808467ffffffffffffffff1667ffffffffffffffff168152602001838152602001828152602001935050505060405180910390f35b341561058357600080fd5b6105d760048035600160a060020a03169060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610f4895505050505050565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561061c578082015183820152602001610604565b50505050905090810190601f1680156106495780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561067f578082015183820152602001610667565b50505050905090810190601f1680156106ac5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156106e25780820151838201526020016106ca565b50505050905090810190601f16801561070f5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61030a60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506111aa95505050505050565b34156107f657600080fd5b61033d600160a060020a03600435166113da565b341561081557600080fd5b61045d60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061147a95505050505050565b341561086657600080fd5b61033d600160a060020a03600435166114a3565b6001602081905260009182526040909120805460048201546007830154600160a060020a039092169383019260028101926003909101919060ff1686565b60028054829081106108c657fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a03331660009081526001602052604081206007015460ff161515610a9257600160a060020a0333166000818152600160208190526040909120805473ffffffffffffffffffffffffffffffffffffffff1916909217825501828051610950929160200190611543565b50600160a060020a0333166000908152600160208190526040909120600701805460ff191682179055600280549091810161098b83826115c1565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03811691821790925583526001909152604091829020600701547fc676fbb134506bee5d4a7d93f37447a0681c2649d76fa3853eeebee5ea21ac1a92859160ff169051600160a060020a0384168152811515604082015260606020820181815290820184818151815260200191508051906020019080838360005b83811015610a4e578082015183820152602001610a36565b50505050905090810190601f168015610a7b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1506001610a96565b5060005b919050565b610aa36115ea565b600160a060020a0382166000908152600160205260409020600781015460ff161515610ace57600080fd5b806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b665780601f10610b3b57610100808354040283529160200191610b66565b820191906000526020600020905b815481529060010190602001808311610b4957829003601f168201915b505050505091505b50919050565b600160a060020a0333166000908152600160205260408120600781015460ff161515610b9f57600080fd5b60028101858051610bb4929160200190611543565b5060038101848051610bca929160200190611543565b50600481018390557fbdf21c90d1a57a57ad21401208d3b694f2496e822588a6a4790a3f4fc2b5bcd933868686604051600160a060020a038516815260608101829052608060208201818152906040830190830186818151815260200191508051906020019080838360005b83811015610c4e578082015183820152602001610c36565b50505050905090810190601f168015610c7b5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015610cb1578082015183820152602001610c99565b50505050905090810190601f168015610cde5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1506001949350505050565b600160a060020a0381166000908152600160205260408120600781015460ff161515610d2557600080fd5b6004015492915050565b6000805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a0316179055565b600160a060020a03331660009081526001602052604081206007810154829060ff161515610d8657600080fd5b610d8f8761147a565b60008181526005840160205260409020909150878051610db3929160200190611543565b5060008181526005830160205260409081902060018101805467ffffffffffffffff191667ffffffffffffffff8a16179055600281018790556003018590557fe3839c6e9286edc7695b5603b2119a05b0e7ef56c8b96d9a6e78ee2878e1da7d903390899089908990899051600160a060020a038616815267ffffffffffffffff84166040820152606081018390526080810182905260a06020820181815290820186818151815260200191508051906020019080838360005b83811015610e85578082015183820152602001610e6d565b50505050905090810190601f168015610eb25780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15060019695505050505050565b600160a060020a03821660009081526001602052604081206007810154829182918290819060ff161515610f0357600080fd5b610f0c8761147a565b60009081526005939093016020525050604090206001810154600282015460039092015467ffffffffffffffff90911697919650945092505050565b610f506115ea565b610f586115ea565b610f606115ea565b600160a060020a03851660009081526001602052604081206007810154909190819060ff161515610f9057600080fd5b610f998761147a565b915082600601600083600019166000191681526020019081526020016000209050806000018160010182600201828054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561105b5780601f106110305761010080835404028352916020019161105b565b820191906000526020600020905b81548152906001019060200180831161103e57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110f75780601f106110cc576101008083540402835291602001916110f7565b820191906000526020600020905b8154815290600101906020018083116110da57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111935780601f1061116857610100808354040283529160200191611193565b820191906000526020600020905b81548152906001019060200180831161117657829003601f168201915b505050505090509550955095505050509250925092565b600160a060020a03331660009081526001602052604081206007810154829060ff1615156111d757600080fd5b6111e08661147a565b60008181526006840160205260409020909150868051611204929160200190611543565b5060008181526006830160205260409020600101858051611229929160200190611543565b506000818152600683016020526040902060020184805161124e929160200190611543565b507f1376bb2e6a30a62bcaee6a469d2c715a9a209c296c1dab21424ec66804f2370233878787604051600160a060020a03851681526080602082018181529060408301906060840190840187818151815260200191508051906020019080838360005b838110156112c95780820151838201526020016112b1565b50505050905090810190601f1680156112f65780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561132c578082015183820152602001611314565b50505050905090810190601f1680156113595780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561138f578082015183820152602001611377565b50505050905090810190601f1680156113bc5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600195945050505050565b6113e26115ea565b600160a060020a0382166000908152600160205260409020600781015460ff16151561140d57600080fd5b806001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b665780601f10610b3b57610100808354040283529160200191610b66565b60006114846115ea565b5081805115156114975760009150610b6e565b60208301519392505050565b6114ab6115ea565b600160a060020a0382166000908152600160205260409020600781015460ff1615156114d657600080fd5b806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b665780601f10610b3b57610100808354040283529160200191610b66565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061158457805160ff19168380011785556115b1565b828001600101855582156115b1579182015b828111156115b1578251825591602001919060010190611596565b506115bd9291506115fc565b5090565b8154818355818115116115e5576000838152602090206115e59181019083016115fc565b505050565b60206040519081016040526000815290565b61161691905b808211156115bd5760008155600101611602565b905600a165627a7a723058209f96816e48dbefad8ce55d84a4c4faee3f8cbe483ca749f359933f0c0a5d5a640029`

// DeployLockboxApp deploys a new Ethereum contract, binding an instance of LockboxApp to it.
func DeployLockboxApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LockboxApp, error) {
	parsed, err := abi.JSON(strings.NewReader(LockboxAppABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockboxAppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockboxApp{LockboxAppCaller: LockboxAppCaller{contract: contract}, LockboxAppTransactor: LockboxAppTransactor{contract: contract}}, nil
}

// LockboxApp is an auto generated Go binding around an Ethereum contract.
type LockboxApp struct {
	LockboxAppCaller     // Read-only binding to the contract
	LockboxAppTransactor // Write-only binding to the contract
}

// LockboxAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockboxAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockboxAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockboxAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockboxAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockboxAppSession struct {
	Contract     *LockboxApp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockboxAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockboxAppCallerSession struct {
	Contract *LockboxAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LockboxAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockboxAppTransactorSession struct {
	Contract     *LockboxAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LockboxAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockboxAppRaw struct {
	Contract *LockboxApp // Generic contract binding to access the raw methods on
}

// LockboxAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockboxAppCallerRaw struct {
	Contract *LockboxAppCaller // Generic read-only contract binding to access the raw methods on
}

// LockboxAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockboxAppTransactorRaw struct {
	Contract *LockboxAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockboxApp creates a new instance of LockboxApp, bound to a specific deployed contract.
func NewLockboxApp(address common.Address, backend bind.ContractBackend) (*LockboxApp, error) {
	contract, err := bindLockboxApp(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockboxApp{LockboxAppCaller: LockboxAppCaller{contract: contract}, LockboxAppTransactor: LockboxAppTransactor{contract: contract}}, nil
}

// NewLockboxAppCaller creates a new read-only instance of LockboxApp, bound to a specific deployed contract.
func NewLockboxAppCaller(address common.Address, caller bind.ContractCaller) (*LockboxAppCaller, error) {
	contract, err := bindLockboxApp(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &LockboxAppCaller{contract: contract}, nil
}

// NewLockboxAppTransactor creates a new write-only instance of LockboxApp, bound to a specific deployed contract.
func NewLockboxAppTransactor(address common.Address, transactor bind.ContractTransactor) (*LockboxAppTransactor, error) {
	contract, err := bindLockboxApp(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &LockboxAppTransactor{contract: contract}, nil
}

// bindLockboxApp binds a generic wrapper to an already deployed contract.
func bindLockboxApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockboxAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockboxApp *LockboxAppRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockboxApp.Contract.LockboxAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockboxApp *LockboxAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockboxApp.Contract.LockboxAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockboxApp *LockboxAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockboxApp.Contract.LockboxAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockboxApp *LockboxAppCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockboxApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockboxApp *LockboxAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockboxApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockboxApp *LockboxAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockboxApp.Contract.contract.Transact(opts, method, params...)
}

// GetCardDetails is a free data retrieval call binding the contract method 0x870ab12d.
//
// Solidity: function getCardDetails(addr address, name string) constant returns(cardno uint64, edate uint256, cvv uint256)
func (_LockboxApp *LockboxAppCaller) GetCardDetails(opts *bind.CallOpts, addr common.Address, name string) (struct {
	Cardno uint64
	Edate  *big.Int
	Cvv    *big.Int
}, error) {
	ret := new(struct {
		Cardno uint64
		Edate  *big.Int
		Cvv    *big.Int
	})
	out := ret
	err := _LockboxApp.contract.Call(opts, out, "getCardDetails", addr, name)
	return *ret, err
}

// GetCardDetails is a free data retrieval call binding the contract method 0x870ab12d.
//
// Solidity: function getCardDetails(addr address, name string) constant returns(cardno uint64, edate uint256, cvv uint256)
func (_LockboxApp *LockboxAppSession) GetCardDetails(addr common.Address, name string) (struct {
	Cardno uint64
	Edate  *big.Int
	Cvv    *big.Int
}, error) {
	return _LockboxApp.Contract.GetCardDetails(&_LockboxApp.CallOpts, addr, name)
}

// GetCardDetails is a free data retrieval call binding the contract method 0x870ab12d.
//
// Solidity: function getCardDetails(addr address, name string) constant returns(cardno uint64, edate uint256, cvv uint256)
func (_LockboxApp *LockboxAppCallerSession) GetCardDetails(addr common.Address, name string) (struct {
	Cardno uint64
	Edate  *big.Int
	Cvv    *big.Int
}, error) {
	return _LockboxApp.Contract.GetCardDetails(&_LockboxApp.CallOpts, addr, name)
}

// GetDob is a free data retrieval call binding the contract method 0x5d936ce0.
//
// Solidity: function getDob(addr address) constant returns(num uint256)
func (_LockboxApp *LockboxAppCaller) GetDob(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockboxApp.contract.Call(opts, out, "getDob", addr)
	return *ret0, err
}

// GetDob is a free data retrieval call binding the contract method 0x5d936ce0.
//
// Solidity: function getDob(addr address) constant returns(num uint256)
func (_LockboxApp *LockboxAppSession) GetDob(addr common.Address) (*big.Int, error) {
	return _LockboxApp.Contract.GetDob(&_LockboxApp.CallOpts, addr)
}

// GetDob is a free data retrieval call binding the contract method 0x5d936ce0.
//
// Solidity: function getDob(addr address) constant returns(num uint256)
func (_LockboxApp *LockboxAppCallerSession) GetDob(addr common.Address) (*big.Int, error) {
	return _LockboxApp.Contract.GetDob(&_LockboxApp.CallOpts, addr)
}

// GetLicenseNo is a free data retrieval call binding the contract method 0x27ebbc2f.
//
// Solidity: function getLicenseNo(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppCaller) GetLicenseNo(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockboxApp.contract.Call(opts, out, "getLicenseNo", addr)
	return *ret0, err
}

// GetLicenseNo is a free data retrieval call binding the contract method 0x27ebbc2f.
//
// Solidity: function getLicenseNo(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppSession) GetLicenseNo(addr common.Address) (string, error) {
	return _LockboxApp.Contract.GetLicenseNo(&_LockboxApp.CallOpts, addr)
}

// GetLicenseNo is a free data retrieval call binding the contract method 0x27ebbc2f.
//
// Solidity: function getLicenseNo(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppCallerSession) GetLicenseNo(addr common.Address) (string, error) {
	return _LockboxApp.Contract.GetLicenseNo(&_LockboxApp.CallOpts, addr)
}

// GetLockboxName is a free data retrieval call binding the contract method 0xb87d8ce6.
//
// Solidity: function getLockboxName(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppCaller) GetLockboxName(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockboxApp.contract.Call(opts, out, "getLockboxName", addr)
	return *ret0, err
}

// GetLockboxName is a free data retrieval call binding the contract method 0xb87d8ce6.
//
// Solidity: function getLockboxName(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppSession) GetLockboxName(addr common.Address) (string, error) {
	return _LockboxApp.Contract.GetLockboxName(&_LockboxApp.CallOpts, addr)
}

// GetLockboxName is a free data retrieval call binding the contract method 0xb87d8ce6.
//
// Solidity: function getLockboxName(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppCallerSession) GetLockboxName(addr common.Address) (string, error) {
	return _LockboxApp.Contract.GetLockboxName(&_LockboxApp.CallOpts, addr)
}

// GetLoginCredentials is a free data retrieval call binding the contract method 0xab23cceb.
//
// Solidity: function getLoginCredentials(addr address, name string) constant returns(service string, uname string, passwd string)
func (_LockboxApp *LockboxAppCaller) GetLoginCredentials(opts *bind.CallOpts, addr common.Address, name string) (struct {
	Service string
	Uname   string
	Passwd  string
}, error) {
	ret := new(struct {
		Service string
		Uname   string
		Passwd  string
	})
	out := ret
	err := _LockboxApp.contract.Call(opts, out, "getLoginCredentials", addr, name)
	return *ret, err
}

// GetLoginCredentials is a free data retrieval call binding the contract method 0xab23cceb.
//
// Solidity: function getLoginCredentials(addr address, name string) constant returns(service string, uname string, passwd string)
func (_LockboxApp *LockboxAppSession) GetLoginCredentials(addr common.Address, name string) (struct {
	Service string
	Uname   string
	Passwd  string
}, error) {
	return _LockboxApp.Contract.GetLoginCredentials(&_LockboxApp.CallOpts, addr, name)
}

// GetLoginCredentials is a free data retrieval call binding the contract method 0xab23cceb.
//
// Solidity: function getLoginCredentials(addr address, name string) constant returns(service string, uname string, passwd string)
func (_LockboxApp *LockboxAppCallerSession) GetLoginCredentials(addr common.Address, name string) (struct {
	Service string
	Uname   string
	Passwd  string
}, error) {
	return _LockboxApp.Contract.GetLoginCredentials(&_LockboxApp.CallOpts, addr, name)
}

// GetSSN is a free data retrieval call binding the contract method 0xfffcd07c.
//
// Solidity: function getSSN(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppCaller) GetSSN(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LockboxApp.contract.Call(opts, out, "getSSN", addr)
	return *ret0, err
}

// GetSSN is a free data retrieval call binding the contract method 0xfffcd07c.
//
// Solidity: function getSSN(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppSession) GetSSN(addr common.Address) (string, error) {
	return _LockboxApp.Contract.GetSSN(&_LockboxApp.CallOpts, addr)
}

// GetSSN is a free data retrieval call binding the contract method 0xfffcd07c.
//
// Solidity: function getSSN(addr address) constant returns(name string)
func (_LockboxApp *LockboxAppCallerSession) GetSSN(addr common.Address) (string, error) {
	return _LockboxApp.Contract.GetSSN(&_LockboxApp.CallOpts, addr)
}

// LockboxAccounts is a free data retrieval call binding the contract method 0x096ce447.
//
// Solidity: function lockboxAccounts( address) constant returns(addr address, name string, ssn string, license string, dob uint256, initialized bool)
func (_LockboxApp *LockboxAppCaller) LockboxAccounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr        common.Address
	Name        string
	Ssn         string
	License     string
	Dob         *big.Int
	Initialized bool
}, error) {
	ret := new(struct {
		Addr        common.Address
		Name        string
		Ssn         string
		License     string
		Dob         *big.Int
		Initialized bool
	})
	out := ret
	err := _LockboxApp.contract.Call(opts, out, "lockboxAccounts", arg0)
	return *ret, err
}

// LockboxAccounts is a free data retrieval call binding the contract method 0x096ce447.
//
// Solidity: function lockboxAccounts( address) constant returns(addr address, name string, ssn string, license string, dob uint256, initialized bool)
func (_LockboxApp *LockboxAppSession) LockboxAccounts(arg0 common.Address) (struct {
	Addr        common.Address
	Name        string
	Ssn         string
	License     string
	Dob         *big.Int
	Initialized bool
}, error) {
	return _LockboxApp.Contract.LockboxAccounts(&_LockboxApp.CallOpts, arg0)
}

// LockboxAccounts is a free data retrieval call binding the contract method 0x096ce447.
//
// Solidity: function lockboxAccounts( address) constant returns(addr address, name string, ssn string, license string, dob uint256, initialized bool)
func (_LockboxApp *LockboxAppCallerSession) LockboxAccounts(arg0 common.Address) (struct {
	Addr        common.Address
	Name        string
	Ssn         string
	License     string
	Dob         *big.Int
	Initialized bool
}, error) {
	return _LockboxApp.Contract.LockboxAccounts(&_LockboxApp.CallOpts, arg0)
}

// LockboxAddresses is a free data retrieval call binding the contract method 0x106061e5.
//
// Solidity: function lockboxAddresses( uint256) constant returns(address)
func (_LockboxApp *LockboxAppCaller) LockboxAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockboxApp.contract.Call(opts, out, "lockboxAddresses", arg0)
	return *ret0, err
}

// LockboxAddresses is a free data retrieval call binding the contract method 0x106061e5.
//
// Solidity: function lockboxAddresses( uint256) constant returns(address)
func (_LockboxApp *LockboxAppSession) LockboxAddresses(arg0 *big.Int) (common.Address, error) {
	return _LockboxApp.Contract.LockboxAddresses(&_LockboxApp.CallOpts, arg0)
}

// LockboxAddresses is a free data retrieval call binding the contract method 0x106061e5.
//
// Solidity: function lockboxAddresses( uint256) constant returns(address)
func (_LockboxApp *LockboxAppCallerSession) LockboxAddresses(arg0 *big.Int) (common.Address, error) {
	return _LockboxApp.Contract.LockboxAddresses(&_LockboxApp.CallOpts, arg0)
}

// LockBoxApp is a paid mutator transaction binding the contract method 0x7bf699dd.
//
// Solidity: function LockBoxApp() returns()
func (_LockboxApp *LockboxAppTransactor) LockBoxApp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockboxApp.contract.Transact(opts, "LockBoxApp")
}

// LockBoxApp is a paid mutator transaction binding the contract method 0x7bf699dd.
//
// Solidity: function LockBoxApp() returns()
func (_LockboxApp *LockboxAppSession) LockBoxApp() (*types.Transaction, error) {
	return _LockboxApp.Contract.LockBoxApp(&_LockboxApp.TransactOpts)
}

// LockBoxApp is a paid mutator transaction binding the contract method 0x7bf699dd.
//
// Solidity: function LockBoxApp() returns()
func (_LockboxApp *LockboxAppTransactorSession) LockBoxApp() (*types.Transaction, error) {
	return _LockboxApp.Contract.LockBoxApp(&_LockboxApp.TransactOpts)
}

// AddCardDetails is a paid mutator transaction binding the contract method 0x84ab686d.
//
// Solidity: function addCardDetails(cname string, cardno uint64, edate uint256, cvv uint256) returns(success bool)
func (_LockboxApp *LockboxAppTransactor) AddCardDetails(opts *bind.TransactOpts, cname string, cardno uint64, edate *big.Int, cvv *big.Int) (*types.Transaction, error) {
	return _LockboxApp.contract.Transact(opts, "addCardDetails", cname, cardno, edate, cvv)
}

// AddCardDetails is a paid mutator transaction binding the contract method 0x84ab686d.
//
// Solidity: function addCardDetails(cname string, cardno uint64, edate uint256, cvv uint256) returns(success bool)
func (_LockboxApp *LockboxAppSession) AddCardDetails(cname string, cardno uint64, edate *big.Int, cvv *big.Int) (*types.Transaction, error) {
	return _LockboxApp.Contract.AddCardDetails(&_LockboxApp.TransactOpts, cname, cardno, edate, cvv)
}

// AddCardDetails is a paid mutator transaction binding the contract method 0x84ab686d.
//
// Solidity: function addCardDetails(cname string, cardno uint64, edate uint256, cvv uint256) returns(success bool)
func (_LockboxApp *LockboxAppTransactorSession) AddCardDetails(cname string, cardno uint64, edate *big.Int, cvv *big.Int) (*types.Transaction, error) {
	return _LockboxApp.Contract.AddCardDetails(&_LockboxApp.TransactOpts, cname, cardno, edate, cvv)
}

// AddLoginCredentials is a paid mutator transaction binding the contract method 0xb6cd1856.
//
// Solidity: function addLoginCredentials(service string, uname string, passwd string) returns(success bool)
func (_LockboxApp *LockboxAppTransactor) AddLoginCredentials(opts *bind.TransactOpts, service string, uname string, passwd string) (*types.Transaction, error) {
	return _LockboxApp.contract.Transact(opts, "addLoginCredentials", service, uname, passwd)
}

// AddLoginCredentials is a paid mutator transaction binding the contract method 0xb6cd1856.
//
// Solidity: function addLoginCredentials(service string, uname string, passwd string) returns(success bool)
func (_LockboxApp *LockboxAppSession) AddLoginCredentials(service string, uname string, passwd string) (*types.Transaction, error) {
	return _LockboxApp.Contract.AddLoginCredentials(&_LockboxApp.TransactOpts, service, uname, passwd)
}

// AddLoginCredentials is a paid mutator transaction binding the contract method 0xb6cd1856.
//
// Solidity: function addLoginCredentials(service string, uname string, passwd string) returns(success bool)
func (_LockboxApp *LockboxAppTransactorSession) AddLoginCredentials(service string, uname string, passwd string) (*types.Transaction, error) {
	return _LockboxApp.Contract.AddLoginCredentials(&_LockboxApp.TransactOpts, service, uname, passwd)
}

// AddPersonalInfo is a paid mutator transaction binding the contract method 0x36ff4a1b.
//
// Solidity: function addPersonalInfo(ssn string, licenseno string, dob uint256) returns(success bool)
func (_LockboxApp *LockboxAppTransactor) AddPersonalInfo(opts *bind.TransactOpts, ssn string, licenseno string, dob *big.Int) (*types.Transaction, error) {
	return _LockboxApp.contract.Transact(opts, "addPersonalInfo", ssn, licenseno, dob)
}

// AddPersonalInfo is a paid mutator transaction binding the contract method 0x36ff4a1b.
//
// Solidity: function addPersonalInfo(ssn string, licenseno string, dob uint256) returns(success bool)
func (_LockboxApp *LockboxAppSession) AddPersonalInfo(ssn string, licenseno string, dob *big.Int) (*types.Transaction, error) {
	return _LockboxApp.Contract.AddPersonalInfo(&_LockboxApp.TransactOpts, ssn, licenseno, dob)
}

// AddPersonalInfo is a paid mutator transaction binding the contract method 0x36ff4a1b.
//
// Solidity: function addPersonalInfo(ssn string, licenseno string, dob uint256) returns(success bool)
func (_LockboxApp *LockboxAppTransactorSession) AddPersonalInfo(ssn string, licenseno string, dob *big.Int) (*types.Transaction, error) {
	return _LockboxApp.Contract.AddPersonalInfo(&_LockboxApp.TransactOpts, ssn, licenseno, dob)
}

// CreateLockbox is a paid mutator transaction binding the contract method 0x23bfa8e5.
//
// Solidity: function createLockbox(lname string) returns(success bool)
func (_LockboxApp *LockboxAppTransactor) CreateLockbox(opts *bind.TransactOpts, lname string) (*types.Transaction, error) {
	return _LockboxApp.contract.Transact(opts, "createLockbox", lname)
}

// CreateLockbox is a paid mutator transaction binding the contract method 0x23bfa8e5.
//
// Solidity: function createLockbox(lname string) returns(success bool)
func (_LockboxApp *LockboxAppSession) CreateLockbox(lname string) (*types.Transaction, error) {
	return _LockboxApp.Contract.CreateLockbox(&_LockboxApp.TransactOpts, lname)
}

// CreateLockbox is a paid mutator transaction binding the contract method 0x23bfa8e5.
//
// Solidity: function createLockbox(lname string) returns(success bool)
func (_LockboxApp *LockboxAppTransactorSession) CreateLockbox(lname string) (*types.Transaction, error) {
	return _LockboxApp.Contract.CreateLockbox(&_LockboxApp.TransactOpts, lname)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_LockboxApp *LockboxAppTransactor) StringToBytes32(opts *bind.TransactOpts, source string) (*types.Transaction, error) {
	return _LockboxApp.contract.Transact(opts, "stringToBytes32", source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_LockboxApp *LockboxAppSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _LockboxApp.Contract.StringToBytes32(&_LockboxApp.TransactOpts, source)
}

// StringToBytes32 is a paid mutator transaction binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(source string) returns(result bytes32)
func (_LockboxApp *LockboxAppTransactorSession) StringToBytes32(source string) (*types.Transaction, error) {
	return _LockboxApp.Contract.StringToBytes32(&_LockboxApp.TransactOpts, source)
}
