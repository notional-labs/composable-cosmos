ROOT=$(pwd)

cd $ROOT/_build/composable

# This start the node
nix run .#zombienet-rococo-local-picasso-dev


# Sleep 20secs for the node to run, before setting up the relayer

sleep 20



ROOT=$(pwd)

cd $ROOT/_build/composable

# This start the node
nix run .#zombienet-rococo-local-picasso-dev
