package utils

import (
	"net/url"
	"path/filepath"
	"strings"
)

func IsHttpUrl(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

func FileName(path string) string {
	if IsHttpUrl(path) {
		myUrl, err := url.Parse(path)
		if err != nil {
			return path
		}
		return filepath.Base(myUrl.Path)
	} else {
		return filepath.Base(path)
	}
}

func FileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
