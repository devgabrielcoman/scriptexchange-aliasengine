package main

import (
	"bee/bbee/program"
	"flag"
)

func main() {
	var register string
	var remove string
	var script bool
	var update bool
	var history bool
	var listSource bool

	flag.StringVar(&register, "register", "", "Register a file of aliases or functions or a script")
	flag.StringVar(&remove, "remove", "", "Remove a file of aliases/functions or a full script")
	flag.BoolVar(&script, "s", false, "Register file as script")
	flag.BoolVar(&update, "u", false, "Update all data")
	flag.BoolVar(&history, "h", false, "Search through bash history")
	flag.BoolVar(&listSource, "ls", false, "List the contents of the source file")
	flag.Parse()

	if register != "" {
		if script {
			program := program.RegisterFileProgram{Path: register, IsScript: true}
			program.Run()
		} else {
			program := program.RegisterFileProgram{Path: register, IsScript: false}
			program.Run()
		}
	} else if remove != "" {
		program := program.RemoveProgram{Name: remove}
		program.Run()
	} else if update {
		program := program.UpdateProgram{}
		program.Run()
	} else if history {
		program := program.HistoryProgram{}
		program.Run()
	} else if listSource {
		program := program.ListSourceProgram{}
		program.Run()
	} else {
		program := program.ProfileProgram{}
		program.Run()
	}
}
