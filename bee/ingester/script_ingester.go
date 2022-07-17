package ingester

import (
	"bee/bbee/models"
)

// The ScriptIngester just ingests a new full script
type ScriptIngester struct {
	Alias       string
	Path        string
	CurrentTime int64
}

func (s ScriptIngester) Process(content string) []models.IndexItem {
	return []models.IndexItem{
		{
			Name:       s.Alias,
			Content:    content,
			Path:       ".scripts",
			Comments:   []string{},
			PathOnDisk: ".scripts", // s.Path,
			Type:       models.ScriptType(models.Script),
			Date:       s.CurrentTime,
			StartLine:  0,
			EndLine:    0,
		},
	}
}
