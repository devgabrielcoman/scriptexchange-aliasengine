//
//  IndexItem.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public enum ScriptType: Int, Codable {
    case alias
    case function
    case script
}

public struct IndexItem: Codable, Equatable {
    let type: ScriptType
    let name: String
    let content: String
    let path: String
    let comments: [String]
    
    var searchString: String {
        return "\(path)/\(name)"
    }
    
    func copy(withComments comments: [String]) -> IndexItem {
        return IndexItem(type: type, name: name, content: content, path: path, comments: comments)
    }
}
