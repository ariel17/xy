package domain

import "time"

// Position Earth coordinates given by GPS device, asociated to a point i
// time.
type Position struct {
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
}

// Subject Represents a person or object to be tracked.
type Subject struct {
	CurrentPosition Position
	History         []Position
}
