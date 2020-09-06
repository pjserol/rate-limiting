# Rate limiting

Rate-limiting module that stops a particular requestor for Go using Redis.

When the limit has been reached, return a 429 error.

The application is configured by default with the strategy "Fixed Window" per IP.

## Run the application

- docker installed
- `docker-compose up`

## Run the application locally

### Requirements

- go installed
- docker installed
- start redis `./script.sh` (port 6380)

### Start the application

Default: FixedWindow

- `export REDIS_ADDRESS=localhost:6380`
- `go run .`

Start with SlidingWindowLogs:

- `export REDIS_ADDRESS=localhost:6380`
- `export RATE_LIMITER_NAME=SlidingWindowLogs`
- `go run .`

### Test locally

- `go test ./... -v`

### Integration test

- `go test ./... -tags=integration`

## Example

GET `http://localhost:8080/`

After too many requests:

>Body
>
>Rate limit exceeded. Try again in 24 seconds.
>
>Headers
>
>HTTP/1.1 429 Too Many Requests  
>Content-Type: text/plain; charset=utf-8  
>X-Content-Type-Options: nosniff  
>X-Rate-Limit-Limit: 5  
>X-Rate-Limit-Remaining: -1  
>X-Rate-Limit-Reset: 19  
>Date: Sun, 06 Sep 2020 05:10:41 GMT  
>Content-Length: 46

## Flush Redis cache

- `docker exec -it rate-app-redis redis-cli`
- `FLUSHALL`

## Access Go logs in Docker

- `docker logs rate-app-go`

## Rebuild docker image

- `docker-compose build`
