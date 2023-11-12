// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleLendingTaskManager

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

// BLSOperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type BLSOperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// BLSOperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type BLSOperatorStateRetrieverOperator struct {
	OperatorId [32]byte
	Stake      *big.Int
}

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleLendingTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleLendingTaskManagerTask struct {
	TaskData                  []byte
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleLendingTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleLendingTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Response           []byte
}

// IIncredibleLendingTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleLendingTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// ContractIncredibleLendingTaskManagerMetaData contains all meta data concerning the ContractIncredibleLendingTaskManager contract.
var ContractIncredibleLendingTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"taskData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIIncredibleLendingTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"}],\"name\":\"NewTaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedSuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedUnsuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIIncredibleLendingTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskResponsedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIIncredibleLendingTaskManager.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"}],\"name\":\"TaskResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allTaskHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allTaskResponses\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubkeyRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSPubkeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96[]\",\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\"}],\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debtAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"createNewTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getCheckSignaturesIndices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structBLSOperatorStateRetriever.CheckSignaturesIndices\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBLSRegistryCoordinatorWithIndices\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structBLSOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskResponseWindowBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIncredibleLendingProtocol\",\"name\":\"_lendingProtocol\",\"type\":\"address\"},{\"internalType\":\"contractOnchainDepthOracle\",\"name\":\"_onchainDepthOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_generator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTaskNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingprotocol\",\"outputs\":[{\"internalType\":\"contractIncredibleLendingProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onchainDepthOracle\",\"outputs\":[{\"internalType\":\"contractOnchainDepthOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"taskData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"internalType\":\"structIIncredibleLendingTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"internalType\":\"structIIncredibleLendingTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskResponsedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"internalType\":\"structIIncredibleLendingTaskManager.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\"}],\"name\":\"raiseAndResolveChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"taskData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"internalType\":\"structIIncredibleLendingTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"internalType\":\"structIIncredibleLendingTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"}],\"name\":\"respondToTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"taskSuccesfullyChallenged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"apk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"}],\"name\":\"trySignatureAndApkVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pairingSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"siganatureIsValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200552038038062005520833981016040819052620000359162000169565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001b0565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001b0565b6001600160a01b031660c0525063ffffffff1660e05250620001d7565b6001600160a01b03811681146200016657600080fd5b50565b600080604083850312156200017d57600080fd5b82516200018a8162000150565b602084015190925063ffffffff81168114620001a557600080fd5b809150509250929050565b600060208284031215620001c357600080fd5b8151620001d08162000150565b9392505050565b60805160a05160c05160e0516152d7620002496000396000818161026701528181610558015261294201526000818161031c015281816112050152611b670152600081816103fc0152818161215401526122f501526000818161042301528181610fd40152611fb201526152d76000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c80636efb46361161011a5780638da5cb5b116100ad578063cefdc1d41161007c578063cefdc1d414610522578063f2fde38b14610543578063f5c9899d14610556578063f63c5bab1461057c578063fabc1cbc1461058457600080fd5b80638da5cb5b146104d8578063bb4ce85f146104e9578063cb931841146104fc578063cc2a9a5b1461050f57600080fd5b80637afa1eed116100e95780637afa1eed1461048f5780637d032815146104a2578063886f1195146104b55780638b00ce7c146104c857600080fd5b80636efb463614610445578063715018a61461046657806372d18e8d1461046e57806376ea50de1461047c57600080fd5b80633563b0d11161019d5780635ac86ab71161016c5780635ac86ab7146103995780635c975abb146103cc5780635decc3f5146103d457806368304835146103f75780636d14a9871461041e57600080fd5b80633563b0d11461033e5780634c36a7751461035e5780634f739f7414610371578063595c6a671461039157600080fd5b8063245a7bfc116101d9578063245a7bfc1461029e5780632cb223d5146102c95780632d89f6fc146102f75780633561deb11461031757600080fd5b806310d67a2f1461020b578063136439dd14610220578063171f1d5b146102335780631ad4318914610262575b600080fd5b61021e610219366004613e0f565b610597565b005b61021e61022e366004613e33565b610653565b610246610241366004613f9d565b610792565b6040805192151583529015156020830152015b60405180910390f35b6102897f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610259565b609b546102b1906001600160a01b031681565b6040516001600160a01b039091168152602001610259565b6102e96102d736600461400b565b60996020526000908152604090205481565b604051908152602001610259565b6102e961030536600461400b565b60986020526000908152604090205481565b6102b17f000000000000000000000000000000000000000000000000000000000000000081565b61035161034c366004614028565b61091c565b6040516102599190614172565b61021e61036c36600461423f565b610c7f565b61038461037f366004614320565b6113fa565b6040516102599190614425565b61021e611a69565b6103bc6103a73660046144e0565b606654600160ff9092169190911b9081161490565b6040519015158152602001610259565b6066546102e9565b6103bc6103e236600461400b565b609a6020526000908152604090205460ff1681565b6102b17f000000000000000000000000000000000000000000000000000000000000000081565b6102b17f000000000000000000000000000000000000000000000000000000000000000081565b610458610453366004614714565b611b30565b6040516102599291906147d5565b61021e612588565b60975463ffffffff16610289565b609e546102b1906001600160a01b031681565b609c546102b1906001600160a01b031681565b609d546102b1906001600160a01b031681565b6065546102b1906001600160a01b031681565b6097546102899063ffffffff1681565b6033546001600160a01b03166102b1565b61021e6104f736600461481e565b61259c565b61021e61050a366004614899565b612751565b61021e61051d366004614920565b612c5c565b6105356105303660046149a2565b612dc9565b6040516102599291906149e4565b61021e610551366004613e0f565b612f5b565b7f0000000000000000000000000000000000000000000000000000000000000000610289565b610289606481565b61021e610592366004613e33565b612fd1565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060e9190614a05565b6001600160a01b0316336001600160a01b0316146106475760405162461bcd60e51b815260040161063e90614a22565b60405180910390fd5b6106508161312d565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561069b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106bf9190614a7a565b6106db5760405162461bcd60e51b815260040161063e90614a97565b606654818116146107545760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161063e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106107da576107da614adf565b60200201518951600160200201518a602001516000600281106107ff576107ff614adf565b60200201518b6020015160016002811061081b5761081b614adf565b602090810291909101518c518d8301516040516108789a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61089b9190614af5565b905061090e6108b46108ad8884613224565b86906132b5565b6108bc61334a565b6109046108f5856108ef604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613224565b6108fe8c61340a565b906132b5565b886201d4c0613499565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561095e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109829190614a05565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e89190614a05565b9050600085516001600160401b03811115610a0557610a05613e4c565b604051908082528060200260200182016040528015610a3857816020015b6060815260200190600190039081610a235790505b50905060005b8651811015610c74576000878281518110610a5b57610a5b614adf565b016020015160405163889ae3e560e01b815260f89190911c6004820181905263ffffffff8916602483015291506000906001600160a01b0386169063889ae3e590604401600060405180830381865afa158015610abc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ae49190810190614b17565b905080516001600160401b03811115610aff57610aff613e4c565b604051908082528060200260200182016040528015610b4457816020015b6040805180820190915260008082526020820152815260200190600190039081610b1d5790505b50848481518110610b5757610b57614adf565b602002602001018190525060005b8151811015610c69576000828281518110610b8257610b82614adf565b6020908102919091018101516040805180820182528281529051631b32722560e01b81526004810183905260ff8816602482015263ffffffff8e166044820152919350918201906001600160a01b038b1690631b32722590606401602060405180830381865afa158015610bfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1e9190614ba7565b6001600160601b0316815250868681518110610c3c57610c3c614adf565b60200260200101518381518110610c5557610c55614adf565b602090810291909101015250600101610b65565b505050600101610a3e565b509695505050505050565b6000610c8e602085018561400b565b63ffffffff8116600090815260996020526040902054909150610cfd5760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b606482015260840161063e565b8383604051602001610d10929190614c7e565b60408051601f19818403018152918152815160209283012063ffffffff84166000908152609990935291205414610daf5760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606482015260840161063e565b63ffffffff81166000908152609a602052604090205460ff1615610e475760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a40161063e565b6064610e56602085018561400b565b610e609190614cd2565b63ffffffff164363ffffffff161115610ee15760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e000000000000000000606482015260840161063e565b60008080610eef8880614cf6565b810190610efc9190614d3c565b609e546040516315d0f54760e31b815260048101859052602481018490526001600160a01b038084166044830152949750929550909350600092169063ae87aa38906064016020604051808303816000875af1158015610f60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f849190614a7a565b9050801515600103610fcd57604051339063ffffffff8716907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a350505050506113f4565b60006110457f000000000000000000000000000000000000000000000000000000000000000061100060408d018d614cf6565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061034c9250505060408e0160208f0161400b565b9050600087516001600160401b0381111561106257611062613e4c565b60405190808252806020026020018201604052801561108b578160200160208202803683370190505b50905060005b88518110156110e1576110bc8982815181106110af576110af614adf565b60200260200101516136b3565b8282815181106110ce576110ce614adf565b6020908102919091010152600101611091565b5060006110f460408d0160208e0161400b565b82604051602001611106929190614d6a565b604051602081830303815290604052805190602001209050896020013581146111b05760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a40161063e565b600089516001600160401b038111156111cb576111cb613e4c565b6040519080825280602002602001820160405280156111f4578160200160208202803683370190505b50905060005b8a5181101561133e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663187548c86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611261573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112859190614a05565b6001600160a01b031663e8bb9ae68583815181106112a5576112a5614adf565b60200260200101516040518263ffffffff1660e01b81526004016112cb91815260200190565b602060405180830381865afa1580156112e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130c9190614a05565b82828151811061131e5761131e614adf565b6001600160a01b03909216602092830291909101909101526001016111fa565b5063ffffffff89166000908152609a6020526040808220805460ff19166001179055609d548151631895c62d60e31b815291516001600160a01b039091169263c4ae3168926004808201939182900301818387803b15801561139f57600080fd5b505af11580156113b3573d6000803e3d6000fd5b505060405133925063ffffffff8c1691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050505b50505050565b6114256040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611465573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114899190614a05565b90506114b66040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516385020d4960e01b81526001600160a01b038a16906385020d49906114e6908b9089908990600401614db2565b600060405180830381865afa158015611503573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261152b9190810190614df9565b815260405163e192e9ad60e01b81526001600160a01b0383169063e192e9ad9061155d908b908b908b90600401614e87565b600060405180830381865afa15801561157a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115a29190810190614df9565b6040820152856001600160401b038111156115bf576115bf613e4c565b6040519080825280602002602001820160405280156115f257816020015b60608152602001906001900390816115dd5790505b50606082015260005b60ff811687111561197a576000856001600160401b0381111561162057611620613e4c565b604051908082528060200260200182016040528015611649578160200160208202803683370190505b5083606001518360ff168151811061166357611663614adf565b602002602001018190525060005b868110156118845760008c6001600160a01b0316633064620d8a8a8581811061169c5761169c614adf565b905060200201358e886000015186815181106116ba576116ba614adf565b60200260200101516040518463ffffffff1660e01b81526004016116f79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611714573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117389190614ea7565b90508a8a8560ff1681811061174f5761174f614adf565b60016001600160c01b038516919093013560f81c1c8216909103905061187b57856001600160a01b031663480858668a8a8581811061179057611790614adf565b905060200201358d8d8860ff168181106117ac576117ac614adf565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611802573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118269190614ed0565b85606001518560ff168151811061183f5761183f614adf565b6020026020010151848151811061185857611858614adf565b63ffffffff909216602092830291909101909101528261187781614eed565b9350505b50600101611671565b506000816001600160401b0381111561189f5761189f613e4c565b6040519080825280602002602001820160405280156118c8578160200160208202803683370190505b50905060005b8281101561193f5784606001518460ff16815181106118ef576118ef614adf565b6020026020010151818151811061190857611908614adf565b602002602001015182828151811061192257611922614adf565b63ffffffff909216602092830291909101909101526001016118ce565b508084606001518460ff168151811061195a5761195a614adf565b60200260200101819052505050808061197290614f06565b9150506115fb565b506000896001600160a01b0316633561deb16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119df9190614a05565b60405163eda1076360e01b81529091506001600160a01b0382169063eda1076390611a12908b908b908e90600401614f25565b600060405180830381865afa158015611a2f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a579190810190614df9565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611ab1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ad59190614a7a565b611af15760405162461bcd60e51b815260040161063e90614a97565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60408051808201909152606080825260208201526040805180820190915260008082526020820181905290815b86811015611d43577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c1af6b24898984818110611ba657611ba6614adf565b9050013560f81c60f81b60f81c888860a001518581518110611bca57611bca614adf565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611c26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4a9190614f4f565b6001600160401b031916611c6d866040015183815181106110af576110af614adf565b67ffffffffffffffff191614611d095760405162461bcd60e51b8152602060048201526061602482015260008051602061528283398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161063e565b611d3985604001518281518110611d2257611d22614adf565b6020026020010151836132b590919063ffffffff16565b9150600101611b5d565b506040805180820190915260608082526020820152866001600160401b03811115611d7057611d70613e4c565b604051908082528060200260200182016040528015611d99578160200160208202803683370190505b506020820152866001600160401b03811115611db757611db7613e4c565b604051908082528060200260200182016040528015611de0578160200160208202803683370190505b5081526020850151516000906001600160401b03811115611e0357611e03613e4c565b604051908082528060200260200182016040528015611e2c578160200160208202803683370190505b50905060008660200151516001600160401b03811115611e4e57611e4e613e4c565b604051908082528060200260200182016040528015611e77578160200160208202803683370190505b5090506000611ebb8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506136f692505050565b905060005b88602001515181101561211c57611ee6896020015182815181106110af576110af614adf565b848281518110611ef857611ef8614adf565b60209081029190910101528015611fb05783611f15600183614f7a565b81518110611f2557611f25614adf565b602002602001015160001c848281518110611f4257611f42614adf565b602002602001015160001c11611fb0576040805162461bcd60e51b815260206004820152602481019190915260008051602061528283398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161063e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633064620d858381518110611ff157611ff1614adf565b60200260200101518c8c60000151858151811061201057612010614adf565b60200260200101516040518463ffffffff1660e01b815260040161204d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561206a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061208e9190614ea7565b6001600160c01b03168382815181106120a9576120a9614adf565b60200260200101818152505061211261210b6120df848685815181106120d1576120d1614adf565b60200260200101511661385a565b6121058c6020015185815181106120f8576120f8614adf565b602002602001015161388b565b90613926565b87906132b5565b9550600101611ec0565b505060005b60ff81168a111561245c5760008b8b8360ff1681811061214357612143614adf565b9050013560f81c60f81b60f81c90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c56828c8c60c001518660ff168151811061219c5761219c614adf565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156121f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061221c9190614ba7565b85602001518360ff168151811061223557612235614adf565b6001600160601b03909216602092830291909101820152850151805160ff841690811061226457612264614adf565b602002602001015185600001518360ff168151811061228557612285614adf565b60200260200101906001600160601b031690816001600160601b03168152505060005b8960200151518163ffffffff1610156124525760006122ee858363ffffffff16815181106122d8576122d8614adf565b60200260200101518460ff161c60019081161490565b1561243f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a43cde89848e898663ffffffff168151811061233c5761233c614adf565b60200260200101518f60e001518960ff168151811061235d5761235d614adf565b60200260200101518663ffffffff168151811061237c5761237c614adf565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156123e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124049190614ba7565b8751805160ff871690811061241b5761241b614adf565b6020026020010181815161242f9190614f8d565b6001600160601b03169052506001015b508061244a81614fad565b9150506122a8565b5050600101612121565b50506000806124758c868a606001518b60800151610792565b91509150816124e65760405162461bcd60e51b8152602060048201526043602482015260008051602061528283398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161063e565b806125475760405162461bcd60e51b8152602060048201526039602482015260008051602061528283398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161063e565b50506000878260405160200161255e929190614d6a565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612590613a0a565b61259a6000613a64565b565b609c546001600160a01b031633146126005760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b606482015260840161063e565b604080516080808201835260608083526000602080850182815285870184905283860192835286518083018e90528088018d90526001600160a01b038c168186015287518082039095018552909401865291845263ffffffff4381169093529187169091528251601f85018290048202810182019093528383529091908490849081908401838280828437600092019190915250505050604080830191909152516126af908290602001615020565b60408051601f1981840301815282825280516020918201206097805463ffffffff9081166000908152609890945293909220555416907f2b864a38ca408a0f5b11acbe8fac005170342bc07398bba9a49febd9e6587a6690612712908490615020565b60405180910390a260975461272e9063ffffffff166001614cd2565b6097805463ffffffff191663ffffffff9290921691909117905550505050505050565b609b546001600160a01b031633146127ab5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c6572000000604482015260640161063e565b60006127bd604085016020860161400b565b90503660006127cf6040870187614cf6565b909250905060006127e6608088016060890161400b565b9050609860006127f9602089018961400b565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016128259190615087565b60405160208183030381529060405280519060200120146128ae5760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e7472616374000000606482015260840161063e565b60006099816128c060208a018a61400b565b63ffffffff1663ffffffff168152602001908152602001600020541461293d5760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b606482015260840161063e565b6129677f000000000000000000000000000000000000000000000000000000000000000085614cd2565b63ffffffff164363ffffffff1611156129d85760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b606482015260840161063e565b60006129e76020880188614cf6565b6040516020016129f8929190615114565b604051602081830303815290604052805190602001209050600080612a208387878a8c611b30565b9150915060005b85811015612b15578460ff1683602001518281518110612a4957612a49614adf565b6020026020010151612a5b9190615128565b6001600160601b0316606484600001518381518110612a7c57612a7c614adf565b60200260200101516001600160601b0316612a97919061514b565b1015612b0d576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d606482015260840161063e565b600101612a27565b5060408051808201825263ffffffff43168152602080820184905291519091612b42918c91849101615162565b60405160208183030381529060405280519060200120609960008c6000016020810190612b6f919061400b565b63ffffffff168152602080820192909252604001600090812092909255612b98908c018c614cf6565b810190612ba59190615195565b905080612c1557609d60009054906101000a90046001600160a01b03166001600160a01b031663c4ae31686040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612bfc57600080fd5b505af1158015612c10573d6000803e3d6000fd5b505050505b7f9b96c981c7c70a9f1702abb044782746c11d090f58ea34b12daf2cc53cf8ab5f8b83604051612c46929190615162565b60405180910390a1505050505050505050505050565b600054610100900460ff1615808015612c7c5750600054600160ff909116105b80612c965750303b158015612c96575060005460ff166001145b612cf95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161063e565b6000805460ff191660011790558015612d1c576000805461ff0019166101001790555b612d27876000613ab6565b612d3084613a64565b609d80546001600160a01b038089166001600160a01b031992831617909255609b8054868416908316179055609c8054858416908316179055609e8054928816929091169190911790558015612dc0576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612e0457612e04614adf565b60209081029190910101526040516385020d4960e01b81526000906001600160a01b038816906385020d4990612e4090889086906004016151b2565b600060405180830381865afa158015612e5d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612e859190810190614df9565b600081518110612e9757612e97614adf565b6020908102919091010151604051633064620d60e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b03891690633064620d90606401602060405180830381865afa158015612f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f279190614ea7565b6001600160c01b031690506000612f3d82613ba0565b905081612f4b8a838a61091c565b9550955050505050935093915050565b612f63613a0a565b6001600160a01b038116612fc85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161063e565b61065081613a64565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613024573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130489190614a05565b6001600160a01b0316336001600160a01b0316146130785760405162461bcd60e51b815260040161063e90614a22565b6066541981196066541916146130f65760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161063e565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610787565b6001600160a01b0381166131bb5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161063e565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613240613d20565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808061326f57fe5b50806132ad5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161063e565b505092915050565b60408051808201909152600080825260208201526132d1613d3e565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808061330c57fe5b50806132ad5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161063e565b613352613d5c565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061343a60008051602061526283398151915286614af5565b90505b61344681613bf5565b9093509150600080516020615262833981519152828309830361347f576040805180820190915290815260208101919091529392505050565b60008051602061526283398151915260018208905061343d565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906134cb613d81565b60005b60028110156136865760006134e482600661514b565b90508482600281106134f8576134f8614adf565b6020020151518361350a836000615208565b600c811061351a5761351a614adf565b602002015284826002811061353157613531614adf565b602002015160200151838260016135489190615208565b600c811061355857613558614adf565b602002015283826002811061356f5761356f614adf565b6020020151515183613582836002615208565b600c811061359257613592614adf565b60200201528382600281106135a9576135a9614adf565b60200201515160016020020151836135c2836003615208565b600c81106135d2576135d2614adf565b60200201528382600281106135e9576135e9614adf565b60200201516020015160006002811061360457613604614adf565b602002015183613615836004615208565b600c811061362557613625614adf565b602002015283826002811061363c5761363c614adf565b60200201516020015160016002811061365757613657614adf565b602002015183613668836005615208565b600c811061367857613678614adf565b6020020152506001016134ce565b5061368f613da0565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000816000015182602001516040516020016136d9929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b60006101008251111561376a5760405162461bcd60e51b815260206004820152603660248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a206044820152756279746573417272617920697320746f6f206c6f6e6760501b606482015260840161063e565b815160000361377b57506000919050565b6000808360008151811061379157613791614adf565b0160200151600160f89190911c81901b92505b8451811015613851578481815181106137bf576137bf614adf565b0160200151600160f89190911c1b9150828216156138455760405162461bcd60e51b815260206004820152603a60248201527f4269746d61705574696c732e62797465734172726179546f4269746d61703a2060448201527f72657065617420656e74727920696e2062797465734172726179000000000000606482015260840161063e565b918117916001016137a4565b50909392505050565b6000805b82156138855761386f600184614f7a565b909216918061387d8161521b565b91505061385e565b92915050565b604080518082019091526000808252602082015281511580156138b057506020820151155b156138ce575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061526283398151915284602001516139019190614af5565b61391990600080516020615262833981519152614f7a565b905292915050565b919050565b60408051808201909152600080825260208201526102008261ffff16106139825760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161063e565b8161ffff16600103613995575081613885565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1611156139ff57600161ffff871660ff83161c811690036139e2576139df84846132b5565b93505b6139ec83846132b5565b92506201fffe600192831b1691016139b1565b509195945050505050565b6033546001600160a01b0316331461259a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161063e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b0316158015613ad757506001600160a01b03821615155b613b595760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161063e565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613b9c8261312d565b5050565b60606000805b610100811015613bee576001811b915083821615613be657828160f81b604051602001613bd4929190615232565b60405160208183030381529060405292505b600101613ba6565b5050919050565b60008080600080516020615262833981519152600360008051602061526283398151915286600080516020615262833981519152888909090890506000613c6b827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615262833981519152613c77565b91959194509092505050565b600080613c82613da0565b613c8a613dbe565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280613cc757fe5b5082613d155760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161063e565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613d6f613ddc565b8152602001613d7c613ddc565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461065057600080fd5b600060208284031215613e2157600080fd5b8135613e2c81613dfa565b9392505050565b600060208284031215613e4557600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613e8457613e84613e4c565b60405290565b60405161010081016001600160401b0381118282101715613e8457613e84613e4c565b604051601f8201601f191681016001600160401b0381118282101715613ed557613ed5613e4c565b604052919050565b600060408284031215613eef57600080fd5b613ef7613e62565b9050813581526020820135602082015292915050565b600082601f830112613f1e57600080fd5b613f26613e62565b806040840185811115613f3857600080fd5b845b81811015613f52578035845260209384019301613f3a565b509095945050505050565b600060808284031215613f6f57600080fd5b613f77613e62565b9050613f838383613f0d565b8152613f928360408401613f0d565b602082015292915050565b6000806000806101208587031215613fb457600080fd5b84359350613fc58660208701613edd565b9250613fd48660608701613f5d565b9150613fe38660e08701613edd565b905092959194509250565b63ffffffff8116811461065057600080fd5b803561392181613fee565b60006020828403121561401d57600080fd5b8135613e2c81613fee565b60008060006060848603121561403d57600080fd5b833561404881613dfa565b92506020848101356001600160401b038082111561406557600080fd5b818701915087601f83011261407957600080fd5b81358181111561408b5761408b613e4c565b61409d601f8201601f19168501613ead565b915080825288848285010111156140b357600080fd5b80848401858401376000848284010152508094505050506140d660408501614000565b90509250925092565b600082825180855260208086019550808260051b8401018186016000805b8581101561416457868403601f19018a52825180518086529086019086860190845b8181101561414f578351805184528901516001600160601b0316898401529288019260409092019160010161411f565b50509a86019a945050918401916001016140fd565b509198975050505050505050565b602081526000613e2c60208301846140df565b60006080828403121561419757600080fd5b50919050565b60006040828403121561419757600080fd5b60006001600160401b038211156141c8576141c8613e4c565b5060051b60200190565b600082601f8301126141e357600080fd5b813560206141f86141f3836141af565b613ead565b8083825260208201915060208460061b87010193508684111561421a57600080fd5b602086015b84811015610c74576142318882613edd565b83529183019160400161421f565b60008060008060a0858703121561425557600080fd5b84356001600160401b038082111561426c57600080fd5b61427888838901614185565b9550602087013591508082111561428e57600080fd5b61429a8883890161419d565b94506142a9886040890161419d565b935060808701359150808211156142bf57600080fd5b506142cc878288016141d2565b91505092959194509250565b60008083601f8401126142ea57600080fd5b5081356001600160401b0381111561430157600080fd5b60208301915083602082850101111561431957600080fd5b9250929050565b6000806000806000806080878903121561433957600080fd5b863561434481613dfa565b9550602087013561435481613fee565b945060408701356001600160401b038082111561437057600080fd5b61437c8a838b016142d8565b9096509450606089013591508082111561439557600080fd5b818901915089601f8301126143a957600080fd5b8135818111156143b857600080fd5b8a60208260051b85010111156143cd57600080fd5b6020830194508093505050509295509295509295565b60008151808452602080850194506020840160005b8381101561441a57815163ffffffff16875295820195908201906001016143f8565b509495945050505050565b60006020808352835160808285015261444160a08501826143e3565b905081850151601f198086840301604087015261445e83836143e3565b9250604087015191508086840301606087015261447b83836143e3565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156144d257848783030184526144c08287516143e3565b958801959388019391506001016144a6565b509998505050505050505050565b6000602082840312156144f257600080fd5b813560ff81168114613e2c57600080fd5b600082601f83011261451457600080fd5b813560206145246141f3836141af565b8083825260208201915060208460051b87010193508684111561454657600080fd5b602086015b84811015610c7457803561455e81613fee565b835291830191830161454b565b600082601f83011261457c57600080fd5b8135602061458c6141f3836141af565b82815260059290921b840181019181810190868411156145ab57600080fd5b8286015b84811015610c745780356001600160401b038111156145ce5760008081fd5b6145dc8986838b0101614503565b8452509183019183016145af565b600061018082840312156145fd57600080fd5b614605613e8a565b905081356001600160401b038082111561461e57600080fd5b61462a85838601614503565b8352602084013591508082111561464057600080fd5b61464c858386016141d2565b6020840152604084013591508082111561466557600080fd5b614671858386016141d2565b60408401526146838560608601613f5d565b60608401526146958560e08601613edd565b60808401526101208401359150808211156146af57600080fd5b6146bb85838601614503565b60a08401526101408401359150808211156146d557600080fd5b6146e185838601614503565b60c08401526101608401359150808211156146fb57600080fd5b506147088482850161456b565b60e08301525092915050565b60008060008060006080868803121561472c57600080fd5b8535945060208601356001600160401b038082111561474a57600080fd5b61475689838a016142d8565b90965094506040880135915061476b82613fee565b9092506060870135908082111561478157600080fd5b5061478e888289016145ea565b9150509295509295909350565b60008151808452602080850194506020840160005b8381101561441a5781516001600160601b0316875295820195908201906001016147b0565b60408152600083516040808401526147f0608084018261479b565b90506020850151603f1984830301606085015261480d828261479b565b925050508260208301529392505050565b60008060008060008060a0878903121561483757600080fd5b8635955060208701359450604087013561485081613dfa565b9350606087013561486081613fee565b925060808701356001600160401b0381111561487b57600080fd5b61488789828a016142d8565b979a9699509497509295939492505050565b6000806000606084860312156148ae57600080fd5b83356001600160401b03808211156148c557600080fd5b6148d187838801614185565b945060208601359150808211156148e757600080fd5b6148f38783880161419d565b9350604086013591508082111561490957600080fd5b50614916868287016145ea565b9150509250925092565b60008060008060008060c0878903121561493957600080fd5b863561494481613dfa565b9550602087013561495481613dfa565b9450604087013561496481613dfa565b9350606087013561497481613dfa565b9250608087013561498481613dfa565b915060a087013561499481613dfa565b809150509295509295509295565b6000806000606084860312156149b757600080fd5b83356149c281613dfa565b92506020840135915060408401356149d981613fee565b809150509250925092565b8281526040602082015260006149fd60408301846140df565b949350505050565b600060208284031215614a1757600080fd5b8151613e2c81613dfa565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b801515811461065057600080fd5b600060208284031215614a8c57600080fd5b8151613e2c81614a6c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614b1257634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614b2a57600080fd5b82516001600160401b03811115614b4057600080fd5b8301601f81018513614b5157600080fd5b8051614b5f6141f3826141af565b81815260059190911b82018301908381019087831115614b7e57600080fd5b928401925b82841015614b9c57835182529284019290840190614b83565b979650505050505050565b600060208284031215614bb957600080fd5b81516001600160601b0381168114613e2c57600080fd5b6000808335601e19843603018112614be757600080fd5b83016020810192503590506001600160401b03811115614c0657600080fd5b80360382131561431957600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008135614c4b81613fee565b63ffffffff168352614c606020830183614bd0565b60406020860152614c75604086018284614c15565b95945050505050565b606081526000614c916060830185614c3e565b90508235614c9e81613fee565b63ffffffff8116602084015250602083013560408301529392505050565b634e487b7160e01b600052601160045260246000fd5b63ffffffff818116838216019080821115614cef57614cef614cbc565b5092915050565b6000808335601e19843603018112614d0d57600080fd5b8301803591506001600160401b03821115614d2757600080fd5b60200191503681900382131561431957600080fd5b600080600060608486031215614d5157600080fd5b833592506020840135915060408401356149d981613dfa565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015614da557815185529382019390820190600101614d89565b5092979650505050505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115614ddf57600080fd5b8260051b8085606085013791909101606001949350505050565b60006020808385031215614e0c57600080fd5b82516001600160401b03811115614e2257600080fd5b8301601f81018513614e3357600080fd5b8051614e416141f3826141af565b81815260059190911b82018301908381019087831115614e6057600080fd5b928401925b82841015614b9c578351614e7881613fee565b82529284019290840190614e65565b63ffffffff84168152604060208201526000614c75604083018486614c15565b600060208284031215614eb957600080fd5b81516001600160c01b0381168114613e2c57600080fd5b600060208284031215614ee257600080fd5b8151613e2c81613fee565b600060018201614eff57614eff614cbc565b5060010190565b600060ff821660ff8103614f1c57614f1c614cbc565b60010192915050565b604081526000614f39604083018587614c15565b905063ffffffff83166020830152949350505050565b600060208284031215614f6157600080fd5b815167ffffffffffffffff1981168114613e2c57600080fd5b8181038181111561388557613885614cbc565b6001600160601b03828116828216039080821115614cef57614cef614cbc565b600063ffffffff808316818103614fc657614fc6614cbc565b6001019392505050565b60005b83811015614feb578181015183820152602001614fd3565b50506000910152565b6000815180845261500c816020860160208601614fd0565b601f01601f19169290920160200192915050565b60208152600082516080602084015261503c60a0840182614ff4565b9050602084015163ffffffff808216604086015260408601519150601f1985840301606086015261506d8383614ff4565b925080606087015116608086015250508091505092915050565b6020815260006150978384614bd0565b608060208501526150ac60a085018284614c15565b91505060208401356150bd81613fee565b63ffffffff80821660408601526150d76040870187614bd0565b868503601f1901606088015292506150f0848483614c15565b9350506060860135915061510382613fee565b166080939093019290925250919050565b6020815260006149fd602083018486614c15565b6001600160601b038181168382160280821691908281146132ad576132ad614cbc565b808202811582820484141761388557613885614cbc565b6060815260006151756060830185614c3e565b905063ffffffff8351166020830152602083015160408301529392505050565b6000602082840312156151a757600080fd5b8135613e2c81614a6c565b60006040820163ffffffff8516835260206040602085015281855180845260608601915060208701935060005b818110156151fb578451835293830193918301916001016151df565b5090979650505050505050565b8082018082111561388557613885614cbc565b600061ffff808316818103614fc657614fc6614cbc565b60008351615244818460208801614fd0565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212201481b81ad5b7b5e0f359ed104d1109c3bb4bd063e381ee11692de5ed652c0d2664736f6c63430008160033",
}

// ContractIncredibleLendingTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleLendingTaskManagerMetaData.ABI instead.
var ContractIncredibleLendingTaskManagerABI = ContractIncredibleLendingTaskManagerMetaData.ABI

// ContractIncredibleLendingTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleLendingTaskManagerMetaData.Bin instead.
var ContractIncredibleLendingTaskManagerBin = ContractIncredibleLendingTaskManagerMetaData.Bin

// DeployContractIncredibleLendingTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleLendingTaskManager to it.
func DeployContractIncredibleLendingTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleLendingTaskManager, error) {
	parsed, err := ContractIncredibleLendingTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleLendingTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleLendingTaskManager{ContractIncredibleLendingTaskManagerCaller: ContractIncredibleLendingTaskManagerCaller{contract: contract}, ContractIncredibleLendingTaskManagerTransactor: ContractIncredibleLendingTaskManagerTransactor{contract: contract}, ContractIncredibleLendingTaskManagerFilterer: ContractIncredibleLendingTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleLendingTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleLendingTaskManager struct {
	ContractIncredibleLendingTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleLendingTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleLendingTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleLendingTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleLendingTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleLendingTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleLendingTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleLendingTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleLendingTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleLendingTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleLendingTaskManagerSession struct {
	Contract     *ContractIncredibleLendingTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractIncredibleLendingTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleLendingTaskManagerCallerSession struct {
	Contract *ContractIncredibleLendingTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// ContractIncredibleLendingTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleLendingTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleLendingTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// ContractIncredibleLendingTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleLendingTaskManagerRaw struct {
	Contract *ContractIncredibleLendingTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleLendingTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleLendingTaskManagerCallerRaw struct {
	Contract *ContractIncredibleLendingTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleLendingTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleLendingTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleLendingTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleLendingTaskManager creates a new instance of ContractIncredibleLendingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleLendingTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleLendingTaskManager, error) {
	contract, err := bindContractIncredibleLendingTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManager{ContractIncredibleLendingTaskManagerCaller: ContractIncredibleLendingTaskManagerCaller{contract: contract}, ContractIncredibleLendingTaskManagerTransactor: ContractIncredibleLendingTaskManagerTransactor{contract: contract}, ContractIncredibleLendingTaskManagerFilterer: ContractIncredibleLendingTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleLendingTaskManagerCaller creates a new read-only instance of ContractIncredibleLendingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleLendingTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleLendingTaskManagerCaller, error) {
	contract, err := bindContractIncredibleLendingTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleLendingTaskManagerTransactor creates a new write-only instance of ContractIncredibleLendingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleLendingTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleLendingTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleLendingTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleLendingTaskManagerFilterer creates a new log filterer instance of ContractIncredibleLendingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleLendingTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleLendingTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleLendingTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleLendingTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleLendingTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleLendingTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleLendingTaskManager.Contract.ContractIncredibleLendingTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.ContractIncredibleLendingTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.ContractIncredibleLendingTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleLendingTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Aggregator(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Aggregator(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleLendingTaskManager.Contract.AllTaskHashes(&_ContractIncredibleLendingTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleLendingTaskManager.Contract.AllTaskHashes(&_ContractIncredibleLendingTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleLendingTaskManager.Contract.AllTaskResponses(&_ContractIncredibleLendingTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleLendingTaskManager.Contract.AllTaskResponses(&_ContractIncredibleLendingTaskManager.CallOpts, arg0)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) BlsPubkeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "blsPubkeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.BlsPubkeyRegistry(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.BlsPubkeyRegistry(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleLendingTaskManager.Contract.CheckSignatures(&_ContractIncredibleLendingTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleLendingTaskManager.Contract.CheckSignatures(&_ContractIncredibleLendingTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Generator(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Generator(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(BLSOperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(BLSOperatorStateRetrieverCheckSignaturesIndices)).(*BLSOperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleLendingTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleLendingTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]BLSOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]BLSOperatorStateRetrieverOperator)).(*[][]BLSOperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetOperatorState(&_ContractIncredibleLendingTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetOperatorState(&_ContractIncredibleLendingTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]BLSOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]BLSOperatorStateRetrieverOperator)).(*[][]BLSOperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetOperatorState0(&_ContractIncredibleLendingTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]BLSOperatorStateRetrieverOperator, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetOperatorState0(&_ContractIncredibleLendingTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.LatestTaskNum(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.LatestTaskNum(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Lendingprotocol is a free data retrieval call binding the contract method 0x7d032815.
//
// Solidity: function lendingprotocol() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) Lendingprotocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "lendingprotocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lendingprotocol is a free data retrieval call binding the contract method 0x7d032815.
//
// Solidity: function lendingprotocol() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Lendingprotocol() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Lendingprotocol(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Lendingprotocol is a free data retrieval call binding the contract method 0x7d032815.
//
// Solidity: function lendingprotocol() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) Lendingprotocol() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Lendingprotocol(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// OnchainDepthOracle is a free data retrieval call binding the contract method 0x76ea50de.
//
// Solidity: function onchainDepthOracle() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) OnchainDepthOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "onchainDepthOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnchainDepthOracle is a free data retrieval call binding the contract method 0x76ea50de.
//
// Solidity: function onchainDepthOracle() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) OnchainDepthOracle() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.OnchainDepthOracle(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// OnchainDepthOracle is a free data retrieval call binding the contract method 0x76ea50de.
//
// Solidity: function onchainDepthOracle() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) OnchainDepthOracle() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.OnchainDepthOracle(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Owner(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Owner(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Paused(&_ContractIncredibleLendingTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Paused(&_ContractIncredibleLendingTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Paused0(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Paused0(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.PauserRegistry(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.PauserRegistry(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.StakeRegistry(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleLendingTaskManager.Contract.StakeRegistry(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TaskNumber(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TaskNumber(&_ContractIncredibleLendingTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleLendingTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleLendingTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractIncredibleLendingTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleLendingTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleLendingTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb4ce85f.
//
// Solidity: function createNewTask(uint256 debtAmt, uint256 collateralAmt, address collateralToken, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, debtAmt *big.Int, collateralAmt *big.Int, collateralToken common.Address, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "createNewTask", debtAmt, collateralAmt, collateralToken, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb4ce85f.
//
// Solidity: function createNewTask(uint256 debtAmt, uint256 collateralAmt, address collateralToken, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) CreateNewTask(debtAmt *big.Int, collateralAmt *big.Int, collateralToken common.Address, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.CreateNewTask(&_ContractIncredibleLendingTaskManager.TransactOpts, debtAmt, collateralAmt, collateralToken, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb4ce85f.
//
// Solidity: function createNewTask(uint256 debtAmt, uint256 collateralAmt, address collateralToken, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) CreateNewTask(debtAmt *big.Int, collateralAmt *big.Int, collateralToken common.Address, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.CreateNewTask(&_ContractIncredibleLendingTaskManager.TransactOpts, debtAmt, collateralAmt, collateralToken, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _pauserRegistry, address _lendingProtocol, address _onchainDepthOracle, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, _lendingProtocol common.Address, _onchainDepthOracle common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, _lendingProtocol, _onchainDepthOracle, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _pauserRegistry, address _lendingProtocol, address _onchainDepthOracle, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Initialize(_pauserRegistry common.Address, _lendingProtocol common.Address, _onchainDepthOracle common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Initialize(&_ContractIncredibleLendingTaskManager.TransactOpts, _pauserRegistry, _lendingProtocol, _onchainDepthOracle, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _pauserRegistry, address _lendingProtocol, address _onchainDepthOracle, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, _lendingProtocol common.Address, _onchainDepthOracle common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Initialize(&_ContractIncredibleLendingTaskManager.TransactOpts, _pauserRegistry, _lendingProtocol, _onchainDepthOracle, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Pause(&_ContractIncredibleLendingTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Pause(&_ContractIncredibleLendingTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.PauseAll(&_ContractIncredibleLendingTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.PauseAll(&_ContractIncredibleLendingTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c36a775.
//
// Solidity: function raiseAndResolveChallenge((bytes,uint32,bytes,uint32) task, (uint32,bytes) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleLendingTaskManagerTask, taskResponse IIncredibleLendingTaskManagerTaskResponse, taskResponseMetadata IIncredibleLendingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c36a775.
//
// Solidity: function raiseAndResolveChallenge((bytes,uint32,bytes,uint32) task, (uint32,bytes) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleLendingTaskManagerTask, taskResponse IIncredibleLendingTaskManagerTaskResponse, taskResponseMetadata IIncredibleLendingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleLendingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x4c36a775.
//
// Solidity: function raiseAndResolveChallenge((bytes,uint32,bytes,uint32) task, (uint32,bytes) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleLendingTaskManagerTask, taskResponse IIncredibleLendingTaskManagerTaskResponse, taskResponseMetadata IIncredibleLendingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleLendingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RenounceOwnership(&_ContractIncredibleLendingTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RenounceOwnership(&_ContractIncredibleLendingTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcb931841.
//
// Solidity: function respondToTask((bytes,uint32,bytes,uint32) task, (uint32,bytes) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleLendingTaskManagerTask, taskResponse IIncredibleLendingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcb931841.
//
// Solidity: function respondToTask((bytes,uint32,bytes,uint32) task, (uint32,bytes) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) RespondToTask(task IIncredibleLendingTaskManagerTask, taskResponse IIncredibleLendingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RespondToTask(&_ContractIncredibleLendingTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xcb931841.
//
// Solidity: function respondToTask((bytes,uint32,bytes,uint32) task, (uint32,bytes) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) RespondToTask(task IIncredibleLendingTaskManagerTask, taskResponse IIncredibleLendingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.RespondToTask(&_ContractIncredibleLendingTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleLendingTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleLendingTaskManager.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TransferOwnership(&_ContractIncredibleLendingTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.TransferOwnership(&_ContractIncredibleLendingTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Unpause(&_ContractIncredibleLendingTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleLendingTaskManager.Contract.Unpause(&_ContractIncredibleLendingTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleLendingTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerInitializedIterator struct {
	Event *ContractIncredibleLendingTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerInitialized)
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
		it.Event = new(ContractIncredibleLendingTaskManagerInitialized)
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
func (it *ContractIncredibleLendingTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleLendingTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerInitializedIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerInitialized)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleLendingTaskManagerInitialized, error) {
	event := new(ContractIncredibleLendingTaskManagerInitialized)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleLendingTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerNewTaskCreated)
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
		it.Event = new(ContractIncredibleLendingTaskManagerNewTaskCreated)
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
func (it *ContractIncredibleLendingTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleLendingTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x2b864a38ca408a0f5b11acbe8fac005170342bc07398bba9a49febd9e6587a66.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes,uint32,bytes,uint32) task)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleLendingTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x2b864a38ca408a0f5b11acbe8fac005170342bc07398bba9a49febd9e6587a66.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes,uint32,bytes,uint32) task)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerNewTaskCreated)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x2b864a38ca408a0f5b11acbe8fac005170342bc07398bba9a49febd9e6587a66.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes,uint32,bytes,uint32) task)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleLendingTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleLendingTaskManagerNewTaskCreated)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleLendingTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractIncredibleLendingTaskManagerOwnershipTransferred)
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
func (it *ContractIncredibleLendingTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleLendingTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleLendingTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleLendingTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerPausedIterator struct {
	Event *ContractIncredibleLendingTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerPaused)
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
		it.Event = new(ContractIncredibleLendingTaskManagerPaused)
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
func (it *ContractIncredibleLendingTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerPaused represents a Paused event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleLendingTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerPausedIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerPaused)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleLendingTaskManagerPaused, error) {
	event := new(ContractIncredibleLendingTaskManagerPaused)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerPauserRegistrySetIterator struct {
	Event *ContractIncredibleLendingTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractIncredibleLendingTaskManagerPauserRegistrySet)
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
func (it *ContractIncredibleLendingTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIncredibleLendingTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerPauserRegistrySetIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerPauserRegistrySet)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIncredibleLendingTaskManagerPauserRegistrySet, error) {
	event := new(ContractIncredibleLendingTaskManagerPauserRegistrySet)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleLendingTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleLendingTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleLendingTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerTaskCompleted)
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
		it.Event = new(ContractIncredibleLendingTaskManagerTaskCompleted)
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
func (it *ContractIncredibleLendingTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleLendingTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerTaskCompletedIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerTaskCompleted)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleLendingTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleLendingTaskManagerTaskCompleted)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleLendingTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerTaskResponded)
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
		it.Event = new(ContractIncredibleLendingTaskManagerTaskResponded)
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
func (it *ContractIncredibleLendingTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleLendingTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleLendingTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x9b96c981c7c70a9f1702abb044782746c11d090f58ea34b12daf2cc53cf8ab5f.
//
// Solidity: event TaskResponded((uint32,bytes) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleLendingTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerTaskRespondedIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x9b96c981c7c70a9f1702abb044782746c11d090f58ea34b12daf2cc53cf8ab5f.
//
// Solidity: event TaskResponded((uint32,bytes) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerTaskResponded)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0x9b96c981c7c70a9f1702abb044782746c11d090f58ea34b12daf2cc53cf8ab5f.
//
// Solidity: event TaskResponded((uint32,bytes) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleLendingTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleLendingTaskManagerTaskResponded)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleLendingTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleLendingTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleLendingTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleLendingTaskManagerUnpaused)
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
		it.Event = new(ContractIncredibleLendingTaskManagerUnpaused)
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
func (it *ContractIncredibleLendingTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleLendingTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleLendingTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleLendingTaskManager contract.
type ContractIncredibleLendingTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleLendingTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleLendingTaskManagerUnpausedIterator{contract: _ContractIncredibleLendingTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleLendingTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleLendingTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleLendingTaskManagerUnpaused)
				if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleLendingTaskManager *ContractIncredibleLendingTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleLendingTaskManagerUnpaused, error) {
	event := new(ContractIncredibleLendingTaskManagerUnpaused)
	if err := _ContractIncredibleLendingTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
