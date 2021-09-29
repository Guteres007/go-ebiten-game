package entities

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const playerWidthAndHeight int = 64

type Player struct {
	Img *ebiten.Image
	PlayX float64
	PlayY float64
	
}

func (p *Player) GetPosX() float64 {
	return p.PlayX
}

func (p *Player) GetPosY() float64 {
	return p.PlayY
}

func (p *Player) SetPosX(newPosition float64)  {
	p.PlayX = newPosition
}
func (p *Player) SetPosY(newPosition float64)  {
	p.PlayY = newPosition
}

func (p *Player) GetPlayerWidthAndHeight() int {
	return playerWidthAndHeight
}