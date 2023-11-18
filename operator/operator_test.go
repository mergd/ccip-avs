package operator

import (
	"context"
	"fmt"
	"testing"
	"time"

	"math/big"

	"encoding/binary"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/mergd/ccip-avs/aggregator"
	aggtypes "github.com/mergd/ccip-avs/aggregator/types"
	cstaskmanager "github.com/mergd/ccip-avs/contracts/bindings/IncredibleLendingTaskManager"
	chainiomocks "github.com/mergd/ccip-avs/core/chainio/mocks"
	operatormocks "github.com/mergd/ccip-avs/operator/mocks"
)

func TestOperator(t *testing.T) {
	operator, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

	t.Run("ProcessNewTaskCreatedLog", func(t *testing.T) {
		newTaskCreatedLog := &cstaskmanager.ContractIncredibleLendingTaskManagerNewTaskCreated{
			TaskIndex: taskIndex,

			// Task: cstaskmanager.IIncredibleLendingTaskManagerTask{
			// 	NumberToBeSquared:         numberToBeSquared,
			// 	TaskCreatedBlock:          1000,
			// 	QuorumNumbers:             aggtypes.QUORUM_NUMBERS,
			// 	QuorumThresholdPercentage: aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			// },
			Raw: types.Log{},
		}

		var data []byte
		// Set the first boolean value to true
		data = append(data, 0) // Append a byte with all bits set to 0
		data[0] |= 1 << 0      // Set the first bit to 1 (true)

		got := operator.ProcessNewTaskCreatedLog(newTaskCreatedLog)
		want := &cstaskmanager.IIncredibleLendingTaskManagerTaskResponse{
			ReferenceTaskIndex: taskIndex,
			Response:           data,
		}
		assert.Equal(t, got, want)
	})

	t.Run("Start", func(t *testing.T) {
		var data []byte

		// Append uint256s
		num1 := uint64(12345)
		data = append(data, make([]byte, 32)...) // Append 32 zero bytes for uint256
		binary.LittleEndian.PutUint64(data[len(data)-32:], num1)

		num2 := uint64(67890)
		data = append(data, make([]byte, 32)...) // Append 32 zero bytes for uint256
		binary.LittleEndian.PutUint64(data[len(data)-32:], num2)

		// Append collateral token address
		address := "0x1234567890123456789012345678901234567890" // Replace with the actual address
		addressBytes := []byte(address)
		data = append(data, addressBytes...)
		// new task event
		newTaskCreatedEvent := &cstaskmanager.ContractIncredibleLendingTaskManagerNewTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.IIncredibleLendingTaskManagerTask{
				TaskData:                  data,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS,
				QuorumThresholdPercentage: aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)
		X, ok := big.NewInt(0).SetString("7926134832136282318561896451042374984489965925674521194255549259381336496956", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("15243507701692917330954619280683582177901049846125926696838777109165913318327", 10)
		assert.True(t, ok)

		signedTaskResponse := &aggregator.SignedTaskResponse{
			TaskResponse: cstaskmanager.IIncredibleLendingTaskManagerTaskResponse{
				ReferenceTaskIndex: taskIndex,
				// NumberSquared:      big.NewInt(0).Mul(numberToBeSquared, numberToBeSquared),
			},
			BlsSignature: bls.Signature{
				G1Point: bls.NewG1Point(X, Y),
			},
			OperatorId: operator.operatorId,
		}
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := operatormocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedTaskResponseToAggregator(signedTaskResponse)
		operator.aggregatorRpcClient = mockAggregatorRpcClient

		mockSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
		mockSubscriber.EXPECT().SubscribeToNewTasks(operator.newTaskCreatedChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}))
		operator.avsSubscriber = mockSubscriber

		mockReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
		mockReader.EXPECT().IsOperatorRegistered(gomock.Any(), operator.operatorAddr).Return(true, nil)
		operator.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			err := operator.Start(ctx)
			assert.Nil(t, err)
		}()
		operator.newTaskCreatedChan <- newTaskCreatedEvent
		time.Sleep(1 * time.Second)

		cancel()
	})

}
