package main

import (
	"errors"
	"flag"
	"fmt"
	_ "image/png"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/adventure-island/conf"
	"github.com/kyeett/adventure-island/render"

	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/game"
	"github.com/kyeett/gameserver/types"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	hair, body = 0, 0
	score      int
	p          entity.Entity
	dummies    []entity.Entity
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

	for i, d := range dummies {
		if (time.Now().Nanosecond()+1000*i)%20000 == 0 {

			switch rand.Intn(4) {
			case 0:
				c.Coord = d.Position.Add(Up)
				c.Theta = 0
			case 1:
				c.Coord = d.Position.Add(Left)
				c.Theta = 3
			case 2:
				c.Coord = d.Position.Add(Down)
				c.Theta = 2
			case 3:
				c.Coord = d.Position.Add(Right)
				c.Theta = 1
			}

			e, err := g.PerformAction(d, c)
			if err != nil {
				log.Errorf("invalid move %s", e)
			} else {
				dummies[i] = e
			}
		}
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

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		hair = (hair + 1) % 12
	case inpututil.IsKeyJustPressed(ebiten.KeyJ):
		body = (body + 1) % 12
	case inpututil.IsKeyJustPressed(ebiten.KeyP):
		e, err := g.NewPlayer()
		if err != nil {
			log.Fatal(err)
		}
		dummies = append(dummies, e)
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

		e, err := g.PerformAction(p, c)
		if err != nil {
			log.Errorf("invalid move %s", p)
		} else {
			p = e
		}
		score = calculateScore(p.ID, g.Entities())
	}

	render.DrawWorld(world, screen)

	for _, e := range g.Entities() {
		if e.ID == p.ID {
			e.ID = fmt.Sprintf(e.ID[:len(e.ID)-4]+"0%X0%X", hair, body)
		}
		render.Draw(e, screen)
	}

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d\n\n\n\n\n\n\n\n\nKeys: WASD + H, J, P", score))
	return nil
}

var (
	g     *game.Game
	world types.World

	dummyPlayer bool
)

func main() {

	var addr, worldName string
	var secure, dev, dummy bool
	flag.BoolVar(&dummy, "dummy", false, "create a dummy player who walks around randomly, mostly for development purposes")
	flag.BoolVar(&dev, "dev", false, "start the development server on local machine on :10001")
	flag.BoolVar(&secure, "secure", false, "enable TLS")
	flag.StringVar(&addr, "addr", "", "address to remote server: default: run local mode")
	flag.StringVar(&worldName, "world", "", "name of the world to play on")
	flag.Parse()

	addr, worldName, dev, secure, dummy = conf.Conf(addr, worldName, dev, secure, dummy)
	dummyPlayer = dummy

	opts := []game.Option{}
	if worldName != "" {
		opts = append(opts, game.World(worldName))
	}

	if dev {
		opts = append(opts, game.DevServer("localhost:10001"))
	}

	if addr != "" {
		opts = append(opts, game.RemoteState(addr, secure))
	}

	var err error
	g, err = game.New(opts...)
	if err != nil {
		log.Fatal(err)
	}

	p, err = g.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}

	if dummyPlayer {
		p2, err := g.NewPlayer()
		if err != nil {
			log.Fatal(err)
		}
		dummies = append(dummies, p2)
	}

	world = g.World()

	if err := ebiten.Run(update, world.Width*16, world.Height*16, 2, "Tiles (Ebiten Demo)"); err != nil {
		log.Fatal("Game exited: ", err)
	}
}
