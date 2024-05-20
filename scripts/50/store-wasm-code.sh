#!/bin/bash
KEY=mykey
DENOM="ppica"
HOME_DIR="mytestnet"
CHAINID="centauri-dev"
BINARY=picad
WASM_CONTRACT_PATH="$PWD/bin/ics10_grandpa_cw.wasm"

# Wait for chain to start 
echo "Waiting for chain to start..."

$BINARY tx ibc-wasm store-code $WASM_CONTRACT_PATH --title "store new wasm code"  --summary "none" --from $KEY --keyring-backend test --chain-id $CHAINID --home $HOME_DIR --deposit 10000000000${DENOM} --gas 20002152622 --fees 20020166${DENOM}  -y

# Fetch proposal id 
sleep 6
# $BINARY query gov proposals -o json > /tmp/proposals.json
# PROPOSAL_ID=$(jq -r '.proposals[-1].id' /tmp/proposals.json)
PROPOSAL_ID=3  ## fix this
echo "Proposal ID is: $PROPOSAL_ID"

# Validator vote yes
$BINARY tx gov vote $PROPOSAL_ID yes --from $KEY --fees 100000${DENOM} --keyring-backend test --home $HOME_DIR --chain-id $CHAINID -y 

#Voting time is 20s, check in localnode.sh
sleep 20

# Check the status 
$BINARY query gov proposal $PROPOSAL_ID -o json > /tmp/proposal.json
STATUS=$(jq -r '.proposal.status' /tmp/proposal.json)
echo "Proposal status is: $STATUS"

# Query wasm checksums 
$BINARY query ibc-wasm checksums
