#!/bin/bash
KEY=mykey
KEY1=mykey1
KEYALGO="secp256k1"
KEYRING="test"
HOME_DIR="mytestnet"
BINARY=_build/old/picad
DENOM=ppica
CHAINID=centauri-dev

MYKEY1_ADDRESS=$($BINARY keys show $KEY1 -a --keyring-backend $KEYRING --home $HOME_DIR)
echo "Address of mykey1: $MYKEY1_ADDRESS"


$BINARY tx transmiddleware add-rly --from $KEY1 $MYKEY1_ADDRESS --keyring-backend test --home $HOME_DIR --chain-id $CHAINID --fees 100000${DENOM} -y
sleep 5