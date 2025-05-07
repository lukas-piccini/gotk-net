package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type LoadFunc func(*ConnectionList)
type FilterFunc func(*ConnectionList, string)

type ConnectionList struct {
	loading          bool
	loadingContainer *gtk.ListBoxRow
	collapse         CollapsableList
	filter           FilterFunc
	load             LoadFunc
	filterRow        map[string]*gtk.ListBoxRow
	Component        *gtk.Box
}

func connectionListNew(title string, loadFunc LoadFunc, filterFunc FilterFunc) *ConnectionList {
	collapse := CollapsableListNew(title, true)

	cl := &ConnectionList{
		loading:   false,
		collapse:  collapse,
		Component: collapse.Component,
		load:      loadFunc,
		filter:    filterFunc,
		filterRow: make(map[string]*gtk.ListBoxRow),
	}

	return cl
}

func (cl *ConnectionList) ToggleLoading() {
	if cl.loading {
		cl.loading = false

		if cl.loadingContainer != nil {
			cl.loadingContainer.Destroy()
			cl.loadingContainer = nil
		}
	} else {
		cl.loading = true

		ClearList(*cl.collapse.List)

		listBoxRow, _ := gtk.ListBoxRowNew()
		AddClass(listBoxRow, "listbox-row")
		box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 16)

		spinner, _ := gtk.SpinnerNew()
		spinner.Start()

		text, _ := gtk.LabelNew("Loading connections")

		box.Add(text)
		box.Add(spinner)
		listBoxRow.Add(box)

		cl.loadingContainer = listBoxRow

		cl.collapse.List.Add(listBoxRow)
	}
}

func (cl *ConnectionList) Load() {
	if cl.load != nil {
		cl.load(cl)
	}
}

func (cl *ConnectionList) Filter(search string) {
	if cl.filter != nil {
		cl.filter(cl, search)
	}
}
