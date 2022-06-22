package main

import (
	"strings"

	"github.com/samber/lo"
)

// The Search Controller contains data & exposes functions related
// to the search through all of the aliases, functions and scripts
// a user has saved
type SearchController struct {
	elems        []IndexItem
	results      []SearchResult
	currentIndex int
	totalLen     int
}

func NewSearchController(elems []IndexItem) *SearchController {
	controller := new(SearchController)
	controller.elems = elems
	controller.results = controller.formResults(elems)
	controller.totalLen = len(elems)
	controller.resetCurrentIndex()
	return controller
}

func (c *SearchController) search(term string) {
	var filtered = lo.Filter(c.elems, func(item IndexItem, i int) bool {
		var key = item.Path + "/" + item.Name
		return strings.Contains(strings.ToLower(key), strings.ToLower(term))
	})
	c.results = c.formResults(filtered)
	c.resetCurrentIndex()
}

func (c *SearchController) formResults(items []IndexItem) []SearchResult {
	var result = []SearchResult{}

	var paths = uniquePaths(items)

	for _, path := range paths {
		result = append(result, NewSearchCategory(path.a, path.b))
		var filtered = filterByPath(items, path.a)
		for _, item := range filtered {
			result = append(result, NewSearchResult(item))
		}
	}

	return result
}

func (c *SearchController) moveDown() {
	// var nextItem = c.getNextItem()
	// var increment int
	// if nextItem.resultType == SearchResultType(Category) {
	// 	increment = 2
	// } else {
	// 	increment = 1
	// }
	c.currentIndex = min(c.currentIndex+1, len(c.results)-1)
}

func (c *SearchController) moveUp() {
	// var nextItem = c.getPrevItem()
	// var increment int
	// if nextItem.resultType == SearchResultType(Category) {
	// 	increment = 2
	// } else {
	// 	increment = 1
	// }
	c.currentIndex = max(c.currentIndex-1, 0) // we "max" with 1 because the 1st element could be a "category" type
}

func (c *SearchController) resetCurrentIndex() {
	if len(c.results) > 1 {
		c.currentIndex = 0 // start from 1st index, so not on the 0th element, which is a "category" type
	} else {
		c.currentIndex = 0
	}
}

func (c SearchController) getCurrentItem() SearchResult {
	if c.currentIndex >= 0 && c.currentIndex < len(c.results) {
		return c.results[c.currentIndex]
	} else {
		return NewEmptySearchResult()
	}
}

// func (c SearchController) getNextItem() SearchResult {
// 	var nextIndex = min(c.currentIndex+1, len(c.results)-1)
// 	if nextIndex >= 0 && nextIndex < len(c.results) {
// 		return c.results[nextIndex]
// 	} else {
// 		return NewEmptySearchResult()
// 	}
// }

// func (c SearchController) getPrevItem() SearchResult {
// 	var prevIndex = max(c.currentIndex-1, 0)
// 	if prevIndex >= 0 && prevIndex < len(c.results) {
// 		return c.results[prevIndex]
// 	} else {
// 		return NewEmptySearchResult()
// 	}
// }

func (c SearchController) getNumberOfSearchResults() int {
	filtered := lo.Filter(c.results, func(result SearchResult, i int) bool {
		return result.resultType == SearchResultType(Item)
	})
	return len(filtered)
}
