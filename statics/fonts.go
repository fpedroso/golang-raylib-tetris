package statics

import rl "github.com/gen2brain/raylib-go/raylib"

type Fonts struct {
	Regular rl.Font
}

func NewFonts() Fonts {
	fonts := Fonts{}
	return fonts
}

func (fonts *Fonts) LoadFonts() {
	fonts.Regular = rl.LoadFontEx("./assets/fonts/Prisma.ttf", 64, nil, 0)
}

func (fonts *Fonts) UnloadFonts() {
	rl.UnloadFont(fonts.Regular)
}
