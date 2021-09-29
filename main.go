package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	img *ebiten.Image
	playX float64
	playY float64
}
// Game implements ebiten.Game interface.
type Game struct{
	player *Player
}

const screenWidth int = 900
const screenHeight int = 600
const playerWidthAndHeight int = 64

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	for _ , k := range inpututil.PressedKeys() {
		if k == ebiten.KeyRight {
			g.player.playX += 1
			if float64(g.player.playX) + float64(playerWidthAndHeight) > float64(screenWidth){
				g.player.playX = float64(screenWidth) - float64(playerWidthAndHeight)
			}
		}
		if k == ebiten.KeyLeft {
			g.player.playX += -1
			if float64(g.player.playX) <= 0 {
				g.player.playX = 0 
			}
		}
	}

    return nil
}



// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate( g.player.playX, g.player.playY)
    screen.DrawImage(g.player.img, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	img , _, err := ebitenutil.NewImageFromFile("./spaceship.png")
	if err != nil {
			log.Fatal(err)
	}
	
	player := Player{img: img, playX: float64(screenWidth / 2) - float64(playerWidthAndHeight /2), playY: float64(screenHeight / 2) - float64(playerWidthAndHeight /2) }

	game.player = &player

    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Your game's title")



    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}