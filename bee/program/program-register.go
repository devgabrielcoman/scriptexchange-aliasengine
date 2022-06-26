package program

import (
	"bee/bbee/data"
	"bee/bbee/ingester"
	"bee/bbee/models"
	"bee/bbee/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RegisterFileProgram struct {
	Path     string
	IsScript bool
}

func (r RegisterFileProgram) Run() {
	if r.IsScript {
		r.registerScript()
	} else {
		r.registerConfigFile()
	}
}

func (r RegisterFileProgram) registerConfigFile() {
	// update sources
	var sources []models.SourceFile = data.ReadSources()
	var source = models.SourceFile{Path: r.Path, Name: utils.FileName(r.Path), Type: models.SourceType(models.Command)}
	sources = append(sources, source)
	sources = models.UniqueSources(sources)

	// read config files liek .bashrc, .profile, etc
	var existingItems = data.ReadItems()

	// open file
	contents, err := data.ReadFile(r.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return
	}

	// process new elements
	time := utils.CurrentTime()
	ingester := ingester.ConfigIngester{FilePath: r.Path, CurrentTime: time}
	var newItems = ingester.Process(contents)
	var items = append(existingItems, newItems...)
	items = models.UniqueItemsByDate(items)
	items = models.SortedItemsByPath(items)

	// write data
	data.WriteSources(sources)
	data.WriteItems(items)

	fmt.Printf("Added %d new elements\n", len(newItems))
}

func (r RegisterFileProgram) registerScript() {
	// get the user to input the alias
	var fileName = utils.FileName(r.Path)
	var initialAlias = utils.FileNameWithoutExtTrimSuffix(fileName)
	fmt.Printf("This script will be registered with alias %s\nPress ENTER to accept or type a new Alias to override it\n", initialAlias)

	buffer := bufio.NewReader(os.Stdin)
	text, _ := buffer.ReadString('\n')

	var alias string = ""
	if text == "\n" {
		alias = initialAlias
	} else {
		alias = strings.Trim(text, "\n")
	}

	// update sources
	var sources []models.SourceFile = data.ReadSources()
	var source = models.SourceFile{Path: r.Path, Name: fileName, Type: models.SourceType(models.File)}
	sources = append(sources, source)
	sources = models.UniqueSources(sources)

	// read script
	var existingItems = data.ReadItems()

	// open file
	contents, err := data.ReadFile(r.Path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return
	}

	time := utils.CurrentTime()
	ingester := ingester.ScriptIngester{Alias: alias, Path: r.Path, CurrentTime: time}
	var newItems = ingester.Process(contents)
	var items = append(existingItems, newItems...)
	items = models.UniqueItemsByDate(items)
	items = models.SortedItemsByPath(items)

	// write data
	data.WriteSources(sources)
	data.WriteItems(items)

	fmt.Printf("Added %d new elements\n", len(newItems))
}
