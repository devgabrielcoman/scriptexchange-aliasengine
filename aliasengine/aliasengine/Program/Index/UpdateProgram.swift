//
//  UpdateProgram.swift
//  aliasengine
//
//  Created by Gabriel Coman on 10/06/2022.
//

import Foundation

class UpdateProgram: Program {
    private let reader = DataReader()
    private let writer = DataWriter()
    
    func run() {
        let sources = reader.readSources()
        var items: [IndexItem] = []
        
        for source in sources {
            switch source.type {
            case .alias:
                items += reloadAliasFiles(source: source)
            case .script:
                items += reloadScriptFiles(source: source)
            }
        }
        
        writer.write(items: items.unique { $0.name })
    }
    
    private func reloadAliasFiles(source: SourceFile) -> [IndexItem] {
        guard let file = try? String(contentsOfFile: source.path) else {
            print("Could not open file \(source.path, color: .red)")
            return []
        }
        
        let aliasIngester = AliasIngester(filePath: source.path)
        let newAliases = aliasIngester.process(fileContents: file)
        
        let functionIngester = FunctionIngester(filePath: source.path)
        let newFunctions = functionIngester.process(fileContents: file)
        
        print("Ingested \(newAliases.count + newFunctions.count) aliases and functions from \(source.path)")
        
        return (newAliases + newFunctions)
    }
    
    private func reloadScriptFiles(source: SourceFile) -> [IndexItem] {
        guard let content = try? String(contentsOfFile: source.path) else {
            print("Could not open file \(source.path, color: .red)")
            return []
        }
        
        let ingester = ScriptIngester(withAlias: source.name, andDiskPath: source.path)
        
        print("Ingested file \(source.name) at \(source.path)")
        
        return ingester.process(fileContents: content)
    }
}
