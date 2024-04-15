package cassandradb

import (
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) AddStream(stream internal.Stream) error {

	query := "INSERT INTO streams.stream (cluster, time_point, points) VALUES ('alone', ?, ?)"

	if err := s.session.Query(query, stream.GetTimestamp(), stream.GetPoints()).Exec(); err != nil {
		return err
	}

	return nil
}
