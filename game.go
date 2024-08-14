package main

import (
	"math/rand"

	blocks "github.com/fpedroso/golang-raylib-tetris/blocks"
	constants "github.com/fpedroso/golang-raylib-tetris/constants"
	statics "github.com/fpedroso/golang-raylib-tetris/statics"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	grid         Grid
	blocks       []blocks.Block
	currentBlock blocks.Block
	nextBlock    blocks.Block
	GameOver     bool
	Score        int
	sounds       statics.Sounds
}

func NewGame(sounds statics.Sounds) Game {
	g := Game{}
	g.grid = Grid{}
	g.sounds = sounds
	g.grid.Speed = 500
	g.blocks = GetAllBlocks()
	g.currentBlock = g.GetRandomBlock()
	g.nextBlock = g.GetRandomBlock()
	g.Score = 0
	g.GameOver = false
	return g
}

func (game *Game) GetRandomBlock() blocks.Block {
	if len(game.blocks) == 0 {
		game.blocks = GetAllBlocks()
	}
	index := rand.Int() % len(game.blocks)
	block := game.blocks[index]
	game.RemoveBlock(index)
	return block
}

func GetAllBlocks() []blocks.Block {
	return []blocks.Block{
		blocks.NewIBlock(),
		blocks.NewJBlock(),
		blocks.NewLBlock(),
		blocks.NewOBlock(),
		blocks.NewSBlock(),
		blocks.NewTBlock(),
		blocks.NewZBlock(),
	}
}

func (game *Game) RemoveBlock(index int) {
	game.blocks = append(game.blocks[:index], game.blocks[index+1:]...)
}

func (game Game) Draw() {
	game.grid.Draw()
	game.currentBlock.Draw(0, 0)
	game.DrawNext()
}

func (game Game) DrawNext() {
	nextBlock := game.nextBlock
	nextBlock.RowOffset = 0
	nextBlock.ColOffset = 0
	nextBlock.Draw(650, 250)
}

func (game *Game) HandleInput(eventTriggered TimeHandler) {
	if game.GameOver {
		if rl.IsKeyPressed(rl.KeyEnter) {
			game.Reset()
		}
		return
	}
	switch {
	case rl.IsKeyPressed(rl.KeyLeft):
		game.MoveBlockLeft()
	case rl.IsKeyPressed(rl.KeyRight):
		game.MoveBlockRight()
	case rl.IsKeyDown(rl.KeyDown):
		if eventTriggered(50) {
			game.MoveBlockDown()
			game.UpdateScore(0, 1)
		}
	case rl.IsKeyPressed(rl.KeyUp):
		game.RotateBlock()
	case rl.IsKeyPressed(rl.KeySpace):
		game.SwitchBlock()
		game.MoveBlockInside()
	}
}

func (game *Game) SwitchBlock() {
	tempBlock := game.currentBlock
	game.currentBlock = game.nextBlock
	game.currentBlock.ColOffset = tempBlock.ColOffset
	game.currentBlock.RowOffset = tempBlock.RowOffset
	game.nextBlock = tempBlock
}

func (game *Game) MoveBlockLeft() {
	if game.GameOver {
		return
	}
	game.currentBlock.Move(0, -1)
	if game.IsBlockOutside() || !game.BlockFits() {
		game.currentBlock.Move(0, 1)
	}
}

func (game *Game) MoveBlockRight() {
	if game.GameOver {
		return
	}
	game.currentBlock.Move(0, 1)
	if game.IsBlockOutside() || !game.BlockFits() {
		game.currentBlock.Move(0, -1)
	}
}

func (game *Game) MoveBlockDown() {
	if game.GameOver {
		return
	}
	game.currentBlock.Move(1, 0)
	if game.IsBlockOutside() || !game.BlockFits() {
		game.currentBlock.Move(-1, 0)
		game.LockBlock()
	}
}

func (game *Game) RotateBlock() {
	if game.GameOver {
		return
	}
	game.currentBlock.Rotate()
	if game.IsBlockOutside() {
		game.MoveBlockInside()
	} else if !game.BlockFits() {
		game.currentBlock.UndoRotate()
	}
}

func (game *Game) IsBlockOutside() bool {
	tiles := game.currentBlock.GetCurrentPositions()
	for _, tile := range tiles {
		if game.grid.IsCellOutside(tile.Row, tile.Column) {
			return true
		}
	}
	return false
}

func (game *Game) MoveBlockInside() {
	lowestColumnOutside := 0
	highestColumnOutside := constants.Cols - 1
	columnFix := 0
	tiles := game.currentBlock.GetCurrentPositions()
	for _, tile := range tiles {
		if tile.Column < lowestColumnOutside {
			lowestColumnOutside = tile.Column
			columnFix = -1 * lowestColumnOutside
		}
		if tile.Column > highestColumnOutside {
			highestColumnOutside = tile.Column
			columnFix = constants.Cols - (highestColumnOutside + 1)
		}
	}
	game.currentBlock.Move(0, columnFix)
}

func (game *Game) LockBlock() {
	if game.GameOver {
		return
	}
	tiles := game.currentBlock.GetCurrentPositions()
	for _, tile := range tiles {
		game.grid.Cells[tile.Row][tile.Column] = game.currentBlock.Color
	}
	game.nextBlock.ResetPosition()
	game.currentBlock = game.nextBlock
	game.nextBlock = game.GetRandomBlock()
	if !game.BlockFits() {
		rl.PlaySound(game.sounds.GameOver)
		game.GameOver = true
	}
	linesCleared := game.grid.ClearFullRows()
	if linesCleared > 0 {
		rl.PlaySound(game.sounds.ClearLine)
		game.UpdateScore(linesCleared, 0)
	}
}

func (game *Game) BlockFits() bool {
	tiles := game.currentBlock.GetCurrentPositions()
	for _, tile := range tiles {
		if !game.grid.IsCellEmpty(tile.Row, tile.Column) {
			return false
		}
	}
	return true
}

func (g *Game) Reset() {
	g.grid = Grid{}
	g.grid.Speed = 500
	g.blocks = GetAllBlocks()
	g.currentBlock = g.GetRandomBlock()
	g.nextBlock = g.GetRandomBlock()
	g.Score = 0
	g.GameOver = false
}

func (game *Game) UpdateScore(linesCleared int, moveDownPoints int) {
	switch linesCleared {
	case 1:
		game.Score += 100
	case 2:
		game.Score += 300
	case 3:
		game.Score += 500
	}
	game.Score += moveDownPoints
}
