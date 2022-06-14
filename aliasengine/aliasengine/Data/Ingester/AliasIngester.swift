//
//  AliasIngester.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public class AliasIngester: FileIngester {
    
    private let COMMENT_PREFIX = "#"
    private let ALIAS_PREFIX = "alias "
    private let START_CHAR_QUOTE = "'"
    private let START_CHAR_DBL_QUOTE = "\""
    private let fileName: String
    
    init(filePath path: String) {
        fileName = path.fileName
    }
    
    override public func process(fileContents contents: String) -> [IndexItem] {
        var result: [IndexItem] = []
        let lines: [String] = contents.split(separator: "\n").map(String.init)
        
        for (index, line) in lines.enumerated() {
            if line.starts(with: ALIAS_PREFIX) {
                if let alias = getAlias(fromLine: line) {
                    let comments = getCommentsForAlias(startIndex: index, lines: lines)
                    result.append(alias.copy(withComments: comments))
                }
            }
        }
        
        return result
    }
    
    /// alias lines should always have this format:
    ///  alias name='command'
    private func getAlias(fromLine line: String) -> IndexItem? {
        let aliasWithoutPrefix = line.deletingPrefix(ALIAS_PREFIX)
        let aliasComponents = aliasWithoutPrefix.split(separator: "=")
        guard aliasComponents.count >= 2 else { return nil }
        
        let aliasName = aliasComponents[0]
        var aliasCommand = aliasComponents[1..<aliasComponents.count].joined(separator: "")
        
        if aliasCommand.starts(with: START_CHAR_QUOTE) {
            aliasCommand = aliasCommand.trimmingCharacters(in: CharacterSet(charactersIn: START_CHAR_QUOTE))
        }
        if aliasCommand.starts(with: START_CHAR_DBL_QUOTE) {
            aliasCommand = aliasCommand.trimmingCharacters(in: CharacterSet(charactersIn: START_CHAR_DBL_QUOTE))
        }
        
        return IndexItem(type: .alias,
                         name: String(aliasName),
                         content: aliasCommand,
                         path: fileName,
                         comments: [],
                         pathOnDisk: fileName)
    }
}
