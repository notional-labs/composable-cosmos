#!/bin/bash

# the upgrade is a fork, "true" otherwise
FORK=${FORK:-"false"}


BINARY=_build/old/centaurid
HOME=mytestnet
ROOT=$(pwd)
DENOM=ppica
CHAIN_ID=centaurid

ADDITIONAL_PRE_SCRIPTS="./scripts/upgrade/old-node-scripts.sh"

SLEEP_TIME=1


echo "running the node"
screen -L -dmS node1 bash scripts/localnode.sh $BINARY $DENOM --Logfile $HOME/log-screen.txt
#scripts/localnode.sh $BINARY


echo "Waiting for the node to start..."
sleep 10 # wait for note to start

# execute additional pre scripts
source $ADDITIONAL_PRE_SCRIPTS 