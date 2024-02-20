package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furya-network/furya/x/house/types"
)

// Deposit defines a method for performing a deposit of coins to become part of the house correspondifg to a market.
func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.validateDeposit(ctx, msg); err != nil {
		return nil, sdkerrors.Wrap(err, "validate deposit")
	}

	participationIndex, err := k.Keeper.Deposit(ctx, msg.Creator, msg.MarketUID, msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "process deposit")
	}

	return &types.MsgDepositResponse{
		MarketUID:          msg.MarketUID,
		ParticipationIndex: participationIndex,
	}, nil
}
