package flib

import "github.com/hajimehoshi/ebiten/v2"

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