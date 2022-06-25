package ingester

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZSHHistoryIngester_Process(t *testing.T) {
	var ingester = ZSHHistoryIngester{Path: "/path/to/.zsh_history"}

	t.Run("it should return empty array of items given empty zsh history", func(t *testing.T) {
		var content = ""
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return empty array of items given non-zsh history", func(t *testing.T) {
		var content = "ls -all\necho \"Hello World\""
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an array of bash commands, in the order they are read", func(t *testing.T) {
		var content = ": 1656006100:0;cd\n: 1656006110:0;vi .zsh_history\n: 1656053625:0;cp .zsh_history /Users/gabriel.coman/Desktop/my_history"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "cp .zsh_history /Users/gabriel.coman/Desktop/my_history",
				Content:    "cp .zsh_history /Users/gabriel.coman/Desktop/my_history",
				Path:       "/path/to/.zsh_history",
				PathOnDisk: "/path/to/.zsh_history",
				Comments:   []string{},
				Type:       models.ScriptType(models.History),
				Date:       1656053625,
				StartLine:  0,
				EndLine:    0,
			},
			{
				Name:       "vi .zsh_history",
				Content:    "vi .zsh_history",
				Path:       "/path/to/.zsh_history",
				PathOnDisk: "/path/to/.zsh_history",
				Comments:   []string{},
				Type:       models.ScriptType(models.History),
				Date:       1656006110,
				StartLine:  1,
				EndLine:    1,
			},
			{
				Name:       "cd",
				Content:    "cd",
				Path:       "/path/to/.zsh_history",
				PathOnDisk: "/path/to/.zsh_history",
				Comments:   []string{},
				Type:       models.ScriptType(models.History),
				Date:       1656006100,
				StartLine:  2,
				EndLine:    2,
			},
		}
		assert.Equal(t, expected, result)
	})
}
