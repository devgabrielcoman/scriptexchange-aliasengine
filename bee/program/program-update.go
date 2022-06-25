package program

import (
	"bee/bbee/data"
	"bee/bbee/ingester"
	"bee/bbee/models"
	"bee/bbee/utils"
	"fmt"
)

type UpdateProgram struct {
}

func (u UpdateProgram) Run() {
	var result = []models.IndexItem{}
	var sources = data.ReadSources()
	sources = models.UniqueSources(sources)
	sources = models.SortedSources(sources)

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

	result = models.UniqueItemsByDate(result)
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

	time := utils.CurrentTime()
	ingester := ingester.ConfigIngester{FilePath: source.Path, CurrentTime: time}
	return ingester.Process(contents)
}

func (u UpdateProgram) updateScriptFiles(source models.SourceFile) []models.IndexItem {
	// open file
	contents, err := data.ReadFile(source.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := utils.CurrentTime()
	ingester := ingester.ScriptIngester{Alias: source.Name, Path: source.Path, CurrentTime: time}
	return ingester.Process(contents)
}
