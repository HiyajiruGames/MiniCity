package module

import (
	"github.com/HiyajiruGames/MiniCity/hgpkg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Moduler interface {
	Update(context *hgpkg.Context)
	Draw(screen *ebiten.Image)
}
