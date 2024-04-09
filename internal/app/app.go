package app

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

func LoadEnv() error {
	return godotenv.Load()
}

func Run(storage storage.Storage) error {

	err := storage.Ping()
	if err != nil {
		fmt.Println("ping is error...")
	}

	return err
}
