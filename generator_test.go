package mandel_test

import (
	"bytes"
	"image/color"
	"io/ioutil"
	"testing"

	"github.com/linuxfreak003/mandel"
	. "github.com/onsi/gomega"
)

func MyColorFunc(iters int) color.RGBA {
	black := color.RGBA{255, 255, 255, 0xff}
	c1 := color.RGBA{0, 100, 255, 0xff}

	if iters == -1 {
		return color.RGBA{0, 0, 0, 0xff}
	}

	if iters < 300 {
		return mandel.Gradient(c1, black, 300, iters)
	}

	if iters < 600 {
		return mandel.Gradient(black, color.RGBA{255, 0, 0, 0xff}, 300, iters-300)
	}

	return mandel.Gradient(color.RGBA{255, 0, 0, 0xff}, color.RGBA{255, 255, 0, 0xff}, 400, iters-600)
}

func TestWrite(t *testing.T) {
	G := NewGomegaWithT(t)
	x, y := mandel.FindInterestingPoint(0, 0)
	m := mandel.NewGenerator(800, 600, x, y).
		WithZoom(900).
		WithAntiAlias(1).
		WithColorizeFunc(MyColorFunc).
		WithLimit(1000)

	m.Generate()
	w := &bytes.Buffer{}
	err := m.Write(w)
	G.Expect(err).To(BeNil())
	err = ioutil.WriteFile("test.png", w.Bytes(), 0644)
	G.Expect(err).To(BeNil())
}
