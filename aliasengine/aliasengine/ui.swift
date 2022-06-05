//
//  ui.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

func drawBottomBar(x: Int32, y: Int32, width: Int32, message: String) {
    attron(reversed)
    mvaddstr(y, x, String(repeating: " ", count: Int(width)))
    mvaddstr(y, x, message)
    attroff(reversed)
}

func drawSearchBar(x: Int32, y: Int32, query: String, cursorIndex: Int32) {
    move(y, x)
    addstr(query)
    move(y, cursorIndex)
}

class WindowManager {
    internal let window: OpaquePointer
    private var x: Int32 = 0
    private var y: Int32 = 0
    private var width: Int32 = 0
    private var height: Int32 = 0
    
    init(window: OpaquePointer) {
        self.window = window
    }
    
    func setPosition(x: Int32, y: Int32, width: Int32, height: Int32) {
        werase(window)
        wresize(window, height, width)
        mvwin(window, y, x)
        self.x = x
        self.y = y
        self.width = width
        self.height = height
    }
    
    func drawBox() {
        box(self.window, 0, 0)
    }
    
    func drawTitle(title: String) {
        wattron(window, reversed)
        let fullTitle = " \(title) "
        mvwaddstr(window, 0, (width / 2) - Int32(fullTitle.count / 2), fullTitle)
        wattroff(window, reversed)
    }
}

extension WindowManager {
    func drawSearchResults(results: [IndexItem], selectedIndex: Int?) {
        let left: Int32 = 1
        let top: Int32 = 1
        
        wmove(window, top, left)
        waddstr(window, "Found \(results.count) results")
        
        for (i, result) in results.enumerated() {
            wmove(window, top + Int32(i) + 1, left)
            if let current = selectedIndex, i == current {
                wattron(window, reversed)
            }
            waddstr(window, result.name)
            wattroff(window, reversed)
        }
    }
    
    func drawResultPreview(selectedItem: IndexItem?) {
        guard let result = selectedItem else { return }
        let content = result.content
        let contentArray = content.split(separator: "\n").map { sub in
            return "\(sub)"
        }

        for (i, contentLine) in contentArray.enumerated() {
            mvwaddstr(window, 1 + Int32(i), 2, contentLine)
        }
    }
}
