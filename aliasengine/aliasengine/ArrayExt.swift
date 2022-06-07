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
}
