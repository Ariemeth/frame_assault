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
	prevX        int
	prevY        int
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(name string, maxStructure, x, y int) *mech {
	newMech := mech{
		name:         name,
		structure:    maxStructure,
		maxStructure: maxStructure,
		entity:       tl.NewEntity(x, y, 1, 1),
	}

	newMech.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'M'})
	return &newMech
}

// StructureLeft Retrieves the amount of remaining structure a mech has.
func (mech mech) StructureLeft() int {
	return mech.structure
}

// Size returns the height and width of the mech
func (mech mech) Size() (int, int) {
	return mech.entity.Size()
}

// Position returns the x,y position of the mech
func (mech mech) Position() (int, int) {
	return mech.entity.Position()
}

// Collide is used called to see if the mech collided with another physical object
func (mech *mech) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		mech.entity.SetPosition(mech.prevX, mech.prevY)
	}
}

// Draw passes the draw call to entity.
func (mech *mech) Draw(screen *tl.Screen) {
	mech.entity.Draw(screen)
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (mech *mech) Tick(event tl.Event) {

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
