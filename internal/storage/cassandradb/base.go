package cassandradb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

type Storage struct {
	session *gocql.Session
	ctx     context.Context
}

const Url = "CASSANDRA_URL"

func New() (*Storage, error) {

	cassandraURL := os.Getenv(Url)
	if len(cassandraURL) == 0 {
		log.Fatalf("postgresURL not find")
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
