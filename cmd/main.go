package main

import (
	"fmt"
	"log"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/app"
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/datagen"
)

func main() {
	fmt.Println("Hello project")

	if errLE := app.LoadEnv(); errLE != nil {
		log.Fatalf("failed load env: %e", errLE)
	}

	//ps, errN := cassandradb.New()
	//if errN != nil {
	//	log.Fatalf("new is failed %e", errN)
	//}

	//if err := app.Run(ps); err != nil {
	//	log.Fatalf("run is failed %e", err)
	//}

	nspg := datagen.NewStreamsGenerator(datagen.NewStreamParams(10, 5, 100))

	points := nspg.GenerateStreamChannel(make(chan struct{}))

	i := 0
	for p := range points {
		i++
		fmt.Println("i:", i, "p:", p.GetTimestamp().UnixMicro())
		fmt.Println(p.GetPoints())
	}

	fmt.Println("ура отлично работает")

	fmt.Println("Bye project")
}
