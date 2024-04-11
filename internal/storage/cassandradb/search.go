package cassandradb

import (
	"time"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) GetStreamRange(from, to time.Time) ([]internal.Stream, error) {

	q := `
		SELECT * FROM streams.stream 
		WHERE cluster = 'alone'
			AND time_point >= ?
			AND time_point <= ?;
	`

	iter := s.session.Query(q, from, to).Iter()

	var streams []internal.Stream

	var cluster string
	var timePoint time.Time
	var points []internal.Point
	for iter.Scan(&cluster, &timePoint, &points) {
		streams = append(streams, internal.NewStream(points, timePoint))
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return streams, nil
}
