package game

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

const (
	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
	frameTime   = 90
)

var (
	screenWidth, screenHeight int
	runnerImage               *ebiten.Image
	once                      sync.Once

	lastTime    time.Time = time.Now()
	readyToDraw bool      = false
	elapsedTime int64     = 0
)

type Game struct {
	count int
}

func init() {
	setup()
}

func (g *Game) Update(screen *ebiten.Image) error {
	t := time.Since(lastTime).Milliseconds()
	lastTime = time.Now()
	elapsedTime += t
	if elapsedTime > frameTime {
		mod := elapsedTime % frameTime
		elapsedTime = mod
		g.count++
	} else if elapsedTime == frameTime {
		elapsedTime = 0
		g.count++
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !readyToDraw {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(2), float64(2))
	op.GeoM.Translate(-float64(frameWidth*2)/2, -float64(frameHeight*2)/2)
	op.GeoM.Translate(float64(screenWidth/2), float64(screenHeight/2))
	i := g.count % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func setup() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	fmt.Println("Setup Complete")
}

func SetupViewport(width, height int) {
	screenWidth = width
	screenHeight = height

	if !readyToDraw {
		readyToDraw = true
	}
}
