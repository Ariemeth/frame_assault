package display

import tl "github.com/JoelOtter/termloop"

//Status represents a player status display
type Status struct {
	background *tl.Rectangle
	level      *tl.BaseLevel
	x          int
	y          int
}

//NewStatus creates a new status display for the specified PlayerMech
func NewStatus(x, y, width, height int, level *tl.BaseLevel) *Status {
	display := Status{
		background: tl.NewRectangle(
			x,
			y,
			width,
			height,
			tl.ColorBlack),
		level: level,
		x:     x,
		y:     y,
	}

	return &display
}

// Draw passes the draw call to entity.
func (display *Status) Draw(screen *tl.Screen) {

	offSetX, offSetY := display.level.Offset()

	display.background.SetPosition(-offSetX+display.x, -offSetY+display.y)

	display.background.Draw(screen)
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (display *Status) Tick(event tl.Event) {

}
