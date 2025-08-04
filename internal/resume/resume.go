package resume

import (
	"errors"
	"sync"
)

// ResumeManager tracks paused positions per playlist using stacks
type ResumeManager struct {
	mu        sync.RWMutex
	playlists map[string][]int // playlist → stack of positions
}

// NewResumeManager initializes a resume manager
func NewResumeManager() *ResumeManager {
	return &ResumeManager{
		playlists: make(map[string][]int),
	}
}

// Pause pushes the current song index onto the stack for a playlist
func (rm *ResumeManager) Pause(playlistName string, index int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.playlists[playlistName] = append(rm.playlists[playlistName], index)
}

// Resume pops the most recent index for a playlist
func (rm *ResumeManager) Resume(playlistName string) (int, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	stack := rm.playlists[playlistName]
	if len(stack) == 0 {
		return -1, errors.New("no paused position found")
	}

	// pop
	last := stack[len(stack)-1]
	rm.playlists[playlistName] = stack[:len(stack)-1]
	return last, nil
}

// Peek returns the last paused index without popping
func (rm *ResumeManager) Peek(playlistName string) (int, error) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	stack := rm.playlists[playlistName]
	if len(stack) == 0 {
		return -1, errors.New("no paused position")
	}
	return stack[len(stack)-1], nil
}

// ClearStack clears the stack for a playlist
func (rm *ResumeManager) ClearStack(playlistName string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.playlists, playlistName)
}
