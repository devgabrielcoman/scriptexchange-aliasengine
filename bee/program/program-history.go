package program

import (
	"bee/bbee/data"
	"bee/bbee/ingester"
	"bee/bbee/models"
)

type HistoryProgram struct{}

func (h HistoryProgram) Run() {
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
	program.Run()
}

func (h HistoryProgram) getBashHistoryData() []models.IndexItem {
	// read history from Bash
	rawHistory, path, err := data.ReadBashHistory()
	if err != nil {
		return []models.IndexItem{}
	}

	// ingest & process data from history
	ingester := ingester.BashHistoryIngester{Path: path}
	return ingester.Process(rawHistory)
}

func (h HistoryProgram) getZSHHistoryData() []models.IndexItem {
	// read history from ZSH
	rawHistory, path, err := data.ReadZSHHistory()
	if err != nil {
		return []models.IndexItem{}
	}

	// ingest & process data from history
	ingester := ingester.ZSHHistoryIngester{Path: path}
	return ingester.Process(rawHistory)
}
