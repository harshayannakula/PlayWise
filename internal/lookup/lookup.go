package lookup

import (
	"errors"
	"strings"
	"sync"

	"PlayWise/models"
)

// LookupService provides fast access to song metadata by ID or title.
type LookupService struct {
	sync.RWMutex
	byID    map[string]*models.Song
	byTitle map[string]*models.Song
}

// NewLookupService initializes the hashmap store.
func NewLookupService() *LookupService {
	return &LookupService{
		byID:    make(map[string]*models.Song),
		byTitle: make(map[string]*models.Song),
	}
}

// AddSong indexes the song by both ID and title.
func (ls *LookupService) AddSong(song *models.Song) {
	ls.Lock()
	defer ls.Unlock()
	ls.byID[song.ID] = song
	ls.byTitle[strings.ToLower(song.Title)] = song
}

// RemoveSong deletes a song from both ID and title maps.
func (ls *LookupService) RemoveSong(song *models.Song) {
	ls.Lock()
	defer ls.Unlock()
	delete(ls.byID, song.ID)
	delete(ls.byTitle, strings.ToLower(song.Title))
}

// GetByID fetches song metadata by ID.
func (ls *LookupService) GetByID(id string) (*models.Song, error) {
	ls.RLock()
	defer ls.RUnlock()
	if song, ok := ls.byID[id]; ok {
		return song, nil
	}
	return nil, errors.New("song not found by ID")
}

// GetByTitle fetches song metadata by Title (case-insensitive).
func (ls *LookupService) GetByTitle(title string) (*models.Song, error) {
	ls.RLock()
	defer ls.RUnlock()
	if song, ok := ls.byTitle[strings.ToLower(title)]; ok {
		return song, nil
	}
	return nil, errors.New("song not found by title")
}

// Size returns total indexed songs
func (ls *LookupService) Size() int {
	ls.RLock()
	defer ls.RUnlock()
	return len(ls.byID)
}
