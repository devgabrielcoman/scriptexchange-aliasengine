package utils

import (
	"path/filepath"
	"strings"
)

func FileName(path string) string {
	return filepath.Base(path)
}

func FileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
