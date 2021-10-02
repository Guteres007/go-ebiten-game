package main

import (
	"fmt"
	entities "hra/entities"
	"hra/uttils"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game implements ebiten.Game interface.
type Game struct{}

var player *entities.Player
var bullet *entities.Bullet
var enemy *entities.Enemy
const screenWidth int = 900
const screenHeight int = 600


var bullets []*entities.Bullet

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	for _ , k := range inpututil.PressedKeys() {
		if k == ebiten.KeyRight {
			player.PlayX = player.PlayX + 1
			if player.PlayX + float64(player.GetPlayerWidthAndHeight()) > float64(screenWidth) {
				player.PlayX = float64(screenWidth) - float64(player.GetPlayerWidthAndHeight())
			}
		}
		if k == ebiten.KeyLeft {
			player.PlayX = player.PlayX + -1
			if player.PlayX <= 0 {
				player.PlayX = 0 
			}
		}
	}

	
	if  inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		bullet = entities.NewBullet(player.PlayX, player.PlayY)
		bullets = append(bullets, bullet)
	}

	if len(bullets) > 0 {
		for _, bullet := range bullets {
			bullet.Y = bullet.Y - 1
		} 
		
		colide := uttils.CalucateDistance(bullet.X, bullet.Y, player.PlayX, player.PlayY)
		fmt.Println(colide)
	}

	


    return nil
}



// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	player.Draw(screen)
	enemy.Draw(screen)
	
	for _, bullet := range bullets {
		bullet.Draw(screen)
	} 

}


// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	player = entities.NewPlayer(float64(screenWidth / 2) - float64(player.GetPlayerWidthAndHeight() /2), float64(screenHeight / 2) - float64(player.GetPlayerWidthAndHeight() /2) )
	enemy = entities.NewEnemy(20, 20)

    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Spaceship")

    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}