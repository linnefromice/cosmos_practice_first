package keeper

import (
	"github.com/linnefromice/cosmos_practice_first/tutorials.cosmos.network/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
