package profilux

import (
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/cjburchell/tools-go/math"

	"github.com/cjburchell/profilux-go/code"
	"github.com/cjburchell/profilux-go/protocol"
	"github.com/cjburchell/profilux-go/protocol/http"
	"github.com/cjburchell/profilux-go/protocol/native"
	"github.com/cjburchell/profilux-go/types"

	"github.com/cjburchell/yasls-client-go"
)

const (
	//Blockitems1To10Vint           = 10
	//BlockitemsIlluminationchannel = 8
	//BlockitemsProglogic           = 8
	blockItemsReminder = 4
	//BlockitemsSensorstates        = 8
	//BlockitemsSwitchplug          = 24
	//BlockitemsTimer               = 12
	//Blocksize1To10Vint            = 3
	//BlocksizeIlluminationchannel  = 28
	//BlocksizeProglogic            = 4
	blockSizeReminder = 12
	//BlocksizeSensorstates         = 8
	//BlocksizeSwitchplug           = 1
	//BlocksizeTimer                = 21
	megaBlockSize  = 1000
	SfFeedPause    = 2
	SfMaintenance  = 1
	SfThunderstorm = 3
	SfWaterChange  = 0
)

func getOffset(index, blockCount, blockSize int) int {
	return ((index % blockCount) * blockSize) + ((index / blockCount) * megaBlockSize)
}

type Controller struct {
	p                 protocol.IProtocol
	reminderCount     *int
	sensorCount       *int
	levelSensorCount  *int
	digitalInputCount *int
	timerCount        *int
	lightCount        *int
	pumpCount         *int
	logicCount        *int
	sPortCount        *int
	lPortCount        *int
	callCount         int
}

const ProtocolHTTP = "HTTP"
const ProtocolSocket = "Socket"

// Settings for controller
type Settings struct {
	// Address of the controller
	Address string

	// Port of the controller
	Port int

	// ControllerAddress of the controller
	ControllerAddress int

	// Protocol to communicate with
	Protocol string
}

// NewController  creates a controller
func NewController(settings Settings) (*Controller, error) {

	var p protocol.IProtocol
	var err error
	if settings.Protocol == ProtocolHTTP {
		p, err = http.NewProtocol(settings.Address, settings.Port)
	} else if settings.Protocol == ProtocolSocket {
		p, err = native.NewProtocol(settings.Address, settings.Port, settings.ControllerAddress)
	} else {
		err = errors.WithStack(errors.New("invalid protocol"))
	}

	if err != nil {
		return nil, err
	}

	var controller Controller
	controller.p = p
	return &controller, nil
}

func (controller *Controller) Disconnect() {
	controller.p.Disconnect()
}

func (controller *Controller) ResetStats() {
	controller.callCount = 0
}

func (controller Controller) GetCallCount() int {
	return controller.callCount
}

// region Protocol

func (controller *Controller) getDataText(code int) (string, error) {
	controller.callCount++
	return controller.p.GetDataText(code)
}

func (controller *Controller) getData(code int) (int, error) {
	controller.callCount++
	return controller.p.GetData(code)
}

func (controller *Controller) getDataDate(code int) (time.Time, error) {
	result, err := controller.getData(code)
	if err != nil {
		return time.Now(), err
	}

	timeString := strconv.Itoa(result)

	if len(timeString) == 6 {
		yearValue, _ := strconv.Atoi(timeString[len(timeString)-2:])
		monthValue, _ := strconv.Atoi(timeString[len(timeString)-4 : len(timeString)-2])
		dateValue, _ := strconv.Atoi(timeString[:len(timeString)-4])
		return time.Date(yearValue+2000, time.Month(monthValue), dateValue, 0, 0, 0, 0, time.UTC), nil
	} else if len(timeString) == 7 {
		yearValue, _ := strconv.Atoi(timeString[len(timeString)-3:])
		monthValue, _ := strconv.Atoi(timeString[len(timeString)-5 : len(timeString)-3])
		dateValue, _ := strconv.Atoi(timeString[:len(timeString)-5])
		return time.Date(yearValue+2000, time.Month(monthValue), dateValue, 0, 0, 0, 0, time.UTC), nil
	}

	return time.Now(), err
}

