package rope

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	ropeSize     = 250
)

type Game struct {
	stage *Stage
}

func NewGame() (*Game, error) {
	g := &Game{
		stage: &Stage{NewRope(ropeSize)},
	}
	g.stage.Rope.makeWave()

	var err error

	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.stage.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stage.Draw(screen)
}
