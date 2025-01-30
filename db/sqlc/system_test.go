package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSystemInitialisation(t *testing.T) System {
	args := MarkInitialisationParams{
		Component:   "system",
		Initialised: true,
	}
	system, err := testQueries.MarkInitialisation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, system)
	require.Equal(t, args.Component, system.Component)
	require.Equal(t, args.Initialised, system.Initialised)
	return system
}

func TestSystemInitialisation(t *testing.T) {
	args := MarkInitialisationParams{
		Component:   "interfaces",
		Initialised: true,
	}
	system, err := testQueries.MarkInitialisation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, system)
	require.Equal(t, args.Component, system.Component)
	require.Equal(t, args.Initialised, system.Initialised)
}

func TestGetInitialisation(t *testing.T) {
	system1 := createRandomSystemInitialisation(t)
	system2, err := testQueries.GetInitialisation(context.Background(), system1.Component)
	require.NoError(t, err)
	require.NotEmpty(t, system2)
	require.Equal(t, system1.Component, system2.Component)
	require.Equal(t, system1.Initialised, system2.Initialised)
}
