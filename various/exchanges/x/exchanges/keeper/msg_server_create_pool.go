package keeper

import (
	"context"
	"fmt"

	"exchanges/x/exchanges/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	count := k.GetPoolCount(ctx)
	base := fmt.Sprintf("%d-%s", count, msg.Denom)
	accAddr := sdk.AccAddress(address.Module(types.ModuleName, []byte(base)))

	accI := k.accountKeeper.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(accAddr),
			accAddr.String(),
		),
	)
	k.accountKeeper.SetAccount(ctx, accI)

	pool := types.Pool{
		Address:         accI.GetAddress().String(),
		IsActive:        true,
		NormalDeposited: 0,
		ConlyDeposited:  0,
		Borrowed:        0,
	}
	k.AppendPool(ctx, pool)

	return &types.MsgCreatePoolResponse{}, nil
}
