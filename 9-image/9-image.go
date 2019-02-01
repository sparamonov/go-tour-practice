package main

import (
	"image/color"
	"image"
	"golang.org/x/tour/pic"
)

type Image struct{}

const (
	dx = 256
	dy = 256
)

func (im Image) Bounds() image.Rectangle { return image.Rect(0, 0, 80, 80) }

func (im Image) ColorModel() color.Model { return color.RGBAModel }

func (im Image) At(x, y int) color.Color {
	data := Pic(dx, dy)
	v := data[y][x]
	return color.RGBA{v, v, dx-1, dy-1} 
}

func Pic(dx, dy int) [][]uint8 {
	sl := make([][]uint8, dy, dy)
	
	for i := range sl {
		sl[i] = make([]uint8, dx, dx)
		
		for j := range sl[i] {
			//sl[i][j] = uint8((i + j) / 2)
			sl[i][j] = uint8(i * j)
			//sl[i][j] = uint8(i ^ j)
			//sl[i][j] = uint8((i - j) / 2)
		}
	}
	
	return sl
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
