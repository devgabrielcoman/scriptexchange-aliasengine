package program

import (
	"bee/bbee/data"
	"bee/bbee/models"
	"fmt"
)

// This Program will remove a source from the user's source list, given
// either the Name or the full Path.
// If no such source is found, nothing will be removed.
type RemoveProgram struct {
	Name string
}

func (r RemoveProgram) Run() {
	fmt.Printf("Removing %s\n", r.Name)

	sources := data.ReadUserSources()
	var result = r.removeSource(sources)

	if len(result) == len(sources) {
		fmt.Printf("Could not find file %s to remove\n", r.Name)
	} else {
		data.WriteSources(result)
		fmt.Printf("Removed %s from sources\n", r.Name)
	}
}

func (r RemoveProgram) removeSource(original []models.SourceFile) []models.SourceFile {
	var result = []models.SourceFile{}
	for _, source := range original {
		if source.Name != r.Name && source.Path != r.Name {
			result = append(result, source)
		}
	}

	return result
}
