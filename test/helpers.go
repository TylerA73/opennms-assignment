package test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tylera73/opennms-assignment/datastore"
	"github.com/tylera73/opennms-assignment/stats"
)

func CleanUp() {
	datastore.Data = make([]interface{}, 0)
}

func CreateStats() {
	datastore.SaveStats(&stats.Stats{
		ID:        fmt.Sprintf("test_id_%d", len(datastore.Data)+1),
		MachineID: 12345,
	})
}

func SetupRouter() *gin.Engine {
	return gin.Default()
}
