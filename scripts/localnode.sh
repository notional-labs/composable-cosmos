#!/bin/bash

KEY="mykey"
KEY1="mykey1"
CHAINID="localpica"
MONIKER="localtestnet"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
BINARY=$1
# to trace evm
#TRACE="--trace"
TRACE=""

HOME_DIR=mytestnet
DENOM=ppica

KEY00="test0"
KEY11="test1"
KEY22="test2"


if [ "$CONTINUE" == "true" ]; then
    echo "\n ->> continuing from previous state"
    $BINARY start --home $HOME_DIR --log_level debug
    exit 0
fi

$BINARY config keyring-backend $KEYRING
$BINARY config chain-id $CHAINID

# remove existing daemon
rm -rf $HOME_DIR

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | $BINARY keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover --home $HOME_DIR
echo "bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort" | $BINARY keys add $KEY1 --keyring-backend $KEYRING --algo $KEYALGO --recover --home $HOME_DIR

$BINARY keys add $KEY00 --keyring-backend $KEYRING --home $HOME_DIR
$BINARY keys add $KEY11 --keyring-backend $KEYRING --home $HOME_DIR
$BINARY keys add $KEY22 --keyring-backend $KEYRING --home $HOME_DIR
$BINARY init $CHAINID --chain-id $CHAINID --default-denom "ppica" --home $HOME_DIR

update_test_genesis () {
    # update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
    cat $HOME_DIR/config/genesis.json | jq "$1" > $HOME_DIR/config/tmp_genesis.json && cp $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json
}

# Allocate genesis accounts (cosmos formatted addresses)
$BINARY add-genesis-account $KEY 100000000000000000000000000ppica --keyring-backend $KEYRING --home $HOME_DIR
$BINARY add-genesis-account $KEY1 100000000000000000000000000ppica --keyring-backend $KEYRING --home $HOME_DIR
$BINARY add-genesis-account $KEY00 "1000000000000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME_DIR
$BINARY add-genesis-account $KEY11 "1000000000000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME_DIR
$BINARY add-genesis-account $KEY22 "1000000000000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME_DIR


# Sign genesis transaction
$BINARY gentx $KEY 10030009994127689ppica --keyring-backend $KEYRING --chain-id $CHAINID --home $HOME_DIR

update_test_genesis '.app_state["gov"]["params"]["voting_period"]="5s"'
update_test_genesis '.app_state["mint"]["params"]["mint_denom"]="'$DENOM'"'
update_test_genesis '.app_state["gov"]["params"]["min_deposit"]=[{"denom":"'$DENOM'","amount": "1"}]'
update_test_genesis '.app_state["crisis"]["constant_fee"]={"denom":"'$DENOM'","amount":"1000"}'

# Collect genesis tx
$BINARY collect-gentxs --home $HOME_DIR

# Run this to ensure everything worked and that the genesis file is setup correctly
$BINARY validate-genesis --home $HOME_DIR

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# update request max size so that we can upload the light client
# '' -e is a must have params on mac, if use linux please delete before run
sed -i'' -e 's/max_body_bytes = /max_body_bytes = 1/g' $HOME_DIR/config/config.toml
sed -i'' -e 's/max_tx_bytes = 1048576/max_tx_bytes = 10000000/g' $HOME_DIR/config/config.toml
sed -i'' -e 's/timeout_commit = "5s"/timeout_commit = "1s"/' $HOME_DIR/config/config.toml


$BINARY start --rpc.unsafe --rpc.laddr tcp://0.0.0.0:26657 --pruning=nothing --minimum-gas-prices=0.000ppica --home=$HOME_DIR  --log_level trace --trace --with-tendermint=true --transport=socket  --grpc.enable=true --grpc-web.enable=false --api.enable=true  --p2p.pex=false --p2p.upnp=false