package postgresdb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn *pgx.Conn
	ctx  context.Context
}

const Url = "POSTGRES_URL"

func New() (*Storage, error) {

	postgresURL := os.Getenv(Url)

	if len(postgresURL) == 0 {
		log.Fatalf("postgresURL not find")
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, postgresURL)
	if err != nil {
		return &Storage{}, err
	}

	log.Println("Postgres conn init")

	return &Storage{
		conn: conn,
		ctx:  ctx,
	}, nil
}

func (s *Storage) Ping() error {
	if err := s.conn.Ping(s.ctx); err != nil {
		return fmt.Errorf("Ping is failed: %w\n", err)
	}
	log.Println("Postgres ping success")
	return nil
}
