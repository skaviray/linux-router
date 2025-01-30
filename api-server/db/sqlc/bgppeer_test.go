package sqlc

import (
	"context"
	"database/sql"
	"gateway-router/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomBgpPeer(t *testing.T) BgpPeer {
	args := CreateBgpPeerParams{
		Name:            utils.RandomName(),
		AsNo:            utils.RandomAsNo(),
		NeighborAddress: utils.RandomIP(),
		LocalAs:         utils.RandomAsNo(),
	}
	peer, err := testQueries.CreateBgpPeer(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Name, peer.Name)
	require.Equal(t, args.AsNo, peer.AsNo)
	require.Equal(t, args.NeighborAddress, peer.NeighborAddress)
	require.Equal(t, args.LocalAs, peer.LocalAs)
	require.NotZero(t, peer.ID)
	require.NotZero(t, peer.CreatedAt)
	return peer
}

func TestCreateBgpPeer(t *testing.T) {
	args := CreateBgpPeerParams{
		Name:            utils.RandomName(),
		AsNo:            utils.RandomAsNo(),
		NeighborAddress: utils.RandomIP(),
		LocalAs:         utils.RandomAsNo(),
	}
	peer, err := testQueries.CreateBgpPeer(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Name, peer.Name)
	require.Equal(t, args.AsNo, peer.AsNo)
	require.Equal(t, args.NeighborAddress, peer.NeighborAddress)
	require.Equal(t, args.LocalAs, peer.LocalAs)
	require.NotZero(t, peer.ID)
	require.NotZero(t, peer.CreatedAt)
}

func TestGetBgpPeer(t *testing.T) {
	peer1 := CreateRandomBgpPeer(t)
	peer2, err := testQueries.GetBgpPeer(context.Background(), peer1.ID)
	require.NoError(t, err)
	require.Equal(t, peer1.Name, peer2.Name)
	require.Equal(t, peer1.AsNo, peer2.AsNo)
	require.Equal(t, peer1.LocalAs, peer2.LocalAs)
	require.Equal(t, peer1.NeighborAddress, peer2.NeighborAddress)
	require.Equal(t, peer1.ID, peer2.ID)
	require.Equal(t, peer1.CreatedAt, peer1.CreatedAt)
}

func TestListBgpPeer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomBgpPeer(t)
	}
	args := ListBgpPeersParams{
		Limit:  10,
		Offset: 0,
	}
	peers, err := testQueries.ListBgpPeers(context.Background(), args)

	require.NoError(t, err)
	for _, peer := range peers {
		require.NotEmpty(t, peer)
	}
}

func TestDeleteBgpPeer(t *testing.T) {
	peer1 := CreateRandomBgpPeer(t)
	err := testQueries.DeleteBgpPeer(context.Background(), peer1.ID)
	require.NoError(t, err)
	peer2, err := testQueries.GetVxlanTunnel(context.Background(), peer1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, peer2)
}
