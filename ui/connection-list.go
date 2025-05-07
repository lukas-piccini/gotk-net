package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type LoadFunc func(*ConnectionList)

type ConnectionList struct {
	loading          bool
	loadingContainer *gtk.ListBoxRow
	collapse         CollapsableList
	load             LoadFunc
	Component        *gtk.Box
}

func connectionListNew(title string, loadFunc LoadFunc) *ConnectionList {
	collapse := CollapsableListNew(title, true)

	cl := &ConnectionList{
		loading:   false,
		collapse:  collapse,
		Component: collapse.Component,
		load:      loadFunc,
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
