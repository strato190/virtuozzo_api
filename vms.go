package main

func config_vm_network(v VM) error {
	commands := make([][]string, 0)
	vmName := v.Name
	for _, n := range v.Networks {
		if n.Device != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device})
		}
		if n.Ip != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device, "--ipadd", n.Ip})
		}
		if n.Gateway != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device, "--gw", n.Gateway})
		}
		if n.Nameserver != "" {
			commands = append(commands, []string{"set", vmName, "--nameserver", n.Nameserver})
		}
	}
	for _, command := range commands {
		err := Prlctl(command...)
		if err != nil {
			return err
		}
	}
	return nil
}

func create_vm(v VM) error {
	vmName := v.Name

	command := []string{"create", vmName, "--ostemplate", "vm-template-centos7"}

	err := Prlctl(command...)

	if err != nil {
		return err
	}
	return nil

}

func config_vm(v VM) error {
	vmName := v.Name
	commands := make([][]string, 4)
	commands[0] = []string{"set", vmName, "--cpus", v.CPU}
	commands[1] = []string{"set", vmName, "--memsize", v.Ram}
	commands[2] = []string{"set", vmName, "--autostart", v.Astart}
	commands[3] = []string{"set", vmName, "--hostname", vmName}

	for _, command := range commands {
		err := Prlctl(command...)
		if err != nil {
			return err
		}
	}
	return nil

}

func config_network(v VM) error {
	commands := make([][]string, 0)
	vmName := v.Name
	for _, n := range v.Networks {
		if n.Device != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device})
		}
		if n.Ip != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device, "--ipadd", n.Ip})
		}
		if n.Gateway != "" {
			commands = append(commands, []string{"set", vmName, "--device-set", n.Device, "--gw", n.Gateway})
		}
		if n.Nameserver != "" {
			commands = append(commands, []string{"set", vmName, "--nameserver", n.Nameserver})
		}
	}
	for _, command := range commands {
		err := Prlctl(command...)
		if err != nil {
			return err
		}
	}
	return nil
}
