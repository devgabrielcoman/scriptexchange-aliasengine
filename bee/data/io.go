package data

import (
	"bee/bbee/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func getHomeUrl() string {
	dirname, err := os.UserHomeDir()
	check(err)
	return dirname
}

func getDataUrl() string {
	var home = getHomeUrl()
	var path = ".local/bin/scripthub/data.json"
	return fmt.Sprintf("%s/%s", home, path)
}

func getSourcesUrl() string {
	var home = getHomeUrl()
	var path = ".local/bin/scripthub/sources.json"
	return fmt.Sprintf("%s/%s", home, path)
}

func getLastCommandUrl() string {
	var home = getHomeUrl()
	var path = ".local/bin/scripthub/lastcommand"
	return fmt.Sprintf("%s/%s", home, path)
}

func getBashHistoryUrl() string {
	var home = getHomeUrl()
	var path = ".bash_history"
	return fmt.Sprintf("%s/%s", home, path)
}

func getZshHistoryUrl() string {
	var home = getHomeUrl()
	var path = ".zsh_history"
	return fmt.Sprintf("%s/%s", home, path)
}

func ReadItems() []models.IndexItem {
	path := getDataUrl()
	dat, err := os.ReadFile(path)

	if err != nil {
		return []models.IndexItem{}
	}

	var items []models.IndexItem
	json.Unmarshal([]byte(dat), &items)
	return items
}

func ReadSources() []models.SourceFile {
	path := getSourcesUrl()
	dat, err := os.ReadFile(path)

	if err != nil {
		return []models.SourceFile{}
	}

	var sources []models.SourceFile
	json.Unmarshal([]byte(dat), &sources)
	return sources
}

func ReadFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func ReadBashHistory() (string, string, error) {
	var path = getBashHistoryUrl()
	data, err := ReadFile(path)
	if err != nil {
		return "", path, err
	}
	return data, path, err
}

func ReadZSHHistory() (string, string, error) {
	var path = getZshHistoryUrl()
	data, err := ReadFile(path)
	if err != nil {
		return "", path, err
	}
	return data, path, err
}

func WriteLastCommand(command string) {
	path := getLastCommandUrl()
	d1 := []byte(command)
	err := os.WriteFile(path, d1, 0644)
	check(err)
}

func WriteItems(items []models.IndexItem) {
	path := getDataUrl()
	json, err := json.Marshal(items)
	check(err)
	ferr := os.WriteFile(path, json, 0644)
	check(ferr)
}

func WriteSources(sources []models.SourceFile) {
	path := getSourcesUrl()
	json, err := json.Marshal(sources)
	check(err)
	ferr := os.WriteFile(path, json, 0644)
	check(ferr)
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
