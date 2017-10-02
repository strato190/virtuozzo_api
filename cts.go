package main

func create_ct(c CT) error {
	cName := c.Name

	command := []string{"create", cName, "--ostemplate", "vm-template-centos7"}

	err := Vzctl(command...)

	if err != nil {
		return err
	}
	return nil

}

func config_ct(c CT) error {
	cName := c.Name
	commands := make([][]string, 4)
	commands[0] = []string{"set", cName, "--cpus", c.CPU}
	commands[1] = []string{"set", cName, "--memsize", c.Ram}
	commands[2] = []string{"set", cName, "--autostart", c.Astart}
	commands[3] = []string{"set", cName, "--hostname", cName}

	for _, command := range commands {
		err := Vzctl(command...)
		if err != nil {
			return err
		}
	}
	return nil

}
