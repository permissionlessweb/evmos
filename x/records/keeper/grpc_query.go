package keeper

import (
	"github.com/evmos/evmos/v14/x/records/types"
)

var _ types.QueryServer = Keeper{}
