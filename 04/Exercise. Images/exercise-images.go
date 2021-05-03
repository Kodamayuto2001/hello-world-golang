package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

//	わからん
// func (img Image) ColorModel() color.RGBA {

// }

// func (img Image) Bounds() image.Rectangle {

// }

// func (img Image) At(x, y int) color.RGBA {

// }

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }
