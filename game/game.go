package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"mota/draw"
	"mota/init"
	"mota/layout"
	"mota/maps"
	"mota/update"
)

type Game struct {
}

func init() {
	init.Init()
}

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

func Run() {
	ebiten.SetWindowSize(maps.MapWidth, maps.MapHeight)
	ebiten.SetWindowTitle("魔塔")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
