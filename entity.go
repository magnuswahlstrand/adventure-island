package main

import (
	"bytes"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure-island/resources"
)

var a = 1

var objectImage *ebiten.Image

var characterImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.Character_png))

	if err != nil {
		log.Fatal(err)
	}
	characterImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(resources.Objects_png))

	if err != nil {
		log.Fatal(err)
	}
	objectImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type EntityType int

type Entity struct {
	ID   string
	Type EntityType
	Position
}

const (
	Coin EntityType = iota + 1
	Score
	Character
	Empty
)

func subObject(typ EntityType, frame int) *ebiten.Image {
	var width, height, offsetX, offsetY, sx, sy int
	switch typ {
	case Coin:
		width, height = 16, 16
		offsetY = 62
		frame = frame % 4
	}
	sx = offsetX + frame*width
	sy = offsetY
	return objectImage.SubImage(image.Rect(sx, sy, sx+width, sy+height)).(*ebiten.Image)
}

func (e Entity) Draw(screen *ebiten.Image) {

	switch e.Type {
	case Score:
		return
	}
	// case Character:

	// 	op.GeoM.Translate(float64(e.X)*tileSize, float64(e.Y-1)*tileSize)
	// 	t := time.Now().Nanosecond() / 1000 / 1000 / 250 // 10th of 2nd
	// 	screen.DrawImage(subCharacter(t%4, p.direction-down), op)

	// case Coin:

	// 	t := time.Now().Nanosecond() / 1000 / 1000 / 100 // 10th of 2nd
	// }
	op := &ebiten.DrawImageOptions{}
	img := e.Sprite()
	_, h := img.Size()
	op.GeoM.Translate(float64(e.Position.X)*tileSize, float64(e.Y-(h/tileSize-1))*tileSize)
	screen.DrawImage(img, op)
}

func (e Entity) Destroy() Entity {
	switch e.Type {
	case Coin:
		return Entity{
			Position: Position{Coord{-1, -1}, 0},
			Type:     Score,
		}
	}

	return Entity{
		Position: Position{Coord{-1, -1}, 0},
		Type:     Empty,
	}
}

func (e Entity) Sprite() *ebiten.Image {
	t := time.Now().Nanosecond() / 1000 / 1000 / 250

	var img *ebiten.Image
	switch e.Type {
	case Character:
		charWidth := 16
		charHeight := 32
		sx := t % 4 * charWidth

		var direction int
		switch e.theta {
		case 0:
			direction = 2
		case 2:
			direction = 0
		default:
			direction = e.theta

		}
		sy := direction * charHeight
		img = characterImage.SubImage(image.Rect(sx, sy, sx+charWidth, sy+charHeight)).(*ebiten.Image)
	case Coin:
		img = subObject(e.Type, t%4)
	}
	return img
}
