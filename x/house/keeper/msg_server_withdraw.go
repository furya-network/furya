package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furya-network/furya/x/house/types"
)

// Withdraw defines a method for performing a withdrawal of coins of unused amount corresponding to a deposit.
func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := k.Keeper.Withdraw(ctx, msg.Creator, msg.MarketUID, msg.ParticipationIndex, msg.Mode, msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "process withdrawal")
	}

	return &types.MsgWithdrawResponse{
		ID:                 id,
		MarketUID:          msg.MarketUID,
		ParticipationIndex: msg.ParticipationIndex,
	}, nil
}
