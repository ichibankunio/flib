package flib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

func IsThereJustReleasedTouch(x0, y0, x1, y1 int) bool {
	for i := 0; i < len(justReleasedTouchIDs); i++ {
		x, y := inpututil.TouchPositionInPreviousTick(justReleasedTouchIDs[i])
		if x > x0 && x < x1 && y > y0 && y < y1 {
			return true
		}
	}
	
	return false
}