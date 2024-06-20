package core

import (
	"image"
)

type Collider interface {
	IsEntered(p image.Point, r image.Rectangle) bool
}
