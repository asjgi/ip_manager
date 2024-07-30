package usecases

import (
	"github.com/stretchr/testify/assert"
	usecase "ip_manager/usecases"
	"testing"
)

func TestCreateSubnet(t *testing.T) {
	manager := usecase.NewIPManager()
	ipRange, err := manager.CreateSubnet("10.0.0.0/24", 0)

	assert.NoError(t, err)
	assert.NotNil(t, ipRange)
	assert.Equal(t, 1, ipRange.ID)
	assert.Equal(t, "10.0.0.0/24", ipRange.Subnet.String())
	assert.Nil(t, ipRange.Parent)
}

func TestDeleteSubnet(t *testing.T) {
	manager := usecase.NewIPManager()
	ipRange, err := manager.CreateSubnet("10.0.0.0/24", 0)
	assert.NoError(t, err)
	assert.NotNil(t, ipRange)

	err = manager.DeleteSubnet(ipRange.ID)
	assert.NoError(t, err)
}

func TestAllocateSubnet(t *testing.T) {
	manager := usecase.NewIPManager()
	_, err := manager.CreateSubnet("10.0.0.0/24", 0)
	assert.NoError(t, err)

	ipRange, err := manager.AllocateSubnet(1, 30)
	assert.NoError(t, err)
	assert.NotNil(t, ipRange)
	assert.Equal(t, "10.0.0.0/30", ipRange.Subnet.String())
}

func TestReleaseSubnet(t *testing.T) {
	manager := usecase.NewIPManager()
	_, err := manager.CreateSubnet("10.0.0.0/24", 0)
	assert.NoError(t, err)

	ipRange, err := manager.AllocateSubnet(1, 30)
	assert.NoError(t, err)
	assert.NotNil(t, ipRange)

	err = manager.ReleaseSubnet(ipRange.ID)
	assert.NoError(t, err)
}
