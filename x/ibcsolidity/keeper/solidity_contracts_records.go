package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/spf13/cast"

	"github.com/evmos/evmos/v14/utils"
	recordstypes "github.com/evmos/evmos/v14/x/records/types"
	"github.com/evmos/evmos/v14/x/ibcsolidity/types"
)

func (k Keeper) CreateSolidityContractRecordsForEpoch(ctx sdk.Context, epochNumber uint64) {

}

// Iterate each transfer record marked TRANSFER_QUEUE and IBC call from the Evmos controller account to the ibc-solidity-call ICAs on each host zone
func (k Keeper) TransferExistingIBCSolidityCallsToHostZones(ctx sdk.Context, epochNumber uint64 callRecords []recordstypes.DepositRecord) {

}


// Iterate each call record marked CALL_QUEUE and IBC call from the Evmos controller account to the ibc-solidity-call ICAs on each host zone
func (k Keeper) CallExistingIBCSolidityCallsOnHostZones(ctx sdk.Context, epochNumber uint64 callRecords []recordstypes.DepositRecord) {

}
