package program

import (
	"bee/bbee/models"
	"strings"
)

// Represents a Search Key formed from an Index Item
type SearchKey struct {
	item models.IndexItem
}

func (k SearchKey) Contains(term string) bool {
	var queries = k.formSearchQueries()
	for _, query := range queries {
		if strings.Contains(strings.ToLower(query), strings.ToLower(term)) {
			return true
		}
	}
	return false
}

func (k SearchKey) formSearchQueries() []string {
	item := k.item
	switch item.Type {
	case models.ScriptType(models.Alias):
		return []string{
			item.Path + "/alias " + item.Name,
			item.Path + "/" + item.Name,
		}
	case models.ScriptType(models.Function):
		return []string{
			item.Path + "/function " + item.Name,
			item.Path + "/" + item.Name,
		}
	case models.ScriptType(models.Export):
		return []string{
			item.Path + "/export " + item.Name,
			item.Path + "/" + item.Name,
		}
	case models.ScriptType(models.Script):
		return []string{
			item.Path + "/./" + item.Name,
			item.Path + "/" + item.Name,
		}
	case models.ScriptType(models.History):
		return []string{
			item.Name,
			item.Content,
		}
	default:
		return []string{}
	}
}
