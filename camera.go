package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	Near = 0.001
	Far = 1000
	Fov = 45
	Sensitivity = 1.0/20.0
)

type Camera struct {
	projectionMatrix mgl32.Mat4	
	viewMatrix mgl32.Mat4
}

func newCamera(x, y, z float32) (*Camera) {
	c := new(Camera)
	c.viewMatrix = mgl32.LookAtV(mgl32.Vec3{x,y,z}, mgl32.Vec3{0,0,0}, mgl32.Vec3{0,1,0})
	c.projectionMatrix = mgl32.Perspective(mgl32.DegToRad(Fov), float32(800)/600, Near, Far)

	return c
}

// NOTE: I don't mind sending uniforms as part of the shaders code, but I resolve
// put it directly here in the camera API
func (c *Camera) sendUniforms(program uint32) {
	transformLoc := gl.GetUniformLocation(program, gl.Str("transformMatrix\x00"))
	transform := c.projectionMatrix.Mul4(c.viewMatrix)

	gl.UniformMatrix4fv(transformLoc, 1, false, &transform[0])
}