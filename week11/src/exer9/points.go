package exer9

// TODO: Point (with everything from exercise 8) and methods that modify them in-place
import (
	"fmt"
	"math"
)

// TODO: The Point struct, NewPoint function, .String and .Norm methods
type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	p := Point{x, y}
	return p
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p Point) Norm() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) Scale(f float64) {
	p.x = p.x * f
	p.y = p.y * f
}

func (p *Point) Rotate(a float64) {
	x := p.x
	y := p.y
	p.x = x*math.Cos(a) - y*math.Sin(a)
	p.y = x*math.Sin(a) + y*math.Cos(a)
}
