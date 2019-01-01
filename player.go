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

type Player struct {
	X, Y int
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X)*tileSize, float64(p.Y-1)*tileSize)
	t := time.Now().Nanosecond() / 1000 / 1000 / 250 // 10th of 2nd
	s := time.Now().Second()
	screen.DrawImage(subCharacter(t%4, s%4), op)
}

func (p *Player) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func (p *Player) MoveUp() {
	p.Move(0, -1)
}
func (p *Player) MoveLeft() {
	p.Move(-1, 0)
}
func (p *Player) MoveDown() {
	p.Move(0, 1)
}
func (p *Player) MoveRight() {
	p.Move(1, 0)
}
