package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-network/furya/utils"
	"github.com/furya-network/furya/x/bet/types"
)

// SetBetStats sets bet statistics in the store
func (k Keeper) SetBetStats(ctx sdk.Context, stats types.BetStats) {
	store := k.getBetStatsStore(ctx)
	b := k.cdc.MustMarshal(&stats)
	store.Set(utils.StrBytes("0"), b)
}

// GetBet returns bet stats
func (k Keeper) GetBetStats(ctx sdk.Context) (val types.BetStats) {
	store := k.getBetStatsStore(ctx)

	b := store.Get(utils.StrBytes("0"))
	if b == nil {
		return val
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}
