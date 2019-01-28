package main

import "os"

func main() {

	//["MyCalendarTwo","book","book","book","book","book","book"]
	//[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]
	//[null,true,true,true,false,true,true]
	cal := Constructor()
	if !cal.Book(10, 20) {
		println("Fail")
		os.Exit(1)
	}

	if !cal.Book(50, 60) {
		println("Fail")
		os.Exit(1)
	}

	if !cal.Book(10, 40) {
		println("Fail")
		os.Exit(1)
	}

	if cal.Book(5, 15) {
		println("Fail")
		os.Exit(1)
	}

	if !cal.Book(5, 10) {
		println("Fail")
		os.Exit(1)
	}

	if !cal.Book(25, 55) {
		println("Fail")
		os.Exit(1)
	}
}

type MyCalendarTwo struct {
	pairs []Pair
}

type Pair struct {
	x int
	y int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	pair := Pair{
		x: start,
		y: end,
	}
	cnt := 0
	for _, p := range this.pairs {
		if intersect(p, pair) {
			cnt ++
			if cnt >= 2 {
				return false
			}
		}
	}
	this.pairs = append(this.pairs, pair)
	return true
}

func intersect(p1, p2 Pair) bool {
	x1 := p1.x
	y1 := p1.y
	x2 := p2.x
	y2 := p2.y
	if ((x1 >= x2 && x1 < y2) || (y1 > x2 && y1 < y2)) ||
		((x2 >= x1 && x2 < y1) || (y2 > x1 && y2 < y1)) {
		return true
	}
	return false
}
