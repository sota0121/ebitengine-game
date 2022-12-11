package core

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game is the core game object
// It implements the ebiten.Game interface
type Game struct {
	count int
}

// Update is called every frame
func (g *Game) Update() error {
	g.count++
	return nil
}

// Draw is called every frame
func (g *Game) Draw(screen *ebiten.Image) {
	lineToPrint := fmt.Sprintf("Runner Animation Demo %d", g.count)
	ebitenutil.DebugPrint(screen, lineToPrint)
}

// Layout is called every frame
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}
