package keeper

import (
	"github.com/furya-network/furya/x/house/types"
)

var _ types.QueryServer = Keeper{}
