//
//  DataReader.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

class DataReader: DataHandler {
    
    private let decoder = JSONDecoder()
    
    func load() -> [IndexItem] {
        guard let data = try? Data(contentsOf: getDataUrl()) else {
            print("No indexed items found.")
            return []
        }
        
        guard let items = try? decoder.decode([IndexItem].self, from: data) else {
            print("Could not parse data.")
            return []
        }
        
        return items
    }
}
