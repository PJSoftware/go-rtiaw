package main

import (
	"fmt"
	"os"

	"github.com/pjsoftware/go-rtiaw/lib/camera"
)

func main() {
	cam := camera.NewCamera("16:9",1280)

	for y := 0; y < cam.ImgHeight; y++ {
		fmt.Printf("\rScan lines remaining: %d   ", cam.ImgHeight-y)
		for x := 0; x < cam.ImgWidth; x++ {
			cam.CalcPixel(x, y)
		}
	}
	fmt.Print("\rBuffer generation complete!          \n")

	os.MkdirAll("./ppm", 0755)
	cam.Image.WriteToPPM("./ppm/camera.ppm")

	fmt.Println("Image buffer written to file.")
}
