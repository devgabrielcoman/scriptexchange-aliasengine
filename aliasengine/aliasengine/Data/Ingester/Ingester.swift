//
//  Ingester.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

public protocol Ingester {
    func process(fileContents contents: String) -> [IndexItem]
}
