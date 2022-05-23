package tiled

import (
	"math"

	"github.com/ichibankunio/flib"
)

type Rectangle struct {
	Pos  *flib.Vec2
	Size *flib.Vec2
}

func (r *Rectangle) IsIntersect(other *Rectangle) bool {
	if math.Abs(r.Pos.X-other.Pos.X) > r.Size.X/2+other.Size.X/2 {
		return false
	}
	if math.Abs(r.Pos.Y-other.Pos.Y) > r.Size.Y/2+other.Size.Y/2 {
		return false
	}
	return true
}
