package internal

import "time"

type Point struct {
	X float64
	Y float64
}

func (p *Point) GetX() float64 {
	return p.X
}

func (p *Point) GetY() float64 {
	return p.Y
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
