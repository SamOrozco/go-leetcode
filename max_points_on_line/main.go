package main

import (
	"os"
)

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

	points1 := []Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 0,
			Y: 0,
		},
	}

	if 2 != maxPoints(points1) {
		println("Failed")
		os.Exit(1)
	}

	points2 := []Point{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 3,
			Y: 2,
		},
		{
			X: 5,
			Y: 3,
		},
		{
			X: 4,
			Y: 1,
		},
		{
			X: 2,
			Y: 3,
		},
		{
			X: 1,
			Y: 4,
		},
	}

	if 4 != maxPoints(points2) {
		println("Failed")
		os.Exit(1)
	}

	//[[0,0],[94911151,94911150],[94911152,94911151]]

	points3 := []Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 94911151,
			Y: 94911150,
		},
		{
			X: 94911152,
			Y: 94911151,
		},
	}


	if 2 != maxPoints(points3) {
		println("Failed")
		os.Exit(1)
	}

	println("pass")
}

func maxPoints(points []Point) int {
	pointLen := len(points)
	if pointLen == 0 {
		return pointLen
	}
	max := 0
	for i, p0 := range points {
		innerMax := 0
		sameCount := 0
		slopes := make(map[float32]int, len(points))
		for _, p1 := range points[i+1:] {
			curSlop := slope(p0, p1)
			if same(p0, p1) {
				sameCount++
				continue
			}
			if val, ok := slopes[curSlop]; ok { // if slop exists
				newVal := val + 1
				slopes[curSlop] = newVal
				if newVal > innerMax {
					innerMax = newVal
				}
			} else {
				if innerMax == 0 {
					innerMax = 1
				}
				slopes[curSlop] = 1
			}
		}
		max = maxInt(max, innerMax+sameCount)
	}
	return max
}

func slope(p0, p1 Point) float32 {
	if p0.X == p1.X { // horizontal
		return 1<<31 - 1
	}

	if p0.Y == p1.Y { // vertical
		return 0
	}
	return float32(p1.X-p0.X) / float32(p1.Y-p0.Y)
}

func same(p0, p1 Point) bool {
	return p0.X == p1.X && p0.Y == p1.Y
}

func maxInt(vals ...int) int {
	max := 0
	for _, i := range vals {
		if i > max {
			max = i
		}
	}
	return max
}
