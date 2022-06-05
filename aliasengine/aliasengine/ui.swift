//
//  ui.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

func drawBottomBar(x: Int32, y: Int32, width: Int32, message: String) {
    Style().inversed {
        mvaddstr(y, x, String(repeating: " ", count: Int(width)))
        mvaddstr(y, x, message)
    }
}

func drawSearchBar(x: Int32, y: Int32, query: String, cursorIndex: Int32) {
    let searchPrefix = "> "
    let prefixLen = Int32(searchPrefix.count)
    move(y, x)
    clrtoeol();
    Style().cyan {
        addstr(searchPrefix)
    }
    move(y, x + prefixLen)
    Style().bold {
        addstr(query)
    }
    move(y, cursorIndex + prefixLen)
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
        let fullTitle = " \(title) "
        Style(window).inversed {
            mvwaddstr(window, 0, (width / 2) - Int32(fullTitle.count / 2), fullTitle)
        }
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
                Style(window).cyan {
                    Style(window).inversed {
                        waddstr(window, result.name)
                    }
                }
                
            } else {
                Style(window).cyan {
                    waddstr(window, result.name)
                }
            }
        }
    }
    
    func drawResultPreview(selectedItem: IndexItem?) {
        guard let result = selectedItem else { return }
        
        // draw comments
        let comments = "# \(result.comment ?? "No comment")"
        Style(window).green {
            mvwaddstr(window, 1, 2, comments)
        }
        
        // draw content
        let content = result.content
        let contentArray = content.split(separator: "\n").map { sub in
            return "\(sub)"
        }

        for (i, contentLine) in contentArray.enumerated() {
            mvwaddstr(window, Int32(i) + 2, 2, contentLine)
        }
    }
}
