package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

type redisStore struct {
}

func (r redisStore) init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func (r redisStore) addToStore(su string, url string) {
	client.Set(su, url, 0)
	client.Set(url, su, 0)
}

func (r redisStore) lookUp(surl string) string {
	v, err := client.Get(surl[7:]).Result()
	if err == redis.Nil {
		return fmt.Sprintf("%v cannot be resolved", surl)
	}
	if err != nil {
		panic(err)
	}
	return v
}

func (r redisStore) lookUpLongURL(url string) (string, bool) {
	val, err := client.Get(url).Result()
	if err == redis.Nil {
		return val, false
	}
	if err != nil {
		panic(err)
	}
	return val, true
}
