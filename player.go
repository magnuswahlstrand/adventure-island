package main

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var characterImage *ebiten.Image

func init() {
	f, err := ebitenutil.OpenFile("resources/character.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
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
	var sx int
	sx = step * charWidth
	sy := direction * charHeight
	return characterImage.SubImage(image.Rect(sx, sy, sx+charWidth, sy+charHeight)).(*ebiten.Image)
}

const (
	down int = iota + 1
	right
	up
	left
)

type Player struct {
	Coord
	direction int
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X)*tileSize, float64(p.Y-1)*tileSize)
	t := time.Now().Nanosecond() / 1000 / 1000 / 250 // 10th of 2nd
	screen.DrawImage(subCharacter(t%4, p.direction-down), op)
}

func (p *Player) Move(c Coord) {
	p.X += c.X
	p.Y += c.Y
}

func (p *Player) MoveTo(c Coord) {
	p.X = c.X
	p.Y = c.Y
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
	return p.Coord.Add(c)
}
