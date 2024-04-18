package boltdb

import (
	"context"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type Storage struct {
	db  *bolt.DB
	ctx context.Context
}

const (
	dbPath     = "./data_volumes/boltdb/data.db"
	bucketName = "streams"
)

func New() (*Storage, error) {
	db, err := bolt.Open(dbPath, 0600, nil)

	if err != nil {
		return &Storage{}, err
	}

	log.Println("Boltdb database create")

	return &Storage{
		db:  db,
		ctx: context.Background(),
	}, nil
}

func (s *Storage) Ping() error {
	log.Println("boltdb always ready")
	return nil
}

func (s *Storage) Close() error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			return fmt.Errorf("failed to close boltDB %w", err)
		}
	}
	return nil
}

func (s *Storage) Init() error {
	if err := s.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketName)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed create bucket %w", err)
	}
	return nil
}

func (s *Storage) Drop() error {
	if err := s.db.Update(func(tx *bolt.Tx) error {
		if bucketExist := tx.Bucket([]byte(bucketName)); bucketExist == nil {
			return nil
		}

		if err := tx.DeleteBucket([]byte(bucketName)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed create bucket %w", err)
	}
	return nil
}

func (s *Storage) Info() {
	log.Println("cassandra")
}

func (s *Storage) Name() string {
	return "boltdb"
}

func (s *Storage) GetUsageMemory() (uint64, error) {
	// TODO: реализовать
	return 0, nil
}
