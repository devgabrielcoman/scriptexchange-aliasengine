package style

type ColorCode string

const (
	ColorDimGray         ColorCode = "[#696969]"
	AliasKeywordColor    ColorCode = "[#659acc]"
	AliasNameColor       ColorCode = "[#8cdbff]"
	FunctionKeywordColor ColorCode = "[#fbdfb5]"
	FunctionNameColor    ColorCode = "[#8bdaff]"
	ExportKeywordColor   ColorCode = "[#ff5e5f]"
	ExportNameColor      ColorCode = "[#d8ea9f]"
	ScriptNameColor      ColorCode = "[#e9fdac]"
	ColorWhite           ColorCode = "[#ededed]"
)

func Color(text string, colorCode ColorCode) string {
	return string(colorCode) + text
}
