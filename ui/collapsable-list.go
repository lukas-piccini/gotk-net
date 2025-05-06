package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type CollapsableList struct {
	Component *gtk.Box
	Header    *gtk.Button
	List      *gtk.ListBox
	revealer  *gtk.Revealer
}

func CollapsableListNew(title string, open bool) CollapsableList {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	header, _ := gtk.ButtonNewWithLabel(title)
	revealer, _ := gtk.RevealerNew()
	list, _ := gtk.ListBoxNew()
	AddClass(list, "listbox")

	headerChild, _ := header.GetChild()
	headerLabel, _ := headerChild.(*gtk.Label)
	headerLabel.SetXAlign(0)

	styleContext, _ := header.GetStyleContext()
	styleContext.AddClass("header-button")

	scroll, _ := gtk.ScrolledWindowNew(nil, nil)
	scroll.SetPropagateNaturalHeight(true)
	scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scroll.Add(list)

	revealer.Add(scroll)
	revealer.SetTransitionType(gtk.REVEALER_TRANSITION_TYPE_SLIDE_DOWN)
	revealer.SetTransitionDuration(250)

	box.Add(header)
	box.Add(revealer)

	finalList := CollapsableList{
		Component: box,
		Header:    header,
		List:      list,
		revealer:  revealer,
	}

	header.Connect("clicked", func() {
		finalList.Toggle()
	})

	if open {
		revealer.SetRevealChild(true)
	}

	return finalList
}

func (collapse *CollapsableList) Add(children gtk.IWidget) {
	collapse.List.Add(children)
}

func (collapse *CollapsableList) Toggle() {
	isOpen := collapse.revealer.GetRevealChild()

	collapse.List.SetCanFocus(!isOpen)
	collapse.revealer.SetRevealChild(!isOpen)
}
