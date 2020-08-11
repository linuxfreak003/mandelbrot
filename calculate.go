package mandel

import (
	"image/color"
	"math/rand"
)

func FindInterestingPoint(x, y float64) (float64, float64) {
	for {
		i := CalcPoint(x, y, 1000)
		if i > 900 {
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

func Gradient(start, end color.RGBA, percentage float64) color.RGBA {
	sR, sG, sB, _ := start.RGBA()
	eR, eG, eB, _ := end.RGBA()
	return color.RGBA{
		R: uint8((float64(eR-sR) * percentage) + float64(sR)),
		G: uint8((float64(eG-sG) * percentage) + float64(sG)),
		B: uint8((float64(eB-sB) * percentage) + float64(sB)),
		A: 0xff,
	}
}
