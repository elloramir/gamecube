package main

import (
	"fmt"
	"runtime"
	_"embed"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func init() {
	runtime.LockOSThread()
}

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

	// window callbacks
	window.SetKeyCallback(KeyboardCallback)
	window.SetCursorPosCallback(MouseCallback)

	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled);
	glfw.SwapInterval(1)

	// load opengl
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	fmt.Println("welcome adventurer, to the 'gamecube'")
	fmt.Println(gl.GoStr(gl.GetString(gl.VERSION)))

	// create chunk
	for z := int32(-4); z < 4; z++ {
		for x := int32(-4); x < 4; x++ {
			CreateChunk(x, z)
		}
	}
	defer NukeChunks()

	// load default shader
	program, err := CreateProgram(vertexSource, fragmentSource)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteProgram(program)

	gl.Enable(gl.DEPTH_TEST)
	gl.UseProgram(program)

	camera := Camera{}
	camera.Init(float32(screenWidth)/float32(screenHeight))

	for !window.ShouldClose() {
		// update
		camera.Update()
		camera.SendUniforms(program)

		if IsKeyDown(glfw.KeyEscape) {
			break
		}

		// render
		gl.Viewport(0, 0, int32(screenWidth), int32(screenHeight))
		gl.ClearColor(0.1, 0.2, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		RenderChunks()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}