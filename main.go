package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"mota/game"
	_ "mota/init"
	"mota/maps"
)

func main() {
	ebiten.SetWindowSize(maps.MapWidth, maps.MapHeight)
	ebiten.SetWindowTitle("魔塔")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(game.G); err != nil {
		log.Fatal(err)
	}
}
