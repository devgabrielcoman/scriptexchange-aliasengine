//
//  IndexScriptProgram.swift
//  aliasengine
//
//  Created by Gabriel Coman on 08/06/2022.
//

import Foundation

class IndexScriptProgram: Program {
    
    private let reader = DataReader()
    private let writer = DataWriter()
    private let path: String
    
    init(path: String) {
        self.path = path
    }
    
    func run() {
        // read data
        let initialAlias = path.fileName.deletingAllExtensions
        print("This script will be registed with Alias \(initialAlias, color: .cyan)")
        print("Type ENTER to accept or type a new Alias to override it")
        let inputedAlias = readLine()
        let alias: String
        if let inputedAlias = inputedAlias, !inputedAlias.isEmpty {
            if inputedAlias.isAlphanumeric {
                alias = inputedAlias
            } else {
                print("\("Please provide an Alias that is alphanumberic (contains characters a-z,A-Z and 0-9)", color: .red)")
                exit(0)
            }
        } else {
            alias = initialAlias
        }
        
        // update sources
        let existingSources = reader.readSources()
        let fileName = path.fileName
        let source = SourceFile(path: path, name: alias, type: .script)
        let sources = (existingSources + [source]).unique { $0.path }
        
        // update the new file
        guard let content = try? String(contentsOfFile: path) else {
            print("Could not open file \(path, color: .red)")
            exit(0)
        }
        
        let existingItems = reader.readItems()
        let ingester = ScriptIngester(withAlias: alias, andDiskPath: self.path)
        let scriptItems = ingester.process(fileContents: content)
        let items = (existingItems + scriptItems).unique { $0.name }
        
        // write everything
        writer.write(sources: sources)
        writer.write(items: items)
        
        print("Ingested file \(fileName) at \(path)")

        exit(0)
    }
    
    private func isAliasNameCorrect(alias: String) -> Bool {
        return !alias.contains(" ")
    }
}
