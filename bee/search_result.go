package main

import (
	"example/bbee/style"
	"strings"
)

type SearchResultType int16

const (
	Item     SearchResultType = 0
	Category SearchResultType = 1
	Empty    SearchResultType = 2
)

// A search result form from an IndexItem a user has  already registered
type SearchResult struct {
	mainText       string
	secondaryText  string
	previewTitle   string
	previewContent string
	command        string
	resultType     SearchResultType
}

func NewSearchResult(item IndexItem) SearchResult {
	switch item.Type {
	case ScriptType(Alias):
		return NewAliasSearchResult(item)
	case ScriptType(Function):
		return NewFunctionSearchResult(item)
	case ScriptType(Script):
		return NewScriptSearchResult(item)
	case ScriptType(Export):
		return NewExportSearchResult(item)
	default:
		return NewEmptySearchResult()
	}
}

func NewAliasSearchResult(item IndexItem) SearchResult {
	var mainText = "   " + style.Color("alias", style.AliasKeywordColor) + " " + style.Color(item.Name, style.AliasNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.Content
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		resultType:     resultType,
	}
}

func NewFunctionSearchResult(item IndexItem) SearchResult {
	var mainText = "   " + style.Color("function", style.FunctionKeywordColor) + " " + style.Color(item.Name, style.FunctionNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.Content + "\n" + item.Name
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		resultType:     resultType,
	}
}

func NewScriptSearchResult(item IndexItem) SearchResult {
	var mainText = "   " + style.Color("./"+item.Name, style.ScriptNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.PathOnDisk
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		resultType:     resultType,
	}
}

func NewExportSearchResult(item IndexItem) SearchResult {
	var mainText = "   " + style.Color("export", style.ExportKeywordColor) + " " + style.Color(item.Name, style.ExportNameColor)
	var secondaryText = ""
	var previewTitle = item.Path + "/" + item.Name
	var previewContent = createPreviewContent(item)
	var command = item.PathOnDisk
	var resultType = SearchResultType(Item)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  secondaryText,
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		resultType:     resultType,
	}
}

func NewEmptySearchResult() SearchResult {
	return SearchResult{
		mainText:       "",
		secondaryText:  "",
		previewTitle:   "",
		previewContent: "",
		command:        "",
		resultType:     Empty,
	}
}

func NewSearchCategory(name string) SearchResult {
	var mainText = style.Color(name, style.ColorDimGray)
	return SearchResult{
		mainText:       mainText,
		secondaryText:  "",
		previewTitle:   "",
		previewContent: "",
		command:        "",
		resultType:     Category,
	}
}

func createPreviewContent(item IndexItem) string {
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
