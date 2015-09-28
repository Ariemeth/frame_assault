package display

import (
	tl "github.com/JoelOtter/termloop"
)

// Notification is a display used to report important information to the player.
type Notification struct {
	Status
}

//NewNotification returns a new Notification display
func NewNotification(x, y, width, height int, level *tl.BaseLevel) *Notification {
	notification := Notification{
		Status: *NewStatus(x, y, width, height, level),
	}

	return &notification
}
