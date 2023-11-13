package main

import "fmt"

const imageWidth = 256
const imageHeight = 256

func main() {
	fmt.Printf("P3\n%d %d\n%d\n", imageWidth, imageHeight, 255)
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			r := float32(x) / float32(imageWidth-1)
			g := float32(y) / float32(imageHeight-1)
			b := float32(0.0)

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}

	}
}
