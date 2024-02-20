package keeper

import (
	"github.com/furya-network/furya/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
