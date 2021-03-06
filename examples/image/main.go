package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Image",
		854,
		480,
		engine.FlagResizable,
		// use Ebiten backend
		ardent.EBITEN,
		// tick function
		func() {},
		// layout function
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	// get component factory
	component := game.Component()

	// create new renderer and image
	renderer := component.NewRenderer()
	image, _ := component.NewImageFromAssetPath("scs.asset")

	// add image to renderer
	renderer.AddImage(image)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
