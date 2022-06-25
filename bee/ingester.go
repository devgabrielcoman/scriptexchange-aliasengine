package main

import (
	"bee/bbee/models"
	"bee/bbee/utils"
	"sort"
	"strings"

	"github.com/samber/lo"
)

// main constants for the file ingeter
const (
	ALIAS_PREFIX         string = "alias "
	ALIAS_SEPARATOR      string = "="
	EXPORT_PREFIX        string = "export "
	EXPORT_SEPARATOR     string = "="
	START_CHAR_QUOTE     string = "'"
	START_CHAR_DBL_QUOTE string = "\""
	COMMENT_PREFIX       string = "#"
	WHITESPACE           string = " "
	FUNCTION_KEYWORD_ONE string = "function"
	FUNCTION_KEYWORD_TWO string = "()"
	OPEN_BRACKET         string = "{"
	CLOSE_BRACKET        string = "}"
	OPEN_PARA            string = "("
	CLOSE_PARA           string = ")"
	SEPARATOR            string = ""
	NEWLINE              string = "\n"
	TAB                  string = "\t"
	ZSH_HISTORY_SEP      string = ";"
	ZSH_HISTORY_SUFFIX   string = "\\"
)

// The ConfigIngester contains methods to ingest
// aliases, functions, etc
// contained in files such as .bashrc, .profile, .zshrc, etc
type ConfigIngester struct {
	filePath    string
	currentTime int64
}

func (c ConfigIngester) process(content string) []models.IndexItem {
	var result = []models.IndexItem{}

	// separate the contents by line
	var lines []string = strings.Split(content, NEWLINE)

	var i = -1
	for i < len(lines)-1 {
		i += 1
		var line = lines[i]
		var trimmedLine = strings.Trim(line, WHITESPACE)

		// found a pontential alias line
		if c.isPotentialAlias(trimmedLine) {
			var item, progress = c.processAlias(line, i, lines)
			if item != nil {
				result = append(result, *item)
				i += progress
			}
		}

		// found a potential export
		if c.isPotentialExport(trimmedLine) {
			var item, process = c.processExport(line, i, lines)
			if item != nil {
				result = append(result, *item)
				i += process
			}
		}

		// found potential function in first style
		if c.isPotentialFunctionStyleOne(trimmedLine) {
			var item, progress = c.processFunctionInStyleOne(line, i, lines)
			if item != nil {
				result = append(result, *item)
				i += progress
			}
		}

		// found potential function in second style
		if c.isPotentialFunctionStyleTwo(trimmedLine) {
			var item, progress = c.processFunctionInStyleTwo(line, i, lines)
			if item != nil {
				result = append(result, *item)
				i += progress
			}
		}
	}

	return result
}

func (c ConfigIngester) isPotentialAlias(line string) bool {
	return strings.Contains(line, ALIAS_PREFIX) && !strings.Contains(line, COMMENT_PREFIX)
}

func (c ConfigIngester) isPotentialExport(line string) bool {
	return strings.Contains(line, EXPORT_PREFIX) && !strings.Contains(line, COMMENT_PREFIX)
}

func (c ConfigIngester) isPotentialFunctionStyleOne(line string) bool {
	return strings.Contains(line, FUNCTION_KEYWORD_ONE) && !strings.Contains(line, COMMENT_PREFIX)
}

func (c ConfigIngester) isPotentialFunctionStyleTwo(line string) bool {
	return strings.Contains(line, FUNCTION_KEYWORD_TWO) && !strings.Contains(line, COMMENT_PREFIX)
}

func (c ConfigIngester) processAlias(line string, startIndex int, allLines []string) (*models.IndexItem, int) {
	var trimmedLine = c.trimLine(line)
	var aliasWithoutPrefix = strings.TrimPrefix(trimmedLine, ALIAS_PREFIX)
	var aliasComponents = strings.Split(aliasWithoutPrefix, ALIAS_SEPARATOR)

	// somehow not a valid alias
	if len(aliasComponents) < 2 {
		return nil, 0
	}

	// get alias name
	var aliasName = c.trimLine(aliasComponents[0])

	// get & parse the alias command
	var aliasCommand = c.trimLine(strings.Join(aliasComponents[1:], ""))

	if strings.HasPrefix(aliasCommand, START_CHAR_QUOTE) {
		aliasCommand = strings.Trim(aliasCommand, START_CHAR_QUOTE)
	}

	if strings.HasPrefix(aliasCommand, START_CHAR_DBL_QUOTE) {
		aliasCommand = strings.Trim(aliasCommand, START_CHAR_DBL_QUOTE)
	}

	// get comments
	var comments = c.getComments(startIndex, allLines)

	// create item
	var indexItem = models.IndexItem{
		Name:       aliasName,
		Content:    aliasCommand,
		Path:       c.getFileName(),
		Comments:   comments,
		PathOnDisk: c.filePath,
		Type:       models.ScriptType(models.Alias),
		Date:       c.currentTime,
	}

	return &indexItem, 0
}

