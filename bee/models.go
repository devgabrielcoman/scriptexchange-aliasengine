package main

// Define source file type & available enum options
type SourceType int16

const (
	Command SourceType = 0
	File    SourceType = 1
)

// SourceFiles represent references to commands of files that a user has registered
type SourceFile struct {
	Path string
	Name string
	Type SourceType
}
