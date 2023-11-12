// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOnChainDepthOracle

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

// ContractOnChainDepthOracleMetaData contains all meta data concerning the ContractOnChainDepthOracle contract.
var ContractOnChainDepthOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumOnchainDepthOracle.ExchangeType\",\"name\":\"exchangeType\",\"type\":\"uint8\"}],\"name\":\"VenueSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tkn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"enumOnchainDepthOracle.ExchangeType\",\"name\":\"exchangeType\",\"type\":\"uint8\"}],\"name\":\"setExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tkn\",\"type\":\"address\"}],\"name\":\"testDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depthOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depthIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deepest\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtAmt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"testDepth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"venues\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"enumOnchainDepthOracle.ExchangeType\",\"name\":\"typeExch\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610bc6380380610bc683398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610b2e610098600039600081816061015281816103f301526105450152610b2e6000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633fc8cef31461005c5780635084efc5146100a0578063a5d2382b146100b5578063ae1d1770146100e3578063ae87aa3814610104575b600080fd5b6100837f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100b36100ae3660046107ce565b610127565b005b6100c86100c336600461081d565b61021d565b60408051938452602084019290925290820152606001610097565b6100f66100f136600461084d565b61072a565b6040516100979291906108b1565b6101176101123660046108ce565b61076d565b6040519015158152602001610097565b6001600160a01b038084166000908152602081815260409182902082518084019093529285168252810183600381111561016357610163610879565b9052815460018101835560009283526020928390208251910180546001600160a01b031981166001600160a01b03909316928317825593830151929390929183916001600160a81b03191617600160a01b8360038111156101c6576101c6610879565b02179055505050816001600160a01b0316836001600160a01b03167fb91203c8fadae2ea949a2213cb1eb05270adb12c6559911a4073464de1b9c7438360405161021091906108fc565b60405180910390a3505050565b6001600160a01b03811660009081526020819052604081205481908190810361024e57506000915081905080610723565b60408051808201909152600080825260208201526001600160a01b038516600081815260208190526040902054906340c10f193061028c848b610920565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b1580156102d257600080fd5b505af11580156102e6573d6000803e3d6000fd5b5050505060005b8181101561071f576001600160a01b038716600090815260208190526040902080548290811061031f5761031f610937565b600091825260209182902060408051808201909152910180546001600160a01b03811683529192909190830190600160a01b900460ff16600381111561036757610367610879565b600381111561037857610378610879565b905250925060008360200151600381111561039557610395610879565b0361050d57825160408051600280825260608201835260009260208301908036833701905050905088816000815181106103d1576103d1610937565b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160018151811061042557610425610937565b6001600160a01b03928316602091820292909201015260405163d06ca61f60e01b815260009184169063d06ca61f90610464908e908690600401610963565b600060405180830381865afa158015610481573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104a991908101906109bc565b90506104cf87826001815181106104c2576104c2610937565b602002602001015161079c565b9650806001815181106104e4576104e4610937565b6020026020010151896104f79190610a7a565b98506105038b89610a7a565b9750505050610717565b60018360200151600381111561052557610525610879565b0361062d5782516040805160a0810182526001600160a01b038a811682527f00000000000000000000000000000000000000000000000000000000000000008116602083019081528284018d8152610bb8606085019081526000608086018181529651636352813560e11b8152865186166004820152935185166024850152915160448401525162ffffff1660648301529351821660848201529192919084169063c6a5026a9060a4016080604051808303816000875af11580156105ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106129190610a8d565b5050509050610621878261079c565b96506104f7818a610a7a565b60028360200151600381111561064557610645610879565b036106fc578251604051630f7c084960e21b815260006004820181905260016024830152604482018b905260648201819052906001600160a01b03831690633df02124906084016020604051808303816000875af11580156106ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106cf9190610adf565b90506106db868261079c565b95506106e78189610a7a565b97506106f38a88610a7a565b96505050610717565b60038360200151600381111561071457610714610879565b50505b6001016102ed565b5050505b9250925092565b6000602052816000526040600020818154811061074657600080fd5b6000918252602090912001546001600160a01b0381169250600160a01b900460ff16905082565b60008061077a858461021d565b9250505083811161078f576000915050610795565b60019150505b9392505050565b60008183116107ab57816107ad565b825b90505b92915050565b6001600160a01b03811681146107cb57600080fd5b50565b6000806000606084860312156107e357600080fd5b83356107ee816107b6565b925060208401356107fe816107b6565b915060408401356004811061081257600080fd5b809150509250925092565b6000806040838503121561083057600080fd5b823591506020830135610842816107b6565b809150509250929050565b6000806040838503121561086057600080fd5b823561086b816107b6565b946020939093013593505050565b634e487b7160e01b600052602160045260246000fd5b600481106108ad57634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b038316815260408101610795602083018461088f565b6000806000606084860312156108e357600080fd5b83359250602084013591506040840135610812816107b6565b602081016107b0828461088f565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176107b0576107b061090a565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60006040820184835260206040602085015281855180845260608601915060208701935060005b818110156109af5784516001600160a01b03168352938301939183019160010161098a565b5090979650505050505050565b600060208083850312156109cf57600080fd5b825167ffffffffffffffff808211156109e757600080fd5b818501915085601f8301126109fb57600080fd5b815181811115610a0d57610a0d61094d565b8060051b604051601f19603f83011681018181108582111715610a3257610a3261094d565b604052918252848201925083810185019188831115610a5057600080fd5b938501935b82851015610a6e57845184529385019392850192610a55565b98975050505050505050565b808201808211156107b0576107b061090a565b60008060008060808587031215610aa357600080fd5b845193506020850151610ab5816107b6565b604086015190935063ffffffff81168114610acf57600080fd5b6060959095015193969295505050565b600060208284031215610af157600080fd5b505191905056fea26469706673582212207e6077bc01629679c28d86e269bb1b0b1cf5f8d6d784139f3f3c2bf24218d8ee64736f6c63430008160033",
}

// ContractOnChainDepthOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOnChainDepthOracleMetaData.ABI instead.
var ContractOnChainDepthOracleABI = ContractOnChainDepthOracleMetaData.ABI

// ContractOnChainDepthOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOnChainDepthOracleMetaData.Bin instead.
var ContractOnChainDepthOracleBin = ContractOnChainDepthOracleMetaData.Bin

// DeployContractOnChainDepthOracle deploys a new Ethereum contract, binding an instance of ContractOnChainDepthOracle to it.
func DeployContractOnChainDepthOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address) (common.Address, *types.Transaction, *ContractOnChainDepthOracle, error) {
	parsed, err := ContractOnChainDepthOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOnChainDepthOracleBin), backend, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOnChainDepthOracle{ContractOnChainDepthOracleCaller: ContractOnChainDepthOracleCaller{contract: contract}, ContractOnChainDepthOracleTransactor: ContractOnChainDepthOracleTransactor{contract: contract}, ContractOnChainDepthOracleFilterer: ContractOnChainDepthOracleFilterer{contract: contract}}, nil
}

// ContractOnChainDepthOracle is an auto generated Go binding around an Ethereum contract.
type ContractOnChainDepthOracle struct {
	ContractOnChainDepthOracleCaller     // Read-only binding to the contract
	ContractOnChainDepthOracleTransactor // Write-only binding to the contract
	ContractOnChainDepthOracleFilterer   // Log filterer for contract events
}

// ContractOnChainDepthOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOnChainDepthOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOnChainDepthOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOnChainDepthOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOnChainDepthOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOnChainDepthOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOnChainDepthOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOnChainDepthOracleSession struct {
	Contract     *ContractOnChainDepthOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractOnChainDepthOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOnChainDepthOracleCallerSession struct {
	Contract *ContractOnChainDepthOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractOnChainDepthOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOnChainDepthOracleTransactorSession struct {
	Contract     *ContractOnChainDepthOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractOnChainDepthOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOnChainDepthOracleRaw struct {
	Contract *ContractOnChainDepthOracle // Generic contract binding to access the raw methods on
}

// ContractOnChainDepthOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOnChainDepthOracleCallerRaw struct {
	Contract *ContractOnChainDepthOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOnChainDepthOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOnChainDepthOracleTransactorRaw struct {
	Contract *ContractOnChainDepthOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOnChainDepthOracle creates a new instance of ContractOnChainDepthOracle, bound to a specific deployed contract.
func NewContractOnChainDepthOracle(address common.Address, backend bind.ContractBackend) (*ContractOnChainDepthOracle, error) {
	contract, err := bindContractOnChainDepthOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOnChainDepthOracle{ContractOnChainDepthOracleCaller: ContractOnChainDepthOracleCaller{contract: contract}, ContractOnChainDepthOracleTransactor: ContractOnChainDepthOracleTransactor{contract: contract}, ContractOnChainDepthOracleFilterer: ContractOnChainDepthOracleFilterer{contract: contract}}, nil
}

// NewContractOnChainDepthOracleCaller creates a new read-only instance of ContractOnChainDepthOracle, bound to a specific deployed contract.
func NewContractOnChainDepthOracleCaller(address common.Address, caller bind.ContractCaller) (*ContractOnChainDepthOracleCaller, error) {
	contract, err := bindContractOnChainDepthOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOnChainDepthOracleCaller{contract: contract}, nil
}

// NewContractOnChainDepthOracleTransactor creates a new write-only instance of ContractOnChainDepthOracle, bound to a specific deployed contract.
func NewContractOnChainDepthOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOnChainDepthOracleTransactor, error) {
	contract, err := bindContractOnChainDepthOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOnChainDepthOracleTransactor{contract: contract}, nil
}

// NewContractOnChainDepthOracleFilterer creates a new log filterer instance of ContractOnChainDepthOracle, bound to a specific deployed contract.
func NewContractOnChainDepthOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOnChainDepthOracleFilterer, error) {
	contract, err := bindContractOnChainDepthOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOnChainDepthOracleFilterer{contract: contract}, nil
}

// bindContractOnChainDepthOracle binds a generic wrapper to an already deployed contract.
func bindContractOnChainDepthOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOnChainDepthOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOnChainDepthOracle.Contract.ContractOnChainDepthOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.ContractOnChainDepthOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.ContractOnChainDepthOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOnChainDepthOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.contract.Transact(opts, method, params...)
}

