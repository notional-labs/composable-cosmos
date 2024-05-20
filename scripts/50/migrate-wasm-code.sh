#!/bin/bash
KEY=mykey
DENOM="ppica"
HOME_DIR="mytestnet"
CHAINID="centauri-dev"
BINARY=picad
WASM_CONTRACT_PATH="/Users/kien6034/notional/projects/composable-ibc/target/wasm32-unknown-unknown/release/ics10_grandpa_cw.wasm"


WASM_CLIENT_ID="08-wasm-0"
WASM_CHECKSUM="3e743bf804a60e5fd1dfab6c61bba0f2e76cda260edc66d6b7b10691fb5096c1"

# Fetch proposal id 
$BINARY tx ibc-wasm migrate-contract $WASM_CLIENT_ID $WASM_CHECKSUM '{}' --from $KEY --keyring-backend test --chain-id $CHAINID --home $HOME_DIR --deposit 10000000000${DENOM} --gas 20002152622 --fees 20020166${DENOM}  -y

# Fetch proposal id 
sleep 6
# $BINARY query gov proposals -o json > /tmp/proposals.json
# PROPOSAL_ID=$(jq -r '.proposals[-1].id' /tmp/proposals.json)
PROPOSAL_ID= 4  ## fix this
echo "Proposal ID is: $PROPOSAL_ID"

# Validator vote yes
$BINARY tx gov vote $PROPOSAL_ID yes --from $KEY --fees 100000${DENOM} --keyring-backend test --home $HOME_DIR --chain-id $CHAINID -y 

#Voting time is 20s, check in localnode.sh
sleep 20

# Check the status 
$BINARY query gov proposal $PROPOSAL_ID -o json > /tmp/proposal.json
STATUS=$(jq -r '.proposal.status' /tmp/proposal.json)
echo "Proposal status is: $STATUS"


