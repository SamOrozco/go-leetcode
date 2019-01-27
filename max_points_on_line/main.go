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

	println("pass")
}

func maxPoints(points []Point) int {
	pointLen := len(points)
	if pointLen == 0 {
		return pointLen
	}
	slopes := make(map[int]int, len(points))
	slopeMax := 0
	sameCount := 0
	for i, p0 := range points {
		for _, p1 := range points[i+1:] {
			curSlop := slope(p0, p1)
			if same(p0, p1) {
				sameCount++
			}
			if val, ok := slopes[curSlop]; ok { // if slop exists
				newVal := val + 1
				slopes[curSlop] = newVal
				if newVal > slopeMax {
					slopeMax = newVal
				}
			} else {
				slopeMax = 1
				slopes[curSlop] = 1
			}
		}
	}
	return slopeMax + sameCount
}

func slope(p0, p1 Point) int {
	xDiff := p0.X - p1.X
	yDiff := p0.Y - p1.Y
	if xDiff == 0 || yDiff == 0 {
		return 0
	}
	return xDiff / yDiff
}

func same(p0, p1 Point) bool {
	return p0.X == p1.X && p0.Y == p1.Y
}
