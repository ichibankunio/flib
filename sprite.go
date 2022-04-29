package flib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Sprite struct {
	Img *ebiten.Image
	Pos Vec2
	// X float64
	// Y float64
	V Vec2

	// Vx float64
	// Vy float64

	Hidden     bool
	Alpha      float64
	DrawOption func(*ebiten.DrawImageOptions)

	JustPressedTouchID ebiten.TouchID
	TouchID            ebiten.TouchID
	lastFrameTouchX    int
	lastFrameTouchY    int
}

func NewSprite(img *ebiten.Image, pos *Vec2) *Sprite {
	spr := &Sprite{
		Img: img,
		Pos: *pos,
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
	s.Pos.X = float64(center-s.Img.Bounds().Dx()/2)
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

	if len(JustPressedTouchIDs) > 0 {
		x, y := ebiten.TouchPosition(s.JustPressedTouchID)
		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			// s.JustPressedTouchID = t
			return true
		}

	}

	return false
}

func (s *Sprite) IsTouchJustReleased() (isJustReleased bool, isStillTouched bool) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true, true
		}
		return true, false

	}

	if inpututil.IsTouchJustReleased(s.JustPressedTouchID) {
		x, y := s.lastFrameTouchX, s.lastFrameTouchY
		DebugInt = x
		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true, true
		}

		return true, false
	}

	return false, false
}

func (s *Sprite) IsTouched() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	if len(TouchIDs) > 0 {
		x, y := ebiten.TouchPosition(s.TouchID)
		if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	return false
}

func (s *Sprite) Update() {
	if len(TouchIDs) > 0 {
		for _, t := range TouchIDs {
			x, y := ebiten.TouchPosition(t)
			if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
				s.TouchID = t
				// break

			} else {
				s.TouchID = 5000
			}
		}
	}

	if len(JustPressedTouchIDs) > 0 {
		for _, t := range JustPressedTouchIDs {
			x, y := ebiten.TouchPosition(t)
			if x >= int(s.Pos.X) && x <= int(s.Pos.X)+s.Img.Bounds().Dx() && y >= int(s.Pos.Y) && y <= int(s.Pos.Y)+s.Img.Bounds().Dy() {
				s.JustPressedTouchID = t
				// break
			} else {
				s.JustPressedTouchID = 5000
			}
		}
	}

	if !inpututil.IsTouchJustReleased(s.JustPressedTouchID) {
		s.lastFrameTouchX, s.lastFrameTouchY = ebiten.TouchPosition(s.JustPressedTouchID)
	}

}
