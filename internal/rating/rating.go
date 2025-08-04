package rating

import (
	"PlayWise/models"
	"errors"
)

// RatingNode represents a BST node for a given rating (1–5)
type RatingNode struct {
	Rating int
	Songs  []*models.Song
	Left   *RatingNode
	Right  *RatingNode
}

// RatingTree is the BST root holder
type RatingTree struct {
	Root *RatingNode
}

// InsertSong inserts a song into the corresponding rating bucket
func (rt *RatingTree) InsertSong(song *models.Song, rating int) {
	rt.Root = insert(rt.Root, song, rating)
	song.Rating = rating // Optional, to sync song struct
}

func insert(node *RatingNode, song *models.Song, rating int) *RatingNode {
	if node == nil {
		return &RatingNode{
			Rating: rating,
			Songs:  []*models.Song{song},
		}
	}
	if rating < node.Rating {
		node.Left = insert(node.Left, song, rating)
	} else if rating > node.Rating {
		node.Right = insert(node.Right, song, rating)
	} else {
		node.Songs = append(node.Songs, song)
	}
	return node
}

// SearchByRating returns all songs with the given rating
func (rt *RatingTree) SearchByRating(rating int) ([]*models.Song, error) {
	node := rt.Root
	for node != nil {
		if rating < node.Rating {
			node = node.Left
		} else if rating > node.Rating {
			node = node.Right
		} else {
			return node.Songs, nil
		}
	}
	return nil, errors.New("rating not found")
}

// DeleteSong removes a song by ID from the corresponding rating bucket
func (rt *RatingTree) DeleteSong(songID string, rating int) error {
	return deleteByID(rt.Root, songID, rating)
}

func deleteByID(node *RatingNode, songID string, rating int) error {
	if node == nil {
		return errors.New("rating not found")
	}
	if rating < node.Rating {
		return deleteByID(node.Left, songID, rating)
	}
	if rating > node.Rating {
		return deleteByID(node.Right, songID, rating)
	}

	// Found the correct node
	for i, s := range node.Songs {
		if s.ID == songID {
			node.Songs = append(node.Songs[:i], node.Songs[i+1:]...)
			return nil
		}
	}
	return errors.New("song not found")
}

func TraverseAndCount(node *RatingNode, counts map[int]int) {
	if node == nil {
		return
	}
	TraverseAndCount(node.Left, counts)
	counts[node.Rating] += len(node.Songs)
	TraverseAndCount(node.Right, counts)
}
