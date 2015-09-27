// Package mech project mech.go
package mech

import (
	"fmt"

	"github.com/Ariemeth/go_game_jam/mech/weapon"
	tl "github.com/JoelOtter/termloop"
)

// Mech is a basic mech type
type Mech struct {
	structure    int
	maxStructure int
	weapons      []weapon.Weapon
	name         string
	entity       *tl.Entity
	prevX        int
	prevY        int
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(name string, maxStructure, x, y int, color tl.Attr, symbol rune) *Mech {
	newMech := Mech{
		name:         name,
		structure:    maxStructure,
		maxStructure: maxStructure,
		entity:       tl.NewEntity(x, y, 1, 1),
	}

	newMech.entity.SetCell(0, 0, &tl.Cell{Fg: color, Ch: symbol})
	return &newMech
}

//Name returns the name of the mech
func (m Mech) Name() string {
	return m.name
}

//Weapons returns the mechs weapons
func (m Mech) Weapons() []weapon.Weapon {
	return m.weapons
}

// StructureLeft Retrieves the amount of remaining structure a mech has.
func (m Mech) StructureLeft() int {
	return m.structure
}

// Size returns the height and width of the mech
func (m Mech) Size() (int, int) {
	return m.entity.Size()
}

// Position returns the x,y position of the mech
func (m Mech) Position() (int, int) {
	return m.entity.Position()
}

// Collide is used called to see if the mech collided with another physical object
func (m *Mech) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		m.entity.SetPosition(m.prevX, m.prevY)
	}
}

// Draw passes the draw call to entity.
func (m *Mech) Draw(screen *tl.Screen) {
	m.entity.Draw(screen)
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (m *Mech) Tick(event tl.Event) {

}

// Hit is call when a mech is hit
func (m *Mech) Hit(damage int) {
	m.structure -= damage
	fmt.Println(m.name, "takes", damage, "damage")
	if m.structure <= 0 {
		fmt.Println(m.name, "destroyed")
	}
}

// AddWeapon adds a Weapon to the mech
func (m *Mech) AddWeapon(weapon weapon.Weapon) {
	m.weapons = append(m.weapons, weapon)
}

// Fire tells the Mech to fire at a Target
func (m *Mech) Fire(rangeToTarget int, target weapon.Target) {
	for _, weapon := range m.weapons {
		weapon.Fire(rangeToTarget, target)
	}
}
