package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DateFormat(t *testing.T) {
	t.Run("should return a formatted date for 0", func(t *testing.T) {
		var timestamp int64 = 0
		var result = DateFormat(timestamp)
		var expected = "1970-01-01 01:00:00"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a formatted date for a valid unix timestamp", func(t *testing.T) {
		var timestamp int64 = 1656153865
		var result = DateFormat(timestamp)
		var expected = "2022-06-25 11:44:25"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a formatted date for a negative timestamp", func(t *testing.T) {
		var timestamp int64 = -5000
		var result = DateFormat(timestamp)
		var expected = "1969-12-31 23:36:40"
		assert.Equal(t, expected, result)
	})
}
