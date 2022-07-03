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
	var sources = data.ReadUserSources()
	sources = models.UniqueSources(sources)

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
	result = models.SortedItemsByPath(result)
	data.WriteItems(result)

	fmt.Printf("Updated %d elements\n", len(result))
}

func (u UpdateProgram) updateConfigFiles(source models.SourceFile) []models.IndexItem {
	// open file
	contents, err := data.ReadResource(source.Path)

	// gently handle error
	if err != nil {
		fmt.Printf("Error updating %s ==> %s\n", source.Name, source.Path)
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := utils.CurrentTime()
	ingester := ingester.ConfigIngester{FilePath: source.Path, CurrentTime: time}
	result := ingester.Process(contents)

	fmt.Printf("Updated %s ==> %s\n", source.Name, source.Path)

	return result
}

func (u UpdateProgram) updateScriptFiles(source models.SourceFile) []models.IndexItem {
	// open file
	contents, err := data.ReadResource(source.Path)

	// gently handle error
	if err != nil {
		fmt.Printf("Error updating %s ==> %s\n", source.Name, source.Path)
		fmt.Println(err)
		return []models.IndexItem{}
	}

	time := utils.CurrentTime()
	ingester := ingester.ScriptIngester{Alias: source.Name, Path: source.Path, CurrentTime: time}
	result := ingester.Process(contents)

	fmt.Printf("Updated %s ==> %s\n", source.Name, source.Path)

	return result
}
