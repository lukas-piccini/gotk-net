package commands

import (
	"fmt"
	"maps"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

type WifiConnection struct {
	Connected bool
	Ssid      string
	Signal    int
	Protected bool
}

type Wifi struct {
	Connections []WifiConnection
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
		signal, err := strconv.Atoi(line[2])
		protected := !(line[3] == "" || line[3] == "--")

		if ssid == "" {
			continue
		}

		if err != nil {
			panic("Error converting signal value to number")
		}

		result[ssid] = WifiConnection{Connected: connected, Ssid: ssid, Signal: signal, Protected: protected}
	}

	w.Connections = slices.Collect(maps.Values(result))
}

func (c *WifiConnection) ToggleConnection() {
	var command string

	if c.Connected {
		command = "disconnect"
	} else {
		command = "connect"
	}

	result, err := exec.Command("nmcli", "dev", "wifi", command, c.Ssid).Output()

	if err != nil {
		fmt.Println("Error connecting to network ", c.Ssid)
	}

	fmt.Println(string(result))
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
