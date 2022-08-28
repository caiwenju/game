package main

import (
	_ "image/png"
	"log"
	. "mota/maps"
	"mota/role"
	"mota/tools"

	"github.com/hajimehoshi/ebiten/v2"
)

var person role.Player

func init() {
	// 初始化任务信息
	person = role.Player{X: 6, Y: 11}
}

type Game struct {
}

func (g *Game) Update() error {

	// 处理人物移动
	person.PlayerStatus()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tools.DrawMap(1, Map, screen)
	person.LoadImages(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MapWidth, MapHeight
}

func main() {
	ebiten.SetWindowSize(MapWidth, MapHeight)
	ebiten.SetWindowTitle("Geometry Matrix")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
