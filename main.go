package main

import (
	"fmt"
	"runtime"
	_"embed"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
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

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	// load opengl
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	fmt.Println("welcome adventurer, to the 'gamecube'")
	fmt.Println(gl.GoStr(gl.GetString(gl.VERSION)))

	// load default shader
	program, err := CreateProgram(vertexSource, fragmentSource)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteProgram(program)

	// create chunk
	CreateChunk(0, 0)
	CreateChunk(-1, 0)
	CreateChunk(-1, 1)
	defer NukeChunks()

	// matrices
	aspect := float32(screenWidth)/float32(screenHeight)
	projectionMatrix := mgl32.Perspective(mgl32.DegToRad(45), aspect, 0.001, 1000)
	viewMatrix := mgl32.LookAtV(mgl32.Vec3{3, 20, -20}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	modelMatrix := mgl32.Ident4()

	projectionLocation := GetUniform(program, "projectionUniform")
	viewLocation := GetUniform(program, "viewUniform")
	modelLocation := GetUniform(program, "modelUniform")

	gl.Enable(gl.DEPTH_TEST)
	gl.UseProgram(program)
	gl.UniformMatrix4fv(projectionLocation, 1, false, &projectionMatrix[0])
	gl.UniformMatrix4fv(viewLocation, 1, false, &viewMatrix[0])

	// time control
	angle := 0.0
	previousTime := glfw.GetTime()

	for !window.ShouldClose() {
		// frame time
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		// update
		angle += elapsed
		modelMatrix = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})
		gl.UniformMatrix4fv(modelLocation, 1, false, &modelMatrix[0])

		// render
		gl.Viewport(0, 0, int32(screenWidth), int32(screenHeight))
		gl.ClearColor(0.1, 0.2, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		RenderChunks()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}