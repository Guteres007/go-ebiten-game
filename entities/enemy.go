package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	X, Y float64
	img  *ebiten.Image
}

func NewEnemy(x float64, y float64) *Enemy {
	img , _, err := ebitenutil.NewImageFromFile("./alien.png")
	if err != nil {
			log.Fatal(err)
	}
	return &Enemy{X: x, Y: y, img: img,}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.X, e.Y)
    screen.DrawImage(e.img, op)
}