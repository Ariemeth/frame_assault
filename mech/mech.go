// Package mech project mech.go
package mech

import (
	"fmt"

	"github.com/Ariemeth/go_game_jam/mech/weapon"
	tl "github.com/JoelOtter/termloop"
)

// Mech is a basic mech type
type mech struct {
	structure    int
	maxStructure int
	weapons      []weapon.Weapon
	name         string
	entity       *tl.Entity
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(name string, maxStructure int, entity *tl.Entity) *mech {
	newMech := mech{name: name, structure: maxStructure,
		maxStructure: maxStructure, entity: entity}
	entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	return &newMech
}

// StructureLeft Retrieves the amount of remaining structure a mech has.
func (mech mech) StructureLeft() int {
	return mech.structure
}

// Draw passes the draw call to entity.
func (mech *mech) Draw(screen *tl.Screen) {
	mech.entity.Draw(screen)
}

func (mech *mech) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := mech.entity.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			mech.entity.SetPosition(x+1, y)
			break
		case tl.KeyArrowLeft:
			mech.entity.SetPosition(x-1, y)
			break
		case tl.KeyArrowUp:
			mech.entity.SetPosition(x, y-1)
			break
		case tl.KeyArrowDown:
			mech.entity.SetPosition(x, y+1)
			break
		}
	}
}

// Hit is call when a mech is hit
func (mech *mech) Hit(damage int) {
	mech.structure -= damage
	fmt.Println(mech.name, "takes", damage, "damage")
	if mech.structure <= 0 {
		fmt.Println(mech.name, "destroyed")
	}
}

// AddWeapon adds a Weapon to the mech
func (mech *mech) AddWeapon(weapon weapon.Weapon) {
	mech.weapons = append(mech.weapons, weapon)
}

// Fire tells the Mech to fire at a Target
func (mech *mech) Fire(rangeToTarget int, target weapon.Target) {
	for _, weapon := range mech.weapons {
		weapon.Fire(rangeToTarget, target)
	}
}
