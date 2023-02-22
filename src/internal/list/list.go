package list

import "github.com/mokan-r/pitch/model/song"

type ISongList interface {
	Append(song song.Song)
}
