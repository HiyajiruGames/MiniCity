package main

import (
	_ "image/png"

	"github.com/HiyajiruGames/MiniCity/app"
	"github.com/HiyajiruGames/MiniCity/hgpkg"
	"github.com/HiyajiruGames/MiniCity/hgpkg/log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {

	// Initialize a Game with Context.
	c := hgpkg.NewContext(screenWidth/2, screenHeight/2)
	g := hgpkg.NewGameWithContext(c, app.NewGame())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("MiniCity(Ebitengine Game Jam Version)")
	if err := ebiten.RunGame(g); err != nil {
		log.Error("", err)
	}
}
