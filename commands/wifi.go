package commands

import (
	"cmp"
	"fmt"
	"maps"
	"os/exec"
	"slices"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

type WifiConnection struct {
	Connected bool
	Ssid      string
	Signal    int
	Protected bool
}

type Wifi struct {
	Connections *[]WifiConnection
}

func WifiNew() *Wifi {
	return &Wifi{}
}

func (w *Wifi) Load() {
	net, err := exec.Command("nmcli", "-t", "-f", "ACTIVE,SSID,SIGNAL,SECURITY", "dev", "wifi", "list", "--rescan", "yes").Output()

	if err != nil {
		panic("Error reading available wi-fi connections.")
	}

	networks := strings.Split(string(net), "\n")

	result := make(map[string]WifiConnection)

	for _, network := range networks {
		if network == "" {
			continue
		}

		line := strings.Split(network, ":")
		connected := line[0] == "yes"
		ssid := line[1]
		signal, _ := strconv.Atoi(line[2])
		protected := !(line[3] == "" || line[3] == "--")

		if ssid == "" {
			continue
		}

		conn := WifiConnection{Connected: connected, Ssid: ssid, Signal: signal, Protected: protected}

		result[ssid] = conn
	}

	resultArray := slices.Collect(maps.Values(result))
	slices.SortFunc(resultArray, sortWifi)

	w.Connections = &resultArray
}

func sortWifi(a, b WifiConnection) int {
	return cmp.Compare(b.Signal, a.Signal)
}

func (c *WifiConnection) ToggleConnection(exists bool) {
	command := []string{"nmcli"}

	if exists {
		command = append(command, "con")

		if c.Connected {
			command = append(command, "down")
		} else {
			command = append(command, "up")
		}
	} else {
		command = append(command, "dev")
		command = append(command, "wifi")

		if c.Connected {
			command = append(command, "disconnect")
		} else {
			command = append(command, "connect")
		}
	}

	command = append(command, c.Ssid)

	fmt.Println(command)

	result, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		fmt.Println("Error connecting to network ", c.Ssid)
	}

	if result != nil {
		gtk.MainQuit()
	}
}

func (c *WifiConnection) CheckIfConnectionAlreadyExists(connections []DefaultConnection) bool {
	contains := false

	for _, x := range connections {
		if x.Name == c.Ssid {
			contains = true
		}
	}

	return contains
}

func (c *WifiConnection) GetPowerClass() string {
	if c.Signal > 85 {
		return "wifi-strong"
	} else if c.Signal > 60 {
		return "wifi-high"
	} else if c.Signal > 30 {
		return "wifi-medium"
	}

	return "wifi-low"
}
