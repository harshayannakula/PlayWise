package snapshot

import (
	"testing"
	"time"

	"PlayWise/internal/rating"
	//"PlayWise/internal/sortengine"
	"PlayWise/models"
)

func newSong(id string, dur time.Duration, rating int, addedOffset, playedOffset int64) *models.Song {
	return &models.Song{
		ID:       id,
		Title:    "Song_" + id,
		Duration: dur,
		Rating:   rating,
		AddedAt:  time.Now().Add(time.Duration(addedOffset) * time.Second),
		PlayedAt: time.Now().Add(time.Duration(playedOffset) * time.Second),
	}
}

func TestExportSnapshot_Integration(t *testing.T) {
	// ✅ Setup songs
	songs := []*models.Song{
		newSong("1", 200, 3, -300, -10),
		newSong("2", 400, 4, -200, -9),
		newSong("3", 500, 5, -100, -8),
		newSong("4", 350, 3, -90, -7),
		newSong("5", 600, 5, -80, -6),
		newSong("6", 150, 2, -70, -5),
	}

	// ✅ Setup playback stack (last played = song 6 → song 2)
	playbackStack := []*models.Song{
		songs[0],
		songs[1],
		songs[2],
		songs[3],
		songs[4],
		songs[5],
	}

	// ✅ Setup RatingTree
	tree := &rating.RatingTree{}
	for _, s := range songs {
		tree.InsertSong(s, s.Rating)
	}

	// ✅ Call Snapshot
	result := ExportSnapshot(songs, playbackStack, tree)

	// 🔍 Test 1: Longest 5 durations
	if len(result.TopLongestSongs) != 5 {
		t.Errorf("Expected 5 longest songs, got %d", len(result.TopLongestSongs))
	}
	if result.TopLongestSongs[0].Duration < result.TopLongestSongs[1].Duration {
		t.Errorf("Top songs not sorted descending by duration")
	}

	// 🔍 Test 2: Recently played (last 5 of stack)
	if len(result.RecentlyPlayed) != 5 {
		t.Errorf("Expected 5 recently played songs, got %d", len(result.RecentlyPlayed))
	}

	if result.RecentlyPlayed[0].ID != "6" || result.RecentlyPlayed[4].ID != "2" {
		t.Errorf("Recently played order incorrect: got %s to %s",
			result.RecentlyPlayed[0].ID, result.RecentlyPlayed[4].ID)
	}

	// 🔍 Test 3: Rating counts
	expected := map[int]int{2: 1, 3: 2, 4: 1, 5: 2}
	for rating, count := range expected {
		if result.SongCountByRating[rating] != count {
			t.Errorf("Expected %d songs with rating %d, got %d",
				count, rating, result.SongCountByRating[rating])
		}
	}
}
