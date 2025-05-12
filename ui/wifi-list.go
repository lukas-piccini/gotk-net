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
	Searchable
	WithPassword
	connection *commands.Connection
}

func WifiListNew(title string, app SearchableWithPassword, connections *commands.Connection) *WifiList {
	w := &WifiList{Searchable: app, WithPassword: app, connection: connections}
	connList := connectionListNew(title, w.load, w.filter)
	w.ConnectionList = *connList
	w.Load()

	return w
}

func (w *WifiList) load(x *ConnectionList) {
	x.ToggleLoading()
	go func() {
		wifiCommand := commands.WifiNew()
		wifiCommand.Load()

		glib.IdleAdd(func() {
			x.ToggleLoading()

			for _, item := range *wifiCommand.Connections {
				w := w
				item := item
				row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)

				nameLabel, _ := gtk.LabelNew(item.Ssid)
				nameLabel.SetXAlign(0)
				row.PackStart(nameLabel, true, true, 0)

				connectionPowerIcon, _ := gtk.LabelNew("\uf1eb")
				//AddClass(connectionPowerIcon, item.GetPowerClass())
				connectionPowerIcon.SetUseMarkup(true)
				connectionPowerIcon.SetXAlign(1)
				row.PackEnd(connectionPowerIcon, false, false, 0)

				if item.Protected {
					protectionIcon, _ := gtk.LabelNew("\uf023")
					protectionIcon.SetUseMarkup(true)
					protectionIcon.SetXAlign(1)
					row.PackEnd(protectionIcon, false, false, 16)
				}
				if item.Connected {
					AddClass(row, "connected")
					AddClass(nameLabel, "connected-label")
				}

				AddClass(row, "listbox-row")

				eventBox, _ := gtk.EventBoxNew()
				eventBox.Add(row)

				listBoxRow, _ := gtk.ListBoxRowNew()
				listBoxRow.Add(eventBox)

				listBoxRow.Connect("key-press-event", func(win *gtk.ListBoxRow, ev *gdk.Event) {
					key := gdk.EventKeyNewFromEvent(ev)

					if key.KeyVal() == gdk.KEY_Return {
						fmt.Println("name: ", item.Ssid)
						connectionAlreadyExists := item.CheckIfConnectionAlreadyExists(w.connection.Connections)

						if item.Protected && !connectionAlreadyExists {
							w.SetSelectedItem(&item)
							w.TogglePassword()
						} else {
							item.ToggleConnection(connectionAlreadyExists, "")
						}
					}
				})

				x.filterRow[item.Ssid] = listBoxRow
				x.collapse.List.Add(listBoxRow)
				listBoxRow.ShowAll()

				x.Filter(w.GetFilter())
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
