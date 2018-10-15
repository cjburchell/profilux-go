package types

import "fmt"

type OperationMode string

const (
	OperationModeNormal             = "Normal"
	OperationModeDiagnostics        = "Diagnostics"
	OperationModeLightTest          = "LightTest"
	OperationModeMaintenance        = "Maintenance"
	OperationModeManualSockets      = "ManualSockets"
	OperationModeManualIllumination = "ManualIllumination"
	OperationModeCanTransparent     = "CanTransparent"
)

var operationModeMap = map[int]string{
	0: OperationModeNormal,
	1: OperationModeDiagnostics,
	3: OperationModeLightTest,
	4: OperationModeMaintenance,
	5: OperationModeManualSockets,
	6: OperationModeManualIllumination,
	7: OperationModeCanTransparent,
}

func GetOperationMode(id int) string {
	if val, ok := operationModeMap[id]; ok {
		return val
	}

	return fmt.Sprintf("Unknown Operation Mode (%d???)", id)
}
