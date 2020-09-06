package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-rate-limiting/config"
	"simple-rate-limiting/redisratelimit"
)

var rdb *redisratelimit.Client
var env config.Environment

func init() {
	env = config.InitEnvironment()
	log.Printf("InitEnvironment::%+v\n", env)

	conf := redisratelimit.Config{
		Address:            env.RedisAddress,
		Password:           env.RedisPassword,
		DB:                 env.RedisDB,
		DialTimeOutSeconds: env.RedisDialTimeOutSeconds,
		PoolSize:           env.RedisPoolSize,
	}

	rdb = redisratelimit.NewClient(conf)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)

	log.Printf("Listening on :%d", env.PortNumber)
	http.ListenAndServe(fmt.Sprintf(":%d", env.PortNumber), limit(mux))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
