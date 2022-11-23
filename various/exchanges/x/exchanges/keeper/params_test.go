package keeper_test

import (
	"testing"

	testkeeper "exchanges/testutil/keeper"
	"exchanges/x/exchanges/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ExchangesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
