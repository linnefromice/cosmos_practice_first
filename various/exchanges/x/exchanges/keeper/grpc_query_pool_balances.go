package keeper

import (
	"context"

	"exchanges/x/exchanges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolBalances(goCtx context.Context, req *types.QueryPoolBalancesRequest) (*types.QueryPoolBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryPoolBalancesResponse{}, nil
}
