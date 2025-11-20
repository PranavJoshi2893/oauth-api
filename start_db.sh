#!/usr/bin/env bash

ENV_FILE=".env"

if [ -f "$ENV_FILE" ]; then
    echo "Loading environment variables...."

    set -a
    . "$ENV_FILE"
    set +a

    echo "Environment variables loaded."

    docker-compose down -v
    docker-compose up -d

    MAX_RETRIES=30
    RETRY_COUNT=0
    until docker exec postgresdb pg_isready -U "$DB_USER" > /dev/null 2>&1; do
        RETRY_COUNT=$((RETRY_COUNT+1))
        
        if [ $RETRY_COUNT -ge $MAX_RETRIES ]; then
            echo "PostgreSQL failed to start after $MAX_RETRIES attempts"
            exit 1
        fi
        
        echo "Connecting... (attempt $RETRY_COUNT)"
        sleep 5
    done

    echo "PostgreSQL is ready to accept connections."

else
    echo "Warning: .env file not found or could not be loaded."
    exit 1
fi
