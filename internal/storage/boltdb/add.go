package boltdb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) AddStream(stream internal.Stream) error {
	if err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		ps, err := convertPointsTypeToBoltdb(stream.GetPoints())
		if err != nil {
			return err
		}
		return bucket.Put(
			convertTimestampTypeToBoltdb(stream.GetTimestamp()),
			ps)
	}); err != nil {
		return fmt.Errorf("failed add stream %w", err)
	}

	return nil
}

func convertTimestampTypeToBoltdb(timestamp time.Time) []byte {
	return []byte(strconv.FormatInt(timestamp.UnixMilli(), 10))
}

func convertPointsTypeToBoltdb(points []internal.Point) ([]byte, error) {
	data, err := json.Marshal(points)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal points %w", err)
	}
	return data, nil
}
