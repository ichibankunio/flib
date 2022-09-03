package raycast

import (
	"fmt"
	"image/color"
	"math"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ichibankunio/flib/vec2"
)

//go:embed shaders/wall.kage
var shaderByte []byte

type Renderer struct {
	Cam *Camera
	Stk *Stick
	Wld *World

	screenWidth  float64
	screenHeight float64

	shader *ebiten.Shader

	floorTexture  *ebiten.Image
	spriteTexture *ebiten.Image
	wallTexture   *ebiten.Image
	texSize       int
}

func (r *Renderer) Init(screenWidth, screenHeight float64, wallTextures []*ebiten.Image, floorTextures []*ebiten.Image, spriteTextures []*ebiten.Image, texSize int) {
	r.Cam = &Camera{}
	r.Cam.Init(screenWidth, screenHeight)

	r.Stk = &Stick{}
	r.Stk.Init(screenWidth, screenHeight)

	r.Wld = &World{}
	r.Wld.Init(screenWidth, screenHeight)

	r.screenWidth = screenWidth
	r.screenHeight = screenHeight

	var err error
	r.shader, err = ebiten.NewShader(shaderByte)
	if err != nil {
		panic(err)
	}

	r.floorTexture = ebiten.NewImage(int(r.screenWidth), int(r.screenHeight))
	for i, t := range floorTextures {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%(int(screenWidth)/texSize))*texSize), float64((i/(int(screenWidth)/texSize))*texSize))

		r.floorTexture.DrawImage(t, op)
	}

	r.wallTexture = ebiten.NewImage(int(screenWidth), int(screenHeight))
	for i, t := range wallTextures {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%(int(screenWidth)/texSize))*texSize), float64((i/(int(screenHeight)/texSize))*texSize))
		r.wallTexture.DrawImage(t, op)
	}

	r.spriteTexture = ebiten.NewImage(int(screenWidth), int(screenHeight))
	for i, t := range spriteTextures {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%(int(screenWidth)/texSize))*texSize), float64((i/(int(screenHeight)/texSize))*texSize))
		r.spriteTexture.DrawImage(t, op)
	}

}

func (r *Renderer) SetShader(b []byte) error {
	var err error
	r.shader, err = ebiten.NewShader(b)
	if err != nil {
		return err
	}

	return nil
}

func (r *Renderer) SetLevel(level [][]float32, width, height int) {
	r.Wld.level = level
	r.Wld.width = width
	r.Wld.height = height
}

func (r *Renderer) GetScreenWidth() float64 {
	return r.screenWidth
}

func (r *Renderer) GetScreenHeight() float64 {
	return r.screenHeight
}

func (r *Renderer) Update() {
	r.updateCamera()
	r.Stk.update()
	r.calcSpriteRenderPos()
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.renderWall(screen)

	r.Wld.DrawTopView(screen)

	ebitenutil.DrawRect(screen, r.Cam.pos.X/2-2, r.Cam.pos.Y/2-2, 4, 4, color.RGBA{255, 0, 0, 255})

	ebitenutil.DrawLine(screen, r.Cam.pos.X/2, r.Cam.pos.Y/2, r.Cam.pos.X/2+r.Cam.dir.X*200, r.Cam.pos.Y/2+r.Cam.dir.Y*200, color.RGBA{255, 0, 0, 255})

	// s.fps.Draw(screen)
	// s.debug.Draw(screen)

	r.Stk.Draw(screen)
}

func (r *Renderer) renderWall(screen *ebiten.Image) {

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]interface{}{
		"ScreenSize": []float32{float32(r.screenWidth), float32(r.screenHeight)},
		"Pos":        []float32{float32(r.Cam.pos.X / float64(r.Wld.gridSize)), float32(r.Cam.pos.Y / float64(r.Wld.gridSize))},
		"Dir":        []float32{float32(r.Cam.dir.X), float32(r.Cam.dir.Y)},
		"Plane":      []float32{float32(r.Cam.plane.X), float32(r.Cam.plane.Y)},
		"TexSize":    float32(r.Wld.texSize) - 0.1,
		// "MapSize":    []float32{float32(len(r.Wld.level[0])), float32(len(r.Wld.level))},
		"WorldSize":   []float32{float32(r.Wld.width), float32(r.Wld.height)},
		"Level":       r.Wld.level[0],
		"FloorLevel":  r.Wld.level[1],
		"SpriteParam": r.Wld.spriteRenderParam,
	}

	op.Images[0] = r.wallTexture
	op.Images[1] = r.floorTexture
	op.Images[2] = r.spriteTexture
	screen.DrawRectShader(int(r.screenWidth), int(r.screenHeight), r.shader, op)

}

