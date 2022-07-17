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
		path := "/full/my_path"

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

func Test_UniquePaths(t *testing.T) {
	t.Run("should return empty given empty input", func(t *testing.T) {
		input := []IndexItem{}
		result := UniquePaths(input)
		expected := []Pair{}
		assert.Equal(t, expected, result)
	})

	t.Run("should return a slice of unique pairs containing paths and paths on disk", func(t *testing.T) {
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
		input := []IndexItem{item1, item2}
		result := UniquePaths(input)
		expected := []Pair{
			{
				A: "/my_path",
				B: "/full/my_path",
			},
			{
				A: "/my_other_path",
				B: "/full/my_other_path",
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("should return a slice of unique pairs with paths even if only paths on disk are different", func(t *testing.T) {
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
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       ScriptType(Alias),
			Date:       123,
		}
		input := []IndexItem{item1, item2}
		result := UniquePaths(input)
		expected := []Pair{
			{
				A: "/my_path",
				B: "/full/my_path",
			},
			{
				A: "/my_path",
				B: "/full/my_other_path",
			},
		}
		assert.Equal(t, expected, result)
	})
}

func Test_UniqueSources(t *testing.T) {
	t.Run("should return empty slice given empty input", func(t *testing.T) {
		input := []SourceFile{}
		result := UniqueSources(input)
		expected := []SourceFile{}
		assert.Equal(t, expected, result)
	})

	t.Run("should return only unique sources by path and name", func(t *testing.T) {
		source1 := SourceFile{
			Name: "my_source",
			Path: "path/to/my_source",
			Type: SourceType(Command),
		}
		source2 := SourceFile{
			Name: "my_source2",
			Path: "path/to/other/my_source2",
			Type: SourceType(Command),
		}
		input := []SourceFile{source1, source2, source2, source1, source2}
		result := UniqueSources(input)
		expected := []SourceFile{source1, source2}
		assert.Equal(t, expected, result)
	})
}

func Test_UniqueItemsByDate(t *testing.T) {
	t.Run("should return empty slice given empty input", func(t *testing.T) {
		input := []IndexItem{}
		result := UniqueItemsByDate(input)
		expected := []IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("should return unique slice of items given non empty input", func(t *testing.T) {
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
		input := []IndexItem{item1, item2, item2, item1, item2}
		result := UniqueItemsByDate(input)
		expected := []IndexItem{item1, item2}
		assert.Equal(t, expected, result)
	})

	t.Run("should return a slice of only one item in case of identical items with different dates", func(t *testing.T) {
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
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       ScriptType(Alias),
			Date:       456,
		}
		input := []IndexItem{item1, item2, item2, item1, item2}
		result := UniqueItemsByDate(input)
		expected := []IndexItem{item2}
		assert.Equal(t, expected, result)
	})
}
