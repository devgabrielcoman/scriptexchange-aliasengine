package main

import (
	"bee/bbee/models"
	"bee/bbee/style"
	"strings"
)

type SearchResultType int16

const (
	Item     SearchResultType = 0
	Category SearchResultType = 1
	Empty    SearchResultType = 2
)

// Represents a Search Result from an IndexItem a user has already registered
type SearchResult struct {
	mainText       string
	secondaryText  string
	previewTitle   string
	previewContent string
	command        string
	pathOnDisk     string
	resultType     SearchResultType
}

func NewSearchResult(item models.IndexItem) SearchResult {
	switch item.Type {
	case models.ScriptType(models.Alias):
		return NewAliasSearchResult(item)
	case models.ScriptType(models.Function):
		return NewFunctionSearchResult(item)
	case models.ScriptType(models.Script):
		return NewScriptSearchResult(item)
	case models.ScriptType(models.Export):
		return NewExportSearchResult(item)
	case models.ScriptType(models.History):
		return NewHistorySearchResult(item)
	default:
		return NewEmptySearchResult()
	}
}

func NewAliasSearchResult(item models.IndexItem) SearchResult {
	var mainText = "   " + style.Color("alias", style.AliasKeywordColor) + " " + style.Color(item.Name, style.AliasNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.Content
	var pathOnDisk = item.PathOnDisk
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		pathOnDisk:     pathOnDisk,
		resultType:     resultType,
	}
}

func NewFunctionSearchResult(item models.IndexItem) SearchResult {
	var mainText = "   " + style.Color("function", style.FunctionKeywordColor) + " " + style.Color(item.Name, style.FunctionNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.Content + "\n" + item.Name
	var pathOnDisk = item.PathOnDisk
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		pathOnDisk:     pathOnDisk,
		resultType:     resultType,
	}
}

func NewScriptSearchResult(item models.IndexItem) SearchResult {
	var mainText = "   " + style.Color("./"+item.Name, style.ScriptNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.PathOnDisk
	var pathOnDisk = item.PathOnDisk
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		pathOnDisk:     pathOnDisk,
		resultType:     resultType,
	}
}

func NewExportSearchResult(item models.IndexItem) SearchResult {
	var mainText = "   " + style.Color("export", style.ExportKeywordColor) + " " + style.Color(item.Name, style.ExportNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.PathOnDisk
	var pathOnDisk = item.PathOnDisk
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		pathOnDisk:     pathOnDisk,
		resultType:     resultType,
	}
}

func NewHistorySearchResult(item models.IndexItem) SearchResult {
	var mainText string
	if item.Date == 0 {
		mainText = "   " + style.Color(item.Content, style.ScriptNameColor)
	} else {
		mainText = "   " + style.Color(dateFormat(item.Date), style.AliasNameColor) + " " + style.Color(item.Content, style.ScriptNameColor)
	}
	return SearchResult{
		mainText:       mainText,
		secondaryText:  "",
		previewTitle:   item.Path,
		previewContent: "",
		command:        item.Content,
		pathOnDisk:     item.Path,
		resultType:     SearchResultType(Item),
	}
}

func NewEmptySearchResult() SearchResult {
	return SearchResult{
		mainText:       "",
		secondaryText:  "",
		previewTitle:   "",
		previewContent: "",
		command:        "",
		pathOnDisk:     "",
		resultType:     Empty,
	}
}

func NewSearchCategory(name string, pathOnDisk string) SearchResult {
	var mainText = style.Color(name+"/", style.ColorDimGray)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  "",
		previewTitle:   "",
		previewContent: "",
		command:        "",
		pathOnDisk:     pathOnDisk,
		resultType:     Category,
	}
}

func createPreviewContent(item models.IndexItem) string {
	var comment = strings.Join(item.Comments[:], "\n")
	var full []string
	if len(item.Comments) > 0 {
		full = []string{comment, "\n", item.Content}
	} else {
		full = []string{item.Content}
	}
	var previewContent = strings.Join(full, "")
	// replace all occurances where we have a variable with one an escaped one
	// this is needed
	previewContent = strings.ReplaceAll(previewContent, "$", "\\$")
	previewContent = strings.ReplaceAll(previewContent, "\"", "\\\"")
	return previewContent
}

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
