package keeper

import (
	"github.com/furya-network/furya/x/mint/types"
)

var _ types.QueryServer = Keeper{}
