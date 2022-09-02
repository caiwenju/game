package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
	"log"
	. "mota/maps"
	"mota/tools"
)

var CurrentMap [FloorNum][YNUm][XNUm]int

type Game struct{}

func init() {
	// 	初始化地图信息
	CurrentMap = Map

	// 初始化人物信息
	person = Player{X: 11, Y: 6, YellowKey: 4, BlueKey: 1, RedKey: 1}
}

type Player struct {
	X         int
	Y         int
	State     int
	Direction int
	MouseX    int
	MouseY    int
	SkillName string
	YellowKey int
	BlueKey   int
	RedKey    int
	image     *embed.FS
}

func (p *Player) LoadImages(screen *ebiten.Image) {
	imagePath := fmt.Sprintf("./image/main/resource_%v.png", 0)
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Y*ImageLarge), float64(p.X*ImageLarge))
	op.GeoM.Scale(1, 1)
	screen.DrawImage(img, op)
}

func (p *Player) PlayerStatus() {
	// 检测人物移动
	p.PlayerMove()
	// 检测
}

func (p *Player) PlayerMove() {

	var nextX, nextY = p.X, p.Y
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		nextX, nextY = p.X-1, p.Y
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		nextX, nextY = p.X+1, p.Y
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		nextX, nextY = p.X, p.Y-1
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		nextX, nextY = p.X, p.Y+1
	}
	if p.MoveDeal(nextX, nextY) {
		// 移动
		p.X = nextX
		p.Y = nextY
		// 消除被销毁物
		CurrentMap[1][nextX][nextY] = 1
	}

}

func (p *Player) MoveDeal(nextX int, nextY int) bool {

	// 获取地图对象的基础信息
	block := Map[1][nextX][nextY]
	imageMap := Image[block]
	imageType := imageMap["type"]
	switch imageType {
	case "wall":
		// 判断是否能破墙并进行
		return p.MoveForWall(imageMap)
	case "door":
		//判断是否能开门并进行开门
		return p.MoveForDoor(imageMap)
	default:
		return true
	}
}

func (p *Player) MoveForWall(imageMap map[string]interface{}) bool {
	// 人物切换图片
	return false
}

func (p *Player) MoveForDoor(imageMap map[string]interface{}) bool {

	doorType := imageMap["type"]
	canMove := false
	fmt.Printf("剩余黄钥匙数：%v", p.YellowKey)
	switch doorType {
	case doorType == "yellowDoor":
		// 判断钥匙数量
		if p.YellowKey >= 1 {
			// 减去钥匙数量
			p.YellowKey -= 1
			canMove = true
		}
	case doorType == "blueDoor":
		// 判断钥匙数量
		if p.BlueKey >= 1 {
			// 减去钥匙数量
			p.BlueKey -= 1
			canMove = true
		}
	case doorType == "RedDoor":
		// 判断钥匙数量
		if p.RedKey >= 1 {
			// 减去钥匙数量
			p.RedKey -= 1
			canMove = true
		}
	}
	fmt.Printf("canMove：%v", canMove)
	return canMove

}

var person Player

func (g *Game) Update() error {

	// 处理人物状态
	person.PlayerStatus()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tools.DrawMap(1, CurrentMap, screen)
	person.LoadImages(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MapWidth, MapHeight
}

func main() {
	ebiten.SetWindowSize(MapWidth, MapHeight)
	ebiten.SetWindowTitle("Geometry Matrix")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
