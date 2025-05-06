package commands

import (
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

func Vpn() []VpnConnection {
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

	return slices.Collect(maps.Values(result))
}
