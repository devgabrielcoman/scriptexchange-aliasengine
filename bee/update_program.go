package main

import (
	"bee/bbee/data"
	"bee/bbee/models"
	"fmt"
)

type UpdateProgram struct {
}

func (u UpdateProgram) run() {
	var result = []models.IndexItem{}
	var sources = data.ReadSources()

	for _, source := range sources {
		switch source.Type {
		case models.SourceType(models.Command):
			items := u.updateConfigFiles(source)
			result = append(result, items...)
		case models.SourceType(models.File):
			items := u.updateScriptFiles(source)
			result = append(result, items...)
		}
	}

	result = uniqueItemsByDate(result)
	data.WriteItems(result)

	fmt.Printf("Updated %d elements\n", len(result))
}

func (u UpdateProgram) updateConfigFiles(source models.SourceFile) []models.IndexItem {
	// open file
	contents, err := data.ReadFile(source.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := CurrentTime()
	ingester := ConfigIngester{filePath: source.Path, currentTime: time}
	return ingester.process(contents)
}

func (u UpdateProgram) updateScriptFiles(source models.SourceFile) []models.IndexItem {
	// open file
	contents, err := data.ReadFile(source.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := CurrentTime()
	ingester := ScriptIngester{alias: source.Name, path: source.Path, currentTime: time}
	return ingester.process(contents)
}
