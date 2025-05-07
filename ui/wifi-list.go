package ui

import (
	"fmt"
	"gotk-net/net/commands"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type WifiList struct {
	ConnectionList
}

func WifiListNew(title string) *WifiList {
	w := &WifiList{}
	connList := connectionListNew(title, w.load, w.filter)
	w.ConnectionList = *connList
	w.Load()

	return w
}

func (w *WifiList) load(x *ConnectionList) {
	x.ToggleLoading()
	go func() {
		connections := commands.Wifi()

		glib.IdleAdd(func() {
			x.ToggleLoading()

			for _, item := range connections {
				item := item
				row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
				nameLabel, _ := gtk.LabelNew(item.Ssid)

				if item.Connected {
					AddClass(row, "connected")
				}

				row.Add(nameLabel)

				AddClass(row, "listbox-row")

				eventBox, _ := gtk.EventBoxNew()
				eventBox.Add(row)

				listBoxRow, _ := gtk.ListBoxRowNew()
				listBoxRow.Add(eventBox)

				listBoxRow.Connect("key-press-event", func(win *gtk.ListBoxRow, ev *gdk.Event) {
					key := gdk.EventKeyNewFromEvent(ev)

					if key.KeyVal() == gdk.KEY_Return {
						fmt.Println("name: ", item.Ssid)
					}
				})

				x.filterRow[item.Ssid] = listBoxRow
				x.collapse.List.Add(listBoxRow)
				listBoxRow.ShowAll()
			}
		})
	}()
}

func (w *WifiList) filter(x *ConnectionList, search string) {
	for key, item := range w.filterRow {
		if strings.Contains(strings.ToLower(key), strings.ToLower(search)) {
			item.Container.ShowAll()
		} else {
			item.Container.Hide()
		}
	}
}
