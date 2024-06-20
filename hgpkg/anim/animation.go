package anim

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

type Animator interface {
	Update()
	Draw(screen *ebiten.Image)
}

type AnimationProvider struct {
	image      *ebiten.Image
	grid       *ganim8.Grid
	animations map[string]*ganim8.Animation
}

func NewAnimationProvider(image *ebiten.Image, grid *ganim8.Grid) *AnimationProvider {
	ap := new(AnimationProvider)
	ap.image = image
	ap.grid = grid
	ap.animations = map[string]*ganim8.Animation{}
	return ap
}

func (ap *AnimationProvider) SetImage(image *ebiten.Image) {
	ap.image = image
}

func (ap *AnimationProvider) PreLoad(name string, interval string, row int) {
	ap.animations[name] = ganim8.New(ap.image, ap.grid.Frames(interval, row), 100*time.Millisecond)
}

func (ap *AnimationProvider) GetAnimation(name string) *ganim8.Animation {
	return ap.animations[name]
}

func (ap *AnimationProvider) HasSomeAnimaion() bool {
	return len(ap.animations) > 0
}
