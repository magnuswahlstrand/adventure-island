package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

func subImage(t Tile) *ebiten.Image {
	sx := (int(t-Water) % tileXNum) * tileSize
	sy := (int(t-Water) / tileXNum) * tileSize
	return tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}

func translateTile(t Tile) (float64, float64) {
	return float64((int(t) % xNum) * tileSize), float64((int(t) / xNum) * tileSize)
}

type Tile int

const (
	Invalid Tile = iota + 1
	Water
	Grass
	GrassUp = 12 + Water
)
