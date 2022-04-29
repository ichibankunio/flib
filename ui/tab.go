package ui

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/flib"
	"golang.org/x/image/font"
)

type Tab struct {
	Selected int
	Buttons []*Button
	ImageSet [2][]*ebiten.Image
}

func NewTab(n int, selected int, centerX int, y int, width int, height int, fontface font.Face, txtClr, clrBound, clrBg color.Color) *Tab {
	t := &Tab{
		Selected: selected,
		Buttons: []*Button{},
		ImageSet: [2][]*ebiten.Image{},
	}

	for i := 0; i < n; i ++ {
		theme := ThemeTabCenter
		if i == 0 {
			theme = ThemeTabLeft
		}else if i == n - 1 {
			theme = ThemeTabRight
		}
		normalImg := NewButtonImg(width, height, theme, clrBound, clrBg)
		selectedImg := NewButtonImg(width, height, theme, clrBound, clrBound)
		if i == 0 {
			drawLevelImage(normalImg, width, height, clrBound)
			drawLevelImage(selectedImg, width, height, clrBg)
			
		}else if i == 1 {
			drawEndlessImage(normalImg, width, height, clrBound)
			drawEndlessImage(selectedImg, width, height, clrBg)
		}

		t.ImageSet[0] = append(t.ImageSet[0], normalImg)
		t.ImageSet[1] = append(t.ImageSet[1], selectedImg)

		b := NewButton("", centerX - width*n/2 + width / 2  + i * width, y, width, height, fontface, theme, txtClr, clrBound, clrBg)
		if i == selected {
			b.Spr.Img = t.ImageSet[1][i]
		}else {
			b.Spr.Img = t.ImageSet[0][i]
		}

		t.Buttons = append(t.Buttons, b)
	}

	return t
}

func (t *Tab) Draw(screen *ebiten.Image) {
	for _, b := range t.Buttons {
		b.Draw(screen)
	}
}

func (t *Tab) Update(g *flib.Game) {
	for i, b := range t.Buttons {
		b.Update(g)
		if b.IsTouchJustReleased() {
			t.Selected = i
			for i, b := range t.Buttons {
				if i == t.Selected {
					b.Spr.Img = t.ImageSet[1][i]
				}else {
					b.Spr.Img = t.ImageSet[0][i]
				}
			}
		}
	}

}

func (t *Tab) IsTouchInProgress() bool {
	for _, b := range t.Buttons {
		if b.IsClickInProgress {
			return true
		}
	}

	return false
}


func (t *Tab) IsTouchJustReleased() bool {
	for _, b := range t.Buttons {
		if isJustReleased, _ := b.Spr.IsTouchJustReleased(); isJustReleased {
			return true
		}
	}

	return false
}



func drawEndlessImage(dst *ebiten.Image, width, height int, clr color.Color) {
	dc := gg.NewContext(width, height)
	dc.DrawCircle(float64(width/4 + width/8), float64(height/2), float64(width/8))
	// dc.DrawCircle(float64(width/4), float64(height/2), float64(width/6))
	dc.SetLineWidth(float64(width/32))
	dc.SetColor(clr)
	dc.Stroke()
	dc.DrawCircle(float64(width*3/4 - width/8), float64(height/2), float64(width/8))
	// dc.DrawCircle(float64(width/4), float64(height/2), float64(width/6))
	// dc.SetLineWidth(float64(width/32))
	dc.SetColor(clr)
	dc.Stroke()
	// dc.Fill()


	op := &ebiten.DrawImageOptions{}
	dst.DrawImage(ebiten.NewImageFromImage(dc.Image()), op)
}

func drawLevelImage(dst *ebiten.Image, width, height int, clr color.Color) {
	dc := gg.NewContext(width, height)
	// dc.SetLineWidth(float64(width/8))
	dc.SetColor(clr)
	dc.DrawRectangle(float64(width/4+4), float64(height/2+width/16), float64(width*3/7), float64(width/7))
	dc.DrawRectangle(float64(width/4+4+width/7), float64(height/2+width/16-width/7), float64(width*2/7), float64(width/7))
	dc.DrawRectangle(float64(width/4+4+width*2/7), float64(height/2+width/16-width*2/7), float64(width/7), float64(width/7))
	dc.Fill()
	// dc.Fill()


	op := &ebiten.DrawImageOptions{}
	dst.DrawImage(ebiten.NewImageFromImage(dc.Image()), op)
}

