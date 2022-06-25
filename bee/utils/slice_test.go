package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	t.Run("should return an empty slice given an empty slice", func(t *testing.T) {
		var data = []string{}
		Reverse(data)
		var expected = []string{}
		assert.Equal(t, expected, data)
	})

	t.Run("should return a reversed slice", func(t *testing.T) {
		var data = []int{1, 2, 3}
		Reverse(data)
		var expected = []int{3, 2, 1}
		assert.Equal(t, expected, data)
	})
}
