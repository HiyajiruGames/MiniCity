package app

import (
	"image"
	"image/color"

	"github.com/HiyajiruGames/MiniCity/app/object"
	"github.com/HiyajiruGames/MiniCity/hgpkg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *object.Player
	tile01 *hgpkg.Object
	test   *hgpkg.Object
	camera *hgpkg.Camera
}

func NewGame() *Game {
	g := new(Game)
	return g
}

func (g *Game) Init(context *hgpkg.Context) {

	// FIXME: プリミティブ型は、ポインタを渡したほうが良いか？実体を渡すべきか？
	var path = "resources/basic_human.png"
	g.test = hgpkg.NewObjectWithFilePath(image.Point{0, 0}, image.Rect(0, 0, int(48), int(48)), &path, 480, 240)

	var tilemap = "resources/tilemap.png"
	g.tile01 = hgpkg.NewObjectWithFilePath(image.Point{0, 0}, image.Rect(0, 0, int(32), int(16)), &tilemap, 192, 80)

	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(s)
	// g.test.Point.X = r.Intn(context.GetScreenWidth())
	// g.test.Point.Y = r.Intn(context.GetScreenHeight())

	//path = "resources/bgm_sample.ogg"
	//g.audio = resource.NewAudio(path, resource.Ogg)

	context.GetLoader().Load(g.test.Image)
	context.GetLoader().Load(g.tile01.Image)
	//g.loader.Load(g.audio)

	g.camera = hgpkg.NewCamera(context.GetScreenWidth(), context.GetScreenHeight())
	g.camera.SetTarget(g.tile01)
	g.camera.SetChaseType(hgpkg.Ease)
	g.camera.SetBackground(color.RGBA{0, 0, 0, 255})

	g.player = object.NewPlayer()
}

func (g *Game) Update(context *hgpkg.Context) {
	//g.test.Update(context)

	/*
		if util.IsPressed(context.GetEnteredKeys(), ebiten.KeyArrowUp) {
			g.test.Point.Y -= 1
		}
		if util.IsPressed(context.GetEnteredKeys(), ebiten.KeyArrowDown) {
			g.test.Point.Y += 1
		}
		if util.IsPressed(context.GetEnteredKeys(), ebiten.KeyArrowLeft) {
			g.test.Point.X -= 1
		}
		if util.IsPressed(context.GetEnteredKeys(), ebiten.KeyArrowRight) {
			g.test.Point.X += 1
		}
	*/

	g.player.Update(context)

	g.camera.Update(context)
}

func (g *Game) Pause(context *hgpkg.Context) {
}

func (g *Game) Finalize(context *hgpkg.Context) {
}

func (g *Game) Draw(context *hgpkg.Context, screen *ebiten.Image) {
	g.camera.PreDraw(context)

	//g.test.Draw(g.camera.Surface)
	g.tile01.Draw(g.camera.Surface)

	// if g.test.Image != nil && g.test.Image.IsLoaded() {
	// 	op := &ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64(g.test.Point.X), float64(g.test.Point.Y))
	// 	screen.DrawImage(g.test.Image.GetImage(), op)
	// }

	// if g.audio != nil && g.audio.IsLoaded() {
	// 	*audio.Player
	// }

	g.player.Draw(screen)

	g.camera.Draw(context, screen)

}
