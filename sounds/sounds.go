package sounds

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var (
	//go:embed bgm1.ogg
	Bgm1Ogg []byte
)

var BackGround *audio.Player
var audioContext *audio.Context

const (
	sampleRate     = 48000 // 采样率
	bytesPerSample = 4     // 每次采样读取的字节数

	introLengthInSecond = 0  // 音频开始的秒数
	loopLengthInSecond  = 58 // 音频循环播放的秒数，bytesPerSample 会影响他
)

func Init() {
	// 初始化背景音乐
	audioContext = audio.NewContext(sampleRate)
	Bgm1Ogg, _ := vorbis.DecodeWithoutResampling(bytes.NewReader(Bgm1Ogg))
	loop := audio.NewInfiniteLoopWithIntro(Bgm1Ogg, introLengthInSecond*bytesPerSample*sampleRate, loopLengthInSecond*bytesPerSample*sampleRate)
	BackGround, _ = audioContext.NewPlayer(loop)
}
