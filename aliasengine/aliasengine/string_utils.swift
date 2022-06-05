//
//  string_utils.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

extension String {
    
    func limitTo(width: Int32) -> String {
        let limit = Int(width) - 3
        var copy = self
        var len = copy.count
        
        while len > limit {
            copy.removeLast()
            len = copy.count
        }
        
        return copy
    }
    
    func splitWordLines(thatFitIn width: Int32) -> [String] {
        let words = self.replacingOccurrences(of: "\n", with: " ")
            .split(separator: " ")
            .map { e in
                return "\(e)"
            }
        
        var result: [String] = []
        var currentLine = "# "
        let limit = Int(width) - 3

        for word in words {
            let nextWord = "\(word) "
            let nextLimit = currentLine.count + nextWord.count
            if nextLimit < limit {
                currentLine += nextWord
            } else {
                result.append(currentLine)
                currentLine = "# "
                currentLine += nextWord
            }
        }
        
        result.append(currentLine)
        
        return result
    }
}
