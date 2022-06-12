//
//  ScriptIngester.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public class ScriptIngester: FileIngester {
    
    private let alias: String
    private let fileName: String = ".scripts"
    
    init(withAlias alias: String) {
        self.alias = alias
    }
    
    public override func process(fileContents contents: String) -> [IndexItem] {
        return [
            IndexItem(type: .script, name: alias, content: contents, path: fileName, comments: [])
        ]
    }
}
