// +build integration

package redisratelimit

import (
	"os"
	"testing"
)

func initRedis() *Client {
	conf := Config{
		Address: os.Getenv("REDIS_ADDRESS"),
	}

	return NewClient(conf)
}
func Test_FixedWindowRateLimiter(t *testing.T) {
	rdb := initRedis()

	tests := []struct {
		name           string
		requests       int
		expectedOutput Limit
	}{
		{
			name:     "test happy path",
			requests: 1,
			expectedOutput: Limit{
				Remaining: 3,
			},
		},
		{
			name:     "test too many requests",
			requests: 5,
			expectedOutput: Limit{
				Remaining: -1,
			},
		},
	}
	for _, test := range tests {
		rdb.flushAll()

		for i := 5; i < test.requests; i++ {
			rdb.FixedWindowRateLimiter("key_other", 5, 20)
		}

		for i := 0; i < test.requests; i++ {
			rdb.FixedWindowRateLimiter("k123", 5, 20)
		}

		limit, err := rdb.FixedWindowRateLimiter("k123", 5, 20)

		if err != nil {
			t.Error(err)
		}

		if test.expectedOutput.Remaining != limit.Remaining {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, limit)
		}
	}
}

func Test_SlidingWindowLogsRateLimiter(t *testing.T) {
	rdb := initRedis()

	tests := []struct {
		name           string
		requests       int
		expectedOutput Limit
	}{
		{
			name:     "test happy path",
			requests: 1,
			expectedOutput: Limit{
				Remaining: 3,
			},
		},
		{
			name:     "test too many requests",
			requests: 5,
			expectedOutput: Limit{
				Remaining: -1,
			},
		},
	}
	for _, test := range tests {
		rdb.flushAll()

		for i := 5; i < test.requests; i++ {
			rdb.SlidingWindowLogsRateLimiter("key_other", 5, 20)
		}

		for i := 0; i < test.requests; i++ {
			rdb.SlidingWindowLogsRateLimiter("k123", 5, 20)
		}

		limit, err := rdb.SlidingWindowLogsRateLimiter("k123", 5, 20)

		if err != nil {
			t.Error(err)
		}

		if test.expectedOutput.Remaining != limit.Remaining {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, limit)
		}
	}
}
