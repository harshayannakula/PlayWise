package models

import "time"

// Song represents the metadata and runtime attributes of a single song.
type Song struct {
	ID          string
	Title       string
	Artist      string
	Duration    time.Duration
	Rating      int
	VolumeLevel float64
	AddedAt     time.Time
	PlayedAt    time.Time
}
