package resource

import (
	"bytes"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image struct {
	image *ebiten.Image
	Resource
}

func NewImage(path string) *Image {
	i := new(Image)
	i.Resource = *NewResource(path)
	return i
}

func (i *Image) Load() {
	i.Resource.Load()
	img, _, err := image.Decode(bytes.NewReader(i.data))
	if err != nil {
		panic(fmt.Sprintf("load error...reason is %s", err.Error()))
	}
	i.image = ebiten.NewImageFromImage(img)
	i.data = nil
}

func (i *Image) Release() {
	i.Resource.Release()
	i.image.Dispose()
	i.image = nil
}

func (i *Image) IsLoaded() bool {
	return i.Resource.IsLoaded() && i.image != nil
}

func (i *Image) GetImage() *ebiten.Image {
	return i.image
}
