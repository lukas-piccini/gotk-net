package ui

import "github.com/gotk3/gotk3/gtk"

func AddClass(widget gtk.IWidget, classname string) {
	ctx, _ := widget.ToWidget().GetStyleContext()
	ctx.AddClass(classname)
}
