package models

type MachinesList []*Machine

type Machine struct {
	ID   uint   `json:"id"`
	Name string `json:"name,omitempty"`
	IP   string `json:"ip,omitempty"`
}
