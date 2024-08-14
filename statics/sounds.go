package statics

import rl "github.com/gen2brain/raylib-go/raylib"

type Sounds struct {
	BackgroundMusic rl.Music
	ClearLine       rl.Sound
	GameOver        rl.Sound
}

func NewSounds() Sounds {
	sounds := Sounds{}
	return sounds
}

func (sounds *Sounds) LoadSounds() {
	rl.InitAudioDevice()
	sounds.BackgroundMusic = rl.LoadMusicStream("./assets/sounds/levelmusic.wav")
	sounds.ClearLine = rl.LoadSound("./assets/sounds/lineclear.wav")
	sounds.GameOver = rl.LoadSound("./assets/sounds/gameover.wav")
}

func (sounds *Sounds) UnloadSounds() {
	rl.CloseAudioDevice()
	rl.UnloadMusicStream(sounds.BackgroundMusic)
	rl.UnloadSound(sounds.ClearLine)
	rl.UnloadSound(sounds.GameOver)
}
