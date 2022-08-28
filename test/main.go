package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	. "mota/maps"
)

type MapDeal struct{}

func (m MapDeal) drawMap(mapInfo [13][13]int) []fyne.CanvasObject {
	var imageArray []fyne.CanvasObject
	for _, rowValue := range mapInfo {
		for _, value := range rowValue {
			imagePath := fmt.Sprintf("./image/main/resource_%v.png", value)
			img := canvas.NewImageFromFile(imagePath)
			img.FillMode = canvas.ImageFillStretch
			img.Resize(fyne.NewSize(ImageLarge, ImageLarge))
			imageArray = append(imageArray, img)
		}
	}
	return imageArray
}

func main() {

	// 创建窗口
	a := app.New()
	w := a.NewWindow("魔塔")
	w.Resize(fyne.NewSize(MapWidth, MapHeight))
	w.CenterOnScreen()

	// 绘制初始地图
	//m := MapDeal{}
	//imageArray := m.drawMap(Map)
	//cont := container.New(layout.NewGridLayout(13), imageArray...)
	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	red := color.NRGBA{R: 0xff, A: 0xff}
	obj2 := canvas.NewRectangle(red)
	obj2.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))
	w.SetContent(container.NewWithoutLayout(obj2))
	//cont := container.New(layout.NewGridLayout(13))
	//cont := container.New(layout.NewGridLayout(13), imageArray...)

	//w.SetContent(cont)
	w.ShowAndRun()

}
