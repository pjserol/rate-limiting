#!/bin/bash

export APP_ENVIRONMENT=local
export REDIS_ADDRESS=localhost:6380

docker pull redis:6.0.7
docker run --name rate-app-redis-local -p 6380:6379 -d redis:6.0.7