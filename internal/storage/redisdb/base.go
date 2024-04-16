package redisdb

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

const streamName = "points_stream"

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

func (s *Storage) Close() error {
	if s.client != nil {
		err := s.client.Close()
		if err != nil {
			return fmt.Errorf("failed to close client: %w", err)
		}
	}
	return nil
}

func (s *Storage) Init() error {
	if s.client.Exists(streamName).Val() == 0 {
		if err := s.client.XGroupCreateMkStream(
			streamName, "points_group", "$").Err(); err != nil {
			return fmt.Errorf("failed to create stream %w", err)
		}
	}

	return nil
}

func (s *Storage) Drop() error {
	if err := s.client.Del(streamName).Err(); err != nil {
		return fmt.Errorf("failed to delete stream %w", err)
	}
	return nil
}

func (s *Storage) Info() {
	log.Println("redis")
}

func (s *Storage) Name() string {
	return "redis"
}
