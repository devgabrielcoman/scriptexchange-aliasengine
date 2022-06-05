//
//  search_term.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

class SearchTerm {
    private let PREFIX = "> "
    private var index = 0 // the length of PREFIX
    private var term: [Int32] = []
    
    func toString() -> String {
        return String(term.map { charCode in
            Character(UnicodeScalar(UInt8(charCode)))
        })
    }
    
    func add(code: Int32) {
        term.insert(code, at: index)
        moveIndexRight()
    }
    
    func toSearchQuery() -> String {
        return "> \(toString())"
    }
    
    func remove() {
        if (index <= 0) {
            return
        }
        term.remove(at: index - 1)
        moveIndexLeft()
    }
    
    func getIndex() -> Int32 {
        return Int32(index) + 2
    }
    
    func moveIndexLeft() {
        index -= 1
        if (index < 0) {
            index = 0
        }
    }
    
    func moveIndexRight() {
        index += 1
        if (index > term.count) {
            index = term.count
        }
    }
}
