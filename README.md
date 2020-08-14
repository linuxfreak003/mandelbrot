# mandel

Mandelbrot Fractal Generator library in Go with accompanying CLI

## Library

Basic example usage of the library:

```go
package main

import (
  "os"
  "image/color"

  "github.com/linuxfreak003/mandel"
)

// Simple example color scheme function
func ColorMe(i int) color.RGBA {
  black := color.RGBA{0,0,0,0xff}
  white := color.RGBA{255,255,255,0xff}
  if i == -1 {
    return black
  }
  return mandel.Gradient(black, white, 1000, i)
}

func main() {
  // Finds an interesting point to use
  x, y := mandel.FindInterestingPoint(0,0)

  // Get a new generator, use ColorMe for
  // the colorscheme
  g := mandel.NewGenerator(1024, 768, x, y).
    WithColorizeFunc(ColorMe)

  // Generate the image
  g.Generate()

  // Save the image to file
  f, err := os.Create("test.png")
  if err != nil {
    panic(err)
  }
  // As a PNG
  err = g.WritePNG(f)
  if err != nil {
    panic(err)
  }
}
```

Refer to the docs for more options like zoom, antialiasing, etc.

## CLI

### Installation

```bash
go get -u github.com/linuxfreak003/mandel/mandel-go
```

### Usage

```bash
Usage of mandel-go:
  -aa int
        AntiAlias level (default 1)
  -colorfunc Color Function
        Color Function (default "SimpleGrayscale")
  -h    Show usage help
  -height int
        Resolution height (default 768)
  -j    Julia Set
  -limit int
        Fractal Calculation limit (default 1000)
  -o string
        Output Filename (default "test.png")
  -r    Use random interesting point (will override x/y)
  -width int
        Resolution width (default 1024)
  -x float
        X
  -y float
        Y
  -zoom float
        Zoom (default 1)
```
