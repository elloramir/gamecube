package main

import (
	"log"
	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	screenWidth = 800
	screenHeight = 600
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatal("Can't initialize GLFW")
	}
	defer glfw.Terminate()

	// create window
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(screenWidth, screenHeight, "", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.SetKeyCallback(KeyboardCallback)
	window.SetCursorPosCallback(MouseCallback)

	glfw.SwapInterval(1) // vsync on

	// loading opengl
	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}

	// program startup
	log.Println("welcome adventurer, to the 'gamecube'")
	log.Println(gl.GoStr(gl.GetString(gl.VERSION)))

	for !window.ShouldClose() {
		glfw.PollEvents()
		InputUpdate()

		if IsKeyDown(glfw.KeyEscape) {
			break
		}

		window.SwapBuffers()
	}
}