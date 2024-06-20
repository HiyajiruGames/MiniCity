package object

import "github.com/HiyajiruGames/MiniCity/app/module"

type Player struct {
	Object
}

func NewPlayer() *Player {
	p := new(Player)
	p.AddModule(&module.Talk{})
	p.AddModule(&module.Test{})
	return p
}
