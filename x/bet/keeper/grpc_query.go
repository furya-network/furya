package keeper

import (
	"github.com/furya-network/furya/x/bet/types"
)

var _ types.QueryServer = Keeper{}
