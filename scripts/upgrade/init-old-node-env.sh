#!/bin/bash

# the upgrade is a fork, "true" otherwise
FORK=${FORK:-"false"}


BINARY=_build/old/centaurid
HOME=mytestnet
ROOT=$(pwd)
DENOM=ppica
CHAIN_ID=centaurid

ADDITIONAL_PRE_SCRIPTS="./scripts/upgrade/v_6_4_8/setup-08-wasm.sh"

SLEEP_TIME=1
# run old node
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "running old node"
    screen -L -dmS node1 bash scripts/localnode.sh $BINARY $DENOM --Logfile $HOME/log-screen.txt
else
    screen -L -Logfile $HOME/log-screen.txt -dmS node1 bash scripts/localnode.sh $BINARY $DENOM
fi

# execute additional pre scripts
if [ ! -z "$ADDITIONAL_PRE_SCRIPTS" ]; then
    # slice ADDITIONAL_SCRIPTS by ,
    SCRIPTS=($(echo "$ADDITIONAL_PRE_SCRIPTS" | tr ',' ' '))
    for SCRIPT in "${SCRIPTS[@]}"; do
         # check if SCRIPT is a file
        if [ -f "$SCRIPT" ]; then
            echo "executing additional pre scripts from $SCRIPT"
            source $SCRIPT $BINARY
            echo "CONTRACT_ADDRESS = $CONTRACT_ADDRESS"
            sleep 5
        else
            echo "$SCRIPT is not a file"
        fi
    done
fi