func (controller *Controller) getDataEnum(code int, convert func(int) string) (string, error) {
	result, err := controller.getData(code)
	if err != nil {
		return "", err
	}

	return convert(result), nil
}

func (controller *Controller) getDataCurrentState(code int) (types.CurrentState, error) {
	result, err := controller.getData(code)
	if err != nil {
		return "", err
	}

	return types.GetCurrentState(result), nil
}

func (controller *Controller) getDataCurrentStateConvert(code int, convert func(int) int) (types.CurrentState, error) {
	result, err := controller.getData(code)
	if err != nil {
		return "", err
	}

	return types.GetCurrentState(convert(result)), nil
}

func (controller *Controller) getDataFloat(code int, multiplier float64) (float64, error) {
	result, err := controller.getData(code)
	if err != nil {
		return 0, err
	}

	return float64(result) * multiplier, nil
}

func (controller *Controller) getDataMultiplier(code int, multiplier int) (int, error) {
	result, err := controller.getData(code)
	if err != nil {
		return 0, err
	}

	return result * multiplier, nil
}

func (controller *Controller) getDataBool(code int) (bool, error) {
	result, err := controller.getData(code)
	if err != nil {
		return false, err
	}

	return result != 0, nil
}

func (controller *Controller) getDataBoolConvert(code int, convert func(int) int) (bool, error) {
	result, err := controller.getData(code)
	if err != nil {
		return false, err
	}

	return convert(result) != 0, nil
}

func (controller *Controller) getDataFloatAndRound(code int, multiplier float64, digits int) (float64, error) {
	result, err := controller.getDataFloat(code, multiplier)
	if err != nil {
		return 0, err
	}

	return math.Round(result, digits), nil
}

// endregion

// region Info

func (controller *Controller) GetSoftwareVersion() (float64, error) {
	return controller.getDataFloatAndRound(code.SOFTWAREVERSION, 0.01, 2)
}

func (controller *Controller) GetModel() (types.Model, error) {
	result, err := controller.getDataEnum(code.PRODUCTID, types.GetModel)
	return types.Model(result), err
}

func (controller *Controller) GetSerialNumber() (int, error) {
	return controller.getData(code.SERIALNUMBER)
}

func (controller *Controller) GetSoftwareDate() (time.Time, error) {
	return controller.getDataDate(code.SOFTWAREDATE)
}

func (controller *Controller) GetDeviceAddress() (int, error) {
	return controller.getData(code.ADDRESS)
}

func (controller *Controller) GetLatitude() (float64, error) {
	return controller.getDataFloat(code.LOC_LATITUDE, 0.1)
}

func (controller *Controller) GetLongitude() (float64, error) {
	return controller.getDataFloat(code.LOC_LONGITUDE, 0.1)
}

func (controller *Controller) GetMoonPhase() (float64, error) {
	return controller.getDataFloat(code.MOON_ACTPHASE, 1)
}

func (controller *Controller) GetAlarm() (types.CurrentState, error) {
	return controller.getDataCurrentState(code.ISALARM)
}

func (controller *Controller) GetOperationMode() (types.OperationMode, error) {
	result, err := controller.getDataEnum(code.OPMODE, types.GetOperationMode)
	return types.OperationMode(result), err
}

// endregion

// region Reminders

func (controller *Controller) GetReminderCount() (int, error) {

	if controller.reminderCount == nil {
		count, err := controller.getData(code.GETREMINDERCOUNT)
		if err != nil {
			return 0, err
		}

		controller.reminderCount = &count
	}

	return *controller.reminderCount, nil
}

func (controller *Controller) IsReminderActive(index int) (bool, error) {
	result, err := controller.getData(code.MEM1_NEXTMEM + getOffset(index, blockItemsReminder, blockSizeReminder))
	return result != 0xffff, err
}

