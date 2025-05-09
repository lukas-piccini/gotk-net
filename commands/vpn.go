package commands

import (
	"fmt"
	"maps"
	"os/exec"
	"slices"
	"strings"
)

type VpnConnection struct {
	Name      string
	Uuid      string
	Connected bool
}

type Vpn struct {
	Connections []VpnConnection
}

func VpnNew() *Vpn {
	return &Vpn{}
}

func (v *Vpn) Load() {
	list, err := exec.Command("nmcli", "-t", "-f", "NAME,UUID,ACTIVE,TYPE", "con").Output()

	if err != nil {
		panic("Error reading vpn connections")
	}

	vpns := strings.Split(string(list), "\n")

	result := make(map[string]VpnConnection)

	for _, item := range vpns {
		if item == "" {
			continue
		}

		split := strings.Split(item, ":")
		name := split[0]
		uuid := split[1]
		connected := split[2] == "yes"
		isVpn := split[3] == "vpn"

		if !isVpn {
			continue
		}

		result[name] = VpnConnection{Name: name, Uuid: uuid, Connected: connected}
	}

	v.Connections = slices.Collect(maps.Values(result))
}

func (c *VpnConnection) ToggleConnection() {
	var command string

	if c.Connected {
		command = "down"
	} else {
		command = "up"
	}

	result, err := exec.Command("nmcli", "con", command, c.Uuid).Output()

	if err != nil {
		fmt.Println("Error connecting to vpn ", c.Name)
	}

	fmt.Println(string(result))
}
