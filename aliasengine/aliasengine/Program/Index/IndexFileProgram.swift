//
//  IndexProgram.swift
//  aliasengine
//
//  Created by Gabriel Coman on 08/06/2022.
//

import Foundation

class IndexFileProgram: Program {
    
    private let reader = DataReader()
    private let writer = DataWriter()
    private let path: String
    
    init(path: String) {
        self.path = path
    }
    
    func run() {
        // update sources
        let existingSources = reader.readSources()
        let fileName = path.fileName
        let source = SourceFile(path: path, name: fileName, type: .alias)
        let sources = (existingSources + [source]).unique { $0.path }
        
        // update items
        let existingItems = reader.readItems()
        
        guard let file = try? String(contentsOfFile: path) else {
            print("Could not open file \(path)")
            return
        }
        let aliasIngester = AliasIngester(filePath: path)
        let newAliases = aliasIngester.process(fileContents: file)
        
        let functionIngester = FunctionIngester(filePath: path)
        let newFunctions = functionIngester.process(fileContents: file)
        
        let items = (existingItems + newAliases + newFunctions).unique { $0.name }
        
        // write everything
        writer.write(sources: sources)
        writer.write(items: items)
        
        print("Ingested \(newAliases.count + newFunctions.count) aliases and functions from \(path)")
        
        exit(0)
    }
}
