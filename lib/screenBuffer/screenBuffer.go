package screenbuffer

import (
	"fmt"
	"log"
	"os"
)

// Pixel defines a pixel colour in rgb format, where each element is in the
// range 0-255
type Pixel struct {
	r uint8
	g uint8
	b uint8
}

// ScreenBuffer defines an array of pixels, 'width' wide by 'height' high
type ScreenBuffer struct {
	width  int
	height int
	buffer [][]Pixel
}

// NewPixel(r,g,b) returns a Pixel{} struct as defined by r,g,b where each
// element has a range of 0-255
func NewPixel(r, g, b uint8) *Pixel {
	col := &Pixel{}
	col.r = r
	col.g = g
	col.b = b
	return col
}

// NewPixelByColour(r,g,b) returns a Pixel{} struct as defined by r,g,b where each
// element is a float32 in the range of 0.0-1.0. Values greater than 1.0 will be
// clipped on conversion to the 0-255 required by the Pixel
func NewPixelByColour(r, g, b float32) *Pixel {
	ir := uint8(255.0 * r)
	ig := uint8(255.0 * g)
	ib := uint8(255.0 * b)
	return NewPixel(ir, ig, ib)
}

// NewScreenBuffer(width, height) returns a ScreenBuffer{} struct
func NewScreenBuffer(width, height int) *ScreenBuffer {
	sb := &ScreenBuffer{}

	sbRGB := make([][]Pixel, height)
	for y := 0; y < height; y++ {
		sbRGB[y] = make([]Pixel, width)
		for x := 0; x < width; x++ {
			sbRGB[y][x] = Pixel{0, 0, 0}
		}
	}

	sb.buffer = sbRGB
	sb.width = width
	sb.height = height

	return sb
}

// sb.SetPixel(x,y, pixel) sets the (x,y) pixel of the ScreenBuffer to the
// Colour{} specified
func (sb *ScreenBuffer) SetPixel(x, y int, pixel *Pixel) {
	bfr := sb.buffer
	bfr[y][x].r = pixel.r
	bfr[y][x].g = pixel.g
	bfr[y][x].b = pixel.b
}

// sb.WriteToPPM(fn) outputs the ScreenBuffer{} to the file specified
func (sb *ScreenBuffer) WriteToPPM(fn string) {
	f, err := os.Create(fn)
	if err != nil {
		log.Fatalf("error creating ppm file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", sb.width, sb.height))
	if err != nil {
		log.Fatalf("error writing to ppm file: %v", err)
	}

	for y := 0; y < sb.height; y++ {
		for x := 0; x < sb.width; x++ {
			px := sb.buffer[y][x]
			_, err = f.WriteString(fmt.Sprintf("%d %d %d\n", px.r, px.g, px.b))
			if err != nil {
				log.Fatalf("error writing pixel to ppm file: %v", err)
			}
		}
	}
}
