package cassandradb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

type Storage struct {
	session *gocql.Session
	ctx     context.Context
}

const (
	url           = "CASSANDRA_URL"
	containerName = "stream-data-processor-cassandra-1"
)

func New() (*Storage, error) {
	cassandraURL := os.Getenv(url)
	if len(cassandraURL) == 0 {
		log.Fatalf("cassandraURL not find")
	}

	cluster := gocql.NewCluster(cassandraURL)
	cluster.Port = 9042
	session, err := cluster.CreateSession()
	if err != nil {
		return &Storage{}, fmt.Errorf("failed to create Cassandra session: %w", err)
	}

	log.Println("Cassandra db init")

	return &Storage{
		session: session,
		ctx:     context.Background(),
	}, nil
}

func (s *Storage) Ping() error {
	pingQuery := "SELECT * FROM system.local"
	if err := s.session.Query(pingQuery).WithContext(s.ctx).Exec(); err != nil {
		return fmt.Errorf("Ping is failed: %w\n", err)
	}

	log.Println("Cassandra ping success")
	return nil
}

func (s *Storage) Close() error {
	if s.session != nil {
		s.session.Close()
	}
	return nil
}

func (s *Storage) Init() error {
	queryCreateKeyspace := `
		CREATE KEYSPACE IF NOT EXISTS streams WITH replication = {
    		'class': 'SimpleStrategy',
    		'replication_factor': 1
    	};
	`

	if err := s.session.Query(queryCreateKeyspace).Exec(); err != nil {
		return fmt.Errorf("failed to create keyspace %w", err)
	}

	queryCreateTable := `
		CREATE TABLE IF NOT EXISTS streams.stream (
    		cluster text,
    		time_point timestamp,
    		points list<tuple<double, double>>,
    		PRIMARY KEY (cluster, time_point)
		);
	`

	if err := s.session.Query(queryCreateTable).Exec(); err != nil {
		return fmt.Errorf("failed to create table %w", err)
	}

	return nil
}

func (s *Storage) Drop() error {
	q := `
		DROP TABLE IF EXISTS streams.stream;
	`

	if err := s.session.Query(q).Exec(); err != nil {
		return fmt.Errorf("failed to drop table %w", err)
	}

	return nil
}

func (s *Storage) Info() {
	log.Println("cassandra")
}

func (s *Storage) Name() string {
	return "cassandra"
}

func (s *Storage) GetUsageMemory() (uint64, error) {
	return storage.GetDockerMemoryUsage(containerName)
}
