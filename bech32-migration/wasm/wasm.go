package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/notional-labs/composable/v6/bech32-migration/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateAddressBech32(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	ctx.Logger().Debug("Migration of address bech32 for wasm module begin")

	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), types.CodeKeyPrefix)
	iter := prefixStore.Iterator(nil, nil)
	defer iter.Close()

	totalMigratedCodeId := uint64(0)
	for ; iter.Valid(); iter.Next() {
		// get code info value
		var c types.CodeInfo
		cdc.MustUnmarshal(iter.Value(), &c)
		c.Creator = utils.ConvertAccAddr(c.Creator)

		// save updated code info
		prefixStore.Set(iter.Key(), cdc.MustMarshal(&c))

		totalMigratedCodeId++
	}

	ctx.Logger().Debug(
		"Migration of address bech32 for wasm module done",
		"total_migrated_code_id", totalMigratedCodeId,
	)
}
