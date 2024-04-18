package postgresdb

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) AddStream(stream internal.Stream) error {

	tx, err := s.conn.BeginTx(s.ctx, pgx.TxOptions{})

	if err != nil {
		return fmt.Errorf("can't begin transaction: %w\n", err)
	}

	q := fmt.Sprintf("INSERT INTO stream (time_point, points) VALUES ($1, array[%s]);",
		getPointsQuery(stream.GetPoints()))

	layout := "2006-01-02 15:04:05.000000"

	_, err = tx.Exec(s.ctx, q, stream.GetTimestamp().Format(layout))

	if err != nil {
		_ = tx.Rollback(s.ctx)
		return fmt.Errorf("can't add a point: %w\n", err)
	}

	err = tx.Commit(s.ctx)

	if err != nil {
		return fmt.Errorf("can't commit transactions %w\n", err)
	}

	return nil
}

func getPointsQuery(points []internal.Point) string {
	var values []string
	for _, p := range points {
		values = append(values, fmt.Sprintf("(point(%f, %f))", p.GetX(), p.GetY()))
	}
	return strings.Join(values, ", ")
}
