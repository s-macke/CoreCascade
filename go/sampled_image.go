package main

import (
	"cascade/filebuffer"
	"fmt"
	"image"
	"image/png"
	"os"
)

type SampledPixel struct {
	Color   Color
	Samples int
}

type SampledImage struct {
	pixels        [][]SampledPixel
	width, height int
}

func NewSampledImage(width, height int) *SampledImage {
	s := &SampledImage{
		width:  width,
		height: height,
	}
	s.pixels = make([][]SampledPixel, height)
	for i := range s.pixels {
		s.pixels[i] = make([]SampledPixel, width)
		for j := range s.pixels[i] {
			s.pixels[i][j] = SampledPixel{Color: Color{R: 0, G: 0, B: 0}, Samples: 0}
		}
	}
	return s
}

func NewSampledImageFromFile(filename string) *SampledImage {
	rb := filebuffer.NewReadBufferFromFile(filename)
	header := rb.ReadSliceAsString(32)
	if header != "CoreCascade Sampled Image V1.0  " {
		panic("Invalid sampled image file format")
	}
	width := rb.ReadInt(4)
	height := rb.ReadInt(4)
	s := NewSampledImage(width, height)

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			sp := &s.pixels[y][x]
			sp.Color.R = rb.ReadFloat64()
			sp.Color.G = rb.ReadFloat64()
			sp.Color.B = rb.ReadFloat64()
			sp.Samples = rb.ReadInt(4)
		}
	}

	return s
}

func (s *SampledImage) AddColorSamples(x, y int, col Color, samples int) {
	s.pixels[y][x].Color.Add(col)
	s.pixels[y][x].Samples += samples
}

func (s *SampledImage) SetColor(x, y int, col Color) {
	s.pixels[y][x].Color = col
	s.pixels[y][x].Samples = 1
}

func (s *SampledImage) ToImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			p := s.pixels[y][x]
			c := p.Color
			c.Div(float64(p.Samples))
			img.Set(x, y, c.ToSRGBA())
		}
	}
	return img
}

func (s *SampledImage) Energy() float64 {
	energy := 0.
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			p := s.pixels[y][x]
			c := p.Color
			c.Div(float64(p.Samples))
			energy += c.Intensity()
		}
	}
	return energy
}

func (s *SampledImage) StoreImage(filename string) {
	img := s.ToImage()
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

func (s *SampledImage) StoreRaw(filename string) {
	wb := filebuffer.NewWriteBuffer(32 + 8 + s.width*s.height*(3*8+4))
	wb.WriteString("CoreCascade Sampled Image V1.0  ")
	wb.WriteInt32(int32(s.width))
	wb.WriteInt32(int32(s.height))
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			sp := s.pixels[y][x]
			wb.WriteFloat64(sp.Color.R)
			wb.WriteFloat64(sp.Color.G)
			wb.WriteFloat64(sp.Color.B)
			wb.WriteInt32(int32(sp.Samples))
		}
	}
	wb.StoreToFile(filename)
}

func (s *SampledImage) Store(filename string) {
	s.StoreRaw(filename + ".raw")
	s.StoreImage(filename + ".png")
}
