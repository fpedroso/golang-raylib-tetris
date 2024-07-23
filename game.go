package main

import (
	"math/rand"

	blocks "github.com/fpedroso/golang-raylib-tetris/blocks"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	grid         Grid
	blocks       []blocks.Block
	currentBlock blocks.Block
	nextBlock    blocks.Block
}

func NewGame() Game {
	g := Game{}
	g.grid = Grid{}
	g.blocks = GetAllBlocks()
	g.currentBlock = g.GetRandomBlock()
	g.nextBlock = g.GetRandomBlock()
	return g
}

func (game Game) GetRandomBlock() blocks.Block {
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
	game.currentBlock.Draw()
}

func (game *Game) HandleInput() {
	keyPressed := rl.GetKeyPressed()
	switch keyPressed {
	case rl.KeyLeft:
		game.MoveBlockLeft()
	case rl.KeyRight:
		game.MoveBlockRight()
	case rl.KeyDown:
		game.MoveBlockDown()
	case rl.KeyUp:
		game.RotateBlock()
	}
}

func (game *Game) MoveBlockLeft() {
	game.currentBlock.Move(0, -1)
	if game.IsBlockOutside() {
		game.currentBlock.Move(0, 1)
	}
}

func (game *Game) MoveBlockRight() {
	game.currentBlock.Move(0, 1)
	if game.IsBlockOutside() {
		game.currentBlock.Move(0, -1)
	}
}

func (game *Game) MoveBlockDown() {
	game.currentBlock.Move(1, 0)
	if game.IsBlockOutside() {
		game.currentBlock.Move(-1, 0)
	}
}

func (game *Game) RotateBlock() {
	game.currentBlock.Rotate()
	for game.IsBlockOutside() {
		game.currentBlock.Rotate()
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
