package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type Password struct {
	Component *gtk.Box
	revelear  *gtk.Revealer
	loading   bool
	visible   bool
}

func PasswordNew() *Password {
	container, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	revealer, _ := gtk.RevealerNew()
	revealer.SetTransitionType(gtk.REVEALER_TRANSITION_TYPE_SLIDE_DOWN)
	revealer.SetRevealChild(false)

	revealerContainer, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)

	passwordLabel, _ := gtk.LabelNew("Password")
	passwordLabel.SetXAlign(0)

	inputBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 8)

	input, _ := gtk.EntryNew()
	input.SetVisibility(false)
	input.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)
	inputBox.PackStart(input, false, false, 0)

	viewPasswordButton, _ := gtk.ButtonNewWithLabel("\uf06e")
	inputBox.PackStart(viewPasswordButton, false, false, 0)

	revealerContainer.Add(passwordLabel)
	revealerContainer.Add(inputBox)
	revealer.Add(revealerContainer)
	container.Add(revealer)

	pass := &Password{revelear: revealer, visible: false, loading: false, Component: container}

	return pass
}

func (p *Password) ToggleDisplay() {
	if p.visible {
		p.revelear.SetRevealChild(false)
		p.visible = false
	} else {
		p.revelear.SetRevealChild(true)
		p.visible = true
	}
}
