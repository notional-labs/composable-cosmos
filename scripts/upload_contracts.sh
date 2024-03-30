#!/bin/bash

KEY="mykey"
CHAINID="test-1"
KEYALGO="secp256k1"
KEYRING="test"

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon
rm -rf ~/.centauri*

~/go/bin/centaurid config keyring-backend $KEYRING
~/go/bin/centaurid config chain-id $CHAINID



~/go/bin/centaurid  tx 08-wasm push-wasm contracts/simple.wasm --from mykey --keyring-backend test --home $HOME/.banksy --gas 10002152622 --fees 10020166stake -y