//
//  main.swift
//  aliasengine
//
//  Created by Liviu Coman on 04.06.2022.
//

import Foundation
import ArgumentParser

struct Arguments: ParsableArguments {
    @Option(name: [.customShort("r"), .long], help: "A path to an alias file to register")
    var register: String?
    
    @Flag(name: [.customShort("s")], help: "Ingest a scriot whole")
    var script: Bool = false
    
    @Argument(help: "Update all data")
    var update: String?
}

let arguments = Arguments.parseOrExit()
if let filePath = arguments.register {
    if arguments.script {
        let program = IndexScriptProgram(path: filePath)
        program.run()
    } else {
        let program = IndexFileProgram(path: filePath)
        program.run()
    }
}
else if let update = arguments.update, update == "update" {
    let program = UpdateProgram()
    program.run()
}
else {
    let program = SearchProgram()
    program.run()
}
