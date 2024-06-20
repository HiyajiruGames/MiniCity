package hgpkg

import (
	"strconv"
	"time"

	"github.com/HiyajiruGames/MiniCity/hgpkg/control"
	"github.com/HiyajiruGames/MiniCity/hgpkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/ebitick"
)

var supportKeys = [...]ebiten.Key{
	ebiten.KeyA,
	ebiten.KeyB,
	ebiten.KeyC,
	ebiten.KeyD,
	ebiten.KeyE,
	ebiten.KeyF,
	ebiten.KeyG,
	ebiten.KeyH,
	ebiten.KeyI,
	ebiten.KeyJ,
	ebiten.KeyK,
	ebiten.KeyL,
	ebiten.KeyM,
	ebiten.KeyN,
	ebiten.KeyO,
	ebiten.KeyP,
	ebiten.KeyQ,
	ebiten.KeyR,
	ebiten.KeyS,
	ebiten.KeyT,
	ebiten.KeyU,
	ebiten.KeyV,
	ebiten.KeyW,
	ebiten.KeyX,
	ebiten.KeyY,
	ebiten.KeyZ,
	ebiten.KeyAltLeft,
	ebiten.KeyAltRight,
	ebiten.KeyArrowDown,
	ebiten.KeyArrowLeft,
	ebiten.KeyArrowRight,
	ebiten.KeyArrowUp,
	ebiten.KeyBackquote,
	ebiten.KeyBackslash,
	ebiten.KeyBackspace,
	ebiten.KeyBracketLeft,
	ebiten.KeyBracketRight,
	ebiten.KeyCapsLock,
	ebiten.KeyComma,
	ebiten.KeyContextMenu,
	ebiten.KeyControlLeft,
	ebiten.KeyControlRight,
	ebiten.KeyDelete,
	ebiten.KeyDigit0,
	ebiten.KeyDigit1,
	ebiten.KeyDigit2,
	ebiten.KeyDigit3,
	ebiten.KeyDigit4,
	ebiten.KeyDigit5,
	ebiten.KeyDigit6,
	ebiten.KeyDigit7,
	ebiten.KeyDigit8,
	ebiten.KeyDigit9,
	ebiten.KeyEnd,
	ebiten.KeyEnter,
	ebiten.KeyEqual,
	ebiten.KeyEscape,
	ebiten.KeyF1,
	ebiten.KeyF2,
	ebiten.KeyF3,
	ebiten.KeyF4,
	ebiten.KeyF5,
	ebiten.KeyF6,
	ebiten.KeyF7,
	ebiten.KeyF8,
	ebiten.KeyF9,
	ebiten.KeyF10,
	ebiten.KeyF11,
	ebiten.KeyF12,
	ebiten.KeyHome,
	ebiten.KeyInsert,
	ebiten.KeyMetaLeft,
	ebiten.KeyMetaRight,
	ebiten.KeyMinus,
	ebiten.KeyNumLock,
	ebiten.KeyNumpad0,
	ebiten.KeyNumpad1,
	ebiten.KeyNumpad2,
	ebiten.KeyNumpad3,
	ebiten.KeyNumpad4,
	ebiten.KeyNumpad5,
	ebiten.KeyNumpad6,
	ebiten.KeyNumpad7,
	ebiten.KeyNumpad8,
	ebiten.KeyNumpad9,
	ebiten.KeyNumpadAdd,
	ebiten.KeyNumpadDecimal,
	ebiten.KeyNumpadDivide,
	ebiten.KeyNumpadEnter,
	ebiten.KeyNumpadEqual,
	ebiten.KeyNumpadMultiply,
	ebiten.KeyNumpadSubtract,
	ebiten.KeyPageDown,
	ebiten.KeyPageUp,
	ebiten.KeyPause,
	ebiten.KeyPeriod,
	ebiten.KeyPrintScreen,
	ebiten.KeyQuote,
	ebiten.KeyScrollLock,
	ebiten.KeySemicolon,
	ebiten.KeyShiftLeft,
	ebiten.KeyShiftRight,
	ebiten.KeySlash,
	ebiten.KeySpace,
	ebiten.KeyTab,
	ebiten.KeyAlt,
	ebiten.KeyControl,
	ebiten.KeyShift,
	ebiten.KeyMeta,
	ebiten.KeyMax,
	ebiten.Key0,
	ebiten.Key1,
	ebiten.Key2,
	ebiten.Key3,
	ebiten.Key4,
	ebiten.Key5,
	ebiten.Key6,
	ebiten.Key7,
	ebiten.Key8,
	ebiten.Key9,
	ebiten.KeyApostrophe,
	ebiten.KeyDown,
	ebiten.KeyGraveAccent,
	ebiten.KeyKP0,
	ebiten.KeyKP1,
	ebiten.KeyKP2,
	ebiten.KeyKP3,
	ebiten.KeyKP4,
	ebiten.KeyKP5,
	ebiten.KeyKP6,
	ebiten.KeyKP7,
	ebiten.KeyKP8,
	ebiten.KeyKP9,
	ebiten.KeyKPAdd,
	ebiten.KeyKPDecimal,
	ebiten.KeyKPDivide,
	ebiten.KeyKPEnter,
	ebiten.KeyKPEqual,
	ebiten.KeyKPMultiply,
	ebiten.KeyKPSubtract,
	ebiten.KeyLeft,
	ebiten.KeyLeftBracket,
	ebiten.KeyMenu,
	ebiten.KeyRight,
	ebiten.KeyRightBracket,
	ebiten.KeyUp,
}

