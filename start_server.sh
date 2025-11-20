#!/usr/bin/env bash

ENV_FILE=".env"

if [ -f "$ENV_FILE" ]; then
    echo "Loading environment variables...."

    set -a
    . "$ENV_FILE"
    set +a

    echo "Environment variables loaded."

    go run cmd/main.go

else
    echo "Warning: .env file not found or could not be loaded."
    exit 1
fi
