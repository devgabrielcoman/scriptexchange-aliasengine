package main

import "flag"

func main() {
	var register string
	var script bool
	var update bool
	var history bool

	flag.StringVar(&register, "register", "", "Register a file of aliases or functions or a script")
	flag.BoolVar(&script, "s", false, "Register file as script")
	flag.BoolVar(&update, "u", false, "Update all data")
	flag.BoolVar(&history, "h", false, "Search through bash history")
	flag.Parse()

	if register != "" {
		if script {
			program := RegisterFileProgram{path: register, isScript: true}
			program.run()
		} else {
			program := RegisterFileProgram{path: register, isScript: false}
			program.run()
		}
	} else if update {
		program := UpdateProgram{}
		program.run()
	} else if history {
		program := HistoryProgram{}
		program.run()
	} else {
		program := ProfileProgram{}
		program.run()
	}
}
