package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sota0121/ebitengine-game/core"
)

func execSampleGame() {
	// Set up the game window
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sample Game Title")
	if err := ebiten.RunGame(&core.SampleGame{}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execSampleGame()
}
