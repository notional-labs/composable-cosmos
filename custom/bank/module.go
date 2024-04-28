package bank

import (
	"cosmossdk.io/core/address"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankmodule "github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	custombankkeeper "github.com/notional-labs/composable/v6/custom/bank/keeper"
)

type AppModule struct {
	bankmodule.AppModule
	keeper       custombankkeeper.Keeper
	subspace     exported.Subspace
	addressCodec address.Codec
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper custombankkeeper.Keeper, accountKeeper types.AccountKeeper, ss exported.Subspace) AppModule {
	bankModule := bankmodule.NewAppModule(cdc, keeper, accountKeeper, ss)
	return AppModule{
		AppModule:    bankModule,
		keeper:       keeper,
		subspace:     ss,
		addressCodec: accountKeeper.AddressCodec(),
	}
}

// RegisterServices registers module services.
// NOTE: Overriding this method as not doing so will cause a panic
// when trying to force this custom keeper into a bankkeeper.BaseKeeper
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), custombankkeeper.NewMsgServerImpl(am.keeper, am.addressCodec))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := bankkeeper.NewMigrator(am.keeper.BaseKeeper, am.subspace)
	if err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1to2); err != nil {
		panic(fmt.Sprintf("failed to migrate x/bank from version 1 to 2: %v", err))
	}

	if err := cfg.RegisterMigration(types.ModuleName, 2, m.Migrate2to3); err != nil {
		panic(fmt.Sprintf("failed to migrate x/bank from version 2 to 3: %v", err))
	}

	if err := cfg.RegisterMigration(types.ModuleName, 3, m.Migrate3to4); err != nil {
		panic(fmt.Sprintf("failed to migrate x/bank from version 3 to 4: %v", err))
	}
}
