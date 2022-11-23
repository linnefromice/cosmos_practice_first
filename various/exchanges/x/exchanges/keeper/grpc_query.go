package keeper

import (
	"exchanges/x/exchanges/types"
)

var _ types.QueryServer = Keeper{}
