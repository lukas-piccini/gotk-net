package ui

import (
	"fmt"
	"gotk-net/net/commands"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type WifiList struct {
	ConnectionList
}

func WifiListNew(title string) *WifiList {
	w := &WifiList{}
	connList := connectionListNew(title, func(x *ConnectionList) {
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

					//listRows = append(listRows, FilterRow{Text: item.Name, Container: listBoxRow})
					x.collapse.List.Add(listBoxRow)
					listBoxRow.ShowAll()
				}
			})
		}()
	})

	w.ConnectionList = *connList
	w.Load()

	return w
}
