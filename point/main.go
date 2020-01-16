package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "x = 42, "
	ptr.y = "y = 21"

}

func main() {
	points := &point{}

	setPoint(points)
	res := points.x + points.y
	for _, e := range res {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
