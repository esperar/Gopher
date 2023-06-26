package main

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // 인터페이스 메서드 호출
		println(a)
	}
}

func main() {
	r := Rect{10, 20}
	c := Circle{10}

	showArea(r, c)
}
