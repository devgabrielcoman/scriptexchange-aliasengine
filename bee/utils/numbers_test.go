package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LenientAtoi(t *testing.T) {
	t.Run("should return a valid number given a valid string representing a positive int", func(t *testing.T) {
		var result = LenientAtoi("123")
		var expected = 123
		assert.Equal(t, expected, result)
	})

	t.Run("should return a valid number given a valid string representing a negative int", func(t *testing.T) {
		var result = LenientAtoi("-123")
		var expected = -123
		assert.Equal(t, expected, result)
	})

	t.Run("should return a valid number given a valid string representing zero", func(t *testing.T) {
		var result = LenientAtoi("0")
		var expected = 0
		assert.Equal(t, expected, result)
	})

	t.Run("should return zero in case string is not a number", func(t *testing.T) {
		var result = LenientAtoi("abc")
		var expected = 0
		assert.Equal(t, expected, result)
	})
}

func Test_LenientAtoi64(t *testing.T) {
	t.Run("should return a valid number given a valid string representing a positive int", func(t *testing.T) {
		var result = LenientAtoi64("123")
		var expected = int64(123)
		assert.Equal(t, expected, result)
	})

	t.Run("should return a valid number given a valid string representing a negative int", func(t *testing.T) {
		var result = LenientAtoi64("-123")
		var expected = int64(-123)
		assert.Equal(t, expected, result)
	})

	t.Run("should return a valid number given a valid string representing zero", func(t *testing.T) {
		var result = LenientAtoi64("0")
		var expected = int64(0)
		assert.Equal(t, expected, result)
	})

	t.Run("should return zero in case string is not a number", func(t *testing.T) {
		var result = LenientAtoi64("abc")
		var expected = int64(0)
		assert.Equal(t, expected, result)
	})
}
