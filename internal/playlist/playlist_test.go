package playlist

import (
	"testing"
)

func TestAddSong(t *testing.T) {
	p := NewPlaylist("TestAdd")

	p.AddSong("Song One", "chapri", 180)

	if p.Size != 1 {
		t.Errorf("Expected size 1, got %d", p.Size)
	}
	if p.Head == nil || p.Head.Song.Title != "Song One" {
		t.Errorf("Head song not added correctly")
	}
}

func TestDeleteSong(t *testing.T) {
	p := NewPlaylist("TestDelete")

	p.AddSong("Song One", "chapri", 180)
	p.AddSong("Song Two", "charith", 200)
	p.AddSong("Song Three", "colour", 150)

	err := p.DeleteSong(1)
	if err != nil {
		t.Errorf("DeleteSong failed: %v", err)
	}

	if p.Size != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", p.Size)
	}

	if p.Head.Next.Song.Title != "Song Three" {
		t.Errorf("Wrong song remained after deletion")
	}
}

func TestMoveSong(t *testing.T) {
	p := NewPlaylist("TestMove")

	p.AddSong("A", "a", 180)
	p.AddSong("B", "b", 190)
	p.AddSong("C", "c", 200)

	err := p.MoveSong(0, 2)
	if err != nil {
		t.Errorf("MoveSong failed: %v", err)
	}

	if p.Tail.Song.Title != "A" {
		t.Errorf("Song not moved to the correct position")
	}
}

func TestReversePlaylist(t *testing.T) {
	p := NewPlaylist("TestReverse")

	p.AddSong("One", "a", 100)
	p.AddSong("Two", "b", 200)
	p.AddSong("Three", "c", 300)

	p.Reverse()

	if p.Head.Song.Title != "Three" || p.Tail.Song.Title != "One" {
		t.Errorf("Playlist was not reversed properly")
	}
}
