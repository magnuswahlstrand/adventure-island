package main

import (
	"errors"
	"flag"
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/kyeett/adventure-island/render"
	"github.com/kyeett/gameserver"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/grpc"
	"github.com/kyeett/gameserver/localserver"
	"github.com/kyeett/gameserver/types"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 160
	screenHeight = 160
)

var (
	score int
	p     entity.Entity
	p2    entity.Entity
)

var (
	Up    = types.Coord{0, -1}
	Left  = types.Coord{-1, 0}
	Down  = types.Coord{0, 1}
	Right = types.Coord{1, 0}
)

func calculateScore(ownerID string, entities []entity.Entity) int {
	var score int
	for _, e := range entities {
		if e.Owner != ownerID {
			continue
		}

		switch e.Type {
		case entity.Score:
			score++
		}
	}
	return score
}

func randomWalk() {

	var c types.Position

	if time.Now().Nanosecond()%10 == 0 {

		switch rand.Intn(4) {
		case 0:
			c.Coord = p2.Position.Add(Up)
			c.Theta = 0
		case 1:
			c.Coord = p2.Position.Add(Left)
			c.Theta = 3
		case 2:
			c.Coord = p2.Position.Add(Down)
			c.Theta = 2
		case 3:
			c.Coord = p2.Position.Add(Right)
			c.Theta = 1
		}
		p2, _ = s.PerformAction(p2, c)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Game terminated by player")
	}

	if dummyPlayer {
		randomWalk()
	}

	var c types.Position
	if leftPressed() || rightPressed() || upPressed() || downPressed() {
		switch {
		case upPressed():
			c.Coord = p.Position.Add(Up)
			c.Theta = 0
		case leftPressed():
			c.Coord = p.Position.Add(Left)
			c.Theta = 3
		case downPressed():
			c.Coord = p.Position.Add(Down)
			c.Theta = 2
		case rightPressed():
			c.Coord = p.Position.Add(Right)
			c.Theta = 1

		}

		p, _ = s.PerformAction(p, c)
		score = calculateScore(p.ID, s.Entities())
	}

	render.DrawWorld(world, screen)
	for _, e := range s.Entities() {
		render.Draw(&e, screen)
	}

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", score))
	return nil
}

var (
	s     gameserver.GameServer
	world types.World

	dummyPlayer bool
)

func main() {

	var remoteState bool

	flag.BoolVar(&dummyPlayer, "dummy", false, "create a dummy player who walks around randomly, mostly for development purposes")
	flag.BoolVar(&remoteState, "remote", false, "run to remote server")
	flag.Parse()

	switch remoteState {
	case true:
		var err error

		ss, err := grpc.NewServer()
		go ss.Run() //Dummy

		if err != nil {
			log.Fatal(err)
		}

		s, err = grpc.NewClient()
		if err != nil {
			log.Fatal(err)
		}
	default:
		s = localserver.New()
	}

	p, _ = s.NewPlayer()
	if dummyPlayer {
		p2, _ = s.NewPlayer()
	}

	world = s.World()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
