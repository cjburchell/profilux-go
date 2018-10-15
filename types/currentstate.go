package types

type CurrentState string

const CurrentStateOff = "Off"
const CurrentStateOn = "On"

func GetCurrentState(value int) CurrentState {
	if value != 0 {
		return CurrentStateOn
	}

	return CurrentStateOff
}
