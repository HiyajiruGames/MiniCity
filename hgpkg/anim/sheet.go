package anim

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sheet struct {
	path          *string
	image         *ebiten.Image
	frameRect     image.Rectangle
	speed         float64 // sec
	actionType    int
	maxIndex      int
	maxActionType int
	current       int
	updateCount   int
}

func NewSheet(path *string, frameRect image.Rectangle, speed float64) *Sheet {
	s := new(Sheet)
	s.path = path
	s.frameRect = frameRect
	s.speed = speed
	s.actionType = 0
	s.current = 0
	s.updateCount = 0
	return s
}

func (s *Sheet) Load() error {
	var err error
	s.image, _, err = ebitenutil.NewImageFromFile(*s.path)
	s.maxActionType = s.image.Bounds().Dy() / s.frameRect.Dy()
	s.maxIndex = s.image.Bounds().Dx() / s.frameRect.Dx()
	return err
}

func (s *Sheet) SetActionType(actionType int) {
	if s.maxActionType >= actionType && actionType >= 0 {
		s.actionType = actionType
		//s.calcMaxIndex()
	}
}

func (s *Sheet) Update() {
	if s.updateCount >= int(s.speed*ebiten.ActualTPS()) {
		s.updateCount = 0
		s.current++
		if s.maxIndex <= s.current {
			s.current = 0
		}
	}
	s.updateCount++
}

func (s *Sheet) GetCurrentImage() *ebiten.Image {
	return s.GetImageAt(s.current)
}

func (s *Sheet) GetImageAt(index int) *ebiten.Image {
	x, y := s.frameRect.Dx()*index, s.frameRect.Dy()*s.actionType
	width := x + s.frameRect.Dx()
	height := y + s.frameRect.Dy()
	return s.image.SubImage(image.Rect(x, y, width, height)).(*ebiten.Image)
}

func (s *Sheet) String() string {
	return fmt.Sprintf("Sheet {path: %s, frameRect: %s, speed: %f, actionType: %d, current: %d, updateCount: %d}", *s.path, s.frameRect.String(), s.speed, s.actionType, s.current, s.updateCount)
}

func (s *Sheet) calcMaxIndex() int {
	image := s.GetImageAt(0)
	buffer := make([]byte, 4*image.Bounds().Dx()*image.Bounds().Dy())
	image.ReadPixels(buffer)
	return 0
}
