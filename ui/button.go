package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/ichibankunio/flib"
	"github.com/ichibankunio/fvec/vec2"
	"golang.org/x/image/font"
)

type UITheme int

const (
	ThemeRect UITheme = iota
	ThemeRound
	ThemeShadow
	ThemeTabLeft
	ThemeTabRight
	ThemeTabCenter
)

type Button struct {
	Spr *flib.Sprite
	// Txt *flib.Text
	Txt string
	fontFace text.Face
	OnClick func(*flib.Game)
	OnRelease func(*flib.Game)
	IsClickInProgress bool

	isJustReleased bool
	
}

func NewButton(txt string, centerX int, y int, width int, height int, font font.Face, theme UITheme, txtClr, clrBound, clrBg color.Color) *Button {
	return &Button{
		Spr: flib.NewSprite(NewButtonImg(width, height, theme, clrBound, clrBg), vec2.New(float64(centerX - width / 2), float64(y))),
		Txt: txt,
		// Txt: flib.NewText(txt, vec2.New(float64(centerX - text.BoundString(fontface, txt).Dx()/2), float64(y + height / 2 - text.BoundString(fontface, txt).Dy()/2)), txtClr, fontface),
		OnClick: func(*flib.Game){},
		OnRelease: func(*flib.Game){},
		IsClickInProgress: false,
		fontFace: text.NewGoXFace(font),
	}
}

func (b *Button) Translate(x, y float64) {
	b.Spr.Pos = vec2.New(x, y)

	// b.Txt.SetCenter(int(x) + b.Spr.Img.Bounds().Dx() / 2)
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.Spr.Draw(screen)
	// b.Txt.Draw(screen)
	op := &text.DrawOptions{}
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	text.Draw(screen, b.Txt, b.fontFace, op)
}

func (b *Button) Update(g *flib.Game) {
	b.Spr.Update()

	if b.Spr.IsTouched() {
		// b.spr.Alpha = 0.7
	}
	if b.Spr.IsJustTouched() {
		// b.spr.Alpha = 0.7
		b.Spr.DrawOption = func(op *ebiten.DrawImageOptions) {
			op.ColorM.ChangeHSV(0, 1, 0.7)
		} 

		b.OnClick(g)

		b.IsClickInProgress = true
	}
	if b.isJustReleased {
		b.isJustReleased = false
	}
	if isTouchJustReleased, isStillTouched := b.Spr.IsTouchJustReleased(); isTouchJustReleased {
		if isStillTouched && b.IsClickInProgress {
			b.OnRelease(g)
			b.isJustReleased = true
		}
		// b.spr.Alpha = 1
		b.Spr.DrawOption = func(op *ebiten.DrawImageOptions) {
			op.ColorM.ChangeHSV(0, 1, 1)
		}
		b.IsClickInProgress = false

	}
}

func (b *Button) IsTouchJustReleased() bool {
	isTouchJustReleased, isStillTouched := b.Spr.IsTouchJustReleased()
	return (isTouchJustReleased && isStillTouched && b.IsClickInProgress) || b.isJustReleased
}

