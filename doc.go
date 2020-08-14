// Package mandel implements a simple library for generating mandelbrot fractals.
//
// Functions are provided for writing a png or jpeg
// Basic usage of package is
//
// // Creates a new Generator
// g := NewGenerator(1024, 768, 0,0)
// // Generates the image
// g.Generate()
// // Writes the image as a png to "test.png"
// f, _ := os.Create("test.png")
// g.WritePNG(f)
package mandel

const VERSION = "v1.0"
