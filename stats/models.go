package stats

type Stats struct {
	ID           string
	MachineID    int
	CPUTemp      *float64
	FanSpeed     *float64
	HDDSpace     *float64
	InternalTemp *float64
	LastLoggedIn string
	SysTime      string
}
