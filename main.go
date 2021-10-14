package main

import (
	"fmt"
	entities "hra/entities"
	listeners "hra/listeners"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Game implements ebiten.Game interface.
type Game struct{
	counter        int
}

var background *ebiten.Image
var player *entities.Player
var bullet *entities.Bullet
var enemy *entities.Enemy
const screenWidth int = 900
const screenHeight int = 600
var op *ebiten.DrawImageOptions
var bullets []*entities.Bullet

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).

const sampleText = `The quick brown fox jumps over the lazy dog.`
var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	background , _, _ = ebitenutil.NewImageFromFile("./background.png")
	op = &ebiten.DrawImageOptions{}
	
	op.GeoM.Scale(0.6,0.6)
	op.GeoM.Translate(-300,0)
}

func (g *Game) Update() error {

	listeners.PlayerMoving(player)
	enemy.Moving()
	
	if  inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		bullet = entities.NewBullet(player.PlayX, player.PlayY)
		bullets = append(bullets, bullet)
	}

	if len(bullets) > 0 {
		for i, bullet := range bullets {

			//tady by měl být nějaký update bullet function
			bullet.Update()
				
			//collision detection (hit) pro bullet 
			collisonDetect(i, enemy, bullet)
			//budeli mimo screen
			
			
		} 
		
	
	}

    return nil
}



// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {


	screen.DrawImage(background, op)

	player.Draw(screen)
	enemy.Draw(screen)
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	text.Draw(screen, msg, mplusNormalFont, 20, 40, color.White)

	// Draw the sample text
	text.Draw(screen, sampleText, mplusNormalFont, 90, 80, color.White)
	
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
	player = entities.NewPlayer(float64(screenWidth / 2) - float64(player.GetPlayerWidthAndHeight() /2), float64(screenHeight) - float64(player.GetPlayerWidthAndHeight() + 20) )
	enemy = entities.NewEnemy(400.0, 20.0)

    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Spaceship")

    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}



func collisonDetect(i int, enemy *entities.Enemy, bullet *entities.Bullet) {
	if  ( math.Abs( enemy.Op.GeoM.Element(1,2) - bullet.Op.GeoM.Element(1,2) + 64) ) <= 0  &&
			 ( (bullet.Op.GeoM.Element(0,2) - enemy.Op.GeoM.Element(0,2)) > 0 && (bullet.Op.GeoM.Element(0,2) - enemy.Op.GeoM.Element(0,2)) < 64) {
				bullets = bullet.RemoveBullet(bullets, i)
			}

	if int(bullet.Y) <= 0  {
		bullets = bullet.RemoveBullet(bullets, i)
	}
}