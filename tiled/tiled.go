package tiled

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ichibankunio/flib"
)

var Gravity = 0.3

type TiledRenderer struct {
	Map      TiledMap
	TileNumX int
	TileNumY int

	Source   *ebiten.Image

	Camera *TiledCamera

	Player *Character
}

type TiledCamera struct {
	Pos *flib.Vec2

	ScrollDeltaX int
	ScrollDeltaY int
	TileX int
	TileY int
}

type TiledMap struct {
	Layers   []TiledLayer `json:"layers"`
	Height   int          `json:"height"`
	Width    int          `json:"width"`
	TileSize int          `json:"tilewidth"`
}

type TiledLayer struct {
	Data []int `json:"data"`
	// ID int `json:"id"`
	Name    string `json:"name"`
	Visible bool   `json:"visible"`
}

func (tr *TiledRenderer) Init(x, y float64, source *ebiten.Image, tileNumX, tileNumY int) {
	tr.Camera = &TiledCamera{
		Pos: flib.NewVec(x, y),
		ScrollDeltaX: 0,
		ScrollDeltaY: 0,
		TileX: int(x/float64(tr.Map.TileSize)),
		TileY: int(y/float64(tr.Map.TileSize)),
	}

	tr.Source = source
	tr.TileNumX = tileNumX
	tr.TileNumY = tileNumY

	tr.Player = &Character{
		Rect: &Rectangle{
			Pos:  flib.NewVec(16*11, 16*10.5),
			Size: flib.NewVec(float64(tr.Map.TileSize), float64(tr.Map.TileSize)),
		},
		V:      flib.NewVec(0, 0),
		OldPos: flib.NewVec(16*11, 16*10.5),
	}
}

func (tr *TiledRenderer) GetTileCoordAtPos(pos *flib.Vec2) *flib.Vec2i {
	return flib.NewVeci(int((pos.X+float64(tr.Map.TileSize)*1.5) / float64(tr.Map.TileSize)), int((pos.Y + tr.Camera.Pos.Y+float64(tr.Map.TileSize)*1.5) / float64(tr.Map.TileSize)))
	// return flib.NewVeci(int((pos.X + tr.Camera.Pos.X+float64(tr.Map.TileSize)*1.5) / float64(tr.Map.TileSize)), int((pos.Y + tr.Camera.Pos.Y+float64(tr.Map.TileSize)*1.5) / float64(tr.Map.TileSize)))
}

func (tr *TiledRenderer) GetPosAtTileCoord(tileCoord *flib.Vec2i) *flib.Vec2 {
	return flib.NewVec(float64(tileCoord.X*tr.Map.TileSize)-tr.Camera.Pos.X, float64(tileCoord.Y*tr.Map.TileSize)-tr.Camera.Pos.Y)
}

func (tr *TiledRenderer) GetPosXAtTileCoordX(tileCoordX int) float64 {
	return float64(tileCoordX*tr.Map.TileSize)-float64(tr.Map.TileSize)*1.5
	// return float64(tileCoordX*tr.Map.TileSize)-tr.Camera.Pos.X-float64(tr.Map.TileSize)*1.5
}

func (tr *TiledRenderer) GetPosYAtTileCoordY(tileCoordY int) float64 {
	return float64(tileCoordY*tr.Map.TileSize)-float64(tr.Map.TileSize)*1.5
	// return float64(tileCoordY*tr.Map.TileSize)-tr.Camera.Pos.Y-float64(tr.Map.TileSize)*1.5
}

func (tr *TiledRenderer) GetTileIDAtTileCoord(tileCoord *flib.Vec2i) int {
	return tr.Map.Layers[0].Data[tileCoord.X+tileCoord.Y*tr.Map.Width]
}


