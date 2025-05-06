package ui

import "github.com/gotk3/gotk3/gtk"

func AddClass(widget gtk.IWidget, classname string) {
	ctx, _ := widget.ToWidget().GetStyleContext()
	ctx.AddClass(classname)
}

func ClearList(widget gtk.ListBox) {
	children := widget.GetChildren()
	for l := children; l != nil; l = l.Next() {
		w := l.Data().(*gtk.Widget)
		widget.Remove(w)
	}
}
