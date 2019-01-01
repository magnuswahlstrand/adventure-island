package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Coord struct{ X, Y int }

type Map struct {
	tiles  []Tile
	width  int
	height int
}

func (m *Map) Draw(screen *ebiten.Image) {
	// do bounds checks here
	// m.tiles[p.Y*m.width+p.X] = t

	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			c := Coord{x, y}
			t := m.At(c)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)*tileSize, float64(y)*tileSize)
			screen.DrawImage(subImage(t), op)

		}
	}

	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {

			c := Coord{x, y}
			t := m.At(c)

			if t != Water {
				continue
			}

			// Up
			if t := m.At(Coord{x, y + 1}); t == Grass {
				r0 := rotatedBorder(0)
				r0.GeoM.Translate(float64(x)*tileSize, float64(y+1)*tileSize)
				r0.GeoM.Translate(+tileSize/2, 0)
				screen.DrawImage(subImage(GrassUp), r0)
			}

			// Down
			if t := m.At(Coord{x, y - 1}); t == Grass {
				r180 := rotatedBorder(math.Pi)
				r180.GeoM.Translate(float64(x)*tileSize, float64(y)*tileSize)
				r180.GeoM.Translate(+tileSize/2, 0)
				screen.DrawImage(subImage(GrassUp), r180)
			}

			// Left
			if t := m.At(Coord{x - 1, y}); t == Grass {
				r270 := rotatedBorder(math.Pi / 2)
				r270.GeoM.Translate(float64(x)*tileSize, float64(y)*tileSize)
				r270.GeoM.Translate(0, +tileSize/2)
				screen.DrawImage(subImage(GrassUp), r270)
			}

			// Right
			if t := m.At(Coord{x + 1, y}); t == Grass {
				r270 := rotatedBorder(3 * math.Pi / 2)
				r270.GeoM.Translate(float64(x+1)*tileSize, float64(y)*tileSize)
				r270.GeoM.Translate(0, +tileSize/2)
				screen.DrawImage(subImage(GrassUp), r270)
			}
		}
	}
}

func (m *Map) At(p Coord) Tile {
	if p.X < 0 || p.X >= m.width || p.Y < 0 || p.Y >= m.height {
		return Invalid
	}
	// do bounds checks here
	return m.tiles[p.Y*m.width+p.X]
}

func (m *Map) Set(p Coord, t Tile) {
	if p.X < 0 || p.X >= m.width || p.Y < 0 || p.Y >= m.height {
		return
	}

	// do bounds checks here
	m.tiles[p.Y*m.width+p.X] = t
}

func rotatedBorder(angle float64) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-tileSize/2, -2)
	op.GeoM.Rotate(angle)
	return op
}

var (
	world = Map{
		tiles: []Tile{
			Water, Water, Water, Water, Water, Water,
			Water, Water, Grass, Grass, Water, Water,
			Water, Grass, Grass, Water, Grass, Water,
			Water, Grass, Grass, Grass, Grass, Water,
			Water, Water, Grass, Grass, Water, Grass,
			Water, Water, Water, Water, Water, Grass,
		},
		width:  6,
		height: 6,
	}
)
