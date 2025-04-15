package task

import "fmt"

type Tri1 struct {
	base   float64
	height float64
}
type Square1 struct {
	side float64
}

type Shapes1 interface {
	getArea() float64
}

func TaskShapes() {
	t := Tri1{base: 10, height: 10}
	s := Square1{side: 10}
	printArea(t)
	printArea(s)
}

func (s Square1) getArea() float64 {
	return s.side * s.side
}
func (t Tri1) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s Shapes1) {
	fmt.Println(s.getArea())
}
