package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	epochtypes "github.com/Evmos-Labs/evmos/v16/x/epochs/types"
)

// RegisterInvariants registers all governance invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
}

// AllInvariants runs all invariants of the ibcsolidity module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// msg, broke := RedemptionRateInvariant(k)(ctx)
		// note: once we have >1 invariant here, follow the pattern from staking module invariants here: https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/x/staking/keeper/invariants.go
		return "", false
	}
}


// TODO: Consider removing evmos and day epochs completely and using a single hourly epoch
// Confirm the number of evmos epochs in 1 day epoch
func (k Keeper) AssertEvmosAndDayEpochRelationship(ctx sdk.Context) {
	evmosEpoch, found := k.GetEpochTracker(ctx, epochtypes.STRIDE_EPOCH)
	if !found || evmosEpoch.Duration == 0 {
		return
	}
	dayEpoch, found := k.GetEpochTracker(ctx, epochtypes.DAY_EPOCH)
	if !found || dayEpoch.Duration == 0 {
		return
	}
	if dayEpoch.Duration/evmosEpoch.Duration != EvmosEpochsPerDayEpoch {
		panic("The evmos epoch must be 1/4th the length of the day epoch")
	}
}
