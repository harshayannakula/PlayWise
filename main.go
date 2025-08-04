package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"PlayWise/internal/history"
	"PlayWise/internal/lookup"
	"PlayWise/internal/playlist"
	"PlayWise/internal/rating"
	"PlayWise/internal/resume"
	"PlayWise/internal/snapshot"
	//"PlayWise/internal/sortengine"
	"PlayWise/internal/volume"
	"PlayWise/models"
)

var (
	pl         = playlist.NewPlaylist("MyPlaylist")
	historyMgr = history.NewHistory()
	lookupMgr  = lookup.NewLookupService()
	ratingTree = &rating.RatingTree{}
	resumeMgr  = resume.NewResumeManager()
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🎵 Welcome to PlayWise CLI")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Song")
		fmt.Println("2. Delete Song")
		fmt.Println("3. Move Song")
		fmt.Println("4. Reverse Playlist")
		fmt.Println("5. Play Song (simulate)")
		fmt.Println("6. Undo Last Play")
		fmt.Println("7. Normalize Volume")
		fmt.Println("8. Snapshot Dashboard")
		fmt.Println("9. Pause Playlist")
		fmt.Println("10. Resume Playlist")
		fmt.Println("11. View Playlist")
		fmt.Println("0. Exit")
		fmt.Print("> ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter ID, Title, Artist, Duration(sec), Rating(1-5), Volume: ")
			scanner.Scan()
			parts := strings.Split(scanner.Text(), ",")
			if len(parts) != 6 {
				fmt.Println("Invalid input format")
				continue
			}
			dur, _ := strconv.Atoi(strings.TrimSpace(parts[3]))
			rating, _ := strconv.Atoi(strings.TrimSpace(parts[4]))
			vol, _ := strconv.ParseFloat(strings.TrimSpace(parts[5]), 64)
			dur1 := time.Duration(dur) * time.Second
			song := &models.Song{
				ID:          strings.TrimSpace(parts[0]),
				Title:       strings.TrimSpace(parts[1]),
				Artist:      strings.TrimSpace(parts[2]),
				Duration:    dur1,
				Rating:      rating,
				VolumeLevel: vol,
			}

			pl.Addsong(song)
			lookupMgr.AddSong(song)
			ratingTree.InsertSong(song, rating)
			fmt.Println("✅ Song added")

		case "2":
			fmt.Print("Enter song index to delete: ")
			scanner.Scan()
			idx, _ := strconv.Atoi(scanner.Text())
			song := getSongAtIndex(idx)
			if song != nil {
				lookupMgr.RemoveSong(song)
				ratingTree.DeleteSong(song.ID, song.Rating)
				_ = pl.DeleteSong(idx)
				fmt.Println("✅ Song deleted")
			}

		case "3":
			fmt.Print("From Index, To Index: ")
			scanner.Scan()
			parts := strings.Split(scanner.Text(), ",")
			from, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			to, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			_ = pl.MoveSong(from, to)
			fmt.Println("✅ Song moved")

		case "4":
			pl.Reverse()
			fmt.Println("✅ Playlist reversed")

		case "5":
			fmt.Print("Enter song index to play: ")
			scanner.Scan()
			idx, _ := strconv.Atoi(scanner.Text())
			song := getSongAtIndex(idx)
			if song != nil {
				historyMgr.PushPlayedSong(song)
				fmt.Printf("🎶 Now Playing: %s\n", song.Title)
			}

		case "6":
			song, err := historyMgr.UndoLastPlay()
			if err != nil {
				fmt.Println("Nothing to undo")
				continue
			}
			pl.Addsong(song)
			fmt.Printf("🔁 Re-added to playlist: %s\n", song.Title)

		case "7":
			avg := volume.NormalizeVolume(getAllSongs())
			fmt.Printf("✅ Normalized all volumes to %.2f\n", avg)

		case "8":
			snapshot := snapshot.ExportSnapshot(getAllSongs(), historyMgr.GetStack(), ratingTree)
			fmt.Println("📊 Snapshot Dashboard")
			for _, s := range snapshot.TopLongestSongs {
				fmt.Printf("🔊 %s (%ds)\n", s.Title, s.Duration)
			}
			fmt.Println("Recently Played:")
			for _, s := range snapshot.RecentlyPlayed {
				fmt.Printf("🎵 %s\n", s.Title)
			}
			fmt.Println("Song Count by Rating:")
			for rating, count := range snapshot.SongCountByRating {
				fmt.Printf("⭐ %d: %d songs\n", rating, count)
			}

		case "9":
			fmt.Print("Enter index to pause: ")
			scanner.Scan()
			idx, _ := strconv.Atoi(scanner.Text())
			resumeMgr.Pause("MyPlaylist", idx)
			fmt.Println("⏸️ Playlist paused")

		case "10":
			idx, err := resumeMgr.Resume("MyPlaylist")
			if err != nil {
				fmt.Println("No paused state found")
			} else {
				song := getSongAtIndex(idx)
				fmt.Printf("▶️ Resuming from index %d: %s\n", idx, song.Title)
			}

		case "11":
			pl.Print()

		case "0":
			fmt.Println("👋 Exiting PlayWise. Bye!")
			return
		default:
			fmt.Println("❌ Invalid option")
		}
	}
}

// Utility to get song at a specific index from the playlist
func getSongAtIndex(index int) *models.Song {
	current := pl.Head
	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}
	if current != nil {
		return current.Song
	}
	fmt.Println("⚠️ Song not found at index")
	return nil
}

// Utility to collect all songs from playlist into a slice
func getAllSongs() []*models.Song {
	songs := []*models.Song{}
	current := pl.Head
	for current != nil {
		songs = append(songs, current.Song)
		current = current.Next
	}
	return songs
}
