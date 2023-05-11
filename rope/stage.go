package rope

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Stage struct {
	Rope *Rope
}

func (s *Stage) Update() {
	s.Rope.updateVectors()
	s.Rope.updateBlocksPos()
}

func (s *Stage) Draw(stageScreen *ebiten.Image) {
	for _, v := range s.Rope.blocks {
		ebitenutil.DrawRect(stageScreen, float64(v.X), ScreenHeight/2+float64(v.Y), BlockWidth, BlockWidth, color.NRGBA{0x80, 0xff, 0x80, 0x80})
	}
}
