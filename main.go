package main

import (
	"time"

	constants "github.com/fpedroso/golang-raylib-tetris/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var lastUpdateMilli = int64(0)

func EventTriggered(intervalMilli int64) bool {
	timeMilli := time.Now().UnixMilli()
	if timeMilli-lastUpdateMilli >= intervalMilli {
		lastUpdateMilli = timeMilli
		return true
	}
	return false
}

func main() {
	rl.InitWindow((constants.Cols*constants.CellSize)+1, (constants.Rows*constants.CellSize)+1, "Golang Raylib Tetris")
	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() {
		game.HandleInput()
		if EventTriggered(500) {
			game.MoveBlockDown()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkBlue)
		game.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