func (controller *Controller) GetReminderNext(index int) (time.Time, error) {
	result, err := controller.getData(code.MEM1_NEXTMEM + getOffset(index, blockItemsReminder, blockSizeReminder))
	if err != nil {
		return time.Now(), err
	}
	nextReminder := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	nextReminder = nextReminder.AddDate(0, 0, result)
	return nextReminder, nil
}

func (controller *Controller) GetReminderIsRepeating(index int) (bool, error) {
	return controller.getDataBool(code.MEM1_REPEAT + getOffset(index, blockItemsReminder, blockSizeReminder))
}

func (controller *Controller) GetReminderPeriod(index int) (int, error) {
	return controller.getData(code.MEM1_DAYS + getOffset(index, blockItemsReminder, blockSizeReminder))
}

func (controller *Controller) GetReminderText(index int) (string, error) {
	return controller.getDataText(code.MEM1_TEXT01 + getOffset(index, blockItemsReminder, blockSizeReminder))
}

// endregion

// region Maintenance

func (controller *Controller) IsMaintenanceActive(index int) (bool, error) {
	return controller.getDataBool(code.MAINTENANCE_ISACTIVE + getOffset(index, 1, 2))
}

func (controller *Controller) GetMaintenanceDuration(index int) (int, error) {
	return controller.getDataMultiplier(code.MAINTENANCE_TIMEOUT+getOffset(index, 1, 27), 60)
}

func (controller *Controller) GetMaintenanceTimeLeft(index int) (int, error) {
	return controller.getDataMultiplier(code.MAINTENANCE_REMATINGTIME+getOffset(index, 1, 2), 60)
}
func (controller *Controller) GetMaintenanceText(index int) (string, error) {
	return controller.getDataText(code.MAINT_NAME + getOffset(index, 64, 1))
}

// endregion

// region Probes

func getSensorOffset(index int) int {
	return getOffset(index, 8, 24)
}

func (controller *Controller) GetSensorCount() (int, error) {

	if controller.sensorCount == nil {
		count, err := controller.getData(code.GETSENSORCOUNT)
		if err != nil {
			return 0, err
		}

		controller.sensorCount = &count
	}

	return *controller.sensorCount, nil
}

func (controller *Controller) GetSensorType(index int) (types.SensorType, error) {
	result, err := controller.getDataEnum(code.SENSORPARA1_SENSORTYPE+getSensorOffset(index), types.GetSensorType)
	return types.SensorType(result), err
}

func (controller *Controller) GetSensorMode(index int) (types.SensorMode, error) {
	result, err := controller.getDataEnum(code.SENSORPARA1_CAL1ADC+getSensorOffset(index), func(value int) string {
		return types.GetSensorMode(value >> 12)
	})
	return types.SensorMode(result), err
}

func (controller *Controller) GetSensorActive(index int) (bool, error) {
	props, err := controller.getData(code.SENSORPARA1_PROPS + getSensorOffset(index))
	return props&0x1 == 1, err
}

func (controller *Controller) GetProbeName(index int) (string, error) {
	return controller.getDataText(code.SENSOR1_NAME + getOffset(index, 32, 1))
}

func (controller *Controller) GetSensorNominalValue(index int, multiplier float64) (float64, error) {
	return controller.getDataFloat(code.SENSORPARA1_DESVALUE+getSensorOffset(index), multiplier)
}

func (controller *Controller) GetSensorAlarmDeviation(index int, multiplier float64) (float64, error) {
	return controller.getDataFloat(code.SENSORPARA1_ALARMDEVIATION+getSensorOffset(index), multiplier)
}

func (controller *Controller) GetSensorAlarmEnable(index int) (bool, error) {
	return controller.getDataBool(code.SENSORPARA1_ALARMMODE + getSensorOffset(index))
}

func (controller *Controller) GetSensorAlarm(index int) (types.CurrentState, error) {
	return controller.getDataCurrentStateConvert(code.SENSORPARA1_ACTSTATE+getOffset(index, 8, 8), func(i int) int {
		return i & 0x100
	})
}

