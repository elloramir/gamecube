package main

import (
	"fmt"
	_"embed"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	screenWidth = 800
	screenHeight = 600
)

//go:embed res/game.vert
var vertexSource string
//go:embed res/game.frag
var fragmentSource string

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// create window
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(screenWidth, screenHeight, "", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	// load opengl
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	fmt.Println("welcome adventurer, to the 'gamecube'")
	fmt.Println(gl.GoStr(gl.GetString(gl.VERSION)))

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}