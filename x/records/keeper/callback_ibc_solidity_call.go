package keeper

import (
	"github.com/evmos/evmos/v14/utils"
	icacallbackstypes "github.com/evmos/evmos/v14/x/icacallbacks/types"
	"github.com/evmos/evmos/v14/x/records/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

// Callback after an IBC Solidity Call is IBC tranferred to the host zone
//
//	If successful: mark the IBC Solidity Call status as DE_QUEUE
//	If failure: mark the IBC Solidity Call status as FAILED
//	If timeout: revert the IBC Solidity Call status back to TRANSFER_QUEUE so it gets resubmitted
func (k Keeper) IBCSolidityCallCallback(ctx sdk.Context, packet channeltypes.Packet, ackResponse *icacallbackstypes.AcknowledgementResponse, args []byte) error {

	
}