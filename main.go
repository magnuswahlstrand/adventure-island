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

var p = Player{Coord{1, 2}, right}

var (
	Up    = Coord{0, -1}
	Left  = Coord{-1, 0}
	Down  = Coord{0, 1}
	Right = Coord{1, 0}
)

func leftPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func rightPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func upPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func downPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// if time.Now().Nanosecond()%10 == 0 {
	// 	var c Coord
	// 	switch rand.Intn(4) {
	// 	case 0:
	// 		c = p.PrepareMove(Up)
	// 	case 1:
	// 		c = p.PrepareMove(Left)
	// 	case 2:
	// 		c = p.PrepareMove(Down)
	// 	case 3:
	// 		c = p.PrepareMove(Right)
	// 	}

	// 	if world.ValidTarget(c) == true {
	// 		p.MoveTo(c)
	// 	}
	// }

	var c Coord

	if leftPressed() || rightPressed() || upPressed() || downPressed() {
		switch {
		case upPressed():
			c = p.PrepareMove(Up)
		case leftPressed():
			c = p.PrepareMove(Left)
		case downPressed():
			c = p.PrepareMove(Down)
		case rightPressed():
			c = p.PrepareMove(Right)
		}

		if world.ValidTarget(c) == true {
			p.MoveTo(c)
		}
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
