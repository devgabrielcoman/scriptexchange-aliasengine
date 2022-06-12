//
//  ScreenManager.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

class ScreenManager {
    
    static func drawBottomBar(x: Int32, y: Int32, width: Int32, message: String) {
        Style().inversed {
            mvaddstr(y, x, String(repeating: " ", count: Int(width)))
            mvaddstr(y, x, message)
        }
    }

    static func drawSearchBar(x: Int32, y: Int32, query: String, cursorIndex: Int32) {
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
}
