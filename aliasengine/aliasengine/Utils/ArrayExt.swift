//
//  array_utils.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

extension Array {
    
    func limitTo(height: Int32) -> Self {
        let copy = self
        if copy.count > height {
            return Array(copy.prefix(upTo: Int(height)))
        }
        
        return copy
    }
    
    func unique<T:Hashable>(by: ((Element) -> (T)))  -> [Element] {
        var set = Set<T>() //the unique list kept in a Set for fast retrieval
        var arrayOrdered = [Element]() //keeping the unique list of elements but ordered
        for value in self {
            if !set.contains(by(value)) {
                set.insert(by(value))
                arrayOrdered.append(value)
            }
        }
        
        return arrayOrdered
    }
}

extension Collection {

    /// Returns the element at the specified index if it is within bounds, otherwise nil.
    subscript (safe index: Index) -> Element? {
        return indices.contains(index) ? self[index] : nil
    }
}
