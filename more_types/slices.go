package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mypic := make([][]uint8, dy)
	for y := range mypic {
		mypic[y] = make([]uint8, dx)
		for x := range mypic[y] {
			mypic[y][x] = uint8(x * y)
		}
	}

	return mypic
}

func main() {
	pic.Show(Pic)
}
