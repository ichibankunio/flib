package raycast

import (
	"image/color"
	"math"

	_"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ichibankunio/flib/vec2"
)

//go:embed shaders/wall.kage
var shaderByte []byte

type Renderer struct {
	cam *camera
	stk *stick
	wld *world

	screenWidth float64
	screenHeight float64

	shader *ebiten.Shader

	floorTexture *ebiten.Image
	spriteTexture *ebiten.Image
	wallTexture *ebiten.Image
	
}

func (r *Renderer) Init(screenWidth, screenHeight float64, wallTexture, floorTexture, spriteTexture *ebiten.Image) {
	r.cam = &camera{}
	r.cam.Init(screenWidth, screenHeight)

	r.stk = &stick{}
	r.stk.Init(screenWidth, screenHeight)

	r.wld = &world{}
	r.wld.Init(screenWidth, screenHeight)

	r.screenWidth = screenWidth
	r.screenHeight = screenHeight

	var err error
	r.shader, err = ebiten.NewShader(shaderByte)
	if err != nil {
		panic(err)
	}

	r.wallTexture = wallTexture
	r.floorTexture = floorTexture
	r.spriteTexture = spriteTexture
}

func (r *Renderer) GetScreenWidth() float64 {
	return r.screenWidth
}

func (r *Renderer) GetScreenHeight() float64 {
	return r.screenHeight
}

func (r *Renderer) Update() {
	r.updateCamera()
	r.stk.update()
	r.calcSpriteRenderPos()
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.RenderWall(screen)

	r.wld.DrawTopView(screen)

	ebitenutil.DrawRect(screen, r.cam.pos.X/2-2, r.cam.pos.Y/2-2, 4, 4, color.RGBA{255, 0, 0, 255})

	ebitenutil.DrawLine(screen, r.cam.pos.X/2, r.cam.pos.Y/2, r.cam.pos.X/2+r.cam.dir.X*200, r.cam.pos.Y/2+r.cam.dir.Y*200, color.RGBA{255, 0, 0, 255})

	// s.fps.Draw(screen)
	// s.debug.Draw(screen)

	r.stk.Draw(screen)
}

func (r *Renderer) RenderWall(screen *ebiten.Image) {

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]interface{}{
		"ScreenSize": []float32{float32(r.screenWidth), float32(r.screenHeight)},
		"Pos":        []float32{float32(r.cam.pos.X / float64(r.wld.gridSize)), float32(r.cam.pos.Y / float64(r.wld.gridSize))},
		"Dir":        []float32{float32(r.cam.dir.X), float32(r.cam.dir.Y)},
		"Plane":      []float32{float32(r.cam.plane.X), float32(r.cam.plane.Y)},
		"TexSize":    float32(r.wld.texSize) - 0.1,
		// "MapSize":    []float32{float32(len(r.wld.level[0])), float32(len(r.wld.level))},
		"WorldSize":   []float32{float32(r.wld.width), float32(r.wld.height)},
		"Level":       r.wld.level,
		"FloorLevel":  r.wld.floorLevel,
		"SpriteParam": r.wld.spriteRenderParam,
	}

	op.Images[0] = r.wallTexture
	op.Images[1] = r.floorTexture
	op.Images[2] = r.spriteTexture
	screen.DrawRectShader(int(r.screenWidth), int(r.screenHeight), r.shader, op)

}

func (r *Renderer) calcSpriteRenderPos() {
	invDet := 1.0 / (r.cam.plane.X*r.cam.dir.Y - r.cam.dir.X*r.cam.plane.Y) // 1/(ad-bc)
	for i, pos := range r.wld.spritePos {
		relPos := pos.Sub(r.cam.pos).Scale(1.0 / float64(r.wld.gridSize))
		transPos := vec2.New(r.cam.dir.Y*relPos.X-r.cam.dir.X*relPos.Y, -r.cam.plane.Y*relPos.X+r.cam.plane.X*relPos.Y).Scale(invDet)
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
			r.wld.spriteRenderParam[5*i] = float32(relPos.SquaredLength())
			// r.wld.spriteRenderParam[5*i] = float32(relPos.SquaredLength() *SCREEN_HEIGHT/SCREEN_HEIGHT*3/4)
		} else {
			r.wld.spriteRenderParam[5*i] = float32(-1)
		}

		// fmt.Printf("%f, %f, %f, %f\n", drawStart, spriteSize, relPos.SquaredLength(), transPos.Y)

		r.wld.spriteRenderParam[5*i+1] = float32(drawStart.X)
		r.wld.spriteRenderParam[5*i+2] = float32(drawStart.Y)
		r.wld.spriteRenderParam[5*i+3] = float32(spriteSize.X)
		r.wld.spriteRenderParam[5*i+4] = float32(spriteSize.Y)
	}
}

