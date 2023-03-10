package music_player

import (
	"github.com/mokan-r/pitch/model/song"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TimeOutOnFunc(t *testing.T, timeout time.Duration, function func() bool) {
	timer := time.NewTimer(timeout)

	for {
		select {
		case <-timer.C:
			t.Errorf("failed on timeout: %s", timeout)
			return
		default:
			if function() {
				return
			}
		}
	}
}

// TestAddSong Test that the MusicPlayer.AddSong() function correctly appends a song to the Playlist.
func TestAddSong(t *testing.T) {
	musicPlayer := New()
	testSong := &song.Song{Id: "1", Title: "TestAddSong song", Duration: time.Second * 3}
	musicPlayer.AddSong(testSong)
	if musicPlayer.Playlist.Back() != testSong {
		t.Errorf("MusicPlayer.AddSong() failed to add the song to the Playlist")
	}
}

// TestPlay Test that the MusicPlayer.Play() function correctly prints out the name of the song playing.
func TestPlay(t *testing.T) {
	musicPlayer := New()
	testSong := &song.Song{Id: "1", Title: "TestPlay song", Duration: time.Second * 2}
	musicPlayer.AddSong(testSong)
	err := musicPlayer.Play()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	if musicPlayer.Playlist.current.song.Title != testSong.Title {
		t.Errorf("MusicPlayer.Play() failed to print the correct name of the song playing")
	}
}

// TestPause Test that the MusicPlayer.Pause() function successfully pauses the song.
func TestPause(t *testing.T) {
	musicPlayer := New()
	testSong := &song.Song{Id: "1", Title: "TestPause song", Duration: time.Second * 2}
	musicPlayer.AddSong(testSong)
	err := musicPlayer.Play()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	musicPlayer.Pause()
	if !musicPlayer.Paused {
		t.Errorf("MusicPlayer.Pause() failed to pause the song")
	}
}

// TestNext Test that the MusicPlayer.Next() function correctly advances to the next song in the Playlist.
func TestNext(t *testing.T) {
	musicPlayer := New()
	testSong1 := &song.Song{Id: "1", Title: "TestNext song 1", Duration: time.Second * 2}
	testSong2 := &song.Song{Id: "2", Title: "TestNext song 2", Duration: time.Second * 3}
	musicPlayer.AddSong(testSong1)
	musicPlayer.AddSong(testSong2)
	err := musicPlayer.Play()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	musicPlayer.Next()
	if musicPlayer.Playlist.current.song.Title != testSong2.Title {
		t.Errorf("MusicPlayer.Next() failed to advance to the next song in the Playlist")
	}
	musicPlayer.Pause()
}

// TestPrev Test that the MusicPlayer.Prev() function correctly returns to the previous song in the Playlist.
func TestPrev(t *testing.T) {
	musicPlayer := New()
	testSong1 := &song.Song{Id: "1", Title: "TestPrev song 1", Duration: time.Second * 3}
	testSong2 := &song.Song{Id: "2", Title: "TestPrev song 2", Duration: time.Second * 2}
	musicPlayer.AddSong(testSong1)
	musicPlayer.AddSong(testSong2)
	err := musicPlayer.Play()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	musicPlayer.Next()
	err = musicPlayer.Prev()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	if musicPlayer.Playlist.current.song.Title != testSong1.Title {
		t.Errorf("MusicPlayer.Prev() failed to return to the previous song in the Playlist")
	}
}

// TestAppend Test that the Playlist.Append() function successfully adds a song to the end of the Playlist.
func TestAppend(t *testing.T) {
	playlist := Playlist{}
	testSong := &song.Song{Id: "1", Title: "TestAppend song", Duration: time.Second * 2}
	testSong1 := &song.Song{Id: "2", Title: "TestAppend song1", Duration: time.Second * 3}
	testSong2 := &song.Song{Id: "3", Title: "TestAppend song2", Duration: time.Second * 4}
	testSong3 := &song.Song{Id: "4", Title: "TestAppend song3", Duration: time.Second * 5}
	testSong4 := &song.Song{Id: "5", Title: "TestAppend song4", Duration: time.Second * 1}
	playlist.Append(testSong)
	playlist.Append(testSong1)
	playlist.Append(testSong2)
	playlist.Append(testSong3)
	playlist.Append(testSong4)
	if playlist.Back().Title != testSong4.Title {
		t.Errorf("Playlist.Append() failed to add the song to the end of the Playlist")
	}
}

// TesSongChanging Test that the after playing a song it changes to the next one
func TestAutoSongChange(t *testing.T) {
	for i := 0; i < 20; i++ {
		musicPlayer := New()

		waitSongChange := func() {
			expected := musicPlayer.Playlist.current.next.song.Title
			TimeOutOnFunc(t, time.Second, func() bool {
				if musicPlayer.Playlist.current.song.Title == expected {
					return true
				}
				return false
			})
		}

		testSong := &song.Song{Id: "1", Title: "song", Duration: time.Millisecond * 20}
		testSong1 := &song.Song{Id: "2", Title: "song1", Duration: time.Millisecond * 20}
		testSong2 := &song.Song{Id: "3", Title: "song2", Duration: time.Millisecond * 20}
		musicPlayer.AddSong(testSong)
		musicPlayer.AddSong(testSong1)
		musicPlayer.AddSong(testSong2)

		err := musicPlayer.Play()
		if err != nil {
			t.Errorf("got error when called play\ngot: %s", err)
			return
		}

		waitSongChange()

		if musicPlayer.Playlist.current.song.Title != testSong1.Title {
			t.Errorf("failed to auto change song to next\ngot: %s\nexpected: %s",
				musicPlayer.Playlist.current.song.Title,
				testSong1.Title)
			return
		}

		waitSongChange()

		if musicPlayer.Playlist.current.song.Title != testSong2.Title {
			t.Errorf("failed to auto change song to next\ngot: %s\nexpected: %s",
				musicPlayer.Playlist.current.song.Title,
				testSong2.Title)
			return
		}
		musicPlayer.Pause()
	}
}

func TestAutoPauseOnPlaylistEnd(t *testing.T) {
	for i := 0; i < 20; i++ {
		musicPlayer := New()
		testSong := &song.Song{Id: "1", Title: "AutoSongChange song", Duration: time.Millisecond}
		musicPlayer.AddSong(testSong)
		err := musicPlayer.Play()
		if err != nil {
			t.Errorf("got error when called play\ngot: %s", err)
			return
		}

		TimeOutOnFunc(t, time.Second, func() bool {
			if musicPlayer.Paused {
				return true
			}
			return false
		})
	}
}

func TestContinue(t *testing.T) {
	musicPlayer := New()
	testSong := &song.Song{Id: "1", Title: "Continue song", Duration: time.Millisecond * 5}
	testSong2 := &song.Song{Id: "2", Title: "Continue song2", Duration: time.Millisecond * 20}
	musicPlayer.AddSong(testSong)
	musicPlayer.AddSong(testSong2)
	err := musicPlayer.Play()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	time.Sleep(testSong.Duration)
	musicPlayer.Pause()
	err = musicPlayer.Play()
	if err != nil {
		t.Errorf("got error when called play\ngot: %s", err)
		return
	}
	time.Sleep(testSong.Duration)
	musicPlayer.Pause()

	if musicPlayer.Playlist.current.song.Title != testSong2.Title {
		t.Errorf("failed to countinue playing song after pause. Elapsed: %s, expected: %s", musicPlayer.Playlist.current.song.Title, testSong2.Title)
	}
}

func TestConcurrentAddSong(t *testing.T) {
	musicPlayer := New()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := strconv.Itoa(i)
			musicPlayer.AddSong(&song.Song{
				Id:       id,
				Title:    "song" + id,
				Duration: time.Second * 3,
			})
		}(i)
	}
	wg.Wait()
	addedMap := make(map[string]struct{}, 100)
	cur := musicPlayer.Playlist.head
	for cur != nil {
		addedMap["song"+cur.song.Id] = struct{}{}
		cur = cur.next
	}
	for i := 0; i < 100; i++ {
		songName := "song" + strconv.Itoa(i)
		if _, ok := addedMap[songName]; !ok {
			t.Errorf("failed to concurrently add song: %s", songName)
			return
		}
	}
}

