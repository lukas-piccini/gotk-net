package ui

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Searchable interface {
	GetFilter() string
}

type App struct {
	Window *gtk.Window
	vpn    *VpnList
	wifi   *WifiList
	filter string
}

func AppNew() *App {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

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
	AddClass(mainBox, "main")

	inputBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	searchInput, _ := gtk.EntryNew()
	searchInput.SetPlaceholderText("Type here")
	filter := new(string)
	AddClass(searchInput, "search-input")

	inputBox.Add(searchInput)
	mainBox.Add(inputBox)

	separator, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	mainBox.Add(separator)

	scroll, _ := gtk.ScrolledWindowNew(nil, nil)
	scroll.SetVExpand(true)
	scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scrollBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	app := &App{}

	vpn := VpnListNew("Vpn", app)
	wifi := WifiListNew("Wifi", app)
	scrollBox.Add(vpn.Component)
	scrollBox.Add(wifi.Component)
	scroll.Add(scrollBox)

	info, _ := gtk.LabelNew("Reload connections (Ctrl + r)")
	infoBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	infoBox.SetVExpand(true)
	mainBox.Add(scroll)
	mainBox.Add(info)
	win.Add(mainBox)

	LoadTheme()

	app.Window = win
	app.vpn = vpn
	app.wifi = wifi
	app.filter = *filter

	app.setFilter(searchInput)
	return app
}

func (w *App) Run() {
	w.attachDefaultEvents()
	w.Window.ShowAll()
	gtk.Main()
}

func (w *App) attachDefaultEvents() {
	w.Window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	w.Window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		key := gdk.EventKeyNewFromEvent(ev)

		if key.KeyVal() == gdk.KEY_Escape {
			gtk.MainQuit()
		}
	})
}

func (w *App) setFilter(entry *gtk.Entry) {
	entry.Connect("changed", func() {
		text, _ := entry.GetText()
		w.filter = strings.ToLower(text)
		w.wifi.Filter(text)
		w.vpn.Filter(text)
	})
}

func LoadTheme() {
	var font string
	flag.StringVar(&font, "font", "Monospace,12", "Set the font family and size of the application, separated by comma")
	flag.Parse()

	css := fmt.Sprintf(`
	* {
		font-family: %s;
		font-size: %spx;
	}

	button, entry, row {
		all: unset;
	}

	.search-input {
		padding: 8px;
	}

	.header-button {
		font-weight: bold;
		padding: 12px 4px;
		color: white;
	}

	spinner {
		min-width: 16px;
		min-height: 16px;
	}

	separator {
		background: #4d6c88;
	}

	.header-button:active, header-button:focus,
	row:active, row:selected {
		background: #1D2D44;
	}

	.main {
		padding: 12px;
		background: #0D1321;
	}

	.listbox {
		background: #0D1321;
	}

	.listbox-row .connected-label {
		color: green;
	}

	.listbox-row {
		padding: 8px 16px;
	}

	.wifi-low {
		color: red;
	}

	.wifi-medium {
		color: orange;
	}

	.wifi-high {
		color: yellow;
	}

	.wifi-strong {
		color: green;
	}
`, strings.Split(font, ",")[0], strings.Split(font, ",")[1])

	provider, _ := gtk.CssProviderNew()
	provider.LoadFromData(css)

	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

func (a *App) GetFilter() string {
	return a.filter
}
