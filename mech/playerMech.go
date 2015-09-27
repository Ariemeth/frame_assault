package mech

import (
	"strconv"
	"strings"

	"github.com/Ariemeth/go_game_jam/util"
	tl "github.com/JoelOtter/termloop"
)

//PlayerMech represents a player controlled mech
type PlayerMech struct {
	Mech
	level   *tl.BaseLevel
	enemies []*Mech
}

// NewPlayerMech is used to create a new instance of a mech with default structure.
func NewPlayerMech(name string, maxStructure, x, y int, level *tl.BaseLevel) *PlayerMech {
	newPlayerMech := PlayerMech{
		Mech: *NewMech(
			name,
			maxStructure,
			x,
			y,
			tl.ColorRed,
			'M'),
		level: level,
	}

	return &newPlayerMech
}

//SetEnemyList sets the list of enemies the player can interact
func (pMech *PlayerMech) SetEnemyList(enemies []*Mech) {
	pMech.enemies = enemies
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (pMech *PlayerMech) Tick(event tl.Event) {
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
		case tl.KeyCtrlA:
			pMech.attackTarget("A")
			break
		case tl.KeyCtrlB:
			pMech.attackTarget("B")
			break
		case tl.KeyCtrlC:
			pMech.attackTarget("C")
			break
		case tl.KeyCtrlD:
			pMech.attackTarget("D")
			break
		case tl.KeyCtrlE:
			pMech.attackTarget("E")
			break
		case tl.KeyCtrlF:
			pMech.attackTarget("F")
			break
		case tl.KeyCtrlG:
			pMech.attackTarget("G")
			break
		case tl.KeyCtrlH:
			pMech.attackTarget("H")
			break
		}
	}
}

// Draw passes the draw call to entity.
func (pMech *PlayerMech) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := pMech.entity.Position()
	pMech.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	pMech.entity.Draw(screen)
}

func (pMech *PlayerMech) getTargetEnemy(name string) *Mech {
	for i, mech := range pMech.enemies {
		if strings.HasSuffix(mech.Name(), name) {
			pMech.game.Log("enemy found: %s", mech.Name())
			return pMech.enemies[i]
		}
	}
	return nil
}

func (pMech *PlayerMech) attackTarget(name string) {
	target := pMech.getTargetEnemy(name)
	if target == nil {
		return
	}

	targetX, targetY := target.Position()
	distance := util.CalculateDistance(pMech.prevX, pMech.prevY, targetX, targetY)
	pMech.Fire((int)(distance), target)
	pMech.game.Log("distance " + strconv.Itoa((int)(distance)))
	pMech.game.Log("player (%d,%d), target (%d,%d)", pMech.prevX, pMech.prevY, targetX, targetY)
}
