package program

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
	startLine      int
	endLine        int
	noHighlight    bool
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
		startLine:      item.StartLine,
		endLine:        item.EndLine,
		noHighlight:    false,
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
		startLine:      item.StartLine,
		endLine:        item.EndLine,
		noHighlight:    false,
	}
}

func NewScriptSearchResult(item models.IndexItem) SearchResult {
	var mainText = "   " + style.Color("./"+item.Name, style.ScriptNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	// the script execution command is its content
	// so we can execute both remore and local scripts
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
		startLine:      item.StartLine,
		endLine:        item.EndLine,
		noHighlight:    true,
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
		startLine:      item.StartLine,
		endLine:        item.EndLine,
		noHighlight:    false,
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
		startLine:      0,
		endLine:        0,
		noHighlight:    false,
	}
}

func NewSearchCategory(name string, pathOnDisk string) SearchResult {
	var mainText = style.Color(name+"/", style.ColorDimGray)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  "",
		previewTitle:   pathOnDisk,
		previewContent: "",
		command:        "",
		pathOnDisk:     pathOnDisk,
		resultType:     Category,
		startLine:      0,
		endLine:        0,
		noHighlight:    true,
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
