package hgpkg

import (
	"github.com/HiyajiruGames/MiniCity/hgpkg/control"
	"github.com/HiyajiruGames/MiniCity/hgpkg/resource"
)

type Context struct {
	screenHeight int
	screenWidth  int
	keys         []*control.Key
	loader       *resource.Loader
}

func NewContext(screenWidth int, screenHeight int) *Context {
	g := new(Context)
	g.screenWidth = screenWidth
	g.screenHeight = screenHeight
	g.keys = make([]*control.Key, 0)
	g.loader = resource.NewLoader()
	return g
}

func (c *Context) GetScreenHeight() int {
	return c.screenHeight
}

func (c *Context) GetScreenWidth() int {
	return c.screenWidth
}

func (c *Context) GetEnteredKeys() []*control.Key {
	return c.keys
}

func (c *Context) SetEnteredKeys(keys []*control.Key) {
	c.keys = keys
}

func (c *Context) GetLoader() *resource.Loader {
	return c.loader
}
