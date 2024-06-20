package module

import (
	"github.com/HiyajiruGames/MiniCity/hgpkg"
	"github.com/HiyajiruGames/MiniCity/hgpkg/log"
	"github.com/hajimehoshi/ebiten/v2"
)

type Test struct {
}

func (t *Test) Update(context *hgpkg.Context) {
	log.Debug("Call Update of Test!!!")
}

func (t *Test) Draw(screen *ebiten.Image) {
	log.Debug("Call Draw of Test!!!")
}
