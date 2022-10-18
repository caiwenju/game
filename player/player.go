package player

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"math"
	"mota/image"
	"mota/maps"
	"mota/sysconf"
	"strconv"
)

var lastPosition [2]int

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
	Attack    int
	Defense   int
	Health    int
	image     *embed.FS
}

var Person Player

func (p *Player) LoadImages(screen *ebiten.Image) {
	imagePath := fmt.Sprintf("main/resource_%v.png", 0)
	img, _, err := ebitenutil.NewImageFromFileSystem(image.MainFs, imagePath)
	if err != nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.X*maps.ImageLarge), float64(p.Y*maps.ImageLarge))
	op.GeoM.Scale(1, 1)
	screen.DrawImage(img, op)
}

func (p *Player) PlayerStatus() {
	// 检测人物移动
	p.PlayerMove()
}

func (p *Player) PlayerMove() {

	var nextX, nextY int
	var press = false
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		nextX, nextY = p.X, p.Y-1
		press = true
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		nextX, nextY = p.X, p.Y+1
		press = true
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		nextX, nextY = p.X-1, p.Y
		press = true
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		nextX, nextY = p.X+1, p.Y
		press = true
	}
	if press && p.MoveDeal(nextX, nextY) {
		// 移动
		p.X = nextX
		p.Y = nextY
		// 消除被销毁物
		maps.Map[sysconf.Floor][nextY][nextX] = 1
	}
}

func (p *Player) MoveDeal(nextX int, nextY int) bool {

	// 获取地图对象的基础信息
	block := maps.Map[sysconf.Floor][nextY][nextX]
	imageMap := maps.Image[block]
	imageType := imageMap["type"]
	switch imageType {
	case "road":
		// 路
		return true
	case "wall":
		// 判断是否能破墙并进行
		return p.MoveForWall(imageMap)
	case "door":
		//判断是否能开门并进行开门
		return p.MoveForDoor(imageMap)
	case "Key":
		//拾取钥匙
		return p.MoveForAddKey(imageMap)
	case "master":
		// 判断是否能够击杀怪物,并击杀
		return p.MoveForFight(imageMap)
	case "medicine":
		// 增加血量
		return p.MoveForAddHealth(imageMap)
	case "tone":
		// 增加攻击力或者防御力
		return p.MoveForAddAttackOrDefense(imageMap)
	case "stairs":
		// 增切换地图
		return p.MoveForCheckMap(imageMap)
	default:
		return false
	}
}

func (p *Player) MoveForWall(imageMap map[string]interface{}) bool {
	// 人物切换图片
	return false
}

func (p *Player) MoveForDoor(imageMap map[string]interface{}) bool {

	doorColor := imageMap["color"]
	canMove := false
	switch doorColor {
	case "yellow":
		// 判断钥匙数量
		if p.YellowKey >= 1 {
			// 减去钥匙数量
			p.YellowKey -= 1
			canMove = true
		}
	case "blue":
		// 判断钥匙数量
		if p.BlueKey >= 1 {
			// 减去钥匙数量
			p.BlueKey -= 1
			canMove = true
		}
	case "red":
		// 判断钥匙数量
		if p.RedKey >= 1 {
			// 减去钥匙数量
			p.RedKey -= 1
			canMove = true
		}
	}
	return canMove

}

func (p *Player) MoveForFight(imageMap map[string]interface{}) bool {
	canMove := false
	if p.CanKillMaster(imageMap["attack"].(int), imageMap["defense"].(int), imageMap["health"].(int)) {
		// 削减血量 (假定人物先手攻击，回合数向下取整)
		p.Health -= int(math.Floor(float64(imageMap["health"].(int)/(p.Attack-imageMap["defense"].(int))))) * imageMap["attack"].(int)
		canMove = true
	}
	return canMove
}

func (p *Player) MoveForAddHealth(imageMap map[string]interface{}) bool {
	p.Health += imageMap["health"].(int)
	return true
}

func (p *Player) MoveForAddKey(imageMap map[string]interface{}) bool {
	switch imageMap["color"] {
	case "yellow":
		p.YellowKey += imageMap["num"].(int)
	case "blue":
		p.BlueKey += imageMap["num"].(int)
	case "red":
		p.RedKey += imageMap["num"].(int)
	}
	return true
}

func (p *Player) MoveForAddAttackOrDefense(imageMap map[string]interface{}) bool {
	val, ok := imageMap["Attack"]
	if ok {
		p.Attack += val.(int)
	}
	val, ok = imageMap["Defence"]
	if ok {
		p.Defense += val.(int)
	}
	return true
}

func (p *Player) CanKillMaster(masterAttack, masterDefense, masterHealth int) bool {
	if p.Attack > masterDefense && p.Defense >= masterAttack {
		return true
	} else if p.Attack > masterDefense && masterHealth/(p.Attack-masterDefense) <= p.Health/(masterAttack-p.Defense) {
		return true
	}
	return false
}

func (p *Player) MoveForCheckMap(imageMap map[string]interface{}) bool {

	if imageMap["direct"] == "up" {
		sysconf.Floor += 1
		lastPosition[0] = p.X
		lastPosition[1] = p.Y
		p.X, p.Y = maps.PersonPosition[sysconf.Floor][0], maps.PersonPosition[sysconf.Floor][1]
	} else {
		sysconf.Floor -= 1
		p.X, p.Y = lastPosition[0], lastPosition[1]
	}
	return false
}

func (p *Player) DrawPersonInfo(screen *ebiten.Image) {
	// 绘制战斗数字信息
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	normalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	text.Draw(screen, "魔塔", normalFont, 0, 14*maps.ImageLarge-20, color.Black)
	text.Draw(screen, strconv.Itoa(sysconf.Floor), normalFont, maps.ImageLarge+10, 14*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(p.Attack), normalFont, maps.ImageLarge, 15*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(p.Defense), normalFont, maps.ImageLarge, 16*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(p.Health), normalFont, maps.ImageLarge, 17*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(100), normalFont, 8*maps.ImageLarge, 14*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(p.YellowKey), normalFont, 8*maps.ImageLarge, 15*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(p.BlueKey), normalFont, 8*maps.ImageLarge, 16*maps.ImageLarge-15, color.White)
	text.Draw(screen, strconv.Itoa(p.RedKey), normalFont, 8*maps.ImageLarge, 17*maps.ImageLarge-15, color.White)

}

func Init() {
	// 初始化人物信息
	Person = Player{X: maps.PersonPosition[sysconf.Floor][0], Y: maps.PersonPosition[sysconf.Floor][1], YellowKey: 111, BlueKey: 111, RedKey: 111, Attack: 100, Defense: 100, Health: 1000}
}
