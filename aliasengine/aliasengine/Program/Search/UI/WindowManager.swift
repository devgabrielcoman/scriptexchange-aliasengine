//
//  WindowManager.swift
//  aliasengine
//
//  Created by Gabriel Coman on 09/06/2022.
//

import Foundation

class WindowManager {
    internal let window: OpaquePointer
    private var x: Int32 = 0
    private var y: Int32 = 0
    internal var width: Int32 = 0
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
    func drawSearchResults(results: [IndexItem], selectedIndex: Int?, total: Int, full: Int) {
        let left: Int32 = 1
        let top: Int32 = 1
        
        wmove(window, top, left)
        waddstr(window, "Showing ")
        Style(window).bold {
            waddstr(window, "\(total)/\(full)")
        }
        waddstr(window, " results")
        
        for (i, result) in results.enumerated() {
            wmove(window, top + Int32(i) + 1, left)
            if let current = selectedIndex, i == current {
                Style(window).inversed {
                    waddstr(window, "\(result.path)/\(result.name)")
                }
            } else {
                waddstr(window, "\(result.path)/")
                switch result.type {
                case .alias:
                    Style(window).cyan {
                        waddstr(window, result.name)
                    }
                case .function:
                    Style(window).yellow {
                        waddstr(window, result.name)
                    }
                case .script:
                    Style(window).magenta {
                        waddstr(window, result.name)
                    }
                }
            }
        }
    }
    
    func drawResultPreview(selectedItem: IndexItem?) {
        guard let result = selectedItem else { return }
        var vlimit = height - 2
        
        // draw comments
        var t = 0
        let comments = result.comments.count > 0 ? result.comments : ["No comments"]
        for (i, commentLine) in comments.enumerated() {
            for (j, fitCommentLine) in commentLine.splitWordLines(thatFitIn: width).enumerated() {
                Style(window).green {
                    t = i + j + 1
                    mvwaddstr(window, Int32(t), 2, fitCommentLine)
                }
            }
        }
        // update vlimit
        vlimit -= Int32(t)
        
        // draw content
        let content = result.content
        let contentArray = content.split(separator: "\n").map { line in
            return "\(line)".limitTo(width: width)
        }.limitTo(height: vlimit)
        
        for (i, contentLine) in contentArray.enumerated() {
            mvwaddstr(window, Int32(i) + Int32(t) + 1, 2, contentLine)
        }
    }
    
//    func drawResults(selectedItem: IndexItem?) {
//        guard let result = selectedItem else { return }
//        
//        let res = try? shell("bat /Users/gabriel.coman/Workspace/my-scripts/setup_docker.sh --paging=never --color=always")
//        
////        let task = Process()
////        let pipe = Pipe()
////
////        task.standardOutput = pipe
////        task.standardError = pipe
////        task.arguments = ["-i", "-c", "bat /Users/gabriel.coman/Workspace/my-scripts/setup_docker.sh --paging=never"]
////        task.launchPath = "/bin/sh"
////        task.standardInput = nil
////        task.launch()
////
////        let data = pipe.fileHandleForReading.readDataToEndOfFile()
////        let output = String(data: data, encoding: .utf8)!
//        
////        return output
//        
//        if let res = res {
//            let t = 0
//            let contentArray = res.split(separator: "\n")
//            for (i, contentLine) in contentArray.enumerated() {
////                mvwaddstr(window, Int32(i) + Int32(t) + 1, 2, "\(contentLine)")
//                wmove(window, Int32(i) + Int32(t) + 1, 2)
//                let args: [CVarArg] = []
////                vw_printw(window, "\(contentLine)", getVaList(args))
//                start_color();
//                vwprintw(window, "\(contentLine)", getVaList(args))
//            }
////            wmove(window, 0, 0)
////            waddstr(window, "\(res.split(separator: "\n").count)")
////            let args: [CVarArg] = []
////            vw_printw(window, res, getVaList(args))
////            vwprintf(window, res)
//        }
//
//////        let res = try? exec("/bin/sh", "-c", "bat /Users/gabriel.coman/Workspace/my-scripts/setup_docker.sh --paging=never")
//////        if let res = res {
//////            let pipe = Pipe()
//////            wmove(window, 0, 0)
//////            waddstr(window, "ABC")
//////        }
//////        print(res!)
//    }
}
