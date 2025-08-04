package history

import (
	"PlayWise/models"
	"errors"
)

// PlaybackHistory represents a LIFO stack of played songs
type PlaybackHistory struct {
	stack []*models.Song
}

// NewHistory initializes a new empty history stack
func NewHistory() *PlaybackHistory {
	return &PlaybackHistory{}
}

// PushPlayedSong adds a song to the playback history stack
func (h *PlaybackHistory) PushPlayedSong(song *models.Song) {
	h.stack = append(h.stack, song)
}

// UndoLastPlay pops the last played song and returns it to re-add to playlist
func (h *PlaybackHistory) UndoLastPlay() (*models.Song, error) {
	if len(h.stack) == 0 {
		return nil, errors.New("no song to undo")
	}
	// Pop from stack
	last := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return last, nil
}

// Size returns the number of songs in the playback history
func (h *PlaybackHistory) Size() int {
	return len(h.stack)
}

// Peek returns the last played song without removing it
func (h *PlaybackHistory) Peek() (*models.Song, error) {
	if len(h.stack) == 0 {
		return nil, errors.New("history is empty")
	}
	return h.stack[len(h.stack)-1], nil
}

// GetStack returns the current playback history stack (read-only)
func (h *PlaybackHistory) GetStack() []*models.Song {
	return h.stack
}
