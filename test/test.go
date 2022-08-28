package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"log"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func windowsControlExp() {
	/*
		控制多窗口的示例
	*/
	a := app.New()
	w := a.NewWindow("Clock")
	w.Resize(fyne.NewSize(800, 800))
	w.SetMaster()

	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w2 := a.NewWindow("w2")
	w2.Resize(fyne.NewSize(500, 600))
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
	w2.Show()
	w.Show()
	a.Run()
}

func CanvasExp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	go func() {
		setContentToText(myCanvas)
	}()

	go func() {
		setContentToCircle(myCanvas)
	}()

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

func newEntry() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")
	myWindow.Resize(fyne.NewSize(700, 700))
	myWindow.SetContent(widget.NewEntry())
	myWindow.ShowAndRun()
}

func containerExp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2)
	//content := container.New(layout.NewGridLayout(2), text1, text2)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func dbExp() {
	a := app.NewWithID("com.example.tutorial.preferences")
	w := a.NewWindow("Timeout")

	var timeout time.Duration

	timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
		switch selected {
		case "10 seconds":
			timeout = 10 * time.Second
		case "30 seconds":
			timeout = 30 * time.Second
		case "1 minute":
			timeout = time.Minute
		}

		a.Preferences().SetString("AppTimeout", selected)
	})

	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))

	go func() {
		time.Sleep(timeout)
		a.Quit()
	}()

	w.SetContent(timeoutSelector)
	w.ShowAndRun()
}

func backRun() {
	a := app.New()
	w := a.NewWindow("SysTray")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.ShowAndRun()
}

func dataBind() {
	a := app.New()
	w := a.NewWindow("Hello")

	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()

	w.SetContent(widget.NewLabelWithData(str))
	w.ShowAndRun()
}

func imageExp() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	//image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal // 图片比例不可以变化
	//image.FillMode = canvas.ImageFillContain // 随窗口放大缩小，但是位置是永远按比例
	//image.FillMode = canvas.ImageFillStretch // 随窗口放大缩小
	w.SetContent(image)

	w.ShowAndRun()
}

func Raster() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	raster := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)), 0xff}
		})
	// raster := canvas.NewRasterFromImage()
	w.SetContent(raster)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}

func Gradient() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")

	//gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}

func cartoon() {
	a := app.New()
	w := a.NewWindow("Hello")

	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}
	canvas.NewColorRGBAAnimation(red, blue, time.Second*4, func(c color.Color) {
		obj.FillColor = c
		canvas.Refresh(obj)
	}).Start()
	move := canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(200, 0), time.Second, obj.Move)
	move.AutoReverse = true
	move.Start()
	w.Resize(fyne.NewSize(250, 50))
	w.SetPadded(false)
	w.ShowAndRun()
}

func Grid() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	text4 := canvas.NewText("4", color.Black)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3, text4)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

func Layout() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}

func Border() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.Black)
	left := canvas.NewText("left", color.Black)
	bottom := canvas.NewText("left", color.Black)
	middle := canvas.NewText("content", color.Black)
	content := container.New(layout.NewBorderLayout(top, bottom, left, nil),
		top, left, bottom, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func Center() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Center Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func Max() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewMaxLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func TabContainer() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func Toolbar() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar Widget")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func main() {
	//windowsControlExp()
	//CanvasExp()
	//newEntry()
	//containerExp()
	//dbExp()
	//backRun()
	//dataBind()
	//imageExp()
	//Raster()
	//Gradient()
	cartoon()
	//Grid()
	//Layout()
	//Border()
	//Center()
	//Max()
	//TabContainer() // 楼层飞行器
	//Toolbar() // 各种道具
}
