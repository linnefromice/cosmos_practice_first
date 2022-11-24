package keeper

import (
	"context"
	"fmt"

	"exchanges/x/exchanges/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	lpDenom := fmt.Sprintf("share%s", msg.Amount.Denom)
	lpCoin := sdk.NewCoin(lpDenom, msg.Amount.Amount)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(lpCoin))
	if err != nil {
		return &types.MsgWithdrawResponse{}, err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
	if err != nil {
		return &types.MsgWithdrawResponse{}, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(msg.Amount))
	if err != nil {
		return &types.MsgWithdrawResponse{}, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
