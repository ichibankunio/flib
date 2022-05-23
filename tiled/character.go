package tiled

import (

	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ichibankunio/flib"
)


type Character struct {
	// Pos *flib.Vec2
	V *flib.Vec2
	// Size *flib.Vec2
	Rect *Rectangle
	// OldPos flib.Vec2
	OldPos *flib.Vec2
	
}

func (c *Character) Update() {
	// fmt.Printf("%v\n", c.V)

	c.UpdateV()
	
	c.OldPos.X = c.Rect.Pos.X
	c.Rect.Pos.Add(c.V)
	// fmt.Printf("%v\n %v\n",c.OldPos, *c.Rect.Pos)
}

func (c *Character) UpdateV() {
	dL := inpututil.KeyPressDuration(ebiten.KeyLeft)
	dR := inpututil.KeyPressDuration(ebiten.KeyRight)
	dU := inpututil.KeyPressDuration(ebiten.KeyUp)
	dD := inpututil.KeyPressDuration(ebiten.KeyDown)
	duration := []int{dL, dR, dU, dD}
	sort.Slice(duration, func(i, j int) bool { return duration[i] < duration[j] }) //小さい順
	latestPressedKeyIndex := 0
	for i := 0; i < len(duration); i++ {
		if duration[i] > 0 {
			latestPressedKeyIndex = i
			break
		}
	}

	if dL > 0 && dL == duration[latestPressedKeyIndex] {
		c.V.X = -1

	} else if dR > 0 && dR == duration[latestPressedKeyIndex] {
		c.V.X = 1

	} else if dU > 0 && dU == duration[latestPressedKeyIndex] {
		// tr.Player.V.Y = -1

	} else if dD > 0 && dD == duration[latestPressedKeyIndex] {
		// tr.Player.V.Y = 1

	} else {
		c.V.X = 0
		// tr.Player.V.Y = 0

	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && c.V.Y == 0 {
		c.V.Y = -6
	}

	c.V.Y += Gravity
}



func (c *Character) IsIntersect(other *Rectangle) {

}

