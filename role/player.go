package role

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	. "mota/maps"
)

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
	op.GeoM.Translate(float64(p.X*ImageLarge), float64(p.Y*ImageLarge))
	op.GeoM.Scale(1, 1)
	screen.DrawImage(img, op)
}

func (p *Player) PlayerStatus() {
	// 检测人物移动
	p.PlayerMove()
	// 检测
}

func (p *Player) PlayerMove() {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		nextX, nextY := p.X, p.Y-1
		p.MoveDeal(nextX, nextY)
		p.Y -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		nextX, nextY := p.X, p.Y+1
		p.MoveDeal(nextX, nextY)
		p.Y += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		nextX, nextY := p.X-1, p.Y
		p.MoveDeal(nextX, nextY)
		p.X -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		nextX, nextY := p.X+1, p.Y
		p.MoveDeal(nextX, nextY)
		p.X += 1
	}
}

func (p *Player) MoveDeal(nextX int, nextY int) {

	// 获取地图对象的基础信息
	block := Map[1][nextX][nextY]
	imageMap := Image[block]
	imageType := imageMap["type"]
	switch imageType {
	case "wall":
		// 判断是否能破墙并进行
		p.MoveForWall(imageMap, nextX, nextY)
	case "door":
		//判断是否能开门并进行开门
		p.MoveForDoor(imageMap, nextX, nextY)
	}

}

func (p *Player) MoveForWall(imageMap map[string]interface{}, nextX, nextY int) {
	// 人物切换图片
	fmt.Println("人物切换图片")
}

func (p *Player) MoveForDoor(imageMap map[string]interface{}, nextX, nextY int) {

	doorType := imageMap["type"]
	switch doorType {

	case doorType == "yellowDoor":
		// 判断钥匙数量
		if p.YellowKey >= 1 {

			// 更换门的图片
			CurrentMap[1][nextX][nextY] = 1
			// 减去钥匙数量
			p.YellowKey -= 1
		}
	case doorType == "blueDoor":
		// 判断钥匙数量
		if p.BlueKey >= 1 {
			// 更换门的图片
			CurrentMap[1][nextX][nextY] = 1
			// 减去钥匙数量
			p.YellowKey -= 1
		}
	case doorType == "RedDoor":
		// 判断钥匙数量
		if p.RedKey >= 1 {
			// 更换门的图片
			CurrentMap[1][nextX][nextY] = 1
			// 减去钥匙数量
			p.YellowKey -= 1
		}
	}

}
