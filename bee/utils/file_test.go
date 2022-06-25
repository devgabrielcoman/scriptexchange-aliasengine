package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FileName(t *testing.T) {
	t.Run("should return empty dot given empty input", func(t *testing.T) {
		var path = ""
		var result = FileName(path)
		var expected = "."
		assert.Equal(t, expected, result)
	})

	t.Run("should return the same input given garbage input", func(t *testing.T) {
		var path = "\\\\//aasaas;;;s;a"
		var result = FileName(path)
		var expected = "aasaas;;;s;a"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a file name given valid path", func(t *testing.T) {
		var path = "/Users/test.test/my/file.text"
		var result = FileName(path)
		var expected = "file.text"
		assert.Equal(t, expected, result)
	})
}