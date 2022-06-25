package ingester

import (
	"bee/bbee/models"
	"strings"
)

// The BashHistoryIngester ingests a .bash_history type file
type BashHistoryIngester struct {
	Path string
}

func (h BashHistoryIngester) Process(content string) []models.IndexItem {
	// separate the contents by line
	var lines []string = strings.Split(content, NEWLINE)
	var result = []models.IndexItem{}

	for _, line := range lines {
		item := models.IndexItem{
			Name:       line,
			Content:    line,
			Comments:   []string{},
			Path:       h.Path,
			PathOnDisk: h.Path,
			Type:       models.ScriptType(models.History),
			Date:       0, // special case here, for bash we don't really have date info
		}
		result = append(result, item)
	}

	return models.UniqueItems(result)
}
