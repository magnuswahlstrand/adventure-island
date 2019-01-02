# Lessons learnt

[<- Back to index](README.md)

## 1. Structures for handling byte arrays in a simple way

Egon recommended this simple setup, after I complained about that byte slices were cumbersome to work with:

```go
package main

type Coord struct{ X, Y int }

type Tile byte

const (
	Invalid Tile = iota
	Grass
	Water
)

type Map struct {
	tiles  []Tile
	width  int
	height int
}

func (m *Map) At(p Coord) Tile {
	// do bounds checks here
	return m.tiles[p.Y*m.width+p.X]
}

func (m *Map) At(p Coord, t Tile) {
	// do bounds checks here
	m.tiles[p.Y*m.width+p.X] = tile
}

```
