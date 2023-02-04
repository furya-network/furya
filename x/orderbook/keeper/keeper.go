package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sge-network/sge/x/orderbook/types"
)

// keeper of the orderbook store
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	paramstore    paramtypes.Subspace
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
}

// NewKeeper creates a new orderbook Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, ps paramtypes.Subspace, bnkKeeper types.BankKeeper, accKeeper types.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		storeKey:      key,
		cdc:           cdc,
		paramstore:    ps,
		bankKeeper:    bnkKeeper,
		accountKeeper: accKeeper,
	}
}
