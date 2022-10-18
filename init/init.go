package init

import (
	"mota/player"
	"mota/sounds"
)

func init() {
	// 初始化人物信息
	player.Init()

	// 初始化背景音乐
	sounds.Init()
}
