#!/bin/bash

CONTAINER="app1"

echo "An error occurred in container '$CONTAINER'..."
echo ""

NETWORK=$(docker inspect -f '{{range $k, $v := .NetworkSettings.Networks}}{{$k}}{{end}}' "$CONTAINER")
CURRENT_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "$CONTAINER")

if [ -z "$NETWORK" ] || [ -z "$CURRENT_IP" ]; then
  echo "Error getting IP from container '$CONTAINER'"
  exit 1
fi

IFS='.' read -r IP1 IP2 IP3 IP4 <<< "$CURRENT_IP"
NEW_IP4=$((IP4 + 3))

if [ "$NEW_IP4" -gt 254 ]; then
  NEW_IP4=100
fi

NEW_IP="$IP1.$IP2.$IP3.$NEW_IP4"

echo "Removing container '$CONTAINER' from network '$NETWORK'..."
echo ""
docker network disconnect "$NETWORK" "$CONTAINER"

echo "Connecting container '$CONTAINER' to network '$NETWORK'..."
echo ""
docker network connect --ip "$NEW_IP" "$NETWORK" "$CONTAINER"

ACTUAL_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "$CONTAINER")
echo "Container '$CONTAINER' successfully connected to network '$NETWORK'."
echo ""