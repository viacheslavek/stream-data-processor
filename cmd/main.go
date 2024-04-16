package main

import (
	"fmt"
	"log"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/app"
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage/boltdb"
)

func main() {
	fmt.Println("Hello project")

	if errLE := app.LoadEnv(); errLE != nil {
		log.Fatalf("failed load env: %e", errLE)
	}

	ps, errN := boltdb.New()
	if errN != nil {
		log.Fatalf("new is failed %e", errN)
	}

	if err := app.Run(ps); err != nil {
		log.Fatalf("run is failed %e", err)
	}

	fmt.Println("Bye project")
}
