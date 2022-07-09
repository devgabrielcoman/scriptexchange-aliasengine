package program

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSearchController(t *testing.T) {
	t.Run("should initiate the controller in an empty state given empty input", func(t *testing.T) {
		items := []models.IndexItem{}
		controller := NewSearchController(items)

		assert.Equal(t, []models.IndexItem{}, controller.elems)
		assert.Equal(t, []SearchResult{}, controller.results)
		assert.Equal(t, 0, controller.totalLen)
		assert.Equal(t, 0, controller.currentIndex)
	})

	t.Run("should initiate the controller in a valid state given non empty input", func(t *testing.T) {
		item1 := models.IndexItem{
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       models.ScriptType(models.Alias),
			Date:       123,
		}
		item2 := models.IndexItem{
			Name:       "Two",
			Path:       "/my_other_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       models.ScriptType(models.Alias),
			Date:       123,
		}

		items := []models.IndexItem{item1, item2}
		controller := NewSearchController(items)

		assert.Equal(t, items, controller.elems)
		assert.Equal(t, 2, controller.totalLen)
		assert.Equal(t, 0, controller.currentIndex)

		expectedResult := []SearchResult{
			{
				mainText:       "[#696969]/my_path/",
				secondaryText:  "",
				previewTitle:   "",
				previewContent: "",
				command:        "",
				pathOnDisk:     "/full/my_path",
				startLine:      0,
				endLine:        0,
				noHighlight:    true,
				resultType:     SearchResultType(Category),
			},
			{
				mainText:       "   [#659acc]alias [#8cdbff]One",
				secondaryText:  "",
				previewTitle:   "/my_path/One",
				previewContent: "Content",
				command:        "Content",
				pathOnDisk:     "/full/my_path",
				startLine:      0,
				endLine:        0,
				noHighlight:    false,
				resultType:     SearchResultType(Item),
			},
			{
				mainText:       "[#696969]/my_other_path/",
				secondaryText:  "",
				previewTitle:   "",
				previewContent: "",
				command:        "",
				pathOnDisk:     "/full/my_other_path",
				startLine:      0,
				endLine:        0,
				noHighlight:    true,
				resultType:     SearchResultType(Category),
			},
			{
				mainText:       "   [#659acc]alias [#8cdbff]Two",
				secondaryText:  "",
				previewTitle:   "/my_other_path/Two",
				previewContent: "Content",
				command:        "Content",
				pathOnDisk:     "/full/my_other_path",
				startLine:      0,
				endLine:        0,
				noHighlight:    false,
				resultType:     SearchResultType(Item),
			},
		}

		assert.Equal(t, expectedResult, controller.results)
	})
}

func Test_search(t *testing.T) {
	t.Run("should not update results given empty input", func(t *testing.T) {
		items := []models.IndexItem{}
		controller := NewSearchController(items)

		controller.search("term")

		assert.Equal(t, []SearchResult{}, controller.results)
		assert.Equal(t, 0, controller.currentIndex)
	})

	t.Run("should filter results if search is called", func(t *testing.T) {
		item1 := models.IndexItem{
			Name:       "One",
			Path:       "/my_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_path",
			Type:       models.ScriptType(models.Alias),
			Date:       123,
		}
		item2 := models.IndexItem{
			Name:       "Two",
			Path:       "/my_other_path",
			Content:    "Content",
			Comments:   []string{},
			PathOnDisk: "/full/my_other_path",
			Type:       models.ScriptType(models.Alias),
			Date:       123,
		}

		items := []models.IndexItem{item1, item2}
		controller := NewSearchController(items)

		controller.search("other")

		expected := []SearchResult{
			{
				mainText:       "[#696969]/my_other_path/",
				secondaryText:  "",
				previewTitle:   "",
				previewContent: "",
				command:        "",
				pathOnDisk:     "/full/my_other_path",
				startLine:      0,
				endLine:        0,
				noHighlight:    true,
				resultType:     1,
			},
			{
				mainText:       "   [#659acc]alias [#8cdbff]Two",
				secondaryText:  "",
				previewTitle:   "/my_other_path/Two",
				previewContent: "Content",
				command:        "Content",
				pathOnDisk:     "/full/my_other_path",
				startLine:      0,
				endLine:        0,
				noHighlight:    false,
				resultType:     0,
			},
		}

		assert.Equal(t, 0, controller.currentIndex)
		assert.Equal(t, 2, controller.totalLen)
		assert.Equal(t, 2, len(controller.results))
		assert.Equal(t, expected, controller.results)
	})
}
