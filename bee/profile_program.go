package main

import "bee/bbee/data"

type ProfileProgram struct{}

func (p ProfileProgram) run() {
	// get saved items
	items := data.ReadItems()
	controller := NewSearchController(items)

	// get saved source + form cache
	sources := data.ReadSources()
	cache := NewSearchCache(sources)

	// run the search program
	program := SearchProgram{controller: *controller, cache: *cache, showPreview: true}
	program.run()
}
