package flib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ichibankunio/flib/vec2"
)

var justPressedTouchIDs []ebiten.TouchID
var touchIDs []ebiten.TouchID
var justReleasedTouchIDs []ebiten.TouchID

func GetJustPressedTouchIDs() []ebiten.TouchID {
	return justPressedTouchIDs
}

func GetTouchIDs() []ebiten.TouchID {
	return touchIDs
}

func GetJustReleasedTouchIDs() []ebiten.TouchID {
	return justReleasedTouchIDs
}

func IsThereJustReleasedTouch(pos vec2.Vec2, size vec2.Vec2) bool {
	for i := 0; i < len(inpututil.AppendJustReleasedTouchIDs(nil)); i++ {
		x, y := inpututil.TouchPositionInPreviousTick(inpututil.AppendJustReleasedTouchIDs(nil)[i])
		if x > int(pos.X) && x < int(pos.X + size.X) && y > int(pos.Y) && y < int(pos.Y + size.Y) {
			return true

		}
	}
	
	return false
}