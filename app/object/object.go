package object

import (
	"image"

	"github.com/HiyajiruGames/MiniCity/app/module"
	"github.com/HiyajiruGames/MiniCity/hgpkg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Objector interface {
	Update(context *hgpkg.Context)
	Draw(screen *ebiten.Image)
}

type Object struct {
	point   image.Point
	modules []module.Moduler
}

func NewObject(point image.Point) *Object {
	o := new(Object)
	o.point = point
	o.modules = make([]module.Moduler, 0, 999) // Capacityは、動的に変えたい。
	return o
}

func (o *Object) AddModule(moduler module.Moduler) {
	if moduler != nil {
		o.modules = append(o.modules, moduler)
	}
}

// Implemented Objector.

func (o *Object) Update(context *hgpkg.Context) {
	for _, v := range o.modules {
		v.Update(context)
	}
}

func (o *Object) Draw(screen *ebiten.Image) {
	for _, v := range o.modules {
		v.Draw(screen)
	}
}
