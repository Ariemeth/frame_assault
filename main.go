// main.go
package main

import (
	"github.com/Ariemeth/go_game_jam/display"
	"github.com/Ariemeth/go_game_jam/mech"
	"github.com/Ariemeth/go_game_jam/mech/weapon"
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
	weapon1 := weapon.Create(5, 1, "Rifle", .75)
	player.AddWeapon(weapon1)
	level.AddEntity(player)

	status := display.NewStatusDisplay(0, 0, 20, 13, player, level)
	level.AddEntity(status)

	game.Screen().SetLevel(level)

	game.Start()
}
