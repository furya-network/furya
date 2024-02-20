package keeper

import (
	"github.com/furya-network/furya/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
