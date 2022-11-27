package keeper

import (
	"atomic/x/atomic/types"
)

var _ types.QueryServer = Keeper{}
