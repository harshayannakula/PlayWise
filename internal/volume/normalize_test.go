package volume

import (
	"testing"

	"PlayWise/models"
)

func TestNormalizeVolume(t *testing.T) {
	songs := []*models.Song{
		{ID: "1", Title: "A", VolumeLevel: 0.8},
		{ID: "2", Title: "B", VolumeLevel: 1.2},
		{ID: "3", Title: "C", VolumeLevel: 0.6},
	}

	avg := NormalizeVolume(songs)

	if avg < 0.86 || avg > 0.88 {
		t.Errorf("Expected average around 0.87, got %f", avg)
	}

	for _, s := range songs {
		if s.VolumeLevel != roundToTwoDecimal(avg) {
			t.Errorf("Song %s not normalized correctly", s.ID)
		}
	}
}
