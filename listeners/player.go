package listeners

import (
	"hra/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)



func PlayerMoving(player *entities.Player) {
	for _ , key := range inpututil.PressedKeys() {
		if key == ebiten.KeyRight {
			player.PlayX = player.PlayX + 1
			if player.PlayX + float64(player.GetPlayerWidthAndHeight()) > float64(900) {
				player.PlayX = float64(900) - float64(player.GetPlayerWidthAndHeight())
			}
		}
		if key == ebiten.KeyLeft {
			player.PlayX = player.PlayX + -1
			if player.PlayX <= 0 {
				player.PlayX = 0 
			}
		}
	}
}
