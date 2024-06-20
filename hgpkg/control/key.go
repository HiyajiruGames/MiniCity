package control

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Key struct {
	key            ebiten.Key
	isPressed      bool
	isJustPressed  bool
	isJustReleased bool
}

func NewKey(key ebiten.Key) *Key {
	return NewKeyWith(key, false, false, false)
}

func NewKeyWith(key ebiten.Key, isPressed bool, isJustPressed bool, isJustRelease bool) *Key {
	k := new(Key)
	k.key = key
	k.isPressed = isPressed
	k.isJustPressed = isJustPressed
	k.isJustReleased = isJustRelease
	return k
}

func (k *Key) Contains(key ebiten.Key) bool {
	return k.key == key
}

func (k *Key) IsPressed() bool {
	return k.isPressed
}

func (k *Key) IsJustPressed() bool {
	return k.isJustPressed
}
func (k *Key) isJustRelease() bool {
	return k.isJustReleased
}
