package main

import (
	"bee/bbee/data"
	"bee/bbee/models"
)

type HistoryProgram struct{}

func (h HistoryProgram) run() {
	// read history data
	bash_data := h.getBashHistoryData()
	zsh_data := h.getZSHHistoryData()
	data := append(bash_data, zsh_data...)

	// init the controller
	controller := NewSearchController(data)

	// create an empty search cache
	sources := []models.SourceFile{}
	cache := NewSearchCache(sources)

	// run search program
	program := SearchProgram{controller: *controller, cache: *cache, showPreview: false}
	program.run()
}

func (h HistoryProgram) getBashHistoryData() []models.IndexItem {
	// read history from Bash
	rawHistory, path, err := data.ReadBashHistory()
	if err != nil {
		return []models.IndexItem{}
	}

	// ingest & process data from history
	ingester := BashHistoryIngester{path: path}
	return ingester.process(rawHistory)
}

func (h HistoryProgram) getZSHHistoryData() []models.IndexItem {
	// read history from ZSH
	rawHistory, path, err := data.ReadZSHHistory()
	if err != nil {
		return []models.IndexItem{}
	}

	// ingest & process data from history
	ingester := ZSHHistoryIngester{path: path}
	return ingester.process(rawHistory)
}
