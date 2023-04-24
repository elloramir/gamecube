// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package world

import "github.com/elloramir/gamecube/gfx"

type BlockKind uint8

const (
	BlockEmpty BlockKind = iota
	BlockVoid
	BlockGrass
	BlockWater
)

type QuadSide uint32

const (
	SideNorth QuadSide = iota
	SideSouth
	SideEast
	SideWest
	SideTop
	SideBottom
)

func generateQuad(mesh *gfx.Mesh, side QuadSide, i, j, k int32) {
	xp := 0.5 + float32(i)
	xn := -0.5 + float32(i)
	yp := 0.5 + float32(j)
	yn := -0.5 + float32(j)
	zp := 0.5 + float32(k)
	zn := -0.5 + float32(k)

	switch side {
	case SideNorth:
		mesh.Vertex(xn, yp, zn, 0, 0, -1, 1, 1)
		mesh.Vertex(xp, yn, zn, 0, 0, -1, 0, 0)
		mesh.Vertex(xn, yn, zn, 0, 0, -1, 1, 0)
		mesh.Vertex(xp, yp, zn, 0, 0, -1, 0, 1)
		mesh.Quad(0, 1, 2, 0, 3, 1)
	case SideSouth:
		mesh.Vertex(xn, yp, zp, 0, 0, 1, 0, 1)
		mesh.Vertex(xn, yn, zp, 0, 0, 1, 0, 0)
		mesh.Vertex(xp, yn, zp, 0, 0, 1, 1, 0)
		mesh.Vertex(xp, yp, zp, 0, 0, 1, 1, 1)
		mesh.Quad(0, 1, 2, 0, 2, 3)
	case SideEast:
		mesh.Vertex(xp, yp, zn, 1, 0, 0, 1, 1)
		mesh.Vertex(xp, yn, zp, 1, 0, 0, 0, 0)
		mesh.Vertex(xp, yn, zn, 1, 0, 0, 1, 0)
		mesh.Vertex(xp, yp, zp, 1, 0, 0, 0, 1)
		mesh.Quad(0, 1, 2, 3, 1, 0)
	case SideWest:
		mesh.Vertex(xn, yp, zn, -1, 0, 0, 0, 1)
		mesh.Vertex(xn, yn, zn, -1, 0, 0, 0, 0)
		mesh.Vertex(xn, yn, zp, -1, 0, 0, 1, 0)
		mesh.Vertex(xn, yp, zp, -1, 0, 0, 1, 1)
		mesh.Quad(0, 1, 2, 3, 0, 2)
	case SideTop:
		mesh.Vertex(xn, yp, zp, 0, 1, 0, 0, 0)
		mesh.Vertex(xp, yp, zp, 0, 1, 0, 1, 0)
		mesh.Vertex(xn, yp, zn, 0, 1, 0, 0, 1)
		mesh.Vertex(xp, yp, zn, 0, 1, 0, 1, 1)
		mesh.Quad(0, 1, 2, 2, 1, 3)
	case SideBottom:
		mesh.Vertex(xn, yn, zp, 0, -1, 0, 0, 1)
		mesh.Vertex(xn, yn, zn, 0, -1, 0, 0, 0)
		mesh.Vertex(xp, yn, zp, 0, -1, 0, 1, 1)
		mesh.Vertex(xp, yn, zn, 0, -1, 0, 1, 0)
		mesh.Quad(0, 1, 2, 1, 3, 2)
	}
}
