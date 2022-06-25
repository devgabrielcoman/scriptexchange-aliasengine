package models

// Define script type & enum values
type ScriptType int16

const (
	Alias    ScriptType = 0
	Function ScriptType = 1
	Script   ScriptType = 2
	Export   ScriptType = 3
	History  ScriptType = 4
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
	Date       int64
	StartLine  int
	EndLine    int
}
