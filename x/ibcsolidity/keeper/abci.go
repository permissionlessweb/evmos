package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/evmos/evmos/v14/x/ibcsolidity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker of ibcsolidity module
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// Iterate over all host zone and verify ...?


	k.AssertEvmosAndDayEpochRelationship(ctx)
}

func (k Keeper) EndBlocker(ctx sdk.Context) {
	// Submit an IBC Solidity Call or ...? for all queued IBC Solidity Calls across each host


	k.SendAllIBCSolidityCalls(ctx)
}
