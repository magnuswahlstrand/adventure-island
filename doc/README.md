# Adventure Island

Special thanks to @peterhellberg, @hajimehosh, and @Egon for support and guidance

## Day 1

### Part 1 - Draw world

**Estimated (actual):** 5h (4h)

- [x] Draw small 2d world, 10x10 grid
- [x] Single player character in middle
      **Comments**: Initial setup always takes longer than expected. Had to (re)learn both [faiface/pixel](https://github.com/faiface/pixel) and [hajimehoshi/ebiten](https://github.com/hajimehoshi/ebiten). Spent an hour on reading sprites, and another couple of hours on good ways of drawing transitions between `Tile` types. Got lot of help from `#gamedev`.

### Part 2 - Interaction

**Estimated (actual):** 2h (1h)

- [x] Move player character
- [x] Using WASD or arrow keys

**Comments**: Very easy using `ebiten`. Will need to refactor at some point, when game gets more complex. Introduced animation and player direction while I was at it.

### Part 3 - Collision

**Estimated (actual):** 4h (30min)

**Comments**: Only implemented collision with water. Works well.

### Progress - Day 1

Drawing a 2D map, simple player character controlled by WASD or keys, and collision with water.

![Result day 1](day-1.gif)

## Planned

#### Part 4 - Collect coins + score **Estimated (actual):** 3h ()

#### Part 5 - Sync game state over network **Estimated (actual):** 8h ()

#### Part 6 - Multiplayer **Estimated (actual):** 8h ()

#### Part 7 - Animation between tiles - **Estimated (actual):** 4h ()

#### Part 8 - Add trees - **Estimated (actual):** 1h ()

#### Part 9 - Character animation **Estimated (actual):** 6h ()

#### Part 10 - Several Z-levels **Estimated (actual):** 6h ()

#### Part 11 - Game chat **Estimated (actual):** 3h ()

#### Part 12 - Compile and publish javascript **Estimated (actual):** 2h ()

### Estimated (actual): 55h (5h30+)

## Resources

- `#gamedev` and `#ebiten` on `https://gophers.slack.com`
- <https://www.mapeditor.org/> - mapeditor for tile games
- [faiface/pixel](https://github.com/faiface/pixel) - 2D game library in Go, used initially
- [hajimehoshi/ebiten](https://github.com/hajimehoshi/ebiten) - Another 2D game library in Go. Can be compiled to javascript using WASM or GopherJS
- [Marching squares](https://en.wikipedia.org/wiki/Marching_squares) - Algorithm for generating contours in 2D maps

## Lessons learnt

[Lessons learnt along the way](LESSONS_LEARNT.md)