func (controller *Controller) GetSensorValue(index int, multiplier float64) (float64, error) {
	return controller.getDataFloat(code.SENSORPARA1_ACTVALUE+getOffset(index, 8, 8), multiplier)
}

func (controller *Controller) GetProbeOperationHours(index int) (int, error) {
	return controller.getData(code.SENSORPARA1_OHM + getOffset(index, 8, 8))
}

func (controller *Controller) GetSensorFormat(index int) (int, error) {
	result, err := controller.getData(code.SENSORPARA1_DISPLAYMODE + getSensorOffset(index))
	return result & 0x0f, err
}

// endregion

// region LevelSensor

func (controller *Controller) GetLevelSenosrCount() (int, error) {
	if controller.levelSensorCount == nil {
		count, err := controller.getData(code.GETLEVELSENSORCOUNT)
		if err != nil {
			return 0, err
		}

		controller.levelSensorCount = &count
	}

	return *controller.levelSensorCount, nil
}

func (controller *Controller) GetLevelSensorMode(index int) (types.LevelSensorOperationMode, error) {
	result, err := controller.getDataEnum(code.LEVEL1_PROPS+getOffset(index, 3, 3), func(value int) string {

		return types.GetLevelSensorOperationMode(value >> 13)
	})
	return types.LevelSensorOperationMode(result), err
}

func (controller *Controller) GetLevelName(index int) (string, error) {
	return controller.getDataText(code.LEVEL_NAME + getOffset(index, 64, 1))
}

type LevelState struct {
	WaterMode types.WaterMode
	Drain     types.CurrentState
	Fill      types.CurrentState
	Alarm     types.CurrentState
}

func (controller *Controller) GetLevelSensorState(index int) (LevelState, error) {
	// 7654 3210
	// AFDW WWWR
	// A - Alarm
	// F - Fill
	// D - Drain
	// W - Water State
	// R - Reserverd
	var levelState LevelState
	state, err := controller.getData(code.LEVEL1_ACTSTATE + getOffset(index, 3, 1))
	state >>= 1
	levelState.WaterMode = types.GetWaterMode(state & 0xf)
	state >>= 4
	levelState.Drain = types.GetCurrentState(state & 0x1)
	state >>= 1
	levelState.Fill = types.GetCurrentState(state & 0x1)
	state >>= 1
	levelState.Alarm = types.GetCurrentState(state & 0x1)

	return levelState, err
}

func (controller *Controller) GetLevelSource1(index int) (int, error) {
	result, err := controller.getData(code.LEVEL1_SOURCES + getOffset(index, 3, 3))
	return result & 0xf, err

}

func (controller *Controller) GetLevelSource2(index int) (int, error) {
	result, err := controller.getData(code.LEVEL1_SOURCES + getOffset(index, 3, 3))
	result >>= 4
	return result & 0xf, err
}

type LevelInputState struct {
	Delayed   types.CurrentState
	Previous  types.CurrentState
	Undelayed types.CurrentState
}

func (controller *Controller) GetLevelSensorCurrentState(index int) (LevelInputState, error) {
	// 7654 3210
	// UPDR RRRR
	// U - undelayed
	// P-  Previous
	// D - Delayed
	// R - Reserverd
	var levelState LevelInputState
	state, err := controller.getData(code.LEVEL1_INPUT_STATE + getOffset(index, 4, 1))
	state >>= 5
	levelState.Delayed = types.GetCurrentState(state & 0x1)
	state >>= 1
	levelState.Previous = types.GetCurrentState(state & 0x1)
	state >>= 1
	levelState.Undelayed = types.GetCurrentState(state & 0x1)

	return levelState, err
}

// endregion

// region DigitalInput
func (controller *Controller) GetDigitalInputCount() (int, error) {

	if controller.digitalInputCount == nil {
		count, err := controller.getData(code.GETDIGITALINPUTCOUNT)
		if err != nil {
			return 0, nil
		}

		controller.digitalInputCount = &count
	}

	return *controller.digitalInputCount, nil
}

