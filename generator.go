package mandel

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type Generator struct {
	Width      int
	Height     int
	X          float64
	Y          float64
	Zoom       float64
	Limit      int
	Image      *image.RGBA
	ColorStart color.RGBA
	ColorEnd   color.RGBA
}

func NewGenerator(width, height int, x, y, zoom float64) *Generator {
	return &Generator{
		Width:      width,
		Height:     height,
		X:          x,
		Y:          y,
		Zoom:       zoom,
		Limit:      1000,
		Image:      image.NewRGBA(image.Rect(0, 0, width, height)),
		ColorStart: color.RGBA{0, 0, 0, 0xff},
		ColorEnd:   color.RGBA{255, 255, 255, 0xff},
	}
}

func (g *Generator) SetWidth(width int)   { g.Width = width }
func (g *Generator) SetHeight(height int) { g.Height = height }
func (g *Generator) SetX(x float64)       { g.X = x }
func (g *Generator) SetY(y float64)       { g.Y = y }
func (g *Generator) SetZoom(zoom float64) { g.Zoom = zoom }
func (g *Generator) SetLimit(limit int)   { g.Limit = limit }

func (g *Generator) Generate() error {
	wInc := 4.0 / float64(g.Width)
	hInc := 4.0 / float64(g.Height)
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			a := (float64(x) * wInc) - 2.0
			b := (float64(y) * hInc) - 2.0
			col := g.GetColor(a, b)
			g.Image.Set(x, y, col)
		}
	}
	return nil
}

func (g *Generator) GetColor(x, y float64) color.RGBA {
	iter := CalcPoint(x, y, g.Limit)
	if iter == -1 {
		return g.ColorStart
	}
	percentage := float64(iter) / float64(g.Limit)
	return Gradient(g.ColorStart, g.ColorEnd, percentage)
}

func (g *Generator) Write(w io.Writer) error {
	return png.Encode(w, g.Image)
}
