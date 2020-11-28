package exer10

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func DrawCircle(outerRadius, innerRadius int, outputFile string) {
	bound := image.Rect(0, 0, 200, 200)
	grid := image.NewRGBA(bound)

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	// adapat the correct way to iterate over an Image m's pixels from
	// https://blog.golang.org/image
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			grid.Set(x, y, white)
			if math.Sqrt(float64((y-100)*(y-100)+(100-x)*(100-x))) <= float64(outerRadius) {
				grid.Set(x, y, black)
			}
		}
	}

	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			if math.Sqrt(float64((y-100)*(y-100)+(100-x)*(100-x))) <= float64(innerRadius) {
				grid.Set(x, y, white)
			}
		}
	}

	myFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	png.Encode(myFile, grid)
}
