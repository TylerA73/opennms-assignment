package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tylera73/opennms-assignment/stats"
)

func main() {
	router := gin.Default()

	stats.RegisterRoutes(router)

	err := router.Run(":4000")
	if err != nil {
		fmt.Println(fmt.Errorf("could not start service: %v", err))
	}
}
