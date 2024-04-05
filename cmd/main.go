package main

import (
	"fmt"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/app"
)

func main() {
	fmt.Println("Hello project")

	_ = app.Run()

	fmt.Println("Bye project")
}
