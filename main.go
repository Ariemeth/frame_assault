// main.go
package main

import (
	"github.com/Ariemeth/go_game_jam/display"
	"github.com/Ariemeth/go_game_jam/mech"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		//		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	player := mech.NewPlayerMech("Player", 2, 1, 1, level)
	level.AddEntity(player)

	status := display.NewStatusDisplay(0, 0, 20, 20, player, level)
	level.AddEntity(status)

	game.Screen().SetLevel(level)

	game.Start()
}
