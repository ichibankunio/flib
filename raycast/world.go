package raycast

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/flib/vec2"
)

type World struct {
	// level    [][]int
	// level          []float32
	level [][]float32
	// floorLevel     []float32
	gridSize       int
	topImage       *ebiten.Image
	texSize        int
	baseLightValue float32

	buffer *ebiten.Image

	// renderMap     [SCREEN_WIDTH]float32
	floorTexture  *ebiten.Image
	wallTexture   *ebiten.Image
	spriteTexture *ebiten.Image

	width  int
	height int

	// SpritePos         []vec2.Vec2
	SpriteRenderParam []float32

	Sprites []*Sprite
}

type Sprite struct {
	Pos vec2.Vec2
	ID int
	TexID int
}

func (w *World) Init(screenWidth, screenHeight float64) {
	w.gridSize = 64
	w.width = 10
	w.height = 10

	// w.spritePos = []vec2.Vec2{
	// 	{128, 128},
	// }

	w.SpriteRenderParam = make([]float32, 60)
	for i := range w.SpriteRenderParam {
		if i%6 == 0 {
			w.SpriteRenderParam[i] = -1
		}
	}

	w.texSize = 64

	

}

func (w *World) NewSprite(pos vec2.Vec2, texID int) {
	if len(w.Sprites) < 10 {
		w.Sprites = append(w.Sprites, &Sprite{
			Pos: pos,
			ID: len(w.Sprites),
			TexID: texID,
		})

		w.SpriteRenderParam[6*(len(w.Sprites)-1)] = float32(texID)
		w.SpriteRenderParam[6*(len(w.Sprites)-1)+1] = 0
		w.SpriteRenderParam[6*(len(w.Sprites)-1)+2] = 0
		w.SpriteRenderParam[6*(len(w.Sprites)-1)+3] = 0
		w.SpriteRenderParam[6*(len(w.Sprites)-1)+4] = 0
		w.SpriteRenderParam[6*(len(w.Sprites)-1)+5] = 0

	}

	// if len(w.SpritePos) < 10 {
	// 	w.SpritePos = append(w.SpritePos, pos)
	// 	w.SpriteRenderParam[6*(len(w.SpritePos)-1)] = float32(texID)
	// 	w.SpriteRenderParam[6*(len(w.SpritePos)-1)+1] = 0
	// 	w.SpriteRenderParam[6*(len(w.SpritePos)-1)+2] = 0
	// 	w.SpriteRenderParam[6*(len(w.SpritePos)-1)+3] = 0
	// 	w.SpriteRenderParam[6*(len(w.SpritePos)-1)+4] = 0
	// 	w.SpriteRenderParam[6*(len(w.SpritePos)-1)+5] = 0

	// }
}

func (w *World) NewTopView() {
	w.topImage = ebiten.NewImage(w.gridSize*w.width, w.gridSize*w.height)
	grid1 := ebiten.NewImage(w.gridSize-2, w.gridSize-2)
	grid1.Fill(color.RGBA{120, 120, 255, 120})
	grid2 := ebiten.NewImage(w.gridSize-2, w.gridSize-2)
	grid2.Fill(color.RGBA{120, 120, 120, 120})

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*w.gridSize+1), float64(y*w.gridSize+1))
			switch w.level[0][y*w.width+x] {
			case 0:
				w.topImage.DrawImage(grid2, op)
			case 1:
				w.topImage.DrawImage(grid1, op)
			}
		}
	}

}

func (w *World) DrawTopView(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(0, 0)
	screen.DrawImage(w.topImage, op)
}
