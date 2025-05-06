package ui

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"gotk-net/net/commands"
)

type VpnList struct {
	ConnectionList
}

func VpnListNew(title string) *ConnectionList {
	vpnList := connectionListNew(title, func(x *ConnectionList) {
		x.ToggleLoading()

		connections := commands.Vpn()

		x.ToggleLoading()

		for _, item := range connections {
			row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
			nameLabel, _ := gtk.LabelNew(item.Name)

			if item.Connected {
				AddClass(row, "connected")
			}

			row.Add(nameLabel)

			AddClass(row, "listbox-row")

			eventBox, _ := gtk.EventBoxNew()
			eventBox.Add(row)
			eventBox.Connect("button-press-event", func() {
				fmt.Println("name: ", item.Uuid)
			})

			listBoxRow, _ := gtk.ListBoxRowNew()
			listBoxRow.Add(eventBox)

			//listRows = append(listRows, FilterRow{Text: item.Name, Container: listBoxRow})
			x.collapse.Add(listBoxRow)
		}
	})

	vpnList.Load()

	return vpnList
}
