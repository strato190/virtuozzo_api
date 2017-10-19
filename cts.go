package main

func getCTS() (string, error) {
	var v string
	var err error
	commands := make([][]string, 1)
	commands[0] = []string{"list", "--vmtype", "ct", "-o", "name", "-H"}
	for _, command := range commands {
		v, err = Vzctl(command...)
		if err != nil {
			return v, err
		}
	}
	return v, nil

}

func createCT(c CT) error {
	cName := c.Name
	command := []string{"create", cName, "--ostemplate", "vm-template-centos7"}
	_, err := Vzctl(command...)
	if err != nil {
		return err
	}
	return nil
}

func configCT(c CT) error {
	cName := c.Name
	commands := make([][]string, 4)
	commands[0] = []string{"set", cName, "--cpus", c.CPU}
	commands[1] = []string{"set", cName, "--memsize", c.RAM}
	commands[2] = []string{"set", cName, "--autostart", c.Astart}
	commands[3] = []string{"set", cName, "--hostname", cName}

	for _, command := range commands {
		_, err := Vzctl(command...)
		if err != nil {
			return err
		}
	}
	return nil

}
