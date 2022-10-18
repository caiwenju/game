package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"mota/game"
	_ "mota/init"
	"mota/maps"
)

func main() {
	newGame := game.NewGame()
	ebiten.SetWindowSize(maps.MapWidth, maps.MapHeight)
	ebiten.SetWindowTitle("魔塔")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(newGame); err != nil {
		log.Fatal(err)
	}
}
