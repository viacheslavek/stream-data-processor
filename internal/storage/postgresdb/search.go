package postgresdb

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

func (s *Storage) GetStreamRange(from, to time.Time) ([]internal.Stream, error) {

	q := `
		SELECT *
		FROM stream
		WHERE time_point BETWEEN $1 AND $2;
	`

	layout := "2006-01-02 15:04:05.000000"
	rows, err := s.conn.Query(s.ctx, q, from.Format(layout), to.Format(layout))

	if err != nil {
		return nil, fmt.Errorf("can't querry points in polygon %w\n", err)
	}

	defer rows.Close()

	return translateRowsStreams(rows)
}

func translateRowsStreams(rows pgx.Rows) ([]internal.Stream, error) {
	streams := make([]internal.Stream, 0)

	for rows.Next() {
		var timestampData time.Time
		var pointData []pgtype.Point
		err := rows.Scan(&timestampData, &pointData)
		if err != nil {
			return nil, fmt.Errorf("can't scan row in streams %w\n", err)
		}

		streams = append(streams, internal.NewStream(convertPGTypePointsToInternalPoints(pointData), timestampData))
	}

	return streams, nil
}

func convertPGTypePointsToInternalPoints(pointData []pgtype.Point) []internal.Point {
	points := make([]internal.Point, len(pointData))

	for i, p := range pointData {
		points[i] = internal.NewPoint(p.P.X, p.P.Y)
	}

	return points
}
