package exer10

import (
	"fmt"
	"math"
)

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

type Triangle struct {
	A, B, C Point
}

func (t Triangle) String() string {
	return fmt.Sprintf("[%s %s %s]", t.A, t.B, t.C)
}

func (t *Triangle) Scale(s float64) {
	t.A.Scale(s)
	t.B.Scale(s)
	t.C.Scale(s)
}

func (t *Triangle) Rotate(a float64) {
	t.A.Rotate(a)
	t.B.Rotate(a)
	t.C.Rotate(a)
}

type Transformable interface {
	Scale(s float64)
	Rotate(a float64)
}

func TurnDouble(t Transformable, angle float64) {
	t.Scale(2)
	t.Rotate(angle)
}
