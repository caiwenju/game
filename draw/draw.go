package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"mota/image"
	"mota/maps"
	"mota/player"
	"mota/sysconf"
	"mota/tip"
	"mota/tools"
)

var (
	fightDialog tip.FightDialog
)

func init() {
	bgImagePath := "fight/fgt_box.png"
	playerImagePath := "main/resource_0.png"
	masterImagePath := "main/resource_27.png"
	bgImg, _, err := ebitenutil.NewImageFromFileSystem(image.FightFs, bgImagePath)
	playerImg, _, err := ebitenutil.NewImageFromFileSystem(image.MainFs, playerImagePath)
	monsterImg, _, err := ebitenutil.NewImageFromFileSystem(image.MainFs, masterImagePath)
	if err != nil {
		println(err)
	}
	fightDialog = tip.FightDialog{
		X:              1 * maps.ImageLarge,
		Y:              6 * maps.ImageLarge,
		BgImage:        bgImg,
		PlayerImage:    playerImg,
		MonsterImage:   monsterImg,
		PlayerHealth:   10,
		PlayerDefense:  10,
		PlayerAttack:   10,
		MonsterHealth:  10,
		MonsterDefense: 10,
		MonsterAttack:  10,
	}
}

func Draw(screen *ebiten.Image) {

	// 绘制地图
	tools.DrawMap(sysconf.Floor, maps.Map, screen)
	// 绘制人物
	player.Person.LoadImages(screen)
	// 绘制功能面板背景
	tools.DrawBackGround(maps.BackGroundPanel, screen)
	// 绘制功能面板数据
	player.Person.DrawPersonInfo(screen)

	// 绘制战斗信息弹窗
	fightDialog.FightWindow(screen)

}
