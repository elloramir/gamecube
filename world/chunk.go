// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package world

import (
	"github.com/elloramir/gamecube/gfx"
	simplex "github.com/ojrac/opensimplex-go"
)

const (
	SizeWidth  = 16
	SizeHeight = 16
	SizeLength = 16
)

// Misc
const (
	NoiseSmooth = 20
	WaterHeight = 3
)

type Chunk struct {
	X, Z    int32
	Data    [SizeWidth][SizeHeight][SizeLength]BlockKind
	Terrain gfx.Mesh
	Water   gfx.Mesh
}

var Noise32 = simplex.New32(0)

func NewChunk(x, z int32) *Chunk {
	c := &Chunk{X: x, Z: z}
	c.generateTerrain()
	c.generateMesh()

	return c
}

func (c *Chunk) generateTerrain() {
	offsetX := int(c.X * SizeWidth)
	offsetZ := int(c.Z * SizeLength)

	for i := 0; i < SizeWidth; i++ {
		for k := 0; k < SizeLength; k++ {
			noiseX := float32(offsetX+i) / NoiseSmooth
			noiseY := float32(offsetZ+k) / NoiseSmooth

			// Normalize from [-1, 1] to [0, 1]
			value := (Noise32.Eval2(noiseX, noiseY) + 1) * 0.5
			height := int32(value * SizeHeight)

			// Grass
			for height >= 0 {
				c.Data[i][height][k] = BlockGrass
				height -= 1
			}

			// Water
			if c.Data[i][WaterHeight][k] == BlockEmpty {
				c.Data[i][WaterHeight][k] = BlockWater
			}
		}
	}
}

func (c *Chunk) GetBlock(x, y, z int32) BlockKind {
	if y < 0 || y >= SizeLength {
		return BlockVoid
	}

	// TODO: Neighbour check
	if x < 0 || x >= SizeWidth || z < 0 || z >= SizeLength {
		return BlockVoid
	}

	return c.Data[x][y][z]
}

func (c *Chunk) isTransparent(i, j, k int32) bool {
	hot := c.GetBlock(i, j, k)

	return hot == BlockEmpty || hot == BlockWater
}

func (c *Chunk) generateMesh() {
	c.Terrain.Unload()
	c.Water.Unload()

	for k := int32(0); k < SizeLength; k++ {
		for j := int32(0); j < SizeHeight; j++ {
			for i := int32(0); i < SizeWidth; i++ {
				currentBlock := c.GetBlock(i, j, k)

				// Skip empty block
				if currentBlock == BlockEmpty {
					continue
				} else if currentBlock == BlockWater {
					// Water
					generateQuad(&c.Water, SideBottom, i, j, k)
					continue
				}

				// All other blocks
				if c.isTransparent(i, j, k-1) {
					generateQuad(&c.Terrain, SideNorth, i, j, k)
				}
				if c.isTransparent(i, j, k+1) {
					generateQuad(&c.Terrain, SideSouth, i, j, k)
				}
				if c.isTransparent(i+1, j, k) {
					generateQuad(&c.Terrain, SideEast, i, j, k)
				}
				if c.isTransparent(i-1, j, k) {
					generateQuad(&c.Terrain, SideWest, i, j, k)
				}
				if c.isTransparent(i, j+1, k) {
					generateQuad(&c.Terrain, SideTop, i, j, k)
				}
				if c.isTransparent(i, j-1, k) {
					generateQuad(&c.Terrain, SideBottom, i, j, k)
				}
			}
		}
	}

	c.Terrain.Upload()
	c.Water.Upload()
}
