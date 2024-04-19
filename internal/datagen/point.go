package datagen

import "github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"

type StreamParams struct {
	countStreams  int
	lenStream     int
	miliSecPeriod int
}

func NewStreamParams(countStreams, lenStream, miniSecPeriod int) StreamParams {
	return StreamParams{
		countStreams:  countStreams,
		lenStream:     lenStream,
		miliSecPeriod: miniSecPeriod,
	}
}

type GetStreams interface {
	GenerateStreamChannel(cancel <-chan struct{}) <-chan internal.Stream
	GenerateStreams() []internal.Stream
}
