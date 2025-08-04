package volume

import (
	"PlayWise/models"
	"math"
)

// NormalizeVolume adjusts all song volumes to the average volume level
func NormalizeVolume(songs []*models.Song) float64 {
	if len(songs) == 0 {
		return 0
	}

	var total float64
	for _, s := range songs {
		total += s.VolumeLevel
	}
	avg := total / float64(len(songs))

	for _, s := range songs {
		s.VolumeLevel = roundToTwoDecimal(avg)
	}

	return avg
}

// Helper to round float to 2 decimal places
func roundToTwoDecimal(val float64) float64 {
	return math.Round(val*100) / 100
}
