// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// This file provides a helper for generating vertex data easily
// for use in mesh rendering. The Vertex struct contains data and
// index buffers used to store the vertex attributes of a mesh.
// The BakeQuad method simplifies the process of adding four vertices
// that define a quad to the vertex buffer.

package game

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Vertices struct {
	Data    []float32
	Indices []uint32
}

func (v *Vertices) BakeQuad(a, b, c, d mgl32.Vec3) {
	i := uint32(len(v.Data) / 3)

	v.Data = append(v.Data,
		a.X(), a.Y(), a.Z(),
		b.X(), b.Y(), b.Z(),
		c.X(), c.Y(), c.Z(),
		d.X(), d.Y(), d.Z())

	v.Indices = append(v.Indices,
		i+0, i+1, i+2,
		i+0, i+2, i+3)
}

func (v *Vertices) ToMesh() *Mesh {
	if len(v.Data) == 0 {
		return nil
	}

	m := &Mesh{}
	m.Upload(v.Data, v.Indices)

	return m
}
