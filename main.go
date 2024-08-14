package main

import (
	"strconv"
	"time"

	constants "github.com/fpedroso/golang-raylib-tetris/constants"
	statics "github.com/fpedroso/golang-raylib-tetris/statics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var lastUpdateMilli = int64(0)

type TimeHandler func(int64) bool

func EventTriggered(intervalMilli int64) bool {
	timeMilli := time.Now().UnixMilli()
	if timeMilli-lastUpdateMilli >= intervalMilli {
		lastUpdateMilli = timeMilli
		return true
	}
	return false
}

func main() {
	rl.InitWindow((constants.ColCount*constants.CellSize)+constants.SidePanelWidth+1, (constants.RowCount*constants.CellSize)+1, constants.TextGameName)
	rl.SetTargetFPS(constants.FramesPerSecond)

	sounds := statics.NewSounds()
	sounds.LoadSounds()
	defer sounds.UnloadSounds()

	fonts := statics.NewFonts()
	fonts.LoadFonts()
	defer fonts.UnloadFonts()

	rl.PlayMusicStream(sounds.BackgroundMusic)
	game := NewGame(sounds)

	for !rl.WindowShouldClose() {
		if !game.GameOver {
			rl.UpdateMusicStream(sounds.BackgroundMusic)
		}
		game.HandleInput(EventTriggered)
		if EventTriggered(game.grid.Speed) {
			game.MoveBlockDown()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkBlue)

		rl.DrawTextEx(fonts.Regular, constants.TextScore, rl.NewVector2(655, 15), 38, 2, rl.White)
		rl.DrawRectangleRounded(rl.NewRectangle(615, 55, 270, 60), 0.3, 6, rl.LightGray)
		rl.DrawTextEx(fonts.Regular, strconv.Itoa(game.Score), rl.NewVector2(655, 65), 38, 2, rl.Black)

		rl.DrawTextEx(fonts.Regular, constants.TextNext, rl.NewVector2(655, 125), 38, 2, rl.White)
		rl.DrawRectangleRounded(rl.NewRectangle(615, 175, 270, 270), 0.3, 6, rl.LightGray)

		if game.GameOver {
			rl.DrawTextEx(fonts.Regular, constants.TextGameOver, rl.NewVector2(605, 455), 38, 2, rl.White)
		}

		game.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
