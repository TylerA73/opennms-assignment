package stats

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	statsGroup := router.Group("/stats")
	statsGroup.POST("", PostMachineStats)
	statsGroup.GET("", GetMachineStats)
}
