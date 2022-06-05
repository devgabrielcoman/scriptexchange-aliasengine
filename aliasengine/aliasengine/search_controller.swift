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
        searchResult = getSearchResult()
        
        // when we have no selected item but we have results,
        // set the selected item to 0
        if selectedItem == nil && searchResult.count > 0 {
            selectedItem = 0
        }
        // when we have no results, set the item to nil
        else if searchResult.count == 0 {
            selectedItem = nil
        }
        // when we have new serch results and the item
        // is out of bounds, set it to 0
        if let term = selectedItem, term < 0 || term > searchResult.count {
            selectedItem = 0
        }
        
        return searchResult
    }
    
    private func getSearchResult() -> [IndexItem] {
        let term = searchTerm.toString()
        guard !term.isEmpty else { return dummyData }
        return dummyData.filter { item in
            return item.name.contains(term)
        }
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
