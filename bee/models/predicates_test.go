package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FilterByPath(t *testing.T) {
	t.Run("should return an empty slice given an empty input", func(t *testing.T) {
		path := "/my_path"
		items := []IndexItem{}
		result := FilterByPath(items, path)
		expected := []IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("should return a slice of IndexItem containing matches that fit the path", func(t *testing.T) {
		path := "/my_path"

		item1 := IndexItem{
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}
		item2 := IndexItem{
			Name:       "Two",
			Path:       "/my_other_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}

		items := []IndexItem{item1, item2}
		result := FilterByPath(items, path)
		expected := []IndexItem{item1}
		assert.Equal(t, expected, result)
	})

	t.Run("should return an empty slice given no matching inputs", func(t *testing.T) {
		path := "/my_non_matching_path"

		item1 := IndexItem{
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}
		item2 := IndexItem{
			Name:       "Two",
			Path:       "/my_other_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}

		items := []IndexItem{item1, item2}
		result := FilterByPath(items, path)
		expected := []IndexItem{}
		assert.Equal(t, expected, result)
	})
}

func Test_UniqueItems(t *testing.T) {
	t.Run("should return an empty slice given an empty input", func(t *testing.T) {
		input := []IndexItem{}
		result := UniqueItems(input)
		expected := []IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("should return same slice in case there are no duplicates", func(t *testing.T) {
		item1 := IndexItem{
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}
		item2 := IndexItem{
			Name:       "Two",
			Path:       "/my_other_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}

		items := []IndexItem{item1, item2}
		expected := []IndexItem{item1, item2}
		result := UniqueItems(items)
		assert.Equal(t, expected, result)
	})

	t.Run("should return a slice of unique items given a slice with duplicates", func(t *testing.T) {
		item1 := IndexItem{
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}
		item2 := IndexItem{
			Name:       "Two",
			Path:       "/my_other_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}

		items := []IndexItem{item1, item1, item2}
		expected := []IndexItem{item1, item2}
		result := UniqueItems(items)
		assert.Equal(t, expected, result)
	})
}
