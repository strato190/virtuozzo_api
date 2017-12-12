package main

import (
	"encoding/json"
	"strings"
)

func getCTS() ([]byte, error) {
	var v string
	var err error
	commands := make([][]string, 1)
	commands[0] = []string{"list", "--vmtype", "ct", "-o", "name", "-Ha"}
	for _, command := range commands {
		v, err = Prlctl(command...)
		if err != nil {
			return []byte(v), err
		}
	}
	jsnout, _ := json.Marshal(strings.Split(v, "\n"))
	return jsnout, nil

}

func createCT(c CT) error {
	cName := c.Name
	cTemplate := c.Template
	command := []string{"create", cName, "--ostemplate", cTemplate}
	_, err := Vzctl(command...)
	if err != nil {
		return err
	}
	return nil
}

func configCT(c CT) error {
	var diskSize string

	cName := c.Name
	commands := make([][]string, 5)

	if c.DISK == "" {
		diskSize = "10G"
	} else {
		diskSize = c.DISK
	}
	commands[0] = []string{"set", cName, "--cpus", c.CPU}
	commands[1] = []string{"set", cName, "--memsize", c.RAM}
	commands[2] = []string{"set", cName, "--autostart", c.Astart}
	commands[3] = []string{"set", cName, "--hostname", cName}
	commands[4] = []string{"set", cName, "--device-set hdd0", "--size", diskSize}

	for _, command := range commands {
		_, err := Vzctl(command...)
		if err != nil {
			return err
		}
	}
	return nil

}
