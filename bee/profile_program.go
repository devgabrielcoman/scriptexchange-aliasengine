package main

import "bee/bbee/models"

type ProfileProgram struct{}

func (p ProfileProgram) run() {
	// get saved items
	var data []models.IndexItem = ReadItems()
	controller := NewSearchController(data)

	// get saved source + form cache
	var sources []SourceFile = ReadSources()
	cache := NewSearchCache(sources)

	// run the search program
	program := SearchProgram{controller: *controller, cache: *cache, showPreview: true}
	program.run()
}
