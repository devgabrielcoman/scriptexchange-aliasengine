package ingester

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScriptIngester_Process(t *testing.T) {
	var ingester = ScriptIngester{Alias: "my-script", Path: "/path/to/my-script.sh", CurrentTime: 123}
	t.Run("it should return a new script given correct input", func(t *testing.T) {
		var content = "echo \"Good\""
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "my-script",
				Content:    "echo \"Good\"",
				Path:       ".scripts",
				Comments:   []string{},
				PathOnDisk: ".scripts",
				Type:       models.ScriptType(models.Script),
				Date:       123,
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})
}
