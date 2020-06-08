package main

import (
	_ "image/png"
	"log"

	"ebiten-example/game"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	game.SetupViewport(screenWidth, screenHeight)
	game.Setup()
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
