package module

import (
	"github.com/HiyajiruGames/MiniCity/hgpkg"
	"github.com/HiyajiruGames/MiniCity/hgpkg/log"
	"github.com/hajimehoshi/ebiten/v2"
)

type Talk struct {
}

func (t *Talk) Update(context *hgpkg.Context) {
	log.Debug("Call Update of Talk!!!")
}

func (t *Talk) Draw(screen *ebiten.Image) {
	log.Debug("Call Draw of Talk!!!")
}
