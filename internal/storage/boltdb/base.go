package boltdb

import (
	"context"
	"log"

	"github.com/boltdb/bolt"
)

type Storage struct {
	db  *bolt.DB
	ctx context.Context
}

const dbPath = "./data_volumes/boltdb/data.db"

func New() (*Storage, error) {
	db, err := bolt.Open(dbPath, 0600, nil)

	if err != nil {
		return &Storage{}, err
	}

	log.Println("Boltdb database init")

	return &Storage{
		db:  db,
		ctx: context.Background(),
	}, nil
}

func (S *Storage) Ping() error {
	log.Println("Boltdb always available")
	return nil
}
