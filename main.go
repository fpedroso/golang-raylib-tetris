package main

import (
	constants "github.com/fpedroso/golang-raylib-tetris/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow((constants.Cols*constants.CellSize)+1, (constants.Rows*constants.CellSize)+1, "Golang Raylib Tetris")
	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkBlue)
		game.HandleInput()
		game.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
