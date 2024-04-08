package redisdb

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

type Storage struct {
	client *redis.Client
}

const Url = "REDIS_URL"

func New() (*Storage, error) {

	redisURL := os.Getenv(Url)

	if len(redisURL) == 0 {
		log.Fatalf("redisURL not find")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
		DB:   0,
	})

	log.Println("Redis db init")

	return &Storage{
		client: rdb,
	}, nil
}

func (s *Storage) Ping() error {
	if _, err := s.client.Ping().Result(); err != nil {
		return fmt.Errorf("Ping is failed: %w\n", err)
	}
	log.Println("Redis ping success")
	return nil
}
