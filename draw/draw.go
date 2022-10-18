package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"mota/maps"
	"mota/player"
	"mota/sysconf"
	"mota/tip"
	"mota/tools"
)

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
	tip.TipsWindow(screen)
}
