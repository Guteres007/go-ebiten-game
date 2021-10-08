package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	X, Y float64
	Img  *ebiten.Image
	Op *ebiten.DrawImageOptions
	
}

func NewEnemy(x float64, y float64) *Enemy {
	img , _, err := ebitenutil.NewImageFromFile("./alien.png")
	if err != nil {
			log.Fatal(err)
	}
	return &Enemy{X: x, Y: y, Img: img,}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.Op = &ebiten.DrawImageOptions{}
	e.Op.GeoM.Translate(e.X, e.Y)
    screen.DrawImage(e.Img, e.Op)
}