package interfaces

import (
	"fmt"
	"math"
)

type shapes interface {
	Area() float64
}

type Circle struct {
	redius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.redius * c.redius
}
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.redius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return 2*math.Pi*r.height + r.width
}
func school(s shapes) {
	fmt.Println("seklin alani : ", s.Area())
}
func DemoShape() {
	r := Rectangle{width: 10, height: 6}
	school(r)
}
