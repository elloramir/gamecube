// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// This file implements the functionality related to chunk
// generation and management in the Voxel Engine. A chunk is
// a cubic section of the game world that contains a number of
// individual voxels. Dividing the game world up into chunks allows
// for efficient loading, unloading, and rendering of the game world.
// This file contains code for generating new chunks as well as managing
// the data structures that represent them in memory.

package game

import (
	"github.com/go-gl/mathgl/mgl32"
	simplex "github.com/ojrac/opensimplex-go"
)

const (
	SizeWidth  = 16
	SizeHeight = 16
	SizeLength = 16
)

const (
	BlockEmpty = iota
	BlockGrass
	BlockVoid
	BlockWater
)

// Misc
const (
	NoiseSmooth = 20
	WaterHeight = 3
)

var Noise32 = simplex.New32(0)

type Chunk struct {
	X, Z    int32
	Data    [SizeWidth][SizeHeight][SizeLength]uint8
	Terrain *Mesh
	Water   *Mesh
}

func NewChunk(x, z int32) *Chunk {
	c := &Chunk{X: x, Z: z}
	c.generateTerrain()
	c.generateMesh()

	return c
}

func (c *Chunk) generateTerrain() {
	offsetX := int(c.X * SizeWidth)
	offsetZ := int(c.Z * SizeLength)

	for x := 0; x < SizeWidth; x++ {
		for z := 0; z < SizeLength; z++ {
			noiseX := float32(offsetX+x) / NoiseSmooth
			noiseY := float32(offsetZ+z) / NoiseSmooth

			// Normalize from [-1, 1] to [0, 1]
			value := (Noise32.Eval2(noiseX, noiseY) + 1) * 0.5
			height := int32(value * SizeHeight)

			// Grass
			for height >= 0 {
				c.Data[x][height][z] = BlockGrass
				height -= 1
			}

			// Water
			if c.Data[x][WaterHeight][z] == BlockEmpty {
				c.Data[x][WaterHeight][z] = BlockWater
			}
		}
	}
}

func (c *Chunk) GetBlock(x, y, z int32) uint8 {
	if y < 0 || y >= SizeLength {
		return BlockVoid
	}

	// TODO: Neighbour check
	if x < 0 || x >= SizeWidth || z < 0 || z >= SizeLength {
		return BlockEmpty
	}

	return c.Data[x][y][z]
}

func (c *Chunk) isVisible(i, j, k int32) bool {
	hot := c.GetBlock(i, j, k)

	return hot == BlockEmpty || hot == BlockWater
}

func (c *Chunk) generateMesh() {
	tVerts := Vertices{} // Terrain vertices
	wVerts := Vertices{} // Water vertices

	for k := int32(0); k < SizeLength; k++ {
		for j := int32(0); j < SizeHeight; j++ {
			for i := int32(0); i < SizeWidth; i++ {
				currentBlock := c.GetBlock(i, j, k)

				// Skip empty block
				if currentBlock == BlockEmpty {
					continue
				}

				// Enumerated vertices
				//   0-------1
				//  /       /|
				// 3-------2 |
				// | 4     | 5
				// |       |/
				// 7-------6

				x := float32(i)
				y := float32(j)
				z := float32(k)

				// The current block can be water, that said, we can
				// create the remaining vertices after checking it
				v4 := mgl32.Vec3{-0.5 + x, +0.5 + y, -0.5 + z}
				v5 := mgl32.Vec3{+0.5 + x, +0.5 + y, -0.5 + z}
				v6 := mgl32.Vec3{+0.5 + x, +0.5 + y, +0.5 + z}
				v7 := mgl32.Vec3{-0.5 + x, +0.5 + y, +0.5 + z}

				if currentBlock == BlockWater {
					wVerts.BakeQuad(v7, v6, v5, v4)
					continue
				}

				v0 := mgl32.Vec3{-0.5 + x, -0.5 + y, -0.5 + z}
				v1 := mgl32.Vec3{+0.5 + x, -0.5 + y, -0.5 + z}
				v2 := mgl32.Vec3{+0.5 + x, -0.5 + y, +0.5 + z}
				v3 := mgl32.Vec3{-0.5 + x, -0.5 + y, +0.5 + z}

				// The orientation order is very specific
				// going from negative to positive based on the face's normal
				if c.isVisible(i, j, k-1) {
					tVerts.BakeQuad(v1, v0, v4, v5)
				}
				if c.isVisible(i, j, k+1) {
					tVerts.BakeQuad(v3, v2, v6, v7)
				}
				if c.isVisible(i-1, j, k) {
					tVerts.BakeQuad(v0, v3, v7, v4)
				}
				if c.isVisible(i+1, j, k) {
					tVerts.BakeQuad(v2, v1, v5, v6)
				}
				if c.isVisible(i, j+1, k) {
					tVerts.BakeQuad(v7, v6, v5, v4)
				}
				if c.isVisible(i, j-1, k) {
					tVerts.BakeQuad(v0, v1, v2, v3)
				}
			}
		}
	}

	c.Terrain = tVerts.ToMesh()
	c.Water = wVerts.ToMesh()
}
