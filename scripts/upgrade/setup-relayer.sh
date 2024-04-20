#!/bin/bash

ROOT=$(pwd)

cd $ROOT/_build/composable/

# Function to run a command with retries
run_with_retries() {
    local max_attempts=20
    local attempt=1
    local command=$@

    while [ $attempt -le $max_attempts ]; do
        echo "Attempt $attempt of $max_attempts: $command"
        $command
        
        # Check if the command was successful
        if [ $? -eq 0 ]; then
            echo "Command executed successfully."
            return 0
        else
            echo "Command failed, retrying..."
            ((attempt++))
        fi
        
        # Optional: sleep for a few seconds before retrying
        sleep 2
    done
    
    echo "All attempts failed for: $command"
    return 1
}

# Initialize clients
run_with_retries nix run .#picasso-centauri-ibc-init
sleep 1 

# Initialize connection
run_with_retries nix run .#picasso-centauri-ibc-connection-init
sleep 1

# Initialize channel
run_with_retries nix run .#picasso-centauri-ibc-channels-init
sleep 1
 
# Run relayer
run_with_retries nix run .#picasso-centauri-ibc-relay
sleep 1
