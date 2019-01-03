package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/kyeett/adventure-island/resources"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 160
	screenHeight = 160
)

const (
	tileSize = 16
	tileXNum = 2
)

const xNum = screenWidth / tileSize

var world Map

var (
	tilesImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.All_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

var p Entity

var (
	Up    = Coord{0, -1}
	Left  = Coord{-1, 0}
	Down  = Coord{0, 1}
	Right = Coord{1, 0}
)

// Up    = Position{Coord{0, -1}, 0}
// Left  = Position{Coord{-1, 0}, 1}
// Down  = Position{Coord{0, 1}, 2}
// Right = Position{Coord{1, 0}, 3}

func randomWalk() {

	var c Position

	if time.Now().Nanosecond()%10 == 0 {

		switch rand.Intn(4) {
		case 0:
			c.Coord = p.Position.Add(Up)
			c.theta = 0
		case 1:
			c.Coord = p.Position.Add(Left)
			c.theta = 3
		case 2:
			c.Coord = p.Position.Add(Down)
			c.theta = 2
		case 3:
			c.Coord = p.Position.Add(Right)
			c.theta = 1
		}
		p = world.MoveTo(p, c)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Game terminated by player")
	}

	randomWalk()

	var c Position
	if leftPressed() || rightPressed() || upPressed() || downPressed() {
		switch {
		case upPressed():
			c.Coord = p.Position.Add(Up)
			c.theta = 0
		case leftPressed():
			c.Coord = p.Position.Add(Left)
			c.theta = 3
		case downPressed():
			c.Coord = p.Position.Add(Down)
			c.theta = 2
		case rightPressed():
			c.Coord = p.Position.Add(Right)
			c.theta = 1
		}

		// if world.ValidTarget(c) == true {
		// 	p.MoveTo(c)
		// }
		p = world.MoveTo(p, c)
	}

	world.CheckCollisions()

	world.Draw(screen)
	for _, o := range world.entities {
		o.Draw(screen)
	}

	// p.Draw(screen)
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", world.GetScore()))

	return nil
}

func main() {

	world = NewMap()
	p = world.AddPlayer()

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
