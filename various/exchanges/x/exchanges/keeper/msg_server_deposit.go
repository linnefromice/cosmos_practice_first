package keeper

import (
	"context"
	"fmt"

	"exchanges/x/exchanges/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return &types.MsgDepositResponse{}, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.PoolId))
	}

	poolAddr, _ := sdk.AccAddressFromBech32(pool.Address)
	err := k.bankKeeper.SendCoins(ctx, sender, poolAddr, sdk.NewCoins(msg.Amount))
	if err != nil {
		return &types.MsgDepositResponse{}, err
	}

	lpDenom := fmt.Sprintf("share%s", msg.Amount.Denom)
	lpCoin := sdk.NewCoin(lpDenom, msg.Amount.Amount)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
	if err != nil {
		return &types.MsgDepositResponse{}, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(lpCoin))
	if err != nil {
		return &types.MsgDepositResponse{}, err
	}

	// update state
	pool.NormalDeposited += msg.Amount.Amount.Uint64()
	k.SetPool(ctx, pool)

	return &types.MsgDepositResponse{}, nil
}
