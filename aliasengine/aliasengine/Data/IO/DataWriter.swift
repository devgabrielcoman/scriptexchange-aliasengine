//
//  DataWriter.swift
//  aliasengine
//
//  Created by Gabriel Coman on 07/06/2022.
//

import Foundation

class DataWriter: DataHandler {
    
    private let encoder = JSONEncoder()
    
    func write(items: [IndexItem]) {
        guard let jsonData = try? JSONEncoder().encode(items) else {
            print("\("Could not save aliases.", color: .red)")
            return
        }
        
        let fullUrl = getDataUrl()
        
        guard let _ = try? jsonData.write(to: fullUrl) else {
            print("\("Could not save aliases to file", color: .red)")
            return
        }
    }
    
    func write(sources: [SourceFile]) {
        guard let jsonData = try? JSONEncoder().encode(sources) else {
            print("\("Could not save sources", color: .red)")
            return
        }
        
        let fullUrl = getSourcesUrl()
        
        guard let _ = try? jsonData.write(to: fullUrl) else {
            print("\("Could not save source to file", color: .red)")
            return
        }
    }
    
    func write(command: String) {
        let fullUrl = getLastCommandUrl()
        guard let _ = try? command.write(to: fullUrl, atomically: true, encoding: .utf8) else {
            print("\("Could not write last command", color: .red)")
            return
        }
    }
}
