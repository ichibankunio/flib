package fui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/flib"
	"github.com/ichibankunio/fvec/vec2"
	"golang.org/x/image/font"
)


type Slider struct {
	target SoundType
	handle *Button
	bar *flib.Sprite
	handleAdjustX int
	moving bool
}

type SoundType int

const (
	SoundBGM = iota
	SoundSE
)


func NewVolumeSlider(centerX, y int, width int, target SoundType, fontFace font.Face, txtClr, clrBound, clrBg color.Color) *Slider {
	label := []string{"BGM", "SE"}[target]
	
	s := &Slider{
		target: target,
		// handle: NewButton(label, centerX + width / 2 - width / 8, y - width / 16 - width/32 , vec2.New(float64(width)/4, float64(width)/4), fontFace, ThemeRound, txtClr, clrBound, clrBg),
		handle: NewButton(label, vec2.New(float64(centerX) + float64(width) / 2 - float64(width) / 8, float64(y) - float64(width) / 16 - float64(width)/32), vec2.New(float64(width)/4, float64(width)/4), fontFace, ThemeRound, txtClr, clrBound, clrBg),
		bar: flib.NewSprite(NewButtonImg(width, width/16, ThemeRound, clrBound, clrBg), vec2.New(float64(centerX - width / 2), float64(y))),
	}

	switch target {
	case SoundBGM:
		s.handle.Spr.Pos.X = s.bar.Pos.X + flib.GetBGMVolume() * float64(width - s.handle.Spr.Img.Bounds().Dx())
	case SoundSE:
		s.handle.Spr.Pos.X = s.bar.Pos.X + flib.GetSEVolume() * float64(width - s.handle.Spr.Img.Bounds().Dx())
	}

	// s.handle.Txt.SetCenter(int(s.handle.Spr.Pos.X) + s.handle.Spr.Img.Bounds().Dx() / 2)


	return s
}

func (s *Slider) Draw(screen *ebiten.Image) {
	s.bar.Draw(screen)
	s.handle.Draw(screen)
}

func (s *Slider) Update(g *flib.Game) {
	s.handle.Spr.Update()
	
	if s.handle.Spr.IsJustTouched() {
		x, _ := ebiten.TouchPosition(s.handle.Spr.JustPressedTouchID)
		// s.handleAdjustX = x - int(s.handle.spr.X)

		if x == 0 {
			x, _ = ebiten.CursorPosition()
		}
		// x, y := ebiten.CursorPosition()
		s.handleAdjustX = x - int(s.handle.Spr.Pos.X)
		// if x == 0 && y == 0 {
		// 	x, _ = ebiten.TouchPosition(s.handle.spr.TouchID)
		// 	s.handleAdjustX = x - int(s.handle.spr.X)
		// }
		s.moving = true

	}

	if is, _ := s.handle.Spr.IsTouchJustReleased(); is {
		s.moving = false
	}

	if s.moving {
		x, _ := ebiten.TouchPosition(s.handle.Spr.JustPressedTouchID)

		if x == 0 {
			x, _ = ebiten.CursorPosition()
		}
		// s.handle.Spr.Pos = complex(float64(x) - float64(s.handleAdjustX), imag(s.handle.Spr.Ps))
		s.handle.Spr.Pos.X = float64(x) - float64(s.handleAdjustX)
		
		if s.handle.Spr.Pos.X < s.bar.Pos.X {
			s.handle.Spr.Pos.X = s.bar.Pos.X
			
		}
		if s.handle.Spr.Pos.X > s.bar.Pos.X + float64(s.bar.Img.Bounds().Dx()) - float64(s.handle.Spr.Img.Bounds().Dx()) { 
			s.handle.Spr.Pos.X = s.bar.Pos.X + float64(s.bar.Img.Bounds().Dx()) - float64(s.handle.Spr.Img.Bounds().Dx())
		}
		// s.handle.Txt.SetCenter(int(s.handle.Spr.Pos.X) + s.handle.Spr.Img.Bounds().Dx() / 2)

		scale := (s.handle.Spr.Pos.X - s.bar.Pos.X) / float64(s.bar.Img.Bounds().Dx() - s.handle.Spr.Img.Bounds().Dx())
		switch s.target {
		case SoundBGM:
			flib.SetBGMVolume(scale)
		case SoundSE:
			flib.SetSEVolume(scale)
		}
		// println(scale)
	}

	s.bar.Update()
	s.handle.Update(g)
}