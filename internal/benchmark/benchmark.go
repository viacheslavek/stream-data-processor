package benchmark

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

type BenchFile struct {
	Durations         map[string]time.Duration
	DbName            string
	CountStreamPoints int
	CountStreams      int
	CountChecks       int
}

func RunBenchNCheck(s storage.Storage, countStreamPoints, countStreams, countChecks int) (BenchFile, error) {
	bf := BenchFile{
		Durations:         make(map[string]time.Duration),
		DbName:            s.Name(),
		CountStreamPoints: countStreamPoints,
		CountStreams:      countStreams,
		CountChecks:       countChecks,
	}

	// TODO: вот тут подгружаю параметры bf в env
	// TODO: потом сделаю красивее названия в env

	for i := 0; i < countChecks; i++ {
		log.Println("step number:", i)
		if err := RunStreamBench(s, &bf); err != nil {
			return bf, fmt.Errorf("can't do checks in bench streams %w", err)
		}
	}

	// TODO: вот тут потом каждый event поделить на N, чтобы в отчете
	// средняя писалась

	return bf, nil
}

func RunStreamBench(s storage.Storage, bf *BenchFile) error {
	if err := runStreamBench(benchInit, s, bf, StreamInit); err != nil {
		return fmt.Errorf("failed init banch: %w", err)
	}

	if err := runStreamBench(benchAdd, s, bf, AddStreams); err != nil {
		return fmt.Errorf("failed add banch: %w", err)
	}

	if err := runStreamBench(benchSearch, s, bf, SearchStreamsInRange); err != nil {
		return fmt.Errorf("failed search banch: %w", err)
	}

	if err := runStreamBench(benchDrop, s, bf, StreamDrop); err != nil {
		return fmt.Errorf("failed drop banch: %w", err)
	}

	return nil
}

func (bf *BenchFile) ConvertToHTML(path string) error {
	wd, errG := os.Getwd()
	if errG != nil {
		return fmt.Errorf("can't get wd %w", errG)
	}

	htmlPath := filepath.Join(wd, "internal", "benchmark", "benchmark.html")
	tmpl, errP := template.ParseFiles(htmlPath)
	if errP != nil {
		return fmt.Errorf("can't parse benchmark.html in convert %w", errP)
	}

	newFilePath := filepath.Join(wd, path)
	file, errC := os.Create(newFilePath)
	if errC != nil {
		return fmt.Errorf("can't create file in convert %w", errC)
	}
	defer func() {
		_ = file.Close()
	}()

	if errE := tmpl.Execute(file, bf); errE != nil {
		return fmt.Errorf("can't execute file in convert %w", errE)
	}

	return nil
}
