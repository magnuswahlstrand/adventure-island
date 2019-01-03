package main

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure-island/resources"
)

var characterImage *ebiten.Image

type Player string

// type Player struct {
// 	entity    Entity
// 	direction int
// }

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.Character_png))

	if err != nil {
		log.Fatal(err)
	}
	characterImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func subCharacter(step int, direction int) *ebiten.Image {
	charWidth := 16
	charHeight := 32
	// sx := (int(t-Water) % tileXNum) * tileSize
	// sy := (int(t-Water) / tileXNum) * tileSize
	sx := step * charWidth
	sy := direction * charHeight
	return characterImage.SubImage(image.Rect(sx, sy, sx+charWidth, sy+charHeight)).(*ebiten.Image)
}

const (
	down int = iota + 1
	right
	up
	left
)

func (p *Player) Move(c Coord) {
	p.entity.X += c.X
	p.entity.Y += c.Y
}

func (p *Player) MoveTo(c Coord) {
	p.entity.X = c.X
	p.entity.Y = c.Y
}

func (p *Player) MoveUp() {
	p.direction = up
	p.Move(Coord{0, -1})
}
func (p *Player) MoveLeft() {
	p.direction = left
	p.Move(Coord{-1, 0})
}
func (p *Player) MoveDown() {
	p.direction = down
	p.Move(Coord{0, 1})
}
func (p *Player) MoveRight() {
	p.direction = right
	p.Move(Coord{1, 0})
}

func (p *Player) PrepareMove(c Coord) Coord {
	switch c {
	case Up:
		p.direction = up
	case Left:
		p.direction = left
	case Down:
		p.direction = down
	case Right:
		p.direction = right

	}
	return p.entity.Coord.Add(c)
}
