package main

import (
	simplex "github.com/ojrac/opensimplex-go"
)

const (
	ChunkWidth = 16
	ChunkHeight = 16
	ChunkLength = 16
	ChunkSmooth = 20.0
)

type Chunk struct {
	x, z int32
	hash uint64
	blocks [ChunkWidth][ChunkHeight][ChunkLength]uint8
	mesh Mesh
}

var chunksMap = make(map[uint64]Chunk)
var noise32 = simplex.New32(0)

// max and min values [-2147483648, 2147483647] 
func GenerateId(x, z int32) uint64 {
	return uint64(x)<<32 | uint64(z);
}

func DecodeId(id uint64) (int32, int32) {
	return int32(id)>>32, int32(id)
}

func CreateChunk(x, z int32) {
	chunk := Chunk{}
	chunk.x, chunk.z = x, z
	chunk.hash = GenerateId(x, z)
	chunk.GenerateTerrain()

	chunksMap[chunk.hash] = chunk
}

func RenderChunks() {
	for _, chunk := range chunksMap {
		chunk.mesh.Render()
	}
}

func NukeChunks() {
	for key, chunk := range chunksMap {
		chunk.mesh.Unload()
		delete(chunksMap, key)
	}
}

func (c *Chunk) GenerateTerrain() {
	offsetX := int(c.x * ChunkWidth)
	offsetZ := int(c.z * ChunkLength)

	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkHeight; z++ {
			noiseX := float32(offsetX+x)/ChunkSmooth
			noiseY := float32(offsetZ+z)/ChunkSmooth

			// norm from [0, 1] to [-1, 1]
			value := (noise32.Eval2(noiseX, noiseY) + 1) * 0.5
			height := int32(value * ChunkHeight)

			for height >= 0 {
				c.blocks[x][height][z] = BlockGrass
				height -= 1;
			}
		}
	}

	c.Update()
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
				absoluteX := x + float32(c.x * ChunkWidth)
				absoluteZ := z + float32(c.z * ChunkLength) 

				if c.GetBlock(x, y, z) != BlockAir {
					if c.GetBlock(x, y, z-1) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideNorth, absoluteX, y, absoluteZ)...)
					}
					if c.GetBlock(x, y, z+1) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideSouth, absoluteX, y, absoluteZ)...)
					}
					if c.GetBlock(x+1, y, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideEast, absoluteX, y, absoluteZ)...)
					}
					if c.GetBlock(x-1, y, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideWest, absoluteX, y, absoluteZ)...)
					}
					if c.GetBlock(x, y+1, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideTop, absoluteX, y, absoluteZ)...)
					}
					if c.GetBlock(x, y-1, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideBottom, absoluteX, y, absoluteZ)...)
					}
				}				
			}
		}
	}

	c.mesh.Unload()
	c.mesh.Upload(vertices)
}