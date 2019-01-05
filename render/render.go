package render

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/adventure-island/sprites"
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

const tileSize = 16

func Draw(e entity.Entity, screen *ebiten.Image) {
	switch e.Type {
	case entity.Score:
		return
	}

	op := &ebiten.DrawImageOptions{}
	img := sprites.Sprite(e)
	_, h := img.Size()
	op.GeoM.Translate(float64(e.Position.X)*tileSize, float64(e.Position.Y-(h/tileSize-1))*tileSize)
	screen.DrawImage(img, op)
}

func DrawWorld(w types.World, screen *ebiten.Image) {
	// do bounds checks here
	// m.tiles[p.Y*m.width+p.X] = t

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			c := types.Coord{x, y}
			t := w.At(c)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)*tileSize, float64(y)*tileSize)
			screen.DrawImage(sprites.SubImage(t), op)

		}
	}

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {

			c := types.Coord{x, y}
			t := w.At(c)

			if t != types.Water {
				continue
			}

			// Up
			if t := w.At(types.Coord{x, y + 1}); t == types.Grass {
				r0 := rotatedBorder(0)
				r0.GeoM.Translate(float64(x)*tileSize, float64(y+1)*tileSize)
				r0.GeoM.Translate(+tileSize/2, 0)
				screen.DrawImage(sprites.SubImage(types.GrassUp), r0)
			}

			// Down
			if t := w.At(types.Coord{x, y - 1}); t == types.Grass {
				r180 := rotatedBorder(math.Pi)
				r180.GeoM.Translate(float64(x)*tileSize, float64(y)*tileSize)
				r180.GeoM.Translate(+tileSize/2, 0)
				screen.DrawImage(sprites.SubImage(types.GrassUp), r180)
			}

			// Left
			if t := w.At(types.Coord{x - 1, y}); t == types.Grass {
				r270 := rotatedBorder(math.Pi / 2)
				r270.GeoM.Translate(float64(x)*tileSize, float64(y)*tileSize)
				r270.GeoM.Translate(0, +tileSize/2)
				screen.DrawImage(sprites.SubImage(types.GrassUp), r270)
			}

			// Right
			if t := w.At(types.Coord{x + 1, y}); t == types.Grass {
				r270 := rotatedBorder(3 * math.Pi / 2)
				r270.GeoM.Translate(float64(x+1)*tileSize, float64(y)*tileSize)
				r270.GeoM.Translate(0, +tileSize/2)
				screen.DrawImage(sprites.SubImage(types.GrassUp), r270)
			}
		}
	}
}

func rotatedBorder(angle float64) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-tileSize/2, -2)
	op.GeoM.Rotate(angle)
	return op
}
