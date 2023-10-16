
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
		Use:   "ibc-solidity-call [host-zone-id] [contract-address]",
		Short: "Broadcast message ibc-solidity-call",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			
			// host-zone-id is the id of the network the solidity contract is stored.
		
			// contract address is the string of the address being called

			// SolidityMsg is the msg being called to the solidity contract

		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}