package benchmark

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/datagen"
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal/storage"
)

func runStreamBench(
	f func(s storage.Storage) error, s storage.Storage, bf *BenchFile, benchName string,
) error {
	start := time.Now()
	if err := f(s); err != nil {
		return err
	}
	bf.Durations[benchName] += time.Since(start)
	return nil
}

func benchInit(s storage.Storage) error {
	if err := s.Init(); err != nil {
		return fmt.Errorf("failed init stream from init bench %w", err)
	}
	return nil
}

func benchAdd(s storage.Storage) error {
	if err := addStreams(s); err != nil {
		return fmt.Errorf("failed add stream from add bench %w", err)
	}
	return nil
}

func addStreams(s storage.Storage) error {

	// TODO: пока параметры делаю хардкодом, потом исправлю
	streamGen := datagen.NewStreamsGenerator(datagen.NewStreamParams(10000, 100, 100))
	streams := streamGen.GenerateStreams()
	for _, stream := range streams {
		if err := s.AddStream(stream); err != nil {
			return err
		}
	}
	return nil
}

func benchSearch(s storage.Storage) error {
	// TODO: заведу еще переменную startAddTime в env как UnixTime, в которую буду класть время начала для добавления
	startAdd := time.Now()
	countStreams, timePeriodMiliSec := 10, 100
	nRangeSearch := 10

	for i := 0; i < nRangeSearch; i++ {
		if _, err := s.GetStreamRange(getRangeTimestamp(startAdd, countStreams, timePeriodMiliSec)); err != nil {
			return fmt.Errorf("failed find streams from find bench %w", err)
		}
	}

	return nil
}

func getRangeTimestamp(start time.Time, countStreams, timePeriodMiliSec int) (time.Time, time.Time) {
	positionSearch := rand.IntN(10) + 1

	startRangeSearch := start.Add(
		time.Millisecond * time.Duration(getPartOfRangeTimestamp(timePeriodMiliSec, countStreams, positionSearch)))
	finishRangeSearch := startRangeSearch.Add(
		time.Millisecond * time.Duration(getCountRangeTimestamp(timePeriodMiliSec, countStreams)))

	return startRangeSearch, finishRangeSearch
}

func getPartOfRangeTimestamp(timePeriodMiliSec, countStreams, positionSearch int) int {
	return timePeriodMiliSec * countStreams / positionSearch
}

// INFO: введем метрику, что число возвращаемых потоков должно не быстро расти и быть около 50-150
// Для этого подойдет функция sqrt(x) + (x mod 50), где x - общее число потоков
func getCountRangeTimestamp(timePeriodMiliSec, countStreams int) int {
	return timePeriodMiliSec * (int(math.Sqrt(float64(countStreams))) + countStreams%50)
}

func benchDrop(s storage.Storage) error {
	if err := s.Drop(); err != nil {
		return fmt.Errorf("failed drop stream from drop bench %w", err)
	}
	return nil
}

func runMemoryBench(s storage.Storage, u *uint64) error {
	memoryUsage, err := s.GetUsageMemory()
	if err != nil {
		return fmt.Errorf("failed get memory usage in bench %w", err)
	}
	*u = memoryUsage
	return nil
}

func getDifferentMemoryUsage(bm BenchMemory) uint64 {
	if bm.FinishMemoryUsage < bm.StartMemoryUsage {
		return 0
	}
	return bm.FinishMemoryUsage - bm.StartMemoryUsage
}
