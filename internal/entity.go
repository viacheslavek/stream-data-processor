package internal

import (
	"time"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

type Stream struct {
	points    []Point
	timestamp time.Time
}

func (s *Stream) GetPoints() []Point {
	return s.points
}

func (s *Stream) GetTimestamp() time.Time {
	return s.timestamp
}

func NewStream(points []Point, timestamp time.Time) Stream {
	return Stream{points, timestamp}
}
