package util

import (
	"github.com/HiyajiruGames/MiniCity/hgpkg/control"
	"github.com/hajimehoshi/ebiten/v2"
)

func IsPressed(inputKeys []*control.Key, key ebiten.Key) bool {
	var ret = false
	for _, v := range inputKeys {
		if v.Contains(key) {
			ret = v.IsPressed()
			break
		}
	}
	return ret
}

func IsJustPressed(inputKeys []*control.Key, key ebiten.Key) bool {
	var ret = false
	for _, v := range inputKeys {
		if v.Contains(key) {
			ret = v.IsJustPressed()
			break
		}
	}
	return ret
}

func IsJustRelease(inputKeys []*control.Key, key ebiten.Key) bool {
	var ret = false
	for _, v := range inputKeys {
		if v.Contains(key) {
			ret = v.IsJustPressed()
			break
		}
	}
	return ret
}
