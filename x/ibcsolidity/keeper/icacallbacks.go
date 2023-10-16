package keeper

import (
	icacallbackstypes "github.com/evmos/evmos/v14/x/icacallbacks/types"
)

const (
	ICACallbackID_SolidityCall       = "soliditycall"
)

func (k Keeper) Callbacks() icacallbackstypes.ModuleCallbacks {
	return []icacallbackstypes.ICACallback{
		{CallbackId: ICACallbackID_SolidityCall, CallbackFunc: icacallbackstypes.ICACallbackFunction(k.SolidityCallCallback)},
	}
}
