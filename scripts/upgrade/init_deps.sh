#!/bin/bash


OLD_VERSION=kien-devnet-651
SOFTWARE_UPGRADE_NAME="v6_6_0"
ROOT=$(pwd)

COMPOSABLE_VERSION="pfm-fix"

if [ ! -f "_build/composable/$COMPOSABLE_VERSION.zip" ] &> /dev/null
then
    mkdir -p _build/composable
    wget -c "https://github.com/notional-labs/composable/archive/refs/tags/${COMPOSABLE_VERSION}.zip" -O _build/composable/${COMPOSABLE_VERSION}.zip
    unzip _build/composable/${COMPOSABLE_VERSION}.zip -d _build/composable
fi



# install old binary if not exist
if [ ! -f "_build/$OLD_VERSION.zip" ] &> /dev/null
then
    mkdir -p _build/old
    wget -c "https://github.com/notional-labs/composable-cosmos/archive/refs/tags/${OLD_VERSION}.zip" -O _build/${OLD_VERSION}.zip
    unzip _build/${OLD_VERSION}.zip -d _build
fi

# reinstall old binary
if [ $# -eq 1 ] && [ $1 == "--reinstall-old" ] || ! command -v _build/old/centaurid &> /dev/null; then
    cd ./_build/composable-cosmos-${OLD_VERSION}
    GOBIN="$ROOT/_build/old" go install -mod=readonly ./...
    cd ../..
fi


# install new binary
if ! command -v _build/new/picad &> /dev/null
then
    mkdir -p _build/new
    GOBIN="$ROOT/_build/new" go install -mod=readonly ./...
fi

