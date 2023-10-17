
import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"

	"github.com/evmos/evmos/v14/x/ibcsolidity/types"
)

func CmdIBCSolidityCall() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ibc-solidity-call [host-zone-id] [contract-address] [solidity-msg]",
		Short: "Broadcast message ibc-solidity-call",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIBCSolidityCall(
				clientCtx.GetFromAddress().String(),
				argChainId,
				solidityContractAddress,
				solidityMsg,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}