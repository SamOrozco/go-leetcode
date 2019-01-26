package main

import "os"

type Point struct {
	X int
	Y int
}

func main() {
	points := []Point{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 2,
			Y: 2,
		},
		{
			X: 3,
			Y: 3,
		},
	}

	if 3 != maxPoints(points) {
		println("Failed")
		os.Exit(1)
	}
}

func maxPoints(points []Point) int {
	matrix := make([][]int, 0)
	for _, p := range points {

	}
}
