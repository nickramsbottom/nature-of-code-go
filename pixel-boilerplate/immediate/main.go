package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Color = colornames.Black

	x := 0.0
	y := 0.0

	for !win.Closed() {
		imd.Clear()
		win.Clear(colornames.White)

		imd.Push(pixel.V(x, y))
		imd.Circle(5, 5)
		imd.Draw(win)

		win.Update()
		x++
		y++
	}
}

func main() {
	pixelgl.Run(run)
}
