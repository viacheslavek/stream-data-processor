package boltdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) GetStreamRange(from, to time.Time) ([]internal.Stream, error) {
	streams := make([]internal.Stream, 0)

	if err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		c := bucket.Cursor()

		startKey := convertTimestampTypeToBoltdb(from)
		endKey := convertTimestampTypeToBoltdb(to)

		for k, v := c.Seek(startKey); k != nil && bytes.Compare(k, endKey) <= 0; k, v = c.Next() {
			points, errP := convertPointsTypeFromBoltdb(v)
			if errP != nil {
				return fmt.Errorf("failed convert points %w", errP)
			}

			timestamp, errT := convertTimestampTypeFromBoltdb(k)
			if errT != nil {
				return fmt.Errorf("failed convert timestamp %w", errT)
			}

			streams = append(streams, internal.NewStream(points, timestamp))
		}
		return nil
	}); err != nil {
		return streams, fmt.Errorf("failed to get streams %w", err)
	}

	return streams, nil
}

func convertTimestampTypeFromBoltdb(data []byte) (time.Time, error) {
	unixTime, err := strconv.Atoi(string(data))
	if err != nil {
		return time.Time{}, fmt.Errorf("failed convert time %w", err)
	}
	return time.UnixMilli(int64(unixTime)), nil
}

func convertPointsTypeFromBoltdb(data []byte) ([]internal.Point, error) {
	points := make([]internal.Point, 0)

	err := json.Unmarshal(data, &points)
	if err != nil {
		return points, fmt.Errorf("failed to unmarshal points %w", err)
	}
	return points, nil
}
