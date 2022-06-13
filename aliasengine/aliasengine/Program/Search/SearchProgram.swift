//
//  SearchProgram.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation
import Darwin.ncurses
import AppKit

class SearchProgram: Program {
    
    func run() {
        initscr()                   // Init window. Must be first
        cbreak()
        noecho()                    // Don't echo user input
        nonl()                      // Disable newline mode
        intrflush(stdscr, true)     // Prevent flush
        keypad(stdscr, true)        // Enable function and arrow keys
        curs_set(1)                 // Set cursor to invisible
        defer { endwin() }
        Style.setup()               // setup colors

        let resultsWindow = newwin(0, 0, 0, 0)!
        let previewWindow = newwin(0, 0, 0, 0)!

        let reader = DataReader()
        let data = reader.readItems()
        let searchTerm = SearchTerm()
        let controller = SearchController(term: searchTerm, initialData: data)
        let resultsWindowManager = WindowManager(window: resultsWindow)
        let previewWindowManager = WindowManager(window: previewWindow)

        var exitCommand: String? = nil
        var quit = false
        while !quit {

            // first draw of the frame
            refresh()

            // Read the environment
            let width = COLS
            let height = LINES

            // do controller actions
            controller.setVLimit(limit: height - 2)
            controller.search()

            // Bottom bar
            ScreenManager.drawBottomBar(x: 0, y: height - 1, width: width, message: "Type to search all indexed aliases, functions and scripts.")

            // draw results window
            resultsWindowManager.setPosition(x: 0, y: 0, width: width / 2, height: height - 2)
            resultsWindowManager.drawSearchResults(results: controller.getResult(), selectedIndex: controller.current(), total: controller.getTotalNumberOfSearchResults(), full: controller.getAllItemCount())

            // draw preview window
            previewWindowManager.setPosition(x: width / 2, y: 0, width:  width / 2, height: height - 2)
            previewWindowManager.drawBox()
            previewWindowManager.drawResultPreview(selectedItem: controller.getSelectedItem())

            // refresh everything
            refresh()
            wrefresh(resultsWindow)
            wrefresh(previewWindow)

            // display the search bar
            ScreenManager.drawSearchBar(x: 0, y: height - 2, query: searchTerm.toString(), cursorIndex: searchTerm.getIndex())

            // do some actions
            let input = getch()

            switch input {
            case 127, 8: // delete & backspace
                searchTerm.remove()
            case KEY_LEFT:
                searchTerm.moveIndexLeft()
            case KEY_RIGHT:
                searchTerm.moveIndexRight()
            case KEY_UP:
                controller.moveUp()
            case KEY_DOWN:
                controller.moveDown()
            case 32...126:
                searchTerm.add(code: input)
            case 13: // enter
                if let item = controller.getSelectedItem() {
                    exitCommand = item.content
                }
                quit = true
            default:
                break
            }
        }

        delwin(resultsWindow)
        delwin(previewWindow)
        endwin()

        if let command = exitCommand {
            exec("/bin/sh", "-c", "eval \"\(command)\"")
        }

        exit(EX_OK)
    }
}
