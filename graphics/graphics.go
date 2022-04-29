package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawCircle(dst *ebiten.Image, cx, cy, r int, clr color.Color) {
	xx := 128 * r
	yy := 0
	x := 0
	y := 0
	for yy <= xx {
		x = xx / 128
		y = yy / 128
		dst.Set(cx + x, cy + y, clr)
		dst.Set(cx - x, cy - y, clr)
		dst.Set(cx - x, cy + y, clr)
		dst.Set(cx + x, cy - y, clr)
		dst.Set(cx + y, cy + x, clr)
		dst.Set(cx - y, cy - x, clr)
		dst.Set(cx - y, cy + x, clr)
		dst.Set(cx + y, cy - x, clr)

		yy += xx >> 7
		xx -= yy >> 7
		// yy += xx / 128
		// xx -= yy / 128
	}

}
// func DrawCircleFilled(dst *ebiten.Image, x, y, r int, clr color.Color) {
// 	for ix := 0; i

// }