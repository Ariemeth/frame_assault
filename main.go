// main.go
package main

import (
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
	status := tl.NewText(0, 10, "blah", tl.ColorWhite, tl.ColorBlack)
	level.AddEntity(status)
	player := mech.NewPlayerMech("Mech1", 2, tl.NewEntity(1, 1, 1, 1))
	level.AddEntity(player)
	game.Screen().SetLevel(level)

	game.Start()
}
