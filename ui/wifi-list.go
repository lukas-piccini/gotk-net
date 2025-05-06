package ui

import (
	"fmt"
	"gotk-net/net/commands"

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
					row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
					nameLabel, _ := gtk.LabelNew(item.Ssid)

					if item.Connected {
						AddClass(row, "connected")
					}

					row.Add(nameLabel)

					AddClass(row, "listbox-row")

					eventBox, _ := gtk.EventBoxNew()
					eventBox.Add(row)
					eventBox.Connect("button-press-event", func() {
						fmt.Println("name: ", item.Ssid)
					})

					listBoxRow, _ := gtk.ListBoxRowNew()
					listBoxRow.Add(eventBox)

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
