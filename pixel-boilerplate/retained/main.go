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

	circle := imdraw.New(nil)

	circle.Color = colornames.Black

	x := 10.0
	y := 10.0

	circle.Push(pixel.V(x, y))
	circle.Circle(5, 5)

	for !win.Closed() {
		win.Clear(colornames.White)

		circle.Draw(win)

		win.Update()
		x++
		y++
	}
}

func main() {
	pixelgl.Run(run)
}
