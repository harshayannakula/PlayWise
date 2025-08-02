package sortengine

import (
	"PlayWise/models"
	"sort"
	"strings"
)

// SortByTitle sorts songs alphabetically by title
func SortByTitle(songs []*models.Song, ascending bool) {
	sort.Slice(songs, func(i, j int) bool {
		if ascending {
			return strings.ToLower(songs[i].Title) < strings.ToLower(songs[j].Title)
		}
		return strings.ToLower(songs[i].Title) > strings.ToLower(songs[j].Title)
	})
}

// SortByDuration sorts songs by duration (int in seconds)
func SortByDuration(songs []*models.Song, ascending bool) {
	sort.Slice(songs, func(i, j int) bool {
		if ascending {
			return songs[i].Duration < songs[j].Duration
		}
		return songs[i].Duration > songs[j].Duration
	})
}

// SortByAddedAt sorts songs by most recently added first or oldest first
func SortByAddedAt(songs []*models.Song, recentFirst bool) {
	sort.Slice(songs, func(i, j int) bool {
		if recentFirst {
			return songs[i].AddedAt.After(songs[j].AddedAt)
		}
		return songs[i].AddedAt.Before(songs[j].AddedAt)
	})
}
