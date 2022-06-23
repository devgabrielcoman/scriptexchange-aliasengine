package main

type HistoryProgram struct{}

func (h HistoryProgram) run() {
	// read history from bash
	rawHistory, err := ReadHistory()
	check(err)

	// ingest & process data and feed it to the controller
	ingester := HistoryIngester{path: getHistoryUrl()}
	data := ingester.process(rawHistory)
	controller := NewSearchController(data)

	// create an empty search cache
	sources := []SourceFile{}
	cache := NewSearchCache(sources)

	// run search program
	program := SearchProgram{controller: *controller, cache: *cache, showPreview: false}
	program.run()
}
