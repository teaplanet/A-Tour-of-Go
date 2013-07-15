package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{
	width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel	
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) % 255)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
