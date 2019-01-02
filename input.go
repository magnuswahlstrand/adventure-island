package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func touched() (int, int) {
	if len(inpututil.JustPressedTouchIDs()) == 0 {
		return -1, -1
	}

	return ebiten.TouchPosition(ebiten.TouchIDs()[0])
}

func leftPressed() bool {
	x, y := touched()
	if (x != -1 && y != -1) && x < y && x < screenWidth-y {
		fmt.Println(x, y, x, screenWidth-y, "Left")
		return true
	}

	return inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func rightPressed() bool {
	x, y := touched()
	if (x != -1 && y != -1) && x > y && y > screenWidth-x {
		fmt.Println(x, y, x, screenWidth-y, "Right")
		return true
	}

	return inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func upPressed() bool {
	x, y := touched()
	if (x != -1 && y != -1) && x > y && x < screenWidth-y {
		fmt.Println(x, y, x, screenWidth-y, "Up")
		return true
	}

	return inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func downPressed() bool {
	x, y := touched()
	if (x != -1 && y != -1) && x < y && y > screenWidth-x {
		fmt.Println(x, y, x, screenWidth-y, "Down")
		return true
	}
	return inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown)
}
