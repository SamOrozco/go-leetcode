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

type Direction func(int, int) (int, int)
type Matrix map[int]map[int]int

var maxX = 0
var maxY = 0

func maxPoints(points []Point) int {
	if len(points) < 1 {
		return 0
	}
	// I want to build a flexible matrix
	// that I can query with indexes
	matrix := make(map[int]map[int]int, 0)
	for _, p := range points {
		if matrix[p.X] == nil {
			matrix[p.X] = make(map[int]int, 0)
		}
		matrix[p.X][p.Y] = matrix[p.X][p.Y] + 1
	}

	max := 0
	for _, p := range points {
		tempMx := maxStraight(p.X, p.Y, matrix, nil, matrix[p.X][p.Y])
		if tempMx > max {
			max = tempMx
		}
	}
	return max
}

func maxStraight(x, y int, matrix Matrix, dir Direction, count int) int {
	if dir == nil {
		t := maxStraight(x, y, matrix, Top(), count)
		tr := maxStraight(x, y, matrix, TopRight(), count)
		r := maxStraight(x, y, matrix, Right(), count)
		br := maxStraight(x, y, matrix, BottomRight(), count)
		b := maxStraight(x, y, matrix, BottomLeft(), count)
		bl := maxStraight(x, y, matrix, Left(), count)
		l := maxStraight(x, y, matrix, TopLet(), count)
		return max(t, tr, r, br, b, bl, l)
	}
	newX, newY := dir(x, y)
	if matrix[newX][newY] > 0 { // if neighbor is a point
		return maxStraight(newX, newY, matrix, dir, count+1)
	} else {
		// if both are out of range
		if x > maxX && y > maxY {

		}
		return count
	}
}

func max(vals ...int) int {
	mx := -1
	for _, val := range vals {
		if val > mx {
			mx = val
		}
	}
	return mx
}

// DIRECTIONS
func Top() Direction {
	return func(x, y int) (int, int) {
		return x, y + 1
	}
}

func Right() Direction {
	return func(x, y int) (int, int) {
		return x + 1, y
	}
}

func TopRight() Direction {
	return func(x, y int) (int, int) {
		return Right()(Top()(x, y))
	}
}

func Bottom() Direction {
	return func(x, y int) (int, int) {
		return x, y - 1
	}
}

func BottomRight() Direction {
	return func(x, y int) (int, int) {
		return Right()(Bottom()(x, y))
	}
}

func Left() Direction {
	return func(x, y int) (int, int) {
		return x - 1, y
	}
}

func BottomLeft() Direction {
	return func(x, y int) (int, int) {
		return Left()(Bottom()(x, y))
	}
}

func TopLet() Direction {
	return func(x, y int) (int, int) {
		return Left()(Top()(x, y))
	}
}
