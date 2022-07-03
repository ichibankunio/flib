package flib

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/ichibankunio/flib/vec2"
	"golang.org/x/image/font"
)

type TxtSpr struct {
	Txt     string
	Spr     *Sprite
	Clr     color.Color
	PadUp   int
	PadLeft int
	Font    font.Face
	isVert  bool
	Hidden  bool
}

type Text struct {
	Txt        string
	Pos        vec2.Vec2
	Clr        color.Color
	Hidden     bool
	Alpha      float64
	Font       font.Face
	DrawOption func(*ebiten.DrawImageOptions)
}

func NewText(txt string, pos vec2.Vec2, clr color.Color, font font.Face) *Text {
	return &Text{
		Txt:        txt,
		Pos:        pos,
		Clr:        clr,
		Hidden:     false,
		Alpha:      1,
		Font:       font,
		DrawOption: func(*ebiten.DrawImageOptions) {},
	}
}

func (t *Text) SetCenter(center int) {
	width := text.BoundString(t.Font, t.Txt).Dx()
	t.Pos.X = float64(center - width/2)
}

func (t *Text) SetCenterY(center int) {
	height := text.BoundString(t.Font, t.Txt).Dy()
	t.Pos.Y = float64(center - height/2)
}

func (t *Text) SetTextWithCenterFixed(txt string) {
	center := t.Pos.X + float64(text.BoundString(t.Font, t.Txt).Dx()/2)
	t.Txt = txt
	width := text.BoundString(t.Font, t.Txt).Dx()
	t.Pos.X = center - float64(width)/2
}

func (t *Text) Draw(screen *ebiten.Image) {
	if !t.Hidden {
		bound := text.BoundString(t.Font, t.Txt).Bounds()
		op := &ebiten.DrawImageOptions{}

		op.ColorM.Scale(ColorToScale(t.Clr))

		t.DrawOption(op)

		op.GeoM.Translate(t.Pos.X+float64(-bound.Min.X), t.Pos.Y+float64(-bound.Min.Y))
		op.ColorM.Scale(1, 1, 1, t.Alpha)

		text.DrawWithOptions(screen, t.Txt, t.Font, op)

	}
}

func ColorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	return float64(uint8(r)) / 255, float64(uint8(g)) / 255, float64(uint8(b)) / 255, float64(uint8(a)) / 255
}
