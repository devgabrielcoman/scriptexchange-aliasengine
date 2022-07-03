package program

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeSources(t *testing.T) {
	program := SourceProgram{Path: "/path/to/new/sources.json"}
	t.Run("should return empty given empty inputs", func(t *testing.T) {
		old := []models.SourceFile{}
		new := []models.SourceFile{}
		result := program.mergeSources(old, new)
		expected := []models.SourceFile{}
		assert.Equal(t, expected, result)
	})

	t.Run("should return valid input given empty new slice", func(t *testing.T) {
		old := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		new := []models.SourceFile{}
		result := program.mergeSources(old, new)
		expected := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("should return valid input given non empty distinct new slice", func(t *testing.T) {
		old := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		new := []models.SourceFile{
			{
				Path: "/path/to/.test2.sh",
				Name: ".test2.sh",
				Type: models.SourceType(models.Command),
			},
		}
		result := program.mergeSources(old, new)
		expected := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
			{
				Path: "/path/to/.test2.sh",
				Name: ".test2.sh",
				Type: models.SourceType(models.Command),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("should return valid input given non empty duplicated new slice", func(t *testing.T) {
		old := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		new := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		result := program.mergeSources(old, new)
		expected := []models.SourceFile{
			{
				Path: "/path/to/.test.sh",
				Name: ".test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		assert.Equal(t, expected, result)
	})
}
