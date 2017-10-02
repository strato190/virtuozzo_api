package main

type Instances struct {
	Vms []VM `json:"vm,omitempty"`
	Cts []CT `json:"ct,omitempty"`
}

type VM struct {
	Name     string `json:"name"`
	State    string `json:"state"`
	Astart   string `json:"autostart"`
	Ram      string `json:"memory"`
	CPU      string `json:"cpu"`
	Networks []Nets `json:"net,omitempty"`
}

type CT struct {
	Name     string `json:"name"`
	State    string `json:"state"`
	Astart   string `json:"autostart"`
	Ram      string `json:"memory"`
	CPU      string `json:"cpu"`
	Networks []Nets `json:"net,omitempty"`
}

type Nets struct {
	Device     string `json:"device"`
	Ip         string `json:"ip,omitempty"`
	Gateway    string `json:"gateway,omitempty"`
	Nameserver string `json:"nameserver,omitempty"`
}
