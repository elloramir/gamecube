// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package gfx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Mesh struct {
	vaoId, vboId, eboId     uint32
	vertexCount, indexCount uint32

	vertices []float32
	indices  []uint32
}

func (m *Mesh) Vertex(x, y, z, n1, n2, n3, u, v float32) {
	m.vertices = append(m.vertices, x, y, z, n1, n2, n3, u, v)
	m.vertexCount += 1
}

// NOTE: Quad function will use the last 4 vertices as reference to begin
func (m *Mesh) Quad(a, b, c, d, e, f uint32) {
	var i uint32 = m.vertexCount - 4
	m.indices = append(m.indices, i+a, i+b, i+c, i+d, i+e, i+f)
	m.indexCount += 6
}

func (m *Mesh) vanish() {
	m.vertices = nil
	m.indices = nil
}

func (m *Mesh) Upload() {
	// Generate OpenGL objects (again)
	gl.GenVertexArrays(1, &m.vaoId)
	gl.GenBuffers(1, &m.vboId)
	gl.GenBuffers(1, &m.eboId)

	gl.BindVertexArray(m.vaoId) // Bind VAO

	// Buffer data
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vboId)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.vertices)*4, gl.Ptr(m.vertices), gl.STATIC_DRAW)

	// Data layout
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 8*4, 0)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 8*4, 3*4)

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, 8*4, 6*4)

	// Indices data
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.eboId)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.indices)*4, gl.Ptr(m.indices), gl.STATIC_DRAW)

	gl.BindVertexArray(0) // Unbind VAO
	m.vanish()            // Voxels are pretty expensive, we don't need to keep track this data anymore
}

func (m *Mesh) Unload() {
	m.vertexCount = 0
	m.indexCount = 0

	gl.DeleteVertexArrays(1, &m.vaoId)
	gl.DeleteBuffers(1, &m.vboId)
	gl.DeleteBuffers(1, &m.eboId)
	m.vanish()
}

func (m *Mesh) Render() {
	// NOTE: I don't know exactly if the EBO must be binded, but yeah, there is!
	gl.BindVertexArray(m.vaoId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.eboId)

	gl.DrawElements(gl.TRIANGLES, int32(m.indexCount), gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}
