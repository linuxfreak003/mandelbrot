package mandel_test

import (
	"bytes"
	"image/color"
	"io/ioutil"
	"testing"

	"github.com/linuxfreak003/mandelbrot"
	. "github.com/onsi/gomega"
)

func TestAverage(t *testing.T) {
	G := NewGomegaWithT(t)

	colors := []color.RGBA{
		color.RGBA{0, 0, 0, 0xff},
		color.RGBA{255, 255, 255, 0xff},
		color.RGBA{255, 255, 255, 0xff},
		color.RGBA{255, 255, 255, 0xff},
	}

	c := mandelbrot.Average(colors...)
	G.Expect(c).To(Equal(color.RGBA{191, 191, 191, 0xff}))
}

func TestWrite(t *testing.T) {
	G := NewGomegaWithT(t)

	MyColorFunc := func(iters int) color.RGBA {
		black := color.RGBA{255, 255, 255, 0xff}
		c1 := color.RGBA{0, 100, 255, 0xff}

		switch {
		case iters == -1:
			return color.RGBA{0, 0, 0, 0xff}
		case iters < 300:
			return mandelbrot.Gradient(c1, black, 300, iters)
		case iters < 600:
			return mandelbrot.Gradient(black, color.RGBA{255, 0, 0, 0xff}, 300, iters-300)
		}

		return mandelbrot.Gradient(color.RGBA{255, 0, 0, 0xff}, color.RGBA{255, 255, 0, 0xff}, 400, iters-600)
	}
	x, y := mandelbrot.FindInterestingPoint(0, 0)
	m := mandelbrot.NewGenerator(1024, 768, x, y).
		WithZoom(900).
		WithAntiAlias(2).
		WithColorizeFunc(MyColorFunc).
		WithLimit(1000)

	m.Generate()
	w := &bytes.Buffer{}
	err := m.Write(w)
	G.Expect(err).To(BeNil())
	err = ioutil.WriteFile("test.png", w.Bytes(), 0644)
	G.Expect(err).To(BeNil())
}
