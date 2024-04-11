package app

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

func LoadEnv() error {
	return godotenv.Load()
}

func Run(storage storage.Storage) error {

	//errP := storage.Ping()
	//if errP != nil {
	//	return fmt.Errorf("ping err: %w", errP)
	//}
	//
	//time.Sleep(time.Millisecond * 500)
	//errI := storage.Init()
	//if errI != nil {
	//	return fmt.Errorf("init err: %w", errI)
	//}
	//
	//time.Sleep(time.Millisecond * 500)
	//log.Println("init success")
	//
	//p1, p2 := internal.NewPoint(0.1, 1.1), internal.NewPoint(2.2, 2.3)
	//errA := storage.AddStream(internal.NewStream([]internal.Point{p1, p2}, time.Now()))
	//if errA != nil {
	//	return fmt.Errorf("add err: %w", errA)
	//}
	//
	//time.Sleep(time.Millisecond * 1000)
	//log.Println("add success")

	timeString1 := "2023-01-01 00:00:00.000000"
	timeString2 := "2025-01-08 00:00:00.000000"
	layout := "2006-01-02 15:04:05.000000"

	t1, err := time.Parse(layout, timeString1)
	if err != nil {
		fmt.Println("Ошибка при парсинге времени:", err)
	}
	t2, err := time.Parse(layout, timeString2)
	if err != nil {
		fmt.Println("Ошибка при парсинге времени:", err)
	}

	streams, errG := storage.GetStreamRange(t1, t2)
	if errG != nil {
		return fmt.Errorf("search err: %w", errG)
	}

	time.Sleep(time.Millisecond * 500)
	fmt.Println("streams:", streams)
	log.Println("get success")

	//errD := storage.Drop()
	//if errD != nil {
	//	return fmt.Errorf("drop err: %w", errD)
	//}
	//
	//time.Sleep(time.Millisecond * 1000)
	//
	//log.Println("drop success")

	errC := storage.Close()
	if errC != nil {
		return fmt.Errorf("close err: %w", errC)
	}

	time.Sleep(time.Millisecond * 500)

	log.Println("close success")

	return nil
}
