//
//  FileIngester.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public class FileIngester: Ingester {
    private let COMMENT_PREFIX = "#"
    
    public func process(fileContents contents: String) -> [IndexItem]{
        return []
    }
    
    internal func getCommentsForAlias(startIndex index: Int, lines: [String]) -> [String] {
        var startIndex = index - 1
        var commentsArray: [String] = []
        var prevLine = lines[safe: startIndex]
        while let previous = prevLine {
            if previous.trimmingCharacters(in: .whitespaces).starts(with: COMMENT_PREFIX) {
                commentsArray.append(previous.deletingPrefix(COMMENT_PREFIX).trimmingCharacters(in: .whitespaces))
                startIndex -= 1
                prevLine = lines[safe: startIndex]
            } else {
                prevLine = nil
            }
        }
        
        return commentsArray.reversed()
    }
}
