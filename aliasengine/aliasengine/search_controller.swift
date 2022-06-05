//
//  search_controller.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

class SearchController {
    private var selectedItem: Int? = nil
    private let searchTerm: SearchTerm
    private var searchResult: [IndexItem] = []
    
    init(term: SearchTerm) {
        searchTerm = term
    }
    
    func search() -> [IndexItem] {
        let term = searchTerm.toString()
        searchResult = dummyData.filter { item in
            return item.name.contains(term)
        }
        if selectedItem == nil && searchResult.count > 0 {
            selectedItem = 0
        }
        else if searchResult.count == 0 {
            selectedItem = nil
        }
        return searchResult
    }
    
    func getSelectedItem() -> IndexItem? {
        guard let current = current() else { return  nil}
        guard searchResult.count > 0 else { return nil }
        guard current >= 0 && current < searchResult.count else { return nil }
        return searchResult[current]
    }
    
    func current() -> Int? {
        return selectedItem
    }
    
    func moveDown() {
        guard var item = selectedItem else { return }
        item += 1
        if item > searchResult.count - 1 {
            item = searchResult.count - 1
        }
        selectedItem = item
    }
    
    func moveUp() {
        guard var item = selectedItem else { return }
        item -= 1
        if item < 0 {
            item = 0
        }
        selectedItem = item
    }
}
