package tools

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"mota/maps"
)

func DrawMap(floorNum int, mapInfo [maps.FloorNum][maps.YNUm][maps.XNUm]int, screen *ebiten.Image) {

	floorInfo := mapInfo[floorNum]
	for y, rowValue := range floorInfo {
		for x, value := range rowValue {
			imagePath := fmt.Sprintf("./image/main/resource_%v.png", value)
			img, _, err := ebitenutil.NewImageFromFile(imagePath)
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*maps.ImageLarge), float64(y*maps.ImageLarge))
			op.GeoM.Scale(1, 1)
			screen.DrawImage(img, op)
		}
	}
}
