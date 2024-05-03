package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Height, Width int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
