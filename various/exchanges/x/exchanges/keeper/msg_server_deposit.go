package keeper

import (
	"context"
	"fmt"

	"exchanges/x/exchanges/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(msg.Amount))
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

	return &types.MsgDepositResponse{}, nil
}
