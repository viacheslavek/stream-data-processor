package app

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

func LoadEnv() error {
	return godotenv.Load()
}

func Run(s storage.Storage) error {

	log.Printf("app start\n")

	//bf, err := benchmark.RunBenchNCheck(storage, 100, 1000000, 1)
	//if err != nil {
	//	log.Fatalf("can't do bench %e\n", err)
	//}
	//
	//fmt.Println("benchmark:", bf)
	//
	//benchPath := fmt.Sprintf("cmd/bench_result_%s.html", storage.Name())
	//if err = bf.ConvertToHTML(benchPath); err != nil {
	//	log.Fatalf("can't convert bf to html, %e", err)
	//}

	stat, err := storage.GetDockerMemoryUsage("stream-data-processor-postgres-1")
	if err != nil {
		return fmt.Errorf("блин блинский %w", err)
	}

	fmt.Println("stat", stat)

	log.Println("все")

	return nil
}
