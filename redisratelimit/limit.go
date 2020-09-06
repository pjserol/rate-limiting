package redisratelimit

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type Limit struct {
	Remaining         int
	RetryAfterSeconds int
}

// FixedWindowRateLimiter implementation of the algorithm fixed window
// increment the counter of a key to limit the number of requests of a key in a time window
func (c *Client) FixedWindowRateLimiter(key string, maxQuotas, windowSecs int) (Limit, error) {
	pipe := c.rdb.TxPipeline()
	pipe.SetNX(key, 0, time.Duration(windowSecs)*time.Second)
	used := pipe.Incr(key)

	_, err := pipe.Exec()
	if err != nil {
		return Limit{}, err
	}

	ttl := c.rdb.TTL(key).Val().Seconds()

	return Limit{
			Remaining:         getRemaining(maxQuotas, int(used.Val())),
			RetryAfterSeconds: int(ttl)},
		nil
}

func getRemaining(maxQuotas, used int) (remaining int) {
	remaining = maxQuotas - used
	if remaining < 0 {
		return -1
	}
	return
}

// SlidingWindowLogsRateLimiter implementation of the algorithm sliding window logs
// add a new log to limit the number of requests of a key in a time window
func (c *Client) SlidingWindowLogsRateLimiter(key string, maxQuotas, windowSecs int) (Limit, error) {
	window := time.Duration(windowSecs) * time.Second

	now := time.Now()
	nowStr := strconv.FormatInt(now.UnixNano(), 10)
	timeWindow := strconv.FormatInt(now.Add(-window).UnixNano(), 10)

	pipe := c.rdb.TxPipeline()

	pipe.ZRemRangeByScore(key, "0", timeWindow)
	resCmdZRange := pipe.ZRange(key, 0, -1)
	pipe.ZAdd(key, redis.Z{Score: float64(now.UnixNano()), Member: nowStr})
	pipe.Expire(key, window)

	_, err := pipe.Exec()
	if err != nil {
		return Limit{}, err
	}

	timestamps, err := resCmdZRange.Result()
	if err != nil {
		return Limit{}, err
	}

	if len(timestamps) >= maxQuotas {
		retry, err := getRetryInSeconds(timestamps[0], window, now)
		if err != nil {
			return Limit{}, err
		}

		return Limit{Remaining: -1, RetryAfterSeconds: retry}, nil
	}

	return Limit{Remaining: maxQuotas - len(timestamps) - 1}, nil
}

func getRetryInSeconds(timestamp string, window time.Duration, now time.Time) (int, error) {
	tstp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return 0, err
	}

	diff := time.Unix(0, tstp).Sub(now)

	return int((window - -diff).Seconds()), nil
}
