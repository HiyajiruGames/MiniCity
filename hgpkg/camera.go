package hgpkg

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ChaseType int

const (
	Linear ChaseType = iota
	EaseIn
	EaseOut
	Ease
)

type Camera struct {
	Surface         *ebiten.Image
	backgroundColor color.Color
	centerPosition  image.Point
	targetObject    *Object
	targetPosition  image.Point
	chaseType       ChaseType
	isChasing       bool
}

func NewCamera(width int, height int) *Camera {
	c := new(Camera)
	c.Surface = ebiten.NewImage(width, height)
	c.SetBackground(color.RGBA{0, 0, 0, 255})
	c.chaseType = Linear
	return c
}

func (c *Camera) PreDraw(Context *Context) {
	c.Surface.Clear()
	c.Surface.Fill(c.backgroundColor)
}

func (c *Camera) Draw(context *Context, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// 画面の中心座標とキャラクター位置の差を求める。
	// 画面の中心座標にキャラクターを近づけるためにSurface位置を調整。
	cx := context.GetScreenWidth() / 2
	cy := context.GetScreenHeight() / 2
	dx := cx - c.targetPosition.X
	dy := cy - c.targetPosition.Y
	op.GeoM.Translate(float64(dx), float64(dy))
	screen.DrawImage(c.Surface, op)
}

func (c *Camera) Update(context *Context) {

	if c.isChasing {

		if c.targetObject != nil {
			c.SetTargetPosition(image.Point{c.targetObject.Point.X + c.targetObject.Rectangle.Dx()/2, c.targetObject.Point.Y + c.targetObject.Rectangle.Dy()/2})
		}

		var a = c.centerPosition
		var b = c.targetPosition

		switch c.chaseType {
		case Linear:
			// a = b
			c.centerPosition = b
		case EaseIn:
		case EaseOut:
		case Ease:
			// a += (b - a) / 2
			c.centerPosition = a.Add(b.Sub(a).Div(2))
		}

	} else {
		c.centerPosition = c.targetPosition
	}

}

func (c *Camera) SetBackground(color color.Color) {
	c.backgroundColor = color
}

func (c *Camera) SetTargetPosition(p image.Point) {
	c.targetPosition = p
	c.isChasing = true
}

func (c *Camera) SetTarget(object *Object) {
	c.targetObject = object
	c.isChasing = true
}

func (c *Camera) SetChaseType(chaseType ChaseType) {
	c.chaseType = chaseType
}
