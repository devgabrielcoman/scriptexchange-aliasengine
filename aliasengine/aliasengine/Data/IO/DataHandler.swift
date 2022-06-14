//
//  DataHandler.swift
//  aliasengine
//
//  Created by Gabriel Coman on 07/06/2022.
//

import Foundation

protocol DataHandler {
    func getDataUrl() -> URL
}

extension DataHandler {
    
    func getHomeUrl() -> URL {
        return FileManager.default.homeDirectoryForCurrentUser
    }
    
    func getDataUrl() -> URL {
        return URL(fileURLWithPath: ".local/bin/scripthub/data.json", relativeTo: getHomeUrl())
    }
    
    func getSourcesUrl() -> URL {
        return URL(fileURLWithPath: ".local/bin/scripthub/sources.json", relativeTo: getHomeUrl())
    }
    
    func getLastCommandUrl() -> URL {
        return URL(fileURLWithPath: ".local/bin/scripthub/lastcommand", relativeTo: getHomeUrl())
    }
}