func (c ConfigIngester) processExport(line string, startIndex int, allLines []string) (*models.IndexItem, int) {
	var trimmedLine = c.trimLine(line)
	var exportWithoutPrefix = strings.TrimPrefix(trimmedLine, EXPORT_PREFIX)
	var exportComponents = strings.Split(exportWithoutPrefix, EXPORT_SEPARATOR)

	// somehow not a valid alias
	if len(exportComponents) != 2 {
		return nil, 0
	}

	// get alias name
	var exportName = c.trimLine(exportComponents[0])

	// get & parse the alias command
	var exportCommand = c.trimLine(strings.Join(exportComponents[1:], ""))

	if strings.HasPrefix(exportCommand, START_CHAR_QUOTE) {
		exportCommand = strings.Trim(exportCommand, START_CHAR_QUOTE)
	}

	if strings.HasPrefix(exportCommand, START_CHAR_DBL_QUOTE) {
		exportCommand = strings.Trim(exportCommand, START_CHAR_DBL_QUOTE)
	}

	// get comments
	var comments = c.getComments(startIndex, allLines)

	// create item
	var indexItem = models.IndexItem{
		Name:       exportName,
		Content:    exportCommand,
		Path:       c.getFileName(),
		Comments:   comments,
		PathOnDisk: c.filePath,
		Type:       models.ScriptType(models.Export),
		Date:       c.currentTime,
	}

	return &indexItem, 0
}

func (c ConfigIngester) processFunctionInStyleOne(line string, startIndex int, allLines []string) (*models.IndexItem, int) {
	// prepare the line by replacing the keyboard and any start and end whitespaces
	var trimmedLine = c.trimLine(line)
	var preparedLine = strings.Trim(strings.ReplaceAll(trimmedLine, FUNCTION_KEYWORD_ONE, SEPARATOR), WHITESPACE)

	var functionName = ""
	var hasSeenFirstBracket = false
	var openBrackets = 0

	var characterArray = strings.Split(preparedLine, SEPARATOR)
	var nextIndex = startIndex
	var allContent = ""
	var i = -1
	var totalLength = len(characterArray)

	for i < totalLength-1 {
		i += 1
		var nextChar = characterArray[i]
		allContent += nextChar

		if nextChar == OPEN_BRACKET {
			// get the function name correctly
			if !hasSeenFirstBracket {
				var prevLimit = utils.Max(i-1, 0)
				var contentSoFar = strings.Join(characterArray[0:prevLimit], SEPARATOR)
				var potentialFunctionName = strings.Split(contentSoFar, WHITESPACE)
				if len(potentialFunctionName) == 1 {
					functionName = potentialFunctionName[0]
				} else {
					return nil, 0
				}
			}

			hasSeenFirstBracket = true
			openBrackets += 1
		}

		if nextChar == CLOSE_BRACKET {
			openBrackets -= 1
		}

		// if we're at the end and we still haven't closed the function
		if i == totalLength-1 && openBrackets != 0 {
			nextIndex += 1

			if nextIndex < len(allLines)-1 {
				var nextLine = allLines[nextIndex]
				var preparedNextLine = strings.Trim(nextLine, NEWLINE)
				var nextLineArray = strings.Split(preparedNextLine, SEPARATOR)
				characterArray = append(characterArray, nextLineArray...)
				allContent += NEWLINE
				totalLength += len(preparedNextLine)
			}
		}
	}

	if functionName == "" || openBrackets != 0 || !hasSeenFirstBracket {
		return nil, 0
	}

	var progress = nextIndex - startIndex
	var scriptType models.ScriptType = models.ScriptType(models.Function)
	var name = functionName
	var content = FUNCTION_KEYWORD_ONE + " " + allContent
	var path = c.getFileName()
	var comments = c.getComments(startIndex, allLines)
	var pathOnDisk = c.filePath
	var indexItem = models.IndexItem{
		Name:       name,
		Content:    content,
		Path:       path,
		Comments:   comments,
		PathOnDisk: pathOnDisk,
		Type:       scriptType,
		Date:       c.currentTime,
	}

	return &indexItem, progress
}

