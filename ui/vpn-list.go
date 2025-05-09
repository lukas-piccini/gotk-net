package ui

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"gotk-net/net/commands"
	"strings"
)

type VpnList struct {
	ConnectionList
	Searchable
}

func VpnListNew(title string, filter Searchable) *VpnList {
	v := &VpnList{Searchable: filter}
	vpnList := connectionListNew(title, v.load, v.filter)

	v.ConnectionList = *vpnList
	vpnList.Load()

	return v
}

func (v *VpnList) load(x *ConnectionList) {
	x.ToggleLoading()

	vpnCommand := commands.VpnNew()
	vpnCommand.Load()

	x.ToggleLoading()

	for _, item := range vpnCommand.Connections {
		v := v
		item := item

		row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		nameLabel, _ := gtk.LabelNew(item.Name)

		if item.Connected {
			AddClass(row, "connected")
			AddClass(nameLabel, "connected-label")
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
				fmt.Println("name: ", item.Name)
				item.ToggleConnection()
			}
		})

		x.filterRow[item.Name] = listBoxRow
		x.collapse.Add(listBoxRow)

		x.Filter(v.GetFilter())
	}
}

func (v *VpnList) filter(x *ConnectionList, search string) {
	for key, item := range v.filterRow {
		if strings.Contains(strings.ToLower(key), strings.ToLower(search)) {
			item.Container.ShowAll()
		} else {
			item.Container.Hide()
		}
	}
}