func NewButtonImg(width, height int, theme UITheme, clrBound, clrBg color.Color) *ebiten.Image {
	bg := ebiten.NewImage(width, height)
	src := ebiten.NewImage(1, 1)
	src.Fill(clrBound)

	src2 := ebiten.NewImage(1, 1)
	src2.Fill(clrBg)

	var path vector.Path
	var path2 vector.Path

	w := float32(width)
	h := float32(height)
	l := float32(math.Min(float64(width), float64(height)) / 16)

	switch theme {
	case ThemeRect:
		path.MoveTo(0, 0)
		path.LineTo(w, 0)
		path.LineTo(w, h)
		path.LineTo(0, h)
		path.LineTo(0, 0)

		l /= 2
		path2.MoveTo(l, l)
		path2.LineTo(w - l, l)
		path2.LineTo(w - l, h - l)
		path2.LineTo(l, h - l)
		path2.LineTo(l, l)

	case ThemeRound:
		path.MoveTo(l, 0)
		path.LineTo(w - l, 0)
		path.ArcTo(w, 0, w, l, l)
		path.LineTo(w, h-l)
		path.ArcTo(w, h, w-l, h, l)
		path.LineTo(l, h)
		path.ArcTo(0, h, 0, h-l, l)
		path.LineTo(0, l)
		path.ArcTo(0, 0, l, 0, l)

		l /= 2
		path2.MoveTo(l*2, l)
		path2.LineTo(w - l*2, l)
		path2.ArcTo(w-l, l, w-l, l*2, l)
		path2.LineTo(w-l, h-l*2)
		path2.ArcTo(w-l, h-l, w-l*2, h-l, l)
		path2.LineTo(l*2, h-l)
		path2.ArcTo(l, h-l, l, h-l*2, l)
		path2.LineTo(l, l*2)
		path2.ArcTo(l, l, l*2, l, l)

	case ThemeShadow:
		path.MoveTo(l, 0)
		path.LineTo(w - l, 0)
		path.ArcTo(w, 0, w, l, l)
		path.LineTo(w, h-l)
		path.ArcTo(w, h, w-l, h, l)
		path.LineTo(l, h)
		path.ArcTo(0, h, 0, h-l, l)
		path.LineTo(0, l)
		path.ArcTo(0, 0, l, 0, l)

		path2.MoveTo(l, 0)
		path2.LineTo(w - l, 0)
		path2.ArcTo(w, 0, w, l, l)
		path2.LineTo(w, h- l*2)
		path2.ArcTo(w, h-l, w-l, h-l, l)
		path2.LineTo(l, h-l)
		path2.ArcTo(0, h-l, 0, h-l*2, l)
		path2.LineTo(0, l)
		path2.ArcTo(0, 0, l, 0, l)
	
	case ThemeTabLeft:
		path.MoveTo(l, 0)
		path.LineTo(w, 0)
		path.LineTo(w, h)
		path.LineTo(l, h)
		path.ArcTo(0, h, 0, h-l, l)
		path.LineTo(0, l)
		path.ArcTo(0, 0, l, 0, l)

		l /= 2
		path2.MoveTo(l*2, l)
		path2.LineTo(w, l)
		path2.LineTo(w, h-l)
		path2.LineTo(l*2, h-l)
		path2.ArcTo(l, h-l, l, h-l*2, l)
		path2.LineTo(l, l*2)
		path2.ArcTo(l, l, l*2, l, l)
	case ThemeTabRight:
		path.MoveTo(0, 0)
		path.LineTo(w - l, 0)
		path.ArcTo(w, 0, w, l, l)
		path.LineTo(w, h-l)
		path.ArcTo(w, h, w-l, h, l)
		path.LineTo(0, h)
		path.LineTo(0, 0)

		l /= 2
		path2.MoveTo(0, l)
		path2.LineTo(w - l*2, l)
		path2.ArcTo(w-l, l, w-l, l*2, l)
		path2.LineTo(w-l, h-l*2)
		path2.ArcTo(w-l, h-l, w-l*2, h-l, l)
		path2.LineTo(0, h-l)
		path2.LineTo(0, l)
	case ThemeTabCenter:
		path.MoveTo(0, 0)
		path.LineTo(w, 0)
		path.LineTo(w, h)
		path.LineTo(0, h)
		path.LineTo(0, 0)

		l /= 2
		path2.MoveTo(0, l)
		path2.LineTo(w, l)
		path2.LineTo(w, h-l)
		path2.LineTo(0, h-l)
		path2.LineTo(0, l)	
	}
	
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	bg.DrawTriangles(vs, is, src, op)

	op2 := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs2, is2 := path2.AppendVerticesAndIndicesForFilling(nil, nil)
	bg.DrawTriangles(vs2, is2, src2, op2)
	
	return bg
}