package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"mota/draw"
	"mota/layout"
	"mota/update"
)

type Game struct{}

func (g *Game) Update() error {
	update.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return layout.Layout()
}

func NewGame() *Game {
	g := &Game{}
	return g
}
