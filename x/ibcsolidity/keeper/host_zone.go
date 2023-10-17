package keeper

import (
	"errors"
	"fmt"
	"math"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/evmos/evmos/v14/utils"
	"github.com/evmos/evmos/v14/x/ibcsolidity/types"
)

// SetHostZone set a specific hostZone in the store
func (k Keeper) SetHostZone(ctx sdk.Context, hostZone types.HostZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))
	b := k.cdc.MustMarshal(&hostZone)
	store.Set([]byte(hostZone.ChainId), b)
}

// GetHostZone returns a hostZone from its id
func (k Keeper) GetHostZone(ctx sdk.Context, chainId string) (val types.HostZone, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))
	b := store.Get([]byte(chainId))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetHostZoneFromHostDenom returns a HostZone from a HostDenom
func (k Keeper) GetHostZoneFromHostDenom(ctx sdk.Context, denom string) (*types.HostZone, error) {
	var matchZone types.HostZone
	k.IterateHostZones(ctx, func(ctx sdk.Context, index int64, zoneInfo types.HostZone) error {
		if zoneInfo.HostDenom == denom {
			matchZone = zoneInfo
			return nil
		}
		return nil
	})
	if matchZone.ChainId != "" {
		return &matchZone, nil
	}
	return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "No HostZone for %s found", denom)
}

// GetHostZoneFromTransferChannelID returns a HostZone from a transfer channel ID
func (k Keeper) GetHostZoneFromTransferChannelID(ctx sdk.Context, channelID string) (hostZone types.HostZone, found bool) {
	for _, hostZone := range k.GetAllActiveHostZone(ctx) {
		if hostZone.TransferChannelId == channelID {
			return hostZone, true
		}
	}
	return types.HostZone{}, false
}


// RemoveHostZone removes a hostZone from the store
func (k Keeper) RemoveHostZone(ctx sdk.Context, chain_id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))
	store.Delete([]byte(chain_id))
}

// GetAllHostZone returns all hostZone
func (k Keeper) GetAllHostZone(ctx sdk.Context) (list []types.HostZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostZone
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllActiveHostZone returns all hostZones that are active (halted = false)
func (k Keeper) GetAllActiveHostZone(ctx sdk.Context) (list []types.HostZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.HostZone
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if !val.Halted {
			list = append(list, val)
		}
	}

	return
}

// IterateHostZones iterates zones
func (k Keeper) IterateHostZones(ctx sdk.Context, fn func(ctx sdk.Context, index int64, zoneInfo types.HostZone) error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		k.Logger(ctx).Debug(fmt.Sprintf("Iterating HostZone %d", i))
		zone := types.HostZone{}
		k.cdc.MustUnmarshal(iterator.Value(), &zone)

		error := fn(ctx, i, zone)

		if error != nil {
			break
		}
		i++
	}
}