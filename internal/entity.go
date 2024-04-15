package internal

import (
	"encoding/json"
	"fmt"
	"time"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
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

type Points []Point

func (pts Points) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(pts)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal point: %w", err)
	}

	return data, nil
}

func (pts *Points) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, pts); err != nil {
		return fmt.Errorf("failed to unmarshal points %w", err)
	}

	return nil
}