func (tr *TiledRenderer) UpdateCamera() {
	if tr.Player.Rect.Pos.X - tr.Camera.Pos.X >= 16*11 {
		tr.Camera.Pos.X += tr.Player.Rect.Pos.X - tr.Player.OldPos.X
	}

	if tr.Camera.Pos.X < 0 {
		tr.Camera.Pos.X = 0
	}
	if tr.Camera.Pos.X > float64((tr.Map.Width-tr.TileNumX-1)*tr.Map.TileSize) {
		tr.Camera.Pos.X = float64((tr.Map.Width-tr.TileNumX-1)*tr.Map.TileSize)
	}
	if tr.Camera.Pos.Y < 0 {
		tr.Camera.Pos.Y = 0
	}
	if tr.Camera.Pos.Y > float64((tr.Map.Height-tr.TileNumY-1)*tr.Map.TileSize) {
		tr.Camera.Pos.Y = float64((tr.Map.Height-tr.TileNumY-1)*tr.Map.TileSize)
	}
	
	tr.Camera.TileX = (int(tr.Camera.Pos.X)+tr.Map.TileSize/2) / tr.Map.TileSize
	tr.Camera.TileY = (int(tr.Camera.Pos.Y)+tr.Map.TileSize/2) / tr.Map.TileSize

	tr.Camera.ScrollDeltaX = (int(tr.Camera.Pos.X)+tr.Map.TileSize/2) % tr.Map.TileSize
	tr.Camera.ScrollDeltaY = (int(tr.Camera.Pos.Y)+tr.Map.TileSize/2) % tr.Map.TileSize
}

func (tr *TiledRenderer) Update(g *flib.Game) {
	tr.Player.Update()//PlayerVを入力から決めて一旦Positionに加算する
	tr.CheckCollision()//衝突を検出して位置を修正、衝突してなければそのまま
	tr.UpdateCamera()//修正後のPositionからカメラをプレイヤーを追いかけるように移動させる
}

func (tr *TiledRenderer) CheckCollision() {
	if tr.Player.V.X > 0 { //right
		dstTileCoord := tr.GetTileCoordAtPos(tr.Player.Rect.Pos.Clone().Add(flib.NewVec(float64(tr.Map.TileSize-1), 0)))
		oldTileCoord := tr.GetTileCoordAtPos(tr.Player.Rect.Pos.Clone().Add(flib.NewVec(float64(tr.Map.TileSize-1), 0)).Sub(tr.Player.V))

		// println(dstTileCoord.X, oldTileCoord.X, tr.GetPosXAtTileCoordX(dstTileCoord.X-1))

		if dstTileCoord.X > oldTileCoord.X {
			for i := 0; i < 2; i++ {
				if tr.GetTileIDAtTileCoord(dstTileCoord.Clone().Add(flib.NewVeci(0, i))) != 0 {
					tr.Player.Rect.Pos.X = tr.GetPosXAtTileCoordX(dstTileCoord.X - 1)
					// println(tr.GetPosXAtTileCoordX(dstTileCoord.X - 1))
					fmt.Printf("RIGHT:HIT(%d,%d)\n", dstTileCoord.X, dstTileCoord.Y+i)
					// tr.Player.V.X = 0
					// if tr.Player.V.Y <= 0 {//上昇中なら上昇をストップ
					// 	tr.Player.V.Y = 0
					// 	// tr.Player.Rect.Pos.Y = tr.GetPosYAtTileCoordY(dstTileCoord.Y+i)
					// }
					break
				}
				// if int(tr.Player.Rect.Pos.Y)%tr.Map.TileSize == tr.Map.TileSize/2 {
				if (int(tr.Player.Rect.Pos.Y)+tr.Map.TileSize/2)%tr.Map.TileSize == 0 {
					break
				}
			}
		}
	}
	if tr.Player.V.X < 0 { //left
		dstTileCoord := tr.GetTileCoordAtPos(tr.Player.Rect.Pos.Clone())
		oldTileCoord := tr.GetTileCoordAtPos(tr.Player.Rect.Pos.Clone().Sub(tr.Player.V))

		if dstTileCoord.X < oldTileCoord.X {
			for i := 0; i < 2; i++ {
				if tr.GetTileIDAtTileCoord(dstTileCoord.Clone().Add(flib.NewVeci(0, i))) != 0 {
					tr.Player.Rect.Pos.X = tr.GetPosXAtTileCoordX(dstTileCoord.X) + float64(tr.Map.TileSize)
					fmt.Printf("LEFT:HIT(%d,%d)\n", dstTileCoord.X, dstTileCoord.Y+i)
					// tr.Player.V.X = 0

					// if tr.Player.V.Y <= 0 {//上昇中なら上昇をストップ
					// 	tr.Player.V.Y = 0
					// 	// tr.Player.Rect.Pos.Y = tr.GetPosYAtTileCoordY(dstTileCoord.Y+i)

					// }
					break
				}
				if (int(tr.Player.Rect.Pos.Y)+tr.Map.TileSize/2)%tr.Map.TileSize == 0 {

				// if int(tr.Player.Rect.Pos.Y)%tr.Map.TileSize == tr.Map.TileSize/2 {
					break
				}
			}
		}
	}

	if tr.Player.V.Y > 0 { //down
		dstTileCoord := tr.GetTileCoordAtPos(tr.Player.Rect.Pos.Clone().Add(flib.NewVec(0, float64(tr.Map.TileSize))))

		for i := 0; i < 2; i++ {
			if tr.GetTileIDAtTileCoord(dstTileCoord.Clone().Add(flib.NewVeci(i, 0))) != 0 {
				tr.Player.Rect.Pos.Y = tr.GetPosYAtTileCoordY(dstTileCoord.Y - 1)
				// fmt.Printf("DOWN:HIT(%d,%d)\n", dstTileCoord.X+i, dstTileCoord.Y)

				tr.Player.V.Y = 0
				break
			}

			if (int(tr.Player.Rect.Pos.X)+tr.Map.TileSize/2)%tr.Map.TileSize == 0 {
				// println("break")
				break
			}
		}

	}
	if tr.Player.V.Y < 0 { //up
		dstTileCoord := tr.GetTileCoordAtPos(tr.Player.Rect.Pos.Clone())
		
		for i := 0; i < 2; i++ {
			if tr.GetTileIDAtTileCoord(dstTileCoord.Clone().Add(flib.NewVeci(i, 0))) != 0 {
				tr.Player.Rect.Pos.Y = tr.GetPosYAtTileCoordY(dstTileCoord.Y) + float64(tr.Map.TileSize)
				fmt.Printf("UP:HIT(%d,%d)\n", dstTileCoord.X+i, dstTileCoord.Y)
				tr.Player.V.Y = 0

				break
			}
			println((int(tr.Player.Rect.Pos.X)+tr.Map.TileSize/2))

			if (int(tr.Player.Rect.Pos.X)+tr.Map.TileSize/2)%tr.Map.TileSize == 0 {

				println("break")
				break
			}
		}
	}
}

