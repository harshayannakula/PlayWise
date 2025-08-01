package rating

import (
	"testing"
	"time"

	"PlayWise/models"
)

func newSong(id, title string) *models.Song {
	return &models.Song{
		ID:       id,
		Title:    title,
		Artist:   "Test Artist",
		Duration: 120,
		AddedAt:  time.Now(),
	}
}

func TestInsertAndSearch(t *testing.T) {
	rt := &RatingTree{}

	s1 := newSong("1", "Alpha")
	s2 := newSong("2", "Bravo")
	s3 := newSong("3", "Charlie")

	rt.InsertSong(s1, 3)
	rt.InsertSong(s2, 3)
	rt.InsertSong(s3, 4)

	songs, err := rt.SearchByRating(3)
	if err != nil || len(songs) != 2 {
		t.Errorf("Expected 2 songs with rating 3, got %d", len(songs))
	}

	songs4, err := rt.SearchByRating(4)
	if err != nil || len(songs4) != 1 || songs4[0].ID != "3" {
		t.Errorf("Rating 4 not working")
	}
}

func TestDeleteSong(t *testing.T) {
	rt := &RatingTree{}

	s := newSong("1", "ToDelete")
	rt.InsertSong(s, 5)

	err := rt.DeleteSong("1", 5)
	if err != nil {
		t.Errorf("DeleteSong failed: %v", err)
	}

	songs, _ := rt.SearchByRating(5)
	if len(songs) != 0 {
		t.Errorf("Song not deleted properly")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	rt := &RatingTree{}

	s := newSong("1", "Only One")
	rt.InsertSong(s, 2)

	err := rt.DeleteSong("999", 2)
	if err == nil {
		t.Errorf("Expected error when deleting non-existent song")
	}
}
