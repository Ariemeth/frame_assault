package mech

import (
	tl "github.com/JoelOtter/termloop"
)

type playerMech struct {
	mech
}

// NewMech is used to create a new instance of a mech with default structure.
func NewPlayerMech(name string, maxStructure int, entity *tl.Entity) *playerMech {
	newPlayerMech := playerMech{*(NewMech(name, maxStructure, entity))}
	//	newPlayerMech.mech =
	return &newPlayerMech
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (mech *playerMech) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		mech.prevX, mech.prevY = mech.entity.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			mech.entity.SetPosition(mech.prevX+1, mech.prevY)
			break
		case tl.KeyArrowLeft:
			mech.entity.SetPosition(mech.prevX-1, mech.prevY)
			break
		case tl.KeyArrowUp:
			mech.entity.SetPosition(mech.prevX, mech.prevY-1)
			break
		case tl.KeyArrowDown:
			mech.entity.SetPosition(mech.prevX, mech.prevY+1)
			break
		}
	}
}
