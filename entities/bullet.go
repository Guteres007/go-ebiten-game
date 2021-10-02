package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	X, Y float64
	img  *ebiten.Image
}


func NewBullet(x float64, y float64) *Bullet {
	
	img, _, err := ebitenutil.NewImageFromFile("bullet.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Bullet{x,y,img}
}


func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
    screen.DrawImage(b.img, op)
}