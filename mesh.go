package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Mesh struct {
	vaoId, vboId uint32
	vertexCount uint32
}

func (m *Mesh) Upload(vertices []float32) {
	gl.GenVertexArrays(1, &m.vaoId)
	gl.GenBuffers(1, &m.vboId)

	gl.BindVertexArray(m.vaoId)

	// buffer data
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vboId)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// attribute layout
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)

	gl.BindVertexArray(0)

	// TODO: this '5' coming out of nowhere?? ;-;
	m.vertexCount = uint32(len(vertices) / 5)
}

func (m *Mesh) Unload() {
	gl.DeleteVertexArrays(1, &m.vaoId)
	gl.DeleteBuffers(1, &m.vboId)
}

func (m *Mesh) Render() {
	gl.BindVertexArray(m.vaoId)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(m.vertexCount))
	gl.BindVertexArray(0)
}