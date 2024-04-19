package app

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/benchmark"
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

func LoadEnv() error {
	return godotenv.Load()
}

func Run(s storage.Storage) error {

	log.Printf("app start\n")

	bf, err := benchmark.RunBenchNCheck(s, 10, 10, 1)
	if err != nil {
		log.Fatalf("can't do bench %e\n", err)
	}

	fmt.Println("benchmark:", bf)

	benchPath := fmt.Sprintf("cmd/bench_result_%s.html", s.Name())
	if err = bf.ConvertToHTML(benchPath); err != nil {
		log.Fatalf("can't convert bf to html, %e", err)
	}

	return nil
}
