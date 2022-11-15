package raycast

import (
	"math"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/flib/vec2"
)

// //go:embed shaders/wall.kage
// var shaderByte []byte

type Renderer struct {
	Cam *Camera
	Stk *Stick
	Wld *World

	screenWidth  float64
	screenHeight float64

	shader *ebiten.Shader

	textures [4]*ebiten.Image

	floorTexture  *ebiten.Image
	spriteTexture *ebiten.Image
	wallTexture   *ebiten.Image
	texSize       int

	textureBytes [4][]byte

	pixels []byte
}

func (r *Renderer) Init(screenWidth, screenHeight float64, texSize int) {
	r.Cam = &Camera{}
	r.Cam.Init(screenWidth, screenHeight)

	r.Stk = &Stick{}
	r.Stk.Init(screenWidth, screenHeight)

	r.Wld = &World{}
	r.Wld.Init(screenWidth, screenHeight)

	r.texSize = texSize

	r.screenWidth = screenWidth
	r.screenHeight = screenHeight

	// var err error
	// r.shader, err = ebiten.NewShader(shaderByte)
	// if err != nil {
	// 	panic(err)
	// }

	// r.textures = textures

	r.pixels = make([]byte, 4*int(screenWidth*screenHeight))
	for i:=0; i < len(r.textureBytes); i++ {
		r.textureBytes[i] = make([]byte, 4*int(screenWidth*screenHeight))
	}
}

func (r *Renderer) SetTextures(textures [4]*ebiten.Image) {
	r.textures = textures
	for i := 0; i < len(textures); i++ {
		textures[i].ReadPixels(r.textureBytes[i])
	}
}

func (r *Renderer) NewTextureSheet(src []*ebiten.Image) *ebiten.Image {
	sheet := ebiten.NewImage(int(r.screenWidth), int(r.screenHeight))
	for i, s := range src {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%(int(r.screenWidth)/r.texSize))*r.texSize), float64((i/(int(r.screenWidth)/r.texSize))*r.texSize))

		sheet.DrawImage(s, op)
	}

	return sheet
}

// func (r *Renderer) SetShader(b []byte) error {
// 	var err error
// 	r.shader, err = ebiten.NewShader(b)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *Renderer) SetShader(s *ebiten.Shader) {
	r.shader = s
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

	for y := 0; y < int(r.screenHeight); y++ {
		for x := 0; x < int(r.screenWidth); x++ {
			c := r.fragment(vec2.New(float64(x)/r.screenWidth, float64(y)/r.screenHeight))
			idx := 4 * (y*int(r.screenWidth) + x)
			// println(idx)
			r.pixels[idx] = c[0]
			r.pixels[idx+1] = c[1]
			r.pixels[idx+2] = c[2]
			r.pixels[idx+3] = c[3]
			// r.pixels[idx] = 255
			// r.pixels[idx+1] = c[0]
			// r.pixels[idx+2] = 0
			// r.pixels[idx+3] = 255
		}
	}
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.renderWall(screen)

	// r.Wld.DrawTopView(screen)

	// ebitenutil.DrawRect(screen, r.Cam.pos.X/2-2, r.Cam.pos.Y/2-2, 4, 4, color.RGBA{255, 0, 0, 255})

	// ebitenutil.DrawLine(screen, r.Cam.pos.X/2, r.Cam.pos.Y/2, r.Cam.pos.X/2+r.Cam.dir.X*200, r.Cam.pos.Y/2+r.Cam.dir.Y*200, color.RGBA{255, 0, 0, 255})

	// s.fps.Draw(screen)
	// s.debug.Draw(screen)

	// r.Stk.Draw(screen)
}

func (r *Renderer) DrawTopView(screen *ebiten.Image) {
	r.Wld.DrawTopView(screen)
}

func (r *Renderer) renderWall(screen *ebiten.Image) {

	/*
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]interface{}{
		"ScreenSize":  []float32{float32(r.screenWidth), float32(r.screenHeight)},
		"Pos":         []float32{float32(r.Cam.pos.X / float64(r.Wld.gridSize)), float32(r.Cam.pos.Y / float64(r.Wld.gridSize))},
		"Dir":         []float32{float32(r.Cam.dir.X), float32(r.Cam.dir.Y)},
		"Plane":       []float32{float32(r.Cam.plane.X), float32(r.Cam.plane.Y)},
		"TexSize":     float32(r.Wld.texSize),
		"MapSize":     []float32{float32(len(r.Wld.level[0])), float32(len(r.Wld.level))},
		"WorldSize":   []float32{float32(r.Wld.width), float32(r.Wld.height)},
		"Level":       r.Wld.level[0],
		"FloorLevel":  r.Wld.level[1],
		"SpriteParam": r.Wld.SpriteRenderParam,
	}

	op.Images[0] = r.textures[0]
	op.Images[1] = r.textures[1]
	op.Images[2] = r.textures[2]
	op.Images[3] = r.textures[3]
	screen.DrawRectShader(int(r.screenWidth), int(r.screenHeight), r.shader, op)
	*/
	screen.WritePixels(r.pixels)


}