func (tr *TiledRenderer) Draw(screen *ebiten.Image) {
	for _, layer := range tr.Map.Layers {
		for i := 0; i < tr.TileNumY; i++ {
			for j := 0; j < tr.TileNumX; j++ {
				idx := layer.Data[(j+tr.Camera.TileX)+tr.Map.Width*(i+tr.Camera.TileY)]
				if idx == 0 {
					continue
				}

				op := &ebiten.DrawImageOptions{}


				op.GeoM.Translate(float64(j*tr.Map.TileSize-tr.Map.TileSize*2/2-tr.Camera.ScrollDeltaX), float64(i*tr.Map.TileSize-tr.Map.TileSize*2/2-tr.Camera.ScrollDeltaY))

				cropX := ((idx - 1) % (tr.Source.Bounds().Dx() / tr.Map.TileSize)) * tr.Map.TileSize
				cropY := ((idx - 1) / (tr.Source.Bounds().Dx() / tr.Map.TileSize)) * tr.Map.TileSize

				// cropY := ((gameTileMap.Layers[0].Data[(i+f.cameraY)*GAME_WIDTH/gameTileMap.TileSize+(j+f.cameraX)]-1) / (f.tileSource.Bounds().Dy() / gameTileMap.TileSize)) * gameTileMap.TileSize

				screen.DrawImage(tr.Source.SubImage(image.Rect(cropX, cropY, cropX+tr.Map.TileSize, cropY+tr.Map.TileSize)).(*ebiten.Image), op)

				if j == 2 {
					ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", i), j * tr.Map.TileSize - tr.Map.TileSize * 2 / 2 - tr.Camera.ScrollDeltaX, i * tr.Map.TileSize - tr.Map.TileSize - tr.Camera.ScrollDeltaY)
				}
				if i == tr.TileNumY - 1 {
					ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", (j+tr.Camera.TileX)), j * tr.Map.TileSize - tr.Map.TileSize * 2 / 2 - tr.Camera.ScrollDeltaX, i * tr.Map.TileSize - tr.Map.TileSize - tr.Camera.ScrollDeltaY)
					
				}
			}
		}
	}
}
