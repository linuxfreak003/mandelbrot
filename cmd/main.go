package main

import (
	"image/color"
	"log"
	"os"
	"time"

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

func SimpleGrayscale(i int) color.RGBA {
	if i == -1 {
		return color.RGBA{0, 0, 0, 0xff}
	}
	return color.RGBA{
		uint8(i % 255),
		uint8(i % 255),
		uint8(i % 255),
		0xff,
	}
}

func GradientGrayscale(i int) color.RGBA {
	white := color.RGBA{255, 255, 255, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}
	if i == -1 {
		return black
	}

	return mandel.Gradient(black, white, 1000, i)
}

func main() {
	x, y := mandel.FindInterestingPoint(0, 0)
	m := mandel.NewGenerator(2560, 1400, x, y).
		WithZoom(1400).
		WithAntiAlias(3).
		WithColorizeFunc(GradientGrayscale).
		WithLimit(1000)

	log.Printf("Generating fractal...")
	t := time.Now()
	m.Generate()
	log.Printf("Parameters\nX: %f Y: %f\n"+
		"Zoom: %f\nLimit: %d\nAntiAlias: %d\n"+
		"Took: %v", m.X, m.Y, m.Zoom, m.Limit, m.AntiAlias, time.Now().Sub(t))

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
