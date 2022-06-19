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
	var mainText string
	switch item.Type {
	case ScriptType(Alias):
		mainText = "   " + style.Color("alias", style.AliasKeywordColor) + " " + style.Color(item.Name, style.AliasNameColor)
	case ScriptType(Function):
		mainText = "   " + style.Color("function", style.FunctionKeywordColor) + " " + style.Color(item.Name, style.FunctionNameColor)
	case ScriptType(Script):
		mainText = "   " + style.Color("./"+item.Name, style.ScriptNameColor)
	}
	var previewTitle = item.Path + "/" + item.Name
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

	var command string = ""
	switch item.Type {
	case ScriptType(Alias):
		command = item.Content
	case ScriptType(Function):
		command = item.Content + "\n" + item.Name
	case ScriptType(Script):
		command = item.PathOnDisk
	}

	return SearchResult{
		mainText:       mainText,
		secondaryText:  "",
		previewTitle:   previewTitle,
		previewContent: previewContent,
		command:        command,
		resultType:     Item,
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
