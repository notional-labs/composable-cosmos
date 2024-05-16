package types

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	types "github.com/cosmos/cosmos-sdk/types"
)

// ParamsKey is the key to use for the keeper store.
var (
	ParamsKey = []byte{0x01} // key for global staking middleware params in the keeper store

	SequenceFeeKey = []byte{0x21} // prefix for sequence fee
)

const (
	// module name
	ModuleName = "ibctransferparamsmodule"

	// StoreKey is the default store key for ibctransfermiddleware module that store params when apply validator set changes and when allow to unbond/redelegate

	StoreKey = "customibcparams" // not using the module name because of collisions with key "staking"

	RouterKey = ModuleName
)

// GetSequenceKey returns a key prefix for indexing HistoricalInfo objects.
func GetSequenceKey(sequence uint64) []byte {
	return append(SequenceFeeKey, []byte(strconv.FormatUint(sequence, 10))...)
}

func MustMarshalCoin(cdc codec.BinaryCodec, coin *types.Coin) []byte {
	return cdc.MustMarshal(coin)
}

// unmarshal a redelegation from a store value
func MustUnmarshalCoin(cdc codec.BinaryCodec, value []byte) types.Coin {
	validator, err := UnmarshalCoin(cdc, value)
	if err != nil {
		panic(err)
	}

	return validator
}

// unmarshal a redelegation from a store value
func UnmarshalCoin(cdc codec.BinaryCodec, value []byte) (v types.Coin, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}
