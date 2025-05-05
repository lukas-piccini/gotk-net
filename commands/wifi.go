package commands

import (
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
}

func Wifi() []WifiConnection {
	net, err := exec.Command("nmcli", "-t", "-f", "ACTIVE,SSID,SIGNAL", "dev", "wifi", "list", "--rescan", "yes").Output()

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

		if ssid == "" {
			continue
		}

		if err != nil {
			panic("Error converting signal value to number")
		}

		result[ssid] = WifiConnection{Connected: connected, Ssid: ssid, Signal: signal}
	}

	return slices.Collect(maps.Values(result))
}
