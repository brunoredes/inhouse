#!/bin/bash

# Get all container IDs (including stopped ones)
container_ids=($(docker ps -a -q))

# Check if there are any containers
if [ ${#container_ids[@]} -eq 0 ]; then
    echo "No containers found."
    exit 1
fi

# Iterate over each container ID and start it
for container_id in "${container_ids[@]}"; do
    echo "Starting container: $container_id"
    docker start "$container_id"
done

echo "All containers started."
