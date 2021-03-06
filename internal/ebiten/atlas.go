package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type Atlas struct {
	img     *ebiten.Image
	regions map[string]common.AtlasRegion
	cache   map[string]engine.Image
}

func (a *Atlas) GetImage(k string) engine.Image {
	region, ok := a.regions[k]
	if !ok {
		return nil
	}

	eImg, ok := a.cache[k]
	if ok {
		return eImg
	}

	img := a.img.SubImage(
		image.Rect(
			int(region.X),
			int(region.Y),
			int(region.X+region.W),
			int(region.Y+region.H),
		),
	)

	a.cache[k] = &Image{
		img: img.(*ebiten.Image),
		sx:  1,
		sy:  1,
	}

	return a.cache[k]
}