func (r *Renderer) calcSpriteRenderPos() {
	invDet := 1.0 / (r.Cam.plane.X*r.Cam.dir.Y - r.Cam.dir.X*r.Cam.plane.Y) // 1/(ad-bc)
	for i, spr := range r.Wld.Sprites {
		relPos := spr.Pos.Sub(r.Cam.pos).Scale(1.0 / float64(r.Wld.gridSize))
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
			r.Wld.SpriteRenderParam[6*i+1] = float32(relPos.SquaredLength())
			// r.Wld.spriteRenderParam[5*i] = float32(relPos.SquaredLength() *SCREEN_HEIGHT/SCREEN_HEIGHT*3/4)
		} else {
			r.Wld.SpriteRenderParam[6*i+1] = float32(-1)
		}

		r.Wld.SpriteRenderParam[6*i+2] = float32(drawStart.X)
		r.Wld.SpriteRenderParam[6*i+3] = float32(drawStart.Y)
		r.Wld.SpriteRenderParam[6*i+4] = float32(spriteSize.X)
		r.Wld.SpriteRenderParam[6*i+5] = float32(spriteSize.Y)
	}

	r.Wld.sortSpriteRenderParam()

	// for i := 0; i < 24; i++ {
	// 	fmt.Printf("%.2f,", r.Wld.SpriteRenderParam[i])
	// }
	// println("")
}

func (w *World) sortSpriteRenderParam() {
	if len(w.Sprites) < 2 {
		return
	}

	for i := 0; i < len(w.Sprites)-1; i++ {
		for j := 0; j < len(w.Sprites)-i-1; j++ {
			if w.SpriteRenderParam[6*j+1] > w.SpriteRenderParam[6*(j+1)+1] {
				for k := 0; k < 6; k++ {
					// fmt.Printf("%d, %d\n", 6*j+k, 6*(j+1)+k)

					// tmp := w.SpriteRenderParam[6*j+k]
					w.SpriteRenderParam[6*j+k], w.SpriteRenderParam[6*(j+1)+k] = w.SpriteRenderParam[6*(j+1)+k], w.SpriteRenderParam[6*j+k]
					// w.SpriteRenderParam[6*j+k] = w.SpriteRenderParam[6*(j+1)+k]
					// w.SpriteRenderParam[6*(j+1)+k] = tmp
				}
				w.Sprites[j], w.Sprites[j+1] = w.Sprites[j+1], w.Sprites[j]
			}
		}
	}
}

func (r *Renderer) updateCamera() {
	rotateV := 0.02
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxisValue(0, 3) > 0.3 || r.Stk.Input[1] == STICK_RIGHT {
		r.Cam.dir = vec2.New(math.Cos(rotateV)*r.Cam.dir.X-math.Sin(rotateV)*r.Cam.dir.Y, math.Sin(rotateV)*r.Cam.dir.X+math.Cos(rotateV)*r.Cam.dir.Y)

		r.Cam.plane = vec2.New(math.Cos(rotateV)*r.Cam.plane.X-math.Sin(rotateV)*r.Cam.plane.Y, math.Sin(rotateV)*r.Cam.plane.X+math.Cos(rotateV)*r.Cam.plane.Y)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxisValue(0, 3) < -0.3 || r.Stk.Input[1] == STICK_LEFT {
		r.Cam.dir = vec2.New(math.Cos(-rotateV)*r.Cam.dir.X-math.Sin(-rotateV)*r.Cam.dir.Y, math.Sin(-rotateV)*r.Cam.dir.X+math.Cos(-rotateV)*r.Cam.dir.Y)

		r.Cam.plane = vec2.New(math.Cos(-rotateV)*r.Cam.plane.X-math.Sin(-rotateV)*r.Cam.plane.Y, math.Sin(-rotateV)*r.Cam.plane.X+math.Cos(-rotateV)*r.Cam.plane.Y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.GamepadAxisValue(0, 1) < -0.1 || r.Stk.Input[0] == STICK_UP {
		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.dir.Scale(r.Cam.v)))

		// r.Cam.pos = r.GetValidPos(r.Cam.por.X + r.Cam.dir.X*v, r.Cam.por.Y + r.Cam.dir.Y*v)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxisValue(0, 1) > 0.1 || r.Stk.Input[0] == STICK_DOWN {

		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.dir.Scale(-r.Cam.v)))

	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.GamepadAxisValue(0, 0) > 0.1 || r.Stk.Input[0] == STICK_RIGHT {

		r.Cam.pos = r.Cam.pos.Add(r.collisionCheckedDelta(r.Cam.plane.Scale(-r.Cam.v)))

	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.GamepadAxisValue(0, 0) < -0.1 || r.Stk.Input[0] == STICK_LEFT {

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