func TestConcurrentNext(t *testing.T) {
	for i := 0; i < 20; i++ {
		musicPlayer := New()
		songsQty := 30

		for i := 0; i < songsQty; i++ {
			musicPlayer.AddSong(&song.Song{Id: "1", Title: "Next song" + strconv.Itoa(i)})
		}

		testSong := &song.Song{Title: "Final song", Duration: time.Second * 3}
		musicPlayer.AddSong(testSong)

		w := sync.WaitGroup{}

		for i := 0; i <= songsQty; i++ {
			w.Add(1)
			go func() {
				defer w.Done()
				musicPlayer.Next()
			}()
		}
		w.Wait()

		if musicPlayer.Playlist.current.song.Title != testSong.Title {
			t.Errorf("failed to concurently change song next:\ngot %s, expected: %s", musicPlayer.Playlist.current.song.Title, testSong.Title)
		}
	}
}

func TestConcurrentPrev(t *testing.T) {
	for i := 0; i < 20; i++ {
		musicPlayer := New()
		songsQty := 30

		for i := 0; i < songsQty; i++ {
			id := strconv.Itoa(i)
			musicPlayer.AddSong(&song.Song{Id: id, Title: "Next song" + id})
		}

		testSong := &song.Song{Title: "Final song", Duration: time.Second}
		musicPlayer.AddSong(testSong)

		for i := 0; i <= songsQty; i++ {
			musicPlayer.Next()
		}

		w := sync.WaitGroup{}

		for i := 0; i <= songsQty; i++ {
			w.Add(1)
			go func() {
				defer w.Done()
				err := musicPlayer.Prev()
				if err != nil {
					t.Errorf("failed to change song to prev: %s", err)
					return
				}
			}()
		}
		w.Wait()

		if musicPlayer.Playlist.current.song.Title != "Next song0" {
			t.Errorf("failed to concurently change song next:\ngot %s, expected: %s", musicPlayer.Playlist.current.song.Title, "Next song0")
		}
	}
}

func TestConcurrentChangeSong(t *testing.T) {
	for i := 0; i < 20; i++ {
		musicPlayer := New()
		songsQty := 30

		for i := 0; i < songsQty; i++ {
			musicPlayer.AddSong(&song.Song{Title: "Next song" + strconv.Itoa(i)})
		}

		for i := 0; i <= songsQty/2; i++ {
			musicPlayer.Next()
		}

		w := sync.WaitGroup{}

		for i := 0; i <= songsQty; i++ {
			w.Add(2)
			go func() {
				defer w.Done()
				err := musicPlayer.Prev()
				if err != nil {
					t.Errorf("failed to change song to prev: %s", err)
					return
				}
			}()
			go func() {
				defer w.Done()
				musicPlayer.Next()
			}()
		}
		w.Wait()

		if musicPlayer.Playlist.current.song.Title != "Next song16" {
			t.Errorf("failed to concurently change song:\ngot %s, expected: %s",
				musicPlayer.Playlist.current.song.Title, "Next song16")
			return
		}
	}
}
