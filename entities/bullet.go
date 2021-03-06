package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type Bullet struct {
	X, Y float64
	img  *ebiten.Image
	Op *ebiten.DrawImageOptions
	speed float64 
}

var Op ebiten.DrawImageOptions 

func NewBullet(x float64, y float64) *Bullet {
	
	img, _, err := ebitenutil.NewImageFromFile("bullet.png")
	if err != nil {
		log.Fatal(err)
	}
	
	return &Bullet{x,y,img, &ebiten.DrawImageOptions{}, 3}
}


func (b *Bullet) Draw(screen *ebiten.Image) {
	b.Op = &ebiten.DrawImageOptions{}
	//b.Op = &ebiten.DrawImageOptions{}
	b.Op.GeoM.Translate(b.X + 16, b.Y)
    screen.DrawImage(b.img, b.Op)
}

func (e *Bullet) Update() {
	e.Y = e.Y - 1 * e.speed
}

func (b *Bullet) RemoveBullet(s []*Bullet, index int) []*Bullet {
	return append(s[:index], s[index+1:]...)
}