//
//  main.swift
//  aliasengine
//
//  Created by Liviu Coman on 04.06.2022.
//

import Foundation
import Darwin.ncurses
import ArgumentParser

struct Arguments: ParsableArguments {
    @Option(name: [.customShort("i"), .long], help: "A path to an alias file to ingest")
    var ingest: String?
}

///
/// THE INGESTION PART OF THE PROGRAM
///

let arguments = Arguments.parseOrExit()
if let filePath = arguments.ingest {
    let ingester = DataIngest()
    ingester.ingest(path: filePath)
    exit(0)
}

///
/// THE INTERACTIVE PART OF THE PROGRAM
///

initscr()                   // Init window. Must be first
cbreak()
noecho()                    // Don't echo user input
nonl()                      // Disable newline mode
intrflush(stdscr, true)     // Prevent flush
keypad(stdscr, true)        // Enable function and arrow keys
curs_set(1)                 // Set cursor to invisible
defer { endwin() }
Style.setup()               // setup colors

private let resultsWindow = newwin(0, 0, 0, 0)!
private let previewWindow = newwin(0, 0, 0, 0)!

private let reader = DataReader()
private let data = reader.load()
private let searchTerm = SearchTerm()
private let controller = SearchController(term: searchTerm, initialData: data)
private let resultsWindowManager = WindowManager(window: resultsWindow)
private let previewWindowManager = WindowManager(window: previewWindow)

var quit = false;
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
    drawBottomBar(x: 0, y: height - 1, width: width, message: "Type to search all indexed aliases, functions and scripts.")
    
    // draw results window
    resultsWindowManager.setPosition(x: 0, y: 0, width: width / 2, height: height - 2)
    resultsWindowManager.drawSearchResults(results: controller.getResult(), selectedIndex: controller.current(), total: controller.getTotalNumberOfResults())
    
    // draw preview window
    previewWindowManager.setPosition(x: width / 2, y: 0, width:  width / 2, height: height - 2)
    previewWindowManager.drawBox()
    previewWindowManager.drawResultPreview(selectedItem: controller.getSelectedItem())
    
    // refresh everything
    refresh()
    wrefresh(resultsWindow)
    wrefresh(previewWindow)
    
    // display the search bar
    drawSearchBar(x: 0, y: height - 2, query: searchTerm.toString(), cursorIndex: searchTerm.getIndex())
    
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
    default:
        break
    }
}

delwin(resultsWindow)
delwin(previewWindow)
endwin()
exit(EX_OK)
