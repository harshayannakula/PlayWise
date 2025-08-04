package snapshot

import (
	"PlayWise/internal/rating"
	"PlayWise/internal/sortengine"
	"PlayWise/models"
)

// Snapshot dashboard data
type Snapshot struct {
	TopLongestSongs   []*models.Song
	RecentlyPlayed    []*models.Song
	SongCountByRating map[int]int
}

// ExportSnapshot integrates other modules to build dashboard
func ExportSnapshot(
	allSongs []*models.Song,
	playbackStack []*models.Song,
	rTree *rating.RatingTree,
) Snapshot {
	// 1. Sort longest songs using sortengine
	sorted := sortengine.SortByDuration(allSongs, false)
	longest := topN(sorted, 5)

	// 2. Recently played stack (reverse order)
	recent := getRecentlyPlayed(playbackStack, 5)

	// 3. Rating counts using tree traversal
	counts := make(map[int]int)
	rating.TraverseAndCount(rTree.Root, counts)

	return Snapshot{
		TopLongestSongs:   longest,
		RecentlyPlayed:    recent,
		SongCountByRating: counts,
	}
}

// topN returns top N elements from a list
func topN(songs []*models.Song, n int) []*models.Song {
	if len(songs) < n {
		return songs
	}
	return songs[:n]
}

// getRecentlyPlayed returns last N from stack
func getRecentlyPlayed(stack []*models.Song, n int) []*models.Song {
	recent := make([]*models.Song, 0, n)
	for i := len(stack) - 1; i >= 0 && len(recent) < n; i-- {
		recent = append(recent, stack[i])
	}
	return recent
}
