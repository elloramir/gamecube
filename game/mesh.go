// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// This file defines the Mesh struct used for rendering and building meshes
// in the Voxel Engine.

package game

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Mesh struct {
	vaoId, vboId, eboId     uint32
	vertexCount, indexCount uint32
}

func (m *Mesh) Upload(vertices []float32, indices []uint32) {
	m.Unload()

	m.vertexCount = uint32(len(vertices) / 3)
	m.indexCount = uint32(len(indices))

	// Generate OpenGL objects (again)
	gl.GenVertexArrays(1, &m.vaoId)
	gl.GenBuffers(1, &m.vboId)
	gl.GenBuffers(1, &m.eboId)

	gl.BindVertexArray(m.vaoId) // Bind VAO

	// Buffer data
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vboId)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Data layout
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 0, 0)

	// Indices data
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.eboId)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindVertexArray(0) // Unbind VAO
}

func (m *Mesh) Unload() {
	m.vertexCount = 0
	m.indexCount = 0

	gl.DeleteVertexArrays(1, &m.vaoId)
	gl.DeleteBuffers(1, &m.vboId)
	gl.DeleteBuffers(1, &m.eboId)
}

func (m *Mesh) Render() {
	// NOTE: I don't know exactly if the EBO must be binded, but yeah, there is!
	gl.BindVertexArray(m.vaoId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.eboId)

	gl.DrawElements(gl.TRIANGLES, int32(m.indexCount), gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}