func (r *Renderer) calcSpriteRenderPos() {
	invDet := 1.0 / (r.Cam.plane.X*r.Cam.dir.Y - r.Cam.dir.X*r.Cam.plane.Y) // 1/(ad-bc)
	for i, pos := range r.Wld.spritePos {
		relPos := pos.Sub(r.Cam.pos).Scale(1.0 / float64(r.Wld.gridSize))
		transPos := vec2.New(r.Cam.dir.Y*relPos.X-r.Cam.dir.X*relPos.Y, -r.Cam.plane.Y*relPos.X+r.Cam.plane.X*relPos.Y).Scale(invDet)
		screenX := (r.screenWidth / 2) * (1.0 - transPos.X/transPos.Y)

		//calculate height of the sprite on screen
		spriteSize := vec2.New(math.Abs(r.screenHeight/transPos.Y), math.Abs(r.screenHeight/transPos.Y))
		// spriteHeight := math.Abs(SCREEN_HEIGHT / transPos.Y) //using 'transformY' instead of the real distance prevents fisheye
		// spriteWidth := math.Abs(SCREEN_HEIGHT / transPos.Y)

		//calculate lowest and highest pixel to fill in current stripe
		drawStart := vec2.New(-spriteSize.X/2+screenX, -spriteSize.Y/2+r.screenHeight/2)
		// drawEnd := vec2.New(spriteWidth/2+screenX, spriteHeight/2+SCREEN_HEIGHT/2)

		if transPos.Y > 0 {
			// s.wld.spriteRenderParam[5*i] = float32(relPos.SquaredLength()*math.Min(SCREEN_HEIGHT/SCREEN_WIDTH*3/4, SCREEN_WIDTH/SCREEN_HEIGHT*4/3))
			r.Wld.spriteRenderParam[6*i+1] = float32(relPos.SquaredLength())
			// r.Wld.spriteRenderParam[5*i] = float32(relPos.SquaredLength() *SCREEN_HEIGHT/SCREEN_HEIGHT*3/4)
		} else {
			r.Wld.spriteRenderParam[6*i+1] = float32(-1)
		}

		// fmt.Printf("%f, %f, %f, %f\n", drawStart, spriteSize, relPos.SquaredLength(), transPos.Y)

		r.Wld.spriteRenderParam[6*i+2] = float32(drawStart.X)
		r.Wld.spriteRenderParam[6*i+3] = float32(drawStart.Y)
		r.Wld.spriteRenderParam[6*i+4] = float32(spriteSize.X)
		r.Wld.spriteRenderParam[6*i+5] = float32(spriteSize.Y)
	}

	r.Wld.sortSpriteRenderParam()

	for i := 0; i < 18; i++ {
		fmt.Printf("%.2f,", r.Wld.spriteRenderParam[i])
	}
	println("")
}

func (w *World) sortSpriteRenderParam() {
	if len(w.spritePos) < 2 {
		return
	}

	for i := 0; i < len(w.spritePos); i++ {
		for j := 0; j < len(w.spritePos)-i; j++ {
			if w.spriteRenderParam[6*j] < 0 || w.spriteRenderParam[6*(j+1)] < 0 {
				return
			}

			fmt.Printf("%f, %f\n", w.spriteRenderParam[6*j+1], w.spriteRenderParam[6*(j+1)+1])
			if w.spriteRenderParam[6*j+1] > w.spriteRenderParam[6*(j+1)+1] {
				for k := 0; k < 6; k++ {
					tmp := w.spriteRenderParam[6*j+k]
					w.spriteRenderParam[6*j+k] = w.spriteRenderParam[6*(j+1)+k]
					w.spriteRenderParam[6*(j+1)+k] = tmp
				}
				println("sort performed")
			}
		}
	}
}

