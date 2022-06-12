//
//  FunctionIngesterTest_EdgeCases.swift
//  aliasengine-tests
//
//  Created by Liviu Coman on 12.06.2022.
//

import XCTest

class FunctionIngesterTest_EdgeCases: XCTestCase {
    
    let ingester = FunctionIngester(filePath: ".my_file")
    
    override func setUpWithError() throws {
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }
    
    func test_ShouldReturnACorrectFunction() {
        let content = """
        
        # Change directories and view the contents at the same time
        cl() {
            DIR="$*";
                # if no DIR given, go home
                if [ $# -lt 1 ]; then
                        DIR=$HOME;
            fi;
            builtin cd "${DIR}" && \
            # use your preferred ls command
                ls -F --color=auto
        }
        """
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function,
                      name: "cl",
                      content: "cl() {\n    DIR=\"$*\";\n        # if no DIR given, go home\n        if [ $# -lt 1 ]; then\n                DIR=$HOME;\n    fi;\n    builtin cd \"${DIR}\" &&     # use your preferred ls command\n        ls -F --color=auto\n}\n\ncl",
                      path: ".my_file",
                      comments: ["Change directories and view the contents at the same time"])
        ]
        XCTAssertEqual(result, expected)
    }
}
