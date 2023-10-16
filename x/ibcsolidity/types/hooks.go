package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// combine multiple staking hooks, all hook functions are run in array sequence
type MultiIBCSolidityHooks []IBCSolidityHooks

func NewMultiIBCSolidityHooks(hooks ...IBCSolidityHooks) MultiIBCSolidityHooks {
	return hooks
}

func (h MultiIBCSolidityHooks) AfterIBCSolidity(ctx sdk.Context, addr sdk.AccAddress) {
	for i := range h {
		h[i].AfterIBCSolidity(ctx, addr)
	}
}
