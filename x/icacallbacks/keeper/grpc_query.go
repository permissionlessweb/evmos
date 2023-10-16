package keeper

import (
	"github.com/evmos/evmos/v14/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
