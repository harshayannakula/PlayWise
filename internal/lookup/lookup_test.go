package lookup

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
		Duration: 100,
		AddedAt:  time.Now(),
	}
}

func TestAddAndGetByID(t *testing.T) {
	ls := NewLookupService()
	s := newSong("id1", "Rockstar")
	ls.AddSong(s)

	got, err := ls.GetByID("id1")
	if err != nil || got.Title != "Rockstar" {
		t.Errorf("Failed to fetch song by ID")
	}
}

func TestAddAndGetByTitle(t *testing.T) {
	ls := NewLookupService()
	s := newSong("id2", "Kal Ho Naa Ho")
	ls.AddSong(s)

	got, err := ls.GetByTitle("kal ho naa ho")
	if err != nil || got.ID != "id2" {
		t.Errorf("Failed to fetch song by title (case-insensitive)")
	}
}

func TestRemoveSong(t *testing.T) {
	ls := NewLookupService()
	s := newSong("id3", "DeleteMe")
	ls.AddSong(s)
	ls.RemoveSong(s)

	_, err1 := ls.GetByID("id3")
	_, err2 := ls.GetByTitle("deleteme")
	if err1 == nil || err2 == nil {
		t.Errorf("Song was not properly removed")
	}
}

func TestSize(t *testing.T) {
	ls := NewLookupService()
	ls.AddSong(newSong("a", "A"))
	ls.AddSong(newSong("b", "B"))
	if ls.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ls.Size())
	}
}
