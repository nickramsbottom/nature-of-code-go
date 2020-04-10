package walker

import (
	"github.com/aquilax/go-perlin"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Walker a circle that walks around the screen
type Walker struct {
	x, y   float64
	tx, ty float64
}

var imd *imdraw.IMDraw = imdraw.New(nil)
var p *perlin.Perlin = perlin.NewPerlin(2, 4, 3, 0)

// New create a new walker
func New() *Walker {
	imd.Color = colornames.Black

	return &Walker{
		tx: 0,
		ty: 10000,
	}
}

// Display add a circle to the canvas
func (w *Walker) Display(win *pixelgl.Window, width, height int) {
	w.x = transform(p.Noise1D(w.tx), -0.5, 0.5, 0.0, float64(width))
	w.y = transform(p.Noise1D(w.ty), -0.5, 0.5, 0.0, float64(height))

	// imd.Clear()
	imd.Push(pixel.V(w.x, w.y))
	imd.Circle(5, 1)
	imd.Draw(win)

	w.tx += 0.001
	w.ty += 0.001
}

// transform maps one range to another
func transform(x, inMin, inMax, outMin, outMax float64) float64 {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
