package ui

import (
	"fmt"
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

	connections := commands.Vpn()

	x.ToggleLoading()

	for _, item := range connections {
		v := v
		item := item

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
