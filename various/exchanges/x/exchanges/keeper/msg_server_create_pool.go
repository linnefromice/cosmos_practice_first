package keeper

import (
	"context"
	"fmt"

	"exchanges/x/exchanges/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// create ModuleAccount
	count := k.GetPoolCount(ctx)
	base := fmt.Sprintf("%d-%s", count, msg.DenomName)
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
		Denom:           msg.DenomSymbol,
		IsActive:        true,
		NormalDeposited: 0,
		ConlyDeposited:  0,
		Borrowed:        0,
	}
	id := k.AppendPool(ctx, pool)

	// register coin's denom
	k.bankKeeper.SetDenomMetaData(ctx, banktypes.Metadata{
		Base:    fmt.Sprintf("%d-%s", id, msg.DenomSymbol),
		Display: msg.DenomSymbol,
		Name:    msg.DenomName,
		Symbol:  msg.DenomSymbol,
	})
	k.bankKeeper.SetDenomMetaData(ctx, banktypes.Metadata{
		Base:    fmt.Sprintf("%d-share-%s", id, msg.DenomSymbol),
		Display: pool.ShareCoinDenom(),
		Name:    fmt.Sprintf("Share %s", msg.DenomName),
		Symbol:  pool.ShareCoinDenom(),
	})

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolDepositedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(id)),
			sdk.NewAttribute(types.PoolEventDenom, fmt.Sprint(msg.DenomSymbol))),
	)

	return &types.MsgCreatePoolResponse{
		PoolId: pool.Id,
	}, nil
}
