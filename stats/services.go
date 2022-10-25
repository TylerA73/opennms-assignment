package stats

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tylera73/opennms-assignment/datastore"
)

func CreateMachineStats(req *PostMachineStatsRequest) error {
	if req.MachineID == 0 || req.Stats == nil || req.LastLoggedIn == "" || req.SysTime == "" {
		return errors.New("machineID, stats, lastLoggedIn, and sysTime cannot be empty")
	}

	stats := &Stats{
		ID:           uuid.New().String(),
		MachineID:    req.MachineID,
		CPUTemp:      req.Stats.CPUTemp,
		FanSpeed:     req.Stats.FanSpeed,
		HDDSpace:     req.Stats.HDDSpace,
		InternalTemp: req.Stats.InternalTemp,
		LastLoggedIn: req.LastLoggedIn,
		SysTime:      req.SysTime,
	}

	datastore.SaveStats(stats)

	return nil
}

func FetchMachineStats() []*Stats {
	values := datastore.GetStats()
	stats := make([]*Stats, len(values))
	for i, v := range values {
		stats[i] = v.(*Stats)
	}

	return stats
}
