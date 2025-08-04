package playlist

import (
	"errors"
	"fmt"
	"time"

	"PlayWise/models"
)

// SongNode represents a node in the doubly linked list for the playlist
type SongNode struct {
	Song *models.Song
	Prev *SongNode
	Next *SongNode
}

// Playlist represents a playlist containing songs
// It uses a doubly linked list to store the songs, allowing for efficient insertion and deletion
type Playlist struct {
	Name string
	Head *SongNode
	Tail *SongNode
	Size int
}

func NewPlaylist(name string) *Playlist {
	return &Playlist{Name: name}
}

// AddSong adds a new song to the end of the playlist
func (p *Playlist) AddSong(title string, artist string, duration time.Duration) {
	song := &models.Song{
		ID:       models.UniqueID(),
		Title:    title,
		Artist:   artist,
		Duration: duration,
	}
	song.AddedAt = time.Now()
	newNode := &SongNode{Song: song}
	if p.Head == nil {
		p.Head = newNode
		p.Tail = newNode
	} else {
		p.Tail.Next = newNode
		newNode.Prev = p.Tail
		p.Tail = newNode
	}
	p.Size++
}

// Addsong adds an existing song to the playlist
func (p *Playlist) Addsong(song *models.Song) {
	song.ID = models.UniqueID()
	song.AddedAt = time.Now()
	newNode := &SongNode{Song: song}
	if p.Head == nil {
		p.Head = newNode
		p.Tail = newNode
	} else {
		p.Tail.Next = newNode
		newNode.Prev = p.Tail
		p.Tail = newNode
	}
	p.Size++
}

// DeleteSong removes a song from the playlist by its index
// It returns an error if the index is invalid
func (p *Playlist) DeleteSong(index int) error {
	if index < 0 || index >= p.Size {
		return errors.New("invalid index")
	}

	current := p.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	if current.Prev != nil {
		current.Prev.Next = current.Next
	} else {
		p.Head = current.Next
	}

	if current.Next != nil {
		current.Next.Prev = current.Prev
	} else {
		p.Tail = current.Prev
	}

	p.Size--
	return nil
}

// MoveSong moves a song from one index to another in the playlist
func (p *Playlist) MoveSong(fromIndex, toIndex int) error {
	if fromIndex < 0 || fromIndex >= p.Size || toIndex < 0 || toIndex >= p.Size {
		return errors.New("invalid index")
	}
	if fromIndex == toIndex {
		return nil
	}

	var node *SongNode
	current := p.Head
	for i := 0; i < fromIndex; i++ {
		current = current.Next
	}
	node = current

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		p.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		p.Tail = node.Prev
	}

	current = p.Head
	for i := 0; i < toIndex; i++ {
		current = current.Next
	}

	if toIndex == 0 {
		node.Next = p.Head
		node.Prev = nil
		p.Head.Prev = node
		p.Head = node
	} else if toIndex == p.Size-1 {
		node.Prev = p.Tail
		node.Next = nil
		p.Tail.Next = node
		p.Tail = node
	} else {
		node.Next = current
		node.Prev = current.Prev
		if current.Prev != nil {
			current.Prev.Next = node
		}
		current.Prev = node
	}

	return nil
}

// Reverse reverses the order of songs in the playlist
func (p *Playlist) Reverse() {
	current := p.Head
	var prev *SongNode
	p.Tail = p.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}
	p.Head = prev
}

func (p *Playlist) Print() {
	current := p.Head
	fmt.Printf("🎵 Playlist: %s\n", p.Name)
	for current != nil {
		s := current.Song
		fmt.Printf(" - [%s] %s (%d sec)\n", s.ID, s.Title, s.Duration)
		current = current.Next
	}
	fmt.Println()
}
