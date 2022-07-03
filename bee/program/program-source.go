package program

import (
	"bee/bbee/data"
	"bee/bbee/models"
	"fmt"
)

// This program takes in a valid local or remote Path to a new Source file
// and concatenates it with the user's existing ones
type SourceProgram struct {
	Path string
}

func (s SourceProgram) Run() {
	fmt.Printf("Reading source %s\n", s.Path)

	old := data.ReadUserSources()
	new := data.ReadSourceFile(s.Path)

	var sources = s.mergeSources(old, new)

	fmt.Printf("Added %d valid new sources\n", len(new))

	data.WriteSources(sources)
}

func (s SourceProgram) mergeSources(old []models.SourceFile, new []models.SourceFile) []models.SourceFile {
	var sources = append(old, new...)
	return models.UniqueSources(sources)
}
