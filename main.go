package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	_ "image/png"
	"io"
	"log"
	"math"
	. "mota/maps"
	"mota/tools"
	"os"
	"strconv"
)

var CurrentMap [FloorNum][YNUm][XNUm]int

var Floor = 1

var lastPosition [2]int

type Game struct{}

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
		CurrentMap[Floor][nextY][nextX] = 1
	}
}

func (p *Player) MoveDeal(nextX int, nextY int) bool {

	// 获取地图对象的基础信息
	block := CurrentMap[Floor][nextY][nextX]
	imageMap := Image[block]
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
		Floor += 1
		lastPosition[0] = p.X
		lastPosition[1] = p.Y
		p.X, p.Y = PersonPosition[Floor][0], PersonPosition[Floor][1]
	} else {
		Floor -= 1
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
	text.Draw(screen, "魔塔", normalFont, 0, 14*ImageLarge-20, color.Black)
	text.Draw(screen, strconv.Itoa(Floor), normalFont, ImageLarge+10, 14*ImageLarge-20, color.White)
	text.Draw(screen, strconv.Itoa(p.Attack), normalFont, ImageLarge, 15*ImageLarge-20, color.White)
	text.Draw(screen, strconv.Itoa(p.Defense), normalFont, ImageLarge, 16*ImageLarge-20, color.White)
	text.Draw(screen, strconv.Itoa(p.Health), normalFont, ImageLarge, 17*ImageLarge-20, color.White)

}

var person Player

func init() {
	// 	初始化地图信息
	CurrentMap = Map

	// 初始化人物信息
	person = Player{X: PersonPosition[Floor][0], Y: PersonPosition[Floor][1], YellowKey: 111, BlueKey: 111, RedKey: 111, Attack: 100, Defense: 100, Health: 1000}
}

type File struct {
}

type ReadSeeker struct {
	io.ReadSeeker
}

func (f File) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (f File) re() []byte {
	// 打开文件
	file, err := os.ReadFile("./sounds/bgm1.ogg")
	if err != nil {
		fmt.Println("ReadFile err :", err)
		return []byte{}
	}
	return file
}

func (g *Game) Update() error {

	// 处理人物状态
	person.PlayerStatus()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// 绘制地图
	tools.DrawMap(Floor, CurrentMap, screen)
	// 绘制人物
	person.LoadImages(screen)
	// 绘制功能面板背景
	tools.DrawBackGround(BackGroundPanel, screen)
	// 绘制功能面板数据
	person.DrawPersonInfo(screen)
	//// 播放背景音乐
	//var reader io.Reader
	//f := File{}
	//reader = f
	//_, _ = mp3.DecodeWithSampleRate(44100, reader)
	//bgmStream, _ := reader.Read(f.re())
	//s := audio.NewInfiniteLoop(1, int64(bgmStream))
	//d := audio.Player{}
	//d.Play()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MapWidth, MapHeight
}

func main() {
	ebiten.SetWindowSize(MapWidth, MapHeight)
	ebiten.SetWindowTitle("魔塔")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum)
	mobile.SetGame(&Game{})
}
