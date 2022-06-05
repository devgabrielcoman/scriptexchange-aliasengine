//
//  utils.swift
//  aliasengine
//
//  Created by Liviu Coman on 05.06.2022.
//

import Foundation

extension String {
    var ascii32: Int32 {
        Int32(Character(self).asciiValue!)
    }
}


let reversed = NCURSES_BITS(1, 10)

func NCURSES_BITS(_ mask: UInt32, _ shift: UInt32) -> CInt {
    CInt(mask << (shift + UInt32(NCURSES_ATTR_SHIFT)))
}
