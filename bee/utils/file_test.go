package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsHttpUrl(t *testing.T) {
	t.Run("should return true for HTTP url", func(t *testing.T) {
		var path = "http://my.examole.com"
		var result = IsHttpUrl(path)
		assert.True(t, result)
	})

	t.Run("should return true for HTTPS url", func(t *testing.T) {
		var path = "https://my.examole.com"
		var result = IsHttpUrl(path)
		assert.True(t, result)
	})

	t.Run("should return false for file path", func(t *testing.T) {
		var path = "/usr/path/to/file"
		var result = IsHttpUrl(path)
		assert.False(t, result)
	})

	t.Run("should return false empty", func(t *testing.T) {
		var path = ""
		var result = IsHttpUrl(path)
		assert.False(t, result)
	})

	t.Run("should return false for garbage", func(t *testing.T) {
		var path = "\\////safas"
		var result = IsHttpUrl(path)
		assert.False(t, result)
	})
}

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

	t.Run("should return a file name given valid path", func(t *testing.T) {
		var path = "https://test.com/abc?query=abc"
		var result = FileName(path)
		var expected = "abc"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a file name given valid path", func(t *testing.T) {
		var path = "http://test.com/test/file.sh"
		var result = FileName(path)
		var expected = "file.sh"
		assert.Equal(t, expected, result)
	})
}
