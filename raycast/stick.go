package raycast

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ichibankunio/flib/vec2"
)

type Stick struct {
	// pos     vec2.Vec2
	visible  [2]bool
	img      *ebiten.Image
	input    [2]stickDirection
	touchIDs [2]ebiten.TouchID
	pos      [2]vec2.Vec2
	isMobile bool
	screenWidth float64
	screenHeight float64
}

type stickDirection int

const (
	STICK_NONE stickDirection = iota
	STICK_UP
	STICK_DOWN
	STICK_LEFT
	STICK_RIGHT
)

func (s *Stick) Init(screenWidth, screenHeight float64) {
	s.screenWidth = screenWidth
	s.screenHeight = screenHeight
	// s.pos = vec2.New(0, 0)
	s.pos = [2]vec2.Vec2{
		vec2.New(0, 0),
		vec2.New(0, 0),
	}

	s.visible = [2]bool{false, false}
	s.input = [2]stickDirection{
		STICK_NONE,
		STICK_NONE,
	}
	s.touchIDs = [2]ebiten.TouchID{-1, -1}
	s.isMobile = false

	s.img = ebiten.NewImage(int(s.screenHeight/10), int(s.screenHeight/10))
	ebitenutil.DrawCircle(s.img, float64(s.img.Bounds().Dx()/2), float64(s.img.Bounds().Dy()/2), float64(s.img.Bounds().Dx()/2), color.RGBA{200, 200, 200, 150})
	// s.img.Fill(color.RGBA{200, 200, 200, 50})
}

func (s *Stick) update() {
	if len(inpututil.AppendJustPressedTouchIDs(nil)) > 0 {
		for _, id := range inpututil.AppendJustPressedTouchIDs(nil) {
			x, y := ebiten.TouchPosition(id)
			if s.touchIDs[0] < 0 && x < int(s.screenWidth/2) {
				s.pos[0] = vec2.New(float64(x-s.img.Bounds().Dx()/2), float64(y-s.img.Bounds().Dy()/2))
				s.visible[0] = true
				s.touchIDs[0] = id
				s.isMobile = true
				continue
			}
			if s.touchIDs[1] < 0 && x >= int(s.screenWidth/2) {
				s.pos[1] = vec2.New(float64(x-s.img.Bounds().Dx()/2), float64(y-s.img.Bounds().Dy()/2))
				s.visible[1] = true
				s.touchIDs[1] = id
				s.isMobile = true
				continue
			}
		}
		// x, y := ebiten.TouchPosition(s.touchID)
		// if x < SCREEN_WIDTH/2 {
		// 	s.pos = vec2.New(float64(x-s.img.Bounds().Dx()/2), float64(y-s.img.Bounds().Dy()/2))
		// 	s.visible = true
		// }else {
		// 	s.touchID = -1
		// }
		// s.isMobile = true
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x < int(s.screenWidth/2) {
			s.pos[0] = vec2.New(float64(x-s.img.Bounds().Dx()/2), float64(y-s.img.Bounds().Dy()/2))
			s.visible[0] = true
		}
		if x >= int(s.screenWidth/2) {
			s.pos[1] = vec2.New(float64(x-s.img.Bounds().Dx()/2), float64(y-s.img.Bounds().Dy()/2))
			s.visible[1] = true
		}
	}

	if s.visible[0] {
		x, y := ebiten.CursorPosition()
		if s.isMobile {
			x, y = ebiten.TouchPosition(s.touchIDs[0])
		}
		current := vec2.New(float64(x), float64(y))
		rel := current.Sub(s.pos[0])
		if rel.X > 0 && math.Abs(rel.Y/rel.X) < 0.8 {
			s.input[0] = STICK_RIGHT
		} else if rel.X < 0 && math.Abs(rel.Y/rel.X) < 0.8 {
			s.input[0] = STICK_LEFT
		} else if rel.Y > 0 && math.Abs(rel.X/rel.Y) < 0.8 {
			s.input[0] = STICK_DOWN
		} else if rel.Y < 0 && math.Abs(rel.X/rel.Y) < 0.8 {
			s.input[0] = STICK_UP
		} else {
			s.input[0] = STICK_NONE
		}

		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			s.visible[0] = false
			s.input[0] = STICK_NONE
		}
		if inpututil.IsTouchJustReleased(s.touchIDs[0]) {
			s.visible[0] = false
			s.input[0] = STICK_NONE
			s.touchIDs[0] = -1
		}
	}

	if s.visible[1] {
		x, y := ebiten.CursorPosition()
		if s.isMobile {
			x, y = ebiten.TouchPosition(s.touchIDs[1])
		}
		current := vec2.New(float64(x), float64(y))
		rel := current.Sub(s.pos[1])
		if rel.X > 0 && math.Abs(rel.Y/rel.X) < 0.8 {
			s.input[1] = STICK_RIGHT
		} else if rel.X < 0 && math.Abs(rel.Y/rel.X) < 0.8 {
			s.input[1] = STICK_LEFT
		} else {
			s.input[1] = STICK_NONE
		}

		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			s.visible[1] = false
			s.input[1] = STICK_NONE

		}
		if inpututil.IsTouchJustReleased(s.touchIDs[1]) {
			s.visible[1] = false
			s.input[1] = STICK_NONE
			s.touchIDs[1] = -1
		}
	}

}

func (s *Stick) Draw(screen *ebiten.Image) {
	if s.visible[0] {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(s.pos[0].X, s.pos[0].Y)
		screen.DrawImage(s.img, op)
	}

	if s.visible[1] {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(s.pos[1].X, s.pos[1].Y)
		screen.DrawImage(s.img, op)
	}
}
