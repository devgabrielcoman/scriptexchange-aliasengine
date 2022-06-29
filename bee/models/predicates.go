package models

import "sort"

// String Pair
type Pair struct {
	A, B string
}

func UniqueItems(slice []IndexItem) []IndexItem {
	keys := make(map[string]bool)
	list := []IndexItem{}
	for _, entry := range slice {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}

	return list
}

func UniqueItemsByDate(slice []IndexItem) []IndexItem {
	keys := make(map[string]IndexItem)
	list := []IndexItem{}

	for _, item := range slice {
		key := item.Name + " " + item.Content
		value, ok := keys[key]
		if ok {
			if item.Date > value.Date {
				keys[key] = item
			}
		} else {
			keys[key] = item
		}
	}

	for _, value := range keys {
		list = append(list, value)
	}

	return list
}

func SortedItemsByPath(slice []IndexItem) []IndexItem {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Path < slice[j].Path
	})
	return slice
}

func UniqueSources(slice []SourceFile) []SourceFile {
	keys := make(map[string]bool)
	list := []SourceFile{}
	for _, entry := range slice {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniquePaths(data []IndexItem) []Pair {
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

func FilterByPath(data []IndexItem, path string) []IndexItem {
	var result = []IndexItem{}

	for _, item := range data {
		if item.Path == path {
			result = append(result, item)
		}
	}

	return result
}
