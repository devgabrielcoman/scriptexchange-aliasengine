package ingester

import (
	"bee/bbee/models"
	"bee/bbee/utils"
	"strings"
)

// The ZSHHistoryIngester ingests a .zsh_history type file
type ZSHHistoryIngester struct {
	Path string
}

func (z ZSHHistoryIngester) Process(content string) []models.IndexItem {
	// separate the contents by line
	lines := strings.Split(content, NEWLINE)
	// ZSH history is appended to, so we need to reverse that order
	utils.Reverse(lines)
	var result = []models.IndexItem{}

	for i, line := range lines {
		lineItems := strings.Split(line, ZSH_HISTORY_SEP)

		// not a valid line in the format date;command
		if len(lineItems) < 2 {
			continue
		}

		date := utils.LenientAtoi64(z.parseZSHDateItem(lineItems[0]))
		command := strings.Join(lineItems[1:], SEPARATOR)
		item := models.IndexItem{
			Name:       command,
			Content:    command,
			Comments:   []string{},
			Path:       z.Path,
			PathOnDisk: z.Path,
			Type:       models.ScriptType(models.History),
			Date:       date,
			StartLine:  i,
			EndLine:    i,
		}
		result = append(result, item)
	}

	result = models.UniqueItemsByDate(result)

	return result
}

// parses a ZSH date item in the format: ": 1642437214:0"
func (z ZSHHistoryIngester) parseZSHDateItem(rawDate string) string {
	date := strings.TrimPrefix(strings.TrimSuffix(rawDate, ":0"), ": ")
	return date
}