func (controller *Controller) GetDigitalInputFunction(index int) (types.DigitalInputFunction, error) {
	result, err := controller.getDataEnum(code.DIGITALINPUT1_FUNCTION+getOffset(index, 4, 1), types.GetDigitalInputFunction)
	return types.DigitalInputFunction(result), err
}

func (controller *Controller) GetDigitalInputState(index int) (types.CurrentState, error) {
	state, err := controller.getData(code.DIGITALINPUTSSTATE + getOffset(index, 4, 0))
	if err != nil {
		return types.CurrentStateOff, nil
	}
	switch index % 4 {
	case 0:
		return types.GetCurrentState(state % 0x1), nil
	case 1:
		return types.GetCurrentState(state % 0x2), nil
	case 2:
		return types.GetCurrentState(state % 0x4), nil
	case 3:
		return types.GetCurrentState(state % 0x8), nil
	}

	return types.CurrentStateOff, nil
}

// endregion

// region Timer

func (controller *Controller) GetTimerCount() (int, error) {

	if controller.timerCount == nil {
		count, err := controller.getData(code.GETTIMERCOUNT)
		if err != nil {
			return 0, nil
		}
		controller.timerCount = &count
	}

	return *controller.timerCount, nil
}

func (controller *Controller) GetTimerSettings(index int) (types.TimerSettings, error) {
	config, err := controller.getData(code.TIMER1_PROPS + getOffset(index, 12, 21))

	return types.TimerSettings{
		Mode:              types.GetTimerMode((config >> 7) & 0x7),
		SwitchingCount:    config >> 11,
		FeedPauseIfActive: (config & 0x10) != 0,
		DayMode:           types.GetDayMode((config >> 10) & 0x1),
	}, err
}

func (controller *Controller) GetDosingRate(index int) (int, error) {
	return controller.getData(code.TIMER1_RATEPERDOSING + getOffset(index, 12, 21))
}

// endregion

// region Light
func (controller *Controller) GetLightCount() (int, error) {
	if controller.lightCount == nil {
		count, err := controller.getData(code.GETILLUMINATIONCOUNT)
		if err != nil {
			return 0, nil
		}

		controller.lightCount = &count
	}

	return *controller.lightCount, nil
}

func (controller *Controller) GetIsLightActive(index int) (bool, error) {
	return controller.getDataBoolConvert(code.ILLUMINATION1_PROPS+getOffset(index, 8, 28), func(config int) int {
		return (config >> 4) & 0xF
	})
}
func (controller *Controller) GetIsLightDimmable(index int) (bool, error) {
	return controller.getDataBoolConvert(code.ILLUMINATION1_PROPS+getOffset(index, 8, 28), func(config int) int {
		return config & 0x8
	})
}
func (controller *Controller) GetLightOperationHours(index int) (int, error) {
	return controller.getData(code.ILLUMINATION1_OHM + getOffset(index, 8, 4))
}
func (controller *Controller) GetLightValue(index int) (float64, error) {
	return controller.getDataFloat(code.ILLUMINATION1_ACTPERCENT+getOffset(index, 8, 4), 1)
}
func (controller *Controller) GetLightName(index int) (string, error) {
	return controller.getDataText(code.ILLUMINATION1_NAME + getOffset(index, 32, 1))
}

// endregion

// region Pumps

func (controller *Controller) GetCurrentPumpCount() int {

	if controller.pumpCount == nil {
		count, _ := controller.getData(code.GETCURRENTPUMPCOUNT)
		controller.pumpCount = &count
	}

	return *controller.pumpCount
}

func (controller *Controller) GetIsCurrentPumpAssigned(index int) (bool, error) {
	group1Mask, err := controller.getData(code.CURRENTCONTROL_GROUP1PUMPCOUNT)
	if err != nil {
		return false, err
	}
	i := uint(index)
	if (group1Mask >> i & 0x1) == 1 {
		return true, nil
	}

	group2Mask, err := controller.getData(code.CURRENTCONTROL_GROUP2PUMPCOUNT)
	if err != nil {
		return false, err
	}

	return (group2Mask >> i & 0x1) == 1, nil
}

