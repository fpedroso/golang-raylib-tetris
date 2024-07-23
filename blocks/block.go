package blocks

import (
	constants "github.com/fpedroso/golang-raylib-tetris/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Color     int
	state     int
	cells     [4][4]Position
	rowOffset int
	colOffset int
}

func NewBlock(color int) Block {
	b := Block{Color: color}
	b.colOffset = 3
	return b
}

func (block *Block) Rotate() {
	block.state++
	if block.state > 3 {
		block.state = 0
	}
}

func (block *Block) UndoRotate() {
	block.state--
	if block.state < 0 {
		block.state = 3
	}
}

func (block Block) Draw() {
	positions := block.GetCurrentPositions()
	for _, position := range positions {
		posX := (position.Column * constants.CellSize) + 1
		posY := (position.Row * constants.CellSize) + 1
		width := constants.CellSize - 1
		height := constants.CellSize - 1
		color := constants.Colors[block.Color]
		rl.DrawRectangle(int32(posX), int32(posY), int32(width), int32(height), color)
	}
}

func (block *Block) Move(rows int, columns int) {
	block.rowOffset += rows
	block.colOffset += columns
}

func (block Block) GetCurrentPositions() []Position {
	positions := block.cells[block.state]
	movedPositions := make([]Position, 0)
	for _, tile := range positions {
		movedPositions = append(movedPositions, Position{
			Row:    tile.Row + block.rowOffset,
			Column: tile.Column + block.colOffset,
		})
	}
	return movedPositions
}
