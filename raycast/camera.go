package raycast

import (
	// "image/color"
	// "math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/flib"
	"github.com/ichibankunio/flib/vec2"
)

type Camera struct {

	//--camera position, init to start position--//
	pos vec2.Vec2

	// vertical camera strafing up/down, for jumping/crouching
	posZ float64

	//--current facing direction, init to values coresponding to FOV--//
	dir vec2.Vec2

	//--the 2d raycaster version of camera plane, adjust y component to change FOV (ratio between this and dir x resizes FOV)--//
	plane vec2.Vec2

	collisionDistance float64

	v float64
}

func (c *Camera) Init(screenWidth, screenHeight float64) {
	c.pos = vec2.New(360, 360)
	c.dir = vec2.New(-1, 0)
	c.plane = vec2.New(0, 0.66*screenWidth/screenHeight*3/4)
	// c.plane = vec2.New(0, 0.66 * SCREEN_WIDTH / 960 * 720 / SCREEN_HEIGHT)
	c.collisionDistance = 0.2

	c.v = 2.0
}

func (c *Camera) Update(g *flib.Game) {

}

func (c *Camera) Draw(screen *ebiten.Image) {

}
