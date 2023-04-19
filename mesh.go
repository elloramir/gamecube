package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Mesh struct {
	vaoId, vboId, eboId uint32
	vertexCount int32
}

func (m *Mesh) upload(vertices []float32, indices []uint32) {
	m.vertexCount = int32(len(indices))

	// NOTE: since our buffers are static, we need to recreate
	// them every time this function is called
	gl.GenVertexArrays(1, &m.vaoId)
	gl.GenBuffers(1, &m.vboId)
	gl.GenBuffers(1, &m.eboId)

	gl.BindVertexArray(m.vaoId) // bind

	// vbo data
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vboId)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// attribute layout
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 0, 0)

	// ebo data
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.eboId)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindVertexArray(0) // unbind
}

func (m *Mesh) unload() {
	m.vertexCount = 0 // reset count

	gl.DeleteVertexArrays(1, &m.vaoId)
	gl.DeleteBuffers(1, &m.vboId)
	gl.DeleteBuffers(1, &m.eboId)
}

func (m *Mesh) render() {
	gl.BindVertexArray(m.vaoId)
	gl.DrawElements(gl.TRIANGLES, m.vertexCount, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}