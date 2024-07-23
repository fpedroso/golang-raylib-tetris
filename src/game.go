package main

import (
	"math/rand"

	blocks "github.com/fpedroso/golang-raylib-tetris/blocks"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	grid          Grid
	blocks        []blocks.Block
	currentBlock  blocks.Block
	nextBlock     blocks.Block
	GameOver      bool
	Score         int
	ClearSound    rl.Sound
	GameOverSound rl.Sound
}

func NewGame() Game {
	g := Game{}
	g.grid = Grid{}
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
	game.nextBlock.Draw(450, 270)
}

func (game *Game) HandleInput() {
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
		game.MoveBlockDown()
		game.UpdateScore(0, 1)
	case rl.IsKeyPressed(rl.KeyUp):
		game.RotateBlock()
	}
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
	if game.IsBlockOutside() || !game.BlockFits() {
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

func (game *Game) LockBlock() {
	if game.GameOver {
		return
	}
	tiles := game.currentBlock.GetCurrentPositions()
	for _, tile := range tiles {
		game.grid.Cells[tile.Row][tile.Column] = game.currentBlock.Color
	}
	game.currentBlock = game.nextBlock
	game.nextBlock = game.GetRandomBlock()
	if !game.BlockFits() {
		rl.PlaySound(game.GameOverSound)
		game.GameOver = true
	}
	linesCleared := game.grid.ClearFullRows()
	if linesCleared > 0 {
		rl.PlaySound(game.ClearSound)
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
