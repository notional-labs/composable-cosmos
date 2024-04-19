package transfermiddleware_test

import (
	"encoding/json"
	"os"
	"testing"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	wasmkeeper "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/keeper"
	wasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	"github.com/stretchr/testify/suite"

	customibctesting "github.com/notional-labs/composable/v6/app/ibctesting"
)

// NOTE: This is the address of the gov authority on the chain that is being tested.
// This means that we need to check bech32 .... everywhere.
var govAuthorityAddress = "centauri10556m38z4x6pqalr9rl5ytf3cff8q46nk85k9m"

// ORIGINAL NOTES:
// convert from: centauri10556m38z4x6pqalr9rl5ytf3cff8q46nk85k9m

type TransferTestSuite struct {
	suite.Suite

	coordinator *customibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *customibctesting.TestChain
	chainB *customibctesting.TestChain

	ctx      sdk.Context
	store    storetypes.KVStore
	testData map[string]string

	wasmKeeper wasmkeeper.Keeper
}

func (suite *TransferTestSuite) SetupTest() {
	suite.coordinator = customibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(customibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(customibctesting.GetChainID(1))
	suite.chainB.SetWasm(true)
	suite.coordinator.CommitNBlocks(suite.chainA, 2)
	suite.coordinator.CommitNBlocks(suite.chainB, 2)

	data, err := os.ReadFile("../../app/ibctesting/test_data/raw.json")
	suite.Require().NoError(err)
	err = json.Unmarshal(data, &suite.testData)
	suite.Require().NoError(err)

	suite.ctx = suite.chainB.GetContext().WithBlockGasMeter(storetypes.NewInfiniteGasMeter())
	suite.store = suite.chainB.App.GetIBCKeeper().ClientKeeper.ClientStore(suite.ctx, "08-wasm-0")

	wasmContract, err := os.ReadFile("../../contracts/ics10_grandpa_cw.wasm")
	suite.Require().NoError(err)

	suite.wasmKeeper = suite.chainB.GetTestSupport().Wasm08Keeper()

	msg := wasmtypes.NewMsgStoreCode(govAuthorityAddress, wasmContract)

	response, err := suite.wasmKeeper.StoreCode(suite.ctx, msg)

	suite.Require().NoError(err)
	suite.Require().NotNil(response.Checksum)
	suite.coordinator.CodeID = response.Checksum
}

func TestTransferTestSuite(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}
