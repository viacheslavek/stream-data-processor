package redisdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) GetStreamRange(from, to time.Time) ([]internal.Stream, error) {

	result, err := s.client.XRange(streamName,
		strconv.FormatInt(from.UnixMilli(), 10),
		strconv.FormatInt(to.UnixMilli(), 10)).Result()

	if err != nil {
		return []internal.Stream{}, fmt.Errorf("failed to get range %w", err)
	}

	return convertRedisTypeStreamToInternal(result)
}

func convertRedisTypeStreamToInternal(result []redis.XMessage) ([]internal.Stream, error) {

	streams := make([]internal.Stream, len(result))

	for i, r := range result {
		redisPoints, ok := r.Values["points"].(string)
		if !ok {
			return []internal.Stream{},
				fmt.Errorf("incorrect points type: required string, recived: %T", r.Values["points"])
		}

		points, errP := convertRedisTypePointsToInternal(redisPoints)
		if errP != nil {
			return []internal.Stream{}, errP
		}
		timestamp, errT := convertRedisTypeTimeToInternal(r.ID)
		if errT != nil {
			return []internal.Stream{}, errT
		}

		streams[i] = internal.NewStream(points, timestamp)
	}

	return streams, nil
}

func convertRedisTypePointsToInternal(redisPoints string) ([]internal.Point, error) {
	points := make(internal.Points, 0)

	err := points.UnmarshalBinary([]byte(redisPoints))
	if err != nil {
		return internal.Points{}, fmt.Errorf("failed unmarshal points: %w", err)
	}

	return points, nil
}

func convertRedisTypeTimeToInternal(redisTime string) (time.Time, error) {
	unixTime, err := strconv.ParseInt(strings.Split(redisTime, "-")[0], 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed parse unixTime %w", err)
	}

	timestamp := time.Unix(0, unixTime*int64(time.Millisecond))
	return timestamp, nil
}
