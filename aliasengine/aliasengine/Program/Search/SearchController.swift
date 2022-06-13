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
    private let data: [IndexItem]
    private var vLimit: Int = 0
    private var startFrom: Int = 0
    
    init(term: SearchTerm, initialData: [IndexItem]) {
        searchTerm = term
        data = SearchController.sorted(initialData)
    }
    
    func setVLimit(limit: Int32) {
        vLimit = Int(limit) - 3 /** to account for the search result title / magic number */
    }
    
    func search() {
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
            startFrom = 0
        }
    }
    
    func getResult() -> [IndexItem] {
        return boxed(searchResult)
    }
    
    func getTotalNumberOfSearchResults() -> Int {
        return searchResult.count
    }
    
    func getAllItemCount() -> Int {
        return data.count
    }
    
    private func getSearchResult() -> [IndexItem] {
        let term = searchTerm.toString()
        guard !term.isEmpty else { return data }
        return data.filter { item in
            return item.searchString.contains(term)
        }
    }
    
    func getSelectedItem() -> IndexItem? {
        guard let current = current() else { return  nil}
        guard searchResult.count > 0 else { return nil }
        guard current >= 0 && current < searchResult.count else { return nil }
        return searchResult[current]
    }
    
    func current() -> Int? {
        guard let item = selectedItem else { return nil }
        return item - startFrom
    }
    
    func moveDown() {
        guard var item = selectedItem else { return }
        item += 1
        if item > searchResult.count - 1 {
            item = searchResult.count - 1
        }
        selectedItem = item
        startFrom = max(item - vLimit, 0)
    }
    
    func moveUp() {
        guard var item = selectedItem else { return }
        item -= 1
        if item < 0 {
            item = 0
        }
        selectedItem = item
        startFrom = max(item - vLimit, 0)
    }
    
    private static func sorted(_ data: [IndexItem]) -> [IndexItem] {
        return data.sorted { l, r in
            if l.path == r.path {
                return l.name < r.name
            }
            return l.path < r.path
        }
    }
    
    private func boxed(_ data: [IndexItem]) -> [IndexItem] {
        var result: [IndexItem] = []
        
        for (i, item) in data.enumerated() {
            if i >= startFrom && i <= startFrom + vLimit {
                result.append(item)
            }
        }
        
        return result
    }
}
