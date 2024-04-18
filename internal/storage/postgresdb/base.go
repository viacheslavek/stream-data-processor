package postgresdb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

type Storage struct {
	conn *pgx.Conn
	ctx  context.Context
}

const (
	url           = "POSTGRES_URL"
	containerName = "stream-data-processor-postgres-1"
)

func New() (*Storage, error) {

	postgresURL := os.Getenv(url)

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

	var greeting string
	err := s.conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return fmt.Errorf("QueryRow failed: %w\n", err)
	}

	log.Println("QueryRow success")

	return nil
}

func (s *Storage) Close() error {
	err := s.conn.Close(s.ctx)
	if err != nil {
		return fmt.Errorf("can't close connection %e\n", err)
	}
	return nil
}

func (s *Storage) Init() error {
	q := `
		CREATE TABLE IF NOT EXISTS stream (
			time_point timestamptz primary key,
			points point[]
		);
	`

	return s.initBase(s.ctx, q)
}

func (s *Storage) initBase(ctx context.Context, query string) error {
	_, err := s.conn.Exec(
		ctx,
		query,
	)

	if err != nil {
		return fmt.Errorf("can't create tables %w", err)
	}

	return nil
}

func (s *Storage) Drop() error {
	q := `
		DROP TABLE IF EXISTS stream;
	`

	return s.drop(s.ctx, q)
}

func (s *Storage) drop(ctx context.Context, query string) error {
	_, err := s.conn.Exec(
		ctx,
		query,
	)
	if err != nil {
		return fmt.Errorf("can't drop tables %w", err)
	}

	return nil
}

func (s *Storage) Info() {
	q := `
		SELECT table_catalog, table_schema, table_name
		FROM information_schema.tables
		WHERE table_name = 'stream';
	`

	err := s.info(s.ctx, q)
	if err != nil {
		log.Fatalf("info failed: %e", err)
	}
}

func (s *Storage) info(ctx context.Context, query string) error {
	output, err := s.conn.Query(
		ctx,
		query,
	)
	if err != nil {
		return fmt.Errorf("can't drop tables %w", err)
	}

	log.Println("info output")
	log.Println(output.RawValues())

	return nil
}

func (s *Storage) Name() string {
	return "postgres"
}

func (s *Storage) GetUsageMemory() (uint64, error) {
	return storage.GetDockerMemoryUsage(containerName)
}
