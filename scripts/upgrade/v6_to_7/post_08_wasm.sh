#!/bin/bash
KEY="mykey"
KEYALGO="secp256k1"
KEYRING="test"
HOME_DIR="mytestnet"


sleep 2

./_build/new/centaurid query ibc-wasm checksums --home $HOME_DIR