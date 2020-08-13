package mandel

import (
	"image/color"
	"math/rand"
	"time"
)

func FindInterestingPoint(x, y float64) (float64, float64) {
	rand.Seed(time.Now().UnixNano())
	for {
		i := CalcPoint(x, y, 1000)
		if i > 990 {
			return x, y
		}
		x = (rand.Float64() * 4) - 2
		y = (rand.Float64() * 4) - 2
	}
}

// Mandel Calculation
func CalcPoint(x, y float64, limit int) int {
	i := 0
	for a, b := x, y; a*a+b*b < 4.0; a, b = a*a-b*b+x, 2*a*b+y {
		if i > limit {
			return -1
		}
		i++
	}

	return i
}

func Average(colors ...color.RGBA) color.RGBA {
	if len(colors) == 0 {
		return color.RGBA{}
	}

	var r, g, b int

	for _, c := range colors {
		r += int(c.R)
		g += int(c.G)
		b += int(c.B)
	}

	l := len(colors)

	return color.RGBA{
		uint8(r / l),
		uint8(g / l),
		uint8(b / l),
		0xff,
	}
}

func Gradient(start, end color.RGBA, max, steps int) color.RGBA {
	fi := float64(steps)
	return color.RGBA{
		uint8(fi*(float64(end.R-start.R)/float64(max-1))) + start.R,
		uint8(fi*(float64(end.G-start.G)/float64(max-1))) + start.G,
		uint8(fi*(float64(end.B-start.B)/float64(max-1))) + start.B,
		0xff,
	}
}
