package ingester

import (
	"bee/bbee/models"
	"bee/bbee/utils"
	"strings"
)

// The BashHistoryIngester ingests a .bash_history type file
type BashHistoryIngester struct {
	Path string
}

func (h BashHistoryIngester) Process(content string) []models.IndexItem {
	// separate the contents by line
	var lines []string = strings.Split(content, NEWLINE)
	utils.Reverse(lines)
	var result = []models.IndexItem{}

	for i, line := range lines {
		if line == WHITESPACE || line == SEPARATOR {
			continue
		}

		item := models.IndexItem{
			Name:       line,
			Content:    line,
			Comments:   []string{},
			Path:       h.Path,
			PathOnDisk: h.Path,
			Type:       models.ScriptType(models.History),
			Date:       0, // special case here, for bash we don't really have date info
			StartLine:  i,
			EndLine:    i,
		}
		result = append(result, item)
	}

	return models.UniqueItems(result)
}
