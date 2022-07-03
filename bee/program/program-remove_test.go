package program

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeSource_WithName(t *testing.T) {
	program := RemoveProgram{Name: ".test.sh"}
	t.Run("should return empty given empty input", func(t *testing.T) {
		sources := []models.SourceFile{}
		result := program.removeSource(sources)
		expected := []models.SourceFile{}
		assert.Equal(t, expected, result)
	})

	t.Run("should not remove anything if input does not contain source to be removed", func(t *testing.T) {
		sources := []models.SourceFile{
			{
				Name: "test",
				Path: "/path/to/test",
				Type: models.SourceType(models.Command),
			},
		}
		result := program.removeSource(sources)
		expected := []models.SourceFile{
			{
				Name: "test",
				Path: "/path/to/test",
				Type: models.SourceType(models.Command),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("should remove item if input contains source to be removed", func(t *testing.T) {
		sources := []models.SourceFile{
			{
				Name: ".test.sh",
				Path: "/path/to/.test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		result := program.removeSource(sources)
		expected := []models.SourceFile{}
		assert.Equal(t, expected, result)
	})
}

func Test_removeSource_WithPath(t *testing.T) {
	program := RemoveProgram{Name: "/path/to/.test.sh"}
	t.Run("should return empty given empty input", func(t *testing.T) {
		sources := []models.SourceFile{}
		result := program.removeSource(sources)
		expected := []models.SourceFile{}
		assert.Equal(t, expected, result)
	})

	t.Run("should not remove anything if input does not contain source to be removed", func(t *testing.T) {
		sources := []models.SourceFile{
			{
				Name: "test",
				Path: "/path/to/test",
				Type: models.SourceType(models.Command),
			},
		}
		result := program.removeSource(sources)
		expected := []models.SourceFile{
			{
				Name: "test",
				Path: "/path/to/test",
				Type: models.SourceType(models.Command),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("should remove item if input contains source to be removed", func(t *testing.T) {
		sources := []models.SourceFile{
			{
				Name: ".test.sh",
				Path: "/path/to/.test.sh",
				Type: models.SourceType(models.Command),
			},
		}
		result := program.removeSource(sources)
		expected := []models.SourceFile{}
		assert.Equal(t, expected, result)
	})
}
