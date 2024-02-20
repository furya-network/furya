package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/furya-network/furya/x/house/types"
	"github.com/spf13/cobra"
)

func CmdDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [market_uid] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Deposit tokens to be part of a house",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Deposit coins to be part of a house corresponding to a market.

				Example:
				$ %s tx house deposit bc79a72c-ad7e-4cf5-91a2-98af2751e812 1000ufury --from mykey
				`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			MarketUID := args[0]

			argAmountCosmosInt, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return types.ErrInvalidAmount
			}

			depAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgDeposit(depAddr.String(), MarketUID, argAmountCosmosInt)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
