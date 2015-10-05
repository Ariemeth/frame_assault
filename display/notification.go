package display

import (
	tl "github.com/Ariemeth/termloop"
)

// Notification is a display used to report important information to the player.
type Notification struct {
	Status
	textLine1 *tl.Text
	textLine2 *tl.Text
	textLine3 *tl.Text
	textLine4 *tl.Text
}

//NewNotification returns a new Notification display
func NewNotification(x, y, width, height int, level *tl.BaseLevel) *Notification {
	notification := Notification{
		Status: *NewStatus(x, y, width, height, level),
	}

	notification.textLine1 = tl.NewText(x, y, "No notifications", tl.ColorWhite, tl.ColorBlack)
	notification.textLine2 = tl.NewText(x, y, "", tl.ColorWhite, tl.ColorBlack)
	notification.textLine3 = tl.NewText(x, y, "", tl.ColorWhite, tl.ColorBlack)
	notification.textLine4 = tl.NewText(x, y, "", tl.ColorWhite, tl.ColorBlack)

	return &notification
}

// Draw passes the draw call to entity.
func (display *Notification) Draw(screen *tl.Screen) {
	display.Status.Draw(screen)

	offSetX, offSetY := display.level.Offset()

	display.textLine1.SetPosition(-offSetX+1+display.x, -offSetY+1+display.y)
	display.textLine2.SetPosition(-offSetX+1+display.x, -offSetY+2+display.y)
	display.textLine3.SetPosition(-offSetX+1+display.x, -offSetY+3+display.y)
	display.textLine4.SetPosition(-offSetX+1+display.x, -offSetY+4+display.y)

	display.textLine1.Draw(screen)
	display.textLine2.Draw(screen)
	display.textLine3.Draw(screen)
	display.textLine4.Draw(screen)

}

// AddMessage adds a notification to the notification list.
func (display *Notification) AddMessage(message string) {
	display.textLine1.SetText(display.textLine2.Text())
	display.textLine2.SetText(display.textLine3.Text())
	display.textLine3.SetText(display.textLine4.Text())
	display.textLine4.SetText(message)
}

// Clear clears all entries from the notification display
func (display *Notification) Clear() {
	display.textLine1.SetText("")
	display.textLine2.SetText("")
	display.textLine3.SetText("")
	display.textLine4.SetText("")
}
