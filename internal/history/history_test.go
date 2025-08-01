package history

import (
	"testing"
	"time"

	"PlayWise/models"
)

// helper function to create test songs
func newSong(id, title string) *models.Song {
	return &models.Song{
		ID:       id,
		Title:    title,
		Artist:   "Test Artist",
		Duration: 120,
		Rating:   3,
		AddedAt:  time.Now(),
	}
}

func TestPushAndSize(t *testing.T) {
	h := NewHistory()

	h.PushPlayedSong(newSong("1", "A"))
	h.PushPlayedSong(newSong("2", "B"))

	if h.Size() != 2 {
		t.Errorf("Expected size 2, got %d", h.Size())
	}
}

func TestUndoLastPlay(t *testing.T) {
	h := NewHistory()
	s1 := newSong("1", "A")
	s2 := newSong("2", "B")

	h.PushPlayedSong(s1)
	h.PushPlayedSong(s2)

	last, err := h.UndoLastPlay()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if last.ID != "2" {
		t.Errorf("Expected last played song to be B, got %s", last.Title)
	}

	if h.Size() != 1 {
		t.Errorf("Expected size 1 after undo, got %d", h.Size())
	}
}

func TestPeek(t *testing.T) {
	h := NewHistory()
	s := newSong("1", "Peek Song")
	h.PushPlayedSong(s)

	peeked, err := h.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if peeked.ID != "1" {
		t.Errorf("Peeked song mismatch")
	}
	if h.Size() != 1 {
		t.Errorf("Peek should not change size")
	}
}

func TestUndoEmpty(t *testing.T) {
	h := NewHistory()
	_, err := h.UndoLastPlay()
	if err == nil {
		t.Errorf("Expected error when undoing empty stack")
	}
}
