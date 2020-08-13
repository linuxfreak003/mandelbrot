package mandel

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type ColorFunc func(int) color.RGBA

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
	Colorize   ColorFunc
}

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

func NewGenerator(width, height int, x, y, zoom float64) *Generator {
	if zoom < 1 {
		zoom = 1
	}
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
		Colorize:   DefaultColorize,
	}
}

func (g *Generator) SetWidth(width int)       { g.Width = width }
func (g *Generator) SetHeight(height int)     { g.Height = height }
func (g *Generator) SetX(x float64)           { g.X = x }
func (g *Generator) SetY(y float64)           { g.Y = y }
func (g *Generator) SetZoom(zoom float64)     { g.Zoom = zoom }
func (g *Generator) SetLimit(limit int)       { g.Limit = limit }
func (g *Generator) SetColorize(cf ColorFunc) { g.Colorize = cf }

func (g *Generator) Generate() error {
	inc := 4.0 / (float64(g.Height) * g.Zoom)
	x0 := g.X - inc*float64(g.Width/2)
	y0 := g.Y - inc*float64(g.Height/2)
	for x, a := 0, x0; x < g.Width; x, a = x+1, a+inc {
		for y, b := 0, y0; y < g.Height; y, b = y+1, b+inc {
			col := g.GetColor(a, b)
			g.Image.Set(x, y, col)
		}
	}
	return nil
}

func (g *Generator) AntiAlias(level int, x, y, inc float64) color.RGBA {
	// Get points in middle of pixel
	// Return Average of those points
	return Average()
}

func (g *Generator) GetColor(x, y float64) color.RGBA {
	iter := CalcPoint(x, y, g.Limit)
	return g.Colorize(iter)
}

func (g *Generator) Write(w io.Writer) error {
	return png.Encode(w, g.Image)
}
