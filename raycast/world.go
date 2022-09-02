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

	spritePos         []vec2.Vec2
	spriteRenderParam []float32
}

func (w *World) Init(screenWidth, screenHeight float64) {
	w.gridSize = 64
	w.width = 10
	w.height = 10

	// w.spritePos = []vec2.Vec2{
	// 	{128, 128},
	// }

	w.spriteRenderParam = make([]float32, 60)
	for i := range w.spriteRenderParam {
		if i%6 == 0 {
			w.spriteRenderParam[i] = -1
		}
	}

	w.texSize = 64

	// w.floorTexture = ebiten.NewImage(int(screenWidth), int(screenHeight))
	// op := &ebiten.DrawImageOptions{}
	// w.floorTexture.DrawImage(images[IMG_GROUND], op)

	// for i, t := range floorTextures {
	// 	op := &ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64((i%(int(screenWidth)/w.texSize))*w.texSize), float64((i/(int(screenWidth)/w.texSize))*w.texSize))

	// 	w.floorTexture.DrawImage(t, op)
	// }

	// w.wallTexture = ebiten.NewImage(int(screenWidth), int(screenHeight))
	// // w.wallTexture.Fill(color.RGBA{255, 0, 0, 255})
	// // op := &ebiten.DrawImageOptions{}
	// // w.wallTexture.DrawImage(images[IMG_CITY], op)
	// for i, t := range wallTextures {
	// 	op := &ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64((i%(int(screenWidth)/w.texSize))*w.texSize), float64((i/(int(screenWidth)/w.texSize))*w.texSize))
	// 	w.wallTexture.DrawImage(t, op)
	// }

	// w.spriteTexture = ebiten.NewImage(int(screenWidth), int(screenHeight))
	// // w.wallTexture.Fill(color.RGBA{255, 0, 0, 255})
	// // op := &ebiten.DrawImageOptions{}
	// w.spriteTexture.DrawImage(images[IMG_POLE], op)

	// ebitenutil.DrawLine(w.wallTexture, float64(w.texSize), 0, float64(w.texSize), float64(w.texSize), color.RGBA{0, 0, 255, 255})
	// op.GeoM.Translate(float64(images[IMG_WALL].Bounds().Dx()), 0)
	// w.wallTexture.DrawImage(images[IMG_WALL], op)

}

func (w *World) NewSprite(pos vec2.Vec2, texID int) {
	if len(w.spritePos) < 10 {
		w.spritePos = append(w.spritePos, pos)
		w.spriteRenderParam[6*(len(w.spritePos)-1)] = float32(texID)
		w.spriteRenderParam[6*(len(w.spritePos)-1)+1] = 0
		w.spriteRenderParam[6*(len(w.spritePos)-1)+2] = 0
		w.spriteRenderParam[6*(len(w.spritePos)-1)+3] = 0
		w.spriteRenderParam[6*(len(w.spritePos)-1)+4] = 0
		w.spriteRenderParam[6*(len(w.spritePos)-1)+5] = 0
		w.spriteRenderParam[6*(len(w.spritePos)-1)+6] = 0
	}
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
