package entities

import "net"

type IPRange struct {
	ID        int        `json:"id"`
	Subnet    *net.IPNet `json:"subnet"`
	Parent    *IPRange   `json:"parent,omitempty"`
	Allocated bool       `json:"allocated"`
}
