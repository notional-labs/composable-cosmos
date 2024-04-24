package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"

	storetypes "cosmossdk.io/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibctransfermiddleware "github.com/notional-labs/composable/v6/x/ibctransfermiddleware/keeper"
)

type Keeper struct {
	ibctransferkeeper.Keeper
	cdc                   codec.BinaryCodec
	IbcTransfermiddleware *ibctransfermiddleware.Keeper
	bank                  *custombankkeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	ics4Wrapper porttypes.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper,
	bk types.BankKeeper,
	scopedKeeper exported.ScopedKeeper,
	ibcTransfermiddleware *ibctransfermiddleware.Keeper,
	authority string,
) Keeper {
	keeper := Keeper{
		Keeper:                ibctransferkeeper.NewKeeper(cdc, key, paramSpace, ics4Wrapper, channelKeeper, portKeeper, authKeeper, bk, scopedKeeper, authority),
		IbcTransfermiddleware: ibcTransfermiddleware,
		cdc:                   cdc,
		bank:                  bankKeeper,
	}
	return keeper
}

// Transfer is the server API around the Transfer method of the IBC transfer module.
// It checks if the sender is allowed to transfer the token and if the channel has fees.
// If the channel has fees, it will charge the sender and send the fees to the fee address.
// If the sender is not allowed to transfer the token because this tokens does not exists in the allowed tokens list, it just return without doing anything.
// If the sender is allowed to transfer the token, it will call the original transfer method.
// If the transfer amount is less than the minimum fee, it will charge the full transfer amount.
// If the transfer amount is greater than the minimum fee, it will charge the minimum fee and the percentage fee.
func (k Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.IbcTransfermiddleware.GetParams(ctx)
	if params.ChannelFees != nil && len(params.ChannelFees) > 0 {
		channelFee := findChannelParams(params.ChannelFees, msg.SourceChannel)
		if channelFee != nil {
			if channelFee.MinTimeoutTimestamp > 0 {

				goCtx := sdk.UnwrapSDKContext(goCtx)
				blockTime := goCtx.BlockTime()

				timeoutTimeInFuture := time.Unix(0, int64(msg.TimeoutTimestamp))
				if timeoutTimeInFuture.Before(blockTime) {
					return nil, fmt.Errorf("incorrect timeout timestamp found during ibc transfer. timeout timestamp is in the past")
				}

				difference := timeoutTimeInFuture.Sub(blockTime).Nanoseconds()
				if difference < channelFee.MinTimeoutTimestamp {
					return nil, fmt.Errorf("incorrect timeout timestamp found during ibc transfer. too soon")
				}
			}
			coin := findCoinByDenom(channelFee.AllowedTokens, msg.Token.Denom)
			if coin == nil {
				return nil, fmt.Errorf("token not allowed to be transferred in this channel")
			}
			minFee := coin.MinFee.Amount
			charge := minFee
			if charge.GT(msg.Token.Amount) {
				charge = msg.Token.Amount
			}

			newAmount := msg.Token.Amount.Sub(charge)

			if newAmount.IsPositive() {
				percentageCharge := newAmount.QuoRaw(coin.Percentage)
				newAmount = newAmount.Sub(percentageCharge)
				charge = charge.Add(percentageCharge)
			}

			msgSender, err := sdk.AccAddressFromBech32(msg.Sender)
			if err != nil {
				return nil, err
			}

			feeAddress, err := sdk.AccAddressFromBech32(channelFee.FeeAddress)
			if err != nil {
				return nil, err
			}

			send_err := k.bank.SendCoins(ctx, msgSender, feeAddress, sdk.NewCoins(sdk.NewCoin(msg.Token.Denom, charge)))
			if send_err != nil {
				return nil, send_err
			}

			if newAmount.LTE(sdk.ZeroInt()) {
				return &types.MsgTransferResponse{}, nil
			}
			msg.Token.Amount = newAmount
		}
	}
	return k.Keeper.Transfer(goCtx, msg)
}
