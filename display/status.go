package display

import (
	"github.com/Ariemeth/go_game_jam/mech"
	tl "github.com/JoelOtter/termloop"
)

//Status represents a player status display
type Status struct {
	background *tl.Rectangle
	textLine1  *tl.Text
	textLine2  *tl.Text
	textLine3  *tl.Text
	textLine4  *tl.Text
	level      *tl.BaseLevel
	player     *mech.PlayerMech
	x          int
	y          int
}

//NewStatusDisplay creates a new status display for the specified PlayerMech
func NewStatusDisplay(x, y, width, height int, playerMech *mech.PlayerMech, level *tl.BaseLevel) *Status {
	display := Status{
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

	return &display
}

// Draw passes the draw call to entity.
func (display *Status) Draw(screen *tl.Screen) {

	offSetX, offSetY := display.level.Offset()

	display.background.SetPosition(-offSetX+display.x, -offSetY+display.y)
	display.textLine1.SetPosition(-offSetX+1+display.x, -offSetY+1+display.y)
	display.textLine2.SetPosition(-offSetX+1+display.x, -offSetY+2+display.y)
	display.textLine3.SetPosition(-offSetX+1+display.x, -offSetY+3+display.y)
	display.textLine4.SetPosition(-offSetX+1+display.x, -offSetY+4+display.y)

	display.background.Draw(screen)
	display.textLine1.Draw(screen)
	display.textLine2.Draw(screen)
	display.textLine3.Draw(screen)
	display.textLine4.Draw(screen)
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (display *Status) Tick(event tl.Event) {

}
