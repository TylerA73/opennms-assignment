package stats_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tylera73/opennms-assignment/datastore"
	"github.com/tylera73/opennms-assignment/stats"
	"github.com/tylera73/opennms-assignment/test"
)

var (
	cpuTemp      float64 = 75
	fanSpeed     float64 = 100
	hddSpace     float64 = 500
	internalTemp float64 = 20
)

func TestCreateMachineStats(t *testing.T) {
	tt := map[string]struct {
		give            *stats.PostMachineStatsRequest
		want            error
		numberOfRecords int
	}{
		"should create new machine stats": {
			give: &stats.PostMachineStatsRequest{
				MachineID: 12345,
				Stats: &stats.MachineStats{
					CPUTemp:      &cpuTemp,
					FanSpeed:     &fanSpeed,
					HDDSpace:     &hddSpace,
					InternalTemp: &internalTemp,
				},
				LastLoggedIn: "admin/User",
				SysTime:      time.Now().Format("2006-01-02"),
			},
			want:            nil,
			numberOfRecords: 1,
		},
		"should not create new machine stats": {
			give:            &stats.PostMachineStatsRequest{},
			want:            errors.New("machineID, stats, lastLoggedIn, and sysTime cannot be empty"),
			numberOfRecords: 0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(test.CleanUp)

			err := stats.CreateMachineStats(tc.give)
			assert.Equal(t, tc.want, err)
			assert.Equal(t, tc.numberOfRecords, len(datastore.Data))
		})
	}
}

func TestFetchMachineStats(t *testing.T) {
	t.Cleanup(test.CleanUp)

	tt := map[string]struct {
		want            []*stats.Stats
		numberOfRecords int
	}{
		"should fetch an empty array of stats": {
			want:            make([]*stats.Stats, 0),
			numberOfRecords: 0,
		},
		"should fetch an array containing one set of stats": {
			want: []*stats.Stats{
				{
					ID:        "test_id_1",
					MachineID: 12345,
				},
			},
			numberOfRecords: 1,
		},
		"should fetch an array containing multiple set of stats": {
			want: []*stats.Stats{
				{
					ID:        "test_id_1",
					MachineID: 12345,
				},
				{
					ID:        "test_id_2",
					MachineID: 12345,
				},
			},
			numberOfRecords: 2,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// cleanup the previous test and set up the next
			t.Cleanup(test.CleanUp)
			for i := 0; i < tc.numberOfRecords; i++ {
				test.CreateStats()
			}

			s := stats.FetchMachineStats()
			assert.Equal(t, tc.numberOfRecords, len(s))
			for i, stat := range s {
				assert.Equal(t, tc.want[i].ID, stat.ID)
			}
		})
	}
}
