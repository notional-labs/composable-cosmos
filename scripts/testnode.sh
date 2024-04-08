#!/bin/bash

KEY="mykey"
CHAINID="test-1"
MONIKER="localtestnet"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

HOME_DIR=~/.banksy
DENOM=${2:-ppica}


# remove existing daemon
rm -rf ~/.banksy*

centaurid config keyring-backend $KEYRING
centaurid config chain-id $CHAINID

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | centaurid keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

centaurid init $MONIKER --chain-id $CHAINID

update_test_genesis () {
    # update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
    cat $HOME_DIR/config/genesis.json | jq "$1" > $HOME_DIR/config/tmp_genesis.json && mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json
}

# Allocate genesis accounts (cosmos formatted addresses)
centaurid add-genesis-account $KEY 100000000000000000000000000ppica --keyring-backend $KEYRING

# Sign genesis transaction
centaurid gentx $KEY 10030009994127689ppica --keyring-backend $KEYRING --chain-id $CHAINID

update_test_genesis '.app_state["gov"]["params"]["voting_period"]="20s"'
update_test_genesis '.app_state["mint"]["params"]["mint_denom"]="'$DENOM'"'
update_test_genesis '.app_state["gov"]["params"]["min_deposit"]=[{"denom":"'$DENOM'","amount": "1"}]'
update_test_genesis '.app_state["crisis"]["constant_fee"]={"denom":"'$DENOM'","amount":"1000"}'
update_test_genesis '.app_state["staking"]["params"]["bond_denom"]="'$DENOM'"'

# Collect genesis tx
centaurid collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
centaurid validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# update request max size so that we can upload the light client
# '' -e is a must have params on mac, if use linux please delete before run
sed -i'' -e 's/max_body_bytes = /max_body_bytes = 1/g' ~/.banksy/config/config.toml

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
# centaurid start --pruning=nothing  --minimum-gas-prices=0.0001ppica --rpc.laddr tcp://0.0.0.0:26657
