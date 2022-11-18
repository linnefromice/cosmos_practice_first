package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/linnefromice/checkers/testutil/keeper"
	"github.com/linnefromice/checkers/x/checkers"
	"github.com/linnefromice/checkers/x/checkers/keeper"
	"github.com/linnefromice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneGameForPlayMove(t *testing.T) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	return server, *k, context
}

func TestPlayMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgPlayMoveResponse{
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

// func TestPlayMoveGameNotFound
// func TestPlayMoveSameBlackRed
// func TestPlayMoveSavedGame
// func TestPlayMoveNotPlayer
// func TestPlayMoveCannotParseGame
// func TestPlayMoveWrongOutOfTurn
// func TestPlayMoveWrongPieceAtDestination
// func TestPlayMove2
// func TestPlayMove2SavedGame
// func TestPlayMove3
// func TestPlayMove3SavedGame
