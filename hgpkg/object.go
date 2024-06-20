package hgpkg

import (
	"image"

	"github.com/HiyajiruGames/MiniCity/hgpkg/anim"
	"github.com/HiyajiruGames/MiniCity/hgpkg/resource"
	"github.com/HiyajiruGames/MiniCity/hgpkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

type Objector interface {
	Update(context *Context)
	Draw(screen *ebiten.Image)
}

type Object struct {
	Point     image.Point
	Rectangle image.Rectangle
	Image     *resource.Image
	provider  *anim.AnimationProvider
	animation *ganim8.Animation
}

func NewObject(point image.Point, rectangle image.Rectangle) {
	NewObjectWithFilePath(point, rectangle, nil, 0, 0)
}

func NewObjectWithFilePath(point image.Point, rectangle image.Rectangle, path *string, imageWidth int, imageHeight int) *Object {
	o := new(Object)
	o.Point = point
	o.Rectangle = rectangle
	if !util.IsEmpty(path) {
		o.Image = resource.NewImage(*path)
	}
	o.provider = anim.NewAnimationProvider(nil, ganim8.NewGrid(rectangle.Dx(), rectangle.Dy(), imageWidth, imageHeight))
	return o
}

// Implemented Objector.

func (o *Object) Update(context *Context) {

	if o.provider == nil {
		return
	}

	if o.Image != nil && o.Image.IsLoaded() && !o.provider.HasSomeAnimaion() {
		o.provider.SetImage(o.Image.GetImage())
		o.provider.PreLoad("IDEL", "1-10", 1)
		o.provider.PreLoad("WALK", "1-8", 2)
		o.provider.PreLoad("RUN", "1-8", 3)
		o.animation = o.provider.GetAnimation("RUN")
		o.animation.GoToFrame(1)
	} else if o.provider.HasSomeAnimaion() {
		o.animation.Update()
	}

}

func (o *Object) Draw(screen *ebiten.Image) {
	if o.animation != nil {
		o.animation.Draw(screen, ganim8.DrawOpts(float64(o.Point.X), float64(o.Point.Y), 0, 1, 1))
	} else {
		if o.Image.IsLoaded() {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(o.Point.X), float64(o.Point.Y))
			// ganim8を使ってアニメーションがないものでもフレームを切り出して描画できるようにすると楽かも。
			screen.DrawImage(o.Image.GetImage(), op)
		}
	}
}

// Implemented Mover.

func (o *Object) Up() {
	print("Up")
}

func (o *Object) Down() {
	print("Down")
}

func (o *Object) Left() {
	print("Left")
}

func (o *Object) Right() {
	print("Right")
}

// Implemented Collider.

func (o *Object) IsEntered(p image.Point, r image.Rectangle) bool {
	print("IsEntered")
	return false
}
