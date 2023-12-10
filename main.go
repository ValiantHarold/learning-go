package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func pixelate(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	blockSize := int(width / 200)

	pixelated := image.NewRGBA(bounds)

	for y := 0; y < height; y += blockSize {
		for x := 0; x < width; x += blockSize {
			avgColor := averageColor(img, x, y, blockSize)

			fillBlock(pixelated, x, y, blockSize, avgColor)
		}
	}

	return pixelated
}

func averageColor(img image.Image, startX, startY, blockSize int) color.RGBA {
	var totalR, totalG, totalB, count int

	for y := startY; y < startY+blockSize && y < img.Bounds().Dy(); y++ {
		for x := startX; x < startX+blockSize && x < img.Bounds().Dx(); x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			totalR += int(r >> 8)
			totalG += int(g >> 8)
			totalB += int(b >> 8)
			count++
		}
	}

	avgR := uint8(totalR / count)
	avgG := uint8(totalG / count)
	avgB := uint8(totalB / count)

	return color.RGBA{avgR, avgG, avgB, 255}
}

func fillBlock(img *image.RGBA, startX, startY, blockSize int, avgColor color.RGBA) {
	for y := startY; y < startY+blockSize && y < img.Bounds().Dy(); y++ {
		for x := startX; x < startX+blockSize && x < img.Bounds().Dx(); x++ {
			img.Set(x, y, avgColor)
		}
	}
}

func main() {
	// Open image file
	file, err := os.Open("images/chad.png")

	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Pixelate image
	pixelated := pixelate(img)

	// Make new file
	outFile, err := os.Create("images/output.png")

	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Encode image
	err = png.Encode(outFile, pixelated)
	if err != nil {
		fmt.Println("Error encoding image:", err)
	}
}
