package stats_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tylera73/opennms-assignment/stats"
	"github.com/tylera73/opennms-assignment/test"
)

func TestPostMachineStats(t *testing.T) {
	tt := map[string]struct {
		want       string
		give       stats.PostMachineStatsRequest
		statusCode int
	}{
		"should receive a 201": {
			want: `{"message":"Machine stats successfully created"}`,
			give: stats.PostMachineStatsRequest{
				MachineID: 12345,
				Stats: &stats.MachineStats{
					CPUTemp:  &cpuTemp,
					FanSpeed: &fanSpeed,
					HDDSpace: &hddSpace,
				},
				LastLoggedIn: "admin/Paul",
				SysTime:      "2022-04-23T18:25:43.511Z",
			},
			statusCode: http.StatusCreated,
		},
		"should receive a 400 with missing required fields": {
			want: `{"message":"Could not create the machine stats: machineID, stats, lastLoggedIn, and sysTime cannot be empty"}`,
			give: stats.PostMachineStatsRequest{
				MachineID:    12345,
				LastLoggedIn: "admin/Paul",
				SysTime:      "2022-04-23T18:25:43.511Z",
			},
			statusCode: http.StatusBadRequest,
		},
	}

	r := test.SetupRouter()
	r.POST("/stats", stats.PostMachineStats)

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(test.CleanUp)

			body, _ := json.Marshal(tc.give)

			req, err := http.NewRequest(http.MethodPost, "/stats", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			response, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(tc.give)

			assert.Equal(t, tc.want, string(response))
			assert.Equal(t, tc.statusCode, rr.Code)
		})
	}
}

func TestGetMachineStats(t *testing.T) {
	tt := map[string]struct {
		want            string
		statusCode      int
		numberOfRecords int
	}{
		"should receive empty array": {
			want:            "[]",
			statusCode:      http.StatusOK,
			numberOfRecords: 0,
		},
		"should receive an array with stats": {
			want:            `[{"id":"test_id_1","machineId":12345,"stats":{"cpuTemp":null,"fanSpeed":null,"HDDSpace":null},"lastLoggedIn":"","sysTime":""}]`,
			statusCode:      http.StatusOK,
			numberOfRecords: 1,
		},
	}

	r := test.SetupRouter()
	r.GET("/", stats.GetMachineStats)

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(test.CleanUp)

			for i := 0; i < tc.numberOfRecords; i++ {
				test.CreateStats()
			}

			req, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			response, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.want, string(response))
			assert.Equal(t, tc.statusCode, rr.Code)
		})
	}
}
