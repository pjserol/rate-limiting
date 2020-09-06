package config

// Environment environment variable of the application
type Environment struct {
	AppEnvironment string
	PortNumber     int

	// Redis
	RedisAddress            string
	RedisPassword           string
	RedisDB                 int
	RedisDialTimeOutSeconds int
	RedisPoolSize           int

	// Limit requests
	MaxQuotas       int
	WindowSeconds   int
	RateLimiterName string
}

// InitEnvironment initialise environment variable of the application
func InitEnvironment() Environment {
	return Environment{
		AppEnvironment: getEnvString("APP_ENVIRONMENT", "local"), // local, dev, stag, prod.
		PortNumber:     getEnvInt("PORT_NUMBER", 8080),

		// Redis
		RedisAddress:            getEnvString("REDIS_ADDRESS", "localhost:6379"),
		RedisPassword:           getEnvString("REDIS_PASSWORD", ""),
		RedisDB:                 getEnvInt("REDIS_DB", 0),
		RedisDialTimeOutSeconds: getEnvInt("REDIS_DIAL_TIMEOUT_SECONDS", 3),
		RedisPoolSize:           getEnvInt("REDIS_POOL_SIZE", 20),

		// Limit requests
		MaxQuotas:       getEnvInt("MAX_QUOTAS", 3),
		WindowSeconds:   getEnvInt("WINDOW_SECONDS", 10),
		RateLimiterName: getEnvString("RATE_LIMITER_NAME", "FixedWindow"), // FixedWindow, SlidingWindowLogs
	}
}
