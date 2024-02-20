package keeper

import (
	"github.com/furya-network/furya/x/market/types"
)

var _ types.QueryServer = Keeper{}