// Venues is a free data retrieval call binding the contract method 0xae1d1770.
//
// Solidity: function venues(address , uint256 ) view returns(address exchange, uint8 typeExch)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleCaller) Venues(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Exchange common.Address
	TypeExch uint8
}, error) {
	var out []interface{}
	err := _ContractOnChainDepthOracle.contract.Call(opts, &out, "venues", arg0, arg1)

	outstruct := new(struct {
		Exchange common.Address
		TypeExch uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exchange = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TypeExch = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Venues is a free data retrieval call binding the contract method 0xae1d1770.
//
// Solidity: function venues(address , uint256 ) view returns(address exchange, uint8 typeExch)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleSession) Venues(arg0 common.Address, arg1 *big.Int) (struct {
	Exchange common.Address
	TypeExch uint8
}, error) {
	return _ContractOnChainDepthOracle.Contract.Venues(&_ContractOnChainDepthOracle.CallOpts, arg0, arg1)
}

// Venues is a free data retrieval call binding the contract method 0xae1d1770.
//
// Solidity: function venues(address , uint256 ) view returns(address exchange, uint8 typeExch)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleCallerSession) Venues(arg0 common.Address, arg1 *big.Int) (struct {
	Exchange common.Address
	TypeExch uint8
}, error) {
	return _ContractOnChainDepthOracle.Contract.Venues(&_ContractOnChainDepthOracle.CallOpts, arg0, arg1)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOnChainDepthOracle.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleSession) Weth() (common.Address, error) {
	return _ContractOnChainDepthOracle.Contract.Weth(&_ContractOnChainDepthOracle.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleCallerSession) Weth() (common.Address, error) {
	return _ContractOnChainDepthOracle.Contract.Weth(&_ContractOnChainDepthOracle.CallOpts)
}

// SetExchange is a paid mutator transaction binding the contract method 0x5084efc5.
//
// Solidity: function setExchange(address tkn, address exchange, uint8 exchangeType) returns()
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactor) SetExchange(opts *bind.TransactOpts, tkn common.Address, exchange common.Address, exchangeType uint8) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.contract.Transact(opts, "setExchange", tkn, exchange, exchangeType)
}

// SetExchange is a paid mutator transaction binding the contract method 0x5084efc5.
//
// Solidity: function setExchange(address tkn, address exchange, uint8 exchangeType) returns()
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleSession) SetExchange(tkn common.Address, exchange common.Address, exchangeType uint8) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.SetExchange(&_ContractOnChainDepthOracle.TransactOpts, tkn, exchange, exchangeType)
}

// SetExchange is a paid mutator transaction binding the contract method 0x5084efc5.
//
// Solidity: function setExchange(address tkn, address exchange, uint8 exchangeType) returns()
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactorSession) SetExchange(tkn common.Address, exchange common.Address, exchangeType uint8) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.SetExchange(&_ContractOnChainDepthOracle.TransactOpts, tkn, exchange, exchangeType)
}

// TestDepth is a paid mutator transaction binding the contract method 0xa5d2382b.
//
// Solidity: function testDepth(uint256 amount, address tkn) returns(uint256 depthOut, uint256 depthIn, uint256 deepest)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactor) TestDepth(opts *bind.TransactOpts, amount *big.Int, tkn common.Address) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.contract.Transact(opts, "testDepth", amount, tkn)
}

// TestDepth is a paid mutator transaction binding the contract method 0xa5d2382b.
//
// Solidity: function testDepth(uint256 amount, address tkn) returns(uint256 depthOut, uint256 depthIn, uint256 deepest)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleSession) TestDepth(amount *big.Int, tkn common.Address) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.TestDepth(&_ContractOnChainDepthOracle.TransactOpts, amount, tkn)
}

