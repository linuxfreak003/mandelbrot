package main

import (
	"flag"
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
		return mandel.Gradient(black, c1, 300, iters-300)
	}

	return mandel.Gradient(c1, color.RGBA{255, 255, 0, 0xff}, 400, iters-600)
}

func CustomColorFunc(iters int) color.RGBA {
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{255, 255, 255, 0xff}
	red := color.RGBA{255, 0, 0, 0xff}

	switch {
	case iters == -1:
		return black
	case iters < 300:
		return mandel.Gradient(red, white, 300, iters)
	case iters < 600:
		return mandel.Gradient(white, red, 300, iters-300)
	}

	return mandel.Gradient(red, color.RGBA{255, 255, 0, 0xff}, 400, iters-600)
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

var FuncMap = map[string]func(int) color.RGBA{
	"SimpleGrayscale":   SimpleGrayscale,
	"GradientGrayscale": GradientGrayscale,
	"CustomColorFunc":   CustomColorFunc,
	"MyColorFunc":       MyColorFunc,
}

func main() {
	var colorfunc, filename string
	var zoom, x, y float64
	var aa, width, height, limit int
	var random, julia bool

	flag.StringVar(&colorfunc, "colorfunc", "SimpleGrayscale", "`Color Function`")
	flag.StringVar(&filename, "o", "test.png", "Output Filename")
	flag.Float64Var(&zoom, "zoom", 1.0, "Zoom")
	flag.Float64Var(&x, "x", 0.0, "X")
	flag.Float64Var(&y, "y", 0.0, "Y")
	flag.IntVar(&width, "width", 1024, "Resolution width")
	flag.IntVar(&height, "height", 768, "Resolution height")
	flag.IntVar(&aa, "aa", 1, "AntiAlias level")
	flag.IntVar(&limit, "limit", 1000, "Fractal Calculation limit")
	flag.BoolVar(&random, "r", false, "Use random interesting point (will override x/y)")
	flag.BoolVar(&julia, "j", false, "Julia Set")

	flag.Parse()

	if random {
		x, y = mandel.FindInterestingPoint(0, 0)
	}

	F := FuncMap[colorfunc]

	m := mandel.NewGenerator(width, height, x, y).
		WithZoom(zoom).
		WithAntiAlias(aa).
		WithColorizeFunc(F).
		WithLimit(limit)

	log.Printf("Generating fractal...")
	t := time.Now()
	m.Generate()
	log.Printf("Parameters\nX: %f Y: %f\n"+
		"Zoom: %f\nLimit: %d\nAntiAlias: %d\n"+
		"Took: %v", m.X, m.Y, m.Zoom, m.Limit, m.AntiAlias, time.Now().Sub(t))

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = m.Write(f)
	if err != nil {
		panic(err)
	}
}
