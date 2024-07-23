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

func (grid Grid) IsCellEmpty(row int, column int) bool {
	return grid.Cells[row][column] == 0
}

func (grid Grid) IsRowFull(row int) bool {
	for col := range constants.Cols {
		if grid.IsCellEmpty(row, col) {
			return false
		}
	}
	return true
}

func (grid *Grid) ClearRow(row int) {
	for col := range constants.Cols {
		grid.Cells[row][col] = 0
	}
}

func (grid *Grid) MoveRowDown(row int, numRows int) {
	for col := range constants.Cols {
		grid.Cells[row+numRows][col] = grid.Cells[row][col]
		grid.Cells[row][col] = 0
	}
}

func (grid *Grid) ClearFullRows() int {
	completed := 0
	for row := constants.Rows - 1; row >= 0; row-- {
		if grid.IsRowFull(row) {
			grid.ClearRow(row)
			completed++
		} else if completed > 0 {
			grid.MoveRowDown(row, completed)
		}
	}
	return completed
}
