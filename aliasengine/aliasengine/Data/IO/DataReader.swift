//
//  DataReader.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

class DataReader: DataHandler {
    
    private let decoder = JSONDecoder()
    
    func readItems() -> [IndexItem] {
        guard let data = try? Data(contentsOf: getDataUrl()) else {
            print("\("No indexed items found.", color: .yellow)")
            return []
        }
        
        guard let items = try? decoder.decode([IndexItem].self, from: data) else {
            print("\("Could not parse index items.", color: .red)")
            return []
        }
        
        return items
    }
    
    func readSources() -> [SourceFile] {
        guard let data = try? Data(contentsOf: getSourcesUrl()) else {
            print("\("No sources registered", color: .yellow)")
            return []
        }
        
        guard let sources = try? decoder.decode([SourceFile].self, from: data) else {
            print("\("Could not parse sources.", color: .red)")
            return []
        }
        
        return sources
    }
}
