package sortengine

import (
	"testing"
	"time"

	"PlayWise/models"
)

func newSong(id, title string, dur time.Duration, offsetSec int64) *models.Song {
	return &models.Song{
		ID:       id,
		Title:    title,
		Duration: dur, // this is correct
		AddedAt:  time.Now().Add(time.Duration(offsetSec) * time.Second),
	}
}

func TestSortByTitle(t *testing.T) {
	songs := []*models.Song{
		newSong("1", "Zebra", 100, -100),
		newSong("2", "apple", 100, -50),
		newSong("3", "Banana", 100, -10),
	}

	SortByTitle(songs, true)

	if songs[0].Title != "apple" || songs[2].Title != "Zebra" {
		t.Errorf("Title sorting failed")
	}
}

func TestSortByDurationAsc(t *testing.T) {
	songs := []*models.Song{
		newSong("1", "A", 300, 0),
		newSong("2", "B", 100, 0),
		newSong("3", "C", 200, 0),
	}

	result := SortByDuration(songs, true)

	if result[0].Duration != 100 || result[2].Duration != 300 {
		t.Errorf("Duration ascending sort failed")
	}
}

func TestSortByDurationDesc(t *testing.T) {
	songs := []*models.Song{
		newSong("1", "A", 150, 0),
		newSong("2", "B", 450, 0),
		newSong("3", "C", 250, 0),
	}

	result := SortByDuration(songs, false)

	if result[0].Duration != 450 || result[2].Duration != 150 {
		t.Errorf("Duration descending sort failed")
	}
}

func TestSortByRecentlyAdded(t *testing.T) {
	songs := []*models.Song{
		newSong("1", "Old", 100, -1000),
		newSong("2", "Newer", 100, -200),
		newSong("3", "Newest", 100, -10),
	}

	SortByRecentlyAdded(songs, true)

	if songs[0].Title != "Newest" || songs[2].Title != "Old" {
		t.Errorf("Recently added sort failed")
	}
}
