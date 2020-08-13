package main

import (
	"image/color"
	"log"
	"os"

	"github.com/linuxfreak003/mandel"
)

func MyColorFunc(iters int) color.RGBA {
	black := color.RGBA{255, 255, 255, 0xff}
	c1 := color.RGBA{0, 100, 255, 0xff}

	switch {
	case iters == -1:
		return color.RGBA{0, 0, 0, 0xff}
	case iters < 300:
		return mandel.Gradient(c1, black, 300, iters)
	case iters < 600:
		return mandel.Gradient(black, color.RGBA{255, 0, 0, 0xff}, 300, iters-300)
	}

	return mandel.Gradient(color.RGBA{255, 0, 0, 0xff}, color.RGBA{255, 255, 0, 0xff}, 400, iters-600)
}

func main() {
	x, y := mandel.FindInterestingPoint(0, 0)
	m := mandel.NewGenerator(2560, 1440, x, y).
		WithZoom(1400).
		WithAntiAlias(3).
		WithColorizeFunc(MyColorFunc).
		WithLimit(1000)

	log.Printf("Generating fractat at %f,%f with parameters\n"+
		"Zoom: %f\nLimit: %d\nAntiAlias: %d", m.X, m.Y, m.Zoom, m.Limit, m.AntiAlias)
	m.Generate()

	f, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = m.Write(f)
	if err != nil {
		panic(err)
	}
}
