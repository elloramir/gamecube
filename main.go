// Copyright (c) 2023 Ellora.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// This file implements the initialization/entry of the Voxel Engine.

package main

import (
	"fmt"
	"github.com/elloramir/gamecube/game"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"runtime"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See GLFW (go-gl) documentation for functions that are only allowed to
	// be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize glow:", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.DepthFunc(gl.LESS)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0.1, 0.2, 0.3, 1.0)

	// Playground
	chunk := game.NewChunk(0, 0)
	camera := game.NewCamera()
	program, err := game.LoadShader("shaders/game.vert", "shaders/game.frag")
	if err != nil {
		log.Fatalln(err)
	}
	wprogram, err := game.LoadShader("shaders/game.vert", "shaders/water.frag")
	if err != nil {
		log.Fatalln(err)
	}

	for !window.ShouldClose() {
		// Render
		gl.Viewport(0, 0, windowWidth, windowHeight)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if chunk.Terrain != nil {
			camera.SendUniforms(program)
			gl.UseProgram(program)
			chunk.Terrain.Render()
		}
		if chunk.Water != nil {
			camera.SendUniforms(wprogram)
			gl.UseProgram(wprogram)
			chunk.Water.Render()
		}

		// Maintenance
		camera.Update()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
