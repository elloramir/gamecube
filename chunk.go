package main

import (
	simplex "github.com/ojrac/opensimplex-go"
)

const (
	ChunkWidth = 16
	ChunkHeight = 16
	ChunkLength = 16
)

type Chunk struct {
	blocks [ChunkWidth][ChunkHeight][ChunkLength]uint8
	mesh Mesh
}

var noise32 = simplex.New32(0)

func (c *Chunk) GenerateTerrain() {
	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkHeight; z++ {
			value := (noise32.Eval2(float32(x)/20.0, float32(z)/20.0) + 1) * 0.5
			height := int32(value * ChunkHeight)

			for height >= 0 {
				c.blocks[x][height][z] = BlockGrass
				height -= 1;
			}
		}
	}
}

// TODO: neighbour chunks in case of out bounds
func (c *Chunk) GetBlock(x, y, z float32) (uint8) {
	if x < 0 || x >= ChunkWidth {
		return BlockVoid
	}
	if y < 0 || y >= ChunkHeight {
		return BlockVoid
	}
	if z < 0 || z >= ChunkLength {
		return BlockVoid
	}

	return c.blocks[uint32(x)][uint32(y)][uint32(z)]
}

func (c *Chunk) Update() {
	var vertices []float32

	for z := float32(0); z < ChunkLength; z++ {
		for y := float32(0); y < ChunkHeight; y++ {
			for x := float32(0); x < ChunkWidth; x++ {
				if c.GetBlock(x, y, z) != BlockAir {
					if c.GetBlock(x, y, z-1) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideNorth, x, y, z)...)
					}
					if c.GetBlock(x, y, z+1) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideSouth, x, y, z)...)
					}
					if c.GetBlock(x+1, y, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideEast, x, y, z)...)
					}
					if c.GetBlock(x-1, y, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideWest, x, y, z)...)
					}
					if c.GetBlock(x, y+1, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideTop, x, y, z)...)
					}
					if c.GetBlock(x, y-1, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideBottom, x, y, z)...)
					}
				}				
			}
		}
	}

	c.mesh.Unload()
	c.mesh.Upload(vertices)
}

func (c *Chunk) Nuke() {
	c.mesh.Unload()
}