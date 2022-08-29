package role

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"mota/maps"
)

type Player struct {
	X         float64
	Y         float64
	State     int
	Direction int
	MouseX    int
	MouseY    int
	SkillName string
	image     *embed.FS
}

func (p *Player) LoadImages(screen *ebiten.Image) {
	imagePath := fmt.Sprintf("./image/main/resource_%v.png", 0)
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X*maps.ImageLarge, p.Y*maps.ImageLarge)
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
		p.Y -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		p.Y += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.X -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		p.X += 1
	}
}

func (p *Player) CheckCanMove() {

}
