package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RegisterFileProgram struct {
	path     string
	isScript bool
}

func (r RegisterFileProgram) run() {
	if r.isScript {
		r.registerScript()
	} else {
		r.registerConfigFile()
	}
}

func (r RegisterFileProgram) registerConfigFile() {
	// update sources
	var sources []SourceFile = ReadSources()
	var source = SourceFile{Path: r.path, Name: fileName(r.path), Type: SourceType(Command)}
	sources = append(sources, source)
	sources = uniqueSources(sources)

	// read config files liek .bashrc, .profile, etc
	var existingItems = ReadItems()

	// open file
	contents, err := ReadFile(r.path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return
	}

	// process new elements
	ingester := ConfigIngester{filePath: r.path}
	var newItems = ingester.process(contents)
	var items = append(existingItems, newItems...)
	items = uniqueItems(items)

	// write data
	WriteSources(sources)
	WriteItems(items)

	fmt.Printf("Added %d new elements\n", len(newItems))
}

func (r RegisterFileProgram) registerScript() {
	// get the user to input the alias
	var fileName = fileName(r.path)
	var initialAlias = fileNameWithoutExtTrimSuffix(fileName)
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
	var sources []SourceFile = ReadSources()
	var source = SourceFile{Path: r.path, Name: fileName, Type: SourceType(Script)}
	sources = append(sources, source)
	sources = uniqueSources(sources)

	// read script
	var existingItems = ReadItems()

	// open file
	contents, err := ReadFile(r.path)

	// gently handle error
	if err != nil {
		fmt.Println(err)
		return
	}

	ingester := ScriptIngester{alias: alias, path: r.path}
	var newItems = ingester.process(contents)
	var items = append(existingItems, newItems...)
	items = uniqueItems(items)

	// write data
	WriteSources(sources)
	WriteItems(items)

	fmt.Printf("Added %d new elements\n", len(newItems))
}
