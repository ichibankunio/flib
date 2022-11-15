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

func NewArray(arr [][2]float64) []Vec2 {
	out := make([]Vec2, len(arr))
	for i, _ := range arr {
		out[i] = New(arr[i][0], arr[i][1])
	}

	return out
}

func (v Vec2) Heading() float64 {
	return math.Atan2(v.X, v.Y)
}

func (v Vec2) Scale(s float64) Vec2 {
	return New(v.X*s, v.Y*s)
}

func (v Vec2) Mul(other Vec2) Vec2 {
	return New(v.X * other.X, v.Y * other.Y)
}

func (v Vec2) Floor() Vec2 {
	return New(math.Floor(v.X), math.Floor(v.Y))
}

func (v Vec2) Abs() Vec2 {
	return New(math.Abs(v.X), math.Abs(v.Y))
}

func (v Vec2) Sign() Vec2 {
	sign := New(1, 1)
	if v.X < 0 {
		sign.X = -1
	}
	if v.Y < 0 {
		sign.Y = -1
	}

	return sign
}