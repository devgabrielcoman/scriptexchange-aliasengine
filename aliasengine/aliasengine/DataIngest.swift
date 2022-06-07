//
//  DataIngest.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

enum ScriptType: Int, Codable {
    case alias
    case function
    case script
}

struct IndexItem: Codable {
    let type: ScriptType
    let name: String
    let content: String
    let path: String
    let comment: String
}

protocol Ingester {
    func process(fileContents contents: String) -> [IndexItem]
}

protocol DataHandler {
    func getDataUrl() -> URL
}

extension DataHandler {
    func getDataUrl() -> URL {
        let homeUrl: URL = FileManager.default.homeDirectoryForCurrentUser
        return URL(fileURLWithPath: ".local/bin/scripthub/data.json", relativeTo: homeUrl)
    }
}

class DataIngest: DataHandler {
    
    func ingest(path: String) {
        guard let file = try? String(contentsOfFile: path) else {
            print("Could not open file \(path)")
            return
        }
        let aliasIngester = AliasIngester(filePath: path)
        let items = aliasIngester.process(fileContents: file)
        guard let jsonData = try? JSONEncoder().encode(items) else {
            print("Could not save aliases from file \(path)")
            return
        }
        
        let fullUrl = getDataUrl()
        
        guard let _ = try? jsonData.write(to: fullUrl) else {
            print("Could not save aliases from file \(path)")
            return
        }
        
        print("Ingested file \(path)")
        print("Found \(items.count) aliases")
        items.forEach { item in
            print("  Ingested \(item.name)")
        }
    }
}

class AliasIngester: Ingester {
    
    private let ALIAS_PREFIX = "alias "
    private let fileName: String
    
    init(filePath path: String) {
        fileName = path.fileName()
    }
    
    func process(fileContents contents: String) -> [IndexItem] {
        return contents.split(separator: "\n")
            .map(String.init)
            .filter { $0.starts(with: ALIAS_PREFIX) }
            .compactMap(getAlias)
    }
    
    /// alias lines should always have this format:
    ///  alias name='command'
    private func getAlias(fromLine line: String) -> IndexItem? {
        let aliasWithoutPrefix = line.deletingPrefix(ALIAS_PREFIX)
        let aliasComponents = aliasWithoutPrefix.split(separator: "=")
        guard aliasComponents.count >= 2 else { return nil }
        
        let aliasName = aliasComponents[0]
        let aliasCommand = aliasComponents[1..<aliasComponents.count].joined(separator: "")
        
        return IndexItem(type: .alias,
                        name: String(aliasName),
                         content: aliasCommand,
                         path: fileName,
                         comment: "No comment")
    }
}