func (controller *Controller) GetCurrentPumpValue(index int) (int, error) {
	return controller.getData(code.CURRENTPUMP1_ACTPERCENT + getOffset(index, 4, 1))
}

// endregion

// region ProgrammableLogic
func (controller *Controller) GetProgrammableLogicCount() (int, error) {

	if controller.logicCount == nil {
		count, err := controller.getData(code.GETPROGLOGICCOUNT)
		if err != nil {
			return 0, nil
		}
		controller.logicCount = &count
	}

	return *controller.logicCount, nil
}

func (controller *Controller) GetProgramLogicInput(input, index int) (types.PortMode, error) {
	mode, err := controller.getData(code.PROGLOGIC1_INPUT1 + input + getOffset(index, 8, 4))
	var portMode types.PortMode
	if err != nil {
		return portMode, err
	}

	// mode format
	// 1234 1234 1234 1234
	// RRRR RRPP PPPT TTTT
	// R = Reserved
	// P = Port Number
	// T = Type
	mode >>= 6
	portMode.Port = (mode & 0x1F) + 1
	mode >>= 5
	portMode.DeviceMode = types.GetDeviceMode(mode & 0x1F)
	portMode.IsProbe = getIsProbe(portMode.DeviceMode)
	return portMode, nil
}

func (controller *Controller) GetProgramLogicFunction(index int) (types.LogicFunction, error) {
	mode, err := controller.getData(code.PROGLOGIC1_FUNCTION + getOffset(index, 8, 4))

	// mode format
	// 1234 1234
	// RRMM MMMM
	// R = Reserved
	// P = Port Number
	// T = Type
	var function types.LogicFunction
	if err != nil {
		return function, err
	}
	function.Invert2 = (mode & 0x1) == 1
	mode >>= 1
	function.Invert1 = (mode & 0x1) == 1
	mode >>= 1
	function.LogicMode = types.GetLogicMode(mode & 0x3F)

	return function, nil
}

// endregion

// region sport

func (controller *Controller) GetSPortCount() (int, error) {

	if controller.sPortCount == nil {
		count, err := controller.getData(code.GETSWITCHCOUNT)
		if err != nil {
			return 0, err
		}

		controller.sPortCount = &count
	}

	return *controller.sPortCount, nil
}

func getIsProbe(mode types.DeviceMode) bool {
	return mode == types.DeviceModeDecrease ||
		mode == types.DeviceModeIncrease ||
		mode == types.DeviceModeSubstrate ||
		mode == types.DeviceModeProbeAlarm
}

func (controller *Controller) GetSPortFunction(index int) (types.PortMode, error) {
	mode, err := controller.getData(code.SWITCHPLUG1_FUNCTION + getOffset(index, 24, 1))

	// mode format
	// 1234 1234 1234 1234
	// IBBB BBPP PPPT TTTT
	// I = Invert
	// B = Blackout time
	// P = Port Number
	// T = Type
	var portMode types.PortMode

	if err != nil {
		return portMode, err
	}

	portMode.Invert = (mode & 0x01) != 0
	mode >>= 1
	portMode.BlackOut = mode & 0x1F
	mode >>= 5
	portMode.Port = (mode & 0x1F) + 1
	mode >>= 5
	portMode.DeviceMode = types.GetDeviceMode(mode & 0x1F)
	portMode.IsProbe = getIsProbe(portMode.DeviceMode)
	return portMode, nil
}

func (controller *Controller) GetSPortValue(index int) (types.CurrentState, error) {
	return controller.getDataCurrentState(code.SP1_STATE + getOffset(index, 24, 1))
}

func (controller *Controller) GetSPortName(index int) (string, error) {
	return controller.getDataText(code.SWITCHPLUG1_NAME + getOffset(index, 64, 1))
}

// endregion

// region lport

func (controller *Controller) GetLPortCount() (int, error) {

	if controller.lPortCount == nil {
		count, err := controller.getData(code.GETONTETOTENVINTCOUNT)
		if err != nil {
			return 0, err
		}
		controller.lPortCount = &count
	}

	return *controller.lPortCount, nil
}

