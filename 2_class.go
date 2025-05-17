package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (q *Point) Distance(P Point) float64 {
	return math.Hypot(q.X-P.X, q.Y-P.Y)
}
func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}

func CalculateDistance() {
	p := Point{1, 2}
	q := Point{4, 6}
	distance1 := p.Distance(q)
	fmt.Printf("Distance: %f\n", distance1)

	distance2 := (*Point).Distance
	fmt.Println(distance2(&p, q))
	fmt.Printf("Distance: %T\n", distance2)
}

type Path []Point

func (path Path) TranslateBy(anther Point, add bool) {
	var op func(q, p Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], anther)
		path[i].Print()
	}
}

func test() {
	points := Path{
		{1, 2},
		{3, 4},
	}

	anotherPoint := Point{5, 5}
	points.TranslateBy(anotherPoint, true)
	fmt.Println("-------------")
	points.TranslateBy(anotherPoint, false)
}

func main() {
	// CalculateDistance()
	test()
}
