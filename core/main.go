package core

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	// ScreenWidth is the width of the game screen
	ScreenWidth = 640
	// ScreenHeight is the height of the game screen
	ScreenHeight = 480
)

// CharacterImage is the image status container
type characterImage struct {
	image       *ebiten.Image
	frame0X     int
	frame0Y     int
	frameWidth  int
	frameHeight int
	frameCount  int
	drawOpt     *ebiten.DrawImageOptions
}

const (
	// FrameOX is the X offset of the first frame
	runnerFrameOX = 0
	// FrameOY is the Y offset of the first frame
	runnerFrameOY = 32
	// FrameWidth is the width of each frame
	runnerFrameWidth = 32
	// FrameHeight is the height of each frame
	runnerFrameHeight = 32
	// FrameCount is the number of frames
	runnerFrameCount = 8
)

// NewCharacterImage is the constructor for CharacterImage
func NewCharacterImage(img *ebiten.Image, frame0X, frame0Y, frameWidth, frameHeight, frameCount int) *characterImage {
	drawOpt := &ebiten.DrawImageOptions{}
	drawOpt.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	drawOpt.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)

	return &characterImage{
		image:       img,
		frame0X:     frame0X,
		frame0Y:     frame0Y,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		frameCount:  frameCount,
		drawOpt:     drawOpt,
	}
}

const (
	mvInvalid int = iota // 0
	mvStay               // 1
	mvJump               // 2
)

const (
	healthMax int = 100
	energyMax int = 100
)

// characterStatus is the character status container
type characterStatus struct {
	moving int
	health int
	energy int
}

// NewCharacterStatus is the constructor for CharacterStatus
func NewCharacterStatus() *characterStatus {
	return &characterStatus{
		moving: mvStay,
		health: healthMax,
		energy: energyMax,
	}
}

// Runner is the player character
type Runner struct {
	chImage   *characterImage
	positionX int
	positionY int
	pause     bool
	status    *characterStatus
	drawOpt   *ebiten.DrawImageOptions
}

// getCurFrameImage returns the current frame image
func (r *Runner) getCurFrameImage(frameCount int) *ebiten.Image {
	i := (frameCount / 5) % r.chImage.frameCount
	sx, sy := r.chImage.frame0X+i*r.chImage.frameWidth, r.chImage.frame0Y
	return r.chImage.image.SubImage(image.Rect(sx, sy, sx+r.chImage.frameWidth, sy+r.chImage.frameHeight)).(*ebiten.Image)
}

// getRunnerImage is the constructor for Runner Image
func getRunnerImage() *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage := ebiten.NewImageFromImage(img)
	return runnerImage
}

// Game is the core game object
// It implements the ebiten.Game interface
type Game struct {
	count  int
	keys   []ebiten.Key
	runner *Runner
}

// NewGame is the constructor for Game
func NewGame() *Game {
	chImage := NewCharacterImage(
		getRunnerImage(),
		runnerFrameOX,
		runnerFrameOY,
		runnerFrameWidth,
		runnerFrameHeight,
		runnerFrameCount,
	)
	status := NewCharacterStatus()

	return &Game{
		count: 0,
		keys:  []ebiten.Key{},
		runner: &Runner{
			chImage:   chImage,
			positionX: 0,
			positionY: 0,
			pause:     false,
			status:    status,
		},
	}
}

// Update is called every frame
func (g *Game) Update() error {
	g.count++
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

// Draw is called every frame
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the game title and frame count
	lineToPrint := fmt.Sprintf("Runner Animation Demo %d\n", g.count)
	ebitenutil.DebugPrint(screen, lineToPrint)

	// Draw the pressed keys
	keyStrs := []string{}
	for _, key := range g.keys {
		keyStrs = append(keyStrs, key.String())
	}
	linePressedKeys := fmt.Sprintf("\nPressed Keys: %s", strings.Join(keyStrs, ", "))
	ebitenutil.DebugPrint(screen, linePressedKeys)

	// Draw the runner
	screen.DrawImage(g.runner.getCurFrameImage(g.count), g.runner.chImage.drawOpt)
}

// Layout is called every frame
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
