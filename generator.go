package mandelbrot

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
)

// ColorFucn is what is used to generate the colorscheme
// It takes the number of iterations from the mandelbrot
// calculation and returns a color DefaultColorize is an
// extremely basic example
type ColorFunc func(int) color.RGBA

// DefaultColorze is a very simple Greyscale ColorFunc
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

// Generator is the used to generate the fractal
type Generator struct {
	// Width and Height specify the resolution to use
	Width  int
	Height int
	// X and Y specify what point on the fractal to center on
	X float64
	Y float64
	// Zoom specifies how much to zoom in
	Zoom float64
	// Limit specifies when the mandelbrot calculation
	// should bail out and return -1 instead of
	// the number of iterations
	Limit int
	// AntiAlias specifies what level of antialiasing to use
	// An AntiAlias of 2 will average 4 points for each pixel
	// 3 will average 9 points. The increase is exponential
	AntiAlias int
	// Colorize is the ColorFunc used to generate the colorscheme
	Colorize ColorFunc
	//img is the underlying image
	img *image.RGBA
}

// NewGenerator creates a new *Generator
// it should be used to ensure all fields
// are filled.
func NewGenerator(width, height int, x, y float64) *Generator {
	return &Generator{
		Width:     width,
		Height:    height,
		X:         x,
		Y:         y,
		Zoom:      1,
		Limit:     1000,
		AntiAlias: 1,
		Colorize:  DefaultColorize,
		img:       image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// Sets the AntiAliasing level
func (g *Generator) WithAntiAlias(aa int) *Generator {
	if aa < 1 {
		aa = 1
	}
	g.AntiAlias = aa
	return g
}

// Sets the zoom level
func (g *Generator) WithZoom(z float64) *Generator {
	if z < 1 {
		z = 1
	}
	g.Zoom = z
	return g
}

// Sets the bailout limit for fractal
func (g *Generator) WithLimit(l int) *Generator {
	g.Limit = l
	return g
}

// Sets the Colorize function used to generate colorscheme
func (g *Generator) WithColorizeFunc(f ColorFunc) *Generator {
	if f == nil {
		return g
	}
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

// Generate does the mandelbrot calculation
// and stores the fractal into an image
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
		g.img.Set(p.X, p.Y, p.Color)
	}
	return
}

// AntiAliasedColor breaks a pixel down into parts
// and gets the color for each point, then averages
// them out for the pixel color
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

// GetColor gets the mandelbrot calculation iterations
// and Uses the defined Colorize function turn into a color
func (g *Generator) GetColor(x, y float64) color.RGBA {
	iter := Calculate(x, y, g.Limit)
	return g.Colorize(iter)
}

// WritePNG writes the underlying image to a an io.Writer
// as a PNG
func (g *Generator) WritePNG(w io.Writer) error {
	return png.Encode(w, g.img)
}

// WriteJPG writes the underlying image to a an io.Writer
// as a JPG
func (g *Generator) WriteJPG(w io.Writer) error {
	return jpeg.Encode(w, g.img, nil)
}
