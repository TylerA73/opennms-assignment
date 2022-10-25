package stats

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostMachineStatsRequest struct {
	MachineID    int           `json:"machineId"`
	Stats        *MachineStats `json:"stats"`
	LastLoggedIn string        `json:"lastLoggedIn"`
	SysTime      string        `json:"sysTime"`
}

type MachineStats struct {
	CPUTemp      *float64 `json:"cpuTemp"`
	FanSpeed     *float64 `json:"fanSpeed"`
	HDDSpace     *float64 `json:"HDDSpace"`
	InternalTemp *float64 `json:"internalTemp,omitempty"`
}

type PostMachineStatsResponse struct {
	Message string `json:"message"`
}

func PostMachineStats(ctx *gin.Context) {
	req := new(PostMachineStatsRequest)

	if err := ctx.Bind(req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, PostMachineStatsResponse{
			Message: "Could not bind the request body",
		})
		return
	}

	err := CreateMachineStats(req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, PostMachineStatsResponse{
			Message: fmt.Sprintf("Could not create the machine stats: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, PostMachineStatsResponse{
		Message: "Machine stats successfully created",
	})
}

type GetMachineStatsResponse struct {
	ID           string       `json:"id"`
	MachineID    int          `json:"machineId"`
	Stats        MachineStats `json:"stats"`
	LastLoggedIn string       `json:"lastLoggedIn"`
	SysTime      string       `json:"sysTime"`
}

func GetMachineStats(ctx *gin.Context) {
	resp := make([]GetMachineStatsResponse, 0)
	stats := FetchMachineStats()

	for _, s := range stats {
		resp = append(resp, GetMachineStatsResponse{
			ID:        s.ID,
			MachineID: s.MachineID,
			Stats: MachineStats{
				CPUTemp:      s.CPUTemp,
				FanSpeed:     s.FanSpeed,
				HDDSpace:     s.HDDSpace,
				InternalTemp: s.InternalTemp,
			},
			LastLoggedIn: s.LastLoggedIn,
			SysTime:      s.SysTime,
		})
	}

	ctx.JSON(http.StatusOK, resp)

}
