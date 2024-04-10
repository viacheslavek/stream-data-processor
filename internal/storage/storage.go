package storage

import (
	"github.com/VyacheslavIsWorkingNow/stream-data-processor/internal"
	"time"
)

type Storage interface {
	Ping() error
	Close() error
	Init() error
	Drop() error
	AddStream(stream internal.Stream) error
	GetStreamRange(from, to time.Time) ([]internal.Stream, error)
	Info()
}
