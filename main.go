package main

import (
	"flag"
	"fmt"
	"gotk-net/net/commands"
	"gotk-net/net/ui"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type FilterRow struct {
	Text      string
	Container *gtk.ListBoxRow
}

var font string

func main() {
	flag.StringVar(&font, "font", "Monospace,12", "Set the font family and size of the application, separated by comma")
	flag.Parse()

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	connections := commands.Wifi()

	if err != nil {
		fmt.Println(err)
	}

	win.SetTitle("gotk-net")
	win.SetDefaultSize(600, 400)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetKeepAbove(true)
	win.SetTypeHint(gdk.WINDOW_TYPE_HINT_DIALOG)
	win.SetModal(true)
	win.SetDecorated(false)

	mainBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)
	ui.AddClass(mainBox, "main")

	inputBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	searchInput, _ := gtk.EntryNew()
	searchInput.SetPlaceholderText("Type here")
	ui.AddClass(searchInput, "search-input")

	inputBox.Add(searchInput)
	mainBox.Add(inputBox)

	separator, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	mainBox.Add(separator)

	vpnBox := ui.CollapsableListNew("Vpn", true)
	listBox := ui.CollapsableListNew(fmt.Sprintf("Wi-fi (%d)", len(connections)), true)
	ui.AddClass(listBox.List, "listbox")

	var listRows []FilterRow

	for _, item := range connections {
		row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		ssidLabel, _ := gtk.LabelNew(item.Ssid)

		if item.Connected {
			ui.AddClass(row, "connected")
		}

		row.Add(ssidLabel)

		ui.AddClass(row, "listbox-row")

		eventBox, _ := gtk.EventBoxNew()
		eventBox.Add(row)
		eventBox.Connect("button-press-event", func() {
			fmt.Println("SSID: ", item.Ssid)
		})

		listBoxRow, _ := gtk.ListBoxRowNew()
		listBoxRow.Add(eventBox)

		listRows = append(listRows, FilterRow{Text: item.Ssid, Container: listBoxRow})
		listBox.Add(listBoxRow)
	}

	searchInput.Connect("changed", func() {
		text, _ := searchInput.GetText()

		for _, row := range listRows {
			if strings.Contains(strings.ToLower(row.Text), strings.ToLower(text)) {
				row.Container.ShowAll()
			} else {
				row.Container.Hide()
			}
		}
	})

	scroll, _ := gtk.ScrolledWindowNew(nil, nil)
	scroll.SetVExpand(true)
	scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scroll.Add(listBox.Component)

	mainBox.Add(vpnBox.Component)
	mainBox.Add(scroll)

	win.Add(mainBox)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		key := gdk.EventKeyNewFromEvent(ev)

		if key.KeyVal() == gdk.KEY_Escape {
			gtk.MainQuit()
		}
	})

	css := fmt.Sprintf(`
	* {
		font-family: %s;
		font-size: %spx;
	}

	button, entry, row {
		all: unset;
	}

	.header-button {
		font-weight: bold;
		padding: 12px 4px;
		color: white;
	}

	separator {
		background: #4d6c88;
	}

	row:active, row:hover, row:selected {
		background: #1D2D44;
	}
	
	.main {
		padding: 12px;
		background: #0D1321;
	}

	.listbox {
		background: #0D1321;
	}

	.listbox-row.connected {
		color: green;
		font-weight: bold;
	}

	.listbox-row {
		padding: 8px 12px;
	}
		`, strings.Split(font, ",")[0], strings.Split(font, ",")[1])
	provider, _ := gtk.CssProviderNew()
	provider.LoadFromData(css)

	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	win.ShowAll()
	gtk.Main()
}
