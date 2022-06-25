package ingester

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBashHistoryIngester_Process(t *testing.T) {
	var ingester = BashHistoryIngester{Path: "/path/to/.bash_history"}

	t.Run("it should return empty array of items given empty bash history", func(t *testing.T) {
		var content = ""
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an array of bash commands, in the order they are read", func(t *testing.T) {
		var content = "ls -all\necho \"Hello World\""
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "ls -all",
				Content:    "ls -all",
				Path:       "/path/to/.bash_history",
				PathOnDisk: "/path/to/.bash_history",
				Comments:   []string{},
				Type:       models.ScriptType(models.History),
				Date:       0,
			},
			{
				Name:       "echo \"Hello World\"",
				Content:    "echo \"Hello World\"",
				Path:       "/path/to/.bash_history",
				PathOnDisk: "/path/to/.bash_history",
				Comments:   []string{},
				Type:       models.ScriptType(models.History),
				Date:       0,
			},
		}
		assert.Equal(t, expected, result)
	})
}
