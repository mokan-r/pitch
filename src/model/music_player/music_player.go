package music_player

import (
	"fmt"
	"github.com/mokan-r/pitch/model/song"
	"sync"
	"time"
)

type IMusicPlayer interface {
	Play() error
	Pause()
	AddSong(song *song.Song)
	Next()
	Prev() error
}

type MusicPlayer struct {
	IMusicPlayer
	Playlist        *Playlist
	Elapsed         time.Duration
	songEnded       chan bool
	songPaused      chan bool
	addSongMutex    sync.Mutex
	changeSongMutex sync.Mutex
	pauseMutex      sync.Mutex
	Paused          bool
}

func (m *MusicPlayer) Play() error {
	if m.Playlist == nil {
		return fmt.Errorf("playlist is empty")
	}
	if m.Paused {
		m.Paused = false
		go func() {
			time.Sleep(m.Playlist.current.song.Duration - m.Elapsed)
			m.songEnded <- true
		}()
		go func() {
			ticker := time.NewTicker(time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-m.songPaused:
					return
				case <-m.songEnded:
					m.Next()
					m.Play()
				case t := <-ticker.C:
					m.Elapsed = time.Duration(t.Second())
				}
			}
		}()
	}

	return nil
}

func (m *MusicPlayer) Pause() {
	if !m.Paused {
		m.pauseMutex.Lock()
		m.Paused = true
		m.songPaused <- true
		m.pauseMutex.Unlock()
	}
}

func (m *MusicPlayer) AddSong(song *song.Song) {
	m.changeSongMutex.Lock()
	m.Playlist.Append(song)
	m.changeSongMutex.Unlock()
}

func (m *MusicPlayer) Next() {
	m.changeSongMutex.Lock()
	m.Pause()
	if m.Playlist.NextSong() {
		m.Elapsed = 0
	}
	m.changeSongMutex.Unlock()
}

func (m *MusicPlayer) Prev() error {
	m.changeSongMutex.Lock()
	if m.Playlist == nil {
		return fmt.Errorf("Playlist is empty")
	}
	m.Pause()
	m.Playlist.PrevSong()
	m.Elapsed = 0
	m.changeSongMutex.Unlock()
	return nil
}

func New() MusicPlayer {
	return MusicPlayer{
		songEnded:  make(chan bool, 1),
		songPaused: make(chan bool, 1),
		Playlist:   &Playlist{},
		Paused:     true,
	}
}
