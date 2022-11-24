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
	if msg.Amount.Denom != pool.Denom {
		return &types.MsgDepositResponse{}, sdkerrors.Wrapf(types.ErrIncorrectDenom, "input: %s, supported: %s", msg.Amount.Denom, pool.Denom)
	}

	poolAddr, _ := sdk.AccAddressFromBech32(pool.Address)
	err := k.bankKeeper.SendCoins(ctx, sender, poolAddr, sdk.NewCoins(msg.Amount))
	if err != nil {
		return &types.MsgDepositResponse{}, err
	}

	lpCoin := sdk.NewCoin(pool.ShareCoinDenom(), msg.Amount.Amount)
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

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolDepositedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(pool.Id)),
			sdk.NewAttribute(types.PoolEventAmount, fmt.Sprint(msg.Amount.Amount)),
			sdk.NewAttribute(types.PoolEventDenom, fmt.Sprint(msg.Amount.Denom))),
	)

	return &types.MsgDepositResponse{}, nil
}
