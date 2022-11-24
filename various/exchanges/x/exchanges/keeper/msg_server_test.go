package keeper_test

import (
	"context"
	"testing"

	keepertest "exchanges/testutil/keeper"
	"exchanges/x/exchanges/keeper"
	"exchanges/x/exchanges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ExchangesKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
