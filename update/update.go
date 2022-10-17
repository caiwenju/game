package update

import (
	"mota/player"
	"mota/sounds"
)

func Update() {

	// 处理人物状态
	player.Person.PlayerStatus()

	// 播放背景音乐
	sounds.BackGround.Play()

}
