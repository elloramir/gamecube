// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package game

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	Near = 0.001
	Far  = 1000
	Fov  = 45
)

type Camera struct {
	projMat  mgl32.Mat4
	viewMat  mgl32.Mat4
	transMat mgl32.Mat4
}

// Util vectors (unlike, not provide by mgl32)
var (
	zeroVec3 = mgl32.Vec3{0, 0, 0}
	upVec3   = mgl32.Vec3{0, 1, 0}
)

func NewCamera() *Camera {
	c := &Camera{}
	c.Update() // First interation here!

	return c
}

func (c *Camera) Update() {
	c.projMat = mgl32.Perspective(mgl32.DegToRad(Fov), float32(800)/600, Near, Far)
	c.viewMat = mgl32.LookAtV(mgl32.Vec3{10, 10, 20}, zeroVec3, upVec3)
	c.transMat = c.projMat.Mul4(c.viewMat)
}

// NOTE: I don't mind sending uniforms as part of the shaders code, but I resolve
// put it directly here in the camera API
func (c *Camera) SendUniforms(program uint32) {
	loc := gl.GetUniformLocation(program, gl.Str("transform\x00"))

	gl.UseProgram(program)
	gl.UniformMatrix4fv(loc, 1, false, &c.transMat[0])
	gl.UseProgram(0)
}
