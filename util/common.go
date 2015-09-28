package util

import (
	"math"
)

//CalculateDistance returns the distance between points x1,y1 and x2,y2
func CalculateDistance(x1, y1, x2, y2 int) float64 {
	x := math.Pow((float64)(x2-x1), 2)
	y := math.Pow((float64)(y2-y1), 2)

	d := math.Sqrt(x + y)
	return d
}

// Notifier is an interface that can be implemented to recieve messages
type Notifier interface {
	AddMessage(string)
}
