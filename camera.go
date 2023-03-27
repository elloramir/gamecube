package main

import (
	"github.com/chewxy/math32"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	Near = 0.001
	Far = 1000
	Fov = 45
	Sensitivity = 1.0/30.0
)

type Camera struct {
	projectionMatrix mgl32.Mat4	
	viewMatrix mgl32.Mat4
	yaw, pitch float32
	position, front, up mgl32.Vec3
}

func (c *Camera) Init(aspect float32) {
	c.yaw = -90
	c.up = mgl32.Vec3{0, 1, 0}
	c.projectionMatrix = mgl32.Perspective(mgl32.DegToRad(Fov), aspect, Near, Far)
}

func (c *Camera) Update() {
	// keyboard movement
	if IsKeyDown(glfw.KeyW) {
		c.position = c.position.Add(c.front)
	}
	if IsKeyDown(glfw.KeyS) {
		c.position = c.position.Sub(c.front)
	}
	if IsKeyDown(glfw.KeyA) {
		c.position = c.position.Sub(c.front.Cross(c.up).Normalize())
	}
	if IsKeyDown(glfw.KeyD) {
		c.position = c.position.Add(c.front.Cross(c.up).Normalize())
	}
	if IsKeyDown(glfw.KeySpace) {
		c.position = c.position.Add(c.up)
	}
	if IsKeyDown(glfw.KeyLeftShift) {
		c.position = c.position.Sub(c.up)
	}

	// mouse looking
	mouseDeltaX, mouseDeltaY := MouseRelativePosition()
	c.yaw += mouseDeltaX * Sensitivity;
	c.pitch = math32.Max(math32.Min(c.pitch - mouseDeltaY * Sensitivity, 89.0), -89.0)

	// update matrices
	target := mgl32.Vec3{
		math32.Cos(mgl32.DegToRad(c.yaw)) * math32.Cos(mgl32.DegToRad(c.pitch)),
		math32.Sin(mgl32.DegToRad(c.pitch)),
		math32.Sin(mgl32.DegToRad(c.yaw)) * math32.Cos(mgl32.DegToRad(c.pitch))}

	c.front = target.Normalize()
	c.viewMatrix = mgl32.LookAtV(c.position, c.position.Add(c.front), mgl32.Vec3{0, 1, 0})
}

func (c *Camera) SendUniforms(program uint32) {
	viewProjLocation := GetUniform(program, "viewProjMatrix")
	viewProj := c.projectionMatrix.Mul4(c.viewMatrix)

	gl.UniformMatrix4fv(viewProjLocation, 1, false, &viewProj[0])
}