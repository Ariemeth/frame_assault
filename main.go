// main.go
package main

import (
	"math/rand"
	"time"

	"github.com/Ariemeth/go_game_jam/display"
	"github.com/Ariemeth/go_game_jam/mech"
	"github.com/Ariemeth/go_game_jam/mech/weapon"
	tl "github.com/JoelOtter/termloop"
)

// GenerateEnemyMechs creates a slice of mechs to be used as enemies
func GenerateEnemyMechs(number int) []*mech.Mech {
	enemyMechs := make([]*mech.Mech, 0, number)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < number; i++ {
		var m *mech.Mech

		chance := i % 7
		x := -15 + r.Intn(30)
		y := -15 + r.Intn(30)

		switch chance {
		case 0:
			//			m = mech.NewMech("Mech A", i, x, y, tl.ColorRed, rune('A'))
			//			m.AddWeapon(weapon.CreateRifle())
			break
		case 1:
			m = mech.NewMech("Mech B", i, x, y, tl.ColorRed, rune('B'))
			m.AddWeapon(weapon.CreateRifle())
			break
		case 2:
			m = mech.NewMech("Mech C", i, x, y, tl.ColorRed, rune('C'))
			m.AddWeapon(weapon.CreateShotgun())
			break
		case 3:
			m = mech.NewMech("Mech D", i, x, y, tl.ColorRed, rune('D'))
			m.AddWeapon(weapon.CreateShotgun())
			break
		case 4:
			m = mech.NewMech("Mech E", i, x, y, tl.ColorRed, rune('E'))
			m.AddWeapon(weapon.CreateFist())
			break
		case 5:
			m = mech.NewMech("Mech F", i, x, y, tl.ColorRed, rune('F'))
			m.AddWeapon(weapon.CreateFist())
			break
		case 6:
			m = mech.NewMech("Mech G", i, x, y, tl.ColorRed, rune('G'))
			m.AddWeapon(weapon.CreateSword())
			break
		case 7:
			m = mech.NewMech("Mech H", i, x, y, tl.ColorRed, rune('H'))
			m.AddWeapon(weapon.CreateSword())
			break
		}

		if m != nil {
			enemyMechs = append(enemyMechs, m)
		}
	}

	return enemyMechs
}

func main() {
	//Create the game
	game := tl.NewGame()

	//Create the level for the game
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		//		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	//Create the enemy mechs
	enemies := GenerateEnemyMechs(8)
	for _, enemy := range enemies {
		enemy.AttachGame(game)
		level.AddEntity(enemy)
	}

	//Create the player's mech
	player := mech.NewPlayerMech("Player", 10, 1, 1, level)
	weapon1 := weapon.CreateRifle()
	player.AddWeapon(weapon1)
	player.SetEnemyList(enemies)
	player.AttachGame(game)
	level.AddEntity(player)

	//Create the players mech status display
	status := display.NewPlayerStatus(0, 0, 20, 13, player, level)
	level.AddEntity(status)

	//Set the level to be the current game level
	game.Screen().SetLevel(level)

	game.SetDebugOn(true)

	//Start the game engine
	game.Start()
}
