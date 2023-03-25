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
	x, z int32
	hash uint64
	blocks [ChunkWidth][ChunkHeight][ChunkLength]uint8
	mesh Mesh
}

var chunksMap = make(map[uint64]Chunk)
var noise32 = simplex.New32(0)

// 16 bits for each coordinates
// side effect: max and min values [-32768, 32767] 
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
			value := (noise32.Eval2(float32(offsetX+x)/20.0, float32(offsetZ+z)/20.0) + 1) * 0.5
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

	offsetX := float32(c.x * ChunkWidth)
	offsetZ := float32(c.z * ChunkLength) 

	for z := float32(0); z < ChunkLength; z++ {
		for y := float32(0); y < ChunkHeight; y++ {
			for x := float32(0); x < ChunkWidth; x++ {
				if c.GetBlock(x, y, z) != BlockAir {
					if c.GetBlock(x, y, z-1) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideNorth, offsetX + x, y, offsetZ + z)...)
					}
					if c.GetBlock(x, y, z+1) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideSouth, offsetX + x, y, offsetZ + z)...)
					}
					if c.GetBlock(x+1, y, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideEast, offsetX + x, y, offsetZ + z)...)
					}
					if c.GetBlock(x-1, y, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideWest, offsetX + x, y, offsetZ + z)...)
					}
					if c.GetBlock(x, y+1, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideTop, offsetX + x, y, offsetZ + z)...)
					}
					if c.GetBlock(x, y-1, z) == BlockAir {
						vertices = append(vertices, BlockGenFace(SideBottom, offsetX + x, y, offsetZ + z)...)
					}
				}				
			}
		}
	}

	c.mesh.Unload()
	c.mesh.Upload(vertices)
}