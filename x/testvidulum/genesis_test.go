package testvidulum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vidulum/testvidulum/testutil/keeper"
	"github.com/vidulum/testvidulum/x/testvidulum"
	"github.com/vidulum/testvidulum/x/testvidulum/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestvidulumKeeper(t)
	testvidulum.InitGenesis(ctx, *k, genesisState)
	got := testvidulum.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
