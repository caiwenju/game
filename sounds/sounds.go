package sounds

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var (
	BackGround *audio.Player
)

var (
	MusicObj Music
)

var audioContext *audio.Context

var (
	//go:embed bgm1.ogg
	Bgm1OggFile []byte // 背景音乐
	//go:embed hit.ogg
	HitOggFile []byte // 战斗音效
	//go:embed get.ogg
	GetOggFile []byte // 获取到物资音效
	//go:embed open_door.ogg
	OpenDoorOggFile []byte // 开门音效
)

const (
	sampleRate     = 48000 // 采样率
	bytesPerSample = 4     // 每次采样读取的字节数

	introLengthInSecond = 0  // 音频开始的秒数
	loopLengthInSecond  = 58 // 音频循环播放的秒数，bytesPerSample 会影响他
)

// Music 专门播放音乐的类
type Music struct {
	musicMap map[string]*audio.Player
}

func (m *Music) Play(musicType string) {
	music := m.musicMap[musicType]
	if !music.IsPlaying() {
		_ = music.Rewind()
		music.Play()
	} else {
		_ = music.Rewind()
		music.Play()
	}
}

func Init() {
	audioContext = audio.NewContext(sampleRate)

	// 初始化背景音乐
	Bgm1Ogg, _ := vorbis.DecodeWithoutResampling(bytes.NewReader(Bgm1OggFile))
	loop := audio.NewInfiniteLoopWithIntro(Bgm1Ogg, introLengthInSecond*bytesPerSample*sampleRate, loopLengthInSecond*bytesPerSample*sampleRate)
	BackGround, _ = audioContext.NewPlayer(loop)

	// 初始化战斗音乐
	HitOgg, _ := vorbis.DecodeWithoutResampling(bytes.NewReader(HitOggFile))
	HitMusic, _ := audioContext.NewPlayer(HitOgg)

	// 初始化 获取物资 音乐
	GetOgg, _ := vorbis.DecodeWithoutResampling(bytes.NewReader(GetOggFile))
	GetOggMusic, _ := audioContext.NewPlayer(GetOgg)

	// 初始 化开门 音乐
	OpenDoorOgg, _ := vorbis.DecodeWithoutResampling(bytes.NewReader(OpenDoorOggFile))
	OpenDoorOggMusic, _ := audioContext.NewPlayer(OpenDoorOgg)

	musicMap := map[string]*audio.Player{
		"hit":      HitMusic,
		"get":      GetOggMusic,
		"openDoor": OpenDoorOggMusic,
	}

	MusicObj = Music{musicMap: musicMap}
}
