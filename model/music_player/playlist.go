package music_player

import (
	"github.com/mokan-r/pitch/model/song"
)

type Playlist struct {
	current *SongLink
	head    *SongLink
	tail    *SongLink
}

type SongLink struct {
	song *song.Song
	next *SongLink
	prev *SongLink
}

func (pl *Playlist) Append(song *song.Song) {
	if pl.current == nil {
		pl.current = &SongLink{
			song: song,
			next: nil,
			prev: nil,
		}
		pl.head = pl.current
		pl.tail = pl.current
	} else {
		pl.tail.next = &SongLink{
			song: song,
			next: nil,
			prev: pl.tail,
		}
		pl.tail = pl.tail.next
	}
}

func (pl *Playlist) NextSong() bool {
	if pl.current.next == nil {
		return false
	}
	pl.current = pl.current.next
	return true
}

func (pl *Playlist) PrevSong() {
	if pl.current.prev != nil {
		pl.current = pl.current.prev
	}
}

func (pl *Playlist) Back() *song.Song {
	return pl.tail.song
}

func (pl *Playlist) Current() *song.Song {
	return pl.current.song
}
