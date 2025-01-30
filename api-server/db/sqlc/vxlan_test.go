package sqlc

import (
	"context"
	"database/sql"
	"gateway-router/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomVxlan(t *testing.T) VxlanTunnel {
	args := CreateVxlanTunnelParams{
		Name:      utils.RandomName(),
		Tag:       utils.RandomVxlan(),
		TunnelIp:  utils.RandomIpInCidr(),
		RemoteIp:  utils.RandomIP(),
		RemoteMac: utils.RandomMAC(),
	}
	tunnel, err := testQueries.CreateVxlanTunnel(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Name, tunnel.Name)
	require.Equal(t, args.Tag, tunnel.Tag)
	require.Equal(t, args.TunnelIp, tunnel.TunnelIp)
	require.Equal(t, args.RemoteIp, tunnel.RemoteIp)
	require.Equal(t, args.RemoteMac, tunnel.RemoteMac)
	require.NotZero(t, tunnel.ID)
	require.NotZero(t, tunnel.CreatedAt)
	return tunnel

}
func TestCreateVxlan(t *testing.T) {
	args := CreateVxlanTunnelParams{
		Name:      utils.RandomName(),
		Tag:       utils.RandomVxlan(),
		TunnelIp:  utils.RandomIpInCidr(),
		RemoteIp:  utils.RandomIP(),
		RemoteMac: utils.RandomMAC(),
	}
	tunnel, err := testQueries.CreateVxlanTunnel(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Name, tunnel.Name)
	require.Equal(t, args.Tag, tunnel.Tag)
	require.Equal(t, args.TunnelIp, tunnel.TunnelIp)
	require.Equal(t, args.RemoteIp, tunnel.RemoteIp)
	require.Equal(t, args.RemoteMac, tunnel.RemoteMac)
	require.NotZero(t, tunnel.ID)
	require.NotZero(t, tunnel.CreatedAt)
}

func TestGetVxlan(t *testing.T) {
	tunnel1 := CreateRandomVxlan(t)
	tunnel2, err := testQueries.GetVxlanTunnel(context.Background(), tunnel1.ID)
	require.NoError(t, err)
	require.Equal(t, tunnel1.Name, tunnel2.Name)
	require.Equal(t, tunnel1.Tag, tunnel2.Tag)
	require.Equal(t, tunnel1.TunnelIp, tunnel2.TunnelIp)
	require.Equal(t, tunnel1.RemoteIp, tunnel2.RemoteIp)
	require.Equal(t, tunnel1.RemoteMac, tunnel2.RemoteMac)
	require.Equal(t, tunnel1.ID, tunnel2.ID)
	require.Equal(t, tunnel1.CreatedAt, tunnel2.CreatedAt)
}

func TestListVxlan(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomVxlan(t)
	}
	args := ListVxlanTunnelParams{
		Limit:  10,
		Offset: 0,
	}
	tunnels, err := testQueries.ListVxlanTunnel(context.Background(), args)

	require.NoError(t, err)
	for _, tunnel := range tunnels {
		require.NotEmpty(t, tunnel)
	}
}

func TestDeleteVxlan(t *testing.T) {
	tunnel1 := CreateRandomVxlan(t)
	err := testQueries.DeleteVxlanTunnel(context.Background(), tunnel1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetVxlanTunnel(context.Background(), tunnel1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestUpdateVxlanStatus(t *testing.T) {
	tunnel1 := CreateRandomVxlan(t)
	args := UpdateVxlanStatusParams{
		ID:     tunnel1.ID,
		Status: "created",
	}
	err := testQueries.UpdateVxlanStatus(context.Background(), args)
	require.NoError(t, err)
	tunnel2, err := testQueries.GetVxlanTunnel(context.Background(), tunnel1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tunnel2)
	require.Equal(t, tunnel1.Name, tunnel2.Name)
	require.Equal(t, tunnel1.Tag, tunnel2.Tag)
	require.Equal(t, tunnel1.TunnelIp, tunnel2.TunnelIp)
	require.Equal(t, tunnel1.RemoteIp, tunnel2.RemoteIp)
	require.Equal(t, tunnel1.RemoteMac, tunnel2.RemoteMac)
	require.Equal(t, tunnel1.ID, tunnel2.ID)
	require.Equal(t, args.Status, tunnel2.Status)
	require.Equal(t, tunnel1.CreatedAt, tunnel2.CreatedAt)

}
