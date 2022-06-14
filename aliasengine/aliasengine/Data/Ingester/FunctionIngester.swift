//
//  FunctionIngester.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public class FunctionIngester: FileIngester {
    
    private let FUNCTION_KEYWORD_ONE = "function"
    private let FUNCTION_KEYWORD_TWO = "()"
    private let OPEN_BRACKET = Character("{")
    private let CLOSE_BRACKET = Character("}")
    private let OPEN_PARA = Character("(")
    private let CLOSE_PARA = Character(")")
    private let SEPARATOR = Character(" ")
    private let fileName: String
    private let diskPath: String
    
    init(filePath path: String) {
        diskPath = path
        fileName = path.fileName
    }
    
    override public func process(fileContents contents: String) -> [IndexItem] {
        let lines = contents.split(separator: "\n").map(String.init)
        var result: [IndexItem] = []
        
        var i = -1
        let totalLength = lines.count
        
        while i < totalLength - 1 {
            i += 1
            let line = lines[i]
            let trimmedLine = line.trimmingCharacters(in: .whitespaces)
            
            // found potential function in first style
            if trimmedLine.starts(with: FUNCTION_KEYWORD_ONE) {
                let resultPair = processStyleOne(startLine: line, startIndex: i, allLines: lines)
                if let function = resultPair.item {
                    let comments = getCommentsForAlias(startIndex: i, lines: lines)
                    result.append(function.copy(withComments: comments))
                    i += resultPair.progress
                }
            }
            
            // found potential function in second style
            if trimmedLine.contains(FUNCTION_KEYWORD_TWO) {
                let resultPair = processStyleTwo(startLine: line, startIndex: i, allLines: lines)
                if let function = resultPair.item {
                    let comments = getCommentsForAlias(startIndex: i, lines: lines)
                    result.append(function.copy(withComments: comments))
                    i += resultPair.progress
                }
            }
        }
        
        return result
    }
    
    private func processStyleTwo(startLine line: String, startIndex: Int, allLines: [String]) -> (item: IndexItem?, progress: Int) {
        // prepare this data
        let preparedLine = line.trimmingCharacters(in: .whitespaces)
        
        var functionName: String = ""
        var hasSeenFirstPara = false
        var parenthesesNumber = 0
        var openBrackets = 0
        var hasSeenFirstBracket = false
        
        var characterArray = Array(preparedLine)
        var nextIndex = startIndex
        var allContent = ""
        var i = -1
        var totalLength = characterArray.count
        
        while i < totalLength - 1 {
            i += 1
            let nextChar = characterArray[i]
            allContent.append(nextChar)
            
            if nextChar == OPEN_PARA {
                hasSeenFirstPara = true
                parenthesesNumber += 1
            }
            
            if nextChar == CLOSE_PARA {
                parenthesesNumber += 1
            }
            
            if !hasSeenFirstPara {
                functionName.append(nextChar)
            }
            
            if nextChar == OPEN_BRACKET {
                hasSeenFirstBracket = true
                openBrackets += 1
            }
            
            if nextChar == CLOSE_BRACKET {
                openBrackets -= 1
            }
            
            // if we're at the end AND we still haven't closed the function
            if i == totalLength - 1 && openBrackets != 0 {
                
                nextIndex += 1
                if let nextLine = allLines[safe: nextIndex] {
                    let preparedNextLine = nextLine.trimmingCharacters(in: .newlines)
                    characterArray += Array(preparedNextLine)
                    allContent.append("\n")
                    totalLength += preparedNextLine.count
                }
            }
        }
        
        let name = functionName.trimmingCharacters(in: .whitespaces)
        
        guard parenthesesNumber >= 2, openBrackets == 0, hasSeenFirstBracket, name.split(separator: " ").count == 1 else {
            return (nil, nextIndex)
        }
        
        let progress = nextIndex - startIndex
        
        return (
            IndexItem(type: .function,
                      name: name,
                      content: allContent,
                      path: fileName,
                      comments: [],
                      pathOnDisk: diskPath),
            progress
        )
    }
    
    private func processStyleOne(startLine line: String, startIndex: Int, allLines: [String]) -> (item: IndexItem?, progress: Int) {
        // prepare this data
        let preparedLine = line.replacingOccurrences(of: FUNCTION_KEYWORD_ONE, with: "").trimmingCharacters(in: .whitespaces)
        
        var functionName: String? = nil
        var hasSeenFirstBracket = false
        var openBrackets = 0
        
        var characterArray = Array(preparedLine)
        var nextIndex = startIndex
        var allContent = ""
        var i = -1
        var totalLength = characterArray.count
        
        while i < totalLength - 1 {
            i += 1
            let nextChar = characterArray[i]
            allContent.append(nextChar)
            
            if nextChar == OPEN_BRACKET {
                // get the function name correctly
                if !hasSeenFirstBracket {
                    let potentialFunctionName = characterArray[0..<i].split(separator: SEPARATOR)
                    if potentialFunctionName.count == 1 {
                        functionName = String(potentialFunctionName[0])
                    } else {
                        return (nil, nextIndex)
                    }
                }
                
                hasSeenFirstBracket = true
                openBrackets += 1
            }
            
            if nextChar == CLOSE_BRACKET {
                openBrackets -= 1
            }
            
            // if we're at the end AND we still haven't closed the function
            if i == totalLength - 1 && openBrackets != 0 {
                nextIndex += 1
                if let nextLine = allLines[safe: nextIndex] {
                    let preparedNextLine = nextLine.trimmingCharacters(in: .newlines)
                    characterArray += Array(preparedNextLine)
                    allContent.append("\n")
                    totalLength += preparedNextLine.count
                }
            }
        }
        
        guard let name = functionName, openBrackets == 0, hasSeenFirstBracket else {
            return (nil, nextIndex)
        }
        
        let progress = nextIndex - startIndex
        
        return (
            IndexItem(type: .function,
                      name: name,
                      content: "\(FUNCTION_KEYWORD_ONE) \(allContent)",
                      path: fileName,
                      comments: [],
                      pathOnDisk: diskPath),
            progress
        )
    }
}
