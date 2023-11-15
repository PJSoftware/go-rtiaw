package main

import (
	"os"

	sb "github.com/pjsoftware/go-rtiaw/lib/screenBuffer"
)

const imageWidth = 256
const imageHeight = 256

func main() {
  screen := sb.NewScreenBuffer(imageWidth,imageHeight)
	
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			pixel := sb.NewColour(
				float32(x) / float32(imageWidth-1),
				float32(y) / float32(imageHeight-1),
				0.0 )
			screen.SetPixel(x,y, pixel)
		}
	}

	os.MkdirAll("./ppm", 0755)
	screen.WriteToPPM("./ppm/first.ppm")
}
