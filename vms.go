package main

import (
	"encoding/json"
	"strings"
)

func getVMS() ([]byte, error) {
	var v string
	var err error
	commands := make([][]string, 1)
	commands[0] = []string{"list", "--vmtype", "vm", "-o", "name", "-H"}
	for _, command := range commands {
		v, err = Prlctl(command...)
		if err != nil {
			return []byte(v), err
		}
	}
	jsnout, _ := json.Marshal(strings.Split(v, "\n"))
	return jsnout, nil

}

func configVMNetwork(v VM) error {
	commands := make([][]string, 0)
	vmName := v.Name
	for _, n := range v.Networks {
		if n.Device != "" {
			if n.Device != "net0" {
				commands = append(commands, []string{"set", vmName, "--device-add", "net"})
			}
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device})
		}
		if n.IP != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device, "--ipadd", n.IP})
		}
		if n.Gateway != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device, "--gw", n.Gateway})
		}
		if n.Nameserver != "" {
			commands = append(commands, []string{"set", vmName, "--nameserver", n.Nameserver})
		}
	}
	for _, command := range commands {
		_, err := Prlctl(command...)
		if err != nil {
			return err
		}
	}
	return nil
}

func createVM(v VM) error {
	vmName := v.Name
	vmTemplate := v.Template

	command := []string{"create", vmName, "--ostemplate", vmTemplate}

	_, err := Prlctl(command...)

	if err != nil {
		return err
	}
	return nil

}

func startVM(v VM) error {
	vmName := v.Name

	command := []string{"start", vmName}

	_, err := Prlctl(command...)

	if err != nil {
		return err
	}
	return nil

}

func stopVM(v VM) error {
	vmName := v.Name

	command := []string{"start", vmName}

	_, err := Prlctl(command...)

	if err != nil {
		return err
	}
	return nil

}

func destroyVM(v VM) error {
	vmName := v.Name

	if err := stopVM(v); err != nil {
		return err
	}

	command := []string{"destroy", vmName}

	_, err := Prlctl(command...)

	if err != nil {
		return err
	}
	return nil

}

func configVM(v VM) error {
	var diskSize string
	vmName := v.Name
	commands := make([][]string, 5)
	if v.DISK == "" {
		diskSize = "10G"
	} else {
		diskSize = v.DISK
	}
	commands[0] = []string{"set", vmName, "--cpus", v.CPU}
	commands[1] = []string{"set", vmName, "--memsize", v.RAM}
	commands[2] = []string{"set", vmName, "--autostart", v.Astart}
	commands[3] = []string{"set", vmName, "--hostname", vmName}
	commands[4] = []string{"set", vmName, "--device-set hdd0", "--size", diskSize}

	for _, command := range commands {
		_, err := Prlctl(command...)
		if err != nil {
			return err
		}
	}
	return nil

}
