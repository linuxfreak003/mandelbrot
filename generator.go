package mandel

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type ColorFunc func(int) color.RGBA

type Generator struct {
	Width     int
	Height    int
	X         float64
	Y         float64
	Zoom      float64
	Limit     int
	AntiAlias int
	Image     *image.RGBA
	Colorize  ColorFunc
}

// Simple Greyscale colorscheme
func DefaultColorize(iter int) color.RGBA {
	if iter == -1 {
		return color.RGBA{0, 0, 0, 0xff}
	}

	return color.RGBA{
		uint8(iter % 255),
		uint8(iter % 255),
		uint8(iter % 255),
		0xff,
	}
}

func NewGenerator(width, height int, x, y float64) *Generator {
	return &Generator{
		Width:     width,
		Height:    height,
		X:         x,
		Y:         y,
		Zoom:      1,
		Limit:     1000,
		AntiAlias: 1,
		Image:     image.NewRGBA(image.Rect(0, 0, width, height)),
		Colorize:  DefaultColorize,
	}
}

func (g *Generator) WithAntiAlias(aa int) *Generator {
	if aa < 1 {
		aa = 1
	}
	g.AntiAlias = aa
	return g
}
func (g *Generator) WithZoom(z float64) *Generator {
	if z < 1 {
		z = 1
	}
	g.Zoom = z
	return g
}
func (g *Generator) WithLimit(l int) *Generator {
	g.Limit = l
	return g
}
func (g *Generator) WithColorizeFunc(f ColorFunc) *Generator {
	g.Colorize = f
	return g
}

func (g *Generator) SetWidth(width int)       { g.Width = width }
func (g *Generator) SetHeight(height int)     { g.Height = height }
func (g *Generator) SetX(x float64)           { g.X = x }
func (g *Generator) SetY(y float64)           { g.Y = y }
func (g *Generator) SetZoom(zoom float64)     { g.Zoom = zoom }
func (g *Generator) SetLimit(limit int)       { g.Limit = limit }
func (g *Generator) SetColorize(cf ColorFunc) { g.Colorize = cf }

func (g *Generator) Generate() {
	type pixel struct {
		X, Y  int
		Color color.RGBA
	}

	ch := make(chan pixel, 0)

	inc := 4.0 / (float64(g.Height) * g.Zoom)
	x0 := g.X - inc*float64(g.Width/2)
	y0 := g.Y - inc*float64(g.Height/2)
	for x, a := 0, x0; x < g.Width; x, a = x+1, a+inc {
		for y, b := 0, y0; y < g.Height; y, b = y+1, b+inc {
			go func(a, b, inc float64, x, y int) {
				col := g.AntiAliasedColor(a, b, inc)
				ch <- pixel{x, y, col}
			}(a, b, inc, x, y)
		}
	}

	for c := 0; c < g.Width*g.Height; c++ {
		p := <-ch
		g.Image.Set(p.X, p.Y, p.Color)
	}
	return
}

func (g *Generator) AntiAliasedColor(x, y, inc float64) color.RGBA {
	colors := []color.RGBA{}
	smallInc := inc / float64(g.AntiAlias)
	for i := x + smallInc/2; i < x+inc; i += smallInc {
		for j := y + smallInc/2; j < y+inc; j += smallInc {
			colors = append(colors, g.GetColor(i, j))
		}
	}
	return Average(colors...)
}

func (g *Generator) GetColor(x, y float64) color.RGBA {
	iter := CalcPoint(x, y, g.Limit)
	return g.Colorize(iter)
}

func (g *Generator) Write(w io.Writer) error {
	return png.Encode(w, g.Image)
}
