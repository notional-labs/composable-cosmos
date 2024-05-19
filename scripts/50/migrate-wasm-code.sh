#!/bin/bash
KEY=mykey
DENOM="ppica"
HOME_DIR="mytestnet"
CHAINID="centauri-dev"
BINARY=picad
WASM_CONTRACT_PATH="/Users/kien6034/notional/projects/composable-ibc/target/wasm32-unknown-unknown/release/ics10_grandpa_cw.wasm"

# Create a proposale to store wasm code
$BINARY tx ibc-wasm store-code $WASM_CONTRACT_PATH --title "migrate new contract" --summary "none" --from $KEY --keyring-backend test --home $HOME_DIR --deposit 10000000000${DENOM} --gas 20002152622 --fees 20020166${DENOM}  -y 

# Fetch proposal id 
sleep 6
$BINARY query gov proposals -o json > tmp-proposals.json
PROPOSAL_ID=$(jq -r '.proposals[-1].id' tmp-proposals.json)
echo "Proposal ID is: $PROPOSAL_ID"
rm -rf tmp-proposals.json

# Validator vote yes
$BINARY tx gov vote $PROPOSAL_ID yes --from $KEY --fees 100000${DENOM} --keyring-backend test --home $HOME_DIR --chain-id $CHAINID -y 

#Voting time is 20s, check in localnode.sh
sleep 20

# Check the status 
$BINARY query gov proposal $PROPOSAL_ID -o json > tmp-proposal.json
STATUS=$(jq -r '.proposal.status' tmp-proposal.json)
echo "Proposal status is: $STATUS"

# Query newly wasm checksums 
CHECKSUM=$($BINARY query ibc-wasm checksums -o json | jq -r '.checksums[-1]')



##### Migrate the contract 
