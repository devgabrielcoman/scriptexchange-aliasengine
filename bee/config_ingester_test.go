package main

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ingester = ConfigIngester{filePath: "test.sh", currentTime: 0}

func TestConfigIngester_Process(t *testing.T) {
	t.Run("it should return an empty slice given an empty input", func(t *testing.T) {
		var content = ""
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given garbage input", func(t *testing.T) {
		var content = "\\\asa/////"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given incorrectly formatted alias", func(t *testing.T) {
		var content = "alias test'ls -all'"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given incorrectly formatted alias", func(t *testing.T) {
		var content = "alias test 'ls -all'"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should a slice with one alias even if format is slighlty incorrect", func(t *testing.T) {
		var content = "alias test = 'ls -all'"
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one alias in one line content", func(t *testing.T) {
		var content = "alias test='ls -all'"
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should not return a slice with alises if it actually starts with a comment", func(t *testing.T) {
		var content = `# this is my
		# write down like: alias test='ls -all'
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one alias and comments in multi line content", func(t *testing.T) {
		var content = `# this is my
		# comment
		alias test='ls -all'
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my", "# comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two aliases in multi line content", func(t *testing.T) {
		var content = `
		# this is my comment
		alias test1='ls -all'

		# this is the
		# second comment
		alias test2='run this'
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
			{
				Name:       "test2",
				Content:    "run this",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two aliases even if the format is slightly incorrect", func(t *testing.T) {
		var content = `
		# this is my comment
		alias test1 = 'ls -all'

		# this is the
		# second comment
		alias test2 = run this
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
			{
				Name:       "test2",
				Content:    "run this",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without keywords", func(t *testing.T) {
		var content = "{ echo \"abc\"; }"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without body of type one", func(t *testing.T) {
		var content = "function abc"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without name of type one", func(t *testing.T) {
		var content = "function { echo \"abc\" }"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one function given one liner with correct function of type one", func(t *testing.T) {
		var content = "function test { echo \"hello world\"; }"
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with a function and comments given one liner with correct function of type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test { echo "hello world"; }
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two functions and comments given content with multiple one line functions of type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test1 { echo "hello"; }

		# this is my second function
		function test2 { echo "world"; }
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "function test1 { echo \"hello\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
			{
				Name:       "test2",
				Content:    "function test2 { echo \"world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with a function and comments given a multi line function of type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test { 
			echo "hello world"; 
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test {\n\t\t\techo \"hello world\"; \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice created from nested functions if type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test { 
			function world { echo "world"'; }
			echo "hello";
			world 
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test {\n\t\t\tfunction world { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice created with multiple nested functions of type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test { 
			eval "${command}"
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test {\n\t\t\teval \"${command}\"\n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice created from nested functions of type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test1 { 
			function world { echo "world"'; }
			echo "hello";
			world 
		}

		# this is my second function
		function test2 {
			function world { echo "hello world"; }
			world 
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "function test1 {\n\t\t\tfunction world { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
			{
				Name:       "test2",
				Content:    "function test2 {\n\t\t\tfunction world { echo \"hello world\"; }\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without body of type two", func(t *testing.T) {
		var content = "abc()"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without name of type two", func(t *testing.T) {
		var content = "() { echo \"abc\" }"
		var result = ingester.process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one function given one liner with correct function definition of type two", func(t *testing.T) {
		var content = "test() { echo \"hello world\"; }"
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with a function and comments given one liner with correct function definition of type two", func(t *testing.T) {
		var content = `
		# this is my function
		test() { echo "hello world"; }
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two functions and comments given content with multiple one line functions of type twi", func(t *testing.T) {
		var content = `
		# this is my function
		test1() { echo "hello"; }

		# this is my second function
		test2() { echo "world"; }
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "test1() { echo \"hello\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
			{
				Name:       "test2",
				Content:    "test2() { echo \"world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with a function and comments given a multi line function of type two", func(t *testing.T) {
		var content = `
		# this is my function
		test() { 
			echo "hello world"; 
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() {\n\t\t\techo \"hello world\"; \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice created from nested functions if type two", func(t *testing.T) {
		var content = `
		# this is my function
		test() { 
			world() { echo "world"'; }
			echo "hello";
			world 
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() {\n\t\t\tworld() { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice created with multiple nested functions of type two", func(t *testing.T) {
		var content = `
		# this is my function
		test() { 
			eval "${command}"
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() {\n\t\t\teval \"${command}\"\n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice created from nested functions of type two", func(t *testing.T) {
		var content = `
		# this is my function
		test1() { 
			world() { echo "world"'; }
			echo "hello";
			world 
		}

		# this is my second function
		test2() {
			world() { echo "hello world"; }
			world 
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "test1() {\n\t\t\tworld() { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
			{
				Name:       "test2",
				Content:    "test2() {\n\t\t\tworld() { echo \"hello world\"; }\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should parse a file with variois content in it", func(t *testing.T) {
		var content = `
		# this is my function
		test1() { 
			function world() { echo "world"'; }
			echo "hello";
			world 
		}

		# this is my first alias
		alias test='ls -all'

		# this is my second alias
		alias second = echo

		function test2 {
			echo "${command}"
		}
		`
		var result = ingester.process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "test1() {\n\t\t\tfunction world() { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my first alias"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
			{
				Name:       "second",
				Content:    "echo",
				Path:       "test.sh",
				Comments:   []string{"# this is my second alias"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
			},
			{
				Name:       "test2",
				Content:    "function test2 {\n\t\t\techo \"${command}\"\n\t\t}",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
			},
		}
		assert.Equal(t, expected, result)
	})
}

func TestConfigIngester_isPotentialAlias(t *testing.T) {
	t.Run("it should return false if alias keyword is not present", func(t *testing.T) {
		var content = ""
		var result = ingester.isPotentialAlias(content)
		assert.False(t, result)
	})

	t.Run("it should return false if alias is present but behind a comment", func(t *testing.T) {
		var content = "# like this: alias my-alias='ll -all'"
		var result = ingester.isPotentialAlias(content)
		assert.False(t, result)
	})

	t.Run("it should return false if alias is incorrect", func(t *testing.T) {
		var content = "alias='ll -all'"
		var result = ingester.isPotentialAlias(content)
		assert.False(t, result)
	})

	t.Run("it should return true if alias is present at the start of line", func(t *testing.T) {
		var content = "alias my-alias='ll -all'"
		var result = ingester.isPotentialAlias(content)
		assert.True(t, result)
	})

	t.Run("it should return true if alias is present after a tab", func(t *testing.T) {
		var content = "	alias my-alias='ll -all'"
		var result = ingester.isPotentialAlias(content)
		assert.True(t, result)
	})
}

func TestConfigIngester_isPotentialExport(t *testing.T) {
	t.Run("it should return false if export keyword is not present", func(t *testing.T) {
		var content = ""
		var result = ingester.isPotentialExport(content)
		assert.False(t, result)
	})

	t.Run("it should return false if export is present but behind a comment", func(t *testing.T) {
		var content = "# like this: export VAR=123"
		var result = ingester.isPotentialExport(content)
		assert.False(t, result)
	})

	t.Run("it should return false if export is incorrect", func(t *testing.T) {
		var content = "export=123"
		var result = ingester.isPotentialExport(content)
		assert.False(t, result)
	})

	t.Run("it should return true if export is present at the start of line", func(t *testing.T) {
		var content = "export VAR=123"
		var result = ingester.isPotentialExport(content)
		assert.True(t, result)
	})

	t.Run("it should return true if export is present after a tab", func(t *testing.T) {
		var content = "	export VAR=123"
		var result = ingester.isPotentialExport(content)
		assert.True(t, result)
	})
}

func TestConfigIngester_isPotentialFunctionStyleOne(t *testing.T) {
	t.Run("it should return false if function keyword is not present", func(t *testing.T) {
		var content = ""
		var result = ingester.isPotentialFunctionStyleOne(content)
		assert.False(t, result)
	})

	t.Run("it should return false if function is present but behind a comment", func(t *testing.T) {
		var content = "# like this: function test { echo; }"
		var result = ingester.isPotentialFunctionStyleOne(content)
		assert.False(t, result)
	})

	t.Run("it should return true even if function is incorrect", func(t *testing.T) {
		var content = "function { echo; }"
		var result = ingester.isPotentialFunctionStyleOne(content)
		assert.True(t, result)
	})

	t.Run("it should return true if function is present at the start of line", func(t *testing.T) {
		var content = "function test { echo; }"
		var result = ingester.isPotentialFunctionStyleOne(content)
		assert.True(t, result)
	})

	t.Run("it should return true if function is present after a tab", func(t *testing.T) {
		var content = "		function test { echo; }"
		var result = ingester.isPotentialFunctionStyleOne(content)
		assert.True(t, result)
	})
}

func TestConfigIngester_isPotentialFunctionStyleTwo(t *testing.T) {
	t.Run("it should return false if function keyword is not present", func(t *testing.T) {
		var content = ""
		var result = ingester.isPotentialFunctionStyleTwo(content)
		assert.False(t, result)
	})

	t.Run("it should return false if function is present but behind a comment", func(t *testing.T) {
		var content = "# like this: test() { echo; }"
		var result = ingester.isPotentialFunctionStyleTwo(content)
		assert.False(t, result)
	})

	t.Run("it should return true even if function is incorrect", func(t *testing.T) {
		var content = "test()() { echo; }"
		var result = ingester.isPotentialFunctionStyleTwo(content)
		assert.True(t, result)
	})

	t.Run("it should return true if function is present at the start of line", func(t *testing.T) {
		var content = "test() { echo; }"
		var result = ingester.isPotentialFunctionStyleTwo(content)
		assert.True(t, result)
	})

	t.Run("it should return true if function is present after a tab", func(t *testing.T) {
		var content = "		test() { echo; }"
		var result = ingester.isPotentialFunctionStyleTwo(content)
		assert.True(t, result)
	})
}
