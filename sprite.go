package flib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ichibankunio/fvec/vec2"
)

type Sprite struct {
	Img *ebiten.Image
	Pos vec2.Vec2
	// X float64
	// Y float64
	V vec2.Vec2

	// Vx float64
	// Vy float64

	Hidden     bool
	Alpha      float64
	DrawOption func(*ebiten.DrawImageOptions)

	JustPressedTouchID ebiten.TouchID
	TouchID            ebiten.TouchID
	lastFrameTouchX    int
	lastFrameTouchY    int

	isJustReleased bool
	isStillTouched bool
}

func NewSprite(img *ebiten.Image, pos vec2.Vec2) *Sprite {
	spr := &Sprite{
		Img: img,
		Pos: pos,
		// X:     x,
		// Y:     y,
		Alpha:              1,
		DrawOption:         func(*ebiten.DrawImageOptions) {},
		TouchID:            5000,
		JustPressedTouchID: 5000,
	}

	return spr
}

func (s *Sprite) SetCenter(center int) {
	s.Pos.X = float64(center - s.Img.Bounds().Dx()/2)
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if !s.Hidden {
		op := &ebiten.DrawImageOptions{}
		s.DrawOption(op)

		op.GeoM.Translate(s.Pos.X, s.Pos.Y)
		op.ColorM.Scale(1, 1, 1, s.Alpha)
		screen.DrawImage(s.Img, op)
	}
}

func (s *Sprite) IsJustTouched() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	if len(justPressedTouchIDs) > 0 {
		x, y := ebiten.TouchPosition(s.JustPressedTouchID)
		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			// s.JustPressedTouchID = t
			return true
		}

	}

	return false
}

func (s *Sprite) checkIsTouchJustReleased() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			s.isStillTouched = true
			s.isJustReleased = true
			return
		}

		s.isJustReleased = true
		s.isStillTouched = false
		return
	}

	// if s.JustPressedTouchID != 5000 {
	if inpututil.IsTouchJustReleased(s.JustPressedTouchID) {
		s.JustPressedTouchID = 5000

		x, y := s.lastFrameTouchX, s.lastFrameTouchY
		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			s.isStillTouched = true
			s.isJustReleased = true
			return
		}

		s.isJustReleased = true
		s.isStillTouched = false
		return
	}

	s.isJustReleased = false
	s.isStillTouched = false
	return
}

func (s *Sprite) IsTouchJustReleased() (isJustReleased bool, isStillTouched bool) {
	/*
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()

			if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
				return true, true
			}

			return true, false

		}

		// if s.JustPressedTouchID != 5000 {
		if inpututil.IsTouchJustReleased(s.JustPressedTouchID) {
			s.JustPressedTouchID = 5000

			x, y := s.lastFrameTouchX, s.lastFrameTouchY
			if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
				return true, true
			}

			return true, false
		}

		return false, false
	*/

	return s.isJustReleased, s.isStillTouched
}

func (s *Sprite) IsTouched() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	if len(touchIDs) > 0 {
		x, y := ebiten.TouchPosition(s.TouchID)
		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	return false
}

func (s *Sprite) Update() {
	if len(touchIDs) > 0 && s.TouchID == 5000 {
		for _, t := range touchIDs {
			x, y := ebiten.TouchPosition(t)
			if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
				s.TouchID = t
				// break
			}
		}
	}

	if s.TouchID != 5000 && inpututil.IsTouchJustReleased(s.TouchID) {
		s.TouchID = 5000
	}

	if len(justPressedTouchIDs) > 0 {
		for _, t := range justPressedTouchIDs {
			x, y := ebiten.TouchPosition(t)
			if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
				s.JustPressedTouchID = t
				// break
			}
			// else {
			// 	s.JustPressedTouchID = 5000
			// }
		}
	}

	if s.JustPressedTouchID != 5000 && inpututil.IsTouchJustReleased(s.JustPressedTouchID) {
	}

	if !inpututil.IsTouchJustReleased(s.JustPressedTouchID) {
		s.lastFrameTouchX, s.lastFrameTouchY = ebiten.TouchPosition(s.JustPressedTouchID)
	}

	s.checkIsTouchJustReleased()
}
