//
//  ScriptIngester.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public class ScriptIngester: FileIngester {
    
    private let alias: String
    private let diskPath: String
    private let fileName: String = ".scripts"
    
    init(withAlias alias: String, andDiskPath path: String) {
        self.alias = alias
        self.diskPath = path
    }
    
    public override func process(fileContents contents: String) -> [IndexItem] {
        return [
            IndexItem(type: .script,
                      name: alias,
                      content: contents,
                      path: fileName,
                      comments: [],
                      pathOnDisk: diskPath)
        ]
    }
}
