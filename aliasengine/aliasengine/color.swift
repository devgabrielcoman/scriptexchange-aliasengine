//
//  color.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

enum ColorModes: Int16 {
    case Cyan = 1
}

struct Color {
    
    static func setup() {
        start_color();
        init_pair(ColorModes.Cyan.rawValue, Int16(COLOR_CYAN), Int16(COLOR_BLACK))
    }
    
    static func cyan(commands: () -> Void) {
        attron(COLOR_PAIR(Int32(ColorModes.Cyan.rawValue)))
        commands()
        attroff(COLOR_PAIR(Int32(ColorModes.Cyan.rawValue)))
    }
}
