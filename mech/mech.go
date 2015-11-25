// Package mech project mech.go
package mech

import (
	"strconv"

	"github.com/Ariemeth/frame_assault/mech/weapon"
	"github.com/Ariemeth/frame_assault/util"
	tl "github.com/Ariemeth/termloop"
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
	game         *tl.Game
	notifier     util.Notifier
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

//AttachGame is used to attach the termloop game struct for logging
func (m *Mech) AttachGame(game *tl.Game) {
	m.game = game
}

//AttachNotifier is used to attach a notification display
func (m *Mech) AttachNotifier(notifier util.Notifier) {
	m.notifier = notifier
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
	} else if _, ok := collision.(*Mech); ok {
		m.entity.SetPosition(m.prevX, m.prevY)
	}
}

// Draw passes the draw call to entity.
func (m *Mech) Draw(screen *tl.Screen) {
	if m.StructureLeft() > 0 {
		m.entity.Draw(screen)
	}
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (m *Mech) Tick(event tl.Event) {

}

// Hit is call when a mech is hit
func (m *Mech) Hit(damage int) {
	
	//check if the mech is already destroyed
	if m.structure <= 0 {
		return
	}
	
	m.structure -= damage
	message1 := m.name + " takes " + strconv.Itoa(damage)
	m.game.Log(message1)
	m.notifier.AddMessage(message1)

	if m.structure <= 0 {
		message2 := m.name + " has been destroyed"
		m.game.Log(message2)
		m.notifier.AddMessage(message2)
		m.game.Screen().Level().RemoveEntity(m)
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
