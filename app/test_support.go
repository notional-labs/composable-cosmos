package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

func (app *ComposableApp) GetStakingKeeper() *stakingkeeper.Keeper {
	return &app.StakingKeeper.Keeper
}

func (app *ComposableApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

func (app *ComposableApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

func (app *ComposableApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *ComposableApp) GetBankKeeper() bankkeeper.Keeper {
	return app.BankKeeper
}

func (app *ComposableApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.AccountKeeper
}

func (app *ComposableApp) GetWasmKeeper() wasmkeeper.Keeper {
	return app.WasmKeeper
}

// GetTransferKeeper implements the TestingApp interface.
func (app *ComposableApp) GetTransferKeeper() *ibctransferkeeper.Keeper {
	return &app.TransferKeeper.Keeper
}

// GetTxConfig implements the TestingApp interface.
func (app *ComposableApp) GetTxConfig() client.TxConfig {
	cfg := MakeEncodingConfig()
	return cfg.TxConfig
}
