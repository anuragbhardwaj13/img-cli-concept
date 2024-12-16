package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"

	"os"
)

func main() {
	imagePath := "test.jpg"

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("os error: %s\n", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("image error: %s\n", err)
		return
	}

	// Print image bounds and format
	fmt.Println("Image Bounds:", img.Bounds())
	fmt.Println("Image Format:", format)

	// saveAsPng(img)
	// resizeImage(img, 500, 300)
	// rotate90(img)

}

func rotate90(img image.Image) {
	bounds := img.Bounds()
	rotatedImg := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rotatedImg.Set(bounds.Max.Y-y-1, x, img.At(x, y))
		}
	}
	outputFile, err := os.Create("rotated.jpg")
	if err != nil {
		fmt.Printf("output image error: %s\n", err)
		return
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, rotatedImg, nil)
	if err != nil {
		fmt.Printf("save image error: %s\n", err)
		return
	}
	fmt.Printf("Image saved as %s\n", outputFile)

}

func saveAsPng(img image.Image) {
	outputFile, err := os.Create("test.png")
	if err != nil {
		fmt.Printf("output image error: %s\n", err)
		return
	}

	defer outputFile.Close()
	err = png.Encode(outputFile, img)
	if err != nil {
		fmt.Printf("Error encoding PNG: %v\n", err)
		return
	}
	fmt.Printf("Image saved as %s\n", outputFile)
}

func resizeImage(img image.Image, newH, newW int) {
	newImage := image.NewRGBA(image.Rect(0, 0, newW, newH))

	draw.Draw(newImage, newImage.Bounds(), img, image.Point{}, draw.Over)

	outputFile, err := os.Create("resize.png")
	if err != nil {
		fmt.Printf("output image error: %s\n", err)
		return
	}
	defer outputFile.Close()
	err = png.Encode(outputFile, newImage)
	if err != nil {
		fmt.Printf("Error encoding PNG: %v\n", err)
		return
	}
	fmt.Printf("Image saved as %s\n", outputFile.Name())
}
