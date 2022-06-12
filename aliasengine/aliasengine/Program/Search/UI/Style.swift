//
//  color.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

private let A_BOLD = NCURSES_BITS(1, 13)
private let A_DIM = NCURSES_BITS(1, 12)
private let A_REVERSED = NCURSES_BITS(1, 10)

func NCURSES_BITS(_ mask: UInt32, _ shift: UInt32) -> CInt {
    CInt(mask << (shift + UInt32(NCURSES_ATTR_SHIFT)))
}

enum ColorModes: Int32 {
    case Cyan = 1
    case Green = 2
    case Yellow = 3
    case Magenta = 4
}

struct Style {
    private var window: OpaquePointer? = nil
    
    init() {
        self.window = nil
    }
    
    init(_ window: OpaquePointer) {
        self.window = window
    }
    
    static func setup() {
        start_color();
        init_pair(Int16(ColorModes.Cyan.rawValue), Int16(COLOR_CYAN), Int16(COLOR_BLACK))
        init_pair(Int16(ColorModes.Green.rawValue), Int16(COLOR_GREEN), Int16(COLOR_BLACK))
        init_pair(Int16(ColorModes.Yellow.rawValue), Int16(COLOR_YELLOW), Int16(COLOR_BLACK))
        init_pair(Int16(ColorModes.Magenta.rawValue), Int16(COLOR_MAGENTA), Int16(COLOR_BLACK))
    }
    
    func cyan(commands: () -> Void) {
        let pair = COLOR_PAIR(ColorModes.Cyan.rawValue)
        
        if let window = window {
            wattron(window, pair)
        } else {
            attron(pair)
        }
        commands()
        if let window = window {
            wattroff(window, pair)
        } else {
            attroff(pair)
        }
    }
    
    func green(commands: () -> Void) {
        let pair = COLOR_PAIR(ColorModes.Green.rawValue)
        if let window = window {
            wattron(window, pair)
        } else {
            attron(pair)
        }
        commands()
        if let window = window {
            wattroff(window, pair)
        } else {
            attroff(pair)
        }
    }
    
    func magenta(commands: () -> Void) {
        let pair = COLOR_PAIR(ColorModes.Magenta.rawValue)
        if let window = window {
            wattron(window, pair)
        } else {
            attron(pair)
        }
        commands()
        if let window = window {
            wattroff(window, pair)
        } else {
            attroff(pair)
        }
    }
    
    func yellow(commands: () -> Void) {
        let pair = COLOR_PAIR(ColorModes.Yellow.rawValue)
        if let window = window {
            wattron(window, pair)
        } else {
            attron(pair)
        }
        commands()
        if let window = window {
            wattroff(window, pair)
        } else {
            attroff(pair)
        }
    }
    
    func bold(commands: () -> Void) {
        if let window = window {
            wattron(window, A_BOLD)
        } else {
            attron(A_BOLD)
        }
        commands()
        if let window = window {
            wattroff(window, A_BOLD)
        } else {
            attroff(A_BOLD)
        }
    }
    
    func dim(commands: () -> Void) {
        if let window = window {
            wattron(window, A_DIM)
        } else {
            attron(A_DIM)
        }
        commands()
        if let window = window {
            wattroff(window, A_DIM)
        } else {
            attroff(A_DIM)
        }
    }
    
    func inversed(commands: () -> Void) {
        if let window = window {
            wattron(window, A_REVERSED)
        } else {
            attron(A_REVERSED)
        }
        commands()
        if let window = window {
            wattroff(window, A_REVERSED)
        } else {
            attroff(A_REVERSED)
        }
    }
}
