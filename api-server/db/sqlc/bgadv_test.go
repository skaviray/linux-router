package sqlc

import (
	"context"
	"database/sql"
	"gateway-router/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomBGPAdv(t *testing.T) BgpAdvertisement {
	args := CreateBgpAdvertisementParams{
		Name:            utils.RandomName(),
		DestinationCidr: utils.RandomCIDR(),
	}
	adv, err := testQueries.CreateBgpAdvertisement(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Name, adv.Name)
	require.Equal(t, args.DestinationCidr, adv.DestinationCidr)
	require.NotZero(t, adv.ID)
	require.NotZero(t, adv.CreatedAt)
	return adv
}
func TestCreateBGPAdv(t *testing.T) {
	args := CreateBgpAdvertisementParams{
		Name:            utils.RandomName(),
		DestinationCidr: utils.RandomCIDR(),
	}
	adv, err := testQueries.CreateBgpAdvertisement(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.Name, adv.Name)
	require.Equal(t, args.DestinationCidr, adv.DestinationCidr)
	require.NotZero(t, adv.ID)
	require.NotZero(t, adv.CreatedAt)
}

func TestGetBGPAdv(t *testing.T) {
	adv1 := CreateRandomBGPAdv(t)
	adv2, err := testQueries.GetBgpAdvertisement(context.Background(), adv1.ID)
	require.NoError(t, err)
	require.Equal(t, adv1.Name, adv2.Name)
	require.Equal(t, adv1.DestinationCidr, adv2.DestinationCidr)
	require.Equal(t, adv1.ID, adv2.ID)
	require.Equal(t, adv1.CreatedAt, adv2.CreatedAt)
}

func TestListBGPAdv(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomBGPAdv(t)
	}
	args := ListBgpAdvertisementsParams{
		Limit:  10,
		Offset: 0,
	}
	advs, err := testQueries.ListBgpAdvertisements(context.Background(), args)

	require.NoError(t, err)
	for _, adv := range advs {
		require.NotEmpty(t, adv)
	}
}

func TestDeleteBGPAdv(t *testing.T) {
	adv1 := CreateRandomBGPAdv(t)
	err := testQueries.DeleteBgpAdvertisement(context.Background(), adv1.ID)
	require.NoError(t, err)
	adv2, err := testQueries.GetBgpAdvertisement(context.Background(), adv1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, adv2)
}
