package exchanges_test

import (
	"testing"

	keepertest "exchanges/testutil/keeper"
	"exchanges/testutil/nullify"
	"exchanges/x/exchanges"
	"exchanges/x/exchanges/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExchangesKeeper(t)
	exchanges.InitGenesis(ctx, *k, genesisState)
	got := exchanges.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
