package types

import (
	"errors"

	cstaskmanager "github.com/mergd/ccip-avs/contracts/bindings/IncredibleLendingTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleLendingTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleLendingTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
