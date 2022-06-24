package main

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// String Pair
type Pair struct {
	a, b string
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func uniqueItems(slice []IndexItem) []IndexItem {
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

func uniqueSources(slice []SourceFile) []SourceFile {
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

func uniquePaths(data []IndexItem) []Pair {
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

func filterByPath(data []IndexItem, path string) []IndexItem {
	var result = []IndexItem{}

	for _, item := range data {
		if item.Path == path {
			result = append(result, item)
		}
	}

	return result
}

func fileName(path string) string {
	return filepath.Base(path)
}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func dateFormat(d int) string {
	intTime := int64(d)
	t := time.Unix(intTime, 0)
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func lenientAtoi(stringDate string) int {
	num, err := strconv.Atoi(stringDate)
	if err == nil {
		return num
	} else {
		return 0
	}
}
