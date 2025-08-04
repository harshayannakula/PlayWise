package sortengine

import (
	"PlayWise/models"
	"sort"
	"strings"
)

// SortByTitle sorts songs alphabetically using built-in sort
func SortByTitle(songs []*models.Song, asc bool) {
	sort.SliceStable(songs, func(i, j int) bool {
		if asc {
			return strings.ToLower(songs[i].Title) < strings.ToLower(songs[j].Title)
		}
		return strings.ToLower(songs[i].Title) > strings.ToLower(songs[j].Title)
	})
}

// SortByDuration sorts by song duration using Merge Sort (manual)
func SortByDuration(songs []*models.Song, asc bool) []*models.Song {
	if len(songs) <= 1 {
		return songs
	}
	mid := len(songs) / 2
	left := SortByDuration(songs[:mid], asc)
	right := SortByDuration(songs[mid:], asc)

	return mergeByDuration(left, right, asc)
}

func mergeByDuration(left, right []*models.Song, asc bool) []*models.Song {
	result := make([]*models.Song, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if asc {
			if left[i].Duration <= right[j].Duration {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		} else {
			if left[i].Duration >= right[j].Duration {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// SortByRecentlyAdded sorts by AddedAt (newest first by default)
func SortByRecentlyAdded(songs []*models.Song, newestFirst bool) {
	sort.SliceStable(songs, func(i, j int) bool {
		if newestFirst {
			return songs[i].AddedAt.After(songs[j].AddedAt)
		}
		return songs[i].AddedAt.Before(songs[j].AddedAt)
	})
}
