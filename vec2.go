package flib

import (
	"math"
	"math/rand"
)

type Vec2 struct {
	X, Y float64
}

type Vec2i struct {
	X, Y int
}

func NewVeci(x, y int) *Vec2i {
	return &Vec2i{
		X: x,
		Y: y,
	}
}

func NewVec(x, y float64) *Vec2 {
	return &Vec2{
		X: x,
		Y: y,
	}
}

func (v Vec2) Clone() *Vec2 {
	return &Vec2{
		X: v.X,
		Y: v.Y,
	}
}

func (v Vec2i) Clone() *Vec2i {
	return &Vec2i{
		X: v.X,
		Y: v.Y,
	}
}

func (v *Vec2) Add(other *Vec2) *Vec2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vec2i) Add(other *Vec2i) *Vec2i {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vec2) AddScalar(scalar float64) *Vec2 {
	v.X += scalar
	v.Y += scalar
	return v
}

func (v *Vec2) Sub(other *Vec2) *Vec2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v *Vec2) SubScalar(scalar float64) *Vec2 {
	v.X -= scalar
	v.Y -= scalar
	return v
}

func (v *Vec2) Normalize() *Vec2 {
	len := v.Length()
	if len != 0 {
		v.X /= len
		v.Y /= len
	}
	return v
}

func (v *Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vec2) Scale(s float64) *Vec2 {
	v.X *= s
	v.Y *= s
	return v
}

func VecFromAngle(angle, magnitude float64) *Vec2 {
	return &Vec2{
		math.Cos(angle) * magnitude,
		math.Sin(angle) * magnitude,
	}
}

func (v *Vec2) Equals(other *Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

func RandomDirection() *Vec2 {
	return (&Vec2{
		rand.Float64() - 0.5,
		rand.Float64() - 0.5,
	}).Normalize()
}

func (v *Vec2) Heading() float64 {
	return math.Atan2(v.X, v.Y)
}
