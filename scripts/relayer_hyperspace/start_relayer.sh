./scripts/relayer_hyperspace/run-composable.sh
./scripts/relayer_hyperspace/run-picasso.sh

# Waiting for nodes to start 
sleep 20 

RUST_LOG=hyperspace_cosmos=trace,hyperspace=trace,info /Users/kien6034/notional/projects/composable-ibc/target/release/hyperspace create-clients --config-a scripts/relayer_hyperspace/config-chain-a.toml --config-b scripts/relayer_hyperspace/config-chain-b.toml --config-core scripts/relayer_hyperspace/config-core.toml --delay-period 50
