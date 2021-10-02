package entities

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const playerWidthAndHeight int = 64
type Player struct {
	img *ebiten.Image
	PlayX float64
	PlayY float64
}

func NewPlayer(x float64, y float64) *Player {
	img , _, err := ebitenutil.NewImageFromFile("./spaceship.png")
	if err != nil {
			log.Fatal(err)
	}
	return &Player{
		img: img,
		PlayX: x,
		PlayY: y,
	}
}

func (p *Player) PlayerUpdate(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.PlayX, p.PlayY)
    screen.DrawImage(p.img, op)
}

func (p *Player) GetPlayerWidthAndHeight() int {
	return playerWidthAndHeight
}