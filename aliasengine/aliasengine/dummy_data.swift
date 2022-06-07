//
//  dummy_data.swift
//  aliasengine
//
//  Created by Liviu Coman on 04.06.2022.
//

import Foundation

let dummyData: [IndexItem] = [
    IndexItem(type: .alias,
              name: "lt",
              content: "du -sh * | sort -h",
              path: ".simple_aliases",
              comment: "sort file by size"),
    IndexItem(type: .alias,
              name: "mnt",
              content: "mount | grep -E ^/dev | column -t",
              path: ".simple_aliases", comment: "succint mnt"),
    IndexItem(type: .alias,
              name: "bash-history",
              content: "history|grep bash",
              path: ".simple_aliases",
              comment: "see all bash history"),
    IndexItem(type: .alias,
              name: "count-files",
              content: "find . -type f | wc -l",
              path: ".simple_aliases",
              comment: "returns file count"),
    IndexItem(type: .alias,
              name: "trash",
              content: "mv --force -t ~/.Trash",
              path: ".simple_aliases",
              comment: "move to trash"),
    IndexItem(type: .alias,
              name: "preview",
              content: "find . | fzf --preview 'bat --theme={} --color=always {}'",
              path: ".simple_aliases",
              comment: "find with preview"),
    IndexItem(type: .function,
              name: "cl",
              content: """
function cl() {
    DIR="$*";
        # if no DIR given, go home
        if [ $# -lt 1 ]; then
                DIR=$HOME;
    fi;
    builtin cd "${DIR}" && \
    # use your preferred ls command
        ls -F --color=auto
}
""",
              path: "functions",
              comment: "Change directories and view the contents at the same time"),
    IndexItem(type: .alias,
              name: "empty1",
              content: "",
              path: ".simple_aliases",
              comment: "find with preview"),
    IndexItem(type: .alias,
              name: "empty2",
              content: "",
              path: ".simple_aliases",
              comment: "find with preview"),
    IndexItem(type: .alias,
              name: "empty3",
              content: "",
              path: ".simple_aliases",
              comment: "find with preview"),
    IndexItem(type: .alias,
              name: "empty4",
              content: "",
              path: ".simple_aliases",
              comment: "find with preview"),
    IndexItem(type: .alias,
              name: "empty5",
              content: "",
              path: ".simple_aliases",
              comment: "find with preview"),
]
