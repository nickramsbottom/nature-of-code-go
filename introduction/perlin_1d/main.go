package main

import (
	"perlinwalk/walker"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	width, height := 800, 800
	cfg := pixelgl.WindowConfig{
		Title:  "1D Perlin Walker",
		Bounds: pixel.R(0, 0, float64(width), float64(height)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	walker := walker.New()

	for !win.Closed() {
		win.Clear(colornames.White)
		walker.Display(win, 800, 800)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
