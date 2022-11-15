package raycast

import (
	"math"

	"github.com/ichibankunio/flib/vec2"
)

func (r *Renderer) fragment(position vec2.Vec2) []byte {
	cameraX := 2.0*(1.0-position.X/r.screenWidth) - 1.0
	rayDir := r.Cam.dir.Add(r.Cam.plane.Scale(cameraX))
	pos := r.Cam.pos.Scale(1 / float64(r.texSize))
	mapPos := pos.Floor()
	deltaDist := vec2.New(1/rayDir.X, 1/rayDir.Y).Abs()
	perpWallDist := 0.0

	unit := rayDir.Sign()
	side := -1.0
	sideDist := vec2.New(0, 0)
	if rayDir.X < 0 {
		unit.X = -1
		sideDist.X = (pos.X - mapPos.X) * deltaDist.X
	} else {
		unit.X = 1
		sideDist.X = (mapPos.X + 1.0 - pos.X) * deltaDist.X
	}

	if rayDir.Y < 0 {
		unit.Y = -1
		sideDist.Y = (pos.Y - mapPos.Y) * deltaDist.Y
	} else {
		unit.Y = 1
		sideDist.Y = (mapPos.Y + 1.0 - pos.Y) * deltaDist.Y
	}

	mapIndex := 0

	for i := 0; i < 20; i++ {
		if sideDist.X < sideDist.Y {
			sideDist.X += deltaDist.X
			mapPos.X += unit.X
			side = 0.0
		} else {
			sideDist.Y += deltaDist.Y
			mapPos.Y += unit.Y
			side = 1.0
		}

		// fmt.Printf("%f\n", mapPos)
		mapIndex = r.getWallMap(int(mapPos.X), int(mapPos.Y))
		if mapIndex >= 1 {
			break
		}
	}

	wallX := 0.0
	if side == 0 {
		perpWallDist = sideDist.X - deltaDist.X
		wallX = pos.Y + perpWallDist*rayDir.Y //rayposY + perpwallDist * sin(angle)みたいなイメージ
	} else {
		perpWallDist = sideDist.Y - deltaDist.Y
		wallX = pos.X + perpWallDist*rayDir.X //rayposX + perpwallDist * cos(angle)みたいなイメージ
	}

	lineHeight := r.screenHeight / perpWallDist
	drawStart := -lineHeight/2.0 + r.screenHeight/2.0
	drawEnd := lineHeight/2.0 + r.screenHeight/2.0

	// fmt.Printf("%f\n", perpWallDist)

	if position.Y >= drawStart && position.Y <= drawEnd {
		texX := int((wallX + float64(int(mapIndex)%(int(r.screenWidth))/r.texSize)) * float64(r.texSize))
		texY := int((position.Y - drawStart) / lineHeight * float64(r.texSize))

		idx := texX + texY*int(r.screenWidth)

		return r.textureBytes[0][idx : idx+3]
		// return []byte{byte(255), 0, 0, byte(255)}

	}

	return []byte{255, 0, 0, 255}
}

func (r *Renderer) getWallMap(x, y int) int {
	return int(r.Wld.level[0][10*y+x])
}

func (r *Renderer) getFloorMap(x, y int) int {
	return int(r.Wld.level[1][10*y+x])
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
