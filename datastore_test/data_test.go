package datastore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tylera73/opennms-assignment/datastore"
	"github.com/tylera73/opennms-assignment/stats"
)

func TestSaveStats(t *testing.T) {
	t.Run("should start with 0 records", func(t *testing.T) {
		assert.Equal(t, 0, len(datastore.Data))
	})

	t.Run("should successfully add a record to the datastore", func(t *testing.T) {
		datastore.SaveStats(stats.Stats{})
		assert.Equal(t, 1, len(datastore.Data))
	})
}
