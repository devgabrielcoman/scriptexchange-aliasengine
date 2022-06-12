//
//  FunctionIngesterTest_Type2.swift
//  aliasengine-tests
//
//  Created by Gabriel Coman on 08/06/2022.
//

import XCTest

class FunctionIngesterTest_Type2: XCTestCase {

    let ingester = FunctionIngester(filePath: ".my_file")
    
    override func setUpWithError() throws {
        // Put setup code here. This method is called before the invocation of each test method in the class.
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }
    
    func test_ShouldReturnNoItems_Given_AFileWithNoFunctions() {
        let content = "ls -a"
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnNoItems_Given_AFileWithGibberish() {
        let content = "sasassa////\\\\"
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnNoItems_Given_AFileWithInvalidFunctionName() {
        let content = "hello world () { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnNoItems_GivenAFileWithNoBrackets() {
        let content = "test() echo \"ABC\""
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnNoItems_GivenAFileWithNoOpeningBrackets() {
        let content = "test() echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnNoItems_Given_AFileWithNoClosingBrackets() {
        let content = "test() echo { \"ABC\""
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = []
        XCTAssertEqual(result, expected)
    }
    
//    func test_ShouldReturnNoItems_Given_AFileWithInvalidFunctionParentheses() {
//        let content = "test() () { echo \"ABC\" }"
//        let result = ingester.process(fileContents: content)
//        let expected: [IndexItem] = []
//        XCTAssertEqual(result, expected)
//    }
    
    func test_ShouldReturnAFunctionTypeItem_Given_AOneLineFile() {
        let content = "test() { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "test", content: "test() { echo \"ABC\" }\n\ntest", path: ".my_file", comments: [])
        ]
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnAFunctionItem_Given_OneLineFunctionStartsWithSpaces() {
        let content = "   test() { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expect: [IndexItem] = [
            IndexItem(type: .function, name: "test", content: "test() { echo \"ABC\" }\n\ntest", path: ".my_file", comments: [])
        ]
        XCTAssertEqual(result, expect)
    }
    
    func test_ShouldReturnAFunctionTypeItem_Given_AOneLineFileWithSpacesBeforeParentheses() {
        let content = "test () { echo \"ABC\" }"
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "test", content: "test () { echo \"ABC\" }\n\ntest", path: ".my_file", comments: [])
        ]
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnAFunctionTypeItem_Given_AMultiLineFile() {
        let content = """
            test() {
                echo "ABC"
            }
        """
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "test", content: "test() {\n        echo \"ABC\"\n    }\n\ntest", path: ".my_file", comments: [])
        ]
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnTwoFunctionItems_Given_AFileWithMultipleOneLineFunctions() {
        let content = """
        hello() { echo "HELLO" }
        world() { echo "WORLD" }
        """
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "hello", content: "hello() { echo \"HELLO\" }\n\nhello", path: ".my_file", comments: []),
            IndexItem(type: .function, name: "world", content: "world() { echo \"WORLD\" }\n\nworld", path: ".my_file", comments: [])
        ]
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnOneFunctionItem_Given_AFileWithNestedFunctions() {
        let content = """
        hello() {
            echo "HELLO"
            world () {
            }
        }
        """
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "hello", content: "hello() {\n    echo \"HELLO\"\n    world () {\n    }\n}\n\nhello", path: ".my_file", comments: []),
        ]
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnTwoFunctionItems_Given_AFileWithMultipleMultiLineFunctions() {
        let content = """
        ls -a
        echo "ABC"
        hello() {
            echo "HELLO"
            function foo {
                echo "FOO"
            }
        }
        
        echo "I am here"
        world() {
            echo "WORLD"
        }
        
        lt
        """
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "hello", content: "hello() {\n    echo \"HELLO\"\n    function foo {\n        echo \"FOO\"\n    }\n}\n\nhello", path: ".my_file", comments: []),
            IndexItem(type: .function, name: "world", content: "world() {\n    echo \"WORLD\"\n}\n\nworld", path: ".my_file", comments: [])
        ]
        XCTAssertEqual(result, expected)
    }
    
    func test_ShouldReturnOneFunctionItemWithComments_GivenAFileWithComments() {
        let content = """
        # this is my first
        # comment
        hello() {
            echo "World"
        }
        """
        let result = ingester.process(fileContents: content)
        let expected: [IndexItem] = [
            IndexItem(type: .function, name: "hello", content: "hello() {\n    echo \"World\"\n}\n\nhello", path: ".my_file", comments: ["this is my first", "comment"])
        ]
        XCTAssertEqual(result, expected)
    }
}