func (r *Renderer) castRay(dir, plane vec2.Vec2) float64 {
	cameraX := 0.0 //x-coordinate in camera space
	rayDir := dir.Add(plane.Scale(cameraX))
	rayPos := vec2.New(r.Cam.pos.X/float64(r.Wld.gridSize), r.Cam.pos.Y/float64(r.Wld.gridSize))
	mapPos := vec2.New(math.Floor(rayPos.X), math.Floor(rayPos.Y))
	deltaDist := vec2.New(math.Abs(1.0/rayDir.X), math.Abs(1.0/rayDir.Y))
	perpWallDist := 0.0
	unit := vec2.New(1, 1)
	sideDist := vec2.New(0, 0)
	if rayDir.X < 0 {
		unit.X = -1
		sideDist.X = (rayPos.X - mapPos.X) * deltaDist.X
	} else {
		unit.X = 1
		sideDist.X = (mapPos.X + 1.0 - rayPos.X) * deltaDist.X
	}

	if rayDir.Y < 0 {
		unit.Y = -1
		sideDist.Y = (rayPos.Y - mapPos.Y) * deltaDist.Y
	} else {
		unit.Y = 1
		sideDist.Y = (mapPos.Y + 1.0 - rayPos.Y) * deltaDist.Y
	}
	side := -1.0
	for i := 0; i < 20; i++ {
		//jump to next map square, OR in x-direction, OR in y-direction
		if sideDist.X < sideDist.Y {
			sideDist.X += deltaDist.X
			mapPos.X += unit.X
			side = 0.0
		} else {
			sideDist.Y += deltaDist.Y
			mapPos.Y += unit.Y
			side = 1.0
		}

		if r.Wld.level[0][int(mapPos.Y)*r.Wld.width+int(mapPos.X)] >= 1 {
			// hit = 1
			break
		}

		//Calculate distance of perpendicular ray (oblique distance will give fisheye effect!)
	}

	if side == 0 {
		perpWallDist = sideDist.X - deltaDist.X
	} else {
		perpWallDist = sideDist.Y - deltaDist.Y

	}

	return perpWallDist
}

func (r *Renderer) updateCamera() {
	rotateV := 0.02
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxisValue(0, 3) > 0.3 || r.Stk.input[1] == STICK_RIGHT {
		r.Cam.dir = vec2.New(math.Cos(rotateV)*r.Cam.dir.X-math.Sin(rotateV)*r.Cam.dir.Y, math.Sin(rotateV)*r.Cam.dir.X+math.Cos(rotateV)*r.Cam.dir.Y)

		r.Cam.plane = vec2.New(math.Cos(rotateV)*r.Cam.plane.X-math.Sin(rotateV)*r.Cam.plane.Y, math.Sin(rotateV)*r.Cam.plane.X+math.Cos(rotateV)*r.Cam.plane.Y)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxisValue(0, 3) < -0.3 || r.Stk.input[1] == STICK_LEFT {
		r.Cam.dir = vec2.New(math.Cos(-rotateV)*r.Cam.dir.X-math.Sin(-rotateV)*r.Cam.dir.Y, math.Sin(-rotateV)*r.Cam.dir.X+math.Cos(-rotateV)*r.Cam.dir.Y)

		r.Cam.plane = vec2.New(math.Cos(-rotateV)*r.Cam.plane.X-math.Sin(-rotateV)*r.Cam.plane.Y, math.Sin(-rotateV)*r.Cam.plane.X+math.Cos(-rotateV)*r.Cam.plane.Y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxisValue(0, 1) < -0.1 || r.Stk.input[0] == STICK_UP {
		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.dir.Scale(r.Cam.v)))

		// r.Cam.pos = r.GetValidPos(r.Cam.por.X + r.Cam.dir.X*v, r.Cam.por.Y + r.Cam.dir.Y*v)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxisValue(0, 1) > 0.1 || r.Stk.input[0] == STICK_DOWN {

		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.dir.Scale(-r.Cam.v)))

	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.GamepadAxisValue(0, 0) > 0.1 || r.Stk.input[0] == STICK_RIGHT {

		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.plane.Scale(-r.Cam.v)))

	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.GamepadAxisValue(0, 0) < -0.1 || r.Stk.input[0] == STICK_LEFT {

		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.plane.Scale(r.Cam.v)))

	}

}

func (r *Renderer) collisionCheckedDelta(delta vec2.Vec2) vec2.Vec2 {
	if (delta.X > 0 && r.castRay(vec2.New(1, 0), r.Cam.plane) <= r.Cam.collisionDistance) || (delta.X < 0 && r.castRay(vec2.New(-1, 0), r.Cam.plane) <= r.Cam.collisionDistance) {
		delta.X = 0
	}
	if (delta.Y > 0 && r.castRay(vec2.New(0, 1), r.Cam.plane) <= r.Cam.collisionDistance) || (delta.Y < 0 && r.castRay(vec2.New(0, -1), r.Cam.plane) <= r.Cam.collisionDistance) {
		delta.Y = 0
	}

	return delta
}
