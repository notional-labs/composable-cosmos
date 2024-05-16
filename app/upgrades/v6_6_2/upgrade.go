package v6_6_2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibchookstypes "github.com/notional-labs/composable/v6/x/ibc-hooks/types"
	ibctransfermiddlewaretypes "github.com/notional-labs/composable/v6/x/ibctransfermiddleware/types"

	"github.com/notional-labs/composable/v6/app/keepers"
	"github.com/notional-labs/composable/v6/app/upgrades"
	bech32IbcHooksMigration "github.com/notional-labs/composable/v6/bech32-migration/ibchooks"
	bench32ibctransfermiddleware "github.com/notional-labs/composable/v6/bech32-migration/ibctransfermiddleware"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	codec codec.Codec,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		keys := keepers.GetKVStoreKey()
		// Migration prefix
		ctx.Logger().Info("First step: Migrate addresses stored in bech32 form to use new prefix")
		bench32ibctransfermiddleware.MigrateAddressBech32(ctx, keys[ibctransfermiddlewaretypes.StoreKey], codec)
		bech32IbcHooksMigration.MigrateAddressBech32(ctx, keys[ibchookstypes.StoreKey], codec)
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
