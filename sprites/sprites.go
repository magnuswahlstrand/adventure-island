package sprites

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"

	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure-island/resources"
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

var (
	objectImage         *ebiten.Image
	characterImage      *ebiten.Image
	characterColorImage *ebiten.Image
	tilesImage          *ebiten.Image
)

var hairColors = []struct {
	name  string
	color color.NRGBA
}{
	// {"Black", color.NRGBA{9, 8, 6, 255}},
	// {"Off Black", color.NRGBA{44, 34, 43, 255}},
	// {"Dark Gray", color.NRGBA{113, 99, 90, 255}},
	// {"Medium Gray", color.NRGBA{183, 166, 158, 255}},
	// {"Light Gray", color.NRGBA{214, 196, 194, 255}},
	// {"Platinum Blonde", color.NRGBA{202, 191, 177, 255}},
	// {"Bleached Blonde", color.NRGBA{220, 208, 186, 255}},
	// {"White Blonde", color.NRGBA{255, 245, 225, 255}},
	// {"Light Blonde", color.NRGBA{230, 206, 168, 255}},
	// {"Golden Blonde", color.NRGBA{229, 200, 168, 255}},
	// {"Ash Blonde", color.NRGBA{222, 188, 153, 255}},
	// {"Honey Blonde", color.NRGBA{184, 151, 120, 255}},
	// {"Strawberry Blonde", color.NRGBA{165, 107, 70, 255}},
	// {"Light Red", color.NRGBA{181, 82, 57, 255}},
	// {"Dark Red", color.NRGBA{141, 74, 67, 255}},
	// {"Light Auburn", color.NRGBA{145, 85, 61, 255}},
	// {"Dark Auburn", color.NRGBA{83, 61, 50, 255}},
	// {"Dark Brown", color.NRGBA{59, 48, 36, 255}},
	// {"Golden Brown", color.NRGBA{85, 72, 56, 255}},
	// {"Medium Brown", color.NRGBA{78, 67, 63, 255}},
	// {"Chestnut Brown", color.NRGBA{80, 68, 68, 255}},
	// {"Brown", color.NRGBA{106, 78, 66, 255}},
	{"Light Brown", color.NRGBA{167, 133, 106, 255}},
	{"Ash Brown", color.NRGBA{151, 121, 97, 255}},
}

func addFrame(img image.Image) image.Image {

	hairColor := color.NRGBA{106, 72, 52, 255}
	hairDarkColor := color.NRGBA{67, 46, 39, 255}
	shirtColor := color.NRGBA{196, 60, 60, 255}
	shirtDarkColor := color.NRGBA{136, 46, 46, 255}

	newHair := hairColors[rand.Int31n(int32(len(hairColors)))].color
	newDark := color.NRGBA{uint8(uint32(newHair.R) * 5 / 10), uint8(uint32(newHair.G) * 5 / 10), uint8(uint32(newHair.B) * 5 / 10), 255}

	pimg := imaging.AdjustFunc(
		img,
		func(c color.NRGBA) color.NRGBA {
			// shift the red channel by 16

			switch c {
			case shirtColor, shirtDarkColor:
				return color.NRGBA{c.G, c.B, c.R, c.A}
			case hairColor:
				return newHair
			case hairDarkColor:
				return newDark
			}
			return c
			// return color.NRGBA{c.R + 130, c.G + 160, c.B, c.A}
		},
	)

	return pimg
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(resources.All_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(resources.Objects_png))
	if err != nil {
		log.Fatal(err)
	}
	objectImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(resources.Character_png))
	if err != nil {
		log.Fatal(err)
	}
	img = addFrame(img)
	characterImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(resources.Character_color_png))
	if err != nil {
		log.Fatal(err)
	}
	characterColorImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
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

func subObject(e entity.Entity, frame int) *ebiten.Image {
	var width, height, offsetX, offsetY, sx, sy int
	switch e.Type {
	case entity.Coin:
		width, height = 16, 16
		offsetY = 62
		frame = frame % 4
	case entity.Bridge:
		width, height = 48, 48
		offsetX = 224
		offsetY = 96
		frame = e.Position.Theta % 2
	}
	sx = offsetX + frame*width
	sy = offsetY
	return objectImage.SubImage(image.Rect(sx, sy, sx+width, sy+height)).(*ebiten.Image)
}

func subCharacter(e entity.Entity, frame int) *ebiten.Image {
	charWidth := 16
	charHeight := 32
	sx := frame * charWidth

	var direction int
	switch e.Position.Theta {
	case 0:
		direction = 2
	case 2:
		direction = 0
	default:
		direction = e.Position.Theta
	}
	sy := direction * charHeight

	// offset := frame % 2

	// Get color based on user ID
	var headIndex, bodyIndex int
	l := len(e.ID)
	fmt.Sscanf(e.ID[l-3:l-2], "%X", &headIndex)
	fmt.Sscanf(e.ID[l-1:l], "%X", &bodyIndex)

	// Offset based on color
	bodyColorX := charWidth * 4 * (bodyIndex % 12)
	headColorX := charWidth * 4 * (headIndex % 12)

	tmpImg := image.NewNRGBA(image.Rect(0, 0, 16, 32))
	draw.Draw(tmpImg, image.Rect(0, 16, 16, 32), characterColorImage, image.Pt(bodyColorX+sx, 16+sy), draw.Src)
	draw.Draw(tmpImg, image.Rect(0, 0, 16, 16), characterColorImage, image.Pt(headColorX+sx, sy), draw.Src)

	tmpImg2, err := ebiten.NewImageFromImage(tmpImg, ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Error while converting image", err)
	}

	return tmpImg2
}

func Sprite(e entity.Entity) *ebiten.Image {
	t := time.Now().Nanosecond() / 1000 / 1000 / 250

	var img *ebiten.Image
	switch e.Type {
	case entity.Character:
		img = subCharacter(e, t%4)
	case entity.Coin, entity.Bridge:
		img = subObject(e, t%4)
	default:
		log.Fatalf("Entity type %s is not valid, shutting down", e)
	}
	return img
}
