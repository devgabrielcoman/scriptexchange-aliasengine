package program

import (
	"bee/bbee/data"
	"fmt"
)

// Program will read all of a user's local sources (from the sources file path)
// and list them to standard output
type ListSourceProgram struct{}

func (ls ListSourceProgram) Run() {
	source := data.ReadSourcesRaw()
	fmt.Printf("%s\n", source)
}
