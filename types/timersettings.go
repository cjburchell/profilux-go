package types

type TimerSettings struct {
	FeedPauseIfActive bool
	Mode              TimerMode
	DayMode           DayMode
	SwitchingCount    int
}
