package main

//the code can run in the go tour background uncommenting the pic library and the
//execution in the func main
import (
	//	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h           int
	colorB, colorA uint8
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), i.colorB, i.colorA}
}

func main() {
	//m := &Image{111, 111, 255, 255}
	//pic.ShowImage(m)
}
