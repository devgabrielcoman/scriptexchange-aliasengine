package program

import (
	"bee/bbee/data"
	"bee/bbee/models"
	"strings"
)

// The SearchCache provides preview results for a given search result
type SearchCache struct {
	cache map[string]string
}

func NewSearchCache(sources []models.SourceFile) *SearchCache {
	cache := new(SearchCache)
	cache.cache = fillCache(sources)
	return cache
}

func (c SearchCache) getPreviewForSearchResult(result SearchResult) string {
	switch result.resultType {
	case SearchResultType(Item):
		return result.previewContent
	case SearchResultType(Category):
		return c.getPreviewFromCache(result)
	default:
		return ""
	}
}

func fillCache(sources []models.SourceFile) map[string]string {
	cache := make(map[string]string)
	for _, source := range sources {
		var path = source.Path
		var content, err = data.ReadResource(path)
		if err != nil {
			continue
		}
		content = strings.ReplaceAll(content, "$", "\\$")
		content = strings.ReplaceAll(content, "\"", "\\\"")
		cache[path] = content
	}

	return cache
}

func (c SearchCache) getPreviewFromCache(result SearchResult) string {
	var key = result.pathOnDisk
	return c.cache[key]
}
