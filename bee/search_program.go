package main

import (
	"bee/bbee/data"
	"fmt"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

type SearchProgram struct {
	controller  SearchController
	cache       SearchCache
	showPreview bool
}

func (p SearchProgram) run() {
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
	// get the content
	var item = p.controller.getCurrentItem()
	var content = p.cache.getPreviewForSearchResult(item)

	// command to display with the "bat" utility, if present on the system
	command := "echo \"" + content + "\" | bat -l Bash --color=always --style=numbers --line-range=:500 --paging=never --theme=1337"
	out, _, err := Shellout(command)

	// if error, revert to setting the text
	if err != nil {
		textView.SetText(content)
	} else {
		textView.Clear()
		w := cview.ANSIWriter(textView)
		fmt.Fprintln(w, out)
	}
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
