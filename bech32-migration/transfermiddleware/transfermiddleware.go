package transfermiddleware

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "cosmossdk.io/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/composable/v6/bech32-migration/utils"
	"github.com/notional-labs/composable/v6/x/transfermiddleware/types"
)

func MigrateAddressBech32(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	ctx.Logger().Info("Migration of address bech32 for transfermiddleware module begin")
	allowRelayAddressCount := uint64(0)

	store := ctx.KVStore(storeKey)

	relayAddressPrefix := []byte{1}
	iterator := sdk.KVStorePrefixIterator(store, types.KeyRlyAddress)

	for ; iterator.Valid(); iterator.Next() {
		allowRelayAddressCount++
		trimedAddr := strings.Replace(string(iterator.Key()), "\x04", "", 1)
		newPrefixAddr := utils.ConvertAccAddr(trimedAddr)
		store.Set(types.GetKeyByRlyAddress(newPrefixAddr), relayAddressPrefix)
	}

	ctx.Logger().Info(
		"Migration of address bech32 for transfermiddleware module done",
		"allow_relay_address_count", allowRelayAddressCount,
	)
}
