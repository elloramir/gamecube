// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package world

type BlockKind uint8

const (
	BlockAir BlockKind = iota
	BlockVoid
	BlockGrass
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

func generateQuad(Mesh *mesh, side QuadSide, i, j, k int32) {
	xp := float32( 0.5 + x);
	xn := float32(-0.5 + x);
	yp := float32( 0.5 + y);
	yn := float32(-0.5 + y);
	zp := float32( 0.5 + z);
	zn := float32(-0.5 + z);

	switch side {
		case SideNorth:
			mesh.Vertex(xN, yP, zN, 0, 0, -1, 1, 1)
			mesh.Vertex(xP, yN, zN, 0, 0, -1, 0, 0)
			mesh.Vertex(xN, yN, zN, 0, 0, -1, 1, 0)
			mesh.Vertex(xP, yP, zN, 0, 0, -1, 0, 1)
			mesh.Quad(0, 1, 2, 0, 3, 1)
		case SideSouth:
			mesh.Vertex(xN, yP, zP, 0, 0, 1, 0, 1)
			mesh.Vertex(xN, yN, zP, 0, 0, 1, 0, 0)
			mesh.Vertex(xP, yN, zP, 0, 0, 1, 1, 0)
			mesh.Vertex(xP, yP, zP, 0, 0, 1, 1, 1)
			mesh.Quad(0, 1, 2, 0, 2, 3)
		case SideEast:
			mesh.Vertex(xP, yP, zN, 1, 0, 0, 1, 1)
			mesh.Vertex(xP, yN, zP, 1, 0, 0, 0, 0)
			mesh.Vertex(xP, yN, zN, 1, 0, 0, 1, 0)
			mesh.Vertex(xP, yP, zP, 1, 0, 0, 0, 1)
			mesh.Quad(0, 1, 2, 3, 1, 0)
		case SideWest:
			mesh.Vertex(xN, yP, zN, -1, 0, 0, 0, 1)
			mesh.Vertex(xN, yN, zN, -1, 0, 0, 0, 0)
			mesh.Vertex(xN, yN, zP, -1, 0, 0, 1, 0)
			mesh.Vertex(xN, yP, zP, -1, 0, 0, 1, 1)
			mesh.Quad(0, 1, 2, 3, 0, 2)
		case SideTop:
			mesh.Vertex(xN, yP, zP, 0, 1, 0, 0, 0)
			mesh.Vertex(xP, yP, zP, 0, 1, 0, 1, 0)
			mesh.Vertex(xN, yP, zN, 0, 1, 0, 0, 1)
			mesh.Vertex(xP, yP, zN, 0, 1, 0, 1, 1)
			mesh.Quad(0, 1, 2, 2, 1, 3)
		case SideBottom:
			mesh.Vertex(xN, yN, zP, 0, -1, 0, 0, 1)
			mesh.Vertex(xN, yN, zN, 0, -1, 0, 0, 0)
			mesh.Vertex(xP, yN, zP, 0, -1, 0, 1, 1)
			mesh.Vertex(xP, yN, zN, 0, -1, 0, 1, 0)
			mesh.Quad(0, 1, 2, 1, 3, 2)
	}
}
