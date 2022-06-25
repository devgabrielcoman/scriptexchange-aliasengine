package main

import (
	"bee/bbee/models"
)

// String Pair
type Pair struct {
	a, b string
}

func uniqueItems(slice []models.IndexItem) []models.IndexItem {
	keys := make(map[string]bool)
	list := []models.IndexItem{}
	for _, entry := range slice {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}

	return list
}

func uniqueItemsByDate(slice []models.IndexItem) []models.IndexItem {
	keys := make(map[string]models.IndexItem)
	list := []models.IndexItem{}

	for _, item := range slice {
		value, ok := keys[item.Name]
		if ok {
			if item.Date > value.Date {
				keys[item.Name] = item
			}
		} else {
			keys[item.Name] = item
		}
	}

	for _, value := range keys {
		list = append(list, value)
	}

	return list
}

func uniqueSources(slice []models.SourceFile) []models.SourceFile {
	keys := make(map[string]bool)
	list := []models.SourceFile{}
	for _, entry := range slice {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}
	return list
}

func uniquePaths(data []models.IndexItem) []Pair {
	keys := make(map[string]bool)
	list := []Pair{}

	for _, entry := range data {
		if _, value := keys[entry.Path]; !value {
			keys[entry.Path] = true
			list = append(list, Pair{entry.Path, entry.PathOnDisk})
		}
	}

	return list
}

func filterByPath(data []models.IndexItem, path string) []models.IndexItem {
	var result = []models.IndexItem{}

	for _, item := range data {
		if item.Path == path {
			result = append(result, item)
		}
	}

	return result
}
