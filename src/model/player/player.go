package player

import (
	"github.com/mokan-r/pitch/internal/list"
	"github.com/mokan-r/pitch/model/song"
	"time"
)

type IMusicPlayer interface {
	Play()
	Pause()
	AddSong(song song.Song)
	Next()
	Prev()
}

type Playlist struct {
	list.ISongList
	song         song.Song
	nextSong     *Playlist
	previousSong *Playlist
}

func (sl *Playlist) Append(song song.Song) {
	last := sl.nextSong
	for sl.nextSong != nil {
		last = last.nextSong
	}
	last.nextSong = &Playlist{
		song:         song,
		nextSong:     nil,
		previousSong: last,
	}
}

type MusicPlayer struct {
	IMusicPlayer
	playlist *Playlist
	ticker   *time.Ticker
	status   chan bool
}

func (m *MusicPlayer) Play() {
	go func(status chan bool) {
		time.Sleep(m.playlist.song.Duration)
		status <- true
	}(m.status)
}

func (m *MusicPlayer) Pause() {
	// pause
}

func (m *MusicPlayer) AddSong(song song.Song) {
	m.playlist.Append(song)
}

func (m *MusicPlayer) Next() {
	m.playlist = m.playlist.nextSong
}

func (m *MusicPlayer) Prev() {
	m.playlist = m.playlist.previousSong
}

func New() MusicPlayer {
	return MusicPlayer{}
}
