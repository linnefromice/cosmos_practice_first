package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "nameservice/testutil/keeper"
	"nameservice/x/nameservice"
	"nameservice/x/nameservice/keeper"
	"nameservice/x/nameservice/types"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	// carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerForBuyName(t *testing.T) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.NameserviceKeeper(t)
	nameservice.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

// func TestBuyName(t *testing.T) {
// 	msgServer, _, context := setupMsgServerForBuyName(t)
// 	createResponse, err := msgServer.BuyName(context, &types.MsgBuyName{
// 		Creator: alice,
// 		Name:    "alice",
// 		Bid:     "1000",
// 	})
// 	require.Nil(t, err)
// 	require.EqualValues(t, types.MsgBuyNameResponse{}, *createResponse)
// }

func TestBuyNameWhenBidIsZero(t *testing.T) {
	msgServer, _, context := setupMsgServerForBuyName(t)
	createResponse, err := msgServer.BuyName(context, &types.MsgBuyName{
		Creator: alice,
		Name:    "alice",
		Bid:     "0",
	})
	require.Nil(t, createResponse)
	require.EqualValues(t,
		"Bid is less than min amount: insufficient funds",
		err.Error())
}
