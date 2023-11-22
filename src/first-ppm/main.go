package main

import (
	"fmt"
	"os"

	sb "github.com/pjsoftware/go-rtiaw/lib/screenBuffer"
)

const imageWidth = 400
const imageHeight = 400

func main() {
	screen := sb.NewScreenBuffer(imageWidth, imageHeight)

	for y := 0; y < imageHeight; y++ {
		fmt.Printf("\rScan lines remaining: %d   ", imageHeight-y)
		for x := 0; x < imageWidth; x++ {
			pixel := sb.NewPixelByColour(
				float32(x)/float32(imageWidth-1),
				float32(y)/float32(imageHeight-1),
				0.0)
			screen.SetPixel(x, y, pixel)
		}
	}
	fmt.Print("\rBuffer generation complete!          \n")

	os.MkdirAll("./ppm", 0755)
	screen.WriteToPPM("./ppm/first.ppm")

	fmt.Println("Image buffer written to file.")
}
