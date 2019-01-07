package sprites

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
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
	objectImage    *ebiten.Image
	characterImage *ebiten.Image
	tilesImage     *ebiten.Image
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
	opts := gif.Options{
		NumColors: 4,
		Drawer:    draw.FloydSteinberg,
	}
	fmt.Println(opts)
	// b := img.Bounds()

	// More or less taken from the image/gif package
	// pimg := image.NewPaletted(b, palette.Plan9[:opts.NumColors])
	// pimg := image.NewPaletted(b, palette.WebSafe[:128]),

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

	// // pimg.Palette.Convert(c color.Color)

	// if opts.Quantizer != nil {
	// 	pimg.Palette = opts.Quantizer.Quantize(make(color.Palette, 0, opts.NumColors), img)
	// }

	// draw.Draw(pimg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	// b = b.Union(image.Rect(5, 5, 50, 50))

	// draw.Draw(pimg, b, &image.Uniform{color.White}, image.Point{10, 10})
	// // draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// // paste PNG image OVER to newImage
	// draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)
	// draw.Draw(pimg, b, img, img.Bounds().Min, draw.Over)

	// draw.Draw(pimg, pimg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste PNG image OVER to newImage
	// draw.Draw(pimg, pimg.Bounds(), img, img.Bounds().Min, draw.Over)

	// draw.Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op)
	// opts.Drawer.Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)
	// opts.Drawer.Draw(pimg, b, img, image.ZP)

	return pimg
}

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
	img = addFrame(img)
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
		log.Fatalf("Entity type %s is not valid, shutting down", e)
	}
	return img
}
