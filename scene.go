package flib

import (
	// "image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(*Game) error
	Draw(*ebiten.Image)
	Start()
	GetID() SceneID
	GetStatus() int
	Init()
}

/*
func ShiftSceneWithExpandingSprite(g *Game, spr Sprite, clr color.Color, ID SceneID) *Scene {
	return &Scene{
		Update: func(g *Game) error {
			return nil
		},
		Draw: func(screen *ebiten.Image) {

		},
		ID: ID,
	}
}

func ShiftSceneWithShrinkingSprite(spr Sprite, clr color.Color, ID SceneID) *Scene {
	return &Scene{
		Update: func(g *Game) error {
			return nil
		},
		Draw: func(screen *ebiten.Image) {

		},
		ID: ID,
	}
}
*/

func ShiftSceneWithFadeInOut(g *Game, shiftTo SceneID, duration int) {
	g.IsSceneTransition = true
	g.transitionDuration = duration
	g.transitionTick = NewTickF(1, false, -1, func(cc int, i []interface{}) bool {
		if cc == duration/2 {
			g.Scenes[shiftTo].Start()
			g.State = shiftTo

		} else if cc == duration {
			g.IsSceneTransition = false
			return true
		}

		return false
	})
}

func ShiftSceneWithCircle(g *Game, shiftTo SceneID) {
	g.IsSceneTransition = true
	g.transitionTick = NewTickF(1, false, -1, func(cc int, i []interface{}) bool {
		if cc == 60 {
			g.State = shiftTo
		} else if cc == 120 {
			g.IsSceneTransition = false
			return true
		}

		return false
	})
}
