package walker

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Walker a circle that walks around the screen
type Walker struct {
	x, y   float64
	tx, ty float64
}

// New create a new walker
func New() *Walker {
	return &Walker{
		tx: 0,
		ty: 1000,
	}
}

// Display add a circle to the canvas
func (w Walker) Display(imd *imdraw.IMDraw) {
	imd.Push(pixel.V(w.x, w.y))
	imd.Circle(5, 5)
}

// Step increment the positions on the canvas
func (w *Walker) Step() {
	w.x++
	w.y++
}

// func Transform(x, in_min, in_max, out_min, out_max int64) int64 {
//     return (x - in_min) * (out_max - out_min) / (in_max - in_min) + out_min
// }
