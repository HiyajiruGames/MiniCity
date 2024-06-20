package resource

import (
	"bytes"
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

type MusicType int

const (
	Ogg MusicType = iota
	MP3
	Other
)

type audioStream interface {
	io.ReadSeeker
	Length() int64
}

type Audio struct {
	musicType MusicType
	stream    audioStream
	Resource
}

func NewAudio(path string, musicType MusicType) *Audio {
	a := new(Audio)
	a.Resource = *NewResource(path)
	a.musicType = musicType
	return a
}

func (i *Audio) Load() {
	i.Resource.Load()
	reader := bytes.NewReader(i.data)
	var s audioStream
	var err error
	switch i.musicType {
	case Ogg:
		s, err = vorbis.DecodeWithoutResampling(reader)
	case MP3:
		s, err = mp3.DecodeWithoutResampling(reader)
	default:
		panic("not supported.")
	}
	if err != nil {
		panic(fmt.Sprintf("load error...reason is %s", err.Error()))
	}
	i.stream = s
}

func (i *Audio) Release() {
	i.stream = nil
}
