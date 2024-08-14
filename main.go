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
	rl.InitWindow((constants.Cols*constants.CellSize)+1+300, (constants.Rows*constants.CellSize)+1, "Nitris")
	rl.SetTargetFPS(60)

	sounds := statics.NewSounds()
	sounds.LoadSounds()
	defer sounds.UnloadSounds()

	rl.PlayMusicStream(sounds.BackgroundMusic)

	font := rl.LoadFontEx("./assets/fonts/Prisma.ttf", 64, nil, 0)

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

		rl.DrawTextEx(font, "Score", rl.NewVector2(655, 15), 38, 2, rl.White)
		rl.DrawRectangleRounded(rl.NewRectangle(615, 55, 270, 60), 0.3, 6, rl.LightGray)
		rl.DrawTextEx(font, strconv.Itoa(game.Score), rl.NewVector2(655, 65), 38, 2, rl.Black)

		rl.DrawTextEx(font, "Next", rl.NewVector2(655, 125), 38, 2, rl.White)
		rl.DrawRectangleRounded(rl.NewRectangle(615, 175, 270, 270), 0.3, 6, rl.LightGray)

		if game.GameOver {
			rl.DrawTextEx(font, "GAME OVER", rl.NewVector2(605, 455), 38, 2, rl.White)
		}

		game.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
