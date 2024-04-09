package datagen

import (
	"fmt"
	"time"

	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
)

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

		ticker := time.NewTicker(time.Duration(ssg.miniSecPeriod) * time.Millisecond)
		defer ticker.Stop()

		for i := 0; i < ssg.countStreams; i++ {
			select {
			case out <- generateRealStream(ticker, ssg.lenStream):
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func generateRealStream(ticker *time.Ticker, lenStream int) internal.Stream {
	// TODO: доделать: генерирую timestamp через ticker + массив точек
	return internal.Stream{}
}

func (ssg *StreamsGenerator) GenerateStreams() []internal.Stream {
	fmt.Println("gen streams")
	fmt.Println("ssg:", *ssg)
	return make([]internal.Stream, 0)
}

func generateFakeStream(countStreams, lenStream int) []internal.Stream {
	// TODO: доделать: генерирую timestamp руками + массив точек
	return make([]internal.Stream, 0)
}
