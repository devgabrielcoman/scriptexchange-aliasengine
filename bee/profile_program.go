package main

type ProfileProgram struct{}

func (p ProfileProgram) run() {
	// get saved items
	data := ReadItems()
	controller := NewSearchController(data)

	// get saved source + form cache
	sources := ReadSources()
	cache := NewSearchCache(sources)

	// run the search program
	program := SearchProgram{controller: *controller, cache: *cache, showPreview: true}
	program.run()
}