// TestDepth is a paid mutator transaction binding the contract method 0xa5d2382b.
//
// Solidity: function testDepth(uint256 amount, address tkn) returns(uint256 depthOut, uint256 depthIn, uint256 deepest)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactorSession) TestDepth(amount *big.Int, tkn common.Address) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.TestDepth(&_ContractOnChainDepthOracle.TransactOpts, amount, tkn)
}

// TestDepth0 is a paid mutator transaction binding the contract method 0xae87aa38.
//
// Solidity: function testDepth(uint256 collAmt, uint256 debtAmt, address collateral) returns(bool)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactor) TestDepth0(opts *bind.TransactOpts, collAmt *big.Int, debtAmt *big.Int, collateral common.Address) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.contract.Transact(opts, "testDepth0", collAmt, debtAmt, collateral)
}

// TestDepth0 is a paid mutator transaction binding the contract method 0xae87aa38.
//
// Solidity: function testDepth(uint256 collAmt, uint256 debtAmt, address collateral) returns(bool)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleSession) TestDepth0(collAmt *big.Int, debtAmt *big.Int, collateral common.Address) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.TestDepth0(&_ContractOnChainDepthOracle.TransactOpts, collAmt, debtAmt, collateral)
}

// TestDepth0 is a paid mutator transaction binding the contract method 0xae87aa38.
//
// Solidity: function testDepth(uint256 collAmt, uint256 debtAmt, address collateral) returns(bool)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleTransactorSession) TestDepth0(collAmt *big.Int, debtAmt *big.Int, collateral common.Address) (*types.Transaction, error) {
	return _ContractOnChainDepthOracle.Contract.TestDepth0(&_ContractOnChainDepthOracle.TransactOpts, collAmt, debtAmt, collateral)
}

// ContractOnChainDepthOracleVenueSetIterator is returned from FilterVenueSet and is used to iterate over the raw logs and unpacked data for VenueSet events raised by the ContractOnChainDepthOracle contract.
type ContractOnChainDepthOracleVenueSetIterator struct {
	Event *ContractOnChainDepthOracleVenueSet // Event containing the contract specifics and raw log

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
func (it *ContractOnChainDepthOracleVenueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOnChainDepthOracleVenueSet)
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
		it.Event = new(ContractOnChainDepthOracleVenueSet)
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
func (it *ContractOnChainDepthOracleVenueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOnChainDepthOracleVenueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOnChainDepthOracleVenueSet represents a VenueSet event raised by the ContractOnChainDepthOracle contract.
type ContractOnChainDepthOracleVenueSet struct {
	Token        common.Address
	Exchange     common.Address
	ExchangeType uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVenueSet is a free log retrieval operation binding the contract event 0xb91203c8fadae2ea949a2213cb1eb05270adb12c6559911a4073464de1b9c743.
//
// Solidity: event VenueSet(address indexed token, address indexed exchange, uint8 exchangeType)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleFilterer) FilterVenueSet(opts *bind.FilterOpts, token []common.Address, exchange []common.Address) (*ContractOnChainDepthOracleVenueSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _ContractOnChainDepthOracle.contract.FilterLogs(opts, "VenueSet", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return &ContractOnChainDepthOracleVenueSetIterator{contract: _ContractOnChainDepthOracle.contract, event: "VenueSet", logs: logs, sub: sub}, nil
}

// WatchVenueSet is a free log subscription operation binding the contract event 0xb91203c8fadae2ea949a2213cb1eb05270adb12c6559911a4073464de1b9c743.
//
// Solidity: event VenueSet(address indexed token, address indexed exchange, uint8 exchangeType)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleFilterer) WatchVenueSet(opts *bind.WatchOpts, sink chan<- *ContractOnChainDepthOracleVenueSet, token []common.Address, exchange []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _ContractOnChainDepthOracle.contract.WatchLogs(opts, "VenueSet", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOnChainDepthOracleVenueSet)
				if err := _ContractOnChainDepthOracle.contract.UnpackLog(event, "VenueSet", log); err != nil {
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

// ParseVenueSet is a log parse operation binding the contract event 0xb91203c8fadae2ea949a2213cb1eb05270adb12c6559911a4073464de1b9c743.
//
// Solidity: event VenueSet(address indexed token, address indexed exchange, uint8 exchangeType)
func (_ContractOnChainDepthOracle *ContractOnChainDepthOracleFilterer) ParseVenueSet(log types.Log) (*ContractOnChainDepthOracleVenueSet, error) {
	event := new(ContractOnChainDepthOracleVenueSet)
	if err := _ContractOnChainDepthOracle.contract.UnpackLog(event, "VenueSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
