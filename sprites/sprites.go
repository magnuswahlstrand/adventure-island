package sprites

import (
	"bytes"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure-island/resources"
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

var (
	objectImage    *ebiten.Image
	characterImage *ebiten.Image
	tilesImage     *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.All_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(resources.Character_png))

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

const (
	tileSize = 16
	tileXNum = 2
)

//Todo: how to propagate screenwidth?

const xNum = 480 / tileSize

func SubImage(t types.Tile) *ebiten.Image {
	sx := (int(t-types.Water) % tileXNum) * tileSize
	sy := (int(t-types.Water) / tileXNum) * tileSize
	return tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}

func subObject(typ entity.Type, frame int) *ebiten.Image {
	var width, height, offsetX, offsetY, sx, sy int
	switch typ {
	case entity.Coin:
		width, height = 16, 16
		offsetY = 62
		frame = frame % 4
	}
	sx = offsetX + frame*width
	sy = offsetY
	return objectImage.SubImage(image.Rect(sx, sy, sx+width, sy+height)).(*ebiten.Image)
}

func subCharacter(dir, frame int) *ebiten.Image {
	charWidth := 16
	charHeight := 32
	sx := frame * charWidth

	var direction int
	switch dir {
	case 0:
		direction = 2
	case 2:
		direction = 0
	default:
		direction = dir

	}

	sy := direction * charHeight
	return characterImage.SubImage(image.Rect(sx, sy, sx+charWidth, sy+charHeight)).(*ebiten.Image)
}

func Sprite(e *entity.Entity) *ebiten.Image {
	t := time.Now().Nanosecond() / 1000 / 1000 / 250

	var img *ebiten.Image
	switch e.Type {
	case entity.Character:
		img = subCharacter(e.Position.Theta, t%4)
	case entity.Coin:
		img = subObject(e.Type, t%4)
	default:
		log.Fatalf("Entity type %s is not valid", e)
	}
	return img
}
