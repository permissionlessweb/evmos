package keeper

import (
	icacallbackstypes "github.com/evmos/evmos/v14/x/icacallbacks/types"
)

const IBCCallbacksID_IBCSolidityCall = "ibc-solidity-call"

func (k Keeper) Callbacks() icacallbackstypes.ModuleCallbacks {
	return []icacallbackstypes.ICACallback{
		{CallbackId: IBCCallbacksID_IBCSolidityCall, CallbackFunc: icacallbackstypes.ICACallbackFunction(k.IBCSolidityCallCallback)},
	}
}
