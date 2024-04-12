!/bin/bash
KEY="mykey"
KEYALGO="secp256k1"
CHAINID="centaurid"
KEYRING="test"
HOME_DIR="mytestnet"
BINARY=_build/old/centaurid
DENOM=ppica
 
$BINARY tx gov submit-proposal scripts/08-wasm/ics10_grandpa_cw.wasm.json --from=mykey --fees 100000${DENOM} --gas auto --keyring-backend test  --home $HOME_DIR -y 

sleep 2
# TODO: fetch the propsoal id dynamically 
$BINARY tx gov deposit "1" "20000000ppica" --from $KEY --fees 100000${DENOM} --keyring-backend test --home $HOME_DIR -y 

sleep 2
$BINARY tx gov vote 1 yes --from $KEY --fees 100000${DENOM} --keyring-backend test --home $HOME_DIR -y 
