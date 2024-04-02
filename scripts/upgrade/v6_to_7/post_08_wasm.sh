#!/bin/bash
KEY="mykey"
KEYALGO="secp256k1"
KEYRING="test"
HOME_DIR="mytestnet"


sleep 2

checksum=$(./_build/new/centaurid query ibc-wasm checksums --home $HOME_DIR -o json | jq -r '.checksums[0]')

if ./_build/new/centaurid query ibc-wasm code $checksum --home $HOME_DIR -o json &> /dev/null; then
    echo "Code with checksum $checksum exists."
else
    echo "Code with checksum $checksum does not exist."
fi


