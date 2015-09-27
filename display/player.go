package display

import (
	"strconv"

	"github.com/Ariemeth/go_game_jam/mech"
	tl "github.com/JoelOtter/termloop"
)

//Status represents a player status display
type Player struct {
	Status
	textLine1 *tl.Text
	textLine2 *tl.Text
	textLine3 *tl.Text
	textLine4 *tl.Text
	textLine5 *tl.Text
	textLine6 *tl.Text
	textLine7 *tl.Text
	textLine8 *tl.Text
	player    *mech.PlayerMech
}

//NewStatusDisplay creates a new status display for the specified PlayerMech
func NewPlayerStatus(x, y, width, height int, playerMech *mech.PlayerMech, level *tl.BaseLevel) *Player {
	display := Player{
		Status:
		background: tl.NewRectangle(
			x,
			y,
			width,
			height,
			tl.ColorBlack),
		player: playerMech,
		textLine1: tl.NewText(
			x,
			y,
			"display",
			tl.ColorWhite,
			tl.ColorBlack),
		level: level,
		x:     x,
		y:     y,
	}

	display.textLine1 = tl.NewText(x, y, "line1", tl.ColorWhite, tl.ColorBlack)
	display.textLine2 = tl.NewText(x, y, "line2", tl.ColorWhite, tl.ColorBlack)
	display.textLine3 = tl.NewText(x, y, "line3", tl.ColorWhite, tl.ColorBlack)
	display.textLine4 = tl.NewText(x, y, "line4", tl.ColorWhite, tl.ColorBlack)
	display.textLine5 = tl.NewText(x, y, "line5", tl.ColorWhite, tl.ColorBlack)
	display.textLine6 = tl.NewText(x, y, "line6", tl.ColorWhite, tl.ColorBlack)
	display.textLine7 = tl.NewText(x, y, "line7", tl.ColorWhite, tl.ColorBlack)
	display.textLine8 = tl.NewText(x, y, "line7", tl.ColorWhite, tl.ColorBlack)

	return &display
}

// Draw passes the draw call to entity.
func (display *Player) Draw(screen *tl.Screen) {

	offSetX, offSetY := display.level.Offset()

	display.background.SetPosition(-offSetX+display.x, -offSetY+display.y)
	display.textLine1.SetPosition(-offSetX+1+display.x, -offSetY+1+display.y)
	display.textLine2.SetPosition(-offSetX+1+display.x, -offSetY+3+display.y)
	display.textLine3.SetPosition(-offSetX+1+display.x, -offSetY+5+display.y)
	display.textLine4.SetPosition(-offSetX+1+display.x, -offSetY+7+display.y)
	display.textLine5.SetPosition(-offSetX+1+display.x, -offSetY+8+display.y)
	display.textLine6.SetPosition(-offSetX+1+display.x, -offSetY+9+display.y)
	display.textLine7.SetPosition(-offSetX+1+display.x, -offSetY+10+display.y)
	display.textLine8.SetPosition(-offSetX+1+display.x, -offSetY+11+display.y)

	display.background.Draw(screen)
	display.textLine1.Draw(screen)
	display.textLine2.Draw(screen)
	display.textLine3.Draw(screen)
	display.textLine4.Draw(screen)
	display.textLine5.Draw(screen)
	display.textLine6.Draw(screen)
	display.textLine7.Draw(screen)
	display.textLine8.Draw(screen)
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (display *Player) Tick(event tl.Event) {
	display.textLine1.SetText(display.player.Name())
	display.textLine2.SetText("Struture: " + strconv.Itoa(display.player.StructureLeft()))
	x, y := display.player.Position()
	display.textLine3.SetText("Location: (" + strconv.Itoa(x) + "," + strconv.Itoa(y) + ")")

	//assume for now there is only 1 Weapon
	display.textLine4.SetText("Weapons")
	weapons := display.player.Weapons()
	if len(weapons) > 0 {
		display.textLine5.SetText("    Name: " + weapons[0].Name())
		display.textLine5.SetColor(tl.ColorWhite, tl.ColorBlack)
		display.textLine6.SetText("   Range: " + strconv.Itoa(weapons[0].Range()))
		display.textLine7.SetText("  Damage: " + strconv.Itoa(weapons[0].Damage()))
		display.textLine8.SetText("Accuracy: " + strconv.FormatFloat(weapons[0].Accuracy()*100, 'f', 1, 64) + "%")
	} else {
		display.textLine5.SetText("    None")
		display.textLine5.SetColor(tl.ColorRed, tl.ColorBlack)
		display.textLine6.SetText("")
		display.textLine7.SetText("")
		display.textLine8.SetText("")
	}
}
