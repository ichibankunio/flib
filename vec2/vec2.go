package vec2

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func New(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return New(v.X+other.X, v.Y+other.Y)
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return New(v.X-other.X, v.Y-other.Y)
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y
}

func VecFromAngle(angle, magnitude float64) Vec2 {
	return Vec2{
		math.Cos(angle) * magnitude,
		math.Sin(angle) * magnitude,
	}
}

func (v Vec2) Clone() Vec2 {
	return Vec2{
		X: v.X,
		Y: v.Y,
	}
}