type GameState int

const (
	Initialize GameState = iota
	Run
	Pause
	Finalize
)

type Gamer interface {
	Init(context *Context)
	Update(context *Context)
	Pause(context *Context)
	Finalize(context *Context)
	Draw(context *Context, screen *ebiten.Image)
}

type Game struct {
	context *Context
	game    Gamer
	state   GameState

	// For debug.
	TimerSystem *ebitick.TimerSystem
}

func NewGameWithContext(context *Context, gamer Gamer) *Game {
	g := new(Game)
	g.context = context
	g.state = Initialize
	g.game = gamer
	g.TimerSystem = ebitick.NewTimerSystem()

	var timer = g.TimerSystem.After(10*time.Second, func() {
		util.PrintMemUsage()
	})
	timer.Loop = true

	return g
}

func (g *Game) IsRunning() bool {
	return g.state == Run
}

func (g *Game) IsPausing() bool {
	return g.state == Pause
}

func (g *Game) Update() error {
	var err error = nil
	var keys = g.getEnteredKeys()
	g.context.SetEnteredKeys(keys)

	g.TimerSystem.Update()

	// Turn on/off
	if util.IsJustPressed(keys, ebiten.KeyEscape) {
		switch g.state {
		case Run:
			g.state = Pause
		case Pause:
			g.state = Run
		default:
			// nothing to do...
		}
	}

	switch g.state {
	case Initialize:
		// Initialize Game Logic and loading resources.
		g.context.loader.Start()
		g.game.Init(g.context)
		g.state = Run
	case Run:
		// Main Loop
		g.game.Update(g.context)
		//g.state = Pause
	case Pause:
		// Can be control in Pause Menu Only.
		g.game.Pause(g.context)
		//g.state = Finalize
	case Finalize:
		// FIXME: Dispose all loaded resources.
		g.game.Finalize(g.context)
		g.context.loader.Stop()
	default:
		err = &IllegalStateError{State: strconv.Itoa((int)(g.state))}
	}

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.game.Draw(g.context, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.context.GetScreenWidth(), g.context.GetScreenHeight()
}

func (g *Game) getEnteredKeys() []*control.Key {

	// FIXME: イメージとしては、A,B,Cと入力方式が違う各デバイスで入力された内容をゲームロジックに渡す前に内部用のキーにマッピングし、それをゲームロジックに渡したい。
	// そうすれば、デバイス毎の差異はゲームロジックで吸収する必要がなくなるため。

	keys := []*control.Key{}
	for _, v := range supportKeys {
		isPressed := false
		isJustPressed := false
		isJustReleased := false
		if ebiten.IsKeyPressed(v) {
			isPressed = true
		}
		if inpututil.IsKeyJustPressed(v) {
			isJustPressed = true
		}
		if inpututil.IsKeyJustReleased(v) {
			isJustReleased = true
		}
		keys = append(keys, control.NewKeyWith(v, isPressed, isJustPressed, isJustReleased))
	}
	return keys
}
