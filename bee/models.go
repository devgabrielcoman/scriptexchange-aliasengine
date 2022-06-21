package main

// Define script type & enum values
type ScriptType int16

const (
	Alias    ScriptType = 0
	Function ScriptType = 1
	Script   ScriptType = 2
	Export   ScriptType = 3
)

// IndexItems represent references to aliases, functions, scripts
// that a user has saved
type IndexItem struct {
	Name       string
	Content    string
	Path       string
	Comments   []string
	PathOnDisk string
	Type       ScriptType
}

// Define source type & available enum options
type SourceType int16

const (
	Command SourceType = 0
	File    SourceType = 1
)

// SourceFiels represent references to commands of files that a user has registered
type SourceFile struct {
	Path string
	Name string
	Type SourceType
}
