package tip

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func TipsWindow(screen *ebiten.Image) {
	var emptyImage = ebiten.NewImage(3, 3)
	emptyImage.Fill(color.White)
	op := &ebiten.DrawRectShaderOptions{}
	op.GeoM.Scale(200, 200)
	screen.DrawRectShader(100, 100, shader, op)
}
