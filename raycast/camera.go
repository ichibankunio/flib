package raycast

import (
	// "image/color"
	// "math"

	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/flib"
	"github.com/ichibankunio/flib/vec2"
)

type EngineState int

const (
	ENGINE_NONE = iota
	ENGINE_ACCEL
	ENGINE_BRAKE
)

type camera struct {
	width   int
	height  int
	angle   float64
	halfFov float64

	angleV float64

	//--camera position, init to start position--//
	pos vec2.Vec2

	// vertical camera strafing up/down, for jumping/crouching
	posZ float64

	//--current facing direction, init to values coresponding to FOV--//
	dir vec2.Vec2

	//--the 2d raycaster version of camera plane, adjust y component to change FOV (ratio between this and dir x resizes FOV)--//
	plane vec2.Vec2

	collisionDistance float64

	image     *ebiten.Image
	viewImage *ebiten.Image
	v         float64
	a         float64
	maxV      float64

	engine EngineState

	far  float64
	near float64

	offScreen *ebiten.Image
}

func (c *camera) Init(screenWidth, screenHeight float64) {
	c.width = 720
	c.height = 720
	c.pos = vec2.New(360, 360)
	c.dir = vec2.New(-1, 0)
	c.plane = vec2.New(0, 0.66*screenWidth/screenHeight*3/4)
	// c.plane = vec2.New(0, 0.66 * SCREEN_WIDTH / 960 * 720 / SCREEN_HEIGHT)
	c.collisionDistance = 0.2

	c.angle = math.Pi
	c.angleV = 0.5 / 180 * math.Pi
	c.halfFov = 33 * math.Pi / 180

	c.image = ebiten.NewImage(c.width, c.height)
	c.viewImage = ebiten.NewImage(c.width, c.height)

	c.v = 2.0
	c.a = 0.02
	c.maxV = 1.4

	c.engine = ENGINE_NONE

	c.far = 20
	c.near = 3
}

func (c *camera) Update(g *flib.Game) {

}

func (c *camera) Draw(screen *ebiten.Image) {

}

