package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max(t *testing.T) {
	t.Run("should return the max between a negative and positive number", func(t *testing.T) {
		var result = Max(-5, 5)
		var expected = 5
		assert.Equal(t, expected, result)
	})

	t.Run("should return the max between two positive numbers", func(t *testing.T) {
		var result = Max(8, 5)
		var expected = 8
		assert.Equal(t, expected, result)
	})

	t.Run("should return the max between two negative numbers", func(t *testing.T) {
		var result = Max(-8, -5)
		var expected = -5
		assert.Equal(t, expected, result)
	})

	t.Run("should return the same number for identical inputs", func(t *testing.T) {
		var result = Max(8, 8)
		var expected = 8
		assert.Equal(t, expected, result)
	})
}

func Test_Min(t *testing.T) {
	t.Run("should return the min between a negative and positive number", func(t *testing.T) {
		var result = Min(-5, 5)
		var expected = -5
		assert.Equal(t, expected, result)
	})

	t.Run("should return the min between two positive numbers", func(t *testing.T) {
		var result = Min(8, 5)
		var expected = 5
		assert.Equal(t, expected, result)
	})

	t.Run("should return the min between two negative numbers", func(t *testing.T) {
		var result = Min(-8, -5)
		var expected = -8
		assert.Equal(t, expected, result)
	})

	t.Run("should return the same number for identical inputs", func(t *testing.T) {
		var result = Min(8, 8)
		var expected = 8
		assert.Equal(t, expected, result)
	})
}
