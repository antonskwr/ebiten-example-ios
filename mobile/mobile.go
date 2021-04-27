package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"ebiten-example/game"
)

func init() {
	mobile.SetGame(&game.Game{})
}

func SetScreenSize(screenWidth, screenHeight int) {
	game.SetupViewport(screenWidth, screenHeight)
}
