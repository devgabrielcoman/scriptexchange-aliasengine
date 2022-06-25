package ingester

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

type Ingester struct{}
