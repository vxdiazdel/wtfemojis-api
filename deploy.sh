#!/bin/bash

set -euo pipefail

cd /home/deploy/app

# export environment vars
export DB_HOST=${DB_HOST}
export DB_PORT=${DB_PORT}
export DB_USER=${DB_USER}
export DB_PASSWORD=${DB_PASSWORD}
export DB_NAME=${DB_NAME}
export DB_URL=${DB_URL}
export ENV=production

# rename docker-compose file
mv docker-compose.prod.yml docker-compose.yml

# pull and run docker image
echo "$DOCKERHUB_TOKEN" | docker login -u "$DOCKERHUB_USER" --password-stdin
docker pull "$IMAGE"
docker compose up -d --force-recreate
