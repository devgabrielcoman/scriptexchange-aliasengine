package main

import (
	"encoding/json"
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
	var dataPath = ".local/bin/scripthub/data.json"
	return fmt.Sprintf("%s/%s", home, dataPath)
}

func getSourcesUrl() string {
	var home = getHomeUrl()
	var dataPath = ".local/bin/scripthub/sources.json"
	return fmt.Sprintf("%s/%s", home, dataPath)
}

func getLastCommandUrl() string {
	var home = getHomeUrl()
	var dataPath = ".local/bin/scripthub/lastcommand"
	return fmt.Sprintf("%s/%s", home, dataPath)
}

func ReadItems() []IndexItem {
	path := getDataUrl()
	dat, err := os.ReadFile(path)

	if err != nil {
		return []IndexItem{}
	}

	var items []IndexItem
	json.Unmarshal([]byte(dat), &items)
	return items
}

func ReadSources() []SourceFile {
	path := getSourcesUrl()
	dat, err := os.ReadFile(path)

	if err != nil {
		return []SourceFile{}
	}

	var sources []SourceFile
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

func WriteLastCommand(command string) {
	path := getLastCommandUrl()
	d1 := []byte(command)
	err := os.WriteFile(path, d1, 0644)
	check(err)
}

func WriteItems(items []IndexItem) {
	path := getDataUrl()
	json, err := json.Marshal(items)
	check(err)
	ferr := os.WriteFile(path, json, 0644)
	check(ferr)
}

func WriteSources(sources []SourceFile) {
	path := getSourcesUrl()
	json, err := json.Marshal(sources)
	check(err)
	ferr := os.WriteFile(path, json, 0644)
	check(ferr)
}
