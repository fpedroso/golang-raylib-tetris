package main

import (
	"fmt"

	constants "github.com/fpedroso/golang-raylib-tetris/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Cells [constants.Rows][constants.Cols]int
}

func (grid Grid) Print() {
	for i := 0; i < constants.Rows; i++ {
		for j := 0; j < constants.Cols; j++ {
			fmt.Print(grid.Cells[i][j], " ")
		}
		fmt.Println()
	}
}

func (grid Grid) Draw() {
	for row := 0; row < constants.Rows; row++ {
		for column := 0; column < constants.Cols; column++ {
			cellValue := grid.Cells[row][column]
			posX := (column * constants.CellSize) + 1
			posY := (row * constants.CellSize) + 1
			width := constants.CellSize - 1
			height := constants.CellSize - 1
			color := constants.Colors[cellValue]
			rl.DrawRectangle(int32(posX), int32(posY), int32(width), int32(height), color)
		}
	}
}

func (grid Grid) IsCellOutside(row int, column int) bool {
	if row < 0 || row >= constants.Rows {
		return true
	}

	if column < 0 || column >= constants.Cols {
		return true
	}

	return false
}
