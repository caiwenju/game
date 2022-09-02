package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	. "mota/maps"
)

type MapDeal struct{}

func (m MapDeal) drawMap(mapInfo [13][13]int) []fyne.CanvasObject {
	var imageArray []fyne.CanvasObject
	for _, rowValue := range mapInfo {
		for _, value := range rowValue {
			imagePath := fmt.Sprintf("../image/main/resource_%v.png", value)
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
	Map := [13][13]int{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 8, 1, 20, 21, 20, 1, 1, 1, 1, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2},
		{2, 83, 1, 1, 5, 1, 2, 85, 80, 1, 2, 1, 2},
		{2, 1, 27, 1, 2, 1, 2, 86, 83, 1, 2, 1, 2},
		{2, 2, 5, 2, 2, 1, 2, 2, 2, 5, 2, 1, 2},
		{2, 80, 1, 1, 2, 1, 5, 24, 30, 24, 2, 1, 2},
		{2, 1, 28, 1, 2, 1, 2, 2, 2, 2, 2, 1, 2},
		{2, 2, 5, 2, 2, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 2, 2, 5, 2, 2, 2, 5, 2, 2},
		{2, 83, 1, 80, 2, 80, 1, 1, 2, 1, 24, 1, 2},
		{2, 83, 1, 80, 2, 1, 1, 1, 2, 20, 84, 20, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	}
	m := MapDeal{}
	imageArray := m.drawMap(Map)
	cont := container.New(layout.NewGridLayout(13), imageArray...)
	w.SetContent(cont)
	w.ShowAndRun()

}
