package init

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"mota/maps"
)

func DrawMap(mapInfo [13][13]int, screen *ebiten.Image) {
	for y, rowValue := range mapInfo {
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