func (r *Renderer) castRay(dir, plane vec2.Vec2) float64 {
	cameraX := 0.0 //x-coordinate in camera space
	rayDir := dir.Add(plane.Scale(cameraX))
	rayPos := vec2.New(r.cam.pos.X/float64(r.wld.gridSize), r.cam.pos.Y/float64(r.wld.gridSize))
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

		if r.wld.level[int(mapPos.Y)*r.wld.width+int(mapPos.X)] >= 1 {
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
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxisValue(0, 3) > 0.3 || r.stk.input[1] == STICK_RIGHT {
		r.cam.dir = vec2.New(math.Cos(rotateV)*r.cam.dir.X-math.Sin(rotateV)*r.cam.dir.Y, math.Sin(rotateV)*r.cam.dir.X+math.Cos(rotateV)*r.cam.dir.Y)

		r.cam.plane = vec2.New(math.Cos(rotateV)*r.cam.plane.X-math.Sin(rotateV)*r.cam.plane.Y, math.Sin(rotateV)*r.cam.plane.X+math.Cos(rotateV)*r.cam.plane.Y)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxisValue(0, 3) < -0.3 || r.stk.input[1] == STICK_LEFT {
		r.cam.dir = vec2.New(math.Cos(-rotateV)*r.cam.dir.X-math.Sin(-rotateV)*r.cam.dir.Y, math.Sin(-rotateV)*r.cam.dir.X+math.Cos(-rotateV)*r.cam.dir.Y)

		r.cam.plane = vec2.New(math.Cos(-rotateV)*r.cam.plane.X-math.Sin(-rotateV)*r.cam.plane.Y, math.Sin(-rotateV)*r.cam.plane.X+math.Cos(-rotateV)*r.cam.plane.Y)
	}

	// r.CastRay(r.cam.dir, r.cam.plane)

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxisValue(0, 1) < -0.1 || r.stk.input[0] == STICK_UP {
		r.cam.pos = r.cam.pos.Add(r.collisionCheckedDelta(r.cam.dir.Scale(r.cam.v)))

		// r.cam.pos = r.GetValidPos(r.cam.por.X + r.cam.dir.X*v, r.cam.por.Y + r.cam.dir.Y*v)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxisValue(0, 1) > 0.1 || r.stk.input[0] == STICK_DOWN {

		r.cam.pos = r.cam.pos.Add(r.collisionCheckedDelta(r.cam.dir.Scale(-r.cam.v)))

	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.GamepadAxisValue(0, 0) > 0.1 || r.stk.input[0] == STICK_RIGHT {

		r.cam.pos = r.cam.pos.Add(r.collisionCheckedDelta(r.cam.plane.Scale(-r.cam.v)))

	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.GamepadAxisValue(0, 0) < -0.1 || r.stk.input[0] == STICK_LEFT {

		r.cam.pos = r.cam.pos.Add(r.collisionCheckedDelta(r.cam.plane.Scale(r.cam.v)))

	}

}

func (r *Renderer) collisionCheckedDelta(delta vec2.Vec2) vec2.Vec2 {
	if (delta.X > 0 && r.castRay(vec2.New(1, 0), r.cam.plane) <= r.cam.collisionDistance) || (delta.X < 0 && r.castRay(vec2.New(-1, 0), r.cam.plane) <= r.cam.collisionDistance) {
		delta.X = 0
	}
	if (delta.Y > 0 && r.castRay(vec2.New(0, 1), r.cam.plane) <= r.cam.collisionDistance) || (delta.Y < 0 && r.castRay(vec2.New(0, -1), r.cam.plane) <= r.cam.collisionDistance) {
		delta.Y = 0
	}

	return delta
}