func (controller *Controller) GetLPortFunction(index int) (types.PortMode, error) {
	mode, err := controller.getData(code.L1TO10VINT1_FUNCTION + getOffset(index, 10, 3))

	// mode format
	// 1234 1234
	// PPPT TTTT
	// I = Invert
	// B = Blackout time
	// P = Port Number
	// T = Type
	var portMode types.PortMode
	if err != nil {
		return portMode, err
	}

	portMode.BlackOut = mode & 0x003F
	mode >>= 6
	portMode.Port = (mode & 0x001F) + 1
	mode >>= 5
	portMode.DeviceMode = types.LPortModeToSocketType(mode & 0x003F)
	portMode.IsProbe = getIsProbe(portMode.DeviceMode)
	return portMode, nil
}

const (
	LValueMin = 18.0
	LValueMax = 255.00
)

func (controller *Controller) GetLPortValue(index int) (float64, error) {

	value, err := controller.getData(code.L1TO10VINT1_PWMVALUE + getOffset(index, 10, 1))
	if err != nil {
		return 0, nil
	}

	if value < LValueMin {
		return 0, nil
	} else {
		return (float64(value) - LValueMin) / (LValueMax - LValueMin) * 100.0, nil
	}
}

// endregion

func (controller *Controller) FeedPause(index int, activate bool) error {
	command := (index << 16) | (0 << 8) | SfFeedPause
	if activate {
		command = (index << 16) | (0xFF << 8) | SfFeedPause
	}

	err := controller.p.SendData(code.INVOKESPECIALFUNCTION, command)
	if err != nil {
		log.Errorf(err, "FeedPause: %s", err.Error())
		return err
	}

	return nil
}

func (controller *Controller) ClearReminder(reminder int) error {
	offset := getOffset(reminder, blockItemsReminder, blockSizeReminder)
	err := controller.p.SendData(code.MEM1_NEXTMEM+offset, 0xFFFF)
	if err != nil {
		log.Errorf(err, "ClearReminder: %s", err.Error())
		return err
	}

	return nil
}

func (controller *Controller) ResetReminder(reminder int, period int) error {
	nextReminder := time.Now().AddDate(0, 0, period)
	span := nextReminder.Sub(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	data := int(span.Hours() / 24)

	offset := getOffset(reminder, blockItemsReminder, blockSizeReminder)
	err := controller.p.SendData(code.MEM1_NEXTMEM+offset, data)
	if err != nil {
		log.Errorf(err, "ClearReminder: %s", err.Error())
		return err
	}

	return nil
}

func (controller *Controller) ClearLevelAlarm(index int) error {
	err := controller.p.SendData(code.LEVEL1_ACTSTATE+getOffset(index, 3, 1), 0)

	if err != nil {
		log.Errorf(err, "ClearLevelAlarm: %s", err.Error())
		return err
	}

	return nil
}

func (controller *Controller) Thunderstorm(duration int) error {
	command := (duration << 8) | SfThunderstorm
	err := controller.p.SendData(code.INVOKESPECIALFUNCTION, command)
	if err != nil {
		log.Errorf(err, "Thunderstorm: %s", err.Error())
		return err
	}

	return nil
}

func (controller *Controller) WaterChange(index int) error {
	command := index<<16 | (0xFF << 8) | SfWaterChange
	err := controller.p.SendData(code.INVOKESPECIALFUNCTION, command)
	if err != nil {
		log.Errorf(err, "WaterChange: %s", err.Error())
		return err
	}

	return nil
}

func (controller *Controller) Maintenance(activate bool, index int) error {
	command := (index << 16) | (0 << 8) | SfMaintenance
	if activate {
		command = (index << 16) | (0xFF << 8) | SfMaintenance
	}

	err := controller.p.SendData(code.INVOKESPECIALFUNCTION, command)

	if err != nil {
		log.Errorf(err, "Maintenance: %s", err.Error())
		return err
	}

	return nil
}