func (c ConfigIngester) processFunctionInStyleTwo(line string, startIndex int, allLines []string) (*models.IndexItem, int) {
	// prepare the line
	var trimmedLine = c.trimLine(line)
	var preparedLine = strings.Trim(trimmedLine, WHITESPACE)

	var functioName = ""
	var hasSeenFirstPara = false
	var paranthesesNumber = 0
	var openBrackets = 0
	var hasSeenFirstBracket = false

	var characterArray = strings.Split(preparedLine, SEPARATOR)
	var nextIndex = startIndex
	var allContent = ""
	var i = -1
	var totalLength = len(characterArray)

	for i < totalLength-1 {
		i += 1
		var nextChar = characterArray[i]
		allContent += nextChar

		if nextChar == OPEN_PARA {
			hasSeenFirstPara = true
			paranthesesNumber += 1
		}

		if nextChar == CLOSE_PARA {
			paranthesesNumber += 1
		}

		if !hasSeenFirstPara {
			functioName += nextChar
		}

		if nextChar == OPEN_BRACKET {
			hasSeenFirstBracket = true
			openBrackets += 1
		}

		if nextChar == CLOSE_BRACKET {
			openBrackets -= 1
		}

		// if we're at the end AND we still haven't closed the function
		if i == totalLength-1 && openBrackets != 0 {
			nextIndex += 1

			if nextIndex < len(allLines)-1 {
				var nextLine = allLines[nextIndex]
				var preparedNextLine = strings.Trim(nextLine, NEWLINE)
				var nextLineArray = strings.Split(preparedNextLine, SEPARATOR)
				characterArray = append(characterArray, nextLineArray...)
				allContent += NEWLINE
				totalLength += len(preparedNextLine)
			}
		}
	}

	var name = strings.Trim(functioName, WHITESPACE)

	if name == "" || paranthesesNumber < 2 || openBrackets != 0 || !hasSeenFirstBracket || len(strings.Split(name, WHITESPACE)) != 1 {
		return nil, 0
	}

	var progress = nextIndex - startIndex
	var scriptType models.ScriptType = models.ScriptType(models.Function)
	var content = allContent
	var path = c.getFileName()
	var comments = c.getComments(startIndex, allLines)
	var pathOnDisk = c.filePath
	var indexItem = models.IndexItem{
		Name:       name,
		Content:    content,
		Path:       path,
		Comments:   comments,
		PathOnDisk: pathOnDisk,
		Type:       scriptType,
		Date:       c.currentTime,
	}

	return &indexItem, progress
}

func (c ConfigIngester) getComments(startIndex int, lines []string) []string {
	var index = startIndex - 1
	var comments = []string{}

	for index >= 0 && index < len(lines) {
		var previousLine = c.trimLine(lines[index])
		if strings.HasPrefix(strings.Trim(previousLine, WHITESPACE), COMMENT_PREFIX) {
			var processed = strings.Trim(previousLine, WHITESPACE)
			comments = append(comments, processed)
			index -= 1
		} else {
			index = -1
		}
	}

	return lo.Reverse(comments)
}

func (c ConfigIngester) trimLine(line string) string {
	return strings.TrimSpace(strings.ReplaceAll(line, TAB, ""))
}

func (c ConfigIngester) getFileName() string {
	return utils.FileName(c.filePath)
}

// The ScriptIngester just ingests a new full script
type ScriptIngester struct {
	alias       string
	path        string
	currentTime int64
}

func (s ScriptIngester) process(content string) []models.IndexItem {
	return []models.IndexItem{
		{
			Name:       s.alias,
			Content:    content,
			Path:       ".scripts",
			Comments:   []string{},
			PathOnDisk: s.path,
			Type:       models.ScriptType(models.Script),
			Date:       s.currentTime,
		},
	}
}

// The BashHistoryIngester ingests a .bash_history type file
type BashHistoryIngester struct {
	path string
}

func (h BashHistoryIngester) process(content string) []models.IndexItem {
	// separate the contents by line
	var lines []string = strings.Split(content, NEWLINE)
	var result = []models.IndexItem{}

	for _, line := range lines {
		item := models.IndexItem{
			Name:       line,
			Content:    line,
			Comments:   []string{},
			Path:       h.path,
			PathOnDisk: h.path,
			Type:       models.ScriptType(models.History),
			Date:       0, // special case here, for bash we don't really have date info
		}
		result = append(result, item)
	}

	return uniqueItems(result)
}

// The ZSHHistoryIngester ingests a .zsh_history type file
type ZSHHistoryIngester struct {
	path string
}

func (z ZSHHistoryIngester) process(content string) []models.IndexItem {
	// separate the contents by line
	lines := strings.Split(content, NEWLINE)
	var result = []models.IndexItem{}

	for _, line := range lines {
		lineItems := strings.Split(line, ZSH_HISTORY_SEP)

		// not a valid line in the format date;command
		if len(lineItems) < 2 {
			continue
		}

		date := utils.LenientAtoi64(z.parseZSHDateItem(lineItems[0]))
		command := strings.Join(lineItems[1:], SEPARATOR)
		item := models.IndexItem{
			Name:       command,
			Content:    command,
			Comments:   []string{},
			Path:       z.path,
			PathOnDisk: z.path,
			Type:       models.ScriptType(models.History),
			Date:       date,
		}
		result = append(result, item)
	}

	result = uniqueItemsByDate(result)

	sort.Slice(result, func(i, j int) bool {
		return result[j].Date < result[i].Date
	})

	return result
}

// parses a ZSH date item in the format: ": 1642437214:0"
func (z ZSHHistoryIngester) parseZSHDateItem(rawDate string) string {
	date := strings.TrimPrefix(strings.TrimSuffix(rawDate, ":0"), ": ")
	return date
}
