package linear_image

import (
	"color"
	"filebuffer"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type SampledPixel struct {
	Color   color.Color
	Samples int
}

type SampledImage struct {
	pixels        [][]SampledPixel
	Width, Height int
}

func NewSampledImage(width, height int) *SampledImage {
	s := &SampledImage{
		Width:  width,
		Height: height,
	}
	s.pixels = make([][]SampledPixel, height)
	for i := range s.pixels {
		s.pixels[i] = make([]SampledPixel, width)
		for j := range s.pixels[i] {
			s.pixels[i][j] = SampledPixel{Color: color.Black, Samples: 0}
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

	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			sp := &s.pixels[y][x]
			sp.Color.R = rb.ReadFloat64()
			sp.Color.G = rb.ReadFloat64()
			sp.Color.B = rb.ReadFloat64()
			sp.Samples = rb.ReadInt(4)
		}
	}

	return s
}

func NewSampledImageFromJpeg(filename string) *SampledImage {
	imgFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return nil
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		fmt.Println("Error decoding JPEG image:", err)
		return nil
	}

	bounds := img.Bounds()
	s := NewSampledImage(bounds.Dx(), bounds.Dy())

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			col := color.NewSRGBColor(float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0)
			s.SetColor(x-bounds.Min.X, y-bounds.Min.Y, col)
		}
	}

	return s

}

func (s *SampledImage) Add(s2 *SampledImage) {
	if s.Width != s2.Width || s.Height != s2.Height {
		panic("SampledImage sizes do not match for merging")
	}
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			col1 := s.GetColor(x, y)
			col2 := s2.GetColor(x, y)
			col1.Add(col2)
			s.SetColor(x, y, col1)
		}
	}
}

func (s *SampledImage) Clear() {
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			s.pixels[y][x].Samples = 0
			s.pixels[y][x].Color = color.Black
		}
	}
}

func (s *SampledImage) AddColorSamples(x, y int, col color.Color, samples int) {
	s.pixels[y][x].Color.Add(col)
	s.pixels[y][x].Samples += samples
}

func (s *SampledImage) AddColor(x, y int, col color.Color) {
	s.pixels[y][x].Color.Add(col)
	s.pixels[y][x].Samples = 1
}

func (s *SampledImage) SetColor(x, y int, col color.Color) {
	s.pixels[y][x].Color = col
	s.pixels[y][x].Samples = 1
}

func (s *SampledImage) GetColor(x, y int) color.Color {
	p := s.pixels[y][x]
	c := p.Color
	if p.Samples == 0 {
		return color.Black
	}
	c.Div(float64(p.Samples))
	return c
}

func (s *SampledImage) ToImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, s.Width, s.Height))
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			p := s.pixels[y][x]
			c := p.Color
			c.Div(float64(p.Samples))
			img.Set(x, y, c.ToSRGBAReinhard())
			//img.Set(x, y, c.ToSRGBAOnlyGamma())
		}
	}
	return img
}

func (s *SampledImage) Energy() float64 {
	energy := 0.
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			p := s.pixels[y][x]
			c := p.Color
			c.Div(float64(p.Samples))
			energy += c.Intensity()
		}
	}
	return energy
}

func (s *SampledImage) StorePNG(filename string) {
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

func (s *SampledImage) StoreJPEG(filename string) {
	img := s.ToImage()
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		fmt.Println("Error saving image:", err)
	} else {
		fmt.Println("Image saved as ", filename)
	}
}

func (s *SampledImage) StoreRaw(filename string) {
	wb := filebuffer.NewWriteBuffer(32 + 8 + s.Width*s.Height*(3*8+4))
	wb.WriteString("CoreCascade Sampled Image V1.0  ")
	wb.WriteInt32(int32(s.Width))
	wb.WriteInt32(int32(s.Height))
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
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
	s.StorePNG(filename + ".png")
	//s.StoreJPEG(filename + ".jpg")
}

func (s *SampledImage) Error(img *SampledImage) {
	fmt.Println("Energy of s:", s.Energy())
	fmt.Println("Energy of img:", img.Energy())
	e := 0.0
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			r := (s.Width-img.Width/2)*(s.Width-img.Width/2) + (s.Height-img.Height/2)*(s.Height-img.Height/2)
			if r < 150 {
				continue // skip pixels that are too far from the image
			}
			c := s.GetColor(x, y)
			c2 := img.GetColor(x, y)
			c.Sub(c2)
			c.Abs()
			e += c.Intensity()
			c.Mul(2.)
			s.SetColor(x, y, c)
		}
	}
	fmt.Println("Error:", e/float64(s.Width*s.Height))
}

func (s *SampledImage) Blend(src *SampledImage) {
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			if x >= src.Width || y >= src.Height {
				continue // skip pixels that are out of bounds
			}
			col := s.GetColor(x, y)
			colSrc := src.GetColor(x, y)
			//colSrc := Color{0.5, 0.5, 0.5}
			colSrc.R *= col.R
			colSrc.G *= col.G
			colSrc.B *= col.B
			s.SetColor(x, y, colSrc)
		}
	}
}
