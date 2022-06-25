package main

import (
	"bee/bbee/models"
	"fmt"
)

type UpdateProgram struct {
}

func (u UpdateProgram) run() {
	var result = []models.IndexItem{}
	var sources = ReadSources()

	for _, source := range sources {
		switch source.Type {
		case SourceType(Command):
			items := u.updateConfigFiles(source)
			result = append(result, items...)
		case SourceType(File):
			items := u.updateScriptFiles(source)
			result = append(result, items...)
		}
	}

	result = uniqueItemsByDate(result)
	WriteItems(result)

	fmt.Printf("Updated %d elements\n", len(result))
}

func (u UpdateProgram) updateConfigFiles(source SourceFile) []models.IndexItem {
	// open file
	contents, err := ReadFile(source.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := CurrentTime()
	ingester := ConfigIngester{filePath: source.Path, currentTime: time}
	return ingester.process(contents)
}

func (u UpdateProgram) updateScriptFiles(source SourceFile) []models.IndexItem {
	// open file
	contents, err := ReadFile(source.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := CurrentTime()
	ingester := ScriptIngester{alias: source.Name, path: source.Path, currentTime: time}
	return ingester.process(contents)
}
