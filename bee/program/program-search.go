package program

import (
	"bee/bbee/data"
	"bee/bbee/utils"
	"fmt"
	"strings"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

type SearchProgram struct {
	controller  SearchController
	cache       SearchCache
	showPreview bool
}

func (p SearchProgram) Run() {
	// setup the main app and run it
	app := cview.NewApplication()

	// setup the list
	list := cview.NewList()
	list.ShowSecondaryText(false)
	list.SetBorder(true)
	list.SetBorderColor(tcell.ColorBlack)
	list.SetTitleAlign(0)

	// details box
	details := cview.NewTextView()
	details.SetBorder(true)
	details.SetDynamicColors(true)
	details.SetWordWrap(true)
	details.SetBorderColor(tcell.ColorDimGray)

	// setup the main search field
	searchField := cview.NewInputField()
	searchField.SetLabel("> ")
	searchField.SetBackgroundColor(tcell.ColorBlack)
	searchField.SetFieldBackgroundColor(tcell.ColorBlack)
	searchField.SetFieldBackgroundColorFocused(tcell.ColorBlack)
	searchField.SetFieldTextColor(tcell.ColorDarkBlue)
	searchField.SetPlaceholder("Search for aliases, functions, scripts, etc. Press ESC to clear.")
	searchField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyUp {
			p.controller.moveUp()
		}
		if key == tcell.KeyDown {
			p.controller.moveDown()
		}
		if key == tcell.KeyEnter {
			p.stop()
			app.Stop()
		}
		if key == tcell.KeyEscape {
			searchField.SetText("")
		}
		p.redrawDetails(details)
		p.redrawList(list)
	})
	searchField.SetAutocompleteFunc(func(currentText string) (entries []*cview.ListItem) {
		p.controller.search(currentText)
		p.redrawDetails(details)
		p.redrawList(list)
		return
	})

	// setup the main layout
	layout := cview.NewFlex()
	layout.SetDirection(cview.FlexRow)

	display := cview.NewFlex()
	display.AddItem(list, 0, 1, true)
	if p.showPreview {
		display.AddItem(details, 0, 1, false)
	}
	layout.AddItem(display, 0, 1, false)
	layout.AddItem(searchField, 1, 0, true)

	// run the app with the fleg layout as root
	app.SetRoot(layout, true)
	app.EnableMouse(false)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func (p SearchProgram) redrawDetails(textView *cview.TextView) {
	if !p.showPreview {
		return
	}

	item := p.controller.getCurrentItem()
	category := p.controller.getNearestCategoryResult()
	var content string
	if item.noHighlight {
		content = p.cache.getPreviewForSearchResult(item)
	} else {
		content = p.cache.getPreviewForSearchResult(category)
	}

	var allText string
	if item.noHighlight {
		allText = p.colorize(content)
	} else {
		lines := strings.Split(content, "\n")
		firstLines := lines[:item.startLine]
		selectedLines := lines[item.startLine : item.endLine+1]
		endLines := lines[item.endLine+1:]

		firstContent := p.dehighlight(strings.Join(firstLines, "\n"))
		selectedContent := p.highlight(strings.Join(selectedLines, "\n"))
		endContent := p.dehighlight(strings.Join(endLines, "\n"))

		if len(firstLines) == 0 {
			allText = selectedContent + endContent
		} else {
			allText = firstContent + selectedContent + endContent
		}
	}

	var _, _, _, height = textView.GetRect()

	textView.Clear()
	// smooth scrolling
	if item.startLine > height/2 {
		textView.ScrollTo(item.startLine-height/2, 0)
	} else {
		textView.ScrollToBeginning()
	}

	// actually write things
	w := cview.ANSIWriter(textView)
	fmt.Fprintln(w, allText)

	// set the title
	textView.SetTitle(item.previewTitle)
}

func (p SearchProgram) redrawList(list *cview.List) {
	var data []SearchResult = p.controller.results
	var currentIndex = p.controller.currentIndex

	list.Clear()
	for _, s := range data {
		item := cview.NewListItem(s.mainText)
		list.AddItem(item)
	}
	if len(data) > 0 {
		list.SetCurrentItem(currentIndex)
	}

	list.SetSelectedBackgroundColor(tcell.ColorDarkBlue)
	p.redrawListTitle(list)
}

func (p SearchProgram) redrawListTitle(list *cview.List) {
	list.SetTitle(fmt.Sprintf("Showing %d/%d Results", p.controller.getNumberOfSearchResults(), p.controller.totalLen))
}

func (p SearchProgram) stop() {
	data.WriteLastCommand(p.controller.getCurrentItem().command)
}

func (p SearchProgram) highlight(content string) string {
	command := "echo \"" + content + "\" | bat -l Bash --color=always --style=plain --line-range=:500 --paging=never --theme=1337"
	result, _, err := utils.Shellout(command)
	if err != nil {
		return ""
	}
	result = strings.ReplaceAll("\033[48;2;35;35;45m"+result, "\033[0m", "\033[48;2;35;35;45m")
	return result + "\033[0m"
}

func (p SearchProgram) colorize(content string) string {
	command := "echo \"" + content + "\" | bat -l Bash --color=always --style=plain --line-range=:500 --paging=never --theme=1337"
	result, _, err := utils.Shellout(command)
	if err != nil {
		return ""
	}
	return result
}

func (p SearchProgram) dehighlight(content string) string {
	command := "echo \"" + content + "\" | bat -l Bash --color=always --style=plain --line-range=:500 --paging=never --theme=1337"
	result, _, err := utils.Shellout(command)
	if err != nil {
		return ""
	}
	result = strings.ReplaceAll("\033[2m"+result, "\033[0m", "\033[2m")
	return result + "\033[0m"
}
