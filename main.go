package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize = 16
	tileXNum = 2
)

const xNum = screenWidth / tileSize

var (
	tilesImage *ebiten.Image
)

func init() {
	f, err := ebitenutil.OpenFile("resources/all.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

var p = Player{1, 2}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		p.MoveUp()
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		p.MoveLeft()
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		p.MoveDown()
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		p.MoveRight()

	}
	world.Draw(screen)

	p.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	return nil
}

func main() {

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
