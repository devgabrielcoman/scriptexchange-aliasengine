package ingester

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigIngester_Process_Alias(t *testing.T) {
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}
	t.Run("it should return an empty slice given an empty input", func(t *testing.T) {
		var content = ""
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given garbage input", func(t *testing.T) {
		var content = "\\\asa/////"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given incorrectly formatted alias", func(t *testing.T) {
		var content = "alias test'ls -all'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given incorrectly formatted alias", func(t *testing.T) {
		var content = "alias test 'ls -all'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should a slice with one alias even if format is slighlty incorrect", func(t *testing.T) {
		var content = "alias test = 'ls -all'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one alias in one line content", func(t *testing.T) {
		var content = "alias test='ls -all'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should not return a slice with alises if it actually starts with a comment", func(t *testing.T) {
		var content = `# this is my
		# write down like: alias test='ls -all'
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one alias and comments in multi line content", func(t *testing.T) {
		var content = `# this is my
		# comment
		alias test='ls -all'
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my", "# comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  0,
				EndLine:    2,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "test2",
				Content:    "run this",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  4,
				EndLine:    6,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "test2",
				Content:    "run this",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  4,
				EndLine:    6,
			},
		}
		assert.Equal(t, expected, result)
	})
}

func TestConfigIngester_Process_Export(t *testing.T) {
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}

	t.Run("it should return an empty slice given incorrectly formatted export", func(t *testing.T) {
		var content = "export abc"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given incorrectly formatted export", func(t *testing.T) {
		var content = "export $PATH 'ls -all'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should a slice with one export even if format is slighlty incorrect", func(t *testing.T) {
		var content = "export $PATH = 'my-path'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "$PATH",
				Content:    "my-path",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one export in one line content", func(t *testing.T) {
		var content = "export $PATH='my-path'"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "$PATH",
				Content:    "my-path",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should not return a slice with export if it actually starts with a comment", func(t *testing.T) {
		var content = `# this is my
		# write down like: export $PATH='my-path'
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one export and comments in multi line content", func(t *testing.T) {
		var content = `# this is my
		# comment
		export $PATH='my-path'
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "$PATH",
				Content:    "my-path",
				Path:       "test.sh",
				Comments:   []string{"# this is my", "# comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  0,
				EndLine:    2,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two exports in multi line content", func(t *testing.T) {
		var content = `
		# this is my comment
		export $PATH1='my-path-1'

		# this is the
		# second comment
		export $PATH2='my-path-2'
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "$PATH1",
				Content:    "my-path-1",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "$PATH2",
				Content:    "my-path-2",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  4,
				EndLine:    6,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two exports even if identical names", func(t *testing.T) {
		var content = `
		# this is my comment
		export $PATH='my-path-1'

		# this is the
		# second comment
		export $PATH='my-path-2'
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "$PATH",
				Content:    "my-path-1",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "$PATH",
				Content:    "my-path-2",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  4,
				EndLine:    6,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with two exports even if the format is slightly incorrect", func(t *testing.T) {
		var content = `
		# this is my comment
		export $PATH1 = 'my-path-1'

		# this is the
		# second comment
		export $PATH2 = my-path-2
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "$PATH1",
				Content:    "my-path-1",
				Path:       "test.sh",
				Comments:   []string{"# this is my comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "$PATH2",
				Content:    "my-path-2",
				Path:       "test.sh",
				Comments:   []string{"# this is the", "# second comment"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  4,
				EndLine:    6,
			},
		}
		assert.Equal(t, expected, result)
	})
}

func TestConfigIngester_Process_FunctionStyleOne(t *testing.T) {
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}

	t.Run("it should return an empty slice given function without keywords", func(t *testing.T) {
		var content = "{ echo \"abc\"; }"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without body of type one", func(t *testing.T) {
		var content = "function abc"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without name of type one", func(t *testing.T) {
		var content = "function { echo \"abc\" }"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one function given one liner with correct function of type one", func(t *testing.T) {
		var content = "function test { echo \"hello world\"; }"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with a function and comments given one liner with correct function of type one", func(t *testing.T) {
		var content = `
		# this is my function
		function test { echo "hello world"; }
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    2,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "function test1 { echo \"hello\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "test2",
				Content:    "function test2 { echo \"world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  4,
				EndLine:    5,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test {\n\t\t\techo \"hello world\"; \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    4,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test {\n\t\t\tfunction world { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    6,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "function test {\n\t\t\teval \"${command}\"\n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    4,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "function test1 {\n\t\t\tfunction world { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    6,
			},
			{
				Name:       "test2",
				Content:    "function test2 {\n\t\t\tfunction world { echo \"hello world\"; }\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  8,
				EndLine:    12,
			},
		}
		assert.Equal(t, expected, result)
	})
}

func TestConfigIngester_Process_FunctionStyleTwo(t *testing.T) {
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}

	t.Run("it should return an empty slice given function without body of type two", func(t *testing.T) {
		var content = "abc()"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return an empty slice given function without name of type two", func(t *testing.T) {
		var content = "() { echo \"abc\" }"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with one function given one liner with correct function definition of type two", func(t *testing.T) {
		var content = "test() { echo \"hello world\"; }"
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  0,
				EndLine:    0,
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("it should return a slice with a function and comments given one liner with correct function definition of type two", func(t *testing.T) {
		var content = `
		# this is my function
		test() { echo "hello world"; }
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() { echo \"hello world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    2,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "test1() { echo \"hello\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    2,
			},
			{
				Name:       "test2",
				Content:    "test2() { echo \"world\"; }",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  4,
				EndLine:    5,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() {\n\t\t\techo \"hello world\"; \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    4,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() {\n\t\t\tworld() { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    6,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test",
				Content:    "test() {\n\t\t\teval \"${command}\"\n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    4,
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
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "test1() {\n\t\t\tworld() { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    6,
			},
			{
				Name:       "test2",
				Content:    "test2() {\n\t\t\tworld() { echo \"hello world\"; }\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my second function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  8,
				EndLine:    12,
			},
		}
		assert.Equal(t, expected, result)
	})

}

func TestConfigIngester_Process(t *testing.T) {
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}
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

		# this is an export
		export $PATH="my-new-path"
		`
		var result = ingester.Process(content)
		var expected = []models.IndexItem{
			{
				Name:       "test1",
				Content:    "test1() {\n\t\t\tfunction world() { echo \"world\"'; }\n\t\t\techo \"hello\";\n\t\t\tworld \n\t\t}",
				Path:       "test.sh",
				Comments:   []string{"# this is my function"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  1,
				EndLine:    6,
			},
			{
				Name:       "test",
				Content:    "ls -all",
				Path:       "test.sh",
				Comments:   []string{"# this is my first alias"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  8,
				EndLine:    9,
			},
			{
				Name:       "second",
				Content:    "echo",
				Path:       "test.sh",
				Comments:   []string{"# this is my second alias"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Alias),
				StartLine:  11,
				EndLine:    12,
			},
			{
				Name:       "test2",
				Content:    "function test2 {\n\t\t\techo \"${command}\"\n\t\t}",
				Path:       "test.sh",
				Comments:   []string{},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Function),
				StartLine:  14,
				EndLine:    16,
			},
			{
				Name:       "$PATH",
				Content:    "my-new-path",
				Path:       "test.sh",
				Comments:   []string{"# this is an export"},
				PathOnDisk: "test.sh",
				Type:       models.ScriptType(models.Export),
				StartLine:  18,
				EndLine:    19,
			},
		}
		assert.Equal(t, expected, result)
	})
}

func TestConfigIngester_isPotentialAlias(t *testing.T) {
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}
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
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}
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
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}
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
	var ingester = ConfigIngester{FilePath: "test.sh", CurrentTime: 0}
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
