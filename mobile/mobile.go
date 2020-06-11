package mobile

import (
	"github.com/hajimehoshi/ebiten/mobile"

	"ebiten-example/game"
)

func init() {
	mobile.SetGame(&game.Game{})
}

func InitGame(screenWidth, screenHeight int) {
	game.Setup()
	game.SetupViewport(screenWidth, screenHeight)
}
