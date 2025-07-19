package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func StoreImage(img image.Image, filename string) {
	// Use an image encoder to save the image to a file
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		fmt.Println("Error saving image:", err)
	} else {
		fmt.Println("Image saved as ", filename)
	}
}
