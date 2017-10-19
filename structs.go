package main

//Instances struct contains slices of vm's and container's
type Instances struct {
	Vms []VM `json:"vm,omitempty"`
	Cts []CT `json:"ct,omitempty"`
}

//Hostvm struct for host vm
type Hostvm struct {
	Name string `json:"name"`
}

//VM virtual machine struct
type VM struct {
	Name     string `json:"name"`
	State    string `json:"state"`
	Astart   string `json:"autostart"`
	RAM      string `json:"memory"`
	CPU      string `json:"cpu"`
	Networks []Nets `json:"net,omitempty"`
}

//CT container struct
type CT struct {
	Name     string `json:"name"`
	State    string `json:"state"`
	Astart   string `json:"autostart"`
	RAM      string `json:"memory"`
	CPU      string `json:"cpu"`
	Networks []Nets `json:"net,omitempty"`
}

//Nets struct describes vm or container network interfaces
type Nets struct {
	Device     string `json:"device"`
	IP         string `json:"ip,omitempty"`
	Gateway    string `json:"gateway,omitempty"`
	Nameserver string `json:"nameserver,omitempty"`
}
