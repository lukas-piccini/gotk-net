package ui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Password struct {
	WithPassword
	Component         *gtk.Box
	IsShowingPassword bool
	revelear          *gtk.Revealer
	loading           bool
	visible           bool
	entry             *gtk.Entry
	visibilityButton  *gtk.Button
}

func PasswordNew(app WithPassword) *Password {
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
	AddClass(input, "password-input")
	inputBox.PackStart(input, false, false, 0)

	viewPasswordButton, _ := gtk.ButtonNewWithLabel("\uf06e")
	AddClass(viewPasswordButton, "reveal-password-button")
	inputBox.PackStart(viewPasswordButton, false, false, 0)

	revealerContainer.Add(passwordLabel)
	revealerContainer.Add(inputBox)
	revealer.Add(revealerContainer)
	container.Add(revealer)

	pass := &Password{revelear: revealer, visible: false, loading: false, Component: container, entry: input, IsShowingPassword: false, visibilityButton: viewPasswordButton, WithPassword: app}

	viewPasswordButton.Connect("clicked", func(_ *gtk.Button) {
		pass.setPasswordShow(!pass.IsShowingPassword)
	})

	container.Connect("key-press-event", func(win *gtk.Box, ev *gdk.Event) {
		key := gdk.EventKeyNewFromEvent(ev)

		if key.KeyVal() == gdk.KEY_Escape {
			pass.TogglePassword()
		}
	})

	return pass
}

func (p *Password) setPasswordShow(show bool) {
	p.IsShowingPassword = show
	p.entry.SetVisibility(show)

	if show {
		p.visibilityButton.SetLabel("\uf070")
	} else {
		p.visibilityButton.SetLabel("\uf06e")
	}
}

func (p *Password) SetVisible(visible bool) {
	p.revelear.SetRevealChild(visible)
	p.visible = visible

	if visible {
		p.entry.GrabFocus()
	} else {
		p.entry.SetText("")
		p.setPasswordShow(false)
	}
}
