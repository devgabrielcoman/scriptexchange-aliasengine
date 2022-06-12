//
//  SourceFile.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public enum SourceType: Int, Codable {
    case alias
    case script
}

struct SourceFile: Codable, Equatable {
    let path: String
    let name: String
    let type: SourceType
}
