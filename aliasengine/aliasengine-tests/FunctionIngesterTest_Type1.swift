//
//  aliasengine_tests.swift
//  aliasengine-tests
//
//  Created by Gabriel Coman on 07/06/2022.
//

import XCTest

class FunctionIngesterTest_Type1: XCTestCase {

    let ingester = FunctionIngester(filePath: ".my_file")
    
    override func setUpWithError() throws {
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }
    
    func test_ShouldReturnNoItems_Given_AnEmptyFile() {
        let content = ""
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_AGarbageFile() {
        let content = "saakasljaskl\\\\\\/\'"
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_AFileWithoutFunctions() {
        let content = "echo \"ABC\""
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_AFileWithNoBrackets() {
        let content = "function test echo \"ABC\""
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnNoItems_GivenAFileWithNoClosingBrackets() {
        let content = "function hello_world { "
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_GivenAFileWithNoOpeningBrackets() {
        let content = "function hello_world ehco \"HELLO\" }"
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_AMultiLineFileWithImproperlyFormattedFunction() {
        let content = """
        function hello_world {
            echo "ABC"
            {
            }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_AOneLineFileWithIncorrectFunctionName() {
        let content = "function hello_world xoxox { } "
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_AMultiLineFileWithIncorrectFunctionName() {
        let content = """
        function hello_world xxxxx {
        }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnNoItems_Given_TheFunctionLineStartsWithGibberish() {
        let content = "xxx function { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = []
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunctionItem_Given_AOneLineFunction() {
        let content = "function test { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "test",
                      content: "function test { echo \"ABC\" }\n\ntest",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunctionItem_Given_OneLineFunctionStartsWithSpaces() {
        let content = "   function test { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "test",
                      content: "function test { echo \"ABC\" }\n\ntest",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunctionItem_Given_AMultiLineFunction() {
        let content = """
        function test {
            echo "ABC"
        }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "test",
                      content: "function test {\n    echo \"ABC\"\n}\n\ntest",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunctionItem_Given_AMultilineFunctionWithMultipleBrackets() {
        let content = """
        function hello_world {
            echo "Hello"
            {
                echo "World"
            }
        }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello_world",
                      content: "function hello_world {\n    echo \"Hello\"\n    {\n        echo \"World\"\n    }\n}\n\nhello_world",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnOnlyOutermostFunction_Given_ItHasNestedFunctions() {
        let content = """
        function hello {
            echo "Hello"
            function world {
                echo "World"
            }
        }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello {\n    echo \"Hello\"\n    function world {\n        echo \"World\"\n    }\n}\n\nhello",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunctionItem_Given_OneLineFunctionFromAVariedFile() {
        let content = """
        echo "ABC"

        alias lt='ls -t'

        function hello { echo "Hello" }

        ls -a
        """
        
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello { echo \"Hello\" }\n\nhello",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunction_GivenAVariedFile() {
        let content = """
        echo "ABC"

        alias lt='ls -t'

        function hello {
            echo "Hello"
            function world {
                echo "World"
            }
        }

        ls -a
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello {\n    echo \"Hello\"\n    function world {\n        echo \"World\"\n    }\n}\n\nhello",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnMultipleOneLineFunctionsItems_Given_TheyArePresent() {
        let content = """
        function hello { echo "HELLO" }
        function world { echo "WORLD" }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello { echo \"HELLO\" }\n\nhello",
                      path: ".my_file",
                      comments: []),
            IndexItem(type: .function,
                      name: "world",
                      content: "function world { echo \"WORLD\" }\n\nworld",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnMultipleMultilineFunctions_Given_TheyArePresent() {
        let content = """
        function hello {
            echo "HELLO"
        }

        function world {
            echo "WORLD"
            echo "!"
        }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello {\n    echo \"HELLO\"\n}\n\nhello",
                      path: ".my_file",
                      comments: []),
            IndexItem(type: .function,
                      name: "world",
                      content: "function world {\n    echo \"WORLD\"\n    echo \"!\"\n}\n\nworld",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnMultipleMultilineFunctions_ExcludingNestedFunctions() {
        let content = """
        function hello {
            echo "HELLO"
            function helper {
            }
        }

        function world {
            echo "WORLD"
            function helper {
            }
        }
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello {\n    echo \"HELLO\"\n    function helper {\n    }\n}\n\nhello",
                      path: ".my_file",
                      comments: []),
            IndexItem(type: .function,
                      name: "world",
                      content: "function world {\n    echo \"WORLD\"\n    function helper {\n    }\n}\n\nworld",
                      path: ".my_file",
                      comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnOnlyValidFunctions_Given_AMultiLineMultiCommandFunctionWithSomeSyntaxErrors() {
        let content = """
        echo "HELLP"

        alias lt='ls -a'

        function hello {
            echo "HELLO"

        function world {
            echo "WORLD"
        }

        echo "WORLD"
        """
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function,
                      name: "world",
                      content: "function world {\n    echo \"WORLD\"\n}\n\nworld",
                      path: ".my_file",
                      comments: []),
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnOneFunctionItem_Given_AFileWithComments() {
        let content = """
        # this is my
        # comment
        function hello {
            echo "HELLO"
        }
        """
        let result = ingester.process(fileContents: content)
        let expect = [
            IndexItem(type: .function,
                      name: "hello",
                      content: "function hello {\n    echo \"HELLO\"\n}\n\nhello",
                      path: ".my_file",
                      comments: ["this is my", "comment"])]
        XCTAssertEqual(result, expect)
    }
}
