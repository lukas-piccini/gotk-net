package commands

import (
	"os/exec"
	"strings"
)

type DefaultConnection struct {
	Name      string
	Uuid      string
	Connected bool
}

type Connection struct {
	Connections []DefaultConnection
}

func ConnectionNew() *Connection {
	return &Connection{}
}

func (c *Connection) Load() {
	list, err := exec.Command("nmcli", "-t", "-f", "NAME,UUID,ACTIVE", "con").Output()

	if err != nil {
		panic("Error reading connections")
	}

	connections := strings.Split(string(list), "\n")

	for _, item := range connections {
		if item == "" {
			continue
		}

		split := strings.Split(item, ":")
		name := split[0]
		uuid := split[1]
		connected := split[2] == "yes"

		c.Connections = append(c.Connections, DefaultConnection{Name: name, Uuid: uuid, Connected: connected})
	}
}
