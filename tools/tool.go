package tools

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"mota/maps"
	"runtime"
	"sync"
)

func DrawMap(floorNum int, mapInfo [maps.FloorNum][maps.YNUm][maps.XNUm]int, screen *ebiten.Image) {

	floorInfo := mapInfo[floorNum]
	w := sync.WaitGroup{}
	w.Add(maps.YNUm * maps.XNUm)
	for y, rowValue := range floorInfo {
		for x, imageValue := range rowValue {
			imageValue := imageValue
			x := x
			y := y
			go func() {
				imagePath := fmt.Sprintf("./image/main/resource_%v.png", imageValue)
				img, _, err := ebitenutil.NewImageFromFile(imagePath)
				if err != nil {
					log.Fatal(err)
				}
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*maps.ImageLarge), float64(y*maps.ImageLarge))
				op.GeoM.Scale(1, 1)
				screen.DrawImage(img, op)
				w.Done()
				runtime.GC()
			}()
		}
	}
	w.Wait()
}