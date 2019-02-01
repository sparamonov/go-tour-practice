package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sl := make([][]uint8, dy, dy)
	
	for i := range sl {
		sl[i] = make([]uint8, dx, dx)
		
		for j := range sl[i] {
			//sl[i][j] = uint8((i + j) / 2)
			//sl[i][j] = uint8(i * j)
			sl[i][j] = uint8(i ^ j)
			//sl[i][j] = uint8((i - j) / 2)
		}
	}
	
	return sl
}

func main() {
	pic.Show(Pic)
}