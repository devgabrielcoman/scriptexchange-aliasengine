//
//  SafeShell.swift
//  aliasengine
//
//  Created by Gabriel Coman on 10/06/2022.
//

import Foundation

@discardableResult
func exec(_ path: String, _ args: String...) -> Int32 {
    let task = Process()
    task.launchPath = path
    task.arguments = args
    task.launch()
    task.waitUntilExit()
        
    return task.terminationStatus
}


func shell(_ command: String) -> String {
    let task = Process()
    let pipe = Pipe()
    
    task.standardOutput = pipe
    task.standardError = pipe
    task.arguments = ["-c", command]
    task.launchPath = "/bin/sh"
    task.standardInput = nil
    task.launch()
    
    let data = pipe.fileHandleForReading.readDataToEndOfFile()
    let output = String(data: data, encoding: .utf8)!
    
    return output
}
