package db

import (
	"context"
	"database/sql"
	"gateway-router-db/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomInterface(t *testing.T) Interface {
	ifaceArgs := CreateInterfaceParams{
		Macaddress: utils.RandomMAC(),
		Ipaddress:  utils.RandomIP(),
		Mtu:        utils.RandomMtu(),
		Type:       "vlan",
		Tag:        utils.RandomVlan(),
	}
	iface, err := testQueries.CreateInterface(context.Background(), ifaceArgs)
	require.NoError(t, err)
	require.NotEmpty(t, iface)
	// require.Equal(t, vlaniface.ID, iface.ID)
	require.Equal(t, ifaceArgs.Ipaddress, iface.Ipaddress)
	require.Equal(t, ifaceArgs.Macaddress, iface.Macaddress)
	require.Equal(t, ifaceArgs.Tag, iface.Tag)
	require.Equal(t, ifaceArgs.Mtu, iface.Mtu)
	require.Equal(t, ifaceArgs.Type, iface.Type)
	require.NotZero(t, iface.CreatedAt)
	require.NotZero(t, iface.ID)
	return iface
}
func TestCreateVlanInterface(t *testing.T) {
	ifaceArgs := CreateInterfaceParams{
		Macaddress: utils.RandomMAC(),
		Ipaddress:  utils.RandomIP(),
		Mtu:        utils.RandomMtu(),
		Type:       "vlan",
		Tag:        utils.RandomVlan(),
	}
	iface, err := testQueries.CreateInterface(context.Background(), ifaceArgs)
	require.NoError(t, err)
	require.NotEmpty(t, iface)
	require.Equal(t, ifaceArgs.Ipaddress, iface.Ipaddress)
	require.Equal(t, ifaceArgs.Macaddress, iface.Macaddress)
	require.Equal(t, ifaceArgs.Tag, iface.Tag)
	require.Equal(t, ifaceArgs.Mtu, iface.Mtu)
	require.Equal(t, ifaceArgs.Type, iface.Type)
	require.NotZero(t, iface.CreatedAt)
	require.NotZero(t, iface.ID)
}

func TestUpdateInterface(t *testing.T) {
	iface1 := createRandomInterface(t)
	updateArgs := UpdateInterfaceParams{
		ID:   iface1.ID,
		Name: utils.RandomName(),
	}
	iface2, err := testQueries.UpdateInterface(context.Background(), updateArgs)
	require.NoError(t, err)
	require.NotEmpty(t, iface2)
	require.Equal(t, iface1.Ipaddress, iface2.Ipaddress)
	require.Equal(t, iface1.Macaddress, iface2.Macaddress)
	require.Equal(t, iface1.Tag, iface2.Tag)
	require.Equal(t, iface1.Mtu, iface2.Mtu)
	require.Equal(t, iface1.Type, iface2.Type)
	require.NotZero(t, iface2.CreatedAt)
	require.NotZero(t, iface2.ID)
}

func TestGetInterface(t *testing.T) {
	iface1 := createRandomInterface(t)
	iface2, err := testQueries.GetInterface(context.Background(), iface1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, iface2)
	require.Equal(t, iface1.Ipaddress, iface2.Ipaddress)
	require.Equal(t, iface1.Macaddress, iface2.Macaddress)
	require.Equal(t, iface1.Tag, iface2.Tag)
	require.Equal(t, iface1.Mtu, iface2.Mtu)
	require.Equal(t, iface1.Type, iface2.Type)
	require.NotZero(t, iface2.CreatedAt)
	require.NotZero(t, iface2.ID)
}

func TestListInterface(t *testing.T) {

}

func DeleteInterface(t *testing.T) {
	iface1 := createRandomInterface(t)
	err := testQueries.DeleteInterface(context.Background(), iface1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
