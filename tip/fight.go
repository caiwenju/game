package tip

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	"log"
	"mota/maps"
	"strconv"
)

type FightDialog struct {
	X, Y           int
	BgImage        *ebiten.Image
	PlayerImage    *ebiten.Image
	MonsterImage   *ebiten.Image
	PlayerHealth   int
	PlayerDefense  int
	PlayerAttack   int
	MonsterHealth  int
	MonsterDefense int
	MonsterAttack  int
}

func (f *FightDialog) FightWindow(screen *ebiten.Image) {

	// 绘制背景图片
	bgOptions := &ebiten.DrawImageOptions{}
	bgOptions.GeoM.Translate(float64(f.X), float64(f.Y))
	// 背景图片过大，裁剪图像的坐标和尺寸
	width, height := 10*maps.ImageLarge, 5*maps.ImageLarge
	// 使用SubImage方法裁剪图像
	croppedImage := ebiten.NewImageFromImage(f.BgImage.SubImage(image.Rect(0, 0, width, height)).(*ebiten.Image))
	screen.DrawImage(croppedImage, bgOptions)

	// 设置字体大小和样式
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	normalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	// 绘制怪物图片
	MonsterOptions := &ebiten.DrawImageOptions{}
	MonsterOptions.GeoM.Translate(float64(f.X+20), float64(f.Y+maps.ImageLarge*2))
	screen.DrawImage(f.MonsterImage, MonsterOptions)
	// 绘制怪物数值信息
	fmt.Println(rune(f.MonsterHealth))
	text.Draw(screen, "血量: "+strconv.Itoa(f.MonsterHealth), normalFont, f.X+maps.ImageLarge*2, f.Y+maps.ImageLarge*2, color.White)
	text.Draw(screen, "攻力: "+strconv.Itoa(f.MonsterAttack), normalFont, f.X+maps.ImageLarge*2, f.Y+maps.ImageLarge*3, color.White)
	text.Draw(screen, "防御力: "+strconv.Itoa(f.MonsterDefense), normalFont, f.X+maps.ImageLarge*2, f.Y+maps.ImageLarge*4, color.White)

	// 绘制玩家图片
	playerOptions := &ebiten.DrawImageOptions{}
	playerOptions.GeoM.Translate(float64(f.X+maps.ImageLarge*8+27), float64(f.Y+maps.ImageLarge*2))
	screen.DrawImage(f.PlayerImage, playerOptions)
	// 绘制玩家数值信息
	text.Draw(screen, strconv.Itoa(f.PlayerHealth)+" :血量", normalFont, f.X+maps.ImageLarge*7-27, f.Y+maps.ImageLarge*2, color.White)
	text.Draw(screen, strconv.Itoa(f.PlayerAttack)+" :攻力", normalFont, f.X+maps.ImageLarge*7-27, f.Y+maps.ImageLarge*3, color.White)
	text.Draw(screen, strconv.Itoa(f.PlayerDefense)+" :防御力", normalFont, f.X+maps.ImageLarge*7-27, f.Y+maps.ImageLarge*4, color.White)

}
