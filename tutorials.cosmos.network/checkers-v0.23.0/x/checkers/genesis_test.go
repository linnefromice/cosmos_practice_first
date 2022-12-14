package checkers_test

import (
	"testing"

	keepertest "github.com/linnefromice/cosmos_practice_first/tutorials.cosmos.network/checkers/testutil/keeper"
	"github.com/linnefromice/cosmos_practice_first/tutorials.cosmos.network/checkers/testutil/nullify"
	"github.com/linnefromice/cosmos_practice_first/tutorials.cosmos.network/checkers/x/checkers"
	"github.com/linnefromice/cosmos_practice_first/tutorials.cosmos.network/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: &types.SystemInfo{
			NextId: 89,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
