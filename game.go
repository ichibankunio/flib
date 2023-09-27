package flib

import (
	_ "embed"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ichibankunio/flib/storage"
)

type Game struct {
	lastFrameTime int
	thisFrameTime int
	deltaTime     float64

	Counter int

	Scenes []Scene

	State   SceneID
	Lang    LangID
	Storage storage.Storage

	IsSceneTransition         bool
	transitionTick            *TickF
	transitionDuration        int
	justPressedTouchBeganPosX int
	justPressedTouchBeganPosY int
}

type LangID int

const (
	LANG_JA LangID = iota
	LANG_EN
)

type SceneID int

func (g *Game) GetJustPressedTouchBeganPosX() int {
	return g.justPressedTouchBeganPosX
}

func (g *Game) GetJustPressedTouchBeganPosY() int {
	return g.justPressedTouchBeganPosY
}

func (g *Game) Update() error {
	justReleasedTouchIDs = justReleasedTouchIDs[:0]
	for _, id := range touchIDs {
		if inpututil.IsTouchJustReleased(id) {
			justReleasedTouchIDs = append(justReleasedTouchIDs, id)
		}
	}

	justPressedTouchIDs = inpututil.AppendJustPressedTouchIDs(justPressedTouchIDs[:0])
	if len(justPressedTouchIDs) > 0 {
		x, y := ebiten.TouchPosition(justPressedTouchIDs[0])
		g.justPressedTouchBeganPosX = x
		g.justPressedTouchBeganPosY = y
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.justPressedTouchBeganPosX = x
		g.justPressedTouchBeganPosY = y
	}


	// touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs[:0])
	touchIDs = ebiten.AppendTouchIDs(touchIDs[:0])

	g.thisFrameTime = time.Now().Nanosecond()

	g.deltaTime = float64(g.thisFrameTime-g.lastFrameTime) / math.Pow(10, float64(int(math.Log10(float64(g.thisFrameTime-g.lastFrameTime))+2)))
	if g.deltaTime < 0 || math.IsNaN(g.deltaTime) {
		g.deltaTime = 0
	}
	g.lastFrameTime = g.thisFrameTime

	g.Counter++

	if g.IsSceneTransition {
		g.transitionTick.Update(g)
	} else {
		for _, scene := range g.Scenes {
			if g.State == scene.GetID() {
				err := scene.Update(g)

				return err
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, scene := range g.Scenes {
		if g.State == scene.GetID() {
			scene.Draw(screen)
			break
		}
	}

	if g.IsSceneTransition {
		w, h := screen.Size()
		op := &ebiten.DrawRectShaderOptions{}
		op.Uniforms = map[string]interface{}{
			"Time": float32(g.transitionTick.contentCounter) / float32(g.transitionDuration),
			// "Cursor":     []float32{float32(cx), float32(cy)},
			"ScreenSize": []float32{float32(w), float32(h)},
		}
		screen.DrawRectShader(w, h, shaders[SHADER_FADEINOUT], op)
	}
}

func (g *Game) AddScene(scene interface{}) {
	scene.(Scene).Init(g)
	g.Scenes = append(g.Scenes, scene.(Scene))
}
