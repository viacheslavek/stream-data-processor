package redisdb

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
	"github.com/go-redis/redis"
)

func (s *Storage) AddStream(stream internal.Stream) error {
	if err := s.client.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"points": internal.Points(stream.GetPoints()),
		},
	}).Err(); err != nil {
		return fmt.Errorf("failed to add points %w", err)
	}
	return nil
}
