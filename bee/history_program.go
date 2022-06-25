package main

import "bee/bbee/models"

type HistoryProgram struct{}

func (h HistoryProgram) run() {
	// read history data
	bash_data := h.getBashHistoryData()
	zsh_data := h.getZSHHistoryData()
	data := append(bash_data, zsh_data...)

	// init the controller
	controller := NewSearchController(data)

	// create an empty search cache
	sources := []SourceFile{}
	cache := NewSearchCache(sources)

	// run search program
	program := SearchProgram{controller: *controller, cache: *cache, showPreview: false}
	program.run()
}

func (h HistoryProgram) getBashHistoryData() []models.IndexItem {
	// read history from Bash
	path := getBashHistoryUrl()
	rawHistory, err := ReadFile(path)
	if err != nil {
		return []models.IndexItem{}
	}

	// ingest & process data from history
	ingester := BashHistoryIngester{path: path}
	return ingester.process(rawHistory)
}

func (h HistoryProgram) getZSHHistoryData() []models.IndexItem {
	// read history from ZSH
	path := getZshHistoryUrl()
	rawHistory, err := ReadFile(path)
	if err != nil {
		return []models.IndexItem{}
	}

	// ingest & process data from history
	ingester := ZSHHistoryIngester{path: path}
	return ingester.process(rawHistory)
}
