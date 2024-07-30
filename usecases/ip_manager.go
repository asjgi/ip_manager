package usecases

import (
	"errors"
	"fmt"
	"net"

	"ip_manager/entities"
)

type IPManager struct {
	ipRanges map[int]*entities.IPRange
	nextID   int
}

func NewIPManager() *IPManager {
	return &IPManager{
		ipRanges: make(map[int]*entities.IPRange),
		nextID:   1,
	}
}

func (m *IPManager) CreateSubnet(cidr string, parentID int) (*entities.IPRange, error) {
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var parent *entities.IPRange
	if parentID != 0 {
		var ok bool
		parent, ok = m.ipRanges[parentID]
		if !ok {
			return nil, errors.New("parent ID not found")
		}
	}

	ipRange := &entities.IPRange{
		ID:     m.nextID,
		Subnet: subnet,
		Parent: parent,
	}

	m.ipRanges[ipRange.ID] = ipRange
	m.nextID++
	return ipRange, nil
}

func (m *IPManager) DeleteSubnet(id int) error {
	if _, ok := m.ipRanges[id]; !ok {
		return errors.New("subnet ID not found")
	}

	for _, ipRange := range m.ipRanges {
		if ipRange.Parent != nil && ipRange.Parent.ID == id {
			return errors.New("cannot delete subnet with existing children")
		}
	}

	delete(m.ipRanges, id)
	return nil
}

func (m *IPManager) AllocateSubnet(parentID int, cidr int) (*entities.IPRange, error) {
	parent, ok := m.ipRanges[parentID]
	if !ok {
		return nil, errors.New("parent ID not found")
	}

	parentIP := parent.Subnet.IP
	maskSize, _ := parent.Subnet.Mask.Size()
	for i := 0; i < (1 << (cidr - maskSize)); i++ {
		childIP := make(net.IP, len(parentIP))
		copy(childIP, parentIP)
		childIP[3] += byte(i << (32 - cidr))
		_, childSubnet, err := net.ParseCIDR(fmt.Sprintf("%s/%d", childIP.String(), cidr))
		if err != nil {
			continue
		}

		conflict := false
		for _, ipRange := range m.ipRanges {
			if ipRange.Subnet.Contains(childSubnet.IP) {
				conflict = true
				break
			}
		}

		if !conflict {
			return m.CreateSubnet(childSubnet.String(), parentID)
		}
	}

	return nil, errors.New("no available subnets")
}

func (m *IPManager) ReleaseSubnet(id int) error {
	ipRange, ok := m.ipRanges[id]
	if !ok {
		return errors.New("subnet ID not found")
	}

	if ipRange.Parent == nil {
		return errors.New("cannot release root subnet")
	}

	return m.DeleteSubnet(id)
}

func (m *IPManager) CheckIPAllocated(id int) (bool, error) {
	ipRange, ok := m.ipRanges[id]
	if !ok {
		return false, errors.New("subnet ID not found")
	}
	return ipRange.Allocated, nil
}
