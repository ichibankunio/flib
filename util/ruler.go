package util

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func DrawRuler(screen *ebiten.Image, clr color.Color, scale float64) {
	_, h := screen.Size()

	for i := 2; i < 10; i++ {
		for j := 1; j < i; j++ {
			reducible := false
			for k := 2; k < i-1; k++ {
				if i % (i-k) == 0 && j % (i-k) == 0 {
					reducible = true
					// println("reducible", i, j, i-k)
					break
				}
			}
			
			if reducible {
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scale, scale)
			r, g, b, a := clr.RGBA()
			op.ColorScale.Scale(float32(r), float32(g), float32(b), float32(a))
			s := fmt.Sprintf("%d/%d", j, i)
			bound := text.BoundString(bitmapfont.Face, s)
			op.GeoM.Translate(4+float64(-bound.Min.X)*scale, float64(h*j/i)+float64(-bound.Min.Y)*scale)
			text.DrawWithOptions(screen, fmt.Sprintf("%d/%d", j, i), bitmapfont.Face, op)
		}
	}
}

func drawLineWithShadow(dst *ebiten.Image, x1, y1, x2, y2 float64) {
	ebitenutil.DrawLine(dst, x1, y1+2, x2, y2+2, color.White)
	ebitenutil.DrawLine(dst, x1, y1, x2, y2, color.Black)
}
