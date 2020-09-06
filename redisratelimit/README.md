# Redis rate limit

The key can be an IP, an IP + routes, a token, etc.

## Strategies implemented

### Fixed Window

- start a transaction on the Redis store
- when empty, set a key on Redis with an expiry time
- when the key is set, increment the counter
- execute the transaction

We do these operations atomically to be safe from race conditions.

**Cons**  
If all the requests are executed at the end of the window and if at the beginning of the new window new requests are executed, in the time range, the limit will be overtaken.

### Sliding Window Logs

- start a transaction on the Redis store
- ZRemRangeByScore, remove requests before the time window
- ZRange, get the lists of the requests (timestamps)
- ZAdd, add a log of the current request
- Expire, set a new expire time on the key
- execute the transaction

**Cons**  
We can overtake the limit if we have multiple requests on the same timestamp.

After we reach the limit, we continue to add new logs on Redis. We will have to wait for all the logs to be removed from Redis to execute a new successful request.

**Improvements**  
Lock the number of requests stored in Redis when the limit is reached.

## TODO

Implemtation of the strategy "Sliding window counters"

## Resouces

- `https://konghq.com/blog/how-to-design-a-scalable-rate-limiting-algorithm/`
- `https://www.figma.com/blog/an-alternative-approach-to-rate-limiting/`
