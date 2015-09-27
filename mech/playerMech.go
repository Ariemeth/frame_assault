package mech

import (
	tl "github.com/JoelOtter/termloop"
)

type playerMech struct {
	mech
	level *tl.BaseLevel
}

// NewMech is used to create a new instance of a mech with default structure.
func NewPlayerMech(name string, maxStructure int, entity *tl.Entity, level *tl.BaseLevel) *playerMech {
	newPlayerMech := playerMech{*NewMech(name, maxStructure, entity), level}
	//	newPlayerMech.mech =
	return &newPlayerMech
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (pMech *playerMech) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		pMech.prevX, pMech.prevY = pMech.entity.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			pMech.entity.SetPosition(pMech.prevX+1, pMech.prevY)
			break
		case tl.KeyArrowLeft:
			pMech.entity.SetPosition(pMech.prevX-1, pMech.prevY)
			break
		case tl.KeyArrowUp:
			pMech.entity.SetPosition(pMech.prevX, pMech.prevY-1)
			break
		case tl.KeyArrowDown:
			pMech.entity.SetPosition(pMech.prevX, pMech.prevY+1)
			break
		}
	}
}

// Draw passes the draw call to entity.
func (pMech *playerMech) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := pMech.entity.Position()
	pMech.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	pMech.entity.Draw(screen)
}
