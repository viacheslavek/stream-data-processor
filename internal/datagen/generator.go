package datagen

import (
	"math/rand"
	"time"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

const MaxPointCoordinate = 255999

type StreamsGenerator struct {
	StreamParams
}

func NewStreamsGenerator(sp StreamParams) StreamsGenerator {
	return StreamsGenerator{sp}
}

func (ssg *StreamsGenerator) GenerateStreamChannel(cancel <-chan struct{}) <-chan internal.Stream {
	out := make(chan internal.Stream)

	go func() {
		defer close(out)

		ticker := time.NewTicker(time.Duration(ssg.miliSecPeriod) * time.Millisecond)
		defer ticker.Stop()

		for i := 0; i < ssg.countStreams; i++ {
			select {
			case out <- generateStream(time.Now(), ssg.lenStream):
				<-ticker.C
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func generateStream(timestamp time.Time, lenStream int) internal.Stream {
	return internal.NewStream(getRandomPoints(lenStream), timestamp)
}

func getRandomPoints(lenStream int) []internal.Point {
	points := make([]internal.Point, lenStream)

	for i := 0; i < lenStream; i++ {
		points[i] = getRandomPoint()
	}
	return points
}

// INFO: Получаю совсем случайную точку
func getRandomPoint() internal.Point {
	return internal.NewPoint(
		float64(rand.Intn(MaxPointCoordinate))/1000.0,
		float64(rand.Intn(MaxPointCoordinate))/1000.0,
	)
}

func (ssg *StreamsGenerator) GenerateStreams() []internal.Stream {
	streams := make([]internal.Stream, ssg.countStreams)

	startTimestamp := time.Now()

	for i := 0; i < ssg.countStreams; i++ {
		streams[i] = generateStream(
			startTimestamp.Add(time.Millisecond*time.Duration(ssg.miliSecPeriod*i)),
			ssg.lenStream,
		)
	}
	return streams
}
